package gqttool

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/jinzhu/copier"
	"github.com/suifengpiao14/gqt/v2"
	"goa.design/goa/v3/codegen"
)

// 封装 goa.design/goa/v3/codegen 方便后续可定制
func ToCamel(name string) string {
	return codegen.CamelCase(name, true, true)
}

func ToLowerCamel(name string) string {
	return codegen.CamelCase(name, false, true)
}

func SnakeCase(name string) string {
	return codegen.SnakeCase(name)
}

func GeneratePackageName(dstDir string) (packageName string, err error) {
	if runtime.GOOS == "windows" { // drop driver name
		index := strings.Index(dstDir, ":")
		if index >= 0 {
			dstDir = dstDir[index+1:]
		}
		dstDir = strings.ReplaceAll(dstDir, "\\", "/")
	}
	absoluteDir := dstDir
	if dstDir[0:1] != "/" {
		absoluteDir, err = filepath.Abs(fmt.Sprintf("%s/%s", filepath.Dir(os.Args[0]), dstDir))
		if err != nil {
			return "", err
		}
	}
	if runtime.GOOS == "windows" {
		absoluteDir = strings.ReplaceAll(absoluteDir, "\\", "/")
	}

	basename := path.Base(absoluteDir)
	packageName = strings.ToLower(strcase.ToLowerCamel(basename))
	return

}

// Model2Entity copy model to entity ,some times input used to insert and update ,in this case input mybe model, copy model value to insertEntity and updateEntity
func Model2Entity(from interface{}, to gqt.TplEntity) {
	err := copier.Copy(to, from)
	if err != nil {
		panic(err)
	}
}
