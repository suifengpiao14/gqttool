package main

import (
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"testing"

	"github.com/pkg/errors"
	"github.com/suifengpiao14/errorformatter"
	"github.com/suifengpiao14/gqt/v2/gqttpl"
	"github.com/suifengpiao14/gqttool"
)

func main() {
	var (
		metaDir        string
		sqlFile        string
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
	APISQLCmd := flag.NewFlagSet("apiSql", flag.ExitOnError)
	APISQLCmd.Usage = helpAPISQL
	sqlEntityCmd := flag.NewFlagSet("sqlEntity", flag.ExitOnError)
	sqlEntityCmd.Usage = helpSQLEntity
	curlEntityCmd := flag.NewFlagSet("curlEntity", flag.ExitOnError)
	curlEntityCmd.Usage = helpCURLEntity
	modelCmd.BoolVar(&force, "force", false, "overwrite exist file")
	modelCmd.StringVar(&metaDir, "metaDir", "template/meta", "ddl/config file dir")
	modelCmd.StringVar(&modelFilename, "model", "model.gen.go", "model file name")

	APISQLCmd.BoolVar(&force, "force", false, "overwrite exist file")
	APISQLCmd.StringVar(&metaDir, "metaDir", "template/meta", "ddl/config/sql template file dir")
	APISQLCmd.StringVar(&sqlFile, "sqlFile", "doa.sql", "output sql file")

	SQLCmd.BoolVar(&force, "force", false, "overwrite exist file")
	SQLCmd.StringVar(&metaDir, "metaDir", "template/meta", "ddl/config/sql template file dir")
	SQLCmd.StringVar(&sqlDir, "sqlDir", "template/sql", "sql template dir")

	sqlEntityCmd.BoolVar(&force, "force", false, "overwrite exist file")
	sqlEntityCmd.StringVar(&sqlDir, "sqlDir", "template/sql", "sql template dir")
	sqlEntityCmd.StringVar(&entityFilename, "entity", "sql.entity.gen.go", "entity file")

	curlEntityCmd.BoolVar(&force, "force", false, "overwrite exist file")
	curlEntityCmd.StringVar(&curlDir, "curlDir", "template/sql", "sql template dir")
	curlEntityCmd.StringVar(&entityFilename, "entity", "curl.entity.gen.go", "entity file")
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
	case "doaApi":
		APISQLCmd.Parse(args)
		err = runCmdAPISQL(metaDir, sqlFile, force)
		if err != nil {
			panic(err)
		}

	case "sqlEntity":
		sqlEntityCmd.Parse(args)
		err = runCmdSQLEntity(sqlDir, entityFilename, force)
		if err != nil {
			panic(err)
		}
	case "curlEntity":
		curlEntityCmd.Parse(args)
		err = runCmdCURLEntity(curlDir, entityFilename, force)
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
	sort.Sort(modelStructList) // ?????????????????????

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
		filename := fmt.Sprintf("%s/%s.sql.tpl", sqlDir, sqlTplNamespace.Filename())
		err = saveFile(filename, sqlTplNamespace.String(), force)
		if err != nil {
			return
		}
	}
	return
}

func runCmdAPISQL(metaDir string, apiSQLFile string, force bool) (err error) {
	repo := gqttool.NewRepositoryMeta()
	errChain := errorformatter.NewErrorChain()
	var sqlRaws []string
	errChain.SetError(repo.AddByDir(metaDir, gqttool.MetaTemplatefuncMap)).
		Run(func() (err error) {
			sqlRaw, err := GenerateAPISQL(repo)
			if err != nil {
				return err
			}
			sqlRaws = append(sqlRaws, sqlRaw)
			return nil
		})

	err = errChain.Error()
	if err != nil {
		return
	}
	filename := fmt.Sprintf("%s.sql", apiSQLFile)
	err = saveFile(filename, strings.Join(sqlRaws, gqttpl.EOF), force)
	if err != nil {
		return
	}
	return
}

func runCmdSQLEntity(sqlDir string, entityFilename string, force bool) (err error) {
	repo := gqttool.NewRepositoryMeta()
	errChain := errorformatter.NewErrorChain()
	var entityList = make([]string, 0)
	var sqlTplDefineList = make([]*gqttool.TPLDefineText, 0)
	errChain.SetError(repo.AddByDir(sqlDir, gqttool.MetaTemplatefuncMap)).
		Run(func() (err error) {
			sqlTplDefineList, err = gqttool.ManualParseDirTplDefine(sqlDir, gqttpl.SQLNamespaceSuffix, gqttpl.LeftDelim, gqttpl.RightDelim)
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
	var tplDefineList = make(gqttool.TPLDefineTextList, 0)
	errChain.SetError(repo.AddByDir(curlDir, gqttool.MetaTemplatefuncMap)).
		Run(func() (err error) {
			tplDefineList, err = gqttool.ManualParseDirTplDefine(curlDir, gqttpl.CURLNamespaceSuffix, gqttpl.LeftDelim, gqttpl.RightDelim)
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

const TPL_DEFINE_END = "{{end}}"

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
	defer func() {
		if f != nil {
			f.Close()
		}
	}()
	content = gqttpl.ToEOF(content)
	lineArr := strings.Split(content, gqttpl.EOF)
	newLineArr := make([]string, 0)
	for _, line := range lineArr {
		standLine := gqttpl.StandardizeSpaces(line)
		if standLine == "" {
			continue
		}

		if standLine == TPL_DEFINE_END { // ????????????????????????
			standLine = standLine + gqttpl.EOF
		}
		newLineArr = append(newLineArr, standLine)
	}
	content = strings.Join(newLineArr, gqttpl.EOF)
	_, err = f.WriteString(content)
	if err != nil {
		return
	}

	f.Close()
	f = nil // ?????????????????????defer ?????????
	// ???????????????????????????go??????
	// Format Go source files
	if filepath.Ext(filename) == ".go" {
		if err := gqttool.FinalizeGoSource(filename); err != nil {
			return err
		}
	}
	return
}

func IsExist(path string) bool {
	_, err := os.Stat(path) //os.Stat??????????????????
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

func GenerateCURLEntity(rep *gqttool.RepositoryMeta, curlTplDefineList gqttool.TPLDefineTextList) (entityList []string, err error) {
	entityList = make([]string, 0)
	for _, sqlDefineTpl := range curlTplDefineList {
		entityStruct, err := gqttool.CURLEntity(sqlDefineTpl, curlTplDefineList)
		if err != nil {
			return nil, err
		}
		entityList = append(entityList, entityStruct)
	}
	return
}

func GenerateSQLEntity(rep *gqttool.RepositoryMeta, sqlTplDefineList gqttool.TPLDefineTextList) (entityList []string, err error) {
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
		tableNameList, err := gqttool.ParseSQLTPLTableName(sqlDefineTpl.Text)
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

func GenerateSQLEntityStruct(rep *gqttool.RepositoryMeta, sqlTplDefineList gqttool.TPLDefineTextList) (entityElementList []*gqttool.EntityElement, err error) {
	entityElementList = make([]*gqttool.EntityElement, 0)
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
		tableNameList, err := gqttool.ParseSQLTPLTableName(sqlDefineTpl.Text)
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

		entityStruct, err := gqttool.SQLEntityElement(sqlDefineTpl, relationTableList)
		if err != nil {
			return nil, err
		}
		entityElementList = append(entityElementList, entityStruct)

	}
	return
}

type SourceModel struct {
	SourceId   string
	SourceType string
	Config     string
}

type TemplateModel struct {
	TemplateID  string
	Type        string
	Title       string
	Description string
	SourceID    string
	Version     string
	Tpl         string
}

type APIModel struct {
	APIID       string
	Title       string
	Description string
	Method      string
	Route       string
	TemplateIDs string
	Exec        string
	Input       string
	Output      string
}

const SourceInsertTpl = "insert ignore into `source` (`source_id`,`source_type`,`config`) values('%s','%s','%s');"
const TemplateInsertTpl = "insert ignore into `template` (`template_id`,`type`,`title`,`description`,`source_id`,`tpl`) values('%s','SQL','%s','%s','%s','%s');"
const ApiInsertTpl = "insert ignore into `api` (`api_id`,`title`,`description`,`method`,`route`,`template_ids`,`exec`,`input`,`output`) values('%s','%s','%s','%s','%s','%s','%s','%s','%s');"

func generateTemplateId(dbName string, tableName string, defineName string) (templateId string) {
	templateId = fmt.Sprintf("%s%s%s", dbName, tableName, defineName)
	return templateId
}

func GenerateAPISQL(rep *gqttool.RepositoryMeta) (string, error) {
	sqlTplNamespaceList, err := GenerateSQL(rep)
	if err != nil {
		return "", err
	}
	var sqlRaws []string
	DatabaseConfig := gqttool.DatabaseConfig{
		DatabaseName: "database_db",
		LogLevel:     "debug",
	}
	if len(sqlTplNamespaceList) > 0 {
		table := sqlTplNamespaceList[0].Table
		if table != nil {
			DatabaseConfig = table.DatabaseConfig
		}

	}
	module := DatabaseConfig.DatabaseName
	sourceId := DatabaseConfig.DatabaseName
	moduleCamel := gqttool.ToCamel(module)
	sourceType := "SQL"
	config := fmt.Sprintf(`{"logLevel":"%s","dsn":"root:123456@tcp(mysql_address:3306)/%s?charset=utf8&timeout=1s&readTimeout=5s&writeTimeout=5s&parseTime=False&loc=Local&multiStatements=true","timeout":30}`, DatabaseConfig.LogLevel, DatabaseConfig.DatabaseName)
	sourceInsertSql := fmt.Sprintf(SourceInsertTpl, sourceId, sourceType, config)
	sqlRaws = append(sqlRaws, sourceInsertSql)
	for _, sqlTplNamespace := range sqlTplNamespaceList {
		table := sqlTplNamespace.Table
		entityStructList, err := GenerateSQLEntityStruct(rep, sqlTplNamespace.Defines)
		if err != nil {
			return "", err
		}
		entityStructMap := make(map[string]*gqttool.EntityElement)
		for _, entityStruct := range entityStructList {
			entityStructMap[entityStruct.Name] = entityStruct
		}
		for _, define := range sqlTplNamespace.Defines {
			name, tpl := define.Name, gqttpl.StandardizeSpaces(define.Text)
			defineNameCamel := gqttool.ToCamel(define.Name)
			templatId := generateTemplateId(moduleCamel, table.TableNameCamel(), defineNameCamel)
			tableName := table.TableNameCamel()
			extraTranslatemap := make(map[string]string)
			if table.Comment != "" {
				extraTranslatemap = map[string]string{
					tableName: table.Comment,
				}
			}
			title := fmt.Sprintf("%s%s%s", moduleCamel, tableName, name)
			title = gqttool.Translate(title, extraTranslatemap)
			description := title
			templateInsertSql := fmt.Sprintf(TemplateInsertTpl, templatId, title, description, sourceId, tpl)
			sqlRaws = append(sqlRaws, templateInsertSql)
			defineType := define.Type()
			if defineType == gqttool.TPL_DEFINE_TYPE_SQL_SELECT || defineType == gqttool.TPL_DEFINE_TYPE_SQL_UPDATE || defineType == gqttool.TPL_DEFINE_TYPE_SQL_INSERT {
				entityStruct, ok := entityStructMap[define.FullnameCamel()]
				if !ok {
					err := errors.Errorf("not found tpl define struct:%s", define.FullnameCamel())
					return "", err
				}
				relationEntityStructList := gqttool.GetSamePrefixEntityElements(entityStruct.Name, entityStructList)
				templatIdArr := make([]string, 0)
				for _, relationEntityStruct := range relationEntityStructList {
					tplId := generateTemplateId(moduleCamel, table.TableNameCamel(), gqttool.ToCamel(relationEntityStruct.Name)) // ???????????????ID????????????????????????templateInsertSql ??? `template_id` ??????
					templatIdArr = append(templatIdArr, tplId)
				}
				templatIds := strings.Join(templatIdArr, ",")
				mainName := "main"
				exec, input, output, err := gqttool.GenerateExec(mainName, table, relationEntityStructList)
				if err != nil {
					return "", err
				}
				appId := fmt.Sprintf("%s-%s-%s", module, table.TableName, name)
				tableName := table.TableNameCamel()
				extraTranslatemap := make(map[string]string)
				if table.Comment != "" {
					extraTranslatemap = map[string]string{
						tableName: table.Comment,
					}
				}
				title := fmt.Sprintf("%s%s%s", module, tableName, name)
				title = gqttool.Translate(title, extraTranslatemap)

				apiModel := APIModel{
					APIID:       appId,
					Title:       title,
					Description: title,
					Method:      "POST",
					Route:       fmt.Sprintf("/api/%s/v1/%s/%s", gqttool.SnakeCase(module), gqttool.SnakeCase(table.TableNameCamel()), gqttool.SnakeCase(name)),
					Exec:        exec,
					TemplateIDs: templatIds,
					Input:       input,
					Output:      output,
				}
				apiInsertSql := fmt.Sprintf(ApiInsertTpl, apiModel.APIID, apiModel.Title, apiModel.Description, apiModel.Method, apiModel.Route, apiModel.TemplateIDs, apiModel.Exec, apiModel.Input, apiModel.Output)
				sqlRaws = append(sqlRaws, apiInsertSql)
			}
		}

	}
	return strings.Join(sqlRaws, gqttpl.EOF), nil
}

func Define2Sql(define string) (name string, sql string) {
	define = gqttpl.StandardizeSpaces(define)
	re := regexp.MustCompile(`\{\{define +"(\w+)"\}\}(.+)\{\{end\}\}`)
	matched := re.FindStringSubmatch(define)
	if len(matched) > 2 {
		name = matched[1]
		sql = gqttpl.StandardizeSpaces(matched[2])
		return name, sql
	}
	return "", ""
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
		sqlTplNamespace := &gqttool.SQLTplNamespace{
			Table: table,
		}
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
  gqttool model -metaDir template/meta -model model.gen.go -force true

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
  gqttool sql -metaDir template/meta -sqlDir template/gen -force true
  
Flags:
  -force overwrite exists file
  -metaDir
        template meta dir

  -sqlDir
        save sqlTpl file dir

Example:

  gqttool sql -metaDir template/meta -sqlDir template/gen -force true

`)
	os.Exit(0)
}

func helpAPISQL() {
	fmt.Fprint(os.Stderr, `gqttool doaApi is  generation  data manage plat api sql from mysql ddl

Usage:
  gqttool doaApi -metaDir template/meta -sqlFile doa.sql  -force true
  
Flags:
  -force overwrite exists file
  -metaDir
        template meta dir

  -sqlFile
        save doa api sql filename

Example:

  gqttool doaApi -metaDir template/meta -sqlFile doa.sql  -force true

`)
	os.Exit(0)
}

func helpSQLEntity() {
	fmt.Fprint(os.Stderr, `gqttool sqlEntity is  generation sql template input args entity

Usage:
gqttool sqlEntity -sqlDir template -entity sql.entity.gen.go -force true
  

Flags:
 -force overwrite exists file 
 -sqlDir 
		sqlTpl file dir
 -sqlEntity 
		sqlTpl  entity filename

Example:

  gqttool sqlEntity -sqlDir template -entity sql.entity.gen.go -force true

`)
	os.Exit(0)
}
func helpCURLEntity() {
	fmt.Fprint(os.Stderr, `gqttool curlEntity is  generation curl template input args entity

Usage:
   gqttool curlEntity -curlDir template -entity curl.entity.gen.go -force true
  

Flags:
 -force overwrite exists file 
 -sqlDir 
		sqlTpl file dir
 -curlEntity 
		sqlTpl  entity filename

Example:
  gqttool curlEntity -curlDir template -entity curl.entity.gen.go -force true

`)
	os.Exit(0)
}

func help() {
	fmt.Fprint(os.Stderr, `gqttool is the code generation tool for the gqt package.

Usage:
  gqttool model  -metaDir metaDir -model modelFilename -force true
  gqttool sql -metaDir metaDir -sqlDir sqlTplSaveDir -force true
  gqttool sqlEntity  -sqlDir sqlsqlDir -entity entityFilename -force true
  gqttool curlEntity -curlDir curlDir -entity entityFilename -force true
  gqttool doaApi -metaDir metaDir -sqlFile sqlFile -force true
  
Commands:
  model
  		Generate go struct from  mysql ddl
  sql
        Generate sql sql.tpl from mysql ddl and meta template
  sqlEntity
  		Generate sql.tpl input entity from mysql
  curlEntity
  		Generate curl.tpl input entity from curl
  doaApi
  		Generate doa api  from mysql ddl

Flags:
  -force overwrite exists file
  -metaDir
        template meta dir
  -model
         model filename
 -sqlDir 
		sqlTpl file dir 
 -curlDir 
		curlTpl file dir
 -entity 
		sqlTpl  argument entity filename
 -sqlFile 
 		doa api sql filename

Example:

  gqttool model  -metaDir template/meta -model model.gen.go -force true

`)
	os.Exit(0)
}
