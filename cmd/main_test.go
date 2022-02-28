package main

import (
	"fmt"
	"testing"

	"github.com/suifengpiao14/gqt/v2"
)

func TestRunCmdModel(t *testing.T) {
	ddlFile := "../example/template/ddl.sql.tpl"
	modelDir := "../example/"
	err := runCmdModel(ddlFile, modelDir)
	if err != nil {
		panic(err)
	}

}

func TestGenerateModel(t *testing.T) {
	repo := gqt.NewRepository()
	err := repo.AddByDir("../example", gqt.TemplatefuncMap)
	if err != nil {
		panic(err)
	}
	tableModelList, err := GenerateTableModel(repo)
	if err != nil {
		panic(err)
	}
	fmt.Println(tableModelList)
}
