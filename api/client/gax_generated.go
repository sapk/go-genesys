package client

// Copyright © 2018 Antoine GIRARD <antoine.girard@sapk.fr>

//Generated file DO NOT EDIT

import (
	"github.com/mitchellh/mapstructure"

	"github.com/sapk/go-genesys/api/object"
)
	
//ListDN list all DN
func (c *Client) ListDN() ([]object.CfgDN, error) {
	var apps []object.CfgDN
	_, err := c.ListObject("CfgDN", &apps)
	return apps, err
}

//GetDNByName retrieve a specific DN by name
func (c *Client) GetDNByName(name string) (*object.CfgDN, error) {
	var obj object.CfgDN
	o, _, err := c.GetObjectByName("CfgDN", name)
	if err != nil {
		return nil, err
	}
	err = mapstructure.Decode(o, &obj) //TODO find a better way because mapstructure can use reflect under the hood
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

//ListPerson list all Person
func (c *Client) ListPerson() ([]object.CfgPerson, error) {
	var apps []object.CfgPerson
	_, err := c.ListObject("CfgPerson", &apps)
	return apps, err
}

//GetPersonByName retrieve a specific Person by name
func (c *Client) GetPersonByName(name string) (*object.CfgPerson, error) {
	var obj object.CfgPerson
	o, _, err := c.GetObjectByName("CfgPerson", name)
	if err != nil {
		return nil, err
	}
	err = mapstructure.Decode(o, &obj) //TODO find a better way because mapstructure can use reflect under the hood
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

//ListAgentGroup list all Agent Group
func (c *Client) ListAgentGroup() ([]object.CfgAgentGroup, error) {
	var apps []object.CfgAgentGroup
	_, err := c.ListObject("CfgAgentGroup", &apps)
	return apps, err
}

//GetAgentGroupByName retrieve a specific Agent Group by name
func (c *Client) GetAgentGroupByName(name string) (*object.CfgAgentGroup, error) {
	var obj object.CfgAgentGroup
	o, _, err := c.GetObjectByName("CfgAgentGroup", name)
	if err != nil {
		return nil, err
	}
	err = mapstructure.Decode(o, &obj) //TODO find a better way because mapstructure can use reflect under the hood
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

//ListApplication list all Application
func (c *Client) ListApplication() ([]object.CfgApplication, error) {
	var apps []object.CfgApplication
	_, err := c.ListObject("CfgApplication", &apps)
	return apps, err
}

//GetApplicationByName retrieve a specific Application by name
func (c *Client) GetApplicationByName(name string) (*object.CfgApplication, error) {
	var obj object.CfgApplication
	o, _, err := c.GetObjectByName("CfgApplication", name)
	if err != nil {
		return nil, err
	}
	err = mapstructure.Decode(o, &obj) //TODO find a better way because mapstructure can use reflect under the hood
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

//ListHost list all Host
func (c *Client) ListHost() ([]object.CfgHost, error) {
	var apps []object.CfgHost
	_, err := c.ListObject("CfgHost", &apps)
	return apps, err
}

//GetHostByName retrieve a specific Host by name
func (c *Client) GetHostByName(name string) (*object.CfgHost, error) {
	var obj object.CfgHost
	o, _, err := c.GetObjectByName("CfgHost", name)
	if err != nil {
		return nil, err
	}
	err = mapstructure.Decode(o, &obj) //TODO find a better way because mapstructure can use reflect under the hood
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

//ListDNGroup list all DN Group
func (c *Client) ListDNGroup() ([]object.CfgDNGroup, error) {
	var apps []object.CfgDNGroup
	_, err := c.ListObject("CfgDNGroup", &apps)
	return apps, err
}

//GetDNGroupByName retrieve a specific DN Group by name
func (c *Client) GetDNGroupByName(name string) (*object.CfgDNGroup, error) {
	var obj object.CfgDNGroup
	o, _, err := c.GetObjectByName("CfgDNGroup", name)
	if err != nil {
		return nil, err
	}
	err = mapstructure.Decode(o, &obj) //TODO find a better way because mapstructure can use reflect under the hood
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

//ListAccessGroup list all Access Group
func (c *Client) ListAccessGroup() ([]object.CfgAccessGroup, error) {
	var apps []object.CfgAccessGroup
	_, err := c.ListObject("CfgAccessGroup", &apps)
	return apps, err
}

//GetAccessGroupByName retrieve a specific Access Group by name
func (c *Client) GetAccessGroupByName(name string) (*object.CfgAccessGroup, error) {
	var obj object.CfgAccessGroup
	o, _, err := c.GetObjectByName("CfgAccessGroup", name)
	if err != nil {
		return nil, err
	}
	err = mapstructure.Decode(o, &obj) //TODO find a better way because mapstructure can use reflect under the hood
	if err != nil {
		return nil, err
	}
	return &obj, nil
}
