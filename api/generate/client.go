package main

// Copyright © 2018 Antoine GIRARD <antoine.girard@sapk.fr>

//Generate each function for types
//go generate ./api/client/
//Should be call from $GOPATH/src/github/sapk/go-genesys

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/sapk/go-genesys/api/object"
)

func main() {
	str := `package client

// Copyright © 2018 Antoine GIRARD <antoine.girard@sapk.fr>

//Generated file DO NOT EDIT

import (
	"github.com/mitchellh/mapstructure"

	"github.com/sapk/go-genesys/api/object"
)
	`
	for _, o := range object.TypeListDefined {
		fmt.Printf("Generate methods for: %s\n", o.Name)
		str += `
//List` + strings.TrimPrefix(o.Name, "Cfg") + ` list all ` + o.Desc + `
func (c *Client) List` + strings.TrimPrefix(o.Name, "Cfg") + `() ([]object.` + o.Name + `, error) {
	var apps []object.` + o.Name + `
	_, err := c.ListObject("` + o.Name + `", &apps)
	return apps, err
}

//Get` + strings.TrimPrefix(o.Name, "Cfg") + `ByName retrieve a specific ` + o.Desc + ` by name
func (c *Client) Get` + strings.TrimPrefix(o.Name, "Cfg") + `ByName(name string) (*object.` + o.Name + `, error) {
	var obj object.` + o.Name + `
	o, _, err := c.GetObjectByName("` + o.Name + `", name)
	if err != nil {
		return nil, err
	}
	err = mapstructure.Decode(o, &obj) //TODO find a better way because mapstructure can use reflect under the hood
	if err != nil {
		return nil, err
	}
	return &obj, nil
}
`
	}

	err := ioutil.WriteFile("gax_generated.go", []byte(str), 0644)
	if err != nil {
		log.Panicf("Failed to generate file: %s", "gax_generated.go")
	}
}
