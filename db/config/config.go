package config

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sapk/go-genesys/db"
)

//GetAgentGroup retrive agent_dbid of an gent_group
func GetAgentGroup(d *db.DB, groupFilter string) ([]int, error) {
	e, err := d.CFG()
	if err != nil {
		return []int{}, err
	}
	results, err := e.Query(fmt.Sprintf("SELECT agent_dbid FROM cfg_agent_group INNER JOIN cfg_group ON cfg_group.dbid = cfg_agent_group.group_dbid WHERE cfg_group.name = '%s'", groupFilter))
	if err != nil {
		return []int{}, err
	}
	list := make([]int, len(results))
	for i, r := range results {
		aid, err := strconv.Atoi(strings.TrimSpace(string(r["agent_dbid"])))
		if err != nil {
			return []int{}, err
		}
		list[i] = aid
	}
	return list, nil
}
