package gqttool

import (
	"fmt"
	"strings"

	"github.com/suifengpiao14/gqt/v2"
)

func Backquote(s string) (out string) {
	out = fmt.Sprintf("`%s`", s)
	return
}

// GenerateSQLTpl 此处的模板，是以meta 模版为标准输出的，即分割符号为[[define "xxx"]]...[[end]],一个gqt.TPLDefine 中可能包含多个{{define "xxx"}}...{{end}}
func GenerateSQLTpl(table *Table, repo *RepositoryMeta) (tplDefineTextList TPLDefineTextList, err error) { // list 保证后面输出顺序
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
	tplDefineTextList = TPLDefineTextList{}
	if tableNamespace != "" { // 确保单独文件定义的模板在前面，通用定义在后面
		tableTplDefineList, err := repo.GetByNamespace(tableNamespace, table)
		if err != nil {
			return tplDefineTextList, err
		}
		for _, gqtDefine := range tableTplDefineList {
			subTplDefineTextList := ManualParseDefine(gqtDefine.Output, "", gqt.LeftDelim, gqt.RightDelim)
			tplDefineTextList = append(tplDefineTextList, subTplDefineTextList...)
		}
	}
	if sqlNamespace != "" {
		sqlTplDefineList, err := repo.GetByNamespace(sqlNamespace, table)
		if err != nil {
			return tplDefineTextList, err
		}
		for _, gqtDefine := range sqlTplDefineList {
			subTplDefineTextList := ManualParseDefine(gqtDefine.Output, "", gqt.LeftDelim, gqt.RightDelim)
			tplDefineTextList = append(tplDefineTextList, subTplDefineTextList...)
		}
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
