package main

import (
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path"
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
		force          bool
		err            error
	)
	modelCmd := flag.NewFlagSet("model", flag.ExitOnError)
	modelCmd.Usage = helpModel
	crudCmd := flag.NewFlagSet("crud", flag.ExitOnError)
	crudCmd.Usage = helpCrud
	entityCmd := flag.NewFlagSet("entity", flag.ExitOnError)
	entityCmd.Usage = helpEntity
	modelCmd.BoolVar(&force, "force", false, "overwrite exist file")
	modelCmd.StringVar(&ddlFile, "ddl", "template/ddl.sql.tpl", "ddl template file name")
	modelCmd.StringVar(&modelFilename, "model", "repository.model.go", "ddl template file name")

	crudCmd.BoolVar(&force, "f", false, "overwrite exist file")
	crudCmd.StringVar(&ddlFile, "ddl", "template/ddl.sql.tpl", "ddl template file name")
	crudCmd.StringVar(&tplDir, "tplDir", "template", "sql template dir")

	entityCmd.BoolVar(&force, "f", false, "overwrite exist file")
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
		err = runCmdModel(ddlFile, modelFilename, force)
		if err != nil {
			panic(err)
		}

	case "crud":
		crudCmd.Parse(args)
		err = runCmdCrud(ddlFile, tplDir, force)
		if err != nil {
			panic(err)
		}

	case "entity":
		entityCmd.Parse(args)
		err = runCmdEntity(tplDir, entityFilename, force)
		if err != nil {
			panic(err)
		}

	default:
		help()
	}
}

func runCmdModel(ddlFile string, modelFile string, force bool) (err error) {
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
	packageName, err := gqttool.GeneratePackageName(path.Dir(modelFile))
	if err != nil {
		return
	}
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
	err = saveFile(modelFile, content, force)
	if err != nil {
		return
	}
	return
}

func runCmdCrud(ddlFile string, dstDir string, force bool) (err error) {
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
		err = saveFile(filename, tpl, force)
		if err != nil {
			return
		}
	}
	return
}

func runCmdEntity(sqlTplDir string, entityFilename string, force bool) (err error) {
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
	packageName, err := gqttool.GeneratePackageName(path.Dir(entityFilename))
	if err != nil {
		return
	}
	packageLine := fmt.Sprintf("package %s", packageName)
	contentArr = append(contentArr, packageLine)
	contentArr = append(contentArr, entityList...)
	content := strings.Join(contentArr, "\n")
	err = saveFile(entityFilename, content, force)
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

func saveFile(filename string, content string, force bool) (err error) {
	if IsExist(filename) {
		if !force {
			return
		}
		err = os.Remove(filename)
		if err != nil {
			return
		}
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
	tableMap := make(map[string]*gqttool.Table)
	for _, table := range tableList {
		tableMap[table.TableName] = table
	}

	for _, sqlDefineTpl := range sqlTplDefineList {
		tableNameList, err := gqttool.ParseSQLTPLTableName(sqlDefineTpl.TPL)
		if err != nil {
			return nil, err
		}
		relationTableList := make([]*gqttool.Table, 0)
		for _, tableName := range tableNameList {
			table, ok := tableMap[tableName]
			if ok {
				relationTableList = append(relationTableList, table)
			}
		}

		entityStruct, err := gqttool.RepositoryEntity(sqlDefineTpl, relationTableList)
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
  gqttool model -force true -ddl ddlFilename -model modelFilename

Flags:
  -force overwrite exists file
  -ddl
        mysql ddl file path

  -model
        model filename

Example:

  gqttool model -force true -ddl template/ddl.sql.tpl -model repository.model.go

`)
	os.Exit(0)
}

func helpCrud() {
	fmt.Fprint(os.Stderr, `gqttool crud is  generation crud sql from mysql ddl

Usage:
  gqttool crud  -force true -ddl ddlFilename -tplDir sqlTplSaveDir
  
Flags:
  -force overwrite exists file
  -ddl
        mysql ddl filename

  -tplDir
        save sqlTpl file path

Example:

  gqttool crud -force true -ddl template/ddl.sql.tpl -tplDir template

`)
	os.Exit(0)
}

func helpEntity() {
	fmt.Fprint(os.Stderr, `gqttool entity is  generation sql template inpur args entity

Usage:
  gqttool entity -force true  -tplDir sqlTplDir -entity entityFilename
  

Flags:
 -force overwrite exists file 
 -tplDir 
		sqlTpl file dir
 -entity 
		sqlTpl  entity filename

Example:

  gqttool entity -force true -tplDir template -entity repository.entity.go

`)
	os.Exit(0)
}

func help() {
	fmt.Fprint(os.Stderr, `gqttool is the code generation tool for the gqt package.

Usage:
  gqttool model  -force true -ddl ddlFilename -model modelFilename
  gqttool crud -force true  -ddl ddlFilename -tplDir sqlTplSaveDir
  gqttool entity -force true  -tplDir sqlTplDir -entity entityFilename
  
Commands:
  model
  		Generate go struct from  mysql ddl
  crud
        Generate crud sql.tpl from mysql ddl
  entity
  		Generate sql.tpl input entity from mysql sqlTplDir

Flags:
  -force overwrite exists file
  -ddl
        mysql ddl filename
  -model
        repository model filename
 -tplDir 
		sqlTpl file dir
 -entity 
		sqlTpl  argument entity filename

Example:

  gqttool model -force true -ddl template/ddl.sql.tpl -model repository.model.go

`)
	os.Exit(0)
}
