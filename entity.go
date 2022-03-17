package gqttool

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
	"text/template"

	"github.com/pkg/errors"
	"github.com/suifengpiao14/gqt/v2/gqttpl"
)

type EntityElement struct {
	Tables      []*Table
	Name        string
	VariableMap map[string]*Variable
	FullName    string
}

//RepositoryEntity 根据数据表ddl和sql tpl 生成 sql tpl 调用的输入、输出实体
func RepositoryEntity(sqlTplDefine *gqttpl.TPLDefine, tableList []*Table) (entityStruct string, err error) {
	variableMap := ParsSqlTplVariable(sqlTplDefine.Output, sqlTplDefine.Namespace)

	structName := sqlTplDefine.FullnameCamel()
	entityElement := &EntityElement{
		Tables:      tableList,
		VariableMap: variableMap,
		Name:        structName,
		FullName:    fmt.Sprintf("%s.%s", sqlTplDefine.Namespace, sqlTplDefine.Name),
	}
	entityTplData, err := FormatEntityData(entityElement)
	if err != nil {
		return
	}

	tp, err := template.New("").Parse(EntityTpl())
	if err != nil {
		return
	}
	buf := new(bytes.Buffer)
	err = tp.Execute(buf, entityTplData)
	if err != nil {
		return
	}
	entityStruct = buf.String()
	return
}

func ParseSQLTPLTableName(sqlTpl string) (tableList []string, err error) {

	updateDelim := "update `?(\\w+)`?"
	updateMatchArr, err := regexpMatch(sqlTpl, updateDelim)
	if err != nil {
		return
	}
	selectDelim := "from `?(\\w+)`?"
	fromMatchArr, err := regexpMatch(sqlTpl, selectDelim)
	if err != nil {
		return
	}
	insertDelim := "into `?(\\w+)`?"
	insertMatchArr, err := regexpMatch(sqlTpl, insertDelim)
	if err != nil {
		return
	}

	tableList = append(tableList, updateMatchArr...)
	tableList = append(tableList, fromMatchArr...)
	tableList = append(tableList, insertMatchArr...)
	return
}

func regexpMatch(s string, delim string) (matcheList []string, err error) {
	reg := regexp.MustCompile(delim)
	if reg == nil {
		err = errors.Errorf("regexp.MustCompile %s is nil", delim)
		return
	}
	matchArr := reg.FindAllStringSubmatch(s, -1)
	for _, matchs := range matchArr {
		matcheList = append(matcheList, matchs[1:]...) // index 0 为匹配对象
	}
	return
}

type SQLTplNamespace struct {
	Namespace string
	Defines   []*gqttpl.TPLDefine
}

func (s *SQLTplNamespace) String() (out string) {
	tplArr := make([]string, 0)
	for _, define := range s.Defines {
		tplArr = append(tplArr, define.Output)
	}
	out = strings.Join(tplArr, "\n")
	return
}

func (s *SQLTplNamespace) Filename() (out string) {
	out = SnakeCase(s.Namespace)
	return
}

func ParseDirSqlTplDefine(sqlTplDir string) (sqlTplDefineList []*gqttpl.TPLDefine, err error) {
	allFileList, err := gqttpl.GetTplFilesByDir(sqlTplDir, gqttpl.SQLNamespaceSuffix)
	if err != nil {
		return
	}
	for _, filename := range allFileList {
		b, err := os.ReadFile(filename)
		if err != nil {
			return nil, err
		}
		sqlTplList := ParseDefine(string(b))
		for _, sqlTpl := range sqlTplList {
			name, err := GetDefineName(sqlTpl)
			if err != nil {
				return nil, err
			}
			namespace := gqttpl.FileName2Namespace(filename, sqlTplDir)
			sqlTplDefine := &gqttpl.TPLDefine{
				Name:      name,
				Namespace: namespace,
				Output:    sqlTpl,
				Input:     nil,
			}
			sqlTplDefineList = append(sqlTplDefineList, sqlTplDefine)
		}
	}
	return
}

func ParseDefine(content string) (defineList []string) {
	delim := "{{define "
	delimLen := len(delim)
	for {
		index := strings.Index(content, delim)
		if index >= 0 {
			pos := delimLen + index
			nextIndex := strings.Index(content[pos:], delim)
			if nextIndex >= 0 {
				sepPos := pos + nextIndex
				oneDefine := content[:sepPos]
				defineList = append(defineList, oneDefine)
				content = content[sepPos:]
			} else {
				defineList = append(defineList, content)
				break
			}
		}
	}
	return
}

type EntityTplData struct {
	StructName string
	FullName   string
	Attributes Variables
}

func FormatEntityData(entityElement *EntityElement) (entityTplData *EntityTplData, err error) {
	entityTplData = &EntityTplData{
		StructName: fmt.Sprintf("%sEntity", entityElement.Name),
		FullName:   entityElement.FullName,
		Attributes: make(Variables, 0),
	}
	tableColumnMap := make(map[string]*Column)
	columnNameMap := make(map[string]string)
	for _, table := range entityElement.Tables {
		for _, column := range table.Columns { //todo 多表字段相同，类型不同时，只会取列表中最后一个
			tableColumnMap[column.CamelName] = column
			lname := strings.ToLower(column.CamelName)
			columnNameMap[lname] = column.CamelName // 后续用于检测模板变量拼写错误
		}

	}

	for name, variable := range entityElement.VariableMap { // 使用数据库字段定义校正变量类型

		tableColumn, ok := tableColumnMap[name]
		if ok {
			variable.Type = tableColumn.Type
			entityTplData.Attributes = append(entityTplData.Attributes, variable)
			continue
		}

		// 处理 表字段 +List 后缀类型
		listSuffix := "List"
		basename := name
		suffix := ""
		if strings.HasSuffix(name, listSuffix) {
			suffix = listSuffix
			basename = strings.TrimRight(name, listSuffix)
		}
		tableColumn, ok = tableColumnMap[basename]
		if ok {
			if suffix == listSuffix {
				variable.Type = fmt.Sprintf("[]%s", tableColumn.Type) // 完善列表类型
				entityTplData.Attributes = append(entityTplData.Attributes, variable)
				continue
			}
		}
		// 处理 后缀带类型的变量
		typ := variableSuffix2Type(name)
		if typ != "" {
			variable.Type = typ
			entityTplData.Attributes = append(entityTplData.Attributes, variable)
			continue
		}

		lname := strings.ToLower(name)
		if columnName, ok := columnNameMap[lname]; ok { // 检测模板中大小写拼写错误
			err = errors.Errorf("spelling mistake: have %s, want %s", name, columnName)
			return
		}
		// 不属于数据表字段变量，直接添加
		entityTplData.Attributes = append(entityTplData.Attributes, variable)
	}
	sort.Sort(entityTplData.Attributes)
	return
}

type VariableSuffixType struct {
	Suffix string
	Type   string
}

type VariableSuffixTypes []*VariableSuffixType

func (v VariableSuffixTypes) Len() int { // 重写 Len() 方法
	return len(v)
}
func (v VariableSuffixTypes) Swap(i, j int) { // 重写 Swap() 方法
	v[i], v[j] = v[j], v[i]
}
func (v VariableSuffixTypes) Less(i, j int) bool { // 重写 Less() 方法， 从长到短排序
	return len(v[i].Suffix) > len(v[j].Suffix)
}

var VariableSuffixTypeList = VariableSuffixTypes{
	&VariableSuffixType{Suffix: "ListInt", Type: "[]int"},
	&VariableSuffixType{Suffix: "ListStr", Type: "[]string"},
	&VariableSuffixType{Suffix: "Str", Type: "string"},
	&VariableSuffixType{Suffix: "Int", Type: "int"},
}

func variableSuffix2Type(variableName string) (typ string) {
	typ = ""
	sort.Sort(VariableSuffixTypeList)
	for _, vs2t := range VariableSuffixTypeList {
		if strings.HasSuffix(variableName, vs2t.Suffix) {
			typ = vs2t.Type
			return // 匹配第一个即返回
		}
	}
	return
}

func EntityTpl() (tpl string) {
	tpl = `
		type {{.StructName}} struct{
			{{range .Attributes }}
				{{if  .IsSubDefine}}
					{{.FullnameCamel}}Entity
				{{else}}
					{{.Name}} {{.Type}} 
				{{end}}
				
			{{end}}
			gqt.DataVolumeMap
		}

		func (t *{{.StructName}}) TplName() string{
			return "{{.FullName}}"
		}

	`
	return
}

type Variable struct {
	Namespace   string
	Name        string
	IsSubDefine bool
	Type        string
	AllowEmpty  bool
}

func (v *Variable) FullnameCamel() (fullnameCamel string) {
	fullname := fmt.Sprintf("%s_%s", strings.ReplaceAll(v.Namespace, ".", "_"), v.Name)
	fullnameCamel = ToCamel(fullname)
	return
}

type Variables []*Variable

func (v Variables) Len() int { // 重写 Len() 方法
	return len(v)
}
func (v Variables) Swap(i, j int) { // 重写 Swap() 方法
	v[i], v[j] = v[j], v[i]
}
func (v Variables) Less(i, j int) bool { // 重写 Less() 方法， 从小到大排序
	return v[i].Name < v[j].Name
}

func ParsSqlTplVariable(sqlTpl string, namespace string) (variableMap map[string]*Variable) {
	variableMap = make(map[string]*Variable)
	byteArr := []byte(sqlTpl)
	leftDelim := byte('{')
	rightDelim := byte('}')
	itemBegin := false
	itemArr := make([][]byte, 0)
	item := make([]byte, 0)
	byteLen := len(byteArr)
	for i := 0; i < byteLen; i++ {
		c := byteArr[i]
		if c == leftDelim && i+1 < byteLen && byteArr[i+1] == leftDelim && !itemBegin {
			itemBegin = true
			item = make([]byte, 0)
			i++
			continue
		}
		if c == rightDelim && i+1 < byteLen && byteArr[i+1] == rightDelim && itemBegin {
			itemBegin = false
			itemArr = append(itemArr, item)
			i++
			continue
		}
		if itemBegin {
			item = append(item, c)
		}
	}
	for _, item := range itemArr {
		variable, _ := parsePrefixVariable(item, byte('.'))
		if variable.Name != "" {
			variable.Namespace = namespace
			variableMap[variable.Name] = &variable

		}

	}

	// parse sql variable
	sqlVariableDelim := byte(':')

	for {
		variable, pos := parsePrefixVariable(byteArr, sqlVariableDelim)
		if variable.Name == "" {
			break
		}
		variableMap[variable.Name] = &variable
		variable.Namespace = namespace
		pos += len(variable.Name)
		byteArr = byteArr[pos:]
	}
	limitVariabeMap := GetLimitVariable(sqlTpl)
	for name, variable := range limitVariabeMap {
		variableMap[name] = variable
	}
	// parse sub define variable
	templateNameList := GetTemplateNames(sqlTpl)
	for _, templateName := range templateNameList {
		templateName = ToCamel(templateName)
		variable := &Variable{
			Namespace:   namespace,
			Name:        templateName,
			Type:        "interface{}",
			IsSubDefine: true,
			AllowEmpty:  false,
		}
		variableMap[templateName] = variable
	}

	return
}

func GetLimitVariable(sqlTpl string) (variableMap map[string]*Variable) {
	variableMap = make(map[string]*Variable)
	index := strings.Index(strings.ToLower(sqlTpl), "limit")
	if index < 0 {
		return
	}
	byteArr := []byte(sqlTpl[index:])
	// parse sql variable
	sqlVariableDelim := byte(':')

	for {
		variable, pos := parsePrefixVariable(byteArr, sqlVariableDelim)
		if variable.Name == "" {
			break
		}
		variable.Type = "int"
		variable.AllowEmpty = false
		variableMap[variable.Name] = &variable
		pos += len(variable.Name)
		byteArr = byteArr[pos:]
	}

	return
}

// 找到第一个变量
func parsePrefixVariable(item []byte, variableStart byte) (variable Variable, pos int) {
	variableBegin := false
	pos = 0
	variableNameByte := make([]byte, 0)
	itemLen := len(item)
	for j := 0; j < itemLen; j++ {
		c := item[j]
		if c == variableStart && j+1 < itemLen && IsNameChar(item[j+1]) { // c=变量标示字符，并且后面跟字符，是变量名的必要条件
			if j == 0 || !IsNameChar(item[j-1]) { // 变量符号开始或者前面不为字符，为变量名的充要条件
				variableBegin = true
				pos = j
				continue
			}
		}
		if variableBegin {
			if IsNameChar(c) {
				variableNameByte = append(variableNameByte, c)
			} else {
				break
			}
		}
	}
	variable = Variable{
		Name:       string(variableNameByte),
		Type:       "interface{}",
		AllowEmpty: true,
	}
	return
}

func GetDefineName(sqlTpl string) (defineName string, err error) {
	delim := []byte("{{define \"")
	sqlTplByte := []byte(sqlTpl)
	index := bytes.Index(sqlTplByte, delim)
	nameByte := make([]byte, 0)
	if index >= 0 {
		index += len(delim)
		for i := index; i < len(sqlTplByte); i++ {
			c := sqlTplByte[i]
			if c != '"' {
				nameByte = append(nameByte, sqlTplByte[i])
			} else {
				break
			}

		}
	}
	defineName = string(nameByte)
	if defineName == "" {
		err = errors.Errorf("define name is empty")
	}
	return
}

func GetTemplateNames(sqlTpl string) (templateNameList []string) {
	templateNameList = make([]string, 0)
	reg := regexp.MustCompile(`\{\{template\W+"(\w+)"`)
	if reg == nil {
		panic("regexp err")
	}
	matches := reg.FindAllStringSubmatch(sqlTpl, -1)
	for _, match := range matches {
		templateNameList = append(templateNameList, match[1])
	}
	return
}

// 判断是否可以作为名称的字符
func IsNameChar(c byte) (yes bool) {
	yes = false
	a := byte('a')
	z := byte('z')
	A := byte('A')
	Z := byte('Z')
	underline := byte('_')
	if (a <= c && c <= z) || (A <= c && c <= Z) || c == underline {
		yes = true
	}
	return
}
