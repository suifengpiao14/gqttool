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
	"github.com/suifengpiao14/gqt/v2"
	"github.com/suifengpiao14/jsonschemaline"
	"github.com/suifengpiao14/templatemap"
)

type EntityElement struct {
	StructName string
	Name       string
	Variables  []*Variable
	FullName   string
	Type       string
	OutEntity  *EntityElement // 输出对象
	//ImplementTplEntityInterface bool // 是否需要实现 gqt.TplEntityInterface 接口
}

func GetSamePrefixEntityElements(prefix string, entityElementList []*EntityElement) (samePrefixEntityElementList []*EntityElement) {
	samePrefixEntityElementList = make([]*EntityElement, 0)
	for _, entityElement := range entityElementList {
		if strings.HasPrefix(entityElement.Name, prefix) {
			samePrefixEntityElementList = append(samePrefixEntityElementList, entityElement)
		}
	}
	return samePrefixEntityElementList
}

const STRUCT_DEFINE_NANE_FORMAT = "%sEntity"

// SQLEntity 根据数据表ddl和sql tpl 生成 sql tpl 调用的输入、输出实体
func SQLEntity(sqltplDefineText *TPLDefineText, tableList []*Table) (entityStruct string, err error) {
	entityElement, err := SQLEntityElement(sqltplDefineText, tableList)
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

func SQLEntityElement(sqltplDefineText *TPLDefineText, tableList []*Table) (entityElement *EntityElement, err error) {
	variableList := ParsSqlTplVariable(sqltplDefineText)
	variableList, err = FormatVariableType(variableList, tableList)
	if err != nil {
		return nil, err
	}
	columnArr := ParseSQLSelectColumn(sqltplDefineText.Text)
	columnList, err := FormatColumn(columnArr, tableList)
	if err != nil {
		return nil, err
	}
	camelName := sqltplDefineText.FullnameCamel()
	outName := fmt.Sprintf("%sOut", camelName)
	entityElement = &EntityElement{
		Name:       camelName,
		Type:       sqltplDefineText.Type(),
		StructName: fmt.Sprintf(STRUCT_DEFINE_NANE_FORMAT, camelName),
		Variables:  variableList,
		FullName:   sqltplDefineText.Fullname(),
		OutEntity: &EntityElement{
			Name:       outName,
			Type:       sqltplDefineText.Type(),
			StructName: fmt.Sprintf(STRUCT_DEFINE_NANE_FORMAT, camelName),
			Variables:  columnList,
			FullName:   fmt.Sprintf("%s%sOut", sqltplDefineText.Namespace, sqltplDefineText.Name),
		},
	}
	return entityElement, nil
}

func SqlTplDefineVariable2lineschema(id string, variables []*Variable, direction string) (lineschema string, err error) {
	arr := make([]string, 0)
	if direction == jsonschemaline.LINE_SCHEMA_DIRECTION_IN {
		arr = append(arr, fmt.Sprintf("version=http://json-schema.org/draft-07/schema,id=input,direction=%s", direction))
	} else if direction == jsonschemaline.LINE_SCHEMA_DIRECTION_OUT {
		arr = append(arr, fmt.Sprintf("version=http://json-schema.org/draft-07/schema,id=output,direction=%s", direction))
	}
	for _, v := range variables {
		if v.FieldName == "" { // 过滤匿名字段
			continue
		}
		kvArr := make([]string, 0)

		kvArr = append(kvArr, fmt.Sprintf("fullname=%s", v.FieldName))
		dst := ""
		src := ""
		format := v.Validate.Format
		if direction == jsonschemaline.LINE_SCHEMA_DIRECTION_IN {
			dst = v.Name //此处使用驼峰,v.FieldName 被改成蛇型了
		} else if direction == jsonschemaline.LINE_SCHEMA_DIRECTION_OUT {
			src = v.Validate.DataPathSrc
		}

		if dst != "" {
			kvArr = append(kvArr, fmt.Sprintf("dst=%s", dst))
		}
		if src != "" {
			kvArr = append(kvArr, fmt.Sprintf("src=%s", src))
		}
		if format != "" {
			kvArr = append(kvArr, fmt.Sprintf("format=%s", format))
		}
		kvArr = append(kvArr, "required")

		line := strings.Join(kvArr, ",")
		arr = append(arr, line)
	}
	lineschema = strings.Join(arr, "\n")
	return lineschema, err
}

func SqlTplDefineVariable2Jsonschema(id string, variables []*Variable) (jsonschemaOut string, err error) {
	properties := orderedmap.New()
	//{"$schema":"http://json-schema.org/draft-07/schema#","type":"object","properties":{},"required":[]}
	schema := jsonschema.Schema{
		Version:    "http://json-schema.org/draft-07/schema#",
		Type:       "object",
		ID:         jsonschema.ID(id),
		Properties: properties,
	}
	names := make([]string, 0)
	for _, v := range variables {
		if v.FieldName == "" { // 过滤匿名字段
			continue
		}

		name := v.FieldName
		subSchema := v.Validate
		subSchema.TypeValue = v.Type
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
	Defines   TPLDefineTextList
}

func (s *SQLTplNamespace) String() string { // 这个将第一次模板解析输出的内容，合并成字符串，然后解析出{{define "xxx"}}{{end}}模板
	tplArr := make([]string, 0)
	for _, define := range s.Defines {
		tplArr = append(tplArr, define.Text)
	}
	str := strings.Join(tplArr, gqt.EOF)
	tplDefineList := ManualParseDefine(str, "", gqt.LeftDelim, gqt.RightDelim)
	tplDefineList = tplDefineList.UniqueItems() // 去重
	newTplArr := make([]string, 0)
	for _, tplDefineText := range tplDefineList {
		newTplArr = append(newTplArr, tplDefineText.Text)
	}
	out := strings.Join(newTplArr, gqt.EOF)
	return out
}

func (s *SQLTplNamespace) Filename() (out string) {
	out = SnakeCase(s.Namespace)
	return
}

func ManualParseDirTplDefine(tplDir string, namespaceSuffix string, leftDelim string, rightDelim string) (tplDefineList TPLDefineTextList, err error) {
	allFileList, err := gqt.GetTplFilesByDir(tplDir, namespaceSuffix)
	if err != nil {
		return
	}
	tplDefineList = make(TPLDefineTextList, 0)
	for _, filename := range allFileList {
		b, err := os.ReadFile(filename)
		if err != nil {
			return nil, err
		}
		namespace := gqt.FileName2Namespace(filename, tplDir)
		subTplDefineList := ManualParseDefine(string(b), namespace, leftDelim, rightDelim)
		tplDefineList = append(tplDefineList, subTplDefineList...)
	}
	return
}

func ManualParseDefine(content string, namespace string, leftDelim string, rightDelim string) (tplDefineList TPLDefineTextList) {
	// 解析文本
	delim := leftDelim + "define "
	delimLen := len(delim)
	content = gqt.TrimSpaces(content) // 去除开头结尾的非有效字符
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

	tplDefineList = TPLDefineTextList{}

	// 格式化
	for _, tpl := range defineList {
		name, err := GetDefineName(tpl)
		if err != nil {
			panic(err)
		}

		tplDefineText := &TPLDefineText{
			Name:      name,
			Namespace: namespace,
			Text:      tpl,
		}
		tplDefineList = append(tplDefineList, tplDefineText)
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

func FormatColumn(columns []string, tableList []*Table) (variables Variables, err error) {
	allVariables := make(Variables, 0)
	tableColumnMap := make(map[string]*Column)
	columnNameMap := make(map[string]string)
	for _, table := range tableList {
		for _, column := range table.Columns { //todo 多表字段相同，类型不同时，只会取列表中最后一个
			tableColumnMap[column.CamelName] = column
			lname := strings.ToLower(column.CamelName)
			columnNameMap[lname] = column.CamelName // 后续用于检测模板变量拼写错误
			variable := &Variable{
				Name:      column.Name,
				FieldName: column.Name,
				Type:      column.Type,
			}
			allVariables = append(allVariables, variable)
		}

	}
	variables = make(Variables, 0)
	for _, name := range columns { // 使用数据库字段定义校正变量类型
		if name == "*" {
			sort.Sort(allVariables)
			return allVariables, nil
		}
		variable := &Variable{}
		tableColumn, ok := tableColumnMap[name]
		if ok {
			variable.Name = name
			variable.FieldName = name
			variable.Type = tableColumn.Type
			variables = append(variables, variable)
		}
	}
	sort.Sort(variables)
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
			gqt.TplEmptyEntity
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
	Validate   templatemap.Schema // 验证
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

// UniqueItems 去重
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

func ParseCurlTplVariable(tplDefineText *TPLDefineText) (variableList Variables) {
	content := gqt.ToEOF(tplDefineText.Content()) // 转换换行符
	tplVariableList := ParseTplVariable(content, tplDefineText.Namespace)
	variableList = append(variableList, tplVariableList...)

	if tplDefineText.Type() == TPL_DEFINE_TYPE_CURL_RESPONSE { // parse curl response variable ,curl response 直接采用复制，所以确保 response body 本身符合go语法
		index := strings.Index(content, gqt.HTTP_HEAD_BODY_DELIM)
		if index < 0 {
			return
		}
		body := content[index+len(gqt.HTTP_HEAD_BODY_DELIM):]
		body = strings.ReplaceAll(body, gqt.LeftDelim, "")
		body = strings.ReplaceAll(body, gqt.RightDelim, "")
		variable := &Variable{
			Namespace: tplDefineText.Namespace,
			Name:      "",
			FieldName: "",
			Type:      body,
		}
		variableList = append(variableList, variable)
		return
	}
	return
}

func ParseSQLSelectColumn(sql string) []string {
	grep := regexp.MustCompile(`(?i)select(.+)from`)
	match := grep.FindAllStringSubmatch(sql, -1)
	if len(match) < 1 {
		return make([]string, 0)
	}
	fieldStr := match[0][1]
	out := strings.Split(gqt.StandardizeSpaces(fieldStr), ",")
	return out
}

func ParsSqlTplVariable(tplDefineText *TPLDefineText) (variableList Variables) {
	content := tplDefineText.Content()
	subVariableList := ParseTplVariable(content, tplDefineText.Namespace)
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
		variable.Namespace = tplDefineText.Namespace
		variableList = append(variableList, &variable)
		pos += len(variable.Name)
		byteArr = byteArr[pos:]
	}
	limitVariabeList := GetLimitVariable(content, tplDefineText.Namespace)
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

func GetDefineName(tplDefineText string) (defineName string, err error) {
	delim := []byte("{{define \"")
	tplDefineByte := []byte(tplDefineText)
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

func GetDefineType(tplDefineText string) (defineName string, err error) {
	delim := []byte("{{define \"")
	tplDefineByte := []byte(tplDefineText)
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
