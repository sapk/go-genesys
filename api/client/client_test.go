package client

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var HTTPServer *httptest.Server

func TestMain(m *testing.M) {
	var paths = map[string][]byte{
		"/string/utf-8": []byte(`{"text": "ç"}`),
		//"/string/ansi":  []byte(`{"text": "Ã§"}`),
		"/string/ansi": []byte{0x7b, 0x22, 0x74, 0x65, 0x78, 0x74, 0x22, 0x3a, 0x20, 0x22, 0xc3, 0xa7, 0x22, 0x7d},
		//"/gax/api/cfg/objects": []byte(`[{"firstname":"François"}]`),
		"/gax/api/cfg/objects": []byte{0x5b, 0x7b, 0x22, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3a, 0x22, 0x46, 0x72, 0x61, 0x6e, 0xc3, 0xa7, 0x6f, 0x69, 0x73, 0x22, 0x7d, 0x5d},
		//[]byte(`[{"firstname":"FranÃ§ois"}]`),
		//[]byte{0x46, 0x72, 0x61, 0x6e,
		//0xc3, 0xa7,
		//0x6f, 0x69, 0x73}
	}
	// Start a local HTTP server
	HTTPServer = httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
		fmt.Printf("[%s]: %s\n", req.Method, req.URL.Path)
		if v, exist := paths[req.URL.Path]; exist {
			// Send response to be tested
			fmt.Printf("response: %s\n%#v\n", string(v), v)
			rw.Write(v)
		}
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
	result := make(map[string]string, 0)
	resp, err := client.do(req, &result)
	if err != nil {
		t.Errorf("Failed to do a request: %v\n%v\n%v", err, result, resp)
	}
}

func TestAnsiToUTFEncoding(t *testing.T) {
	client := NewClient(HTTPServer.Listener.Addr().String(), false)
	client.HTTPClient = HTTPServer.Client()
	//client.Decoder = charmap.Windows1252.NewDecoder()
	//client.Encoder = charmap.Windows1252.NewEncoder()
	req, err := client.newRequest("GET", "/string/ansi", "{}")
	if err != nil {
		t.Errorf("Failed to create a request: %v", err)
	}
	result := make(map[string]string, 0)
	resp, err := client.do(req, &result)
	if err != nil {
		t.Errorf("Failed to do a request: %v\n%v\n%v", err, result, resp)
		return
	}
	r, ok := result["text"]
	if !ok || r != "ç" {
		t.Errorf("Failed to decode request:\n%v\n%v", result, resp)
	}
	fmt.Printf("%#v\n", r)
	fmt.Printf("%#v\n", []byte(r))
}

/*
func TestListPersonNoEncoding(t *testing.T) {
	client := NewClient(HTTPServer.Listener.Addr().String(), false)
	client.HTTPClient = HTTPServer.Client()
	personList, err := client.ListPerson()
	if err != nil {
		t.Errorf("Failed to get ListPerson: %v\n", err)
	}
	if len(personList) != 1 {
		t.Errorf("Invalid number of person returned: %v\n", personList)
	}
	fmt.Printf("\n%v\n", personList)
}
*/

func TestListPersonAnsiToUTFEncoding(t *testing.T) {
	client := NewClient(HTTPServer.Listener.Addr().String(), false)
	client.HTTPClient = HTTPServer.Client()
	//client.Decoder = charmap.Windows1252.NewDecoder()
	//client.Encoder = charmap.Windows1252.NewEncoder()
	personList, err := client.ListPerson()
	if err != nil {
		t.Errorf("Failed to get ListPerson: %v\n", err)
	}
	if len(personList) != 1 {
		t.Errorf("Invalid number of person returned: %v\n", personList)
	}
	fmt.Printf("%#v\n", personList[0].Firstname)
	fmt.Printf("%#v\n", toByte(personList[0].Firstname))
	fmt.Printf("test1: %#v\n", string(toByte(personList[0].Firstname)))
}

func toByte(s string) []byte {
	return []byte(s)
}
