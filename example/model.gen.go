package example

type APIModel struct {
	// api 文档标识
	APIID string `json:"apiID"`
	// 服务标识
	ServiceID string `json:"serviceID"`
	// 路由名称(英文)
	Name string `json:"name"`
	// 标题
	Title string `json:"title"`
	// 标签ID集合
	Tags string `json:"tags"`
	// 路径
	URI string `json:"uri"`
	// 摘要
	Summary string `json:"summary"`
	// 介绍
	Description string `json:"description"`
	// 创建时间
	CreatedAt string `json:"createdAt"`
	// 修改时间
	UpdatedAt string `json:"updatedAt"`
	// 删除时间
	DeletedAt string `json:"deletedAt"`
}

func (t *APIModel) TableName() string {
	return "t_api"
}
func (t *APIModel) PrimaryKey() string {
	return "api_id"
}
func (t *APIModel) PrimaryKeyCamel() string {
	return "APIID"
}

type ExampleModel struct {
	// 测试用例标识
	ExampleID string `json:"exampleID"`
	// 服务标识
	ServiceID string `json:"serviceID"`
	// api 文档标识
	APIID string `json:"apiID"`
	// 标签,mock数据时不同接口案例优先返回相同tag案例
	Tag string `json:"tag"`
	// 案例名称
	Title string `json:"title"`
	// 简介
	Summary string `json:"summary"`
	//
	Request string `json:"request"`
	//
	Response string `json:"response"`
	// 创建时间
	CreatedAt string `json:"createdAt"`
	// 修改时间
	UpdatedAt string `json:"updatedAt"`
	// 删除时间
	DeletedAt string `json:"deletedAt"`
}

func (t *ExampleModel) TableName() string {
	return "t_example"
}
func (t *ExampleModel) PrimaryKey() string {
	return "example_id"
}
func (t *ExampleModel) PrimaryKeyCamel() string {
	return "ExampleID"
}

type MarkdownModel struct {
	// markdown 文档标识
	MarkdownID string `json:"markdownID"`
	// 服务标识
	ServiceID string `json:"serviceID"`
	// api 文档标识
	APIID string `json:"apiID"`
	// 唯一名称
	Name string `json:"name"`
	// 文章标题
	Title string `json:"title"`
	//
	Markdown string `json:"markdown"`
	//
	Content string `json:"content"`
	// 作者ID
	OwnerID int `json:"ownerID"`
	// 作者名称
	OwnerName string `json:"ownerName"`
	// 创建时间
	CreatedAt string `json:"createdAt"`
	// 修改时间
	UpdatedAt string `json:"updatedAt"`
	// 删除时间
	DeletedAt string `json:"deletedAt"`
}

func (t *MarkdownModel) TableName() string {
	return "t_markdown"
}
func (t *MarkdownModel) PrimaryKey() string {
	return "markdown_id"
}
func (t *MarkdownModel) PrimaryKeyCamel() string {
	return "MarkdownID"
}

const (
	PARAMETER_TYPE_STRING             = "string"
	PARAMETER_TYPE_INT                = "int"
	PARAMETER_TYPE_NUMBER             = "number"
	PARAMETER_TYPE_ARRAY              = "array"
	PARAMETER_TYPE_OBJECT             = "object"
	PARAMETER_METHOD_GET              = "get"
	PARAMETER_METHOD_POST             = "post"
	PARAMETER_METHOD_PUT              = "put"
	PARAMETER_METHOD_DELETE           = "delete"
	PARAMETER_METHOD_HEAD             = "head"
	PARAMETER_POSITION_BODY           = "body"
	PARAMETER_POSITION_HEAD           = "head"
	PARAMETER_POSITION_PATH           = "path"
	PARAMETER_POSITION_QUERY          = "query"
	PARAMETER_POSITION_COOKIE         = "cookie"
	PARAMETER_DEPRECATED_TRUE         = "true"
	PARAMETER_DEPRECATED_FALSE        = "false"
	PARAMETER_REQUIRED_TRUE           = "true"
	PARAMETER_REQUIRED_FALSE          = "false"
	PARAMETER_EXPLODE_TRUE            = "true"
	PARAMETER_EXPLODE_FALSE           = "false"
	PARAMETER_ALLOW_EMPTY_VALUE_TRUE  = "true"
	PARAMETER_ALLOW_EMPTY_VALUE_FALSE = "false"
	PARAMETER_ALLOW_RESERVED_TRUE     = "true"
	PARAMETER_ALLOW_RESERVED_FALSE    = "false"
)

type ParameterModel struct {
	// 参数标识
	ParameterID string `json:"parameterID"`
	// 服务标识
	ServiceID string `json:"serviceID"`
	// api 文档标识
	APIID string `json:"apiID"`
	// 验证规则标识
	ValidateSchemaID string `json:"validateSchemaID"`
	// 全称
	FullName string `json:"fullName"`
	// 名称(冗余local.en)
	Name string `json:"name"`
	// 名称(冗余local.zh)
	Title string `json:"title"`
	// 参数类型:(string-字符串,int-整型,number-数字,array-数组,object-对象)
	Type string `json:"type"`
	// 所属标签
	Tag string `json:"tag"`
	// 请求方法(get-GET,post-POST,put-PUT,delete-DELETE,head-HEAD)
	Method string `json:"method"`
	// http状态码
	HTTPStatus string `json:"httpStatus"`
	// 参数所在的位置(body-body,head-head,path-path,query-query,cookie-cookie)
	Position string `json:"position"`
	// 案例
	Example string `json:"example"`
	// 是否弃用(true-是,false-否)
	Deprecated string `json:"deprecated"`
	// 是否必须(true-是,false-否)
	Required string `json:"required"`
	// 对数组、对象序列化方法,参照openapi parameters.style
	Serialize string `json:"serialize"`
	// 对象的key,是否单独成参数方式,参照openapi parameters.explode(true-是,false-否)
	Explode string `json:"explode"`
	// 是否容许空值(true-是,false-否)
	AllowEmptyValue string `json:"allowEmptyValue"`
	// 特殊字符是否容许出现在uri参数中(true-是,false-否)
	AllowReserved string `json:"allowReserved"`
	// 简介
	Description string `json:"description"`
	// 创建时间
	CreatedAt string `json:"createdAt"`
	// 修改时间
	UpdatedAt string `json:"updatedAt"`
	// 删除时间
	DeletedAt string `json:"deletedAt"`
}

func (t *ParameterModel) TableName() string {
	return "t_parameter"
}
func (t *ParameterModel) PrimaryKey() string {
	return "parameter_id"
}
func (t *ParameterModel) PrimaryKeyCamel() string {
	return "ParameterID"
}
func (t *ParameterModel) RequiredTitleMap() map[string]string {
	enumMap := make(map[string]string)
	enumMap[PARAMETER_REQUIRED_TRUE] = "是"
	enumMap[PARAMETER_REQUIRED_FALSE] = "否"
	return enumMap
}
func (t *ParameterModel) RequiredTitle() string {
	enumMap := t.RequiredTitleMap()
	title, ok := enumMap[t.Required]
	if !ok {
		msg := "func RequiredTitle not found title by key " + t.Required
		panic(msg)
	}
	return title
}
func (t *ParameterModel) ExplodeTitleMap() map[string]string {
	enumMap := make(map[string]string)
	enumMap[PARAMETER_EXPLODE_TRUE] = "是"
	enumMap[PARAMETER_EXPLODE_FALSE] = "否"
	return enumMap
}
func (t *ParameterModel) ExplodeTitle() string {
	enumMap := t.ExplodeTitleMap()
	title, ok := enumMap[t.Explode]
	if !ok {
		msg := "func ExplodeTitle not found title by key " + t.Explode
		panic(msg)
	}
	return title
}
func (t *ParameterModel) AllowEmptyValueTitleMap() map[string]string {
	enumMap := make(map[string]string)
	enumMap[PARAMETER_ALLOW_EMPTY_VALUE_TRUE] = "是"
	enumMap[PARAMETER_ALLOW_EMPTY_VALUE_FALSE] = "否"
	return enumMap
}
func (t *ParameterModel) AllowEmptyValueTitle() string {
	enumMap := t.AllowEmptyValueTitleMap()
	title, ok := enumMap[t.AllowEmptyValue]
	if !ok {
		msg := "func AllowEmptyValueTitle not found title by key " + t.AllowEmptyValue
		panic(msg)
	}
	return title
}
func (t *ParameterModel) AllowReservedTitleMap() map[string]string {
	enumMap := make(map[string]string)
	enumMap[PARAMETER_ALLOW_RESERVED_TRUE] = "是"
	enumMap[PARAMETER_ALLOW_RESERVED_FALSE] = "否"
	return enumMap
}
func (t *ParameterModel) AllowReservedTitle() string {
	enumMap := t.AllowReservedTitleMap()
	title, ok := enumMap[t.AllowReserved]
	if !ok {
		msg := "func AllowReservedTitle not found title by key " + t.AllowReserved
		panic(msg)
	}
	return title
}
func (t *ParameterModel) TypeTitleMap() map[string]string {
	enumMap := make(map[string]string)
	enumMap[PARAMETER_TYPE_STRING] = "字符串"
	enumMap[PARAMETER_TYPE_INT] = "整型"
	enumMap[PARAMETER_TYPE_NUMBER] = "数字"
	enumMap[PARAMETER_TYPE_ARRAY] = "数组"
	enumMap[PARAMETER_TYPE_OBJECT] = "对象"
	return enumMap
}
func (t *ParameterModel) TypeTitle() string {
	enumMap := t.TypeTitleMap()
	title, ok := enumMap[t.Type]
	if !ok {
		msg := "func TypeTitle not found title by key " + t.Type
		panic(msg)
	}
	return title
}
func (t *ParameterModel) MethodTitleMap() map[string]string {
	enumMap := make(map[string]string)
	enumMap[PARAMETER_METHOD_GET] = "GET"
	enumMap[PARAMETER_METHOD_POST] = "POST"
	enumMap[PARAMETER_METHOD_PUT] = "PUT"
	enumMap[PARAMETER_METHOD_DELETE] = "DELETE"
	enumMap[PARAMETER_METHOD_HEAD] = "HEAD"
	return enumMap
}
func (t *ParameterModel) MethodTitle() string {
	enumMap := t.MethodTitleMap()
	title, ok := enumMap[t.Method]
	if !ok {
		msg := "func MethodTitle not found title by key " + t.Method
		panic(msg)
	}
	return title
}
func (t *ParameterModel) PositionTitleMap() map[string]string {
	enumMap := make(map[string]string)
	enumMap[PARAMETER_POSITION_BODY] = "body"
	enumMap[PARAMETER_POSITION_HEAD] = "head"
	enumMap[PARAMETER_POSITION_PATH] = "path"
	enumMap[PARAMETER_POSITION_QUERY] = "query"
	enumMap[PARAMETER_POSITION_COOKIE] = "cookie"
	return enumMap
}
func (t *ParameterModel) PositionTitle() string {
	enumMap := t.PositionTitleMap()
	title, ok := enumMap[t.Position]
	if !ok {
		msg := "func PositionTitle not found title by key " + t.Position
		panic(msg)
	}
	return title
}
func (t *ParameterModel) DeprecatedTitleMap() map[string]string {
	enumMap := make(map[string]string)
	enumMap[PARAMETER_DEPRECATED_TRUE] = "是"
	enumMap[PARAMETER_DEPRECATED_FALSE] = "否"
	return enumMap
}
func (t *ParameterModel) DeprecatedTitle() string {
	enumMap := t.DeprecatedTitleMap()
	title, ok := enumMap[t.Deprecated]
	if !ok {
		msg := "func DeprecatedTitle not found title by key " + t.Deprecated
		panic(msg)
	}
	return title
}

type ServerModel struct {
	// 服务标识
	ServerID string `json:"serverID"`
	// 服务标识
	ServiceID string `json:"serviceID"`
	// 服务器地址
	URL string `json:"url"`
	// 介绍
	Description string `json:"description"`
	// 代理地址
	Proxy string `json:"proxy"`
	// 扩展字段
	ExtensionIds string `json:"extensionIds"`
	// 创建时间
	CreatedAt string `json:"createdAt"`
	// 修改时间
	UpdatedAt string `json:"updatedAt"`
	// 删除时间
	DeletedAt string `json:"deletedAt"`
}

func (t *ServerModel) TableName() string {
	return "t_server"
}
func (t *ServerModel) PrimaryKey() string {
	return "server_id"
}
func (t *ServerModel) PrimaryKeyCamel() string {
	return "ServerID"
}

type ServiceModel struct {
	// 服务标识
	ServiceID string `json:"serviceID"`
	// 标题
	Title string `json:"title"`
	// 介绍
	Description string `json:"description"`
	// 版本
	Version string `json:"version"`
	// 联系人
	ContactIds string `json:"contactIds"`
	// 协议
	License string `json:"license"`
	// 鉴权
	Security string `json:"security"`
	// 代理地址
	Proxy string `json:"proxy"`
	// json字符串
	Variables string `json:"variables"`
	// 创建时间
	CreatedAt string `json:"createdAt"`
	// 修改时间
	UpdatedAt string `json:"updatedAt"`
	// 删除时间
	DeletedAt string `json:"deletedAt"`
}

func (t *ServiceModel) TableName() string {
	return "t_service"
}
func (t *ServiceModel) PrimaryKey() string {
	return "service_id"
}
func (t *ServiceModel) PrimaryKeyCamel() string {
	return "ServiceID"
}

const (
	VALIDATE_SCHEMA_TYPE_INTEGER            = "integer"
	VALIDATE_SCHEMA_TYPE_ARRAY              = "array"
	VALIDATE_SCHEMA_TYPE_STRING             = "string"
	VALIDATE_SCHEMA_TYPE_OBJECT             = "object"
	VALIDATE_SCHEMA_DEPRECATED_TRUE         = "true"
	VALIDATE_SCHEMA_DEPRECATED_FALSE        = "false"
	VALIDATE_SCHEMA_REQUIRED_TRUE           = "true"
	VALIDATE_SCHEMA_REQUIRED_FALSE          = "false"
	VALIDATE_SCHEMA_NULLABLE_TRUE           = "true"
	VALIDATE_SCHEMA_NULLABLE_FALSE          = "false"
	VALIDATE_SCHEMA_EXCLUSIVE_MAXIMUM_TRUE  = "true"
	VALIDATE_SCHEMA_EXCLUSIVE_MAXIMUM_FALSE = "false"
	VALIDATE_SCHEMA_EXCLUSIVE_MINIMUM_TRUE  = "true"
	VALIDATE_SCHEMA_EXCLUSIVE_MINIMUM_FALSE = "false"
	VALIDATE_SCHEMA_UNIQUE_ITEMS_TRUE       = "true"
	VALIDATE_SCHEMA_UNIQUE_ITEMS_FALSE      = "false"
	VALIDATE_SCHEMA_ALLOW_EMPTY_VALUE_TRUE  = "true"
	VALIDATE_SCHEMA_ALLOW_EMPTY_VALUE_FALSE = "false"
	VALIDATE_SCHEMA_ALLOW_RESERVED_TRUE     = "true"
	VALIDATE_SCHEMA_ALLOW_RESERVED_FALSE    = "false"
	VALIDATE_SCHEMA_READ_ONLY_TRUE          = "true"
	VALIDATE_SCHEMA_READ_ONLY_FALSE         = "false"
	VALIDATE_SCHEMA_WRITE_ONLY_TRUE         = "true"
	VALIDATE_SCHEMA_WRITE_ONLY_FALSE        = "false"
)

type ValidateSchemaModel struct {
	// api schema 标识
	ValidateSchemaID string `json:"validateSchemaID"`
	// 所属服务标识
	ServiceID string `json:"serviceID"`
	// 描述
	Description string `json:"description"`
	// 备注
	Remark string `json:"remark"`
	// 参数类型:(string-字符串,integer-整型,array-数组,object-对象)
	Type string `json:"type"`
	// 案例
	Example string `json:"example"`
	// 是否弃用(true-是,false-否)
	Deprecated string `json:"deprecated"`
	// 是否必须(true-是,false-否)
	Required string `json:"required"`
	// 枚举值
	Enum string `json:"enum"`
	// 枚举名称
	EnumNames string `json:"enumNames"`
	// 枚举标题
	EnumTitles string `json:"enumTitles"`
	// 格式
	Format string `json:"format"`
	// 默认值
	Default string `json:"default"`
	// 是否可以为空(true-是,false-否)
	Nullable string `json:"nullable"`
	// 倍数
	MultipleOf int `json:"multipleOf"`
	// 最大值
	Maxnum int `json:"maxnum"`
	// 是否不包含最大项(true-是,false-否)
	ExclusiveMaximum string `json:"exclusiveMaximum"`
	// 最小值
	Minimum int `json:"minimum"`
	// 是否不包含最小项(true-是,false-否)
	ExclusiveMinimum string `json:"exclusiveMinimum"`
	// 最大长度
	MaxLength int `json:"maxLength"`
	// 最小长度
	MinLength int `json:"minLength"`
	// 正则表达式
	Pattern string `json:"pattern"`
	// 最大项数
	MaxItems int `json:"maxItems"`
	// 最小项数
	MinItems int `json:"minItems"`
	// 所有项是否需要唯一(true-是,false-否)
	UniqueItems string `json:"uniqueItems"`
	// 最多属性项
	MaxProperties int `json:"maxProperties"`
	// 最少属性项
	MinProperties int `json:"minProperties"`
	// 所有
	AllOf string `json:"allOf"`
	// 只满足一个
	OneOf string `json:"oneOf"`
	// 任何一个SchemaID
	AnyOf string `json:"anyOf"`
	// 是否容许空值(true-是,false-否)
	AllowEmptyValue string `json:"allowEmptyValue"`
	// 特殊字符是否容许出现在uri参数中(true-是,false-否)
	AllowReserved string `json:"allowReserved"`
	// 不包含的schemaID
	Not string `json:"not"`
	// boolean
	AdditionalProperties string `json:"additionalProperties"`
	// schema鉴别
	Discriminator string `json:"discriminator"`
	// 是否只读(true-是,false-否)
	ReadOnly string `json:"readOnly"`
	// 是否只写(true-是,false-否)
	WriteOnly string `json:"writeOnly"`
	// xml对象
	XML string `json:"xml"`
	// 附加文档
	ExternalDocs string `json:"externalDocs"`
	// 附加文档
	ExternalPros string `json:"externalPros"`
	// 扩展字段
	Extensions string `json:"extensions"`
	// 创建时间
	CreatedAt string `json:"createdAt"`
	// 修改时间
	UpdatedAt string `json:"updatedAt"`
	// 删除时间
	DeletedAt string `json:"deletedAt"`
	// 简介
	Summary string `json:"summary"`
}

func (t *ValidateSchemaModel) TableName() string {
	return "t_validate_schema"
}
func (t *ValidateSchemaModel) PrimaryKey() string {
	return "validate_schema_id"
}
func (t *ValidateSchemaModel) PrimaryKeyCamel() string {
	return "ValidateSchemaID"
}
func (t *ValidateSchemaModel) WriteOnlyTitleMap() map[string]string {
	enumMap := make(map[string]string)
	enumMap[VALIDATE_SCHEMA_WRITE_ONLY_TRUE] = "是"
	enumMap[VALIDATE_SCHEMA_WRITE_ONLY_FALSE] = "否"
	return enumMap
}
func (t *ValidateSchemaModel) WriteOnlyTitle() string {
	enumMap := t.WriteOnlyTitleMap()
	title, ok := enumMap[t.WriteOnly]
	if !ok {
		msg := "func WriteOnlyTitle not found title by key " + t.WriteOnly
		panic(msg)
	}
	return title
}
func (t *ValidateSchemaModel) TypeTitleMap() map[string]string {
	enumMap := make(map[string]string)
	enumMap[VALIDATE_SCHEMA_TYPE_INTEGER] = "整型"
	enumMap[VALIDATE_SCHEMA_TYPE_ARRAY] = "数组"
	enumMap[VALIDATE_SCHEMA_TYPE_STRING] = "字符串"
	enumMap[VALIDATE_SCHEMA_TYPE_OBJECT] = "对象"
	return enumMap
}
func (t *ValidateSchemaModel) TypeTitle() string {
	enumMap := t.TypeTitleMap()
	title, ok := enumMap[t.Type]
	if !ok {
		msg := "func TypeTitle not found title by key " + t.Type
		panic(msg)
	}
	return title
}
func (t *ValidateSchemaModel) RequiredTitleMap() map[string]string {
	enumMap := make(map[string]string)
	enumMap[VALIDATE_SCHEMA_REQUIRED_TRUE] = "是"
	enumMap[VALIDATE_SCHEMA_REQUIRED_FALSE] = "否"
	return enumMap
}
func (t *ValidateSchemaModel) RequiredTitle() string {
	enumMap := t.RequiredTitleMap()
	title, ok := enumMap[t.Required]
	if !ok {
		msg := "func RequiredTitle not found title by key " + t.Required
		panic(msg)
	}
	return title
}
func (t *ValidateSchemaModel) NullableTitleMap() map[string]string {
	enumMap := make(map[string]string)
	enumMap[VALIDATE_SCHEMA_NULLABLE_TRUE] = "是"
	enumMap[VALIDATE_SCHEMA_NULLABLE_FALSE] = "否"
	return enumMap
}
func (t *ValidateSchemaModel) NullableTitle() string {
	enumMap := t.NullableTitleMap()
	title, ok := enumMap[t.Nullable]
	if !ok {
		msg := "func NullableTitle not found title by key " + t.Nullable
		panic(msg)
	}
	return title
}
func (t *ValidateSchemaModel) ExclusiveMaximumTitleMap() map[string]string {
	enumMap := make(map[string]string)
	enumMap[VALIDATE_SCHEMA_EXCLUSIVE_MAXIMUM_TRUE] = "是"
	enumMap[VALIDATE_SCHEMA_EXCLUSIVE_MAXIMUM_FALSE] = "否"
	return enumMap
}
func (t *ValidateSchemaModel) ExclusiveMaximumTitle() string {
	enumMap := t.ExclusiveMaximumTitleMap()
	title, ok := enumMap[t.ExclusiveMaximum]
	if !ok {
		msg := "func ExclusiveMaximumTitle not found title by key " + t.ExclusiveMaximum
		panic(msg)
	}
	return title
}
func (t *ValidateSchemaModel) UniqueItemsTitleMap() map[string]string {
	enumMap := make(map[string]string)
	enumMap[VALIDATE_SCHEMA_UNIQUE_ITEMS_TRUE] = "是"
	enumMap[VALIDATE_SCHEMA_UNIQUE_ITEMS_FALSE] = "否"
	return enumMap
}
func (t *ValidateSchemaModel) UniqueItemsTitle() string {
	enumMap := t.UniqueItemsTitleMap()
	title, ok := enumMap[t.UniqueItems]
	if !ok {
		msg := "func UniqueItemsTitle not found title by key " + t.UniqueItems
		panic(msg)
	}
	return title
}
func (t *ValidateSchemaModel) DeprecatedTitleMap() map[string]string {
	enumMap := make(map[string]string)
	enumMap[VALIDATE_SCHEMA_DEPRECATED_TRUE] = "是"
	enumMap[VALIDATE_SCHEMA_DEPRECATED_FALSE] = "否"
	return enumMap
}
func (t *ValidateSchemaModel) DeprecatedTitle() string {
	enumMap := t.DeprecatedTitleMap()
	title, ok := enumMap[t.Deprecated]
	if !ok {
		msg := "func DeprecatedTitle not found title by key " + t.Deprecated
		panic(msg)
	}
	return title
}
func (t *ValidateSchemaModel) ExclusiveMinimumTitleMap() map[string]string {
	enumMap := make(map[string]string)
	enumMap[VALIDATE_SCHEMA_EXCLUSIVE_MINIMUM_TRUE] = "是"
	enumMap[VALIDATE_SCHEMA_EXCLUSIVE_MINIMUM_FALSE] = "否"
	return enumMap
}
func (t *ValidateSchemaModel) ExclusiveMinimumTitle() string {
	enumMap := t.ExclusiveMinimumTitleMap()
	title, ok := enumMap[t.ExclusiveMinimum]
	if !ok {
		msg := "func ExclusiveMinimumTitle not found title by key " + t.ExclusiveMinimum
		panic(msg)
	}
	return title
}
func (t *ValidateSchemaModel) AllowEmptyValueTitleMap() map[string]string {
	enumMap := make(map[string]string)
	enumMap[VALIDATE_SCHEMA_ALLOW_EMPTY_VALUE_TRUE] = "是"
	enumMap[VALIDATE_SCHEMA_ALLOW_EMPTY_VALUE_FALSE] = "否"
	return enumMap
}
func (t *ValidateSchemaModel) AllowEmptyValueTitle() string {
	enumMap := t.AllowEmptyValueTitleMap()
	title, ok := enumMap[t.AllowEmptyValue]
	if !ok {
		msg := "func AllowEmptyValueTitle not found title by key " + t.AllowEmptyValue
		panic(msg)
	}
	return title
}
func (t *ValidateSchemaModel) AllowReservedTitleMap() map[string]string {
	enumMap := make(map[string]string)
	enumMap[VALIDATE_SCHEMA_ALLOW_RESERVED_TRUE] = "是"
	enumMap[VALIDATE_SCHEMA_ALLOW_RESERVED_FALSE] = "否"
	return enumMap
}
func (t *ValidateSchemaModel) AllowReservedTitle() string {
	enumMap := t.AllowReservedTitleMap()
	title, ok := enumMap[t.AllowReserved]
	if !ok {
		msg := "func AllowReservedTitle not found title by key " + t.AllowReserved
		panic(msg)
	}
	return title
}
func (t *ValidateSchemaModel) ReadOnlyTitleMap() map[string]string {
	enumMap := make(map[string]string)
	enumMap[VALIDATE_SCHEMA_READ_ONLY_TRUE] = "是"
	enumMap[VALIDATE_SCHEMA_READ_ONLY_FALSE] = "否"
	return enumMap
}
func (t *ValidateSchemaModel) ReadOnlyTitle() string {
	enumMap := t.ReadOnlyTitleMap()
	title, ok := enumMap[t.ReadOnly]
	if !ok {
		msg := "func ReadOnlyTitle not found title by key " + t.ReadOnly
		panic(msg)
	}
	return title
}
