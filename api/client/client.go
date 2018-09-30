package client

// Copyright © 2018 Antoine GIRARD <antoine.girard@sapk.fr>

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"

	"github.com/sapk/go-genesys/api/object"
	"github.com/sirupsen/logrus"
)

//Client Genesys GAX Api client
//TODO add cache
//TODO store current user to known if logged
type Client struct {
	BaseURL   *url.URL
	UserAgent string

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

//Login log the client on the GAX instance linked
func (c *Client) Login(user, pass string) (*object.LoginResponse, error) {
	req, err := c.newRequest("POST", "session/login", object.LoginRequest{Username: user, Password: pass, IsPasswordEncrypted: false})
	if err != nil {
		return nil, err
	}
	_, err = c.do(req, nil)

	//Check logged user
	req, err = c.newRequest("GET", "user/info", nil)
	if err != nil {
		return nil, err
	}
	var u object.LoginResponse
	_, err = c.do(req, &u)
	return &u, err
}

//UpdateObject Update a object. The object could be a json string or a go object
func (c *Client) UpdateObject(objType, objID string, v interface{}) (*http.Response, error) {
	req, err := c.newRequest("PUT", fmt.Sprintf("cfg/objects/%s/%s", objType, objID), v)
	if err != nil {
		return nil, err
	}
	return c.do(req, v)
}

//PostObject Create a object. The object could be a json string or a go object
func (c *Client) PostObject(v interface{}) (*http.Response, error) {
	req, err := c.newRequest("POST", "cfg/objects", v)
	if err != nil {
		return nil, err
	}
	return c.do(req, v)
}

//ListObject Return all the object of a specific type
func (c *Client) ListObject(t string, v interface{}) (*http.Response, error) {
	req, err := c.newRequest("GET", "cfg/objects", nil)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = "brief=false&type=" + t
	return c.do(req, v)
}

/*
func (c *Client) ListApplication() ([]object.CfgApplication, error) {
	var apps []object.CfgApplication
	_, err := c.ListObject("CfgApplication", &apps)
	return apps, err
}

func (c *Client) ListHost() ([]object.CfgHost, error) {
	var apps []object.CfgHost
	_, err := c.ListObject("CfgHost", &apps)
	return apps, err
}

func (c *Client) ListDN() ([]object.CfgDN, error) {
	var apps []object.CfgDN
	_, err := c.ListObject("CfgDN", &apps)
	return apps, err
}

func (c *Client) ListSwitch() ([]object.CfgSwitch, error) {
	var apps []object.CfgSwitch
	_, err := c.ListObject("CfgSwitch", &apps)
	return apps, err
}

func (c *Client) ListPlace() ([]object.CfgPlace, error) {
	var apps []object.CfgPlace
	_, err := c.ListObject("CfgPlace", &apps)
	return apps, err
}
*/

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
	logrus.WithFields(logrus.Fields{
		"Method":  req.Method,
		"Path":    req.URL.Path,
		"Query":   req.URL.RawQuery,
		"Cookies": req.Cookies(),
		"Body":    req.Body,
	}).Debug("Executing request")
	resp, err := c.httpClient.Do(req)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"Method": req.Method,
			"Path":   req.URL.Path,
			"Error":  err,
		}).Debug("Request failed")
		return nil, err
	}
	defer resp.Body.Close()

	//For debug
	//buf := new(bytes.Buffer)
	//buf.ReadFrom(resp.Body)
	logrus.WithFields(logrus.Fields{
		"Method": req.Method,
		"Path":   req.URL.Path,
		"Status": resp.Status,
		"Length": resp.ContentLength,
	}).Debug("Request response")
	if resp.StatusCode != 201 {
		err = json.NewDecoder(resp.Body).Decode(v)
	}
	return resp, err
}
