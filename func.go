package gqttool

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/scanner"
	"go/token"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/suifengpiao14/gqt/v2/gqttpl"
	"golang.org/x/tools/go/ast/astutil"
	"golang.org/x/tools/imports"
)

// 封装 goa.design/goa/v3/codegen 方便后续可定制
func ToCamel(name string) string {
	return gqttpl.ToCamel(name)
}

func ToLowerCamel(name string) string {
	return gqttpl.ToLowerCamel(name)
}

func SnakeCase(name string) string {
	return gqttpl.SnakeCase(name)
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

// FinalizeGoSource removes unneeded imports from the given Go source file and
// runs go fmt on it.
func FinalizeGoSource(path string) error {
	// Make sure file parses and print content if it does not.
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		content, _ := ioutil.ReadFile(path)
		var buf bytes.Buffer
		scanner.PrintError(&buf, err)
		return fmt.Errorf("%s\n========\nContent:\n%s", buf.String(), content)
	}

	// Clean unused imports
	imps := astutil.Imports(fset, file)
	for _, group := range imps {
		for _, imp := range group {
			path := strings.Trim(imp.Path.Value, `"`)
			if !astutil.UsesImport(file, path) {
				if imp.Name != nil {
					astutil.DeleteNamedImport(fset, file, imp.Name.Name, path)
				} else {
					astutil.DeleteImport(fset, file, path)
				}
			}
		}
	}
	ast.SortImports(fset, file)
	w, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}
	if err := format.Node(w, fset, file); err != nil {
		return err
	}
	w.Close()

	// Format code using goimport standard
	bs, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	opt := imports.Options{
		Comments:   true,
		FormatOnly: true,
	}
	bs, err = imports.Process(path, bs, &opt)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, bs, os.ModePerm)
}
