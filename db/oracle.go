//TODO add copyrigths
package db

import "fmt"

func init() {
	handlerMap.Store("Oracle", handler{
		GenerateURL: func(host, user, pass, database string) (string, string, error) {
			return "", "", fmt.Errorf("DB Type Not Implemented")
		},
	})
}
