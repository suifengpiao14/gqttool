package example

	
	type API struct{
		 
		// api 文档标识
		APIID string  `json:"apiID"` 
		 
		// 服务标识
		ServiceID string  `json:"serviceID"` 
		 
		// 路由名称(英文)
		Name string  `json:"name"` 
		 
		// 标题
		Title string  `json:"title"` 
		 
		// 标签ID集合
		Tags string  `json:"tags"` 
		 
		// 路径
		URI string  `json:"uri"` 
		 
		// 摘要
		Summary string  `json:"summary"` 
		 
		// 介绍
		Description string  `json:"description"` 
		 
		// 创建时间
		CreatedAt string  `json:"createdAt"` 
		 
		// 修改时间
		UpdatedAt string  `json:"updatedAt"` 
		 
		// 删除时间
		DeletedAt string  `json:"deletedAt"` 
		
	}
	

	
	type Example struct{
		 
		// 测试用例标识
		ExampleID string  `json:"exampleID"` 
		 
		// 服务标识
		ServiceID string  `json:"serviceID"` 
		 
		// api 文档标识
		APIID string  `json:"apiID"` 
		 
		// 标签,mock数据时不同接口案例优先返回相同tag案例
		Tag string  `json:"tag"` 
		 
		// 案例名称
		Title string  `json:"title"` 
		 
		// 简介
		Summary string  `json:"summary"` 
		 
		// 
		Request string  `json:"request"` 
		 
		// 
		Response string  `json:"response"` 
		 
		// 创建时间
		CreatedAt string  `json:"createdAt"` 
		 
		// 修改时间
		UpdatedAt string  `json:"updatedAt"` 
		 
		// 删除时间
		DeletedAt string  `json:"deletedAt"` 
		
	}
	

	
	type Markdown struct{
		 
		// markdown 文档标识
		MarkdownID string  `json:"markdownID"` 
		 
		// 服务标识
		ServiceID string  `json:"serviceID"` 
		 
		// api 文档标识
		APIID string  `json:"apiID"` 
		 
		// 唯一名称
		Name string  `json:"name"` 
		 
		// 文章标题
		Title string  `json:"title"` 
		 
		// 
		Markdown string  `json:"markdown"` 
		 
		// 
		Content string  `json:"content"` 
		 
		// 作者ID
		OwnerID int64  `json:"ownerID"` 
		 
		// 作者名称
		OwnerName string  `json:"ownerName"` 
		 
		// 创建时间
		CreatedAt string  `json:"createdAt"` 
		 
		// 修改时间
		UpdatedAt string  `json:"updatedAt"` 
		 
		// 删除时间
		DeletedAt string  `json:"deletedAt"` 
		
	}
	

	
	const (
		
			PARAMETER_ALLOW_EMPTY_VALUE_FALSE="false"
		
			PARAMETER_ALLOW_EMPTY_VALUE_TRUE="true"
		
			PARAMETER_ALLOW_RESERVED_FALSE="false"
		
			PARAMETER_ALLOW_RESERVED_TRUE="true"
		
			PARAMETER_DEPRECATED_FALSE="false"
		
			PARAMETER_DEPRECATED_TRUE="true"
		
			PARAMETER_EXPLODE_FALSE="false"
		
			PARAMETER_EXPLODE_TRUE="true"
		
			PARAMETER_METHOD_DELETE="delete"
		
			PARAMETER_METHOD_GET="get"
		
			PARAMETER_METHOD_HEAD="head"
		
			PARAMETER_METHOD_POST="post"
		
			PARAMETER_METHOD_PUT="put"
		
			PARAMETER_POSITION_=""
		
			PARAMETER_POSITION_BODY="body"
		
			PARAMETER_POSITION_COOKIE="cookie"
		
			PARAMETER_POSITION_HEAD="head"
		
			PARAMETER_POSITION_PATH="path"
		
			PARAMETER_POSITION_QUERY="query"
		
			PARAMETER_REQUIRED_FALSE="false"
		
			PARAMETER_REQUIRED_TRUE="true"
		
			PARAMETER_TYPE_ARRAY="array"
		
			PARAMETER_TYPE_INT="int"
		
			PARAMETER_TYPE_NUMBER="number"
		
			PARAMETER_TYPE_OBJECT="object"
		
			PARAMETER_TYPE_STRING="string"
		
		)
	
	type Parameter struct{
		 
		// 参数标识
		ParameterID string  `json:"parameterID"` 
		 
		// 服务标识
		ServiceID string  `json:"serviceID"` 
		 
		// api 文档标识
		APIID string  `json:"apiID"` 
		 
		// 验证规则标识
		ValidateSchemaID string  `json:"validateSchemaID"` 
		 
		// 全称
		FullName string  `json:"fullName"` 
		 
		// 名称(冗余local.en)
		Name string  `json:"name"` 
		 
		// 名称(冗余local.zh)
		Title string  `json:"title"` 
		 
		// 参数类型
		Type string  `json:"type"` 
		 
		// 所属标签
		Tag string  `json:"tag"` 
		 
		// 请求方法
		Method string  `json:"method"` 
		 
		// http状态码
		HTTPStatus string  `json:"httpStatus"` 
		 
		// 参数所在的位置
		Position string  `json:"position"` 
		 
		// 案例
		Example string  `json:"example"` 
		 
		// 是否弃用
		Deprecated string  `json:"deprecated"` 
		 
		// 是否必须(模型类为字符串)
		Required string  `json:"required"` 
		 
		// 对数组、对象序列化方法,参照openapi parameters.style
		Serialize string  `json:"serialize"` 
		 
		// 对象的key,是否单独成参数方式,参照openapi parameters.explode
		Explode string  `json:"explode"` 
		 
		// 是否容许空值
		AllowEmptyValue string  `json:"allowEmptyValue"` 
		 
		// 特殊字符是否容许出现在uri参数中
		AllowReserved string  `json:"allowReserved"` 
		 
		// 简介
		Description string  `json:"description"` 
		 
		// 创建时间
		CreatedAt string  `json:"createdAt"` 
		 
		// 修改时间
		UpdatedAt string  `json:"updatedAt"` 
		 
		// 删除时间
		DeletedAt string  `json:"deletedAt"` 
		
	}
	

	
	type Server struct{
		 
		// 服务标识
		ServerID string  `json:"serverID"` 
		 
		// 服务标识
		ServiceID string  `json:"serviceID"` 
		 
		// 服务器地址
		URL string  `json:"url"` 
		 
		// 介绍
		Description string  `json:"description"` 
		 
		// 代理地址
		Proxy string  `json:"proxy"` 
		 
		// 扩展字段
		ExtensionIds string  `json:"extensionIds"` 
		 
		// 创建时间
		CreatedAt string  `json:"createdAt"` 
		 
		// 修改时间
		UpdatedAt string  `json:"updatedAt"` 
		 
		// 删除时间
		DeletedAt string  `json:"deletedAt"` 
		
	}
	

	
	type Service struct{
		 
		// 服务标识
		ServiceID string  `json:"serviceID"` 
		 
		// 标题
		Title string  `json:"title"` 
		 
		// 介绍
		Description string  `json:"description"` 
		 
		// 版本
		Version string  `json:"version"` 
		 
		// 联系人
		ContactIds string  `json:"contactIds"` 
		 
		// 协议
		License string  `json:"license"` 
		 
		// 鉴权
		Security string  `json:"security"` 
		 
		// 代理地址
		Proxy string  `json:"proxy"` 
		 
		// json字符串
		Variables string  `json:"variables"` 
		 
		// 创建时间
		CreatedAt string  `json:"createdAt"` 
		 
		// 修改时间
		UpdatedAt string  `json:"updatedAt"` 
		 
		// 删除时间
		DeletedAt string  `json:"deletedAt"` 
		
	}
	

	
	const (
		
			VALIDATE_SCHEMA_ALLOW_EMPTY_VALUE_FALSE="false"
		
			VALIDATE_SCHEMA_ALLOW_EMPTY_VALUE_TRUE="true"
		
			VALIDATE_SCHEMA_ALLOW_RESERVED_FALSE="false"
		
			VALIDATE_SCHEMA_ALLOW_RESERVED_TRUE="true"
		
			VALIDATE_SCHEMA_DEPRECATED_FALSE="false"
		
			VALIDATE_SCHEMA_DEPRECATED_TRUE="true"
		
			VALIDATE_SCHEMA_EXCLUSIVE_MAXIMUM_FALSE="false"
		
			VALIDATE_SCHEMA_EXCLUSIVE_MAXIMUM_TRUE="true"
		
			VALIDATE_SCHEMA_EXCLUSIVE_MINIMUM_FALSE="false"
		
			VALIDATE_SCHEMA_EXCLUSIVE_MINIMUM_TRUE="true"
		
			VALIDATE_SCHEMA_NULLABLE_FALSE="false"
		
			VALIDATE_SCHEMA_NULLABLE_TRUE="true"
		
			VALIDATE_SCHEMA_READ_ONLY_FALSE="false"
		
			VALIDATE_SCHEMA_READ_ONLY_TRUE="true"
		
			VALIDATE_SCHEMA_REQUIRED_FALSE="false"
		
			VALIDATE_SCHEMA_REQUIRED_TRUE="true"
		
			VALIDATE_SCHEMA_TYPE_ARRAY="array"
		
			VALIDATE_SCHEMA_TYPE_INTEGER="integer"
		
			VALIDATE_SCHEMA_TYPE_OBJECT="object"
		
			VALIDATE_SCHEMA_TYPE_STRING="string"
		
			VALIDATE_SCHEMA_UNIQUE_ITEMS_FALSE="false"
		
			VALIDATE_SCHEMA_UNIQUE_ITEMS_TRUE="true"
		
			VALIDATE_SCHEMA_WRITE_ONLY_FALSE="false"
		
			VALIDATE_SCHEMA_WRITE_ONLY_TRUE="true"
		
		)
	
	type ValidateSchema struct{
		 
		// api schema 标识
		ValidateSchemaID string  `json:"validateSchemaID"` 
		 
		// 所属服务标识
		ServiceID string  `json:"serviceID"` 
		 
		// 描述
		Description string  `json:"description"` 
		 
		// 备注
		Remark string  `json:"remark"` 
		 
		// 类型
		Type string  `json:"type"` 
		 
		// 案例
		Example string  `json:"example"` 
		 
		// 是否弃用
		Deprecated string  `json:"deprecated"` 
		 
		// 是否必须(模型类为字符串)
		Required string  `json:"required"` 
		 
		// 枚举值
		Enum string  `json:"enum"` 
		 
		// 枚举名称
		EnumNames string  `json:"enumNames"` 
		 
		// 枚举标题
		EnumTitles string  `json:"enumTitles"` 
		 
		// 格式
		Format string  `json:"format"` 
		 
		// 默认值
		Default string  `json:"default"` 
		 
		// 是否可以为空
		Nullable string  `json:"nullable"` 
		 
		// 倍数
		MultipleOf int64  `json:"multipleOf"` 
		 
		// 最大值
		Maxnum int64  `json:"maxnum"` 
		 
		// 是否不包含最大项
		ExclusiveMaximum string  `json:"exclusiveMaximum"` 
		 
		// 最小值
		Minimum int64  `json:"minimum"` 
		 
		// 是否不包含最小项
		ExclusiveMinimum string  `json:"exclusiveMinimum"` 
		 
		// 最大长度
		MaxLength int64  `json:"maxLength"` 
		 
		// 最小长度
		MinLength int64  `json:"minLength"` 
		 
		// 正则表达式
		Pattern string  `json:"pattern"` 
		 
		// 最大项数
		MaxItems int64  `json:"maxItems"` 
		 
		// 最小项数
		MinItems int64  `json:"minItems"` 
		 
		// 所有项是否需要唯一
		UniqueItems string  `json:"uniqueItems"` 
		 
		// 最多属性项
		MaxProperties int64  `json:"maxProperties"` 
		 
		// 最少属性项
		MinProperties int64  `json:"minProperties"` 
		 
		// 所有
		AllOf string  `json:"allOf"` 
		 
		// 只满足一个
		OneOf string  `json:"oneOf"` 
		 
		// 任何一个SchemaID
		AnyOf string  `json:"anyOf"` 
		 
		// 是否容许空值
		AllowEmptyValue string  `json:"allowEmptyValue"` 
		 
		// 特殊字符是否容许出现在uri参数中
		AllowReserved string  `json:"allowReserved"` 
		 
		// 不包含的schemaID
		Not string  `json:"not"` 
		 
		// boolean
		AdditionalProperties string  `json:"additionalProperties"` 
		 
		// schema鉴别
		Discriminator string  `json:"discriminator"` 
		 
		// 是否只读
		ReadOnly string  `json:"readOnly"` 
		 
		// 是否只写
		WriteOnly string  `json:"writeOnly"` 
		 
		// xml对象
		XML string  `json:"xml"` 
		 
		// 附加文档
		ExternalDocs string  `json:"externalDocs"` 
		 
		// 附加文档
		ExternalPros string  `json:"externalPros"` 
		 
		// 扩展字段
		Extensions string  `json:"extensions"` 
		 
		// 创建时间
		CreatedAt string  `json:"createdAt"` 
		 
		// 修改时间
		UpdatedAt string  `json:"updatedAt"` 
		 
		// 删除时间
		DeletedAt string  `json:"deletedAt"` 
		 
		// 简介
		Summary string  `json:"summary"` 
		
	}
	