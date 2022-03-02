package main

import (
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
	"testing"

	"github.com/iancoleman/strcase"
	"github.com/suifengpiao14/errorformatter"
	"github.com/suifengpiao14/gqt/v2"
	"github.com/suifengpiao14/gqttool"
)

func main() {
	var (
		ddlFile        string
		modelFilename  string
		tplDir         string
		entityFilename string
		err            error
	)
	modelCmd := flag.NewFlagSet("model", flag.ExitOnError)
	modelCmd.Usage = helpModel
	crudCmd := flag.NewFlagSet("crud", flag.ExitOnError)
	crudCmd.Usage = helpCrud
	entityCmd := flag.NewFlagSet("entity", flag.ExitOnError)
	entityCmd.Usage = helpEntity
	modelCmd.StringVar(&ddlFile, "ddl", "template/ddl.sql.tpl", "ddl template file name")
	modelCmd.StringVar(&modelFilename, "model", "repository.model.go", "ddl template file name")

	crudCmd.StringVar(&ddlFile, "ddl", "template/ddl.sql.tpl", "ddl template file name")
	crudCmd.StringVar(&tplDir, "tplDir", "template", "sql template dir")

	entityCmd.StringVar(&tplDir, "tplDir", "template", "sql template dir")
	entityCmd.StringVar(&entityFilename, "entity", "repository.entity.go", "entity file")
	testing.Init()

	if len(os.Args) < 3 {
		help()
	}
	args := os.Args[2:]
	switch os.Args[1] {
	case "model":
		modelCmd.Parse(args)
		err = runCmdModel(ddlFile, modelFilename)
		if err != nil {
			panic(err)
		}

	case "crud":
		crudCmd.Parse(args)
		err = runCmdCrud(ddlFile, tplDir)
		if err != nil {
			panic(err)
		}

	case "entity":
		entityCmd.Parse(args)
		err = runCmdEntity(tplDir, entityFilename)
		if err != nil {
			panic(err)
		}

	default:
		help()
	}
}

func generatePackageName(dstDir string) (packageName string) {
	basename := path.Base(dstDir)
	packageName = strings.ToLower(strcase.ToLowerCamel(basename))
	return

}

func runCmdModel(ddlFile string, dstDir string) (err error) {
	repo := gqt.NewRepository()
	errChain := errorformatter.NewErrorChain()
	var content string
	var modelMap map[string]string
	errChain.Run(func() (err error) {
		content, err = GetFileContent(ddlFile)
		return
	}).
		SetError(repo.AddByNamespace("ddl", content, gqt.TemplatefuncMap)).
		Run(func() (err error) {
			modelMap, err = GenerateTableModel(repo)
			return
		})
	err = errChain.Error()
	if err != nil {
		return
	}
	contentArr := make([]string, 0)
	packageName := generatePackageName(dstDir)
	packageLine := fmt.Sprintf("package %s", packageName)
	contentArr = append(contentArr, packageLine)
	tableList := make([]string, 0)
	for name := range modelMap {
		tableList = append(tableList, name)
	}
	sort.Strings(tableList) // 排序后写入文件

	for _, tablename := range tableList {
		model := modelMap[tablename]
		contentArr = append(contentArr, model)
	}

	content = strings.Join(contentArr, "\n")
	filename := fmt.Sprintf("%s/model.go", dstDir)
	err = saveFile(filename, content)
	if err != nil {
		return
	}
	return
}

func runCmdCrud(ddlFile string, dstDir string) (err error) {
	repo := gqt.NewRepository()
	errChain := errorformatter.NewErrorChain()
	var content string
	var tplMap map[string]string
	errChain.Run(func() (err error) {
		content, err = GetFileContent(ddlFile)
		return
	}).
		SetError(repo.AddByNamespace("ddl", content, gqt.TemplatefuncMap)).
		Run(func() (err error) {
			tplMap, err = GenerateCrud(repo)
			return
		})

	err = errChain.Error()
	if err != nil {
		return
	}
	for name, tpl := range tplMap {
		snakeName := strcase.ToSnake(name)
		filename := fmt.Sprintf("%s/%s.sql.tpl", dstDir, snakeName)
		err = saveFile(filename, tpl)
		if err != nil {
			return
		}
	}
	return
}

func runCmdEntity(sqlTplDir string, entityFilename string) (err error) {
	repo := gqt.NewRepository()
	errChain := errorformatter.NewErrorChain()
	var entityList = make([]string, 0)
	var sqlTplDefineList = make([]*gqttool.SQLTPLDefine, 0)
	errChain.SetError(repo.AddByDir(sqlTplDir, gqt.TemplatefuncMap)).
		Run(func() (err error) {
			sqlTplDefineList, err = gqttool.ParseDirSqlTplDefine(sqlTplDir)
			return
		}).
		Run(func() (err error) {
			entityList, err = GenerateEntity(repo, sqlTplDefineList)
			return
		})

	err = errChain.Error()
	if err != nil {
		return
	}
	sort.Strings(entityList)

	contentArr := make([]string, 0)
	packageName := generatePackageName(filepath.Dir(sqlTplDir))
	packageLine := fmt.Sprintf("package %s", packageName)
	contentArr = append(contentArr, packageLine)
	contentArr = append(contentArr, entityList...)
	content := strings.Join(contentArr, "\n")
	err = saveFile(entityFilename, content)
	if err != nil {
		return
	}
	return
}

func GetFileContent(file string) (content string, err error) {
	f, err := os.OpenFile(file, os.O_RDONLY, fs.ModePerm)
	if err != nil {
		return
	}
	b, err := io.ReadAll(f)
	if err != nil {
		return
	}
	content = string(b)
	return

}

func saveFile(filename string, content string) (err error) {
	if IsExist(filename) {
		return
	}
	f, err := os.Create(filename)
	if err != nil {
		return
	}
	_, err = f.WriteString(content)
	if err != nil {
		return
	}
	return

}

func IsExist(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

func GenerateEntity(rep *gqt.Repository, sqlTplDefineList []*gqttool.SQLTPLDefine) (entityList []string, err error) {
	entityList = make([]string, 0)
	ddlList, err := getDDLFromRepository(rep)
	if err != nil {
		return
	}
	tableList, err := gqttool.GenerateTable(ddlList)
	if err != nil {
		return
	}

	for _, sqlDefineTpl := range sqlTplDefineList {
		entityStruct, err := gqttool.RepositoryEntity(sqlDefineTpl, tableList)
		if err != nil {
			return nil, err
		}
		entityList = append(entityList, entityStruct)

	}

	return
}

func GenerateCrud(rep *gqt.Repository) (tplMap map[string]string, err error) {
	ddlList, err := getDDLFromRepository(rep)
	if err != nil {
		return
	}
	tableList, err := gqttool.GenerateTable(ddlList)
	if err != nil {
		return
	}
	tplMap = make(map[string]string)
	for _, table := range tableList {
		oneTableTplList, err := gqttool.Crud(table)
		if err != nil {
			return nil, err
		}
		tableTpl := strings.Join(oneTableTplList, "\n")
		tplMap[table.TableName] = tableTpl
	}

	return
}
func GenerateTableModel(rep *gqt.Repository) (modelMap map[string]string, err error) {
	ddlList, err := getDDLFromRepository(rep)
	if err != nil {
		return
	}
	tableList, err := gqttool.GenerateTable(ddlList)
	if err != nil {
		return
	}
	modelMap, err = gqttool.GenerateModel(tableList)
	return
}

func getDDLFromRepository(rep *gqt.Repository) (ddlList []string, err error) {
	ddlList = make([]string, 0)
	ddlMap, err := rep.GetByNamespace("ddl", nil)
	if err != nil {
		return
	}
	for _, ddl := range ddlMap {
		ddlList = append(ddlList, ddl)
	}
	return
}

func helpModel() {
	fmt.Fprint(os.Stderr, `gqttool model is generate go struct from mysql ddl

Usage:
  gqttool model  -ddl ddlFilename -model modelFilename

Flags:
  -ddl
        mysql ddl file path

  -model
        model file path

Example:

  gqttool model template/ddl.sql.tpl -model repository.model.go

`)
	os.Exit(0)
}

func helpCrud() {
	fmt.Fprint(os.Stderr, `gqttool crud is  generation crud sql from mysql ddl

Usage:
  gqttool crud  -ddl ddlFilename -tplDir sqlTplSaveDir
  
Flags:
  -ddl
        mysql ddl file path

  -tplDir
        save sqlTpl file path

Example:

  gqttool crud template/ddl.sql.tpl -tplDir template

`)
	os.Exit(0)
}

func helpEntity() {
	fmt.Fprint(os.Stderr, `gqttool entity is  generation sql template inpur args entity

Usage:
  gqttool entity  -tplDir sqlTplDir -entity entityFilename
  

Flags:
 -tplDir 
		sqlTpl file dir
 -entity 
		sqlTpl  entity file name

Example:

  gqttool entity -tplDir template -entity repository.entity.go

`)
	os.Exit(0)
}

func help() {
	fmt.Fprint(os.Stderr, `gqttool is the code generation tool for the gqt package.

Usage:
  gqttool model  -ddl ddlFilename -model modelFilename
  gqttool crud  -ddl ddlFilename -tplDir sqlTplSaveDir
  gqttool entity  -tplDir sqlTplDir -entity entityFilename
  
Commands:
  model
  		Generate go struct from  mysql ddl
  crud
        Generate crud sql.tpl from mysql ddl
  entity
  		Generate sql.tpl input entity from mysql sqlTplDir

Flags:
  -ddl
        mysql ddl file path
  -model
        repository model file name
 -tplDir 
		sqlTpl file dir
 -entity 
		sqlTpl  argument entity file name

Example:

  gqttool model template/ddl.sql.tpl -model repository.model.go

`)
	os.Exit(0)
}
