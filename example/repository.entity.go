package example

		type APIDel struct{
			
				APIID string 
			
		}

		func (s *APIDel) DefineName() string{
			return "Del"
		}

	

		type APIGetByAPIID struct{
			
				APIID string 
			
		}

		func (s *APIGetByAPIID) DefineName() string{
			return "GetByAPIID"
		}

	

		type APIInsert struct{
			
				APIID string 
			
				Description string 
			
				Name string 
			
				ServiceID string 
			
				Summary string 
			
				Tags string 
			
				Title string 
			
				URI string 
			
		}

		func (s *APIInsert) DefineName() string{
			return "Insert"
		}

	

		type APIPaginate struct{
			
				Limit int 
			
				Offset int 
			
		}

		func (s *APIPaginate) DefineName() string{
			return "Paginate"
		}

	

		type APITotal struct{
			
		}

		func (s *APITotal) DefineName() string{
			return "Total"
		}

	

		type APIUpdate struct{
			
				APIID string 
			
				Description string 
			
				Name string 
			
				ServiceID string 
			
				Summary string 
			
				Tags string 
			
				Title string 
			
				URI string 
			
		}

		func (s *APIUpdate) DefineName() string{
			return "Update"
		}

	

		type ExampleDel struct{
			
				ExampleID string 
			
		}

		func (s *ExampleDel) DefineName() string{
			return "Del"
		}

	

		type ExampleGetByExampleID struct{
			
				ExampleID string 
			
		}

		func (s *ExampleGetByExampleID) DefineName() string{
			return "GetByExampleID"
		}

	

		type ExampleInsert struct{
			
				APIID string 
			
				ExampleID string 
			
				Request string 
			
				Response string 
			
				ServiceID string 
			
				Summary string 
			
				Tag string 
			
				Title string 
			
		}

		func (s *ExampleInsert) DefineName() string{
			return "Insert"
		}

	

		type ExamplePaginate struct{
			
				Limit int 
			
				Offset int 
			
		}

		func (s *ExamplePaginate) DefineName() string{
			return "Paginate"
		}

	

		type ExampleTotal struct{
			
		}

		func (s *ExampleTotal) DefineName() string{
			return "Total"
		}

	

		type ExampleUpdate struct{
			
				APIID string 
			
				ExampleID string 
			
				Request string 
			
				Response string 
			
				ServiceID string 
			
				Summary string 
			
				Tag string 
			
				Title string 
			
		}

		func (s *ExampleUpdate) DefineName() string{
			return "Update"
		}

	

		type MarkdownDel struct{
			
				MarkdownID string 
			
		}

		func (s *MarkdownDel) DefineName() string{
			return "Del"
		}

	

		type MarkdownGetByMarkdownID struct{
			
				MarkdownID string 
			
		}

		func (s *MarkdownGetByMarkdownID) DefineName() string{
			return "GetByMarkdownID"
		}

	

		type MarkdownInsert struct{
			
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

		func (s *MarkdownInsert) DefineName() string{
			return "Insert"
		}

	

		type MarkdownPaginate struct{
			
				Limit int 
			
				Offset int 
			
		}

		func (s *MarkdownPaginate) DefineName() string{
			return "Paginate"
		}

	

		type MarkdownTotal struct{
			
		}

		func (s *MarkdownTotal) DefineName() string{
			return "Total"
		}

	

		type MarkdownUpdate struct{
			
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

		func (s *MarkdownUpdate) DefineName() string{
			return "Update"
		}

	

		type ParameterDel struct{
			
				ParameterID string 
			
		}

		func (s *ParameterDel) DefineName() string{
			return "Del"
		}

	

		type ParameterGetByParameterID struct{
			
				ParameterID string 
			
		}

		func (s *ParameterGetByParameterID) DefineName() string{
			return "GetByParameterID"
		}

	

		type ParameterInsert struct{
			
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

		func (s *ParameterInsert) DefineName() string{
			return "Insert"
		}

	

		type ParameterPaginate struct{
			
				Limit int 
			
				Offset int 
			
		}

		func (s *ParameterPaginate) DefineName() string{
			return "Paginate"
		}

	

		type ParameterTotal struct{
			
		}

		func (s *ParameterTotal) DefineName() string{
			return "Total"
		}

	

		type ParameterUpdate struct{
			
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

		func (s *ParameterUpdate) DefineName() string{
			return "Update"
		}

	

		type ServerDel struct{
			
				ServerID string 
			
		}

		func (s *ServerDel) DefineName() string{
			return "Del"
		}

	

		type ServerGetByServerID struct{
			
				ServerID string 
			
		}

		func (s *ServerGetByServerID) DefineName() string{
			return "GetByServerID"
		}

	

		type ServerInsert struct{
			
				Description string 
			
				ExtensionIds string 
			
				Proxy string 
			
				ServerID string 
			
				ServiceID string 
			
				URL string 
			
		}

		func (s *ServerInsert) DefineName() string{
			return "Insert"
		}

	

		type ServerPaginate struct{
			
				Limit int 
			
				Offset int 
			
		}

		func (s *ServerPaginate) DefineName() string{
			return "Paginate"
		}

	

		type ServerTotal struct{
			
		}

		func (s *ServerTotal) DefineName() string{
			return "Total"
		}

	

		type ServerUpdate struct{
			
				Description string 
			
				ExtensionIds string 
			
				Proxy string 
			
				ServerID string 
			
				ServiceID string 
			
				URL string 
			
		}

		func (s *ServerUpdate) DefineName() string{
			return "Update"
		}

	

		type ServiceDel struct{
			
				ServiceID string 
			
		}

		func (s *ServiceDel) DefineName() string{
			return "del"
		}

	

		type ServiceGetByServiceID struct{
			
				ServiceID string 
			
		}

		func (s *ServiceGetByServiceID) DefineName() string{
			return "getByServiceId"
		}

	

		type ServiceInsert struct{
			
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

		func (s *ServiceInsert) DefineName() string{
			return "insert"
		}

	

		type ServiceList struct{
			
				Limit int 
			
				Offset int 
			
				Title string 
			
		}

		func (s *ServiceList) DefineName() string{
			return "list"
		}

	

		type ServiceTotal struct{
			
				Title string 
			
		}

		func (s *ServiceTotal) DefineName() string{
			return "total"
		}

	

		type ServiceUpdate struct{
			
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

		func (s *ServiceUpdate) DefineName() string{
			return "update"
		}

	

		type ValidateSchemaDel struct{
			
				ValidateSchemaID string 
			
		}

		func (s *ValidateSchemaDel) DefineName() string{
			return "Del"
		}

	

		type ValidateSchemaGetByValidateSchemaID struct{
			
				ValidateSchemaID string 
			
		}

		func (s *ValidateSchemaGetByValidateSchemaID) DefineName() string{
			return "GetByValidateSchemaID"
		}

	

		type ValidateSchemaInsert struct{
			
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

		func (s *ValidateSchemaInsert) DefineName() string{
			return "Insert"
		}

	

		type ValidateSchemaPaginate struct{
			
				Limit int 
			
				Offset int 
			
		}

		func (s *ValidateSchemaPaginate) DefineName() string{
			return "Paginate"
		}

	

		type ValidateSchemaTotal struct{
			
		}

		func (s *ValidateSchemaTotal) DefineName() string{
			return "Total"
		}

	

		type ValidateSchemaUpdate struct{
			
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

		func (s *ValidateSchemaUpdate) DefineName() string{
			return "Update"
		}

	