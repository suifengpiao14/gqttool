package example

		type APIAutoSQLDelEntity struct{
			
				
					APIID string 
				
				
			
				
					Operator interface{} 
				
				
			
				
					OperatorID interface{} 
				
				
			
		}

		func (t *APIAutoSQLDelEntity) TplName() string{
			return "api.auto.sql.Del"
		}

	

		type APIAutoSQLGetByAPIIDEntity struct{
			
				
					APIID string 
				
				
			
		}

		func (t *APIAutoSQLGetByAPIIDEntity) TplName() string{
			return "api.auto.sql.GetByAPIID"
		}

	

		type APIAutoSQLInsertEntity struct{
			
				
					APIID string 
				
				
			
				
					Description string 
				
				
			
				
					Name string 
				
				
			
				
					ServiceID string 
				
				
			
				
					Summary string 
				
				
			
				
					Tags string 
				
				
			
				
					Title string 
				
				
			
				
					URI string 
				
				
			
		}

		func (t *APIAutoSQLInsertEntity) TplName() string{
			return "api.auto.sql.Insert"
		}

	

		type APIAutoSQLPaginateEntity struct{
			
				
					Limit int 
				
				
			
				
					Offset int 
				
				
			
				
					*APIAutoSQLPaginateWhereEntity
				
				
			
		}

		func (t *APIAutoSQLPaginateEntity) TplName() string{
			return "api.auto.sql.Paginate"
		}

	

		type APIAutoSQLPaginateTotalEntity struct{
			
				
					*APIAutoSQLPaginateWhereEntity
				
				
			
		}

		func (t *APIAutoSQLPaginateTotalEntity) TplName() string{
			return "api.auto.sql.PaginateTotal"
		}

	

		type APIAutoSQLPaginateWhereEntity struct{
			
		}

		func (t *APIAutoSQLPaginateWhereEntity) TplName() string{
			return "api.auto.sql.PaginateWhere"
		}

	

		type APIAutoSQLUpdateEntity struct{
			
				
					APIID string 
				
				
			
				
					Description string 
				
				
			
				
					Name string 
				
				
			
				
					ServiceID string 
				
				
			
				
					Summary string 
				
				
			
				
					Tags string 
				
				
			
				
					Title string 
				
				
			
				
					URI string 
				
				
			
		}

		func (t *APIAutoSQLUpdateEntity) TplName() string{
			return "api.auto.sql.Update"
		}

	

		type ExampleAutoSQLDelEntity struct{
			
				
					ExampleID string 
				
				
			
				
					Operator interface{} 
				
				
			
				
					OperatorID interface{} 
				
				
			
		}

		func (t *ExampleAutoSQLDelEntity) TplName() string{
			return "example.auto.sql.Del"
		}

	

		type ExampleAutoSQLGetByExampleIDEntity struct{
			
				
					ExampleID string 
				
				
			
		}

		func (t *ExampleAutoSQLGetByExampleIDEntity) TplName() string{
			return "example.auto.sql.GetByExampleID"
		}

	

		type ExampleAutoSQLInsertEntity struct{
			
				
					APIID string 
				
				
			
				
					ExampleID string 
				
				
			
				
					Request string 
				
				
			
				
					Response string 
				
				
			
				
					ServiceID string 
				
				
			
				
					Summary string 
				
				
			
				
					Tag string 
				
				
			
				
					Title string 
				
				
			
		}

		func (t *ExampleAutoSQLInsertEntity) TplName() string{
			return "example.auto.sql.Insert"
		}

	

		type ExampleAutoSQLPaginateEntity struct{
			
				
					Limit int 
				
				
			
				
					Offset int 
				
				
			
				
					*ExampleAutoSQLPaginateWhereEntity
				
				
			
		}

		func (t *ExampleAutoSQLPaginateEntity) TplName() string{
			return "example.auto.sql.Paginate"
		}

	

		type ExampleAutoSQLPaginateTotalEntity struct{
			
				
					*ExampleAutoSQLPaginateWhereEntity
				
				
			
		}

		func (t *ExampleAutoSQLPaginateTotalEntity) TplName() string{
			return "example.auto.sql.PaginateTotal"
		}

	

		type ExampleAutoSQLPaginateWhereEntity struct{
			
		}

		func (t *ExampleAutoSQLPaginateWhereEntity) TplName() string{
			return "example.auto.sql.PaginateWhere"
		}

	

		type ExampleAutoSQLUpdateEntity struct{
			
				
					APIID string 
				
				
			
				
					ExampleID string 
				
				
			
				
					Request string 
				
				
			
				
					Response string 
				
				
			
				
					ServiceID string 
				
				
			
				
					Summary string 
				
				
			
				
					Tag string 
				
				
			
				
					Title string 
				
				
			
		}

		func (t *ExampleAutoSQLUpdateEntity) TplName() string{
			return "example.auto.sql.Update"
		}

	

		type MarkdownAutoSQLDelEntity struct{
			
				
					MarkdownID string 
				
				
			
				
					Operator interface{} 
				
				
			
				
					OperatorID interface{} 
				
				
			
		}

		func (t *MarkdownAutoSQLDelEntity) TplName() string{
			return "markdown.auto.sql.Del"
		}

	

		type MarkdownAutoSQLGetByMarkdownIDEntity struct{
			
				
					MarkdownID string 
				
				
			
		}

		func (t *MarkdownAutoSQLGetByMarkdownIDEntity) TplName() string{
			return "markdown.auto.sql.GetByMarkdownID"
		}

	

		type MarkdownAutoSQLInsertEntity struct{
			
				
					APIID string 
				
				
			
				
					Content string 
				
				
			
				
					Markdown string 
				
				
			
				
					MarkdownID string 
				
				
			
				
					Name string 
				
				
			
				
					OwnerID int 
				
				
			
				
					OwnerName string 
				
				
			
				
					ServiceID string 
				
				
			
				
					Title string 
				
				
			
		}

		func (t *MarkdownAutoSQLInsertEntity) TplName() string{
			return "markdown.auto.sql.Insert"
		}

	

		type MarkdownAutoSQLPaginateEntity struct{
			
				
					Limit int 
				
				
			
				
					Offset int 
				
				
			
				
					*MarkdownAutoSQLPaginateWhereEntity
				
				
			
		}

		func (t *MarkdownAutoSQLPaginateEntity) TplName() string{
			return "markdown.auto.sql.Paginate"
		}

	

		type MarkdownAutoSQLPaginateTotalEntity struct{
			
				
					*MarkdownAutoSQLPaginateWhereEntity
				
				
			
		}

		func (t *MarkdownAutoSQLPaginateTotalEntity) TplName() string{
			return "markdown.auto.sql.PaginateTotal"
		}

	

		type MarkdownAutoSQLPaginateWhereEntity struct{
			
		}

		func (t *MarkdownAutoSQLPaginateWhereEntity) TplName() string{
			return "markdown.auto.sql.PaginateWhere"
		}

	

		type MarkdownAutoSQLUpdateEntity struct{
			
				
					APIID string 
				
				
			
				
					Content string 
				
				
			
				
					Markdown string 
				
				
			
				
					MarkdownID string 
				
				
			
				
					Name string 
				
				
			
				
					OwnerID int 
				
				
			
				
					OwnerName string 
				
				
			
				
					ServiceID string 
				
				
			
				
					Title string 
				
				
			
		}

		func (t *MarkdownAutoSQLUpdateEntity) TplName() string{
			return "markdown.auto.sql.Update"
		}

	

		type ParameterAutoSQLDelEntity struct{
			
				
					Operator interface{} 
				
				
			
				
					OperatorID interface{} 
				
				
			
				
					ParameterID string 
				
				
			
		}

		func (t *ParameterAutoSQLDelEntity) TplName() string{
			return "parameter.auto.sql.Del"
		}

	

		type ParameterAutoSQLGetByParameterIDEntity struct{
			
				
					ParameterID string 
				
				
			
		}

		func (t *ParameterAutoSQLGetByParameterIDEntity) TplName() string{
			return "parameter.auto.sql.GetByParameterID"
		}

	

		type ParameterAutoSQLInsertEntity struct{
			
				
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

		func (t *ParameterAutoSQLInsertEntity) TplName() string{
			return "parameter.auto.sql.Insert"
		}

	

		type ParameterAutoSQLPaginateEntity struct{
			
				
					Limit int 
				
				
			
				
					Offset int 
				
				
			
				
					*ParameterAutoSQLPaginateWhereEntity
				
				
			
		}

		func (t *ParameterAutoSQLPaginateEntity) TplName() string{
			return "parameter.auto.sql.Paginate"
		}

	

		type ParameterAutoSQLPaginateTotalEntity struct{
			
				
					*ParameterAutoSQLPaginateWhereEntity
				
				
			
		}

		func (t *ParameterAutoSQLPaginateTotalEntity) TplName() string{
			return "parameter.auto.sql.PaginateTotal"
		}

	

		type ParameterAutoSQLPaginateWhereEntity struct{
			
		}

		func (t *ParameterAutoSQLPaginateWhereEntity) TplName() string{
			return "parameter.auto.sql.PaginateWhere"
		}

	

		type ParameterAutoSQLUpdateEntity struct{
			
				
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

		func (t *ParameterAutoSQLUpdateEntity) TplName() string{
			return "parameter.auto.sql.Update"
		}

	

		type ServerAutoSQLDelEntity struct{
			
				
					Operator interface{} 
				
				
			
				
					OperatorID interface{} 
				
				
			
				
					ServerID string 
				
				
			
		}

		func (t *ServerAutoSQLDelEntity) TplName() string{
			return "server.auto.sql.Del"
		}

	

		type ServerAutoSQLGetByServerIDEntity struct{
			
				
					ServerID string 
				
				
			
		}

		func (t *ServerAutoSQLGetByServerIDEntity) TplName() string{
			return "server.auto.sql.GetByServerID"
		}

	

		type ServerAutoSQLInsertEntity struct{
			
				
					Description string 
				
				
			
				
					ExtensionIds string 
				
				
			
				
					Proxy string 
				
				
			
				
					ServerID string 
				
				
			
				
					ServiceID string 
				
				
			
				
					URL string 
				
				
			
		}

		func (t *ServerAutoSQLInsertEntity) TplName() string{
			return "server.auto.sql.Insert"
		}

	

		type ServerAutoSQLPaginateEntity struct{
			
				
					Limit int 
				
				
			
				
					Offset int 
				
				
			
				
					*ServerAutoSQLPaginateWhereEntity
				
				
			
		}

		func (t *ServerAutoSQLPaginateEntity) TplName() string{
			return "server.auto.sql.Paginate"
		}

	

		type ServerAutoSQLPaginateTotalEntity struct{
			
				
					*ServerAutoSQLPaginateWhereEntity
				
				
			
		}

		func (t *ServerAutoSQLPaginateTotalEntity) TplName() string{
			return "server.auto.sql.PaginateTotal"
		}

	

		type ServerAutoSQLPaginateWhereEntity struct{
			
		}

		func (t *ServerAutoSQLPaginateWhereEntity) TplName() string{
			return "server.auto.sql.PaginateWhere"
		}

	

		type ServerAutoSQLUpdateEntity struct{
			
				
					Description string 
				
				
			
				
					ExtensionIds string 
				
				
			
				
					Proxy string 
				
				
			
				
					ServerID string 
				
				
			
				
					ServiceID string 
				
				
			
				
					URL string 
				
				
			
		}

		func (t *ServerAutoSQLUpdateEntity) TplName() string{
			return "server.auto.sql.Update"
		}

	

		type ServiceAutoSQLDelEntity struct{
			
				
					Operator interface{} 
				
				
			
				
					OperatorID interface{} 
				
				
			
				
					ServiceID string 
				
				
			
		}

		func (t *ServiceAutoSQLDelEntity) TplName() string{
			return "service.auto.sql.Del"
		}

	

		type ServiceAutoSQLGetByServiceIDEntity struct{
			
				
					ServiceID string 
				
				
			
		}

		func (t *ServiceAutoSQLGetByServiceIDEntity) TplName() string{
			return "service.auto.sql.GetByServiceID"
		}

	

		type ServiceAutoSQLInsertEntity struct{
			
				
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

		func (t *ServiceAutoSQLInsertEntity) TplName() string{
			return "service.auto.sql.Insert"
		}

	

		type ServiceAutoSQLPaginateEntity struct{
			
				
					Limit int 
				
				
			
				
					Offset int 
				
				
			
				
					*ServiceAutoSQLPaginateWhereEntity
				
				
			
		}

		func (t *ServiceAutoSQLPaginateEntity) TplName() string{
			return "service.auto.sql.Paginate"
		}

	

		type ServiceAutoSQLPaginateTotalEntity struct{
			
				
					*ServiceAutoSQLPaginateWhereEntity
				
				
			
		}

		func (t *ServiceAutoSQLPaginateTotalEntity) TplName() string{
			return "service.auto.sql.PaginateTotal"
		}

	

		type ServiceAutoSQLPaginateWhereEntity struct{
			
		}

		func (t *ServiceAutoSQLPaginateWhereEntity) TplName() string{
			return "service.auto.sql.PaginateWhere"
		}

	

		type ServiceAutoSQLUpdateEntity struct{
			
				
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

		func (t *ServiceAutoSQLUpdateEntity) TplName() string{
			return "service.auto.sql.Update"
		}

	

		type ValidateSchemaAutoSQLDelEntity struct{
			
				
					Operator interface{} 
				
				
			
				
					OperatorID interface{} 
				
				
			
				
					ValidateSchemaID string 
				
				
			
		}

		func (t *ValidateSchemaAutoSQLDelEntity) TplName() string{
			return "validate_schema.auto.sql.Del"
		}

	

		type ValidateSchemaAutoSQLGetByValidateSchemaIDEntity struct{
			
				
					ValidateSchemaID string 
				
				
			
		}

		func (t *ValidateSchemaAutoSQLGetByValidateSchemaIDEntity) TplName() string{
			return "validate_schema.auto.sql.GetByValidateSchemaID"
		}

	

		type ValidateSchemaAutoSQLInsertEntity struct{
			
				
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
				
				
			
		}

		func (t *ValidateSchemaAutoSQLInsertEntity) TplName() string{
			return "validate_schema.auto.sql.Insert"
		}

	

		type ValidateSchemaAutoSQLPaginateEntity struct{
			
				
					Limit int 
				
				
			
				
					Offset int 
				
				
			
				
					*ValidateSchemaAutoSQLPaginateWhereEntity
				
				
			
		}

		func (t *ValidateSchemaAutoSQLPaginateEntity) TplName() string{
			return "validate_schema.auto.sql.Paginate"
		}

	

		type ValidateSchemaAutoSQLPaginateTotalEntity struct{
			
				
					*ValidateSchemaAutoSQLPaginateWhereEntity
				
				
			
		}

		func (t *ValidateSchemaAutoSQLPaginateTotalEntity) TplName() string{
			return "validate_schema.auto.sql.PaginateTotal"
		}

	

		type ValidateSchemaAutoSQLPaginateWhereEntity struct{
			
		}

		func (t *ValidateSchemaAutoSQLPaginateWhereEntity) TplName() string{
			return "validate_schema.auto.sql.PaginateWhere"
		}

	

		type ValidateSchemaAutoSQLUpdateEntity struct{
			
				
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
				
				
			
		}

		func (t *ValidateSchemaAutoSQLUpdateEntity) TplName() string{
			return "validate_schema.auto.sql.Update"
		}

	