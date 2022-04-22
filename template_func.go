package gqttool

import (
	"fmt"
	"reflect"
	"strings"
	"text/template"

	"github.com/pkg/errors"
	"github.com/suifengpiao14/gqt/v2"
	"github.com/suifengpiao14/gqt/v2/gqttpl"
)

var MetaTemplatefuncMap = template.FuncMap{
	"zeroTime":      gqt.ZeroTime,
	"currentTime":   gqt.CurrentTime,
	"permanentTime": gqt.PermanentTime,
	"contains":      strings.Contains,
	"toCamel":       gqttpl.ToCamel,
	"toLowerCamel":  gqttpl.ToLowerCamel,
	"snakeCase":     gqttpl.SnakeCase,
}

func convertData2Table(data interface{}) (table *Table, err error) {
	var ok bool
	for {
		table, ok = data.(*Table)
		if ok {
			break
		}
		rt := reflect.TypeOf(data)
		if rt.Kind() == reflect.Ptr {
			data = reflect.ValueOf(data).Elem().Interface()
			continue
		}
		// 走到这里，说明数据类型不对
		err = errors.Errorf("expected *gqttool.Table; got %#v ", data)
		return nil, err
	}
	return
}

func TplGetAllByPrimaryKeyList(data interface{}) (tpl string, err error) {
	table, err := convertData2Table(data)
	if err != nil {
		return
	}
	primaryKeyCamel := table.PrimaryKeyCamel()
	name := fmt.Sprintf("GetAllBy%sList", primaryKeyCamel)
	tpl = fmt.Sprintf("{{define \"%s\"}}\nselect * from `%s`  where `%s` in ({{in . .%sList}})", name, table.TableName, table.PrimaryKey, primaryKeyCamel)
	if table.DeleteColumn != "" {
		tpl = fmt.Sprintf("%s  and `%s` is null", tpl, table.DeleteColumn)
	}
	tpl = tpl + ";\n{{end}}\n"
	return
}
