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

	"tplGetByPrimaryKey":        TplGetByPrimaryKey,
	"tplGetAllByPrimaryKeyList": TplGetAllByPrimaryKeyList,
	"tplPaginateWhere":          TplPaginateWhere,
	"tplPaginateTotal":          TplPaginateTotal,
	"tplPaginate":               TplPaginate,
	"tplInsert":                 TplInsert,
	"tplUpdate":                 TplUpdate,
	"tplDel":                    TplDel,
}

var TplDefineNameWithTableName bool

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
	prefix := ""
	if TplDefineNameWithTableName {
		prefix = table.TableNameCamel()
	}
	name := fmt.Sprintf("%sGetAllBy%sList", prefix, primaryKeyCamel)
	tpl = fmt.Sprintf("{{define \"%s\"}}\nselect * from `%s`  where `%s` in ({{in . .%sList}})", name, table.TableName, table.PrimaryKey, primaryKeyCamel)
	if table.DeleteColumn != "" {
		tpl = fmt.Sprintf("%s  and `%s` is null", tpl, table.DeleteColumn)
	}
	tpl = tpl + ";\n{{end}}\n"
	return
}

func TplGetByPrimaryKey(data interface{}) (tpl string, err error) {
	table, err := convertData2Table(data)
	if err != nil {
		return
	}
	primaryKeyCamel := table.PrimaryKeyCamel()
	prefix := ""
	if TplDefineNameWithTableName {
		prefix = table.TableNameCamel()
	}
	name := fmt.Sprintf("%sGetBy%s", prefix, primaryKeyCamel)
	tpl = fmt.Sprintf("{{define \"%s\"}}\nselect * from `%s`  where `%s`=:%s", name, table.TableName, table.PrimaryKey, primaryKeyCamel)
	if table.DeleteColumn != "" {
		tpl = fmt.Sprintf("%s  and `%s` is null", tpl, table.DeleteColumn)
	}
	tpl = tpl + ";\n{{end}}\n"
	return
}

func tplPaginateWhereName(tableNameCamel string) string {
	prefix := ""
	if TplDefineNameWithTableName {
		prefix = tableNameCamel
	}
	return fmt.Sprintf("%sPaginateWhere", prefix)
}

func TplPaginateWhere(data interface{}) (tpl string, err error) {
	table, err := convertData2Table(data)
	if err != nil {
		return
	}
	tpl = fmt.Sprintf("{{define \"%s\"}}\n  ", tplPaginateWhereName(table.TableNameCamel()))

	tpl = tpl + "\n{{end}}\n"
	return
}

func TplPaginateTotal(data interface{}) (tpl string, err error) {
	table, err := convertData2Table(data)
	if err != nil {
		return
	}
	prefix := ""
	if TplDefineNameWithTableName {
		prefix = table.TableNameCamel()
	}
	name := fmt.Sprintf("%sPaginateTotal", prefix)
	tpl = fmt.Sprintf("{{define \"%s\"}}\nselect count(*) as `count` from `%s`  where 1=1 {{template \"%s\" .}} ", name, table.TableName, tplPaginateWhereName(table.TableNameCamel()))
	if table.DeleteColumn != "" {
		tpl = fmt.Sprintf("%s  and `%s` is null", tpl, table.DeleteColumn)
	}
	tpl = tpl + ";\n{{end}}\n"
	return
}

func TplPaginate(data interface{}) (tpl string, err error) {
	table, err := convertData2Table(data)
	if err != nil {
		return
	}
	prefix := ""
	if TplDefineNameWithTableName {
		prefix = table.TableNameCamel()
	}
	name := fmt.Sprintf("%sPaginate", prefix)
	tpl = fmt.Sprintf("{{define \"%s\"}}\nselect * from `%s`  where 1=1 {{template \"%s\" .}} ", name, table.TableName, tplPaginateWhereName(table.TableNameCamel()))
	if table.DeleteColumn != "" {
		tpl = fmt.Sprintf("%s  and `%s` is null", tpl, table.DeleteColumn)
	}
	updatedAtColumn := table.UpdatedAtColumn()
	if updatedAtColumn != nil {
		tpl = fmt.Sprintf(" %s order by `%s` desc ", tpl, updatedAtColumn.Name)
	}
	tpl = fmt.Sprintf(" %s limit :Offset,:Limit ", tpl)
	tpl = tpl + ";\n{{end}}\n"
	return
}

func TplInsert(data interface{}) (tpl string, err error) {
	table, err := convertData2Table(data)
	if err != nil {
		return
	}
	prefix := ""
	if TplDefineNameWithTableName {
		prefix = table.TableNameCamel()
	}
	name := fmt.Sprintf("%sInsert", prefix)
	columns := make([]string, 0)
	values := make([]string, 0)
	for _, column := range table.Columns {
		if isIgnoreColumn(column, table) {
			continue
		}
		columns = append(columns, Backquote(column.Name))
		values = append(values, ":"+column.CamelName)
	}

	columnStr := strings.Join(columns, ",")
	valueStr := strings.Join(values, ",")
	tpl = fmt.Sprintf("{{define \"%s\"}}\ninsert into `%s` (%s)values\n (%s);\n{{end}}\n", name, table.TableName, columnStr, valueStr)

	return
}

func TplUpdate(data interface{}) (tpl string, err error) {
	table, err := convertData2Table(data)
	if err != nil {
		return
	}
	prefix := ""
	if TplDefineNameWithTableName {
		prefix = table.TableNameCamel()
	}
	name := fmt.Sprintf("%sUpdate", prefix)
	updataList := make([]string, 0)
	for _, column := range table.Columns {
		if isIgnoreColumn(column, table) {
			continue
		}
		str := fmt.Sprintf("{{if .%s}} {{$preComma.PreComma}} `%s`=:%s {{end}} ", column.CamelName, column.Name, column.CamelName)
		updataList = append(updataList, str)
	}
	updateTpl := strings.Join(updataList, "\n")
	tpl = fmt.Sprintf("{{define \"%s\"}}\n{{$preComma:=newPreComma}}\n update `%s` set %s where `%s`=:%s;\n{{end}}\n", name, table.TableName, updateTpl, table.PrimaryKey, table.PrimaryKeyCamel())
	return
}

func TplDel(data interface{}) (tpl string, err error) {
	table, err := convertData2Table(data)
	if err != nil {
		return
	}
	prefix := ""
	if TplDefineNameWithTableName {
		prefix = table.TableNameCamel()
	}
	name := fmt.Sprintf("%sDel", prefix)
	tpl = fmt.Sprintf("{{define \"%s\"}}\nupdate `%s` set `%s`={{currentTime .}} where `%s`=:%s;\n{{end}}\n", name, table.TableName, table.DeleteColumn, table.PrimaryKey, table.PrimaryKeyCamel())
	return
}
