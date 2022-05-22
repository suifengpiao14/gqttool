package gqttool

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
	"text/template"

	"github.com/iancoleman/orderedmap"
	"github.com/invopop/jsonschema"
	"github.com/pkg/errors"
	"github.com/suifengpiao14/gqt/v2/gqttpl"
)

type EntityElement struct {
	StructName string
	Name       string
	Variables  []*Variable
	FullName   string
	Type       string
	//ImplementTplEntityInterface bool // 是否需要实现 gqttpl.TplEntityInterface 接口
}

const STRUCT_DEFINE_NANE_FORMAT = "%sEntity"

//SQLEntity 根据数据表ddl和sql tpl 生成 sql tpl 调用的输入、输出实体
func SQLEntity(sqlTplDefine *gqttpl.TPLDefine, tableList []*Table) (entityStruct string, err error) {
	entityElement, err := SQLEntityElement(sqlTplDefine, tableList)
	if err != nil {
		return "", err
	}
	tp, err := template.New("").Parse(InputEntityTpl())
	if err != nil {
		return
	}
	buf := new(bytes.Buffer)
	err = tp.Execute(buf, entityElement)
	if err != nil {
		return
	}
	entityStruct = buf.String()
	return
}

func SQLEntityElement(sqlTplDefine *gqttpl.TPLDefine, tableList []*Table) (entityElement *EntityElement, err error) {
	variableList := ParsSqlTplVariable(sqlTplDefine)
	variableList, err = FormatVariableType(variableList, tableList)
	if err != nil {
		return nil, err
	}
	camelName := sqlTplDefine.FullnameCamel()
	entityElement = &EntityElement{
		Name:       camelName,
		Type:       sqlTplDefine.Type(),
		StructName: fmt.Sprintf(STRUCT_DEFINE_NANE_FORMAT, camelName),
		//ImplementTplEntityInterface: sqlTplDefine.ISSQL(),
		Variables: variableList,
		FullName:  sqlTplDefine.Fullname(),
	}
	return entityElement, nil
}

func SqlTplDefine2Jsonschema(entityElement *EntityElement) (jsonschemaOut string, err error) {
	properties := orderedmap.New()
	schema := jsonschema.Schema{
		ID:         jsonschema.ID(entityElement.FullName),
		Properties: properties,
	}
	names := make([]string, 0)
	for _, v := range entityElement.Variables {
		name := fmt.Sprintf("%s%s", strings.ToLower(v.FieldName[0:1]), v.FieldName[1:])
		subSchema := jsonschema.Schema{
			Type: v.Type,
		}
		properties.Set(name, subSchema)
		names = append(names, name)
	}
	schema.Required = names
	b, err := schema.MarshalJSON()
	jsonschemaOut = string(b)
	return jsonschemaOut, err
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
	Table     *Table
	Defines   []*gqttpl.TPLDefine
}

func (s *SQLTplNamespace) String() string {
	tplArr := make([]string, 0)
	for _, define := range s.Defines {
		tplArr = append(tplArr, define.Output)
	}
	str := strings.Join(tplArr, gqttpl.EOF)
	tplDefineList := ParseDefine(str, "", gqttpl.LeftDelim, gqttpl.RightDelim)
	tplDefineList = tplDefineList.UniqueItems() // 去重
	newTplArr := make([]string, 0)
	for _, tplDefine := range tplDefineList {
		newTplArr = append(newTplArr, tplDefine.Output)
	}
	out := strings.Join(newTplArr, gqttpl.EOF)
	return out
}

func (s *SQLTplNamespace) Filename() (out string) {
	out = SnakeCase(s.Namespace)
	return
}

func ParseDirTplDefine(tplDir string, namespaceSuffix string, leftDelim string, rightDelim string) (tplDefineList []*gqttpl.TPLDefine, err error) {
	allFileList, err := gqttpl.GetTplFilesByDir(tplDir, namespaceSuffix)
	if err != nil {
		return
	}
	tplDefineList = make([]*gqttpl.TPLDefine, 0)
	for _, filename := range allFileList {
		b, err := os.ReadFile(filename)
		if err != nil {
			return nil, err
		}
		namespace := gqttpl.FileName2Namespace(filename, tplDir)
		subTplDefineList := ParseDefine(string(b), namespace, leftDelim, rightDelim)
		tplDefineList = append(tplDefineList, subTplDefineList...)
	}
	return
}

func ParseDefine(content string, namespace string, leftDelim string, rightDelim string) (tplDefineList gqttpl.TPLDefineList) {
	// 解析文本
	delim := leftDelim + "define "
	delimLen := len(delim)
	content = gqttpl.TrimSpaces(content) // 去除开头结尾的非有效字符
	defineList := make([]string, 0)
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
		} else {
			break
		}
	}

	tplDefineList = gqttpl.TPLDefineList{}

	// 格式化
	for _, tpl := range defineList {
		name, err := GetDefineName(tpl)
		if err != nil {
			panic(err)
		}

		tplDefine := &gqttpl.TPLDefine{
			LeftDelim:  leftDelim,
			RightDelim: rightDelim,
			Name:       name,
			Namespace:  namespace,
			Output:     tpl,
			Input:      nil,
		}
		tplDefineList = append(tplDefineList, tplDefine)
	}

	return
}

type EntityTplData struct {
	StructName                  string
	FullName                    string
	ImplementTplEntityInterface bool
	Attributes                  Variables
}

func FormatVariableType(variableList []*Variable, tableList []*Table) (varaibles Variables, err error) {
	varaibles = make(Variables, 0)
	tableColumnMap := make(map[string]*Column)
	columnNameMap := make(map[string]string)
	for _, table := range tableList {
		for _, column := range table.Columns { //todo 多表字段相同，类型不同时，只会取列表中最后一个
			tableColumnMap[column.CamelName] = column
			lname := strings.ToLower(column.CamelName)
			columnNameMap[lname] = column.CamelName // 后续用于检测模板变量拼写错误
		}

	}
	for _, variable := range variableList { // 使用数据库字段定义校正变量类型
		name := variable.Name
		tableColumn, ok := tableColumnMap[name]
		if ok {
			variable.Type = tableColumn.Type
			varaibles = append(varaibles, variable)
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
				varaibles = append(varaibles, variable)
				continue
			}
		}
		// 处理 后缀带类型的变量
		typ := variableSuffix2Type(name)
		if typ != "" {
			variable.Type = typ
			varaibles = append(varaibles, variable)
			continue
		}

		lname := strings.ToLower(name)
		if columnName, ok := columnNameMap[lname]; ok { // 检测模板中大小写拼写错误
			err = errors.Errorf("spelling mistake: have %s, want %s", name, columnName)
			return
		}
		// 不属于数据表字段变量，直接添加
		varaibles = append(varaibles, variable)
	}
	sort.Sort(varaibles)
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

func InputEntityTpl() (tpl string) {
	tpl = `
		type {{.StructName}} struct{
			{{range .Variables }}
				{{.FieldName}} {{.Type}} {{.Tag}} 
			{{end}}
			gqttpl.TplEmptyEntity
		}

		func (t *{{.StructName}}) TplName() string{
			return "{{.FullName}}"
		}

		func (t *{{.StructName}}) TplType() string{
			return "{{.Type}}"
		}
	`
	return
}

type Variable struct {
	Namespace  string
	Name       string // 模板中的原始名称
	FieldName  string // 当变量作为结构体的字段时，字段名称
	Type       string
	Tag        string
	AllowEmpty bool
}

func (v *Variable) FullnameCamel() (fullnameCamel string) {
	fullname := fmt.Sprintf("%s_%s", strings.ReplaceAll(v.Namespace, ".", "_"), v.Name)
	fullnameCamel = ToCamel(fullname)
	return
}
func (v *Variable) Fullname() (fullname string) {
	fullname = fmt.Sprintf("%s.%s", v.Namespace, v.Name)
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

//UniqueItems 去重
func (v Variables) UniqueItems() (uniq []*Variable) {
	vmap := make(map[string]*Variable)
	for _, variable := range v {
		vmap[variable.Name] = variable
	}
	uniq = make([]*Variable, 0)
	for _, variable := range vmap {
		uniq = append(uniq, variable)
	}
	return
}

func ParseTplVariable(tpl string, namespace string) (variableList Variables) {
	variableList = make([]*Variable, 0)
	byteArr := []byte(tpl)
	// template 模板变量提取
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

	// parse define variable
	for _, item := range itemArr {
		variable, _ := parsePrefixVariable(item, byte('.'))
		if variable.Name != "" {
			variable.Namespace = namespace
			variable.FieldName = variable.Name
			variableList = append(variableList, &variable)

		}
	}

	// parse sub define variable
	templateNameList := GetTemplateNames(tpl)
	for _, templateName := range templateNameList {
		templateName = ToCamel(templateName)
		variable := Variable{
			Namespace:  namespace,
			Name:       templateName,
			AllowEmpty: false,
		}
		variable.Type = fmt.Sprintf(STRUCT_DEFINE_NANE_FORMAT, variable.FullnameCamel())
		variableList = append(variableList, &variable)
	}

	variableList = variableList.UniqueItems()
	return
}

func ParseCurlTplVariable(tplDefine *gqttpl.TPLDefine) (variableList Variables) {
	content := gqttpl.ToEOF(tplDefine.Content()) // 转换换行符
	tplVariableList := ParseTplVariable(content, tplDefine.Namespace)
	variableList = append(variableList, tplVariableList...)

	if tplDefine.Type() == gqttpl.TPL_DEFINE_TYPE_CURL_RESPONSE { // parse curl response variable ,curl response 直接采用复制，所以确保 response body 本身符合go语法
		index := strings.Index(content, gqttpl.HTTP_HEAD_BODY_DELIM)
		if index < 0 {
			return
		}
		body := content[index+len(gqttpl.HTTP_HEAD_BODY_DELIM):]
		body = strings.ReplaceAll(body, tplDefine.LeftDelim, "")
		body = strings.ReplaceAll(body, tplDefine.RightDelim, "")
		variable := &Variable{
			Namespace: tplDefine.Namespace,
			Name:      "",
			FieldName: "",
			Type:      body,
		}
		variableList = append(variableList, variable)
		return
	}
	return
}

func ParsSqlTplVariable(tplDefine *gqttpl.TPLDefine) (variableList Variables) {
	content := tplDefine.Content()
	subVariableList := ParseTplVariable(content, tplDefine.Namespace)
	variableList = append(variableList, subVariableList...)
	byteArr := []byte(content)

	// parse sql variable
	sqlVariableDelim := byte(':')

	for {
		variable, pos := parsePrefixVariable(byteArr, sqlVariableDelim)
		if variable.Name == "" {
			break
		}
		variable.FieldName = variable.Name
		variable.Namespace = tplDefine.Namespace
		variableList = append(variableList, &variable)
		pos += len(variable.Name)
		byteArr = byteArr[pos:]
	}
	limitVariabeList := GetLimitVariable(content, tplDefine.Namespace)
	variableList = append(variableList, limitVariabeList...)
	variableList = variableList.UniqueItems()
	return
}

func GetLimitVariable(sqlTpl string, namespace string) (variableList []*Variable) {
	variableList = make([]*Variable, 0)
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
		variable.FieldName = variable.Name
		variable.Type = "int"
		variable.AllowEmpty = false
		variable.Namespace = namespace
		variableList = append(variableList, &variable)
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

func GetDefineName(tplDefine string) (defineName string, err error) {
	delim := []byte("{{define \"")
	tplDefineByte := []byte(tplDefine)
	index := bytes.Index(tplDefineByte, delim)
	nameByte := make([]byte, 0)
	if index >= 0 {
		index += len(delim)
		for i := index; i < len(tplDefineByte); i++ {
			c := tplDefineByte[i]
			if c != '"' {
				nameByte = append(nameByte, tplDefineByte[i])
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

func GetDefineType(tplDefine string) (defineName string, err error) {
	delim := []byte("{{define \"")
	tplDefineByte := []byte(tplDefine)
	index := bytes.Index(tplDefineByte, delim)
	nameByte := make([]byte, 0)
	if index >= 0 {
		index += len(delim)
		for i := index; i < len(tplDefineByte); i++ {
			c := tplDefineByte[i]
			if c != '"' {
				nameByte = append(nameByte, tplDefineByte[i])
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
