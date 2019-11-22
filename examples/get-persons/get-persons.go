package main

import (
	"fmt"
	"os"

	"github.com/sapk/go-genesys/api/client"
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
	//c.Encoder = charmap.Windows1252.NewEncoder()
	_, err := c.Login(user, pass)
	if err != nil {
		fmt.Printf("Failed to login as %s\n", user)
		os.Exit(2)
	}

	personList, err := c.ListPerson()
	if err != nil {
		fmt.Printf("Failed to get ListPerson: %v\n", err)
		os.Exit(3)
	}
	for _, p := range personList {
		fmt.Printf("%s %s\n", p.Firstname, p.Lastname)
		//fmt.Printf("bytes: %#v\n", toByte(p.Firstname))
		//fmt.Printf("test1: %#v\n", string(toByte(p.Firstname)))
		//fmt.Printf("test2: %#v\n", bytes.NewBuffer(toByte(p.Firstname)).String())
	}

	//tBytes := []byte{0x5b, 0x7b, 0x22, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3a, 0x22, 0x46, 0x72, 0x61, 0x6e, 0xc3, 0xa7, 0x6f, 0x69, 0x73, 0x22, 0x7d, 0x5d}
	//fmt.Printf("test1: %#v\n", string(tBytes))
}

/*
func toByte(s string) []byte {
	return []byte(s)
}
*/
