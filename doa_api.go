package gqttool

import (
	"fmt"
	"strings"

	"github.com/invopop/jsonschema"
	"github.com/suifengpiao14/jsonschemaline"
)

type SchemaProperty struct {
	Name   string
	Type   string
	Schema jsonschema.Schema
}

func GenerateExec(defineName string, table *Table, relationEntityStructList []*EntityElement) (exec string, input string, out string, err error) { //简单的getSet 以及调用当前模板
	properties := make([]*Variable, 0)
	strArr := make([]string, 0)
	strArr = append(strArr, fmt.Sprintf(`{{define "%s"}}`, defineName))
	variableMap := make(map[string]string)
	for _, entityStruct := range relationEntityStructList {
		for _, v := range entityStruct.Variables {
			if v.FieldName == "" { // 过滤匿名属性
				continue
			}
			if _, ok := variableMap[v.FieldName]; ok {
				continue
			}
			fieldValidate := v.Validate
			if v.Type == "int" { //接口字段统一使用string，这里增加数字验证
				if fieldValidate.Pattern == "" {
					fieldValidate.Format = `number`
				}
			}
			variableMap[v.FieldName] = v.FieldName
			fieldName := v.FieldName
			column := table.GetColumnByCamelName(fieldName)
			if table.DatabaseConfig.ColumnPrefix != "" && column != nil {
				idx := strings.Index(v.FieldName, table.DatabaseConfig.ColumnPrefix)
				if idx > -1 {
					fieldName = v.FieldName[idx+1:]
				}
			}
			cloneVariable := &Variable{ // 复制一份，不改原有的
				Namespace:  v.Namespace,
				Name:       v.Name,
				FieldName:  fieldName,
				Type:       "string", //接口统一使用string类型
				Validate:   fieldValidate,
				Tag:        v.Tag,
				AllowEmpty: v.AllowEmpty,
			}
			cloneVariable.FieldName = SnakeCase(cloneVariable.FieldName)
			// 解决 HTTPStatus,APPID 写法
			wordArr := strings.Split(cloneVariable.FieldName, "_")
			for i, word := range wordArr {
				if i == 0 {
					continue // 首字母不大写
				}
				wordArr[i] = fmt.Sprintf("%s%s", strings.ToUpper(word[:1]), word[1:])
			}
			cloneVariable.FieldName = strings.Join(wordArr, "")
			properties = append(properties, cloneVariable)
		}
		if entityStruct.Type != TPL_DEFINE_TYPE_TEXT {
			strArr = append(strArr, fmt.Sprintf(`{{execSQLTpl . "%s"}}`, entityStruct.Name))
		}
	}
	strArr = append(strArr, "{{end}}")
	exec = strings.Join(strArr, "\n")
	input, err = SqlTplDefineVariable2lineschema(defineName, properties, jsonschemaline.LINE_SCHEMA_DIRECTION_IN)
	if err != nil {
		return "", "", "", err
	}

	outputVariables := make(Variables, 0)
	outputVariableMap := make(map[string]string)
	for _, inputStruct := range relationEntityStructList {
		outputStruct := inputStruct.OutEntity
		for _, v := range outputStruct.Variables {
			if v.FieldName == "" { // 过滤匿名属性
				continue
			}
			if _, ok := outputVariableMap[v.FieldName]; ok {
				continue
			}
			fieldValidate := v.Validate
			outputVariableMap[v.FieldName] = v.FieldName
			fieldName := v.FieldName
			column := table.GetColumnByCamelName(fieldName)
			if table.DatabaseConfig.ColumnPrefix != "" && column != nil {
				idx := strings.Index(v.FieldName, table.DatabaseConfig.ColumnPrefix)
				if idx > -1 {
					fieldName = v.FieldName[idx+1:]
				}
			}
			fieldValidate.DataPathSrc = fmt.Sprintf("%s.#.%s", outputStruct.Name, v.FieldName)
			cloneVariable := &Variable{ // 复制一份，不改原有的
				Namespace:  v.Namespace,
				Name:       v.Name,
				FieldName:  fieldName,
				Type:       "string", //接口统一使用string类型
				Validate:   fieldValidate,
				Tag:        v.Tag,
				AllowEmpty: v.AllowEmpty,
			}
			cloneVariable.FieldName = SnakeCase(cloneVariable.FieldName)
			// 解决 HTTPStatus,APPID 写法
			wordArr := strings.Split(cloneVariable.FieldName, "_")
			for i, word := range wordArr {
				if i == 0 {
					continue // 首字母不大写
				}
				wordArr[i] = fmt.Sprintf("%s%s", strings.ToUpper(word[:1]), word[1:])
			}
			cloneVariable.FieldName = strings.Join(wordArr, "")
			outputVariables = append(outputVariables, cloneVariable)
		}
	}
	out, err = SqlTplDefineVariable2lineschema(defineName, outputVariables, jsonschemaline.LINE_SCHEMA_DIRECTION_OUT)
	if err != nil {
		return "", "", "", err
	}

	return exec, input, out, nil
}
func GenerateTengoApi(defineName string, table *Table, relationEntityStructList []*EntityElement) (mainScript string, inputSchema string, outSchema string, err error) { //简单的getSet 以及调用当前模板
	properties := make([]*Variable, 0)
	strArr := make([]string, 0)
	variableMap := make(map[string]string)
	for _, entityStruct := range relationEntityStructList {
		for _, v := range entityStruct.Variables {
			if v.FieldName == "" { // 过滤匿名属性
				continue
			}
			if _, ok := variableMap[v.FieldName]; ok {
				continue
			}
			fieldValidate := v.Validate
			if v.Type == "int" { //接口字段统一使用string，这里增加数字验证
				if fieldValidate.Pattern == "" {
					fieldValidate.Format = `number`
				}
			}
			variableMap[v.FieldName] = v.FieldName
			fieldName := v.FieldName
			column := table.GetColumnByCamelName(fieldName)
			if table.DatabaseConfig.ColumnPrefix != "" && column != nil {
				idx := strings.Index(v.FieldName, table.DatabaseConfig.ColumnPrefix)
				if idx > -1 {
					fieldName = v.FieldName[idx+1:]
				}
			}
			cloneVariable := &Variable{ // 复制一份，不改原有的
				Namespace:  v.Namespace,
				Name:       v.Name,
				FieldName:  fieldName,
				Type:       "string", //接口统一使用string类型
				Validate:   fieldValidate,
				Tag:        v.Tag,
				AllowEmpty: v.AllowEmpty,
			}
			cloneVariable.FieldName = SnakeCase(cloneVariable.FieldName)
			// 解决 HTTPStatus,APPID 写法
			wordArr := strings.Split(cloneVariable.FieldName, "_")
			for i, word := range wordArr {
				if i == 0 {
					continue // 首字母不大写
				}
				wordArr[i] = fmt.Sprintf("%s%s", strings.ToUpper(word[:1]), word[1:])
			}
			cloneVariable.FieldName = strings.Join(wordArr, "")
			properties = append(properties, cloneVariable)
		}
		switch entityStruct.Type {
		case TPL_DEFINE_TYPE_SQL_SELECT, TPL_DEFINE_TYPE_SQL_UPDATE, TPL_DEFINE_TYPE_SQL_INSERT:
			outName := fmt.Sprintf("%sOut", entityStruct.Name)
			strArr = append(strArr, fmt.Sprintf(`%s:=execSQLTpl(ctx,"%s",input)`, outName, entityStruct.Name))

		}
	}
	mainScript = strings.Join(strArr, "\n")
	inputSchema, err = SqlTplDefineVariable2lineschema(defineName, properties, jsonschemaline.LINE_SCHEMA_DIRECTION_IN)
	if err != nil {
		return "", "", "", err
	}

	outputVariables := make(Variables, 0)
	outputVariableMap := make(map[string]string)
	for _, inputStruct := range relationEntityStructList {
		outputStruct := inputStruct.OutEntity
		for _, v := range outputStruct.Variables {
			if v.FieldName == "" { // 过滤匿名属性
				continue
			}
			if _, ok := outputVariableMap[v.FieldName]; ok {
				continue
			}
			fieldValidate := v.Validate
			outputVariableMap[v.FieldName] = v.FieldName
			fieldName := v.FieldName
			column := table.GetColumnByCamelName(fieldName)
			if table.DatabaseConfig.ColumnPrefix != "" && column != nil {
				idx := strings.Index(v.FieldName, table.DatabaseConfig.ColumnPrefix)
				if idx > -1 {
					fieldName = v.FieldName[idx+1:]
				}
			}
			fieldValidate.DataPathSrc = fmt.Sprintf("%s.#.%s", outputStruct.Name, v.FieldName)
			cloneVariable := &Variable{ // 复制一份，不改原有的
				Namespace:  v.Namespace,
				Name:       v.Name,
				FieldName:  fieldName,
				Type:       "string", //接口统一使用string类型
				Validate:   fieldValidate,
				Tag:        v.Tag,
				AllowEmpty: v.AllowEmpty,
			}
			cloneVariable.FieldName = SnakeCase(cloneVariable.FieldName)
			// 解决 HTTPStatus,APPID 写法
			wordArr := strings.Split(cloneVariable.FieldName, "_")
			for i, word := range wordArr {
				if i == 0 {
					continue // 首字母不大写
				}
				wordArr[i] = fmt.Sprintf("%s%s", strings.ToUpper(word[:1]), word[1:])
			}
			cloneVariable.FieldName = strings.Join(wordArr, "")
			outputVariables = append(outputVariables, cloneVariable)
		}
	}
	outSchema, err = SqlTplDefineVariable2lineschema(defineName, outputVariables, jsonschemaline.LINE_SCHEMA_DIRECTION_OUT)
	if err != nil {
		return "", "", "", err
	}

	return mainScript, inputSchema, outSchema, nil
}
