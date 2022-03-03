package example

type APIDel struct {
	APIID string
}

func (t *APIDel) TplName() string {
	return "api.Del"
}
func (t *APIDel) TplInput() interface{} {
	return t
}

type APIGetByAPIID struct {
	APIID string
}

func (t *APIGetByAPIID) TplName() string {
	return "api.GetByAPIID"
}
func (t *APIGetByAPIID) TplInput() interface{} {
	return t
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

func (t *APIInsert) TplName() string {
	return "api.Insert"
}
func (t *APIInsert) TplInput() interface{} {
	return t
}

type APIPaginate struct {
	Limit int

	Offset int
}

func (t *APIPaginate) TplName() string {
	return "api.Paginate"
}
func (t *APIPaginate) TplInput() interface{} {
	return t
}

type APITotal struct {
}

func (t *APITotal) TplName() string {
	return "api.Total"
}
func (t *APITotal) TplInput() interface{} {
	return t
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

func (t *APIUpdate) TplName() string {
	return "api.Update"
}
func (t *APIUpdate) TplInput() interface{} {
	return t
}

type ExampleDel struct {
	ExampleID string
}

func (t *ExampleDel) TplName() string {
	return "example.Del"
}
func (t *ExampleDel) TplInput() interface{} {
	return t
}

type ExampleGetByExampleID struct {
	ExampleID string
}

func (t *ExampleGetByExampleID) TplName() string {
	return "example.GetByExampleID"
}
func (t *ExampleGetByExampleID) TplInput() interface{} {
	return t
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

func (t *ExampleInsert) TplName() string {
	return "example.Insert"
}
func (t *ExampleInsert) TplInput() interface{} {
	return t
}

type ExamplePaginate struct {
	Limit int

	Offset int
}

func (t *ExamplePaginate) TplName() string {
	return "example.Paginate"
}
func (t *ExamplePaginate) TplInput() interface{} {
	return t
}

type ExampleTotal struct {
}

func (t *ExampleTotal) TplName() string {
	return "example.Total"
}
func (t *ExampleTotal) TplInput() interface{} {
	return t
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

func (t *ExampleUpdate) TplName() string {
	return "example.Update"
}
func (t *ExampleUpdate) TplInput() interface{} {
	return t
}

type MarkdownDel struct {
	MarkdownID string
}

func (t *MarkdownDel) TplName() string {
	return "markdown.Del"
}
func (t *MarkdownDel) TplInput() interface{} {
	return t
}

type MarkdownGetByMarkdownID struct {
	MarkdownID string
}

func (t *MarkdownGetByMarkdownID) TplName() string {
	return "markdown.GetByMarkdownID"
}
func (t *MarkdownGetByMarkdownID) TplInput() interface{} {
	return t
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

func (t *MarkdownInsert) TplName() string {
	return "markdown.Insert"
}
func (t *MarkdownInsert) TplInput() interface{} {
	return t
}

type MarkdownPaginate struct {
	Limit int

	Offset int
}

func (t *MarkdownPaginate) TplName() string {
	return "markdown.Paginate"
}
func (t *MarkdownPaginate) TplInput() interface{} {
	return t
}

type MarkdownTotal struct {
}

func (t *MarkdownTotal) TplName() string {
	return "markdown.Total"
}
func (t *MarkdownTotal) TplInput() interface{} {
	return t
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

func (t *MarkdownUpdate) TplName() string {
	return "markdown.Update"
}
func (t *MarkdownUpdate) TplInput() interface{} {
	return t
}

type ParameterDel struct {
	ParameterID string
}

func (t *ParameterDel) TplName() string {
	return "parameter.Del"
}
func (t *ParameterDel) TplInput() interface{} {
	return t
}

type ParameterGetByParameterID struct {
	ParameterID string
}

func (t *ParameterGetByParameterID) TplName() string {
	return "parameter.GetByParameterID"
}
func (t *ParameterGetByParameterID) TplInput() interface{} {
	return t
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

func (t *ParameterInsert) TplName() string {
	return "parameter.Insert"
}
func (t *ParameterInsert) TplInput() interface{} {
	return t
}

type ParameterPaginate struct {
	Limit int

	Offset int
}

func (t *ParameterPaginate) TplName() string {
	return "parameter.Paginate"
}
func (t *ParameterPaginate) TplInput() interface{} {
	return t
}

type ParameterTotal struct {
}

func (t *ParameterTotal) TplName() string {
	return "parameter.Total"
}
func (t *ParameterTotal) TplInput() interface{} {
	return t
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

func (t *ParameterUpdate) TplName() string {
	return "parameter.Update"
}
func (t *ParameterUpdate) TplInput() interface{} {
	return t
}

type ServerDel struct {
	ServerID string
}

func (t *ServerDel) TplName() string {
	return "server.Del"
}
func (t *ServerDel) TplInput() interface{} {
	return t
}

type ServerGetByServerID struct {
	ServerID string
}

func (t *ServerGetByServerID) TplName() string {
	return "server.GetByServerID"
}
func (t *ServerGetByServerID) TplInput() interface{} {
	return t
}

type ServerInsert struct {
	Description string

	ExtensionIds string

	Proxy string

	ServerID string

	ServiceID string

	URL string
}

func (t *ServerInsert) TplName() string {
	return "server.Insert"
}
func (t *ServerInsert) TplInput() interface{} {
	return t
}

type ServerPaginate struct {
	Limit int

	Offset int
}

func (t *ServerPaginate) TplName() string {
	return "server.Paginate"
}
func (t *ServerPaginate) TplInput() interface{} {
	return t
}

type ServerTotal struct {
}

func (t *ServerTotal) TplName() string {
	return "server.Total"
}
func (t *ServerTotal) TplInput() interface{} {
	return t
}

type ServerUpdate struct {
	Description string

	ExtensionIds string

	Proxy string

	ServerID string

	ServiceID string

	URL string
}

func (t *ServerUpdate) TplName() string {
	return "server.Update"
}
func (t *ServerUpdate) TplInput() interface{} {
	return t
}

type ServiceDel struct {
	ServiceID string
}

func (t *ServiceDel) TplName() string {
	return "service.del"
}
func (t *ServiceDel) TplInput() interface{} {
	return t
}

type ServiceGetByServiceID struct {
	ServiceID string
}

func (t *ServiceGetByServiceID) TplName() string {
	return "service.getByServiceId"
}
func (t *ServiceGetByServiceID) TplInput() interface{} {
	return t
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

func (t *ServiceInsert) TplName() string {
	return "service.insert"
}
func (t *ServiceInsert) TplInput() interface{} {
	return t
}

type ServiceList struct {
	Limit int

	Offset int

	Title string
}

func (t *ServiceList) TplName() string {
	return "service.list"
}
func (t *ServiceList) TplInput() interface{} {
	return t
}

type ServiceTotal struct {
	Title string
}

func (t *ServiceTotal) TplName() string {
	return "service.total"
}
func (t *ServiceTotal) TplInput() interface{} {
	return t
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

func (t *ServiceUpdate) TplName() string {
	return "service.update"
}
func (t *ServiceUpdate) TplInput() interface{} {
	return t
}

type ValidateSchemaDel struct {
	ValidateSchemaID string
}

func (t *ValidateSchemaDel) TplName() string {
	return "validate_schema.Del"
}
func (t *ValidateSchemaDel) TplInput() interface{} {
	return t
}

type ValidateSchemaGetByValidateSchemaID struct {
	ValidateSchemaID string
}

func (t *ValidateSchemaGetByValidateSchemaID) TplName() string {
	return "validate_schema.GetByValidateSchemaID"
}
func (t *ValidateSchemaGetByValidateSchemaID) TplInput() interface{} {
	return t
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

func (t *ValidateSchemaInsert) TplName() string {
	return "validate_schema.Insert"
}
func (t *ValidateSchemaInsert) TplInput() interface{} {
	return t
}

type ValidateSchemaPaginate struct {
	Limit int

	Offset int
}

func (t *ValidateSchemaPaginate) TplName() string {
	return "validate_schema.Paginate"
}
func (t *ValidateSchemaPaginate) TplInput() interface{} {
	return t
}

type ValidateSchemaTotal struct {
}

func (t *ValidateSchemaTotal) TplName() string {
	return "validate_schema.Total"
}
func (t *ValidateSchemaTotal) TplInput() interface{} {
	return t
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

func (t *ValidateSchemaUpdate) TplName() string {
	return "validate_schema.Update"
}
func (t *ValidateSchemaUpdate) TplInput() interface{} {
	return t
}
