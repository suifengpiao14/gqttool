package example
import "github.com/suifengpiao14/gqt/v2/gqttpl"

		type SQLAPIGenSQLDelEntity struct{
			
				APIID string  
			
				OperatorIDInt int  
			
				OperatorStr string  
			gqttpl.TplEmptyEntity}

		func (t *SQLAPIGenSQLDelEntity) TplName() string{
			return "sql.api.gen.sql.Del"
		}
	

		type SQLAPIGenSQLGetAllByAPIIDListEntity struct{
			
				APIIDList []string  
			gqttpl.TplEmptyEntity}

		func (t *SQLAPIGenSQLGetAllByAPIIDListEntity) TplName() string{
			return "sql.api.gen.sql.GetAllByAPIIDList"
		}
	

		type SQLAPIGenSQLGetByAPIIDEntity struct{
			
				APIID string  
			gqttpl.TplEmptyEntity}

		func (t *SQLAPIGenSQLGetByAPIIDEntity) TplName() string{
			return "sql.api.gen.sql.GetByAPIID"
		}
	

		type SQLAPIGenSQLInsertEntity struct{
			
				APIID string  
			
				Description string  
			
				Name string  
			
				ServiceID string  
			
				Summary string  
			
				Tags string  
			
				Title string  
			
				URI string  
			gqttpl.TplEmptyEntity}

		func (t *SQLAPIGenSQLInsertEntity) TplName() string{
			return "sql.api.gen.sql.Insert"
		}
	

		type SQLAPIGenSQLPaginateEntity struct{
			
				Limit int  
			
				Offset int  
			
				 SQLAPIGenSQLPaginateWhereEntity  
			gqttpl.TplEmptyEntity}

		func (t *SQLAPIGenSQLPaginateEntity) TplName() string{
			return "sql.api.gen.sql.Paginate"
		}
	

		type SQLAPIGenSQLPaginateTotalEntity struct{
			
				 SQLAPIGenSQLPaginateWhereEntity  
			gqttpl.TplEmptyEntity}

		func (t *SQLAPIGenSQLPaginateTotalEntity) TplName() string{
			return "sql.api.gen.sql.PaginateTotal"
		}
	

		type SQLAPIGenSQLPaginateWhereEntity struct{
			
				ID interface{}  
			}

		func (t *SQLAPIGenSQLPaginateWhereEntity) TplName() string{
			return "sql.api.gen.sql.PaginateWhere"
		}
	

		type SQLAPIGenSQLUpdateEntity struct{
			
				APIID string  
			
				Description string  
			
				Name string  
			
				ServiceID string  
			
				Summary string  
			
				Tags string  
			
				Title string  
			
				URI string  
			gqttpl.TplEmptyEntity}

		func (t *SQLAPIGenSQLUpdateEntity) TplName() string{
			return "sql.api.gen.sql.Update"
		}
	

		type SQLExampleGenSQLGetAllByExampleIDListEntity struct{
			
				ExampleIDList []string  
			gqttpl.TplEmptyEntity}

		func (t *SQLExampleGenSQLGetAllByExampleIDListEntity) TplName() string{
			return "sql.example.gen.sql.GetAllByExampleIDList"
		}
	

		type SQLExampleGenSQLGetByExampleIDEntity struct{
			
				ExampleID string  
			gqttpl.TplEmptyEntity}

		func (t *SQLExampleGenSQLGetByExampleIDEntity) TplName() string{
			return "sql.example.gen.sql.GetByExampleID"
		}
	

		type SQLExampleGenSQLInsertEntity struct{
			
				APIID string  
			
				ExampleID string  
			
				Request string  
			
				Response string  
			
				ServiceID string  
			
				Summary string  
			
				Tag string  
			
				Title string  
			gqttpl.TplEmptyEntity}

		func (t *SQLExampleGenSQLInsertEntity) TplName() string{
			return "sql.example.gen.sql.Insert"
		}
	

		type SQLExampleGenSQLPaginateEntity struct{
			
				Limit int  
			
				Offset int  
			
				 SQLExampleGenSQLPaginateWhereEntity  
			gqttpl.TplEmptyEntity}

		func (t *SQLExampleGenSQLPaginateEntity) TplName() string{
			return "sql.example.gen.sql.Paginate"
		}
	

		type SQLExampleGenSQLPaginateTotalEntity struct{
			
				 SQLExampleGenSQLPaginateWhereEntity  
			gqttpl.TplEmptyEntity}

		func (t *SQLExampleGenSQLPaginateTotalEntity) TplName() string{
			return "sql.example.gen.sql.PaginateTotal"
		}
	

		type SQLExampleGenSQLPaginateWhereEntity struct{
			}

		func (t *SQLExampleGenSQLPaginateWhereEntity) TplName() string{
			return "sql.example.gen.sql.PaginateWhere"
		}
	

		type SQLExampleGenSQLUpdateEntity struct{
			
				APIID string  
			
				ExampleID string  
			
				Request string  
			
				Response string  
			
				ServiceID string  
			
				Summary string  
			
				Tag string  
			
				Title string  
			gqttpl.TplEmptyEntity}

		func (t *SQLExampleGenSQLUpdateEntity) TplName() string{
			return "sql.example.gen.sql.Update"
		}
	

		type SQLMarkdownGenSQLGetAllByMarkdownIDListEntity struct{
			
				MarkdownIDList []string  
			gqttpl.TplEmptyEntity}

		func (t *SQLMarkdownGenSQLGetAllByMarkdownIDListEntity) TplName() string{
			return "sql.markdown.gen.sql.GetAllByMarkdownIDList"
		}
	

		type SQLMarkdownGenSQLGetByMarkdownIDEntity struct{
			
				MarkdownID string  
			gqttpl.TplEmptyEntity}

		func (t *SQLMarkdownGenSQLGetByMarkdownIDEntity) TplName() string{
			return "sql.markdown.gen.sql.GetByMarkdownID"
		}
	

		type SQLMarkdownGenSQLInsertEntity struct{
			
				APIID string  
			
				Content string  
			
				Markdown string  
			
				MarkdownID string  
			
				Name string  
			
				OwnerID int  
			
				OwnerName string  
			
				ServiceID string  
			
				Title string  
			gqttpl.TplEmptyEntity}

		func (t *SQLMarkdownGenSQLInsertEntity) TplName() string{
			return "sql.markdown.gen.sql.Insert"
		}
	

		type SQLMarkdownGenSQLPaginateEntity struct{
			
				Limit int  
			
				Offset int  
			
				 SQLMarkdownGenSQLPaginateWhereEntity  
			gqttpl.TplEmptyEntity}

		func (t *SQLMarkdownGenSQLPaginateEntity) TplName() string{
			return "sql.markdown.gen.sql.Paginate"
		}
	

		type SQLMarkdownGenSQLPaginateTotalEntity struct{
			
				 SQLMarkdownGenSQLPaginateWhereEntity  
			gqttpl.TplEmptyEntity}

		func (t *SQLMarkdownGenSQLPaginateTotalEntity) TplName() string{
			return "sql.markdown.gen.sql.PaginateTotal"
		}
	

		type SQLMarkdownGenSQLPaginateWhereEntity struct{
			}

		func (t *SQLMarkdownGenSQLPaginateWhereEntity) TplName() string{
			return "sql.markdown.gen.sql.PaginateWhere"
		}
	

		type SQLMarkdownGenSQLUpdateEntity struct{
			
				APIID string  
			
				Content string  
			
				Markdown string  
			
				MarkdownID string  
			
				Name string  
			
				OwnerID int  
			
				OwnerName string  
			
				ServiceID string  
			
				Title string  
			gqttpl.TplEmptyEntity}

		func (t *SQLMarkdownGenSQLUpdateEntity) TplName() string{
			return "sql.markdown.gen.sql.Update"
		}
	

		type SQLParameterGenSQLGetAllByParameterIDListEntity struct{
			
				ParameterIDList []string  
			gqttpl.TplEmptyEntity}

		func (t *SQLParameterGenSQLGetAllByParameterIDListEntity) TplName() string{
			return "sql.parameter.gen.sql.GetAllByParameterIDList"
		}
	

		type SQLParameterGenSQLGetByParameterIDEntity struct{
			
				ParameterID string  
			gqttpl.TplEmptyEntity}

		func (t *SQLParameterGenSQLGetByParameterIDEntity) TplName() string{
			return "sql.parameter.gen.sql.GetByParameterID"
		}
	

		type SQLParameterGenSQLInsertEntity struct{
			
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
			gqttpl.TplEmptyEntity}

		func (t *SQLParameterGenSQLInsertEntity) TplName() string{
			return "sql.parameter.gen.sql.Insert"
		}
	

		type SQLParameterGenSQLPaginateEntity struct{
			
				Limit int  
			
				Offset int  
			
				 SQLParameterGenSQLPaginateWhereEntity  
			gqttpl.TplEmptyEntity}

		func (t *SQLParameterGenSQLPaginateEntity) TplName() string{
			return "sql.parameter.gen.sql.Paginate"
		}
	

		type SQLParameterGenSQLPaginateTotalEntity struct{
			
				 SQLParameterGenSQLPaginateWhereEntity  
			gqttpl.TplEmptyEntity}

		func (t *SQLParameterGenSQLPaginateTotalEntity) TplName() string{
			return "sql.parameter.gen.sql.PaginateTotal"
		}
	

		type SQLParameterGenSQLPaginateWhereEntity struct{
			}

		func (t *SQLParameterGenSQLPaginateWhereEntity) TplName() string{
			return "sql.parameter.gen.sql.PaginateWhere"
		}
	

		type SQLParameterGenSQLUpdateEntity struct{
			
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
			gqttpl.TplEmptyEntity}

		func (t *SQLParameterGenSQLUpdateEntity) TplName() string{
			return "sql.parameter.gen.sql.Update"
		}
	

		type SQLServerGenSQLGetAllByServerIDListEntity struct{
			
				ServerIDList []string  
			gqttpl.TplEmptyEntity}

		func (t *SQLServerGenSQLGetAllByServerIDListEntity) TplName() string{
			return "sql.server.gen.sql.GetAllByServerIDList"
		}
	

		type SQLServerGenSQLGetByServerIDEntity struct{
			
				ServerID string  
			gqttpl.TplEmptyEntity}

		func (t *SQLServerGenSQLGetByServerIDEntity) TplName() string{
			return "sql.server.gen.sql.GetByServerID"
		}
	

		type SQLServerGenSQLInsertEntity struct{
			
				Description string  
			
				ExtensionIds string  
			
				Proxy string  
			
				ServerID string  
			
				ServiceID string  
			
				URL string  
			gqttpl.TplEmptyEntity}

		func (t *SQLServerGenSQLInsertEntity) TplName() string{
			return "sql.server.gen.sql.Insert"
		}
	

		type SQLServerGenSQLPaginateEntity struct{
			
				Limit int  
			
				Offset int  
			
				 SQLServerGenSQLPaginateWhereEntity  
			gqttpl.TplEmptyEntity}

		func (t *SQLServerGenSQLPaginateEntity) TplName() string{
			return "sql.server.gen.sql.Paginate"
		}
	

		type SQLServerGenSQLPaginateTotalEntity struct{
			
				 SQLServerGenSQLPaginateWhereEntity  
			gqttpl.TplEmptyEntity}

		func (t *SQLServerGenSQLPaginateTotalEntity) TplName() string{
			return "sql.server.gen.sql.PaginateTotal"
		}
	

		type SQLServerGenSQLPaginateWhereEntity struct{
			}

		func (t *SQLServerGenSQLPaginateWhereEntity) TplName() string{
			return "sql.server.gen.sql.PaginateWhere"
		}
	

		type SQLServerGenSQLUpdateEntity struct{
			
				Description string  
			
				ExtensionIds string  
			
				Proxy string  
			
				ServerID string  
			
				ServiceID string  
			
				URL string  
			gqttpl.TplEmptyEntity}

		func (t *SQLServerGenSQLUpdateEntity) TplName() string{
			return "sql.server.gen.sql.Update"
		}
	

		type SQLServiceGenSQLGetAllByServiceIDListEntity struct{
			
				ServiceIDList []string  
			gqttpl.TplEmptyEntity}

		func (t *SQLServiceGenSQLGetAllByServiceIDListEntity) TplName() string{
			return "sql.service.gen.sql.GetAllByServiceIDList"
		}
	

		type SQLServiceGenSQLGetByServiceIDEntity struct{
			
				ServiceID string  
			gqttpl.TplEmptyEntity}

		func (t *SQLServiceGenSQLGetByServiceIDEntity) TplName() string{
			return "sql.service.gen.sql.GetByServiceID"
		}
	

		type SQLServiceGenSQLInsertEntity struct{
			
				ContactIds string  
			
				Description string  
			
				License string  
			
				Proxy string  
			
				Security string  
			
				ServiceID string  
			
				Title string  
			
				Variables string  
			
				Version string  
			gqttpl.TplEmptyEntity}

		func (t *SQLServiceGenSQLInsertEntity) TplName() string{
			return "sql.service.gen.sql.Insert"
		}
	

		type SQLServiceGenSQLPaginateEntity struct{
			
				Limit int  
			
				Offset int  
			
				 SQLServiceGenSQLPaginateWhereEntity  
			gqttpl.TplEmptyEntity}

		func (t *SQLServiceGenSQLPaginateEntity) TplName() string{
			return "sql.service.gen.sql.Paginate"
		}
	

		type SQLServiceGenSQLPaginateTotalEntity struct{
			
				 SQLServiceGenSQLPaginateWhereEntity  
			gqttpl.TplEmptyEntity}

		func (t *SQLServiceGenSQLPaginateTotalEntity) TplName() string{
			return "sql.service.gen.sql.PaginateTotal"
		}
	

		type SQLServiceGenSQLPaginateWhereEntity struct{
			}

		func (t *SQLServiceGenSQLPaginateWhereEntity) TplName() string{
			return "sql.service.gen.sql.PaginateWhere"
		}
	

		type SQLServiceGenSQLUpdateEntity struct{
			
				ContactIds string  
			
				Description string  
			
				License string  
			
				Proxy string  
			
				Security string  
			
				ServiceID string  
			
				Title string  
			
				Variables string  
			
				Version string  
			gqttpl.TplEmptyEntity}

		func (t *SQLServiceGenSQLUpdateEntity) TplName() string{
			return "sql.service.gen.sql.Update"
		}
	

		type SQLValidateSchemaGenSQLGetAllByValidateSchemaIDListEntity struct{
			
				ValidateSchemaIDList []string  
			gqttpl.TplEmptyEntity}

		func (t *SQLValidateSchemaGenSQLGetAllByValidateSchemaIDListEntity) TplName() string{
			return "sql.validate_schema.gen.sql.GetAllByValidateSchemaIDList"
		}
	

		type SQLValidateSchemaGenSQLGetByValidateSchemaIDEntity struct{
			
				ValidateSchemaID string  
			gqttpl.TplEmptyEntity}

		func (t *SQLValidateSchemaGenSQLGetByValidateSchemaIDEntity) TplName() string{
			return "sql.validate_schema.gen.sql.GetByValidateSchemaID"
		}
	

		type SQLValidateSchemaGenSQLInsertEntity struct{
			
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
			gqttpl.TplEmptyEntity}

		func (t *SQLValidateSchemaGenSQLInsertEntity) TplName() string{
			return "sql.validate_schema.gen.sql.Insert"
		}
	

		type SQLValidateSchemaGenSQLPaginateEntity struct{
			
				Limit int  
			
				Offset int  
			
				 SQLValidateSchemaGenSQLPaginateWhereEntity  
			gqttpl.TplEmptyEntity}

		func (t *SQLValidateSchemaGenSQLPaginateEntity) TplName() string{
			return "sql.validate_schema.gen.sql.Paginate"
		}
	

		type SQLValidateSchemaGenSQLPaginateTotalEntity struct{
			
				 SQLValidateSchemaGenSQLPaginateWhereEntity  
			gqttpl.TplEmptyEntity}

		func (t *SQLValidateSchemaGenSQLPaginateTotalEntity) TplName() string{
			return "sql.validate_schema.gen.sql.PaginateTotal"
		}
	

		type SQLValidateSchemaGenSQLPaginateWhereEntity struct{
			}

		func (t *SQLValidateSchemaGenSQLPaginateWhereEntity) TplName() string{
			return "sql.validate_schema.gen.sql.PaginateWhere"
		}
	

		type SQLValidateSchemaGenSQLUpdateEntity struct{
			
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
			gqttpl.TplEmptyEntity}

		func (t *SQLValidateSchemaGenSQLUpdateEntity) TplName() string{
			return "sql.validate_schema.gen.sql.Update"
		}
	