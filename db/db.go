package db

import (
	"fmt"
	"sync"

	"github.com/go-xorm/xorm"
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
	//Engine contane the connection to the Database
	Engine *xorm.Engine
}

//NewDB create DB object
func NewDB(dbType, host, user, pass, base string) (*DB, error) {
	if dbType == "" {
		return nil, fmt.Errorf("Invalid db type")
	}
	dbHandlerStringer, ok := handlerMap.Load(dbType)
	if !ok {
		return nil, fmt.Errorf("Invalid db type")
	}
	t, u, err := dbHandlerStringer.(handler).GenerateURL(host, user, pass, base)
	if err != nil {
		return nil, err
	}

	e, err := initConn(t, u)
	if err != nil {
		return nil, err
	}
	return &DB{
		Engine: e,
	}, nil
}
