package client

import (
	"os"
	"testing"
)

var (
	TestHost = os.Getenv("TEST_HOST")
	TestUser = os.Getenv("TEST_USER")
	TestPass = os.Getenv("TEST_PASS")
)

//To test run : GOCACHE=off TEST_HOST=172.18.0.4:8080 TEST_USER=user TEST_PASS=pass go test -v ./api/client/

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

func TestGetObjectByName(t *testing.T) {
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
