package gqttool

import (
	"fmt"
	"testing"

	"github.com/suifengpiao14/gqt/v2"
)

func TestCrud(t *testing.T) {
	ddlList := GetDDL()
	tableCfg := &Config{}
	repo := NewRepositoryMeta()
	err := repo.AddByDir("example", gqt.TemplatefuncMap)
	if err != nil {
		panic(err)
	}
	tables, err := GenerateTable(ddlList, tableCfg)
	if err != nil {
		panic(err)
	}
	for _, table := range tables {
		tplDefineList, err := Crud(table, repo)
		if err != nil {
			panic(err)
		}
		for _, tplDefine := range tplDefineList {
			fmt.Println(tplDefine.TPL)
		}
	}

}
