// Copyright © 2018 Antoine GIRARD <antoine.girard@sapk.fr>
package object

//Information mostly from https://docs.genesys.com/Documentation/PSDK/9.0.x/ConfigLayerRef/CfgType

//Type Type of object
type Type struct {
	ID         int
	Name       string
	Desc       string
	IsDumpable bool
}

//TypeListShort Contain most common object
var TypeListShort = []Type{
	Type{3, "CfgPerson", "Person", true},
	Type{4, "CfgPlace", "Place", true},
	Type{9, "CfgApplication", "Application", true},
	Type{10, "CfgHost", "Host", true},
}

//TypeListDefined Contain object that have a golang representation
var TypeListDefined = []Type{
	Type{2, "CfgDN", "DN", true},
	Type{3, "CfgPerson", "Person", true},
	Type{4, "CfgPlace", "Place", true},
	Type{5, "CfgAgentGroup", "Agent Group", true},
	Type{9, "CfgApplication", "Application", true},
	Type{10, "CfgHost", "Host", true},
	Type{15, "CfgAgentLogin", "Agent Login", true},
	Type{17, "CfgDNGroup", "DN Group", true},
	Type{20, "CfgAppPrototype", "Application Template", true},
	Type{21, "CfgAccessGroup", "Access Group", true},
	Type{22, "CfgFolder", "Folder", true},
}

//TypeList Contain almost all object
var TypeList = []Type{
	//Type{0, "CfgNoObject", "Unknown Object",true},
	Type{1, "CfgSwitch", "Switch", true},
	Type{2, "CfgDN", "DN", true},
	Type{3, "CfgPerson", "Person", true},
	Type{4, "CfgPlace", "Place", true},
	Type{5, "CfgAgentGroup", "Agent Group", true},
	Type{6, "CfgPlaceGroup", "Place Group", true},
	Type{7, "CfgTenant", "Tenant", true},
	Type{8, "CfgService", "Solution", true},
	Type{9, "CfgApplication", "Application", true},
	Type{10, "CfgHost", "Host", true},
	Type{11, "CfgPhysicalSwitch", "Switching Office", true},
	Type{12, "CfgScript", "Script", true},
	Type{13, "CfgSkill", "Skill", true},
	Type{14, "CfgActionCode", "Action Code", true},
	Type{15, "CfgAgentLogin", "Agent Login", true},
	Type{16, "CfgTransaction", "Transaction", true},
	Type{17, "CfgDNGroup", "DN Group", true},
	Type{18, "CfgStatDay", "Statistical Day", true},
	Type{19, "CfgStatTable", "Statistical Table", true},
	Type{20, "CfgAppPrototype", "Application Template", true},
	Type{21, "CfgAccessGroup", "Access Group", true},
	Type{22, "CfgFolder", "Folder", true},
	Type{23, "CfgField", "Field", true},
	Type{24, "CfgFormat", "Format", true},
	Type{25, "CfgTableAccess", "Table Access", true},
	Type{26, "CfgCallingList", "Calling List", true},
	Type{27, "CfgCampaign", "Campaign", true},
	Type{28, "CfgTreatment", "Treatment", true},
	Type{29, "CfgFilter", "Filter", true},
	Type{30, "CfgTimeZone", "Time Zone", true},
	Type{31, "CfgVoicePrompt", "Voice Prompt", true},
	Type{32, "CfgIVRPort", "IVR Port", true},
	Type{33, "CfgIVR", "IVR", true},
	Type{34, "CfgAlarmCondition", "Alarm Condition", true},
	Type{35, "CfgEnumerator", "Business Attribute", true},
	Type{36, "CfgEnumeratorValue", "Business Attribute Value", true},
	Type{37, "CfgObjectiveTable", "Objective Table", true},
	Type{38, "CfgCampaignGroup", "Campaign Group", true},
	//Type{39, "CfgGVPReseller", "GVP Reseller",true},
	//Type{40, "CfgGVPCustomer", "GVP Customer",true},
	Type{41, "CfgGVPIVRProfile", "GVP IVRProfile", true},
	//Type{42, "CfgScheduledTask ", "Scheduled Task",true},
	Type{43, "CfgRole", "Role", true},
	//	Type{44, "CfgPersonLastLogin", "PersonLastLogin",true},
	//	Type{45, "CfgMaxType", "Shortcut",true},
}

//LoginRequest Data send for a login request
type LoginRequest struct {
	Username            string `json:"username"`
	Password            string `json:"password"`
	IsPasswordEncrypted bool   `json:"isPasswordEncrypted"`
}

//LoginResponse Data receive during a login request
type LoginResponse struct {
	Username       string `json:"username"`
	UserType       string `json:"userType"`
	SessionTimeout int    `json:"sessionTimeout"`
	IsDefaultUser  bool   `json:"isDefaultUser"`
	WriteDefault   bool   `json:"writeDefault"`
}

//CfgObject generic object
type CfgObject struct {
	Dbid string `json:"dbid"`
	Name string `json:"name"`
	Type string `json:"type"`
}

//CfgFolder represent a folder
type CfgFolder struct {
	Folderclass string `json:"folderclass"`
	Subtype     string `json:"subtype"`
	Ownertype   string `json:"ownertype"`
	Dbid        string `json:"dbid"`
	Ownerdbid   string `json:"ownerdbid"`
	Folderid    string `json:"folderid"`
	Name        string `json:"name"`
	Resources   struct {
		Resource []interface{} `json:"resource"` //TODO
	} `json:"resources"`
	State      string `json:"state"`
	Type       string `json:"type"`
	Customtype string `json:"customtype"`
	Objectids  struct {
		Idtype CfgDBIDTypeList `json:"idtype"`
	} `json:"objectids"`
}

//CfgAppPrototype represent a application template
type CfgAppPrototype struct {
	Subtype  string  `json:"subtype"`
	Dbid     string  `json:"dbid"`
	Name     string  `json:"name"`
	Options  Options `json:"options"`
	State    string  `json:"state"`
	Type     string  `json:"type"`
	Version  string  `json:"version"`
	Folderid string  `json:"folderid"`
}

//TODO add get metadata

/*

//CfgSwitch TODO
type CfgSwitch struct {
	*CfgObject
	//TODO
}

*/

//CfgDN represent a DN object
type CfgDN struct {
	Accessnumbers struct {
		Dnaccessnumber []interface{} `json:"dnaccessnumber"`
	} `json:"accessnumbers"`
	Name               string         `json:"name"`
	Contractdbid       string         `json:"contractdbid"`
	Dbid               string         `json:"dbid"`
	Folderid           string         `json:"folderid"`
	Groupdbid          string         `json:"groupdbid"`
	Loginflag          string         `json:"loginflag"`
	Number             string         `json:"number"`
	Registerall        string         `json:"registerall"`
	Routetype          string         `json:"routetype"`
	Sitedbid           string         `json:"sitedbid"`
	State              string         `json:"state"`
	Subtype            string         `json:"subtype"`
	Switchdbid         string         `json:"switchdbid"`
	Switchspecifictype string         `json:"switchspecifictype"`
	Tenantdbid         string         `json:"tenantdbid"`
	Trunks             string         `json:"trunks"`
	Type               string         `json:"type"`
	Useoverride        string         `json:"useoverride"`
	Userproperties     Userproperties `json:"userproperties,omitempty"`
}

//CfgDNGroup represent a group of DN object
type CfgDNGroup struct {
	Capacityruledbid  string `json:"capacityruledbid"`
	Capacitytabledbid string `json:"capacitytabledbid"`
	Contractdbid      string `json:"contractdbid"`
	Dbid              string `json:"dbid"`
	DNS               struct {
		Dninfo []struct {
			Dndbid string `json:"dndbid"`
			Trunks string `json:"trunks"`
		} `json:"dninfo"`
	} `json:"dns"`
	Folderid       string         `json:"folderid"`
	Name           string         `json:"name"`
	Quotatabledbid string         `json:"quotatabledbid"`
	Sitedbid       string         `json:"sitedbid"`
	State          string         `json:"state"`
	Subtype        string         `json:"subtype"`
	Tenantdbid     string         `json:"tenantdbid"`
	Type           string         `json:"type"`
	Userproperties Userproperties `json:"userproperties,omitempty"`
}

//CfgDBIDList represent a generic list of dbid link
type CfgDBIDList []struct {
	Dbid string `json:"dbid"`
}

//CfgDBIDTypeList represent a generic list of dbid with type link
type CfgDBIDTypeList []struct {
	Dbid string `json:"dbid"`
	Type string `json:"type,omitempty"`
}

//CfgAccessGroup represent a AccessGroup
type CfgAccessGroup struct {
	Capacityruledbid  string `json:"capacityruledbid"`
	Capacitytabledbid string `json:"capacitytabledbid"`
	Contractdbid      string `json:"contractdbid"`
	Dbid              string `json:"dbid"`
	Folderid          string `json:"folderid"`
	Memberids         struct {
		Idtype CfgDBIDTypeList `json:"idtype"`
	} `json:"memberids"`
	Name           string `json:"name"`
	Quotatabledbid string `json:"quotatabledbid"`
	Sitedbid       string `json:"sitedbid"`
	State          string `json:"state"`
	Subtype        string `json:"subtype"`
	Tenantdbid     string `json:"tenantdbid"`
	Type           string `json:"type"`
}

//CfgAgentGroup represent a agent group
type CfgAgentGroup struct {
	Agentdbids struct {
		ID CfgDBIDTypeList `json:"id"`
	} `json:"agentdbids"`
	Capacityruledbid  string `json:"capacityruledbid"`
	Capacitytabledbid string `json:"capacitytabledbid"`
	Contractdbid      string `json:"contractdbid"`
	Dbid              string `json:"dbid"`
	Folderid          string `json:"folderid"`
	Managerdbids      struct {
		ID CfgDBIDTypeList `json:"id"`
	} `json:"managerdbids"`
	Name           string `json:"name"`
	Quotatabledbid string `json:"quotatabledbid"`
	Sitedbid       string `json:"sitedbid"`
	State          string `json:"state"`
	Tenantdbid     string `json:"tenantdbid"`
	Type           string `json:"type"`
}

//CfgAgentLogin represent a agent login
type CfgAgentLogin struct {
	Dbid               string         `json:"dbid"`
	Folderid           string         `json:"folderid"`
	Logincode          string         `json:"logincode"`
	State              string         `json:"state"`
	Switchdbid         string         `json:"switchdbid"`
	Switchspecifictype string         `json:"switchspecifictype"`
	Tenantdbid         string         `json:"tenantdbid"`
	Type               string         `json:"type"`
	Useoverride        string         `json:"useoverride"`
	Userproperties     Userproperties `json:"userproperties"`
}

//CfgPerson represent a person
type CfgPerson struct {
	AgentLogins struct {
		Agentlogininfo []struct {
			Agentlogindbid string `json:"agentlogindbid"`
			Wrapuptime     string `json:"wrapuptime"`
		} `json:"agentlogininfo"`
	} `json:"agentlogins"`
	Appranks struct {
		Apprank []interface{} `json:"apprank"`
	} `json:"appranks"`
	Changepasswordonnextlogin string         `json:"changepasswordonnextlogin"`
	Dbid                      string         `json:"dbid"`
	Employeeid                string         `json:"employeeid"`
	Firstname                 string         `json:"firstname"`
	Folderid                  string         `json:"folderid"`
	Isagent                   string         `json:"isagent"`
	Isexternalauth            string         `json:"isexternalauth"`
	Lastname                  string         `json:"lastname"`
	Password                  string         `json:"password"`
	State                     string         `json:"state"`
	Tenantdbid                string         `json:"tenantdbid"`
	Type                      string         `json:"type"`
	Username                  string         `json:"username"`
	Userproperties            Userproperties `json:"userproperties"`
}

//CfgHost represent a server host
type CfgHost struct {
	Dbid      string `json:"dbid"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	Ipaddress string `json:"ipaddress"`
	Scsdbid   string `json:"scsdbid"`
	Subtype   string `json:"subtype"`
	Lcaport   string `json:"lcaport"`
	Ostype    string `json:"ostype"`
	State     string `json:"state"`
	Folderid  string `json:"folderid"`
}

//Userproperties generic annex list
type Userproperties struct {
	Property []Property `json:"property"`
}

//Options generic option list
type Options struct {
	Property []Property `json:"property"`
}

//Property generic prop definition
type Property struct {
	Section string `json:"section"`
	Value   string `json:"value"`
	Key     string `json:"key"`
}

//Appservers generic app server list
type Appservers struct {
	Conninfo []Conninfo `json:"conninfo,omitempty"`
}

//Conninfo generic app server connection definition
type Conninfo struct {
	Mode          string `json:"mode"`
	Appserverdbid string `json:"appserverdbid"`
	Timoutlocal   string `json:"timoutlocal"`
	Longfield1    string `json:"longfield1"`
	Longfield2    string `json:"longfield2"`
	Longfield3    string `json:"longfield3"`
	Longfield4    string `json:"longfield4"`
	Timoutremote  string `json:"timoutremote"`
	ID            string `json:"id"`
}

//CfgPlace represent a place definition
type CfgPlace struct {
	Sitedbid         string         `json:"sitedbid"`
	Dbid             string         `json:"dbid"`
	Tenantdbid       string         `json:"tenantdbid"`
	Name             string         `json:"name"`
	Userproperties   Userproperties `json:"userproperties,omitempty"`
	Contractdbid     string         `json:"contractdbid"`
	State            string         `json:"state"`
	Type             string         `json:"type"`
	Capacityruledbid string         `json:"capacityruledbid"`
	Folderid         string         `json:"folderid"`
	Dndbids          struct {
		ID CfgDBIDList `json:"id,omitempty"`
	} `json:"dndbids,omitempty"`
}

//CfgApplication represent an application definition
type CfgApplication struct {
	Dbid                 string         `json:"dbid,omitempty"`
	Name                 string         `json:"name"`
	Type                 string         `json:"type"`
	Appservers           Appservers     `json:"appservers,omitempty"`
	Autorestart          string         `json:"autorestart,omitempty"`
	Userproperties       Userproperties `json:"userproperties,omitempty"`
	Timeout              string         `json:"timeout,omitempty"`
	Commandline          string         `json:"commandline,omitempty"`
	Folderid             string         `json:"folderid"`
	Commandlinearguments string         `json:"commandlinearguments,omitempty"`
	Subtype              string         `json:"subtype,omitempty"`
	Options              Options        `json:"options,omitempty"`
	State                string         `json:"state,omitempty"`
	Hostdbid             string         `json:"hostdbid,omitempty"`
	Attempts             string         `json:"attempts,omitempty"`
	Portinfos            struct {
		Portinfo []struct {
			Longfield1 string `json:"longfield1"`
			Longfield2 string `json:"longfield2"`
			Longfield3 string `json:"longfield3"`
			Port       string `json:"port"`
			Longfield4 string `json:"longfield4"`
			ID         string `json:"id"`
		} `json:"portinfo,omitempty"`
	} `json:"portinfos,omitempty"`
	Workdirectory string `json:"workdirectory,omitempty"`
	Startuptype   string `json:"startuptype,omitempty"`
	Isserver      string `json:"isserver,omitempty"`
	Resources     struct {
		Resource []interface{} `json:"resource,omitempty"`
	} `json:"resources,omitempty"`
	Tenantdbids struct {
		ID   CfgDBIDList `json:"id,omitempty"`
		Mode string      `json:"mode,omitempty"`
	} `json:"tenantdbids,omitempty"`
	Startuptimeout   string `json:"startuptimeout,omitempty"`
	Backupserverdbid string `json:"backupserverdbid,omitempty"`
	Version          string `json:"version,omitempty"`
	Isprimary        string `json:"isprimary,omitempty"`
	Redundancytype   string `json:"redundancytype,omitempty"`
	Shutdowntimeout  string `json:"shutdowntimeout,omitempty"`
	Componenttype    string `json:"componenttype,omitempty"`
	Appprototypedbid string `json:"appprototypedbid"`
	Port             string `json:"port,omitempty"`
}

//GetOptionValue return option value or default if not found
func (a *CfgApplication) GetOptionValue(section, key, defaultValue string) string {
	for _, ap := range a.Options.Property {
		if ap.Section == section && ap.Key == key {
			return ap.Value
		}
	}
	//If not found return default value
	return defaultValue
}

//SetOptions update or create multiple option passed as arg
func (a *CfgApplication) SetOptions(optList ...Property) {
	for _, o := range optList {
		a.SetOption(o)
	}
}

//SetOption update or create option passed as arg
func (a *CfgApplication) SetOption(o Property) {
	for i, ap := range a.Options.Property {
		if ap.Section == o.Section && ap.Key == o.Key {
			a.Options.Property[i] = o
			return
		}
	}
	//If not found append
	a.Options.Property = append(a.Options.Property, o)
}

//SetAnnex Modify or add an annex
func (props *Userproperties) SetAnnex(newOpt Property) {
	for i, p := range props.Property {
		if newOpt.Section == p.Section && newOpt.Key == p.Key {
			props.Property[i] = newOpt
			return //End here
		}
	}
	//If not found append
	props.Property = append(props.Property, newOpt)
}

//GetAnnex Search an annex by Section and Key of param
func (props *Userproperties) GetAnnex(searchP Property) *Property {
	for _, p := range props.Property {
		if searchP.Section == p.Section && searchP.Key == p.Key {
			return &p
		}
	}
	//If not found
	return nil
}

//Contains Search if an annex with exact Section, Key and Value of param
func (props *Userproperties) Contains(searchP Property) bool {
	p := props.GetAnnex(searchP)
	if p == nil {
		//If not found
		return false
	}
	return p.Value == searchP.Value
}
