package gqttool

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/pkg/errors"
	"github.com/suifengpiao14/gqt/v2"

	executor "github.com/bytewatch/ddl-executor"
)

//map for converting mysql type to golang types
var typeForMysqlToGo = map[string]string{
	"int":                "int",
	"integer":            "int",
	"tinyint":            "int",
	"smallint":           "int",
	"mediumint":          "int",
	"bigint":             "int",
	"int unsigned":       "int",
	"integer unsigned":   "int",
	"tinyint unsigned":   "int",
	"smallint unsigned":  "int",
	"mediumint unsigned": "int",
	"bigint unsigned":    "int",
	"bit":                "int",
	"bool":               "bool",
	"enum":               "string",
	"set":                "string",
	"varchar":            "string",
	"char":               "string",
	"tinytext":           "string",
	"mediumtext":         "string",
	"text":               "string",
	"longtext":           "string",
	"blob":               "string",
	"tinyblob":           "string",
	"mediumblob":         "string",
	"longblob":           "string",
	"date":               "time.Time", // time.Time or string
	"datetime":           "time.Time", // time.Time or string
	"timestamp":          "time.Time", // time.Time or string
	"time":               "time.Time", // time.Time or string
	"float":              "float64",
	"double":             "float64",
	"decimal":            "float64",
	"binary":             "string",
	"varbinary":          "string",
}

func isTimeMysqlType(mysqlType string) bool {
	timeMap := map[string]string{
		"date":      "date",
		"datetime":  "datetime",
		"timestamp": "timestamp",
		"time":      "time",
	}
	_, ok := timeMap[mysqlType]
	return ok
}
func mysql2GoType(mysqlType string, time2str bool) (goType string, err error) {
	if time2str {
		typeForMysqlToGo["date"] = "string"
		typeForMysqlToGo["datetime"] = "string"
		typeForMysqlToGo["timestamp"] = "string"
		typeForMysqlToGo["time"] = "string"
	}
	index := strings.Index(mysqlType, "(")
	if index > -1 {
		mysqlType = mysqlType[:index]
	}
	goType, ok := typeForMysqlToGo[mysqlType]
	if !ok {
		err = errors.Errorf("mysql2GoType: not found mysql type %s to go type", mysqlType)
	}
	return

}

const (
	DEFAULT_VALUE_CURRENT_TIMESTAMP = "current_timestamp"
	DEFAULT_COLUMN_DELETED_AT       = "deleted_at" // 默认删除列名称
)

type Column struct {
	prefix        string
	CamelName     string
	Name          string
	Type          string
	Comment       string
	Tag           string
	Nullable      bool
	Enums         []string
	AutoIncrement bool
	DefaultValue  string
	OnCreate      bool // 根据数据表ddl及默认 值为current_timestap 判断
	OnUpdate      bool // 根据数据表ddl 配置
	OnDelete      bool // 手动设置
}

// IsDefaultValueCurrentTimestamp 判断默认值是否为自动填充时间
func (c *Column) IsDefaultValueCurrentTimestamp() bool {
	return strings.ToLower(c.DefaultValue) == DEFAULT_VALUE_CURRENT_TIMESTAMP
}

type Table struct {
	TablePrefix  string
	ColumnPrefix string
	TableName    string
	PrimaryKey   string
	DeleteColumn string
	Columns      []*Column
	EnumsConst   map[string]string
}

//CamelName 删除表前缀，转换成 camel 格式
func (t *Table) TableNameCamel() (camelName string) {
	name := t.TableName
	if t.TablePrefix != "" {
		name = strings.TrimLeft(name, t.TablePrefix)
	}
	camelName = ToCamel(name)
	return
}
func (t *Table) PrimaryKeyCamel() (camelName string) {
	primaryKey := t.PrimaryKey
	if t.ColumnPrefix != "" {
		primaryKey = strings.TrimLeft(primaryKey, t.TablePrefix)
	}
	camelName = ToCamel(primaryKey)
	return
}

func (t *Table) CreatedAtColumn() (createdAtColumn *Column) {
	for _, column := range t.Columns {
		if column.OnCreate {
			return column
		}
	}
	return
}

// UpdateAtColumn 获取更新列
func (t *Table) UpdatedAtColumn() (updatedAtColumn *Column) {
	for _, column := range t.Columns {
		if column.OnUpdate {
			return column
		}
	}

	return
}

// DeletedAtColumn 获取删除列
func (t *Table) DeletedAtColumn() (deletedAtColumn *Column) {
	for _, column := range t.Columns {
		if column.OnDelete {
			return column
		}
	}
	return
}

type ModelStruct struct {
	Name string
	TPL  string
}

type ModelStructList []*ModelStruct

func (v ModelStructList) Len() int { // 重写 Len() 方法
	return len(v)
}
func (v ModelStructList) Swap(i, j int) { // 重写 Swap() 方法
	v[i], v[j] = v[j], v[i]
}
func (v ModelStructList) Less(i, j int) bool { // 重写 Less() 方法， 从小到大排序
	return v[i].Name < v[j].Name
}
func GenerateModel(tableList []*Table) (modelStructList []*ModelStruct, err error) {
	modelStructList = make([]*ModelStruct, 0)
	tableTpl := structTpl()
	tl, err := template.New("").Parse(tableTpl)
	if err != nil {
		return
	}

	for i := 0; i < len(tableList); i++ {
		buf := new(bytes.Buffer)
		table := tableList[i]
		err = tl.Execute(buf, table)
		if err != nil {
			return
		}
		modelStruct := &ModelStruct{
			Name: ToCamel(table.TableName),
			TPL:  buf.String(),
		}
		modelStructList = append(modelStructList, modelStruct)
	}
	return
}

func GenerateTable(ddlList []string, tableCfg *gqt.Config) (tables []*Table, err error) {
	tables = make([]*Table, 0)
	conf := executor.NewDefaultConfig()
	inst := executor.NewExecutor(conf)
	databaseName := "test"
	ddlDB := fmt.Sprintf("create database `%s`;use `%s`;", databaseName, databaseName)
	ddlAll := make([]string, 0)
	ddlAll = append(ddlAll, ddlDB)
	ddlAll = append(ddlAll, ddlList...)

	ddls := strings.Join(ddlAll, ";")

	err = inst.Exec(ddls)
	if err != nil {
		return nil, err
	}

	tableList, err := inst.GetTables(databaseName)
	if err != nil {
		return
	}

	for i := 0; i < len(tableList); i++ {
		tableName := tableList[i]
		tableDef, err := inst.GetTableDef(databaseName, tableName)
		if err != nil {
			return nil, err
		}

		table := &Table{
			TablePrefix:  tableCfg.TablePrefix,
			TableName:    tableName,
			Columns:      make([]*Column, 0),
			EnumsConst:   make(map[string]string),
			DeleteColumn: tableCfg.DeletedAtColumn,
		}
		for _, indice := range tableDef.Indices {
			if indice.Name == "PRIMARY" {
				table.PrimaryKey = indice.Columns[0] // 暂时取第一个为主键，不支持多字段主键
			}
		}
		for _, columnDef := range tableDef.Columns {

			goType, err := mysql2GoType(columnDef.Type, true)
			if err != nil {
				return nil, err
			}
			if isTimeMysqlType(columnDef.Type) && strings.Contains(columnDef.Name, "deleted_at") {
				table.DeleteColumn = columnDef.Name
			}

			columnPt := &Column{
				CamelName:     ToCamel(columnDef.Name),
				Name:          columnDef.Name,
				Type:          goType,
				Comment:       columnDef.Comment,
				Nullable:      columnDef.Nullable,
				Tag:           fmt.Sprintf("`json:\"%s\"`", ToLowerCamel(columnDef.Name)),
				Enums:         columnDef.Elems,
				AutoIncrement: columnDef.AutoIncrement,
				DefaultValue:  columnDef.DefaultValue,
				OnUpdate:      columnDef.OnUpdate,
			}
			if len(columnDef.Elems) > 0 {
				prefix := fmt.Sprintf("%s_%s", tableName, columnPt.Name)
				subEnumConst := enumsConst(prefix, columnDef.Elems)
				for key, val := range subEnumConst {
					table.EnumsConst[key] = val
				}
			}
			columnPt.OnCreate = columnPt.IsDefaultValueCurrentTimestamp() && !columnPt.OnUpdate    // 自动填充时间，但是更新时不变，认为是创建时间列
			if table.DeleteColumn == columnPt.Name || columnPt.Name == DEFAULT_COLUMN_DELETED_AT { // 删除记录列，通过配置指定，或者列名称为 DEFAULT_COLUMN_DELETED_AT 的值
				columnPt.OnDelete = true
			}

			table.Columns = append(table.Columns, columnPt)
		}
		tables = append(tables, table)
	}
	return
}

func enumsConst(prefix string, enums []string) (enumsConst map[string]string) {
	enumsConst = make(map[string]string)
	for _, enum := range enums {
		name := fmt.Sprintf("%s_%s", prefix, enum)
		name = strings.ToUpper(name)
		enumsConst[name] = enum
	}

	return
}

func structTpl() string {
	return `
	{{if .EnumsConst}}
	const (
		{{range $key, $value := .EnumsConst}}
			{{$key}}="{{$value}}"
		{{end}}
		)
	{{end}}
	type {{.TableNameCamel}}Model struct{
		{{range .Columns }} 
		// {{.Comment}}
		{{.CamelName}} {{.Type}} {{if .Tag}} {{.Tag}} {{end}}
		{{end}}
	}
	func (t *{{.TableNameCamel}}Model) TableName()string{
		return "{{.TableName}}"
	}
	func (t *{{.TableNameCamel}}Model) PrimaryKey()string{
		return "{{.PrimaryKey}}"
	}
	func (t *{{.TableNameCamel}}Model) PrimaryKeyCamel()string{
		return "{{.PrimaryKeyCamel}}"
	}
	`
}
