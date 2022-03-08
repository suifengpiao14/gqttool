package example

		type APIautoDelEntity struct{
			
				APIID string 
			
		}

		func (t *APIautoDelEntity) TplName() string{
			return "API.auto.Del"
		}

	

		type APIautoGetByAPIIDEntity struct{
			
				APIID string 
			
		}

		func (t *APIautoGetByAPIIDEntity) TplName() string{
			return "API.auto.GetByAPIID"
		}

	

		type APIautoInsertEntity struct{
			
				APIID string 
			
				CreatedAt string 
			
				Description string 
			
				Name string 
			
				ServiceID string 
			
				Summary string 
			
				Tags string 
			
				Title string 
			
				URI string 
			
				UpdatedAt string 
			
		}

		func (t *APIautoInsertEntity) TplName() string{
			return "API.auto.Insert"
		}

	

		type APIautoPaginateEntity struct{
			
				Limit int 
			
				Offset int 
			
		}

		func (t *APIautoPaginateEntity) TplName() string{
			return "API.auto.Paginate"
		}

	

		type APIautoTotalEntity struct{
			
		}

		func (t *APIautoTotalEntity) TplName() string{
			return "API.auto.Total"
		}

	

		type APIautoUpdateEntity struct{
			
				APIID string 
			
				CreatedAt string 
			
				Description string 
			
				Name string 
			
				ServiceID string 
			
				Summary string 
			
				Tags string 
			
				Title string 
			
				URI string 
			
				UpdatedAt string 
			
		}

		func (t *APIautoUpdateEntity) TplName() string{
			return "API.auto.Update"
		}

	

		type ExampleAutoDelEntity struct{
			
				ExampleID string 
			
		}

		func (t *ExampleAutoDelEntity) TplName() string{
			return "Example.auto.Del"
		}

	

		type ExampleAutoGetByExampleIDEntity struct{
			
				ExampleID string 
			
		}

		func (t *ExampleAutoGetByExampleIDEntity) TplName() string{
			return "Example.auto.GetByExampleID"
		}

	

		type ExampleAutoInsertEntity struct{
			
				APIID string 
			
				CreatedAt string 
			
				ExampleID string 
			
				Request string 
			
				Response string 
			
				ServiceID string 
			
				Summary string 
			
				Tag string 
			
				Title string 
			
				UpdatedAt string 
			
		}

		func (t *ExampleAutoInsertEntity) TplName() string{
			return "Example.auto.Insert"
		}

	

		type ExampleAutoPaginateEntity struct{
			
				Limit int 
			
				Offset int 
			
		}

		func (t *ExampleAutoPaginateEntity) TplName() string{
			return "Example.auto.Paginate"
		}

	

		type ExampleAutoTotalEntity struct{
			
		}

		func (t *ExampleAutoTotalEntity) TplName() string{
			return "Example.auto.Total"
		}

	

		type ExampleAutoUpdateEntity struct{
			
				APIID string 
			
				CreatedAt string 
			
				ExampleID string 
			
				Request string 
			
				Response string 
			
				ServiceID string 
			
				Summary string 
			
				Tag string 
			
				Title string 
			
				UpdatedAt string 
			
		}

		func (t *ExampleAutoUpdateEntity) TplName() string{
			return "Example.auto.Update"
		}

	

		type MarkdownAutoDelEntity struct{
			
				MarkdownID string 
			
		}

		func (t *MarkdownAutoDelEntity) TplName() string{
			return "Markdown.auto.Del"
		}

	

		type MarkdownAutoGetByMarkdownIDEntity struct{
			
				MarkdownID string 
			
		}

		func (t *MarkdownAutoGetByMarkdownIDEntity) TplName() string{
			return "Markdown.auto.GetByMarkdownID"
		}

	

		type MarkdownAutoInsertEntity struct{
			
				APIID string 
			
				Content string 
			
				CreatedAt string 
			
				Markdown string 
			
				MarkdownID string 
			
				Name string 
			
				OwnerID int 
			
				OwnerName string 
			
				ServiceID string 
			
				Title string 
			
				UpdatedAt string 
			
		}

		func (t *MarkdownAutoInsertEntity) TplName() string{
			return "Markdown.auto.Insert"
		}

	

		type MarkdownAutoPaginateEntity struct{
			
				Limit int 
			
				Offset int 
			
		}

		func (t *MarkdownAutoPaginateEntity) TplName() string{
			return "Markdown.auto.Paginate"
		}

	

		type MarkdownAutoTotalEntity struct{
			
		}

		func (t *MarkdownAutoTotalEntity) TplName() string{
			return "Markdown.auto.Total"
		}

	

		type MarkdownAutoUpdateEntity struct{
			
				APIID string 
			
				Content string 
			
				CreatedAt string 
			
				Markdown string 
			
				MarkdownID string 
			
				Name string 
			
				OwnerID int 
			
				OwnerName string 
			
				ServiceID string 
			
				Title string 
			
				UpdatedAt string 
			
		}

		func (t *MarkdownAutoUpdateEntity) TplName() string{
			return "Markdown.auto.Update"
		}

	

		type ParameterAutoDelEntity struct{
			
				ParameterID string 
			
		}

		func (t *ParameterAutoDelEntity) TplName() string{
			return "Parameter.auto.Del"
		}

	

		type ParameterAutoGetByParameterIDEntity struct{
			
				ParameterID string 
			
		}

		func (t *ParameterAutoGetByParameterIDEntity) TplName() string{
			return "Parameter.auto.GetByParameterID"
		}

	

		type ParameterAutoInsertEntity struct{
			
				APIID string 
			
				AllowEmptyValue string 
			
				AllowReserved string 
			
				CreatedAt string 
			
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
			
				UpdatedAt string 
			
				ValidateSchemaID string 
			
		}

		func (t *ParameterAutoInsertEntity) TplName() string{
			return "Parameter.auto.Insert"
		}

	

		type ParameterAutoPaginateEntity struct{
			
				Limit int 
			
				Offset int 
			
		}

		func (t *ParameterAutoPaginateEntity) TplName() string{
			return "Parameter.auto.Paginate"
		}

	

		type ParameterAutoTotalEntity struct{
			
		}

		func (t *ParameterAutoTotalEntity) TplName() string{
			return "Parameter.auto.Total"
		}

	

		type ParameterAutoUpdateEntity struct{
			
				APIID string 
			
				AllowEmptyValue string 
			
				AllowReserved string 
			
				CreatedAt string 
			
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
			
				UpdatedAt string 
			
				ValidateSchemaID string 
			
		}

		func (t *ParameterAutoUpdateEntity) TplName() string{
			return "Parameter.auto.Update"
		}

	

		type ServerAutoDelEntity struct{
			
				ServerID string 
			
		}

		func (t *ServerAutoDelEntity) TplName() string{
			return "Server.auto.Del"
		}

	

		type ServerAutoGetByServerIDEntity struct{
			
				ServerID string 
			
		}

		func (t *ServerAutoGetByServerIDEntity) TplName() string{
			return "Server.auto.GetByServerID"
		}

	

		type ServerAutoInsertEntity struct{
			
				CreatedAt string 
			
				Description string 
			
				ExtensionIds string 
			
				Proxy string 
			
				ServerID string 
			
				ServiceID string 
			
				URL string 
			
				UpdatedAt string 
			
		}

		func (t *ServerAutoInsertEntity) TplName() string{
			return "Server.auto.Insert"
		}

	

		type ServerAutoPaginateEntity struct{
			
				Limit int 
			
				Offset int 
			
		}

		func (t *ServerAutoPaginateEntity) TplName() string{
			return "Server.auto.Paginate"
		}

	

		type ServerAutoTotalEntity struct{
			
		}

		func (t *ServerAutoTotalEntity) TplName() string{
			return "Server.auto.Total"
		}

	

		type ServerAutoUpdateEntity struct{
			
				CreatedAt string 
			
				Description string 
			
				ExtensionIds string 
			
				Proxy string 
			
				ServerID string 
			
				ServiceID string 
			
				URL string 
			
				UpdatedAt string 
			
		}

		func (t *ServerAutoUpdateEntity) TplName() string{
			return "Server.auto.Update"
		}

	

		type ServiceAutoDelEntity struct{
			
				ServiceID string 
			
		}

		func (t *ServiceAutoDelEntity) TplName() string{
			return "Service.auto.Del"
		}

	

		type ServiceAutoGetByServiceIDEntity struct{
			
				ServiceID string 
			
		}

		func (t *ServiceAutoGetByServiceIDEntity) TplName() string{
			return "Service.auto.GetByServiceID"
		}

	

		type ServiceAutoInsertEntity struct{
			
				ContactIds string 
			
				CreatedAt string 
			
				Description string 
			
				License string 
			
				Proxy string 
			
				Security string 
			
				ServiceID string 
			
				Title string 
			
				UpdatedAt string 
			
				Variables string 
			
				Version string 
			
		}

		func (t *ServiceAutoInsertEntity) TplName() string{
			return "Service.auto.Insert"
		}

	

		type ServiceAutoPaginateEntity struct{
			
				Limit int 
			
				Offset int 
			
		}

		func (t *ServiceAutoPaginateEntity) TplName() string{
			return "Service.auto.Paginate"
		}

	

		type ServiceAutoTotalEntity struct{
			
		}

		func (t *ServiceAutoTotalEntity) TplName() string{
			return "Service.auto.Total"
		}

	

		type ServiceAutoUpdateEntity struct{
			
				ContactIds string 
			
				CreatedAt string 
			
				Description string 
			
				License string 
			
				Proxy string 
			
				Security string 
			
				ServiceID string 
			
				Title string 
			
				UpdatedAt string 
			
				Variables string 
			
				Version string 
			
		}

		func (t *ServiceAutoUpdateEntity) TplName() string{
			return "Service.auto.Update"
		}

	

		type ValidateSchemaAutoDelEntity struct{
			
				ValidateSchemaID string 
			
		}

		func (t *ValidateSchemaAutoDelEntity) TplName() string{
			return "ValidateSchema.auto.Del"
		}

	

		type ValidateSchemaAutoGetByValidateSchemaIDEntity struct{
			
				ValidateSchemaID string 
			
		}

		func (t *ValidateSchemaAutoGetByValidateSchemaIDEntity) TplName() string{
			return "ValidateSchema.auto.GetByValidateSchemaID"
		}

	

		type ValidateSchemaAutoInsertEntity struct{
			
				AdditionalProperties string 
			
				AllOf string 
			
				AllowEmptyValue string 
			
				AllowReserved string 
			
				AnyOf string 
			
				CreatedAt string 
			
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
			
				UpdatedAt string 
			
				ValidateSchemaID string 
			
				WriteOnly string 
			
				XML string 
			
		}

		func (t *ValidateSchemaAutoInsertEntity) TplName() string{
			return "ValidateSchema.auto.Insert"
		}

	

		type ValidateSchemaAutoPaginateEntity struct{
			
				Limit int 
			
				Offset int 
			
		}

		func (t *ValidateSchemaAutoPaginateEntity) TplName() string{
			return "ValidateSchema.auto.Paginate"
		}

	

		type ValidateSchemaAutoTotalEntity struct{
			
		}

		func (t *ValidateSchemaAutoTotalEntity) TplName() string{
			return "ValidateSchema.auto.Total"
		}

	

		type ValidateSchemaAutoUpdateEntity struct{
			
				AdditionalProperties string 
			
				AllOf string 
			
				AllowEmptyValue string 
			
				AllowReserved string 
			
				AnyOf string 
			
				CreatedAt string 
			
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
			
				UpdatedAt string 
			
				ValidateSchemaID string 
			
				WriteOnly string 
			
				XML string 
			
		}

		func (t *ValidateSchemaAutoUpdateEntity) TplName() string{
			return "ValidateSchema.auto.Update"
		}

	