package example

import "github.com/suifengpiao14/gqt/v2/gqttpl"

type SQLAPISQLDelEntity struct {
	APIID         string
	OperatorIDInt int
	OperatorStr   string
	gqttpl.TplEmptyEntity
}

func (t *SQLAPISQLDelEntity) TplName() string {
	return "sql.api.sql.Del"
}
func (t *SQLAPISQLDelEntity) TplType() string {
	return "sql_update"
}

type SQLAPISQLGetAllByAPIIDListEntity struct {
	APIIDList []string
	gqttpl.TplEmptyEntity
}

func (t *SQLAPISQLGetAllByAPIIDListEntity) TplName() string {
	return "sql.api.sql.GetAllByAPIIDList"
}
func (t *SQLAPISQLGetAllByAPIIDListEntity) TplType() string {
	return "sql_select"
}

type SQLAPISQLGetByAPIIDEntity struct {
	APIID string
	gqttpl.TplEmptyEntity
}

func (t *SQLAPISQLGetByAPIIDEntity) TplName() string {
	return "sql.api.sql.GetByAPIID"
}
func (t *SQLAPISQLGetByAPIIDEntity) TplType() string {
	return "sql_select"
}

type SQLAPISQLInsertEntity struct {
	APIID       string
	Description string
	Name        string
	ServiceID   string
	Summary     string
	Tags        string
	Title       string
	URI         string
	gqttpl.TplEmptyEntity
}

func (t *SQLAPISQLInsertEntity) TplName() string {
	return "sql.api.sql.Insert"
}
func (t *SQLAPISQLInsertEntity) TplType() string {
	return "sql_insert"
}

type SQLAPISQLPaginateEntity struct {
	Limit  int
	Offset int
	SQLAPISQLPaginateWhereEntity
	gqttpl.TplEmptyEntity
}

func (t *SQLAPISQLPaginateEntity) TplName() string {
	return "sql.api.sql.Paginate"
}
func (t *SQLAPISQLPaginateEntity) TplType() string {
	return "sql_select"
}

type SQLAPISQLPaginateTotalEntity struct {
	SQLAPISQLPaginateWhereEntity
	gqttpl.TplEmptyEntity
}

func (t *SQLAPISQLPaginateTotalEntity) TplName() string {
	return "sql.api.sql.PaginateTotal"
}
func (t *SQLAPISQLPaginateTotalEntity) TplType() string {
	return "sql_select"
}

type SQLAPISQLPaginateWhereEntity struct {
	gqttpl.TplEmptyEntity
}

func (t *SQLAPISQLPaginateWhereEntity) TplName() string {
	return "sql.api.sql.PaginateWhere"
}
func (t *SQLAPISQLPaginateWhereEntity) TplType() string {
	return "text"
}

type SQLAPISQLUpdateEntity struct {
	APIID       string
	Description string
	Name        string
	ServiceID   string
	Summary     string
	Tags        string
	Title       string
	URI         string
	gqttpl.TplEmptyEntity
}

func (t *SQLAPISQLUpdateEntity) TplName() string {
	return "sql.api.sql.Update"
}
func (t *SQLAPISQLUpdateEntity) TplType() string {
	return "sql_update"
}

type SQLExampleSQLDelEntity struct {
	ExampleID string
	gqttpl.TplEmptyEntity
}

func (t *SQLExampleSQLDelEntity) TplName() string {
	return "sql.example.sql.Del"
}
func (t *SQLExampleSQLDelEntity) TplType() string {
	return "sql_update"
}

type SQLExampleSQLGetAllByExampleIDListEntity struct {
	ExampleIDList []string
	gqttpl.TplEmptyEntity
}

func (t *SQLExampleSQLGetAllByExampleIDListEntity) TplName() string {
	return "sql.example.sql.GetAllByExampleIDList"
}
func (t *SQLExampleSQLGetAllByExampleIDListEntity) TplType() string {
	return "sql_select"
}

type SQLExampleSQLGetByExampleIDEntity struct {
	ExampleID string
	gqttpl.TplEmptyEntity
}

func (t *SQLExampleSQLGetByExampleIDEntity) TplName() string {
	return "sql.example.sql.GetByExampleID"
}
func (t *SQLExampleSQLGetByExampleIDEntity) TplType() string {
	return "sql_select"
}

type SQLExampleSQLInsertEntity struct {
	APIID     string
	ExampleID string
	Request   string
	Response  string
	ServiceID string
	Summary   string
	Tag       string
	Title     string
	gqttpl.TplEmptyEntity
}

func (t *SQLExampleSQLInsertEntity) TplName() string {
	return "sql.example.sql.Insert"
}
func (t *SQLExampleSQLInsertEntity) TplType() string {
	return "sql_insert"
}

type SQLExampleSQLPaginateEntity struct {
	Limit  int
	Offset int
	SQLExampleSQLPaginateWhereEntity
	gqttpl.TplEmptyEntity
}

func (t *SQLExampleSQLPaginateEntity) TplName() string {
	return "sql.example.sql.Paginate"
}
func (t *SQLExampleSQLPaginateEntity) TplType() string {
	return "sql_select"
}

type SQLExampleSQLPaginateTotalEntity struct {
	SQLExampleSQLPaginateWhereEntity
	gqttpl.TplEmptyEntity
}

func (t *SQLExampleSQLPaginateTotalEntity) TplName() string {
	return "sql.example.sql.PaginateTotal"
}
func (t *SQLExampleSQLPaginateTotalEntity) TplType() string {
	return "sql_select"
}

type SQLExampleSQLPaginateWhereEntity struct {
	gqttpl.TplEmptyEntity
}

func (t *SQLExampleSQLPaginateWhereEntity) TplName() string {
	return "sql.example.sql.PaginateWhere"
}
func (t *SQLExampleSQLPaginateWhereEntity) TplType() string {
	return "text"
}

type SQLExampleSQLUpdateEntity struct {
	APIID     string
	ExampleID string
	Request   string
	Response  string
	ServiceID string
	Summary   string
	Tag       string
	Title     string
	gqttpl.TplEmptyEntity
}

func (t *SQLExampleSQLUpdateEntity) TplName() string {
	return "sql.example.sql.Update"
}
func (t *SQLExampleSQLUpdateEntity) TplType() string {
	return "sql_update"
}

type SQLMarkdownSQLDelEntity struct {
	MarkdownID string
	gqttpl.TplEmptyEntity
}

func (t *SQLMarkdownSQLDelEntity) TplName() string {
	return "sql.markdown.sql.Del"
}
func (t *SQLMarkdownSQLDelEntity) TplType() string {
	return "sql_update"
}

type SQLMarkdownSQLGetAllByMarkdownIDListEntity struct {
	MarkdownIDList []string
	gqttpl.TplEmptyEntity
}

func (t *SQLMarkdownSQLGetAllByMarkdownIDListEntity) TplName() string {
	return "sql.markdown.sql.GetAllByMarkdownIDList"
}
func (t *SQLMarkdownSQLGetAllByMarkdownIDListEntity) TplType() string {
	return "sql_select"
}

type SQLMarkdownSQLGetByMarkdownIDEntity struct {
	MarkdownID string
	gqttpl.TplEmptyEntity
}

func (t *SQLMarkdownSQLGetByMarkdownIDEntity) TplName() string {
	return "sql.markdown.sql.GetByMarkdownID"
}
func (t *SQLMarkdownSQLGetByMarkdownIDEntity) TplType() string {
	return "sql_select"
}

type SQLMarkdownSQLInsertEntity struct {
	APIID      string
	Content    string
	Markdown   string
	MarkdownID string
	Name       string
	OwnerID    int
	OwnerName  string
	ServiceID  string
	Title      string
	gqttpl.TplEmptyEntity
}

func (t *SQLMarkdownSQLInsertEntity) TplName() string {
	return "sql.markdown.sql.Insert"
}
func (t *SQLMarkdownSQLInsertEntity) TplType() string {
	return "sql_insert"
}

type SQLMarkdownSQLPaginateEntity struct {
	Limit  int
	Offset int
	SQLMarkdownSQLPaginateWhereEntity
	gqttpl.TplEmptyEntity
}

func (t *SQLMarkdownSQLPaginateEntity) TplName() string {
	return "sql.markdown.sql.Paginate"
}
func (t *SQLMarkdownSQLPaginateEntity) TplType() string {
	return "sql_select"
}

type SQLMarkdownSQLPaginateTotalEntity struct {
	SQLMarkdownSQLPaginateWhereEntity
	gqttpl.TplEmptyEntity
}

func (t *SQLMarkdownSQLPaginateTotalEntity) TplName() string {
	return "sql.markdown.sql.PaginateTotal"
}
func (t *SQLMarkdownSQLPaginateTotalEntity) TplType() string {
	return "sql_select"
}

type SQLMarkdownSQLPaginateWhereEntity struct {
	gqttpl.TplEmptyEntity
}

func (t *SQLMarkdownSQLPaginateWhereEntity) TplName() string {
	return "sql.markdown.sql.PaginateWhere"
}
func (t *SQLMarkdownSQLPaginateWhereEntity) TplType() string {
	return "text"
}

type SQLMarkdownSQLUpdateEntity struct {
	APIID      string
	Content    string
	Markdown   string
	MarkdownID string
	Name       string
	OwnerID    int
	OwnerName  string
	ServiceID  string
	Title      string
	gqttpl.TplEmptyEntity
}

func (t *SQLMarkdownSQLUpdateEntity) TplName() string {
	return "sql.markdown.sql.Update"
}
func (t *SQLMarkdownSQLUpdateEntity) TplType() string {
	return "sql_update"
}

type SQLParameterSQLDelEntity struct {
	ParameterID string
	gqttpl.TplEmptyEntity
}

func (t *SQLParameterSQLDelEntity) TplName() string {
	return "sql.parameter.sql.Del"
}
func (t *SQLParameterSQLDelEntity) TplType() string {
	return "sql_update"
}

type SQLParameterSQLGetAllByParameterIDListEntity struct {
	ParameterIDList []string
	gqttpl.TplEmptyEntity
}

func (t *SQLParameterSQLGetAllByParameterIDListEntity) TplName() string {
	return "sql.parameter.sql.GetAllByParameterIDList"
}
func (t *SQLParameterSQLGetAllByParameterIDListEntity) TplType() string {
	return "sql_select"
}

type SQLParameterSQLGetByParameterIDEntity struct {
	ParameterID string
	gqttpl.TplEmptyEntity
}

func (t *SQLParameterSQLGetByParameterIDEntity) TplName() string {
	return "sql.parameter.sql.GetByParameterID"
}
func (t *SQLParameterSQLGetByParameterIDEntity) TplType() string {
	return "sql_select"
}

type SQLParameterSQLInsertEntity struct {
	APIID            string
	AllowEmptyValue  string
	AllowReserved    string
	Deprecated       string
	Description      string
	Example          string
	Explode          string
	FullName         string
	HTTPStatus       string
	Method           string
	Name             string
	ParameterID      string
	Position         string
	Required         string
	Serialize        string
	ServiceID        string
	Tag              string
	Title            string
	Type             string
	ValidateSchemaID string
	gqttpl.TplEmptyEntity
}

func (t *SQLParameterSQLInsertEntity) TplName() string {
	return "sql.parameter.sql.Insert"
}
func (t *SQLParameterSQLInsertEntity) TplType() string {
	return "sql_insert"
}

type SQLParameterSQLPaginateEntity struct {
	Limit  int
	Offset int
	SQLParameterSQLPaginateWhereEntity
	gqttpl.TplEmptyEntity
}

func (t *SQLParameterSQLPaginateEntity) TplName() string {
	return "sql.parameter.sql.Paginate"
}
func (t *SQLParameterSQLPaginateEntity) TplType() string {
	return "sql_select"
}

type SQLParameterSQLPaginateTotalEntity struct {
	SQLParameterSQLPaginateWhereEntity
	gqttpl.TplEmptyEntity
}

func (t *SQLParameterSQLPaginateTotalEntity) TplName() string {
	return "sql.parameter.sql.PaginateTotal"
}
func (t *SQLParameterSQLPaginateTotalEntity) TplType() string {
	return "sql_select"
}

type SQLParameterSQLPaginateWhereEntity struct {
	gqttpl.TplEmptyEntity
}

func (t *SQLParameterSQLPaginateWhereEntity) TplName() string {
	return "sql.parameter.sql.PaginateWhere"
}
func (t *SQLParameterSQLPaginateWhereEntity) TplType() string {
	return "text"
}

type SQLParameterSQLUpdateEntity struct {
	APIID            string
	AllowEmptyValue  string
	AllowReserved    string
	Deprecated       string
	Description      string
	Example          string
	Explode          string
	FullName         string
	HTTPStatus       string
	Method           string
	Name             string
	ParameterID      string
	Position         string
	Required         string
	Serialize        string
	ServiceID        string
	Tag              string
	Title            string
	Type             string
	ValidateSchemaID string
	gqttpl.TplEmptyEntity
}

func (t *SQLParameterSQLUpdateEntity) TplName() string {
	return "sql.parameter.sql.Update"
}
func (t *SQLParameterSQLUpdateEntity) TplType() string {
	return "sql_update"
}

type SQLServerSQLDelEntity struct {
	ServerID string
	gqttpl.TplEmptyEntity
}

func (t *SQLServerSQLDelEntity) TplName() string {
	return "sql.server.sql.Del"
}
func (t *SQLServerSQLDelEntity) TplType() string {
	return "sql_update"
}

type SQLServerSQLGetAllByServerIDListEntity struct {
	ServerIDList []string
	gqttpl.TplEmptyEntity
}

func (t *SQLServerSQLGetAllByServerIDListEntity) TplName() string {
	return "sql.server.sql.GetAllByServerIDList"
}
func (t *SQLServerSQLGetAllByServerIDListEntity) TplType() string {
	return "sql_select"
}

type SQLServerSQLGetByServerIDEntity struct {
	ServerID string
	gqttpl.TplEmptyEntity
}

func (t *SQLServerSQLGetByServerIDEntity) TplName() string {
	return "sql.server.sql.GetByServerID"
}
func (t *SQLServerSQLGetByServerIDEntity) TplType() string {
	return "sql_select"
}

type SQLServerSQLInsertEntity struct {
	Description  string
	ExtensionIds string
	Proxy        string
	ServerID     string
	ServiceID    string
	URL          string
	gqttpl.TplEmptyEntity
}

func (t *SQLServerSQLInsertEntity) TplName() string {
	return "sql.server.sql.Insert"
}
func (t *SQLServerSQLInsertEntity) TplType() string {
	return "sql_insert"
}

type SQLServerSQLPaginateEntity struct {
	Limit  int
	Offset int
	SQLServerSQLPaginateWhereEntity
	gqttpl.TplEmptyEntity
}

func (t *SQLServerSQLPaginateEntity) TplName() string {
	return "sql.server.sql.Paginate"
}
func (t *SQLServerSQLPaginateEntity) TplType() string {
	return "sql_select"
}

type SQLServerSQLPaginateTotalEntity struct {
	SQLServerSQLPaginateWhereEntity
	gqttpl.TplEmptyEntity
}

func (t *SQLServerSQLPaginateTotalEntity) TplName() string {
	return "sql.server.sql.PaginateTotal"
}
func (t *SQLServerSQLPaginateTotalEntity) TplType() string {
	return "sql_select"
}

type SQLServerSQLPaginateWhereEntity struct {
	gqttpl.TplEmptyEntity
}

func (t *SQLServerSQLPaginateWhereEntity) TplName() string {
	return "sql.server.sql.PaginateWhere"
}
func (t *SQLServerSQLPaginateWhereEntity) TplType() string {
	return "text"
}

type SQLServerSQLUpdateEntity struct {
	Description  string
	ExtensionIds string
	Proxy        string
	ServerID     string
	ServiceID    string
	URL          string
	gqttpl.TplEmptyEntity
}

func (t *SQLServerSQLUpdateEntity) TplName() string {
	return "sql.server.sql.Update"
}
func (t *SQLServerSQLUpdateEntity) TplType() string {
	return "sql_update"
}

type SQLServiceSQLDelEntity struct {
	ServiceID string
	gqttpl.TplEmptyEntity
}

func (t *SQLServiceSQLDelEntity) TplName() string {
	return "sql.service.sql.Del"
}
func (t *SQLServiceSQLDelEntity) TplType() string {
	return "sql_update"
}

type SQLServiceSQLGetAllByServiceIDListEntity struct {
	ServiceIDList []string
	gqttpl.TplEmptyEntity
}

func (t *SQLServiceSQLGetAllByServiceIDListEntity) TplName() string {
	return "sql.service.sql.GetAllByServiceIDList"
}
func (t *SQLServiceSQLGetAllByServiceIDListEntity) TplType() string {
	return "sql_select"
}

type SQLServiceSQLGetByServiceIDEntity struct {
	ServiceID string
	gqttpl.TplEmptyEntity
}

func (t *SQLServiceSQLGetByServiceIDEntity) TplName() string {
	return "sql.service.sql.GetByServiceID"
}
func (t *SQLServiceSQLGetByServiceIDEntity) TplType() string {
	return "sql_select"
}

type SQLServiceSQLInsertEntity struct {
	ContactIds  string
	Description string
	License     string
	Proxy       string
	Security    string
	ServiceID   string
	Title       string
	Variables   string
	Version     string
	gqttpl.TplEmptyEntity
}

func (t *SQLServiceSQLInsertEntity) TplName() string {
	return "sql.service.sql.Insert"
}
func (t *SQLServiceSQLInsertEntity) TplType() string {
	return "sql_insert"
}

type SQLServiceSQLPaginateEntity struct {
	Limit  int
	Offset int
	SQLServiceSQLPaginateWhereEntity
	gqttpl.TplEmptyEntity
}

func (t *SQLServiceSQLPaginateEntity) TplName() string {
	return "sql.service.sql.Paginate"
}
func (t *SQLServiceSQLPaginateEntity) TplType() string {
	return "sql_select"
}

type SQLServiceSQLPaginateTotalEntity struct {
	SQLServiceSQLPaginateWhereEntity
	gqttpl.TplEmptyEntity
}

func (t *SQLServiceSQLPaginateTotalEntity) TplName() string {
	return "sql.service.sql.PaginateTotal"
}
func (t *SQLServiceSQLPaginateTotalEntity) TplType() string {
	return "sql_select"
}

type SQLServiceSQLPaginateWhereEntity struct {
	gqttpl.TplEmptyEntity
}

func (t *SQLServiceSQLPaginateWhereEntity) TplName() string {
	return "sql.service.sql.PaginateWhere"
}
func (t *SQLServiceSQLPaginateWhereEntity) TplType() string {
	return "text"
}

type SQLServiceSQLUpdateEntity struct {
	ContactIds  string
	Description string
	License     string
	Proxy       string
	Security    string
	ServiceID   string
	Title       string
	Variables   string
	Version     string
	gqttpl.TplEmptyEntity
}

func (t *SQLServiceSQLUpdateEntity) TplName() string {
	return "sql.service.sql.Update"
}
func (t *SQLServiceSQLUpdateEntity) TplType() string {
	return "sql_update"
}

type SQLValidateSchemaSQLDelEntity struct {
	ValidateSchemaID string
	gqttpl.TplEmptyEntity
}

func (t *SQLValidateSchemaSQLDelEntity) TplName() string {
	return "sql.validate_schema.sql.Del"
}
func (t *SQLValidateSchemaSQLDelEntity) TplType() string {
	return "sql_update"
}

type SQLValidateSchemaSQLGetAllByValidateSchemaIDListEntity struct {
	ValidateSchemaIDList []string
	gqttpl.TplEmptyEntity
}

func (t *SQLValidateSchemaSQLGetAllByValidateSchemaIDListEntity) TplName() string {
	return "sql.validate_schema.sql.GetAllByValidateSchemaIDList"
}
func (t *SQLValidateSchemaSQLGetAllByValidateSchemaIDListEntity) TplType() string {
	return "sql_select"
}

type SQLValidateSchemaSQLGetByValidateSchemaIDEntity struct {
	ValidateSchemaID string
	gqttpl.TplEmptyEntity
}

func (t *SQLValidateSchemaSQLGetByValidateSchemaIDEntity) TplName() string {
	return "sql.validate_schema.sql.GetByValidateSchemaID"
}
func (t *SQLValidateSchemaSQLGetByValidateSchemaIDEntity) TplType() string {
	return "sql_select"
}

type SQLValidateSchemaSQLInsertEntity struct {
	AdditionalProperties string
	AllOf                string
	AllowEmptyValue      string
	AllowReserved        string
	AnyOf                string
	Default              string
	Deprecated           string
	Description          string
	Discriminator        string
	Enum                 string
	EnumNames            string
	EnumTitles           string
	Example              string
	ExclusiveMaximum     string
	ExclusiveMinimum     string
	Extensions           string
	ExternalDocs         string
	ExternalPros         string
	Format               string
	MaxItems             int
	MaxLength            int
	MaxProperties        int
	Maxnum               int
	MinItems             int
	MinLength            int
	MinProperties        int
	Minimum              int
	MultipleOf           int
	Not                  string
	Nullable             string
	OneOf                string
	Pattern              string
	ReadOnly             string
	Remark               string
	Required             string
	ServiceID            string
	Summary              string
	Type                 string
	UniqueItems          string
	ValidateSchemaID     string
	WriteOnly            string
	XML                  string
	gqttpl.TplEmptyEntity
}

func (t *SQLValidateSchemaSQLInsertEntity) TplName() string {
	return "sql.validate_schema.sql.Insert"
}
func (t *SQLValidateSchemaSQLInsertEntity) TplType() string {
	return "sql_insert"
}

type SQLValidateSchemaSQLPaginateEntity struct {
	Limit  int
	Offset int
	SQLValidateSchemaSQLPaginateWhereEntity
	gqttpl.TplEmptyEntity
}

func (t *SQLValidateSchemaSQLPaginateEntity) TplName() string {
	return "sql.validate_schema.sql.Paginate"
}
func (t *SQLValidateSchemaSQLPaginateEntity) TplType() string {
	return "sql_select"
}

type SQLValidateSchemaSQLPaginateTotalEntity struct {
	SQLValidateSchemaSQLPaginateWhereEntity
	gqttpl.TplEmptyEntity
}

func (t *SQLValidateSchemaSQLPaginateTotalEntity) TplName() string {
	return "sql.validate_schema.sql.PaginateTotal"
}
func (t *SQLValidateSchemaSQLPaginateTotalEntity) TplType() string {
	return "sql_select"
}

type SQLValidateSchemaSQLPaginateWhereEntity struct {
	gqttpl.TplEmptyEntity
}

func (t *SQLValidateSchemaSQLPaginateWhereEntity) TplName() string {
	return "sql.validate_schema.sql.PaginateWhere"
}
func (t *SQLValidateSchemaSQLPaginateWhereEntity) TplType() string {
	return "text"
}

type SQLValidateSchemaSQLUpdateEntity struct {
	AdditionalProperties string
	AllOf                string
	AllowEmptyValue      string
	AllowReserved        string
	AnyOf                string
	Default              string
	Deprecated           string
	Description          string
	Discriminator        string
	Enum                 string
	EnumNames            string
	EnumTitles           string
	Example              string
	ExclusiveMaximum     string
	ExclusiveMinimum     string
	Extensions           string
	ExternalDocs         string
	ExternalPros         string
	Format               string
	MaxItems             int
	MaxLength            int
	MaxProperties        int
	Maxnum               int
	MinItems             int
	MinLength            int
	MinProperties        int
	Minimum              int
	MultipleOf           int
	Not                  string
	Nullable             string
	OneOf                string
	Pattern              string
	ReadOnly             string
	Remark               string
	Required             string
	ServiceID            string
	Summary              string
	Type                 string
	UniqueItems          string
	ValidateSchemaID     string
	WriteOnly            string
	XML                  string
	gqttpl.TplEmptyEntity
}

func (t *SQLValidateSchemaSQLUpdateEntity) TplName() string {
	return "sql.validate_schema.sql.Update"
}
func (t *SQLValidateSchemaSQLUpdateEntity) TplType() string {
	return "sql_update"
}
