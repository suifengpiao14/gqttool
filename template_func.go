package gqttool

import (
	"fmt"
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

	"tplGetByPrimaryKey": TplGetByPrimaryKey,
	"tplPaginateWhere":   TplPaginateWhere,
	"tplPaginateTotal":   TplPaginateTotal,
	"tplPaginate":        TplPaginate,
	"tplInsert":          TplInsert,
	"tplUpdate":          TplUpdate,
	"tplDel":             TplDel,
}

func convertData2Table(data interface{}) (table *Table, err error) {
	for {
		dataI, ok := data.(*interface{})
		if ok {
			data = *dataI
		} else {
			break
		}
	}
	table, ok := data.(*Table)
	if !ok {
		err = errors.Errorf("expected *gqttool.Table; got %#v ", data)
		return nil, err
	}
	return
}
func TplGetByPrimaryKey(data interface{}) (tpl string, err error) {
	table, err := convertData2Table(data)
	if err != nil {
		return
	}
	primaryKeyCamel := table.PrimaryKeyCamel()
	name := fmt.Sprintf("GetBy%s", primaryKeyCamel)
	tpl = fmt.Sprintf("{{define \"%s\"}}\nselect * from `%s`  where `%s`=:%s", name, table.TableName, table.PrimaryKey, primaryKeyCamel)
	if table.DeleteColumn != "" {
		tpl = fmt.Sprintf("%s  and `%s` is null", tpl, table.DeleteColumn)
	}
	tpl = tpl + ";\n{{end}}\n"
	return
}

var tplPaginateWhereName = "PaginateWhere"

func TplPaginateWhere(data interface{}) (tpl string, err error) {
	table, err := convertData2Table(data)
	if err != nil {
		return
	}
	tpl = fmt.Sprintf("{{define \"%s\"}}\n  ", tplPaginateWhereName)
	if table.DeleteColumn != "" {
		tpl = fmt.Sprintf("%s  and `%s` is null", tpl, table.DeleteColumn)
	}
	tpl = tpl + ";\n{{end}}\n"
	return
}

func TplPaginateTotal(data interface{}) (tpl string, err error) {
	table, err := convertData2Table(data)
	if err != nil {
		return
	}
	name := "PaginateTotal"
	tpl = fmt.Sprintf("{{define \"%s\"}}\nselect count(*) as `count` from `%s`  where 1=1 {{template \"%s\" \".\"}} ", name, table.TableName, tplPaginateWhereName)
	tpl = tpl + ";\n{{end}}\n"
	return
}

func TplPaginate(data interface{}) (tpl string, err error) {
	table, err := convertData2Table(data)
	if err != nil {
		return
	}
	name := "Paginate"
	tpl = fmt.Sprintf("{{define \"%s\"}}\nselect * from `%s`  where 1=1 {{template \"%s\" \".\"}} ", name, table.TableName, tplPaginateWhereName)
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
	name := "Insert"
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
	name := "Update"
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
	name := "Del"
	tpl = fmt.Sprintf("{{define \"%s\"}}\nupdate `%s` set `%s`={{currentTime .}} where `%s`=:%s;\n{{end}}\n", name, table.TableName, table.DeleteColumn, table.PrimaryKey, table.PrimaryKeyCamel())
	return
}