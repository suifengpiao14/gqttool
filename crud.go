package gqttool

import (
	"fmt"
	"strings"
)

func Backquote(s string) (out string) {
	out = fmt.Sprintf("`%s`", s)
	return
}

var MetaTplDel = "metaTplDel"

func Crud(table *Table, repo *RepositoryMeta) (sqlTplDefines []*SQLTPLDefine, err error) { // list 保证后面输出顺序
	defineResult, err := repo.GetDefine(MetaTplDel, nil)
	if err != nil {
		return
	}
	sqlTplDefines = make([]*SQLTPLDefine, 0)
	sqlTplDefines = append(sqlTplDefines, TplGetByPrimaryKey(table))
	sqlTplDefines = append(sqlTplDefines, TplTotal(table))
	sqlTplDefines = append(sqlTplDefines, TplPaginate(table))
	sqlTplDefines = append(sqlTplDefines, TplInsert(table))
	sqlTplDefines = append(sqlTplDefines, TplUpdate(table))
	if defineResult == nil {
		sqlTplDefines = append(sqlTplDefines, TplDel(table))
	}
	return
}

func isIgnoreColumn(column *Column, table *Table) (yes bool) {
	if column.AutoIncrement { // 自增列,忽略
		return true
	}
	columnName := column.Name
	if column.IsDefaultValueCurrentTimestamp() { // 自动填充时间的列,忽略
		return true
	}

	ignoreColumnMap := make(map[string]string)
	ignoreColumnMap[table.DeleteColumn] = table.DeleteColumn
	_, yes = ignoreColumnMap[columnName]
	return
}

func TplGetByPrimaryKey(table *Table) (sqlTplDefine *SQLTPLDefine) {
	sqlTplDefine = &SQLTPLDefine{}
	primaryKeyCamel := table.PrimaryKeyCamel()
	sqlTplDefine.Name = fmt.Sprintf("GetBy%s", primaryKeyCamel)
	tpl := fmt.Sprintf("{{define \"%s\"}}\nselect * from `%s`  where `%s`=:%s", sqlTplDefine.Name, table.TableName, table.PrimaryKey, primaryKeyCamel)
	if table.DeleteColumn != "" {
		tpl = fmt.Sprintf("%s  and `%s` is null", tpl, table.DeleteColumn)
	}
	tpl = tpl + ";\n{{end}}\n"
	sqlTplDefine.TPL = tpl

	return
}

func TplTotal(table *Table) (sqlTplDefine *SQLTPLDefine) {
	sqlTplDefine = &SQLTPLDefine{}
	sqlTplDefine.Name = "Total"
	tpl := fmt.Sprintf("{{define \"%s\"}}\nselect count(*) as `count` from `%s`  where 1=1 ", sqlTplDefine.Name, table.TableName)
	if table.DeleteColumn != "" {
		tpl = fmt.Sprintf("%s  and `%s` is null", tpl, table.DeleteColumn)
	}
	tpl = tpl + ";\n{{end}}\n"
	sqlTplDefine.TPL = tpl
	return
}

func TplPaginate(table *Table) (sqlTplDefine *SQLTPLDefine) {
	sqlTplDefine = &SQLTPLDefine{}
	sqlTplDefine.Name = "Paginate"
	tpl := fmt.Sprintf("{{define \"%s\"}}\nselect * from `%s`  where 1=1 ", sqlTplDefine.Name, table.TableName)
	deletedAtColumn := table.DeletedAtColumn()
	if deletedAtColumn != nil {
		tpl = fmt.Sprintf("%s  and `%s` is null ", tpl, deletedAtColumn.Name)
	}
	updatedAtColumn := table.UpdatedAtColumn()
	if updatedAtColumn != nil {
		tpl = fmt.Sprintf(" %s order by `%s` desc ", tpl, updatedAtColumn.Name)
	}
	tpl = fmt.Sprintf(" %s limit :Offset,:Limit ", tpl)
	tpl = tpl + ";\n{{end}}\n"
	sqlTplDefine.TPL = tpl
	return
}

func TplInsert(table *Table) (sqlTplDefine *SQLTPLDefine) {
	sqlTplDefine = &SQLTPLDefine{}
	sqlTplDefine.Name = "Insert"
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
	tpl := fmt.Sprintf("{{define \"%s\"}}\ninsert into `%s` (%s)values\n (%s);\n{{end}}\n", sqlTplDefine.Name, table.TableName, columnStr, valueStr)
	sqlTplDefine.TPL = tpl
	return
}

func TplUpdate(table *Table) (sqlTplDefine *SQLTPLDefine) {
	sqlTplDefine = &SQLTPLDefine{}
	sqlTplDefine.Name = "Update"
	updataList := make([]string, 0)
	for _, column := range table.Columns {
		if isIgnoreColumn(column, table) {
			continue
		}
		str := fmt.Sprintf("{{if .%s}} {{$preComma.PreComma}} `%s`=:%s {{end}} ", column.CamelName, column.Name, column.CamelName)
		updataList = append(updataList, str)
	}
	updateTpl := strings.Join(updataList, "\n")
	tpl := fmt.Sprintf("{{define \"%s\"}}\n{{$preComma:=newPreComma}}\n update `%s` set %s where `%s`=:%s;\n{{end}}\n", sqlTplDefine.Name, table.TableName, updateTpl, table.PrimaryKey, table.PrimaryKeyCamel())
	sqlTplDefine.TPL = tpl
	return
}

func TplDel(table *Table) (sqlTplDefine *SQLTPLDefine) {
	sqlTplDefine = &SQLTPLDefine{}
	sqlTplDefine.Name = "Del"
	tpl := fmt.Sprintf("{{define \"%s\"}}\nupdate `%s` set `%s`={{currentTime .}} where `%s`=:%s;\n{{end}}\n", sqlTplDefine.Name, table.TableName, table.DeleteColumn, table.PrimaryKey, table.PrimaryKeyCamel())
	sqlTplDefine.TPL = tpl
	return
}
