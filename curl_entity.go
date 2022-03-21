package gqttool

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/suifengpiao14/gqt/v2/gqttpl"
)

//CURLEntity curl 请求体
func CURLEntity(sqlTplDefine *gqttpl.TPLDefine) (entityStruct string, err error) {
	variableMap := ParsSqlTplVariable(sqlTplDefine.Output, sqlTplDefine.Namespace)

	structName := sqlTplDefine.FullnameCamel()
	entityElement := &EntityElement{
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
