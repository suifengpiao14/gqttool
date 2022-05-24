package gqttool

import (
	"bytes"
	"fmt"
	"text/template"
)

//CURLEntity curl 请求体
func CURLEntity(curlTplDefine *TPLDefineText, curlTplDefineRelationList TPLDefineTextList) (entityStruct string, err error) {
	variableList := ParseCurlTplVariable(curlTplDefine)
	aliasVariableList := make([]*Variable, 0)
	for _, variable := range variableList {
		if curlTplDefineRelationList.IsDefineNameCamel(variable.Name) {
			// 增加一个简称变量，方便模板中使用简称引用
			fieldType := fmt.Sprintf(STRUCT_DEFINE_NANE_FORMAT, variable.FullnameCamel())
			variable.Type = fieldType
			alias := &Variable{ //复制一份，作为隐式引用属性
				Name:      variable.Name,
				Namespace: variable.Namespace,
				Type:      fieldType,
			}
			variable.Type = fmt.Sprintf("*%s", variable.Type) // 别名没有实际数据传递需求，基本使用TplName()获取模板名称，所以使用引用即可，不需要初始化对象
			aliasVariableList = append(aliasVariableList, alias)
		}
	}
	variableList = append(variableList, aliasVariableList...)
	tableList := make([]*Table, 0)
	variableList, err = FormatVariableType(variableList, tableList)
	if err != nil {
		return
	}
	camelName := curlTplDefine.FullnameCamel()
	entityElement := &EntityElement{
		Name:       camelName,
		Type:       curlTplDefine.Type(),
		StructName: fmt.Sprintf(STRUCT_DEFINE_NANE_FORMAT, camelName),
		Variables:  variableList,
		FullName:   curlTplDefine.Fullname(),
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
