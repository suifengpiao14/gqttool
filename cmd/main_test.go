package main

import (
	"fmt"
	"testing"

	"github.com/suifengpiao14/gqttool"
)

func TestRunCmdModel(t *testing.T) {
	metaDir := "../example/template/"
	modelFile := "../example/model.gen.go"
	err := runCmdModel(metaDir, modelFile, true)
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
	err := runCmdEntity(tplDir, entity, true)
	if err != nil {
		panic(err)
	}
}
func TestRunCmdCrud(*testing.T) {
	metaDir := "../example/template/meta"
	tplDir := "../example/template/sql"
	err := runCmdSQL(metaDir, tplDir, true)
	if err != nil {
		panic(err)
	}
}

func TestGenerateModel(t *testing.T) {
	repo := gqttool.NewRepositoryMeta()
	err := repo.AddByDir("../example/template", gqttool.MetaTemplatefuncMap)
	if err != nil {
		panic(err)
	}
	tableModelList, err := GenerateTableModel(repo)
	if err != nil {
		panic(err)
	}
	for _, tableModel := range tableModelList {
		fmt.Printf("%#v", tableModel)
	}

}
