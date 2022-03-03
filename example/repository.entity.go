package example

		type APIDel struct{
			
				APIID string 
			
		}
	

		type APIGetByAPIID struct{
			
				APIID string 
			
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
	

		type APIPaginate struct{
			
				Limit int 
			
				Offset int 
			
		}
	

		type APITotal struct{
			
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
	

		type ExampleDel struct{
			
				ExampleID string 
			
		}
	

		type ExampleGetByExampleID struct{
			
				ExampleID string 
			
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
	

		type ExamplePaginate struct{
			
				Limit int 
			
				Offset int 
			
		}
	

		type ExampleTotal struct{
			
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
	

		type MarkdownDel struct{
			
				MarkdownID string 
			
		}
	

		type MarkdownGetByMarkdownID struct{
			
				MarkdownID string 
			
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
	

		type MarkdownPaginate struct{
			
				Limit int 
			
				Offset int 
			
		}
	

		type MarkdownTotal struct{
			
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
	

		type ParameterDel struct{
			
				ParameterID string 
			
		}
	

		type ParameterGetByParameterID struct{
			
				ParameterID string 
			
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
	

		type ParameterPaginate struct{
			
				Limit int 
			
				Offset int 
			
		}
	

		type ParameterTotal struct{
			
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
	

		type ServerDel struct{
			
				ServerID string 
			
		}
	

		type ServerGetByServerID struct{
			
				ServerID string 
			
		}
	

		type ServerInsert struct{
			
				Description string 
			
				ExtensionIds string 
			
				Proxy string 
			
				ServerID string 
			
				ServiceID string 
			
				URL string 
			
		}
	

		type ServerPaginate struct{
			
				Limit int 
			
				Offset int 
			
		}
	

		type ServerTotal struct{
			
		}
	

		type ServerUpdate struct{
			
				Description string 
			
				ExtensionIds string 
			
				Proxy string 
			
				ServerID string 
			
				ServiceID string 
			
				URL string 
			
		}
	

		type ServiceDel struct{
			
				ServiceID string 
			
		}
	

		type ServiceGetByServiceID struct{
			
				ServiceID string 
			
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
	

		type ServiceList struct{
			
				Limit int 
			
				Offset int 
			
				Title string 
			
		}
	

		type ServiceTotal struct{
			
				Title string 
			
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
	

		type ValidateSchemaDel struct{
			
				ValidateSchemaID string 
			
		}
	

		type ValidateSchemaGetByValidateSchemaID struct{
			
				ValidateSchemaID string 
			
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
	

		type ValidateSchemaPaginate struct{
			
				Limit int 
			
				Offset int 
			
		}
	

		type ValidateSchemaTotal struct{
			
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
	