package example

type APIDel struct {
	APIID string
}

func (s *APIDel) TplName() string {
	return "api.Del"
}

type APIGetByAPIID struct {
	APIID string
}

func (s *APIGetByAPIID) TplName() string {
	return "api.GetByAPIID"
}

type APIInsert struct {
	APIID string

	Description string

	Name string

	ServiceID string

	Summary string

	Tags string

	Title string

	URI string
}

func (s *APIInsert) TplName() string {
	return "api.Insert"
}

type APIPaginate struct {
	Limit int

	Offset int
}

func (s *APIPaginate) TplName() string {
	return "api.Paginate"
}

type APITotal struct {
}

func (s *APITotal) TplName() string {
	return "api.Total"
}

type APIUpdate struct {
	APIID string

	Description string

	Name string

	ServiceID string

	Summary string

	Tags string

	Title string

	URI string
}

func (s *APIUpdate) TplName() string {
	return "api.Update"
}

type ExampleDel struct {
	ExampleID string
}

func (s *ExampleDel) TplName() string {
	return "example.Del"
}

type ExampleGetByExampleID struct {
	ExampleID string
}

func (s *ExampleGetByExampleID) TplName() string {
	return "example.GetByExampleID"
}

type ExampleInsert struct {
	APIID string

	ExampleID string

	Request string

	Response string

	ServiceID string

	Summary string

	Tag string

	Title string
}

func (s *ExampleInsert) TplName() string {
	return "example.Insert"
}

type ExamplePaginate struct {
	Limit int

	Offset int
}

func (s *ExamplePaginate) TplName() string {
	return "example.Paginate"
}

type ExampleTotal struct {
}

func (s *ExampleTotal) TplName() string {
	return "example.Total"
}

type ExampleUpdate struct {
	APIID string

	ExampleID string

	Request string

	Response string

	ServiceID string

	Summary string

	Tag string

	Title string
}

func (s *ExampleUpdate) TplName() string {
	return "example.Update"
}

type MarkdownDel struct {
	MarkdownID string
}

func (s *MarkdownDel) TplName() string {
	return "markdown.Del"
}

type MarkdownGetByMarkdownID struct {
	MarkdownID string
}

func (s *MarkdownGetByMarkdownID) TplName() string {
	return "markdown.GetByMarkdownID"
}

type MarkdownInsert struct {
	APIID string

	Content string

	Markdown string

	MarkdownID string

	Name string

	OwnerID int64

	OwnerName string

	ServiceID string

	Title string
}

func (s *MarkdownInsert) TplName() string {
	return "markdown.Insert"
}

type MarkdownPaginate struct {
	Limit int

	Offset int
}

func (s *MarkdownPaginate) TplName() string {
	return "markdown.Paginate"
}

type MarkdownTotal struct {
}

func (s *MarkdownTotal) TplName() string {
	return "markdown.Total"
}

type MarkdownUpdate struct {
	APIID string

	Content string

	Markdown string

	MarkdownID string

	Name string

	OwnerID int64

	OwnerName string

	ServiceID string

	Title string
}

func (s *MarkdownUpdate) TplName() string {
	return "markdown.Update"
}

type ParameterDel struct {
	ParameterID string
}

func (s *ParameterDel) TplName() string {
	return "parameter.Del"
}

type ParameterGetByParameterID struct {
	ParameterID string
}

func (s *ParameterGetByParameterID) TplName() string {
	return "parameter.GetByParameterID"
}

type ParameterInsert struct {
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
}

func (s *ParameterInsert) TplName() string {
	return "parameter.Insert"
}

type ParameterPaginate struct {
	Limit int

	Offset int
}

func (s *ParameterPaginate) TplName() string {
	return "parameter.Paginate"
}

type ParameterTotal struct {
}

func (s *ParameterTotal) TplName() string {
	return "parameter.Total"
}

type ParameterUpdate struct {
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
}

func (s *ParameterUpdate) TplName() string {
	return "parameter.Update"
}

type ServerDel struct {
	ServerID string
}

func (s *ServerDel) TplName() string {
	return "server.Del"
}

type ServerGetByServerID struct {
	ServerID string
}

func (s *ServerGetByServerID) TplName() string {
	return "server.GetByServerID"
}

type ServerInsert struct {
	Description string

	ExtensionIds string

	Proxy string

	ServerID string

	ServiceID string

	URL string
}

func (s *ServerInsert) TplName() string {
	return "server.Insert"
}

type ServerPaginate struct {
	Limit int

	Offset int
}

func (s *ServerPaginate) TplName() string {
	return "server.Paginate"
}

type ServerTotal struct {
}

func (s *ServerTotal) TplName() string {
	return "server.Total"
}

type ServerUpdate struct {
	Description string

	ExtensionIds string

	Proxy string

	ServerID string

	ServiceID string

	URL string
}

func (s *ServerUpdate) TplName() string {
	return "server.Update"
}

type ServiceDel struct {
	ServiceID string
}

func (s *ServiceDel) TplName() string {
	return "service.del"
}

type ServiceGetByServiceID struct {
	ServiceID string
}

func (s *ServiceGetByServiceID) TplName() string {
	return "service.getByServiceId"
}

type ServiceInsert struct {
	ContactIds string

	Description string

	License string

	Proxy string

	Security string

	ServiceID string

	Title string

	Variables string

	Version string
}

func (s *ServiceInsert) TplName() string {
	return "service.insert"
}

type ServiceList struct {
	Limit int

	Offset int

	Title string
}

func (s *ServiceList) TplName() string {
	return "service.list"
}

type ServiceTotal struct {
	Title string
}

func (s *ServiceTotal) TplName() string {
	return "service.total"
}

type ServiceUpdate struct {
	ContactIds string

	Description string

	License string

	Proxy string

	Security string

	ServiceID string

	Title string

	Variables string

	Version string
}

func (s *ServiceUpdate) TplName() string {
	return "service.update"
}

type ValidateSchemaDel struct {
	ValidateSchemaID string
}

func (s *ValidateSchemaDel) TplName() string {
	return "validate_schema.Del"
}

type ValidateSchemaGetByValidateSchemaID struct {
	ValidateSchemaID string
}

func (s *ValidateSchemaGetByValidateSchemaID) TplName() string {
	return "validate_schema.GetByValidateSchemaID"
}

type ValidateSchemaInsert struct {
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

	MaxItems int64

	MaxLength int64

	MaxProperties int64

	Maxnum int64

	MinItems int64

	MinLength int64

	MinProperties int64

	Minimum int64

	MultipleOf int64

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
}

func (s *ValidateSchemaInsert) TplName() string {
	return "validate_schema.Insert"
}

type ValidateSchemaPaginate struct {
	Limit int

	Offset int
}

func (s *ValidateSchemaPaginate) TplName() string {
	return "validate_schema.Paginate"
}

type ValidateSchemaTotal struct {
}

func (s *ValidateSchemaTotal) TplName() string {
	return "validate_schema.Total"
}

type ValidateSchemaUpdate struct {
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

	MaxItems int64

	MaxLength int64

	MaxProperties int64

	Maxnum int64

	MinItems int64

	MinLength int64

	MinProperties int64

	Minimum int64

	MultipleOf int64

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
}

func (s *ValidateSchemaUpdate) TplName() string {
	return "validate_schema.Update"
}
