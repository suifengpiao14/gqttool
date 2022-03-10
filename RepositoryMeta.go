package gqttool

import (
	"encoding/json"
	"fmt"
	"strings"
	"text/template"

	"github.com/pkg/errors"
	"github.com/suifengpiao14/gqt/v2/pkg"
	gqtpkg "github.com/suifengpiao14/gqt/v2/pkg"
)

type RepositoryMeta struct {
	templates map[string]*template.Template // namespace: template
}

type Config struct {
	TablePrefix     string `toml:"tablePrefix"`
	ColumnPrefix    string `toml:"columnPrefix"`
	DeletedAtColumn string `toml:"deletedAtColumn"`
}

var MetaSuffix = ".meta.tpl"
var DDLSuffix = ".ddl.tpl"
var LeftDelim = "[["
var RightDelim = "]]"

// ddl namespace sufix . define name prefix
var DDLNamespaceSuffix = "ddl"
var ConfigName = "config"

func NewRepositoryMeta() *RepositoryMeta {
	return &RepositoryMeta{
		templates: make(map[string]*template.Template),
	}
}

func (r *RepositoryMeta) AddByDir(root string, funcMap template.FuncMap) (err error) {
	r.templates, err = pkg.AddTemplateByDir(root, MetaSuffix, funcMap, LeftDelim, RightDelim)
	if err != nil {
		return
	}
	ddlTemplates, err := pkg.AddTemplateByDir(root, DDLSuffix, funcMap, LeftDelim, RightDelim)
	if err != nil {
		return
	}
	for fullname, tpl := range ddlTemplates {
		r.templates[fullname] = tpl
	}
	return
}

func (r *RepositoryMeta) AddByNamespace(namespace string, content string, funcMap template.FuncMap) (err error) {
	t, err := pkg.AddTemplateByStr(namespace, content, funcMap, LeftDelim, RightDelim)
	if err != nil {
		err = errors.WithStack(err)
		return err
	}
	r.templates[namespace] = t
	return
}

// GetByNamespace get all template under namespace
func (r *RepositoryMeta) GetByNamespace(namespace string, data interface{}) (defineResultList []*gqtpkg.DefineResult, err error) {
	defineResultList, err = gqtpkg.ExecuteNamespaceTemplate(r.templates, namespace, data)
	if err != nil {
		return nil, err
	}
	return
}

func (r *RepositoryMeta) GetDefine(fullname string, data interface{}) (defineResult *gqtpkg.DefineResult, err error) {
	defineResult, err = pkg.ExecuteTemplate(r.templates, fullname, data)
	if err != nil {
		return nil, err
	}
	return
}

func (r *RepositoryMeta) GetDDLNamespace() (ddlNamespace string, err error) {
	for namespace := range r.templates {
		if strings.HasSuffix(namespace, DDLNamespaceSuffix) {
			ddlNamespace = namespace
			break
		}
	}
	if ddlNamespace == "" {
		err = errors.Errorf("not find ddl namespace with sufix %s,you can set gqttool.DDLNamespaceSuffix to change sufix", DDLNamespaceSuffix)
		return
	}
	return
}
func (r *RepositoryMeta) GetConfig() (cfg *Config, err error) {
	cfg = &Config{}
	var configDefineResult *gqtpkg.DefineResult
	for namespace, tp := range r.templates {
		if strings.HasSuffix(namespace, DDLNamespaceSuffix) {
			continue
		}
		if configDefineResult != nil {
			break
		}
		subTpList := tp.Templates()
		for _, subTp := range subTpList {
			name := subTp.Name()
			if name == ConfigName {
				fullname := fmt.Sprintf("%s.%s", namespace, name)
				configDefineResult, err = gqtpkg.ExecuteTemplate(r.templates, fullname, nil)
				break
			}
		}
	}

	if configDefineResult != nil {
		var b = []byte(configDefineResult.Output)
		err = json.Unmarshal(b, cfg)
		if err != nil {
			return
		}
	}

	return
}

func (r *RepositoryMeta) GetDDLDefineResult() (defineResultList []*gqtpkg.DefineResult, err error) {
	ddlNamespace, err := r.GetDDLNamespace()
	if err != nil {
		return
	}
	defineResultList, err = r.GetByNamespace(ddlNamespace, nil)
	if err != nil {
		return
	}
	for _, defineResult := range defineResultList {
		defineResult.Output = pkg.StandardizeSpaces(defineResult.Output)
	}
	return
}
