package gqttool

import (
	"fmt"
	"strings"
)

func GenerateExec(defineName string, relationEntityStructList []*EntityElement) (exec string, validateSchema string, err error) { //简单的getSet 以及调用当前模板
	variableList := make([]*Variable, 0)
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
					fieldValidate.Pattern = `\d+`
				}
			}
			variableMap[v.FieldName] = v.FieldName
			cloneVariable := &Variable{ // 复制一份，不改原有的
				Namespace:  v.Namespace,
				Name:       v.Name,
				FieldName:  v.FieldName,
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
			variableList = append(variableList, cloneVariable)
			fn := "getSetValue"
			if v.Type == "int" {
				fn = "getSetValueInt"
			}
			strArr = append(strArr, fmt.Sprintf(`{{%s . "%s" "input.%s"}}`, fn, v.FieldName, cloneVariable.FieldName))
		}
		if entityStruct.Type != TPL_DEFINE_TYPE_TEXT {
			strArr = append(strArr, fmt.Sprintf(`{{execSQLTpl . "%s"}}`, entityStruct.Name))
		}
	}
	strArr = append(strArr, "{{end}}")
	exec = strings.Join(strArr, "\n")
	validateSchema, err = SqlTplDefine2Jsonschema(defineName, variableList)
	if err != nil {
		return "", "", err
	}
	return exec, validateSchema, nil
}
