package db

import (
	_ "github.com/denisenkom/go-mssqldb"
)

func init() {
	handlerMap.Store("Microsoft SQL Server", handler{
		GenerateURL: func(host, user, pass, database string) (string, string, error) {
			return "mssql", "server=" + host + ";user id=" + user + ";password=" + pass + ";database=" + database, nil
		},
	})
}
