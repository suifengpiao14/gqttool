package gqttool

import (
	"fmt"
	"sort"
	"strings"
	"text/template"

	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"
	"github.com/suifengpiao14/gqt/v2"
)

type RepositoryMeta struct {
	templates map[string]*template.Template // namespace: template
}

type DatabaseConfig struct {
	DatabaseName    string `mapstructure:"databaseName"`
	LogLevel        string `mapstructure:"logLevel"`
	TablePrefix     string `mapstructure:"tablePrefix"`
	ColumnPrefix    string `mapstructure:"columnPrefix"`
	DeletedAtColumn string `mapstructure:"deletedAtColumn"`
}
type Config struct {
	Batch string `mapstructure:"batch"`
}

var MetaNameSpaceSuffix = gqt.MetaNamespaceSuffix
var MetaLeftDelim = "[["
var MetaRightDelim = "]]"

var DDLNamespaceSuffix = gqt.DDLNamespaceSuffix
var ConfigNamespaceSuffix = gqt.ConfigNamespaceSuffix

// ddl namespace sufix . define name prefix
var DatabaseConfigName = "database"
var ConfigName = "config"

func NewRepositoryMeta() *RepositoryMeta {
	return &RepositoryMeta{
		templates: make(map[string]*template.Template),
	}
}

func (r *RepositoryMeta) AddByDir(root string, funcMap template.FuncMap) (err error) {
	r.templates, err = gqt.AddTemplateByDir(root, MetaNameSpaceSuffix, funcMap, MetaLeftDelim, MetaRightDelim)
	if err != nil {
		return
	}
	ddlTemplates, err := gqt.AddTemplateByDir(root, DDLNamespaceSuffix, funcMap, gqt.LeftDelim, gqt.RightDelim)
	if err != nil {
		return
	}

	for fullname, tpl := range ddlTemplates {
		r.templates[fullname] = tpl
	}
	configTemplates, err := gqt.AddTemplateByDir(root, ConfigNamespaceSuffix, funcMap, gqt.LeftDelim, gqt.RightDelim)
	if err != nil {
		return
	}
	for fullname, tpl := range configTemplates {
		r.templates[fullname] = tpl
	}
	return
}

func (r *RepositoryMeta) AddByNamespace(namespace string, content string, funcMap template.FuncMap) (err error) {
	leftDelim := gqt.LeftDelim
	rightDelim := gqt.RightDelim
	if strings.HasSuffix(namespace, MetaNameSpaceSuffix) {
		leftDelim = MetaLeftDelim
		rightDelim = MetaRightDelim
	}
	t, err := gqt.AddTemplateByStr(namespace, content, funcMap, leftDelim, rightDelim)
	if err != nil {
		err = errors.WithStack(err)
		return err
	}
	r.templates[namespace] = t
	return
}

// GetByNamespace get all template under namespace
func (r *RepositoryMeta) GetByNamespace(namespace string, data gqt.TplEntityInterface) (tplDefineList []*gqt.TPLDefine, err error) {
	tplDefineList, err = gqt.ExecuteNamespaceTemplate(r.templates, namespace, data)
	if err != nil {
		return nil, err
	}
	return
}

func (r *RepositoryMeta) GetTPLDefine(fullname string, data gqt.TplEntityInterface) (tplDefine *gqt.TPLDefine, err error) {
	tplDefine, err = gqt.ExecuteTemplate(r.templates, fullname, data)
	if err != nil {
		return nil, err
	}
	return
}

func (r *RepositoryMeta) GetNamespaceBySufix(suffix string, notFoundErr bool) (namespaceList []string, err error) {
	namespaceList = make([]string, 0)
	for namespace := range r.templates {
		if strings.HasSuffix(namespace, suffix) {
			namespaceList = append(namespaceList, namespace)
		}
	}
	l := len(namespaceList)
	if l == 0 {
		if notFoundErr { // 找不到报错
			err = errors.Errorf("not find  namespace with sufix %s", suffix)
			return
		}
		return
	}
	sort.Strings(namespaceList)
	return
}

func (r *RepositoryMeta) GetDatabaseConfig() (cfg *DatabaseConfig, err error) {
	cfg = &DatabaseConfig{}
	namespaceList, err := r.GetNamespaceBySufix(ConfigNamespaceSuffix, false)
	if err != nil {
		return
	}
	if len(namespaceList) == 0 {
		return
	}
	namespace := namespaceList[0]
	fullname := fmt.Sprintf("%s.%s", namespace, DatabaseConfigName)
	tplDefine, err := r.GetTPLDefine(fullname, &gqt.TplEmptyEntity{})
	if err != nil {
		return
	}
	_, err = toml.Decode(tplDefine.Output, cfg)
	if err != nil {
		return
	}
	if cfg.DatabaseName == "" {
		err = errors.Errorf("database name required (will as module name)")
		return
	}
	if cfg.LogLevel == "" {
		cfg.LogLevel = "debug"
	}
	return
}
func (r *RepositoryMeta) GetConfig() (cfg *Config, err error) {
	cfg = &Config{}
	namespaceList, err := r.GetNamespaceBySufix(ConfigNamespaceSuffix, false)
	if err != nil {
		return
	}
	if len(namespaceList) == 0 {
		return
	}
	namespace := namespaceList[0]
	fullname := fmt.Sprintf("%s.%s", namespace, ConfigName)
	tplDefine, err := r.GetTPLDefine(fullname, &gqt.TplEmptyEntity{})
	if err != nil {
		return
	}
	_, err = toml.Decode(tplDefine.Output, cfg)
	if err != nil {
		return
	}
	if cfg.Batch == "" {
		err = errors.Errorf("config.batch required")
		return
	}
	return
}

func (r *RepositoryMeta) GetDDLTPLDefine() (defineResultList []*gqt.TPLDefine, err error) {
	ddlNamespaceList, err := r.GetNamespaceBySufix(DDLNamespaceSuffix, true)
	if err != nil {
		return
	}
	ddlNamespace := ddlNamespaceList[0]
	defineResultList, err = r.GetByNamespace(ddlNamespace, nil)
	if err != nil {
		return
	}
	for _, defineResult := range defineResultList {
		defineResult.Output = gqt.StandardizeSpaces(defineResult.Output)
	}
	return
}
