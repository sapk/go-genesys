package client

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"golang.org/x/text/encoding/charmap"
)

var HTTPServer *httptest.Server

func TestMain(m *testing.M) {
	// Start a local HTTP server
	HTTPServer = httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
		// Send response to be tested
		rw.Write([]byte(`{"status": "ok", "ansi": "Ã§"}`))
	}))
	// Close the server when test finishes
	defer HTTPServer.Close()

	//Run server
	exitVal := m.Run()

	os.Exit(exitVal)
}

func TestNoEncoding(t *testing.T) {
	client := NewClient(HTTPServer.Listener.Addr().String(), false)
	client.HTTPClient = HTTPServer.Client()
	req, err := client.newRequest("GET", "/string/utf-8", "{}")
	if err != nil {
		t.Errorf("Failed to create a request: %v", err)
	}
	result := make(map[string]string)
	resp, err := client.do(req, &result)
	if err != nil {
		t.Errorf("Failed to do a request: %v\n%v\n%v", err, result, resp)
	}
}

func TestAnsiToUTFEncoding(t *testing.T) {
	client := NewClient(HTTPServer.Listener.Addr().String(), false)
	client.HTTPClient = HTTPServer.Client()
	client.Encoder = charmap.Windows1252.NewEncoder()
	req, err := client.newRequest("GET", "/string/ansi", "{}")
	if err != nil {
		t.Errorf("Failed to create a request: %v", err)
	}
	result := make(map[string]string)
	resp, err := client.do(req, &result)
	if err != nil {
		t.Errorf("Failed to do a request: %v\n%v\n%v", err, result, resp)
		return
	}
	if r, ok := result["ansi"]; !ok || r != "ç" {
		t.Errorf("Failed to decode request:\n%v\n%v", result, resp)
	}
}
