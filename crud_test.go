package gqttool

import (
	"fmt"
	"testing"

	"github.com/suifengpiao14/gqt/v2"
)

func TestCrud(t *testing.T) {
	ddlList := GetDDL()
	tableCfg := &gqt.Config{}
	tables, err := GenerateTable(ddlList, tableCfg)
	if err != nil {
		panic(err)
	}
	for _, table := range tables {
		tplMap := Crud(table)
		fmt.Println(tplMap)
	}

}
