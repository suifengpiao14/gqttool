package gqttool

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/suifengpiao14/gqt/v2/gqttpl"
)

//CURLEntity curl 请求体
func CURLEntity(sqlTplDefine *gqttpl.TPLDefine, curlTplDefineRelationList gqttpl.TPLDefineList) (entityStruct string, err error) {
	variableMap := ParseTplVariable(sqlTplDefine.Output, sqlTplDefine.Namespace)
	newVariableMap := make(map[string]*Variable)
	for variableName, variable := range variableMap {
		if curlTplDefineRelationList.IsDefineNameCamel(variable.Name) {
			variable.IsSubDefine = true
			// 增加一个简称变量，方便模板中使用简称引用
			alias := &Variable{ //variable.IsSubDefine =true ,不能直接使用
				Name:      variable.Name,
				Namespace: variable.Namespace,
			}
			newVariableMap[variable.Name] = alias
		}
		newVariableMap[variableName] = variable
	}
	structName := sqlTplDefine.FullnameCamel()
	entityElement := &EntityElement{
		VariableMap: newVariableMap,
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
