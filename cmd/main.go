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
		ddlFile   string
		modelDir  string
		tplDir    string
		entiryDir string
		err       error
	)
	modelCmd := flag.NewFlagSet("model", flag.ExitOnError)
	modelCmd.Usage = help
	crudCmd := flag.NewFlagSet("crud", flag.ExitOnError)
	crudCmd.Usage = help
	entiryCmd := flag.NewFlagSet("entity", flag.ExitOnError)
	entiryCmd.Usage = help
	modelCmd.StringVar(&ddlFile, "ddl", "template/ddl.sql.tpl", "ddl template file name")
	modelCmd.StringVar(&modelDir, "dir", "model", "ddl template file name")

	crudCmd.StringVar(&ddlFile, "ddl", "template/ddl.sql.tpl", "ddl template file name")
	crudCmd.StringVar(&tplDir, "dir", "template", "sql template dir")

	entiryCmd.StringVar(&tplDir, "tplDir", "template", "sql template dir")
	entiryCmd.StringVar(&entiryDir, "entiryDir", "entity", "entity dir")
	testing.Init()

	if len(os.Args) < 3 {
		help()
	}
	args := os.Args[2:]
	switch os.Args[1] {
	case "model":
		modelCmd.Parse(args)
		err = runCmdModel(ddlFile, modelDir)
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
		entiryCmd.Parse(args)
		err = runCmdEntity(tplDir, entiryDir)
		if err != nil {
			panic(err)
		}

	default:
		help()
	}
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
	basename := path.Base(dstDir)
	packageName := strings.ToLower(strcase.ToLowerCamel(basename))
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

func runCmdEntity(sqlTplDir string, entiryDir string) (err error) {
	repo := gqt.NewRepository()
	errChain := errorformatter.NewErrorChain()
	var entityMap map[string]string
	errChain.SetError(repo.AddByDir(sqlTplDir, gqt.TemplatefuncMap)).
		Run(func() (err error) {
			entityMap, err = GenerateEntity(repo)
			return
		})

	err = errChain.Error()
	if err != nil {
		return
	}
	for name, entity := range entityMap {
		snakeName := strcase.ToSnake(name)
		filename := fmt.Sprintf("%s/%s.entity.go", entiryDir, snakeName)
		err = saveFile(filename, entity)
		if err != nil {
			return
		}
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

func GenerateEntity(rep *gqt.Repository) (entityMap map[string]string, err error) {
	entityMap = make(map[string]string)
	ddlList, err := getDDLFromRepository(rep)
	if err != nil {
		return
	}
	tableList, err := gqttool.GenerateTable(ddlList)
	if err != nil {
		return
	}
	for _, table := range tableList {
		//todo get sqlTpl
		sqlTpl := ""
		entityStruct, err := gqttool.RepositoryEntity(table, sqlTpl)
		if err != nil {
			return nil, err
		}
		entityMap[table.TableNameCamel] = entityStruct
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
	for _, table := range tableList {
		tplMap, err = gqttool.Crud(table)
		if err != nil {
			return
		}
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
  gqttool model  -ddl ddlFilename -dir modelSaveDir
  gqttool crud  -ddl ddlFilename -dir sqlTplSaveDir
  gqttool entiry  -tplDir sqlTplDir -entiryDir entirySaveDir
  
Commands:
  model
  		Generate go struct from  mysql ddl
  crud
        Generate crud sql.tpl from mysql ddl
  version
  		Generate sql.tpl input entiry from mysql sqlTplDir

Flags:
  -ddl
        mysql ddl file path

  -dir
        save model/sqlTpl file path

 -tplDir 
		sqlTpl file dir
 -entityDir 
		sqlTpl  argument entity save dir

Example:

  gqttool model template/ddl.sql.tpl -o template

`)
	os.Exit(1)
}

func help() {
	fmt.Fprint(os.Stderr, `gqttool is the code generation tool for the gqt package.

Usage:
  gqttool model  -ddl ddlFilename -dir modelSaveDir
  gqttool crud  -ddl ddlFilename -dir sqlTplSaveDir
  gqttool entiry  -tplDir sqlTplDir -entiryDir entirySaveDir
  
Commands:
  model
  		Generate go struct from  mysql ddl
  crud
        Generate crud sql.tpl from mysql ddl
  version
  		Generate sql.tpl input entiry from mysql sqlTplDir

Flags:
  -ddl
        mysql ddl file path

  -dir
        save model/sqlTpl file path

 -tplDir 
		sqlTpl file dir
 -entityDir 
		sqlTpl  argument entity save dir

Example:

  gqttool model template/ddl.sql.tpl -o template

`)
	os.Exit(1)
}
