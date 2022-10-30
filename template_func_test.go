package gqttool

import (
	"fmt"
	"regexp"
	"testing"
)

func TestConvertData2Table(t *testing.T) {
	table := &Table{}
	it := interface{}(table)
	newTable, err := convertData2Table(&it)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v", newTable)
}

func TestRegexpInt(t *testing.T) {
	re := regexp.MustCompile(`\d+`)
	str := "20"
	ok := re.MatchString(str)
	fmt.Println(ok)
}
