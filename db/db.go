package db

import (
	"fmt"
	"sync"

	"xorm.io/xorm"
)

var (
	//handlerMap is a list of handler to generate db url
	handlerMap sync.Map
)

//handler handle some specific code to db
type handler struct {
	GenerateURL func(host, user, pass, database string) (string, string, error)
}

//initConn start the database connection and settings
func initConn(t, s string) (*xorm.Engine, error) {
	engine, err := xorm.NewEngine(t, s)
	if err != nil {
		return nil, err
	}
	//engine.SetLogger(generateSQLLogger())
	//engine.ShowSQL(true)
	engine.ShowSQL(false)
	return engine, engine.Ping()
}

//SupportedTypes list suported db handler
func SupportedTypes() []string {
	dbTypeEntryArr := make([]string, 0)
	handlerMap.Range(func(key, value interface{}) bool {
		dbTypeEntryArr = append(dbTypeEntryArr, key.(string))
		return true
	})
	return dbTypeEntryArr
}

//DB contain the DB
type DB struct {
	dbType string
	//Engine contane the connection to the Database
	cfgEngine  *xorm.Engine
	cfgURL     string
	rtmeEngine *xorm.Engine
	rtmeURL    string
}

//NewDB create DB object
func NewDB(dbType, host, user, pass, cfg, rtme string) (*DB, error) {
	if dbType == "" {
		return nil, fmt.Errorf("Invalid db type")
	}
	dbHandlerStringer, ok := handlerMap.Load(dbType)
	if !ok {
		return nil, fmt.Errorf("Invalid db type")
	}
	_, cu, err := dbHandlerStringer.(handler).GenerateURL(host, user, pass, cfg)
	if err != nil {
		return nil, err
	}
	_, ru, err := dbHandlerStringer.(handler).GenerateURL(host, user, pass, cfg)
	if err != nil {
		return nil, err
	}

	return &DB{
		dbType:  dbType,
		cfgURL:  cu,
		rtmeURL: ru,
	}, nil
}

//NewDBFromURL create DB object directly from url
func NewDBFromURL(dbType, cfgURL, rtmeURL string) *DB {
	return &DB{
		dbType:  dbType,
		cfgURL:  cfgURL,
		rtmeURL: rtmeURL,
	}
}

//RTME handle to rtme database
func (db *DB) RTME() (*xorm.Engine, error) {
	if db.rtmeEngine == nil {
		e, err := generetaDBEngine(db.dbType, db.rtmeURL)
		if err != nil {
			return nil, err
		}
		db.rtmeEngine = e
	}
	return db.rtmeEngine, nil
}

//CFG handle to config database
func (db *DB) CFG() (*xorm.Engine, error) {
	if db.cfgEngine == nil {
		e, err := generetaDBEngine(db.dbType, db.cfgURL)
		if err != nil {
			return nil, err
		}
		db.cfgEngine = e
	}
	return db.cfgEngine, nil
}

//generetaDBEngine create DB object from a specific url
func generetaDBEngine(dbType, dbURL string) (*xorm.Engine, error) {
	e, err := initConn(dbType, dbURL)
	if err != nil {
		return nil, err
	}
	return e, nil
}
