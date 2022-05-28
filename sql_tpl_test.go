package gqttool

import (
	"fmt"
	"testing"
)

func TestCrud(t *testing.T) {
	ddlList := GetDDL()
	tableCfg := &DatabaseConfig{}
	repo := NewRepositoryMeta()
	err := repo.AddByDir("example", MetaTemplatefuncMap)
	if err != nil {
		panic(err)
	}
	tables, err := GenerateTable(ddlList, tableCfg)
	if err != nil {
		panic(err)
	}
	for _, table := range tables {
		tplDefineList, err := GenerateSQLTpl(table, repo)
		if err != nil {
			panic(err)
		}
		for _, tplDefine := range tplDefineList {
			fmt.Println(tplDefine.Text)
		}
	}

}

func TestParseSQLSelectColumn(t *testing.T) {
	sqlDefine := `
	{{define "paginate"}}
		select *,count(*) from t_api where 1=1;
	{{end}}
	`
	fields := ParseSQLSelectColumn(sqlDefine)
	fmt.Sprintln(fields)
}
