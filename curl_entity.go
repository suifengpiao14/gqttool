package gqttool

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/suifengpiao14/gqt/v2/gqttpl"
)

//CURLEntity curl 请求体
func CURLEntity(sqlTplDefine *gqttpl.TPLDefine, curlTplDefineRelationList gqttpl.TPLDefineList) (entityStruct string, err error) {
	variableList := ParseTplVariable(sqlTplDefine.Output, sqlTplDefine.Namespace)
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
			aliasVariableList = append(aliasVariableList, alias)
		}
	}
	variableList = append(variableList, aliasVariableList...)
	structName := sqlTplDefine.FullnameCamel()
	entityElement := &EntityElement{
		Variables: variableList,
		Name:      structName,
		FullName:  fmt.Sprintf("%s.%s", sqlTplDefine.Namespace, sqlTplDefine.Name),
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
