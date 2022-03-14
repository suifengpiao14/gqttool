package gqttool

import (
	"fmt"
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
