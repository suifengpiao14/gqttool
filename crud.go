package gqttool

import (
	"fmt"
	"strings"
)

func Backquote(s string) (out string) {
	out = fmt.Sprintf("`%s`", s)
	return
}
func Crud(table *Table) (tplList []string, err error) { // list 保证后面输出顺序
	tplList = make([]string, 0)
	tplList = append(tplList, TplGetByPrimaryKey(table))
	tplList = append(tplList, TplTotal(table))
	tplList = append(tplList, TplPaginate(table))
	tplList = append(tplList, TplInsert(table))
	tplList = append(tplList, TplUpdate(table))
	tplList = append(tplList, TplDel(table))

	return
}

func isIgnoreColumn(columnName string, table *Table) (yes bool) {
	ignoreColumnMap := make(map[string]string)
	ignoreColumnMap[table.CreatedAtColumn] = table.CreatedAtColumn
	ignoreColumnMap[table.UpdateAtColumn] = table.UpdateAtColumn
	ignoreColumnMap[table.DeleteAtColumn] = table.DeleteAtColumn
	_, yes = ignoreColumnMap[columnName]
	return
}

func TplGetByPrimaryKey(table *Table) (tpl string) {
	blockName := fmt.Sprintf("GetBy%s", table.PrimaryKeyCamel)
	tpl = fmt.Sprintf("{{define \"%s\"}}\nselect * from `%s`  where `%s`=:%s", blockName, table.TableName, table.PrimaryKey, table.PrimaryKeyCamel)
	if table.DeleteAtColumn != "" {
		tpl = fmt.Sprintf("%s  and `%s` is null", tpl, table.DeleteAtColumn)
	}
	tpl = tpl + ";\n{{end}}\n"
	return
}

func TplTotal(table *Table) (tpl string) {
	blockName := "Total"
	tpl = fmt.Sprintf("{{define \"%s\"}}\nselect count(*) as `count` from `%s`  where 1=1 ", blockName, table.TableName)
	if table.DeleteAtColumn != "" {
		tpl = fmt.Sprintf("%s  and `%s` is null", tpl, table.DeleteAtColumn)
	}
	tpl = tpl + ";\n{{end}}\n"
	return
}

func TplPaginate(table *Table) (tpl string) {
	blockName := "Paginate"
	tpl = fmt.Sprintf("{{define \"%s\"}}\nselect * from `%s`  where 1=1 ", blockName, table.TableName)
	if table.DeleteAtColumn != "" {
		tpl = fmt.Sprintf("%s  and `%s` is null ", tpl, table.DeleteAtColumn)
	}
	if table.UpdateAtColumn != "" {
		tpl = fmt.Sprintf(" %s order by `%s` desc ", tpl, table.UpdateAtColumn)
	}
	tpl = fmt.Sprintf(" %s limit :Offset,:Limit ", tpl)
	tpl = tpl + ";\n{{end}}\n"
	return
}

func TplInsert(table *Table) (tpl string) {
	blockName := "Insert"
	columns := make([]string, 0)
	values := make([]string, 0)
	for _, column := range table.Columns {
		if isIgnoreColumn(column.Name, table) {
			continue
		}
		columns = append(columns, Backquote(column.Name))
		values = append(values, ":"+column.CamelName)
	}

	columnStr := strings.Join(columns, ",")
	valueStr := strings.Join(values, ",")
	tpl = fmt.Sprintf("{{define \"%s\"}}\ninsert into `%s` (%s)values\n (%s);\n{{end}}\n", blockName, table.TableName, columnStr, valueStr)
	return
}

func TplUpdate(table *Table) (tpl string) {
	blockName := "Update"
	updataList := make([]string, 0)
	for _, column := range table.Columns {
		if isIgnoreColumn(column.Name, table) {
			continue
		}
		str := fmt.Sprintf("{{if .%s}} {{$preComma.PreComma}} `%s`=:%s {{end}} ", column.CamelName, column.Name, column.CamelName)
		updataList = append(updataList, str)
	}
	updateTpl := strings.Join(updataList, "\n")
	tpl = fmt.Sprintf("{{define \"%s\"}}\n{{$preComma:=newPreComma}}\n update `%s` set %s where `%s`=:%s;\n{{end}}\n", blockName, table.TableName, updateTpl, table.PrimaryKey, table.PrimaryKeyCamel)
	return
}

func TplDel(table *Table) (tpl string) {
	blockName := "Del"
	tpl = fmt.Sprintf("{{define \"%s\"}}\nupdate `%s` set `%s`={{currentTime}} where `%s`=:%s;\n{{end}}\n", blockName, table.TableName, table.DeleteAtColumn, table.PrimaryKey, table.PrimaryKeyCamel)
	return
}
