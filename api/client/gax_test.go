package client

import (
	"os"
	"testing"

	"github.com/sapk/go-genesys/api/object"
)

var (
	TestHost = os.Getenv("TEST_HOST")
	TestUser = os.Getenv("TEST_USER")
	TestPass = os.Getenv("TEST_PASS")
)

//To test run : GOCACHE=off TEST_HOST=host:8080 TEST_USER=user TEST_PASS=pass go test -v ./api/client/
//Tests highly depends on a env test that need to be setup

func TestLogin(t *testing.T) {
	c := NewClient(TestHost, false)
	u, err := c.Login(TestUser, TestPass)

	if err != nil {
		t.Errorf("Failed to login as %s: %v", TestUser, err)
	}
	if u.Username != TestUser {
		t.Errorf("Failed to login as %s: %v", TestUser, u)
	}
}

func TestGetApplicationByName(t *testing.T) {
	c := NewClient(TestHost, false)
	_, err := c.Login(TestUser, TestPass)
	if err != nil {
		t.Errorf("Failed to login as %s!", TestUser)
	}

	app, err := c.GetApplicationByName("GAX")
	if err != nil {
		t.Errorf("Failed to get application %s: %v", "GAX", err)
	}
	if app.Name != "GAX" {
		t.Errorf("Failed to get application %s: %v", "GAX", app)
	}
}

func TestGetObjectByID(t *testing.T) {
	c := NewClient(TestHost, false)
	_, err := c.Login(TestUser, TestPass)
	if err != nil {
		t.Errorf("Failed to login as %s!", TestUser)
	}

	var app object.CfgApplication
	_, err = c.GetObjectByID("CfgApplication", "99", &app)
	if err != nil {
		t.Errorf("Failed to get application %s: %v", "confserv", err)
	}
	if app.Name != "confserv" {
		t.Errorf("Failed to get application %s: %v", "confserv", app)
	}
}

func TestGetObjectByName(t *testing.T) {
	c := NewClient(TestHost, false)
	_, err := c.Login(TestUser, TestPass)
	if err != nil {
		t.Errorf("Failed to login as %s!", TestUser)
	}

	app, _, err := c.GetObjectByName("CfgApplication", "GAX")
	if err != nil {
		t.Errorf("Failed to get application %s: %v", "GAX", err)
	}
	if name, ok := app["name"]; !ok || name != "GAX" {
		t.Errorf("Failed to get application %s: %v", "GAX", app)
	}
}

//TODO get application list and run multiple GetObjectByID, GetObjectByName and GetApplicationByName
