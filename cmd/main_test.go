package main

import (
	"fmt"
	"testing"

	"github.com/suifengpiao14/gqt/v2"
	"github.com/suifengpiao14/gqttool"
)

func TestRunCmdModel(t *testing.T) {
	ddlFile := "../example/template/ddl.sql.tpl"
	modelDir := "../example/"
	err := runCmdModel(ddlFile, modelDir)
	if err != nil {
		panic(err)
	}

}

func TestGeneratePackageName(t *testing.T) {
	dst := ""
	dst = "."
	dst = ".."
	dst = "../../"
	dst = "../example"
	packageName, err := gqttool.GeneratePackageName(dst)
	if err != nil {
		panic(err)
	}
	fmt.Println(packageName)
}

func TestRunCmdEntity(*testing.T) {
	tplDir := "../example/template"
	entity := "../example/repository.entity.go"
	err := runCmdEntity(tplDir, entity)
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
