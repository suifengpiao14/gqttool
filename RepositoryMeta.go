package gqttool

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"
	"github.com/suifengpiao14/gqt/v2"
	"github.com/suifengpiao14/gqt/v2/gqttpl"
)

type RepositoryMeta struct {
	templates map[string]*template.Template // namespace: template
}

type DatabaseConfig struct {
	TablePrefix     string `json:"tablePrefix"`
	ColumnPrefix    string `json:"columnPrefix"`
	DeletedAtColumn string `json:"deletedAtColumn"`
}

var MetaNameSpaceSuffix = gqttpl.MetaNamespaceSuffix
var MetaLeftDelim = "[["
var MetaRightDelim = "]]"

var DDLNamespaceSuffix = gqttpl.DDLNamespaceSuffix
var ConfigNamespaceSuffix = gqttpl.ConfigNamespaceSuffix

// ddl namespace sufix . define name prefix
var DatabaseConfigName = "database"

func NewRepositoryMeta() *RepositoryMeta {
	return &RepositoryMeta{
		templates: make(map[string]*template.Template),
	}
}

func (r *RepositoryMeta) AddByDir(root string, funcMap template.FuncMap) (err error) {
	r.templates, err = gqttpl.AddTemplateByDir(root, MetaNameSpaceSuffix, funcMap, MetaLeftDelim, MetaRightDelim)
	if err != nil {
		return
	}
	ddlTemplates, err := gqttpl.AddTemplateByDir(root, DDLNamespaceSuffix, funcMap, gqt.LeftDelim, gqt.RightDelim)
	if err != nil {
		return
	}

	for fullname, tpl := range ddlTemplates {
		r.templates[fullname] = tpl
	}
	configTemplates, err := gqttpl.AddTemplateByDir(root, ConfigNamespaceSuffix, funcMap, gqt.LeftDelim, gqt.RightDelim)
	if err != nil {
		return
	}
	for fullname, tpl := range configTemplates {
		r.templates[fullname] = tpl
	}
	return
}

func (r *RepositoryMeta) AddByNamespace(namespace string, content string, funcMap template.FuncMap) (err error) {
	t, err := gqttpl.AddTemplateByStr(namespace, content, funcMap, MetaLeftDelim, MetaRightDelim)
	if err != nil {
		err = errors.WithStack(err)
		return err
	}
	r.templates[namespace] = t
	return
}

// GetByNamespace get all template under namespace
func (r *RepositoryMeta) GetByNamespace(namespace string, data interface{}) (defineResultList []*gqttpl.TPLDefine, err error) {
	defineResultList, err = gqttpl.ExecuteNamespaceTemplate(r.templates, namespace, data)
	if err != nil {
		return nil, err
	}
	return
}

func (r *RepositoryMeta) GetTPLDefine(fullname string, data interface{}) (tplDefine *gqttpl.TPLDefine, err error) {
	tplDefine, err = gqttpl.ExecuteTemplate(r.templates, fullname, data)
	if err != nil {
		return nil, err
	}
	return
}

func (r *RepositoryMeta) GetNamespaceBySufix(suffix string) (namespace string, err error) {
	for nsp := range r.templates {
		if strings.HasSuffix(nsp, suffix) {
			namespace = nsp
			break
		}
	}
	if namespace == "" {
		err = errors.Errorf("not find  namespace with sufix %s", suffix)
		return
	}
	return
}

func (r *RepositoryMeta) GetDatabaseConfig() (cfg *DatabaseConfig, err error) {
	cfg = &DatabaseConfig{}
	namespace, err := r.GetNamespaceBySufix(ConfigNamespaceSuffix)
	if err != nil {
		return
	}

	fullname := fmt.Sprintf("%s.%s", namespace, DatabaseConfigName)
	tplDefine, err := r.GetTPLDefine(fullname, nil)
	if err != nil {
		return
	}
	_, err = toml.Decode(tplDefine.Output, cfg)
	if err != nil {
		return
	}
	return
}

func (r *RepositoryMeta) GetDDLTPLDefine() (defineResultList []*gqttpl.TPLDefine, err error) {
	ddlNamespace, err := r.GetNamespaceBySufix(DDLNamespaceSuffix)
	if err != nil {
		return
	}
	defineResultList, err = r.GetByNamespace(ddlNamespace, nil)
	if err != nil {
		return
	}
	for _, defineResult := range defineResultList {
		defineResult.Output = gqttpl.StandardizeSpaces(defineResult.Output)
	}
	return
}
