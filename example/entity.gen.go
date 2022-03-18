package example

import "github.com/suifengpiao14/gqt/v2/gqttpl"

type SQLAPIGenSQLDelEntity struct {
	APIID string

	OperatorIDInt int

	OperatorStr string

	gqttpl.DataVolumeMap
}

func (t *SQLAPIGenSQLDelEntity) TplName() string {
	return "sql.api.gen.sql.Del"
}

type SQLAPIGenSQLGetAllByAPIIDListEntity struct {
	APIIDList []string

	gqttpl.DataVolumeMap
}

func (t *SQLAPIGenSQLGetAllByAPIIDListEntity) TplName() string {
	return "sql.api.gen.sql.GetAllByAPIIDList"
}

type SQLAPIGenSQLGetByAPIIDEntity struct {
	APIID string

	gqttpl.DataVolumeMap
}

func (t *SQLAPIGenSQLGetByAPIIDEntity) TplName() string {
	return "sql.api.gen.sql.GetByAPIID"
}

type SQLAPIGenSQLInsertEntity struct {
	APIID string

	Description string

	Name string

	ServiceID string

	Summary string

	Tags string

	Title string

	URI string

	gqttpl.DataVolumeMap
}

func (t *SQLAPIGenSQLInsertEntity) TplName() string {
	return "sql.api.gen.sql.Insert"
}

type SQLAPIGenSQLPaginateEntity struct {
	Limit int

	Offset int

	SQLAPIGenSQLPaginateWhereEntity

	gqttpl.DataVolumeMap
}

func (t *SQLAPIGenSQLPaginateEntity) TplName() string {
	return "sql.api.gen.sql.Paginate"
}

type SQLAPIGenSQLPaginateTotalEntity struct {
	SQLAPIGenSQLPaginateWhereEntity

	gqttpl.DataVolumeMap
}

func (t *SQLAPIGenSQLPaginateTotalEntity) TplName() string {
	return "sql.api.gen.sql.PaginateTotal"
}

type SQLAPIGenSQLPaginateWhereEntity struct {
	APIIDList interface{}

	gqttpl.DataVolumeMap
}

func (t *SQLAPIGenSQLPaginateWhereEntity) TplName() string {
	return "sql.api.gen.sql.PaginateWhere"
}

type SQLAPIGenSQLUpdateEntity struct {
	APIID string

	Description string

	Name string

	ServiceID string

	Summary string

	Tags string

	Title string

	URI string

	gqttpl.DataVolumeMap
}

func (t *SQLAPIGenSQLUpdateEntity) TplName() string {
	return "sql.api.gen.sql.Update"
}

type SQLExampleGenSQLGetAllByExampleIDListEntity struct {
	ExampleIDList []string

	gqttpl.DataVolumeMap
}

func (t *SQLExampleGenSQLGetAllByExampleIDListEntity) TplName() string {
	return "sql.example.gen.sql.GetAllByExampleIDList"
}

type SQLExampleGenSQLGetByExampleIDEntity struct {
	ExampleID string

	gqttpl.DataVolumeMap
}

func (t *SQLExampleGenSQLGetByExampleIDEntity) TplName() string {
	return "sql.example.gen.sql.GetByExampleID"
}

type SQLExampleGenSQLInsertEntity struct {
	APIID string

	ExampleID string

	Request string

	Response string

	ServiceID string

	Summary string

	Tag string

	Title string

	gqttpl.DataVolumeMap
}

func (t *SQLExampleGenSQLInsertEntity) TplName() string {
	return "sql.example.gen.sql.Insert"
}

type SQLExampleGenSQLPaginateEntity struct {
	Limit int

	Offset int

	SQLExampleGenSQLPaginateWhereEntity

	gqttpl.DataVolumeMap
}

func (t *SQLExampleGenSQLPaginateEntity) TplName() string {
	return "sql.example.gen.sql.Paginate"
}

type SQLExampleGenSQLPaginateTotalEntity struct {
	SQLExampleGenSQLPaginateWhereEntity

	gqttpl.DataVolumeMap
}

func (t *SQLExampleGenSQLPaginateTotalEntity) TplName() string {
	return "sql.example.gen.sql.PaginateTotal"
}

type SQLExampleGenSQLPaginateWhereEntity struct {
	gqttpl.DataVolumeMap
}

func (t *SQLExampleGenSQLPaginateWhereEntity) TplName() string {
	return "sql.example.gen.sql.PaginateWhere"
}

type SQLExampleGenSQLUpdateEntity struct {
	APIID string

	ExampleID string

	Request string

	Response string

	ServiceID string

	Summary string

	Tag string

	Title string

	gqttpl.DataVolumeMap
}

func (t *SQLExampleGenSQLUpdateEntity) TplName() string {
	return "sql.example.gen.sql.Update"
}

type SQLMarkdownGenSQLGetAllByMarkdownIDListEntity struct {
	MarkdownIDList []string

	gqttpl.DataVolumeMap
}

func (t *SQLMarkdownGenSQLGetAllByMarkdownIDListEntity) TplName() string {
	return "sql.markdown.gen.sql.GetAllByMarkdownIDList"
}

type SQLMarkdownGenSQLGetByMarkdownIDEntity struct {
	MarkdownID string

	gqttpl.DataVolumeMap
}

func (t *SQLMarkdownGenSQLGetByMarkdownIDEntity) TplName() string {
	return "sql.markdown.gen.sql.GetByMarkdownID"
}

type SQLMarkdownGenSQLInsertEntity struct {
	APIID string

	Content string

	Markdown string

	MarkdownID string

	Name string

	OwnerID int

	OwnerName string

	ServiceID string

	Title string

	gqttpl.DataVolumeMap
}

func (t *SQLMarkdownGenSQLInsertEntity) TplName() string {
	return "sql.markdown.gen.sql.Insert"
}

type SQLMarkdownGenSQLPaginateEntity struct {
	Limit int

	Offset int

	SQLMarkdownGenSQLPaginateWhereEntity

	gqttpl.DataVolumeMap
}

func (t *SQLMarkdownGenSQLPaginateEntity) TplName() string {
	return "sql.markdown.gen.sql.Paginate"
}

type SQLMarkdownGenSQLPaginateTotalEntity struct {
	SQLMarkdownGenSQLPaginateWhereEntity

	gqttpl.DataVolumeMap
}

func (t *SQLMarkdownGenSQLPaginateTotalEntity) TplName() string {
	return "sql.markdown.gen.sql.PaginateTotal"
}

type SQLMarkdownGenSQLPaginateWhereEntity struct {
	gqttpl.DataVolumeMap
}

func (t *SQLMarkdownGenSQLPaginateWhereEntity) TplName() string {
	return "sql.markdown.gen.sql.PaginateWhere"
}

type SQLMarkdownGenSQLUpdateEntity struct {
	APIID string

	Content string

	Markdown string

	MarkdownID string

	Name string

	OwnerID int

	OwnerName string

	ServiceID string

	Title string

	gqttpl.DataVolumeMap
}

func (t *SQLMarkdownGenSQLUpdateEntity) TplName() string {
	return "sql.markdown.gen.sql.Update"
}

type SQLParameterGenSQLGetAllByParameterIDListEntity struct {
	ParameterIDList []string

	gqttpl.DataVolumeMap
}

func (t *SQLParameterGenSQLGetAllByParameterIDListEntity) TplName() string {
	return "sql.parameter.gen.sql.GetAllByParameterIDList"
}

type SQLParameterGenSQLGetByParameterIDEntity struct {
	ParameterID string

	gqttpl.DataVolumeMap
}

func (t *SQLParameterGenSQLGetByParameterIDEntity) TplName() string {
	return "sql.parameter.gen.sql.GetByParameterID"
}

type SQLParameterGenSQLInsertEntity struct {
	APIID string

	AllowEmptyValue string

	AllowReserved string

	Deprecated string

	Description string

	Example string

	Explode string

	FullName string

	HTTPStatus string

	Method string

	Name string

	ParameterID string

	Position string

	Required string

	Serialize string

	ServiceID string

	Tag string

	Title string

	Type string

	ValidateSchemaID string

	gqttpl.DataVolumeMap
}

func (t *SQLParameterGenSQLInsertEntity) TplName() string {
	return "sql.parameter.gen.sql.Insert"
}

type SQLParameterGenSQLPaginateEntity struct {
	Limit int

	Offset int

	SQLParameterGenSQLPaginateWhereEntity

	gqttpl.DataVolumeMap
}

func (t *SQLParameterGenSQLPaginateEntity) TplName() string {
	return "sql.parameter.gen.sql.Paginate"
}

type SQLParameterGenSQLPaginateTotalEntity struct {
	SQLParameterGenSQLPaginateWhereEntity

	gqttpl.DataVolumeMap
}

func (t *SQLParameterGenSQLPaginateTotalEntity) TplName() string {
	return "sql.parameter.gen.sql.PaginateTotal"
}

type SQLParameterGenSQLPaginateWhereEntity struct {
	gqttpl.DataVolumeMap
}

func (t *SQLParameterGenSQLPaginateWhereEntity) TplName() string {
	return "sql.parameter.gen.sql.PaginateWhere"
}

type SQLParameterGenSQLUpdateEntity struct {
	APIID string

	AllowEmptyValue string

	AllowReserved string

	Deprecated string

	Description string

	Example string

	Explode string

	FullName string

	HTTPStatus string

	Method string

	Name string

	ParameterID string

	Position string

	Required string

	Serialize string

	ServiceID string

	Tag string

	Title string

	Type string

	ValidateSchemaID string

	gqttpl.DataVolumeMap
}

func (t *SQLParameterGenSQLUpdateEntity) TplName() string {
	return "sql.parameter.gen.sql.Update"
}

type SQLServerGenSQLGetAllByServerIDListEntity struct {
	ServerIDList []string

	gqttpl.DataVolumeMap
}

func (t *SQLServerGenSQLGetAllByServerIDListEntity) TplName() string {
	return "sql.server.gen.sql.GetAllByServerIDList"
}

type SQLServerGenSQLGetByServerIDEntity struct {
	ServerID string

	gqttpl.DataVolumeMap
}

func (t *SQLServerGenSQLGetByServerIDEntity) TplName() string {
	return "sql.server.gen.sql.GetByServerID"
}

type SQLServerGenSQLInsertEntity struct {
	Description string

	ExtensionIds string

	Proxy string

	ServerID string

	ServiceID string

	URL string

	gqttpl.DataVolumeMap
}

func (t *SQLServerGenSQLInsertEntity) TplName() string {
	return "sql.server.gen.sql.Insert"
}

type SQLServerGenSQLPaginateEntity struct {
	Limit int

	Offset int

	SQLServerGenSQLPaginateWhereEntity

	gqttpl.DataVolumeMap
}

func (t *SQLServerGenSQLPaginateEntity) TplName() string {
	return "sql.server.gen.sql.Paginate"
}

type SQLServerGenSQLPaginateTotalEntity struct {
	SQLServerGenSQLPaginateWhereEntity

	gqttpl.DataVolumeMap
}

func (t *SQLServerGenSQLPaginateTotalEntity) TplName() string {
	return "sql.server.gen.sql.PaginateTotal"
}

type SQLServerGenSQLPaginateWhereEntity struct {
	gqttpl.DataVolumeMap
}

func (t *SQLServerGenSQLPaginateWhereEntity) TplName() string {
	return "sql.server.gen.sql.PaginateWhere"
}

type SQLServerGenSQLUpdateEntity struct {
	Description string

	ExtensionIds string

	Proxy string

	ServerID string

	ServiceID string

	URL string

	gqttpl.DataVolumeMap
}

func (t *SQLServerGenSQLUpdateEntity) TplName() string {
	return "sql.server.gen.sql.Update"
}

type SQLServiceGenSQLGetAllByServiceIDListEntity struct {
	ServiceIDList []string

	gqttpl.DataVolumeMap
}

func (t *SQLServiceGenSQLGetAllByServiceIDListEntity) TplName() string {
	return "sql.service.gen.sql.GetAllByServiceIDList"
}

type SQLServiceGenSQLGetByServiceIDEntity struct {
	ServiceID string

	gqttpl.DataVolumeMap
}

func (t *SQLServiceGenSQLGetByServiceIDEntity) TplName() string {
	return "sql.service.gen.sql.GetByServiceID"
}

type SQLServiceGenSQLInsertEntity struct {
	ContactIds string

	Description string

	License string

	Proxy string

	Security string

	ServiceID string

	Title string

	Variables string

	Version string

	gqttpl.DataVolumeMap
}

func (t *SQLServiceGenSQLInsertEntity) TplName() string {
	return "sql.service.gen.sql.Insert"
}

type SQLServiceGenSQLPaginateEntity struct {
	Limit int

	Offset int

	SQLServiceGenSQLPaginateWhereEntity

	gqttpl.DataVolumeMap
}

func (t *SQLServiceGenSQLPaginateEntity) TplName() string {
	return "sql.service.gen.sql.Paginate"
}

type SQLServiceGenSQLPaginateTotalEntity struct {
	SQLServiceGenSQLPaginateWhereEntity

	gqttpl.DataVolumeMap
}

func (t *SQLServiceGenSQLPaginateTotalEntity) TplName() string {
	return "sql.service.gen.sql.PaginateTotal"
}

type SQLServiceGenSQLPaginateWhereEntity struct {
	gqttpl.DataVolumeMap
}

func (t *SQLServiceGenSQLPaginateWhereEntity) TplName() string {
	return "sql.service.gen.sql.PaginateWhere"
}

type SQLServiceGenSQLUpdateEntity struct {
	ContactIds string

	Description string

	License string

	Proxy string

	Security string

	ServiceID string

	Title string

	Variables string

	Version string

	gqttpl.DataVolumeMap
}

func (t *SQLServiceGenSQLUpdateEntity) TplName() string {
	return "sql.service.gen.sql.Update"
}

type SQLValidateSchemaGenSQLGetAllByValidateSchemaIDListEntity struct {
	ValidateSchemaIDList []string

	gqttpl.DataVolumeMap
}

func (t *SQLValidateSchemaGenSQLGetAllByValidateSchemaIDListEntity) TplName() string {
	return "sql.validate_schema.gen.sql.GetAllByValidateSchemaIDList"
}

type SQLValidateSchemaGenSQLGetByValidateSchemaIDEntity struct {
	ValidateSchemaID string

	gqttpl.DataVolumeMap
}

func (t *SQLValidateSchemaGenSQLGetByValidateSchemaIDEntity) TplName() string {
	return "sql.validate_schema.gen.sql.GetByValidateSchemaID"
}

type SQLValidateSchemaGenSQLInsertEntity struct {
	AdditionalProperties string

	AllOf string

	AllowEmptyValue string

	AllowReserved string

	AnyOf string

	Default string

	Deprecated string

	Description string

	Discriminator string

	Enum string

	EnumNames string

	EnumTitles string

	Example string

	ExclusiveMaximum string

	ExclusiveMinimum string

	Extensions string

	ExternalDocs string

	ExternalPros string

	Format string

	MaxItems int

	MaxLength int

	MaxProperties int

	Maxnum int

	MinItems int

	MinLength int

	MinProperties int

	Minimum int

	MultipleOf int

	Not string

	Nullable string

	OneOf string

	Pattern string

	ReadOnly string

	Remark string

	Required string

	ServiceID string

	Summary string

	Type string

	UniqueItems string

	ValidateSchemaID string

	WriteOnly string

	XML string

	gqttpl.DataVolumeMap
}

func (t *SQLValidateSchemaGenSQLInsertEntity) TplName() string {
	return "sql.validate_schema.gen.sql.Insert"
}

type SQLValidateSchemaGenSQLPaginateEntity struct {
	Limit int

	Offset int

	SQLValidateSchemaGenSQLPaginateWhereEntity

	gqttpl.DataVolumeMap
}

func (t *SQLValidateSchemaGenSQLPaginateEntity) TplName() string {
	return "sql.validate_schema.gen.sql.Paginate"
}

type SQLValidateSchemaGenSQLPaginateTotalEntity struct {
	SQLValidateSchemaGenSQLPaginateWhereEntity

	gqttpl.DataVolumeMap
}

func (t *SQLValidateSchemaGenSQLPaginateTotalEntity) TplName() string {
	return "sql.validate_schema.gen.sql.PaginateTotal"
}

type SQLValidateSchemaGenSQLPaginateWhereEntity struct {
	gqttpl.DataVolumeMap
}

func (t *SQLValidateSchemaGenSQLPaginateWhereEntity) TplName() string {
	return "sql.validate_schema.gen.sql.PaginateWhere"
}

type SQLValidateSchemaGenSQLUpdateEntity struct {
	AdditionalProperties string

	AllOf string

	AllowEmptyValue string

	AllowReserved string

	AnyOf string

	Default string

	Deprecated string

	Description string

	Discriminator string

	Enum string

	EnumNames string

	EnumTitles string

	Example string

	ExclusiveMaximum string

	ExclusiveMinimum string

	Extensions string

	ExternalDocs string

	ExternalPros string

	Format string

	MaxItems int

	MaxLength int

	MaxProperties int

	Maxnum int

	MinItems int

	MinLength int

	MinProperties int

	Minimum int

	MultipleOf int

	Not string

	Nullable string

	OneOf string

	Pattern string

	ReadOnly string

	Remark string

	Required string

	ServiceID string

	Summary string

	Type string

	UniqueItems string

	ValidateSchemaID string

	WriteOnly string

	XML string

	gqttpl.DataVolumeMap
}

func (t *SQLValidateSchemaGenSQLUpdateEntity) TplName() string {
	return "sql.validate_schema.gen.sql.Update"
}
