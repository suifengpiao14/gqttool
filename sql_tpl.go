package gqttool

import (
	"fmt"
	"strings"

	"github.com/suifengpiao14/gqt/v2/gqttpl"
)

func Backquote(s string) (out string) {
	out = fmt.Sprintf("`%s`", s)
	return
}

func GenerateSQLTpl(table *Table, repo *RepositoryMeta) (tplDefineList []*gqttpl.TPLDefine, err error) { // list 保证后面输出顺序
	metaNamespaceList, err := repo.GetNamespaceBySufix(MetaNameSpaceSuffix, true)
	if err != nil {
		return
	}
	sqlNamespace := ""
	tableNamespace := ""
	for _, namespace := range metaNamespaceList {
		sqlSuffix := fmt.Sprintf("%s.%s", "sql", MetaNameSpaceSuffix)
		if strings.HasSuffix(namespace, sqlSuffix) {
			sqlNamespace = namespace
		}
		tableSuffix := fmt.Sprintf("%s.%s", table.SnakeCase(), MetaNameSpaceSuffix)
		if strings.HasSuffix(namespace, tableSuffix) {
			tableNamespace = namespace
		}
	}
	tplDefineList = make([]*gqttpl.TPLDefine, 0)
	if sqlNamespace != "" {
		sqlTplDefineList, err := repo.GetByNamespace(sqlNamespace, table)
		if err != nil {
			return tplDefineList, err
		}
		tplDefineList = append(tplDefineList, sqlTplDefineList...)
	}

	if tableNamespace != "" {
		tableTplDefineList, err := repo.GetByNamespace(tableNamespace, table)
		if err != nil {
			return tplDefineList, err
		}
		tplDefineList = append(tplDefineList, tableTplDefineList...)
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
