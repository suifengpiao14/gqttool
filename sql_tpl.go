package gqttool

import (
	"fmt"

	"github.com/suifengpiao14/gqt/v2/gqttpl"
)

func Backquote(s string) (out string) {
	out = fmt.Sprintf("`%s`", s)
	return
}

func GenerateSQLTpl(table *Table, repo *RepositoryMeta) (defineResultList []*gqttpl.TPLDefine, err error) { // list 保证后面输出顺序
	metaNamespace, err := repo.GetNamespaceBySufix(MetaNameSpaceSuffix)
	if err != nil {
		return
	}
	defineResultList, err = repo.GetByNamespace(metaNamespace, table)
	if err != nil {
		return
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
