// Copyright © 2018 Antoine GIRARD <antoine.girard@sapk.fr>
package client

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"

	"golang.org/x/text/encoding"
)

//Client Genesys GAX Api client
//TODO add cache
//TODO store current user to known if logged
type Client struct {
	BaseURL    *url.URL
	UserAgent  string
	Decoder    *encoding.Decoder
	httpClient *http.Client
}

//NewClient generate a client to communicate with a specific instance
func NewClient(host string, useSSL bool) *Client {
	cookieJar, _ := cookiejar.New(nil)
	c := &Client{BaseURL: &url.URL{Host: host, Scheme: "http", Path: "/gax/api/"}, UserAgent: "go-genesys/0.0", httpClient: &http.Client{
		Jar: cookieJar,
	}}
	if useSSL {
		c.BaseURL.Scheme = "https"
	}
	return c
}

func (c *Client) newRequest(method, path string, body interface{}) (*http.Request, error) {
	rel := &url.URL{Path: path}
	u := c.BaseURL.ResolveReference(rel)

	var (
		err error
		req *http.Request
	)

	data, ok := body.(string)
	if ok { //Support string body to send direct
		req, err = http.NewRequest(method, u.String(), strings.NewReader(data))
	} else { //Or convert to json
		var buf io.ReadWriter
		if body != nil {
			buf = new(bytes.Buffer)
			e := json.NewEncoder(buf).Encode(body)
			if e != nil {
				return nil, err
			}
		}
		req, err = http.NewRequest(method, u.String(), buf)
	}

	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)

	return req, nil
}

func (c *Client) do(req *http.Request, v interface{}) (*http.Response, error) {
	/*
		logrus.WithFields(logrus.Fields{
			"Method":  req.Method,
			"Path":    req.URL.Path,
			"Query":   req.URL.RawQuery,
			"Cookies": req.Cookies(),
			"Body":    req.Body,
		}).Debug("Executing request")
	*/
	resp, err := c.httpClient.Do(req)
	if err != nil {
		/*
			logrus.WithFields(logrus.Fields{
				"Method": req.Method,
				"Path":   req.URL.Path,
				"Error":  err,
			}).Debug("Request failed")
		*/
		return nil, err
	}
	defer resp.Body.Close()

	//For debug
	//buf := new(bytes.Buffer)
	//buf.ReadFrom(resp.Body)
	/*
		logrus.WithFields(logrus.Fields{
			"Method": req.Method,
			"Path":   req.URL.Path,
			"Status": resp.Status,
			"Length": resp.ContentLength,
		}).Debug("Request response")
	*/
	if resp.StatusCode != 201 {
		var reader io.Reader
		if c.Decoder != nil {
			reader = c.Decoder.Reader(resp.Body)
		} else {
			reader = resp.Body
		}
		err = json.NewDecoder(reader).Decode(v)
	}
	return resp, err
}
