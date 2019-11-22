package main

import (
	"fmt"
	"os"

	"github.com/sapk/go-genesys/api/client"
	"golang.org/x/text/encoding/charmap"
)

var (
	host = os.Getenv("GAX_HOST")
	user = os.Getenv("GAX_USER")
	pass = os.Getenv("GAX_PASS")
)

func main() {
	if host == "" || user == "" {
		fmt.Print("You need to set env variable GAX_HOST, GAX_USER and GAX_PASS\n")
		os.Exit(1)
	}
	c := client.NewClient(host, false)
	c.Encoder = charmap.Windows1252.NewEncoder()
	_, err := c.Login(user, pass)
	if err != nil {
		fmt.Printf("Failed to login as %s\n", user)
		os.Exit(2)
	}

	personList, err := c.ListPerson()
	if err != nil {
		fmt.Printf("Failed to get application %s: %v\n", "GAX", err)
		os.Exit(3)
	}
	for _, p := range personList {
		fmt.Printf("%s %s\n", p.Firstname, p.Lastname)
	}
}
