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

	"github.com/suifengpiao14/errorformatter"
	"github.com/suifengpiao14/gqt/v2/gqttpl"
	"github.com/suifengpiao14/gqttool"
)

func main() {
	var (
		metaDir        string
		modelFilename  string
		sqlDir         string
		curlDir        string
		entityFilename string
		force          bool
		err            error
	)
	modelCmd := flag.NewFlagSet("model", flag.ExitOnError)
	modelCmd.Usage = helpModel
	SQLCmd := flag.NewFlagSet("sql", flag.ExitOnError)
	SQLCmd.Usage = helpSQL
	sqlEntityCmd := flag.NewFlagSet("entity", flag.ExitOnError)
	sqlEntityCmd.Usage = helpSQLEntity
	curlEntityCmd := flag.NewFlagSet("entity", flag.ExitOnError)
	curlEntityCmd.Usage = helpCURLEntity
	modelCmd.BoolVar(&force, "force", false, "overwrite exist file")
	modelCmd.StringVar(&metaDir, "metaDir", "template/meta", "ddl/config file dir")
	modelCmd.StringVar(&modelFilename, "model", "model.gen.go", "model file name")

	SQLCmd.BoolVar(&force, "force", false, "overwrite exist file")
	SQLCmd.StringVar(&metaDir, "metaDir", "template/meta", "ddl/config/sql template file dir")
	SQLCmd.StringVar(&sqlDir, "sqlDir", "template/sql", "sql template dir")

	sqlEntityCmd.BoolVar(&force, "force", false, "overwrite exist file")
	sqlEntityCmd.StringVar(&sqlDir, "sqlDir", "template/sql", "sql template dir")
	sqlEntityCmd.StringVar(&entityFilename, "entity", "sql.entity.gen.go", "entity file")
	curlEntityCmd.BoolVar(&force, "force", false, "overwrite exist file")
	curlEntityCmd.StringVar(&sqlDir, "sqlDir", "template/sql", "sql template dir")
	curlEntityCmd.StringVar(&entityFilename, "entity", "sql.entity.gen.go", "entity file")
	testing.Init()

	if len(os.Args) < 3 {
		help()
	}
	args := os.Args[2:]
	switch os.Args[1] {
	case "model":
		modelCmd.Parse(args)
		err = runCmdModel(metaDir, modelFilename, force)
		if err != nil {
			panic(err)
		}

	case "sql":
		SQLCmd.Parse(args)
		err = runCmdSQL(metaDir, sqlDir, force)
		if err != nil {
			panic(err)
		}

	case "entity":
		sqlEntityCmd.Parse(args)
		err = runCmdSQLEntity(sqlDir, entityFilename, force)
		if err != nil {
			panic(err)
		}
		err = runCmdCURLEntity(sqlDir, entityFilename, force)
		if err != nil {
			panic(err)
		}

	default:
		help()
	}
}

func runCmdModel(metaDir string, modelFile string, force bool) (err error) {
	repo := gqttool.NewRepositoryMeta()
	errChain := errorformatter.NewErrorChain()
	var content string
	var modelStructList gqttool.ModelStructList
	errChain.SetError(repo.AddByDir(metaDir, gqttool.MetaTemplatefuncMap)).
		Run(func() (err error) {
			modelStructList, err = GenerateTableModel(repo)
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
	sort.Sort(modelStructList) // 排序后写入文件

	for _, modelStruct := range modelStructList {
		contentArr = append(contentArr, modelStruct.TPL)
	}
	content = strings.Join(contentArr, "\n")
	err = saveFile(modelFile, content, force)
	if err != nil {
		return
	}
	return
}

func runCmdSQL(metaDir string, sqlDir string, force bool) (err error) {
	repo := gqttool.NewRepositoryMeta()
	errChain := errorformatter.NewErrorChain()
	var sqlTplNamespaceList []*gqttool.SQLTplNamespace
	errChain.SetError(repo.AddByDir(metaDir, gqttool.MetaTemplatefuncMap)).
		Run(func() (err error) {
			sqlTplNamespaceList, err = GenerateSQL(repo)
			return
		})

	err = errChain.Error()
	if err != nil {
		return
	}
	for _, sqlTplNamespace := range sqlTplNamespaceList {
		filename := fmt.Sprintf("%s/%s.gen.sql.tpl", sqlDir, sqlTplNamespace.Filename())
		err = saveFile(filename, sqlTplNamespace.String(), force)
		if err != nil {
			return
		}
	}
	return
}

func runCmdSQLEntity(sqlDir string, entityFilename string, force bool) (err error) {
	repo := gqttool.NewRepositoryMeta()
	errChain := errorformatter.NewErrorChain()
	var entityList = make([]string, 0)
	var sqlTplDefineList = make([]*gqttpl.TPLDefine, 0)
	errChain.SetError(repo.AddByDir(sqlDir, gqttool.MetaTemplatefuncMap)).
		Run(func() (err error) {
			sqlTplDefineList, err = gqttool.ParseDirTplDefine(sqlDir, gqttpl.SQLNamespaceSuffix)
			return
		}).
		Run(func() (err error) {
			entityList, err = GenerateSQLEntity(repo, sqlTplDefineList)
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
	importLine := `import "github.com/suifengpiao14/gqt/v2/gqttpl"`
	contentArr = append(contentArr, packageLine)
	contentArr = append(contentArr, importLine)
	contentArr = append(contentArr, entityList...)
	content := strings.Join(contentArr, "\n")
	err = saveFile(entityFilename, content, force)
	if err != nil {
		return
	}
	return
}

func runCmdCURLEntity(curlDir string, entityFilename string, force bool) (err error) {
	repo := gqttool.NewRepositoryMeta()
	errChain := errorformatter.NewErrorChain()
	var entityList = make([]string, 0)
	var tplDefineList = make([]*gqttpl.TPLDefine, 0)
	errChain.SetError(repo.AddByDir(curlDir, gqttool.MetaTemplatefuncMap)).
		Run(func() (err error) {
			tplDefineList, err = gqttool.ParseDirTplDefine(curlDir, gqttpl.CURLNamespaceSuffix)
			return
		}).
		Run(func() (err error) {
			entityList, err = GenerateCURLEntity(repo, tplDefineList)
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
	importLine := `import "github.com/suifengpiao14/gqt/v2/gqttpl"`
	contentArr = append(contentArr, packageLine)
	contentArr = append(contentArr, importLine)
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

func GenerateCURLEntity(rep *gqttool.RepositoryMeta, curlTplDefineList []*gqttpl.TPLDefine) (entityList []string, err error) {
	entityList = make([]string, 0)
	for _, sqlDefineTpl := range curlTplDefineList {
		entityStruct, err := gqttool.CURLEntity(sqlDefineTpl)
		if err != nil {
			return nil, err
		}
		entityList = append(entityList, entityStruct)
	}
	return
}

func GenerateSQLEntity(rep *gqttool.RepositoryMeta, sqlTplDefineList []*gqttpl.TPLDefine) (entityList []string, err error) {
	entityList = make([]string, 0)
	ddlDefineList, err := rep.GetDDLTPLDefine()
	if err != nil {
		return
	}
	tableCfg, err := rep.GetDatabaseConfig()
	if err != nil {
		return
	}
	ddlList := make([]string, 0)
	for _, ddlDefine := range ddlDefineList {
		ddlList = append(ddlList, ddlDefine.Output)
	}
	tableList, err := gqttool.GenerateTable(ddlList, tableCfg)
	if err != nil {
		return
	}
	tableMap := make(map[string]*gqttool.Table)
	for _, table := range tableList {
		tableMap[table.TableName] = table
	}

	for _, sqlDefineTpl := range sqlTplDefineList {
		tableNameList, err := gqttool.ParseSQLTPLTableName(sqlDefineTpl.Output)
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

		entityStruct, err := gqttool.SQLEntity(sqlDefineTpl, relationTableList)
		if err != nil {
			return nil, err
		}
		entityList = append(entityList, entityStruct)

	}

	return
}

func GenerateSQL(rep *gqttool.RepositoryMeta) (sqlTplNamespaceList []*gqttool.SQLTplNamespace, err error) {
	ddlDefineList, err := rep.GetDDLTPLDefine()
	if err != nil {
		return
	}
	repCfg, err := rep.GetDatabaseConfig()
	if err != nil {
		return nil, err
	}
	ddlList := make([]string, 0)
	for _, ddlDefine := range ddlDefineList {
		ddlList = append(ddlList, ddlDefine.Output)
	}
	tableList, err := gqttool.GenerateTable(ddlList, repCfg)
	if err != nil {
		return
	}

	sqlTplNamespaceList = make([]*gqttool.SQLTplNamespace, 0)
	for _, table := range tableList {
		sqlTplNamespace := &gqttool.SQLTplNamespace{}
		sqlTplNamespace.Namespace = table.TableNameCamel()
		sqlTplDefineList, err := gqttool.GenerateSQLTpl(table, rep)
		if err != nil {
			return nil, err
		}
		sqlTplNamespace.Defines = append(sqlTplNamespace.Defines, sqlTplDefineList...)
		sqlTplNamespaceList = append(sqlTplNamespaceList, sqlTplNamespace)
	}

	return
}
func GenerateTableModel(rep *gqttool.RepositoryMeta) (modelStructList []*gqttool.ModelStruct, err error) {
	ddlDefineList, err := rep.GetDDLTPLDefine()
	if err != nil {
		return
	}
	repCfg, err := rep.GetDatabaseConfig()
	if err != nil {
		return
	}
	ddlList := make([]string, 0)
	for _, ddlDefine := range ddlDefineList {
		ddlList = append(ddlList, ddlDefine.Output)
	}
	tableList, err := gqttool.GenerateTable(ddlList, repCfg)
	if err != nil {
		return
	}
	modelStructList, err = gqttool.GenerateModel(tableList)
	return
}

func helpModel() {
	fmt.Fprint(os.Stderr, `gqttool model is generate go struct from mysql ddl

Usage:
  gqttool model -metaDir metaDir -model modelFilename -force true

Flags:
  -force overwrite exists file
  -metaDir
        template meta dir

  -model
        model filename

Example:

  gqttool model -metaDir template/meta -model model.gen.go -force true

`)
	os.Exit(0)
}

func helpSQL() {
	fmt.Fprint(os.Stderr, `gqttool sql is  generation  sql from mysql ddl

Usage:
  gqttool sql -metaDir metaDirReadDir -sqlDir sqlTplSaveDir -force true
  
Flags:
  -force overwrite exists file
  -metaDir
        template meta dir

  -sqlDir
        save sqlTpl file dir

Example:

  gqttool sql -metaDir template/meta -sqlDir template/sql -force true

`)
	os.Exit(0)
}

func helpSQLEntity() {
	fmt.Fprint(os.Stderr, `gqttool entity is  generation sql template inpur args entity

Usage:
  gqttool entity -sqlDir sqlsqlDir -entity entityFilename -force true
  

Flags:
 -force overwrite exists file 
 -sqlDir 
		sqlTpl file dir
 -entity 
		sqlTpl  entity filename

Example:

  gqttool entity -sqlDir template/sql -entity entity.gen.go -force true

`)
	os.Exit(0)
}

func help() {
	fmt.Fprint(os.Stderr, `gqttool is the code generation tool for the gqt package.

Usage:
  gqttool model  -metaDir metaDir -model modelFilename -force true
  gqttool sql -metaDir metaDir -sqlDir sqlTplSaveDir -force true
  gqttool entity  -sqlDir sqlsqlDir -entity entityFilename -force true
  
Commands:
  model
  		Generate go struct from  mysql ddl
  sql
        Generate sql sql.tpl from mysql ddl and meta template
  entity
  		Generate sql.tpl input entity from mysql sqlsqlDir

Flags:
  -force overwrite exists file
  -metaDir
        template meta dir
  -model
         model filename
 -sqlDir 
		sqlTpl file dir
 -entity 
		sqlTpl  argument entity filename

Example:

  gqttool model  -metaDir template/meta -model model.gen.go -force true

`)
	os.Exit(0)
}
