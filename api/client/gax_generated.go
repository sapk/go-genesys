// Copyright © 2018 Antoine GIRARD <antoine.girard@sapk.fr>
package client

//Generated file DO NOT EDIT

import (
	"github.com/mitchellh/mapstructure"

	"github.com/sapk/go-genesys/api/object"
)

//ListSwitch list all Switch
func (c *Client) ListSwitch() ([]object.CfgSwitch, error) {
	var apps []object.CfgSwitch
	_, err := c.ListObject("CfgSwitch", &apps)
	return apps, err
}

//GetSwitchByID retrieve a specific Switch by id
func (c *Client) GetSwitchByID(id string) (*object.CfgSwitch, error) {
	var obj object.CfgSwitch
	_, err := c.GetObjectByID("CfgSwitch", id, &obj)
	return &obj, err
}

//GetSwitchByName retrieve a specific Switch by name
func (c *Client) GetSwitchByName(name string) (*object.CfgSwitch, error) {
	var obj object.CfgSwitch
	o, _, err := c.GetObjectByName("CfgSwitch", name)
	if err != nil {
		return nil, err
	}
	err = mapstructure.Decode(o, &obj) //TODO find a better way because mapstructure can use reflect under the hood
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

//ListDN list all DN
func (c *Client) ListDN() ([]object.CfgDN, error) {
	var apps []object.CfgDN
	_, err := c.ListObject("CfgDN", &apps)
	return apps, err
}

//GetDNByID retrieve a specific DN by id
func (c *Client) GetDNByID(id string) (*object.CfgDN, error) {
	var obj object.CfgDN
	_, err := c.GetObjectByID("CfgDN", id, &obj)
	return &obj, err
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

//GetPersonByID retrieve a specific Person by id
func (c *Client) GetPersonByID(id string) (*object.CfgPerson, error) {
	var obj object.CfgPerson
	_, err := c.GetObjectByID("CfgPerson", id, &obj)
	return &obj, err
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

//ListPlace list all Place
func (c *Client) ListPlace() ([]object.CfgPlace, error) {
	var apps []object.CfgPlace
	_, err := c.ListObject("CfgPlace", &apps)
	return apps, err
}

//GetPlaceByID retrieve a specific Place by id
func (c *Client) GetPlaceByID(id string) (*object.CfgPlace, error) {
	var obj object.CfgPlace
	_, err := c.GetObjectByID("CfgPlace", id, &obj)
	return &obj, err
}

//GetPlaceByName retrieve a specific Place by name
func (c *Client) GetPlaceByName(name string) (*object.CfgPlace, error) {
	var obj object.CfgPlace
	o, _, err := c.GetObjectByName("CfgPlace", name)
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

//GetAgentGroupByID retrieve a specific Agent Group by id
func (c *Client) GetAgentGroupByID(id string) (*object.CfgAgentGroup, error) {
	var obj object.CfgAgentGroup
	_, err := c.GetObjectByID("CfgAgentGroup", id, &obj)
	return &obj, err
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

//GetApplicationByID retrieve a specific Application by id
func (c *Client) GetApplicationByID(id string) (*object.CfgApplication, error) {
	var obj object.CfgApplication
	_, err := c.GetObjectByID("CfgApplication", id, &obj)
	return &obj, err
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

//GetHostByID retrieve a specific Host by id
func (c *Client) GetHostByID(id string) (*object.CfgHost, error) {
	var obj object.CfgHost
	_, err := c.GetObjectByID("CfgHost", id, &obj)
	return &obj, err
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

//ListAgentLogin list all Agent Login
func (c *Client) ListAgentLogin() ([]object.CfgAgentLogin, error) {
	var apps []object.CfgAgentLogin
	_, err := c.ListObject("CfgAgentLogin", &apps)
	return apps, err
}

//GetAgentLoginByID retrieve a specific Agent Login by id
func (c *Client) GetAgentLoginByID(id string) (*object.CfgAgentLogin, error) {
	var obj object.CfgAgentLogin
	_, err := c.GetObjectByID("CfgAgentLogin", id, &obj)
	return &obj, err
}

//GetAgentLoginByName retrieve a specific Agent Login by name
func (c *Client) GetAgentLoginByName(name string) (*object.CfgAgentLogin, error) {
	var obj object.CfgAgentLogin
	o, _, err := c.GetObjectByProp("CfgAgentLogin", "logincode", name)
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

//GetDNGroupByID retrieve a specific DN Group by id
func (c *Client) GetDNGroupByID(id string) (*object.CfgDNGroup, error) {
	var obj object.CfgDNGroup
	_, err := c.GetObjectByID("CfgDNGroup", id, &obj)
	return &obj, err
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

//ListAppPrototype list all Application Template
func (c *Client) ListAppPrototype() ([]object.CfgAppPrototype, error) {
	var apps []object.CfgAppPrototype
	_, err := c.ListObject("CfgAppPrototype", &apps)
	return apps, err
}

//GetAppPrototypeByID retrieve a specific Application Template by id
func (c *Client) GetAppPrototypeByID(id string) (*object.CfgAppPrototype, error) {
	var obj object.CfgAppPrototype
	_, err := c.GetObjectByID("CfgAppPrototype", id, &obj)
	return &obj, err
}

//GetAppPrototypeByName retrieve a specific Application Template by name
func (c *Client) GetAppPrototypeByName(name string) (*object.CfgAppPrototype, error) {
	var obj object.CfgAppPrototype
	o, _, err := c.GetObjectByName("CfgAppPrototype", name)
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

//GetAccessGroupByID retrieve a specific Access Group by id
func (c *Client) GetAccessGroupByID(id string) (*object.CfgAccessGroup, error) {
	var obj object.CfgAccessGroup
	_, err := c.GetObjectByID("CfgAccessGroup", id, &obj)
	return &obj, err
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

//ListFolder list all Folder
func (c *Client) ListFolder() ([]object.CfgFolder, error) {
	var apps []object.CfgFolder
	_, err := c.ListObject("CfgFolder", &apps)
	return apps, err
}

//GetFolderByID retrieve a specific Folder by id
func (c *Client) GetFolderByID(id string) (*object.CfgFolder, error) {
	var obj object.CfgFolder
	_, err := c.GetObjectByID("CfgFolder", id, &obj)
	return &obj, err
}

//GetFolderByName retrieve a specific Folder by name
func (c *Client) GetFolderByName(name string) (*object.CfgFolder, error) {
	var obj object.CfgFolder
	o, _, err := c.GetObjectByName("CfgFolder", name)
	if err != nil {
		return nil, err
	}
	err = mapstructure.Decode(o, &obj) //TODO find a better way because mapstructure can use reflect under the hood
	if err != nil {
		return nil, err
	}
	return &obj, nil
}
