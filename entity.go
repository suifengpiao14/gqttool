package gqttool

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"text/template"

	"github.com/pkg/errors"
	"github.com/suifengpiao14/gqt/v2"
)

type EntityElement struct {
	Tables      []*Table
	Name        string
	VariableMap map[string]*Variable
	FullName    string
}

//RepositoryEntity 根据数据表ddl和sql tpl 生成 sql tpl 调用的输入、输出实体
func RepositoryEntity(sqlTplDefine *SQLTPLDefine, tableList []*Table) (entityStruct string, err error) {
	variableMap := ParsSqlTplVariable(sqlTplDefine.TPL)

	structName := sqlTplDefine.FullNameCamel
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

type TalbeNameVariable struct {
	Update string `parser:"(?=...)'update ' @String ' '"`
	From   string `parser:"| (?=...)'from ' @Ident ' '"`
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
	if matchArr != nil {
		for _, matchs := range matchArr {
			matcheList = append(matcheList, matchs[1:]...) // index 0 为匹配对象
		}
	}
	return
}

type SQLTPLDefine struct {
	Name          string
	FullNameCamel string
	Namespace     string
	TPL           string
}

func ParseDirSqlTplDefine(sqlTplDir string) (sqlTplDefineList []*SQLTPLDefine, err error) {
	pattern := fmt.Sprintf("%s/*%s", strings.TrimRight(sqlTplDir, "/"), gqt.Suffix)
	allFileList, err := filepath.Glob(pattern)
	if err != nil {
		return
	}
	for _, filename := range allFileList {
		if strings.HasSuffix(filename, "ddl.sql.tpl") { //skep ddl file
			continue
		}
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
			namespace := gqt.FileName2Namespace(filename, sqlTplDir, gqt.Suffix)
			sqlTplDefine := &SQLTPLDefine{
				Name:          name,
				FullNameCamel: fmt.Sprintf("%s%s", ToCamel(namespace), ToCamel(name)),
				Namespace:     namespace,
				TPL:           sqlTpl,
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
		StructName: entityElement.Name,
		FullName:   entityElement.FullName,
		Attributes: make(Variables, 0),
	}
	tableColumnMap := make(map[string]*Column)
	for _, table := range entityElement.Tables {
		for _, column := range table.Columns { //todo 多表字段相同，类型不同时，只会取列表中最后一个
			tableColumnMap[column.CamelName] = column
		}

	}

	for name, variable := range entityElement.VariableMap { // 使用数据库字段定义校正变量类型
		tableColumn, ok := tableColumnMap[name]
		if ok {
			variable.Type = tableColumn.Type
			entityTplData.Attributes = append(entityTplData.Attributes, variable)
			continue
		}
		// 不属于数据表字段变量，直接添加
		entityTplData.Attributes = append(entityTplData.Attributes, variable)
	}
	sort.Sort(entityTplData.Attributes)
	return
}
func EntityTpl() (tpl string) {
	tpl = `
		type {{.StructName}} struct{
			{{range .Attributes }}
				{{.Name}} {{.Type}} 
			{{end}}
		}

		func (t *{{.StructName}}) TplName() string{
			return "{{.FullName}}"
		}
		func (t *{{.StructName}}) TplInput() interface{}{
			return t
		}

	`
	return
}

type Variable struct {
	Name       string
	Type       string
	AllowEmpty bool
}

// 按照 Person.Age 从大到小排序
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

func ParsSqlTplVariable(sqlTpl string) (variableMap map[string]*Variable) {
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
		pos += len(variable.Name)
		byteArr = byteArr[pos:]
	}
	limitVariabeMap := GetLimitVariable(sqlTpl)
	for name, variable := range limitVariabeMap {
		variableMap[name] = variable
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
