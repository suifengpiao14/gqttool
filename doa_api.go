package gqttool

import (
	"fmt"
	"strings"

	"github.com/invopop/jsonschema"
)

type SchemaProperty struct {
	Name   string
	Type   string
	Schema jsonschema.Schema
}

func GenerateExec(defineName string, table *Table, relationEntityStructList []*EntityElement) (exec string, validateSchema string, outputSchema string, err error) { //简单的getSet 以及调用当前模板
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
			if v.Type == "int" { //接口字段统一使用string，这里增加数组验证
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
			fn := "getSetValue"
			if v.Type == "int" {
				fn = "getSetValueNumber"
			}
			strArr = append(strArr, fmt.Sprintf(`{{%s . "%s" "input.%s"}}`, fn, v.FieldName, cloneVariable.FieldName))
		}
		if entityStruct.Type != TPL_DEFINE_TYPE_TEXT {
			strArr = append(strArr, fmt.Sprintf(`{{execSQLTpl . "%s"}}`, entityStruct.Name))
		}
	}
	strArr = append(strArr, "{{end}}")
	exec = strings.Join(strArr, "\n")
	validateSchema, err = SqlTplDefineVariable2Jsonschema(defineName, properties)
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
	outputSchema, err = SqlTplDefineVariable2Jsonschema(defineName, outputVariables)
	if err != nil {
		return "", "", "", err
	}

	return exec, validateSchema, outputSchema, nil
}
