package gqttool

import "strings"

var TranslateMapDefault = map[string]string{
	"insert":   "新增",
	"update":   "修改",
	"del":      "删除",
	"list":     "列表",
	"paginate": "分页列表",
	"total":    "总数",
	"get":      "获取",
	"by":       "通过",
	"id":       "ID",
	"all":      "所有",
}

func Translate(name string, extraMap map[string]string) string {
	snakeName := SnakeCase(name)
	wordArr := strings.Split(snakeName, "_")
	titleArr := make([]string, 0)
	existsWordMap := make(map[string]string)
	if extraMap !=nil{
		for _,extra:=range
	}
	for _, word := range wordArr {
		if _, ok := existsWordMap[word]; ok {
			continue
		}
		if extraMap != nil {
			title, ok := extraMap[word]
			if ok {
				existsWordMap[word] = title
				titleArr = append(titleArr, title)
				continue
			}
		}

		title, ok := TranslateMapDefault[word]
		if ok {
			existsWordMap[word] = title
			titleArr = append(titleArr, title)
			continue
		}
		title = word
		existsWordMap[word] = title
		titleArr = append(titleArr, title)
	}
	out := strings.Join(titleArr, " ")
	return out
}
