package example
import "github.com/suifengpiao14/gqt/v2/gqttpl"

		type SQLAPIGenSQLDelEntity struct{
			
				
					APIID string 
				
				
			
				
					OperatorIDInt int 
				
				
			
				
					OperatorStr string 
				
				
			
			*gqttpl.DataVolumeMap
		}

		func (t *SQLAPIGenSQLDelEntity) TplName() string{
			return "sql.api.gen.sql.Del"
		}

		func (t *SQLAPIGenSQLDelEntity) TplOutput() (out string,err error){
			out,err=gqttpl.TplOutput(t,t.TplName())
			return 
		}

	

		type SQLAPIGenSQLGetAllByAPIIDListEntity struct{
			
				
					APIIDList []string 
				
				
			
			*gqttpl.DataVolumeMap
		}

		func (t *SQLAPIGenSQLGetAllByAPIIDListEntity) TplName() string{
			return "sql.api.gen.sql.GetAllByAPIIDList"
		}

		func (t *SQLAPIGenSQLGetAllByAPIIDListEntity) TplOutput() (out string,err error){
			out,err=gqttpl.TplOutput(t,t.TplName())
			return 
		}

	

		type SQLAPIGenSQLGetByAPIIDEntity struct{
			
				
					APIID string 
				
				
			
			*gqttpl.DataVolumeMap
		}

		func (t *SQLAPIGenSQLGetByAPIIDEntity) TplName() string{
			return "sql.api.gen.sql.GetByAPIID"
		}

		func (t *SQLAPIGenSQLGetByAPIIDEntity) TplOutput() (out string,err error){
			out,err=gqttpl.TplOutput(t,t.TplName())
			return 
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
				
				
			
			*gqttpl.DataVolumeMap
		}

		func (t *SQLAPIGenSQLInsertEntity) TplName() string{
			return "sql.api.gen.sql.Insert"
		}

		func (t *SQLAPIGenSQLInsertEntity) TplOutput() (out string,err error){
			out,err=gqttpl.TplOutput(t,t.TplName())
			return 
		}

	

		type SQLAPIGenSQLPaginateEntity struct{
			
				
					Limit int 
				
				
			
				
					Offset int 
				
				
			
				
					*SQLAPIGenSQLPaginateWhereEntity
				
				
			
			*gqttpl.DataVolumeMap
		}

		func (t *SQLAPIGenSQLPaginateEntity) TplName() string{
			return "sql.api.gen.sql.Paginate"
		}

		func (t *SQLAPIGenSQLPaginateEntity) TplOutput() (out string,err error){
			out,err=gqttpl.TplOutput(t,t.TplName())
			return 
		}

	

		type SQLAPIGenSQLPaginateTotalEntity struct{
			
				
					*SQLAPIGenSQLPaginateWhereEntity
				
				
			
			*gqttpl.DataVolumeMap
		}

		func (t *SQLAPIGenSQLPaginateTotalEntity) TplName() string{
			return "sql.api.gen.sql.PaginateTotal"
		}

		func (t *SQLAPIGenSQLPaginateTotalEntity) TplOutput() (out string,err error){
			out,err=gqttpl.TplOutput(t,t.TplName())
			return 
		}

	

		type SQLAPIGenSQLPaginateWhereEntity struct{
			
			*gqttpl.DataVolumeMap
		}

		func (t *SQLAPIGenSQLPaginateWhereEntity) TplName() string{
			return "sql.api.gen.sql.PaginateWhere"
		}

		func (t *SQLAPIGenSQLPaginateWhereEntity) TplOutput() (out string,err error){
			out,err=gqttpl.TplOutput(t,t.TplName())
			return 
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
				
				
			
			*gqttpl.DataVolumeMap
		}

		func (t *SQLAPIGenSQLUpdateEntity) TplName() string{
			return "sql.api.gen.sql.Update"
		}

		func (t *SQLAPIGenSQLUpdateEntity) TplOutput() (out string,err error){
			out,err=gqttpl.TplOutput(t,t.TplName())
			return 
		}

	

		type SQLExampleGenSQLGetAllByExampleIDListEntity struct{
			
				
					ExampleIDList []string 
				
				
			
			*gqttpl.DataVolumeMap
		}

		func (t *SQLExampleGenSQLGetAllByExampleIDListEntity) TplName() string{
			return "sql.example.gen.sql.GetAllByExampleIDList"
		}

		func (t *SQLExampleGenSQLGetAllByExampleIDListEntity) TplOutput() (out string,err error){
			out,err=gqttpl.TplOutput(t,t.TplName())
			return 
		}

	

		type SQLExampleGenSQLGetByExampleIDEntity struct{
			
				
					ExampleID string 
				
				
			
			*gqttpl.DataVolumeMap
		}

		func (t *SQLExampleGenSQLGetByExampleIDEntity) TplName() string{
			return "sql.example.gen.sql.GetByExampleID"
		}

		func (t *SQLExampleGenSQLGetByExampleIDEntity) TplOutput() (out string,err error){
			out,err=gqttpl.TplOutput(t,t.TplName())
			return 
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
				
				
			
			*gqttpl.DataVolumeMap
		}

		func (t *SQLExampleGenSQLInsertEntity) TplName() string{
			return "sql.example.gen.sql.Insert"
		}

		func (t *SQLExampleGenSQLInsertEntity) TplOutput() (out string,err error){
			out,err=gqttpl.TplOutput(t,t.TplName())
			return 
		}

	

		type SQLExampleGenSQLPaginateEntity struct{
			
				
					Limit int 
				
				
			
				
					Offset int 
				
				
			
				
					*SQLExampleGenSQLPaginateWhereEntity
				
				
			
			*gqttpl.DataVolumeMap
		}

		func (t *SQLExampleGenSQLPaginateEntity) TplName() string{
			return "sql.example.gen.sql.Paginate"
		}

		func (t *SQLExampleGenSQLPaginateEntity) TplOutput() (out string,err error){
			out,err=gqttpl.TplOutput(t,t.TplName())
			return 
		}

	

		type SQLExampleGenSQLPaginateTotalEntity struct{
			
				
					*SQLExampleGenSQLPaginateWhereEntity
				
				
			
			*gqttpl.DataVolumeMap
		}

		func (t *SQLExampleGenSQLPaginateTotalEntity) TplName() string{
			return "sql.example.gen.sql.PaginateTotal"
		}

		func (t *SQLExampleGenSQLPaginateTotalEntity) TplOutput() (out string,err error){
			out,err=gqttpl.TplOutput(t,t.TplName())
			return 
		}

	

		type SQLExampleGenSQLPaginateWhereEntity struct{
			
			*gqttpl.DataVolumeMap
		}

		func (t *SQLExampleGenSQLPaginateWhereEntity) TplName() string{
			return "sql.example.gen.sql.PaginateWhere"
		}

		func (t *SQLExampleGenSQLPaginateWhereEntity) TplOutput() (out string,err error){
			out,err=gqttpl.TplOutput(t,t.TplName())
			return 
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
				
				
			
			*gqttpl.DataVolumeMap
		}

		func (t *SQLExampleGenSQLUpdateEntity) TplName() string{
			return "sql.example.gen.sql.Update"
		}

		func (t *SQLExampleGenSQLUpdateEntity) TplOutput() (out string,err error){
			out,err=gqttpl.TplOutput(t,t.TplName())
			return 
		}

	

		type SQLMarkdownGenSQLGetAllByMarkdownIDListEntity struct{
			
				
					MarkdownIDList []string 
				
				
			
			*gqttpl.DataVolumeMap
		}

		func (t *SQLMarkdownGenSQLGetAllByMarkdownIDListEntity) TplName() string{
			return "sql.markdown.gen.sql.GetAllByMarkdownIDList"
		}

		func (t *SQLMarkdownGenSQLGetAllByMarkdownIDListEntity) TplOutput() (out string,err error){
			out,err=gqttpl.TplOutput(t,t.TplName())
			return 
		}

	

		type SQLMarkdownGenSQLGetByMarkdownIDEntity struct{
			
				
					MarkdownID string 
				
				
			
			*gqttpl.DataVolumeMap
		}

		func (t *SQLMarkdownGenSQLGetByMarkdownIDEntity) TplName() string{
			return "sql.markdown.gen.sql.GetByMarkdownID"
		}

		func (t *SQLMarkdownGenSQLGetByMarkdownIDEntity) TplOutput() (out string,err error){
			out,err=gqttpl.TplOutput(t,t.TplName())
			return 
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
				
				
			
			*gqttpl.DataVolumeMap
		}

		func (t *SQLMarkdownGenSQLInsertEntity) TplName() string{
			return "sql.markdown.gen.sql.Insert"
		}

		func (t *SQLMarkdownGenSQLInsertEntity) TplOutput() (out string,err error){
			out,err=gqttpl.TplOutput(t,t.TplName())
			return 
		}

	

		type SQLMarkdownGenSQLPaginateEntity struct{
			
				
					Limit int 
				
				
			
				
					Offset int 
				
				
			
				
					*SQLMarkdownGenSQLPaginateWhereEntity
				
				
			
			*gqttpl.DataVolumeMap
		}

		func (t *SQLMarkdownGenSQLPaginateEntity) TplName() string{
			return "sql.markdown.gen.sql.Paginate"
		}

		func (t *SQLMarkdownGenSQLPaginateEntity) TplOutput() (out string,err error){
			out,err=gqttpl.TplOutput(t,t.TplName())
			return 
		}

	

		type SQLMarkdownGenSQLPaginateTotalEntity struct{
			
				
					*SQLMarkdownGenSQLPaginateWhereEntity
				
				
			
			*gqttpl.DataVolumeMap
		}

		func (t *SQLMarkdownGenSQLPaginateTotalEntity) TplName() string{
			return "sql.markdown.gen.sql.PaginateTotal"
		}

		func (t *SQLMarkdownGenSQLPaginateTotalEntity) TplOutput() (out string,err error){
			out,err=gqttpl.TplOutput(t,t.TplName())
			return 
		}

	

		type SQLMarkdownGenSQLPaginateWhereEntity struct{
			
			*gqttpl.DataVolumeMap
		}

		func (t *SQLMarkdownGenSQLPaginateWhereEntity) TplName() string{
			return "sql.markdown.gen.sql.PaginateWhere"
		}

		func (t *SQLMarkdownGenSQLPaginateWhereEntity) TplOutput() (out string,err error){
			out,err=gqttpl.TplOutput(t,t.TplName())
			return 
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
				
				
			
			*gqttpl.DataVolumeMap
		}

		func (t *SQLMarkdownGenSQLUpdateEntity) TplName() string{
			return "sql.markdown.gen.sql.Update"
		}

		func (t *SQLMarkdownGenSQLUpdateEntity) TplOutput() (out string,err error){
			out,err=gqttpl.TplOutput(t,t.TplName())
			return 
		}

	

		type SQLParameterGenSQLGetAllByParameterIDListEntity struct{
			
				
					ParameterIDList []string 
				
				
			
			*gqttpl.DataVolumeMap
		}

		func (t *SQLParameterGenSQLGetAllByParameterIDListEntity) TplName() string{
			return "sql.parameter.gen.sql.GetAllByParameterIDList"
		}

		func (t *SQLParameterGenSQLGetAllByParameterIDListEntity) TplOutput() (out string,err error){
			out,err=gqttpl.TplOutput(t,t.TplName())
			return 
		}

	

		type SQLParameterGenSQLGetByParameterIDEntity struct{
			
				
					ParameterID string 
				
				
			
			*gqttpl.DataVolumeMap
		}

		func (t *SQLParameterGenSQLGetByParameterIDEntity) TplName() string{
			return "sql.parameter.gen.sql.GetByParameterID"
		}

		func (t *SQLParameterGenSQLGetByParameterIDEntity) TplOutput() (out string,err error){
			out,err=gqttpl.TplOutput(t,t.TplName())
			return 
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
				
				
			
			*gqttpl.DataVolumeMap
		}

		func (t *SQLParameterGenSQLInsertEntity) TplName() string{
			return "sql.parameter.gen.sql.Insert"
		}

		func (t *SQLParameterGenSQLInsertEntity) TplOutput() (out string,err error){
			out,err=gqttpl.TplOutput(t,t.TplName())
			return 
		}

	

		type SQLParameterGenSQLPaginateEntity struct{
			
				
					Limit int 
				
				
			
				
					Offset int 
				
				
			
				
					*SQLParameterGenSQLPaginateWhereEntity
				
				
			
			*gqttpl.DataVolumeMap
		}

		func (t *SQLParameterGenSQLPaginateEntity) TplName() string{
			return "sql.parameter.gen.sql.Paginate"
		}

		func (t *SQLParameterGenSQLPaginateEntity) TplOutput() (out string,err error){
			out,err=gqttpl.TplOutput(t,t.TplName())
			return 
		}

	

		type SQLParameterGenSQLPaginateTotalEntity struct{
			
				
					*SQLParameterGenSQLPaginateWhereEntity
				
				
			
			*gqttpl.DataVolumeMap
		}

		func (t *SQLParameterGenSQLPaginateTotalEntity) TplName() string{
			return "sql.parameter.gen.sql.PaginateTotal"
		}

		func (t *SQLParameterGenSQLPaginateTotalEntity) TplOutput() (out string,err error){
			out,err=gqttpl.TplOutput(t,t.TplName())
			return 
		}

	

		type SQLParameterGenSQLPaginateWhereEntity struct{
			
			*gqttpl.DataVolumeMap
		}

		func (t *SQLParameterGenSQLPaginateWhereEntity) TplName() string{
			return "sql.parameter.gen.sql.PaginateWhere"
		}

		func (t *SQLParameterGenSQLPaginateWhereEntity) TplOutput() (out string,err error){
			out,err=gqttpl.TplOutput(t,t.TplName())
			return 
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
				
				
			
			*gqttpl.DataVolumeMap
		}

		func (t *SQLParameterGenSQLUpdateEntity) TplName() string{
			return "sql.parameter.gen.sql.Update"
		}

		func (t *SQLParameterGenSQLUpdateEntity) TplOutput() (out string,err error){
			out,err=gqttpl.TplOutput(t,t.TplName())
			return 
		}

	

		type SQLServerGenSQLGetAllByServerIDListEntity struct{
			
				
					ServerIDList []string 
				
				
			
			*gqttpl.DataVolumeMap
		}

		func (t *SQLServerGenSQLGetAllByServerIDListEntity) TplName() string{
			return "sql.server.gen.sql.GetAllByServerIDList"
		}

		func (t *SQLServerGenSQLGetAllByServerIDListEntity) TplOutput() (out string,err error){
			out,err=gqttpl.TplOutput(t,t.TplName())
			return 
		}

	

		type SQLServerGenSQLGetByServerIDEntity struct{
			
				
					ServerID string 
				
				
			
			*gqttpl.DataVolumeMap
		}

		func (t *SQLServerGenSQLGetByServerIDEntity) TplName() string{
			return "sql.server.gen.sql.GetByServerID"
		}

		func (t *SQLServerGenSQLGetByServerIDEntity) TplOutput() (out string,err error){
			out,err=gqttpl.TplOutput(t,t.TplName())
			return 
		}

	

		type SQLServerGenSQLInsertEntity struct{
			
				
					Description string 
				
				
			
				
					ExtensionIds string 
				
				
			
				
					Proxy string 
				
				
			
				
					ServerID string 
				
				
			
				
					ServiceID string 
				
				
			
				
					URL string 
				
				
			
			*gqttpl.DataVolumeMap
		}

		func (t *SQLServerGenSQLInsertEntity) TplName() string{
			return "sql.server.gen.sql.Insert"
		}

		func (t *SQLServerGenSQLInsertEntity) TplOutput() (out string,err error){
			out,err=gqttpl.TplOutput(t,t.TplName())
			return 
		}

	

		type SQLServerGenSQLPaginateEntity struct{
			
				
					Limit int 
				
				
			
				
					Offset int 
				
				
			
				
					*SQLServerGenSQLPaginateWhereEntity
				
				
			
			*gqttpl.DataVolumeMap
		}

		func (t *SQLServerGenSQLPaginateEntity) TplName() string{
			return "sql.server.gen.sql.Paginate"
		}

		func (t *SQLServerGenSQLPaginateEntity) TplOutput() (out string,err error){
			out,err=gqttpl.TplOutput(t,t.TplName())
			return 
		}

	

		type SQLServerGenSQLPaginateTotalEntity struct{
			
				
					*SQLServerGenSQLPaginateWhereEntity
				
				
			
			*gqttpl.DataVolumeMap
		}

		func (t *SQLServerGenSQLPaginateTotalEntity) TplName() string{
			return "sql.server.gen.sql.PaginateTotal"
		}

		func (t *SQLServerGenSQLPaginateTotalEntity) TplOutput() (out string,err error){
			out,err=gqttpl.TplOutput(t,t.TplName())
			return 
		}

	

		type SQLServerGenSQLPaginateWhereEntity struct{
			
			*gqttpl.DataVolumeMap
		}

		func (t *SQLServerGenSQLPaginateWhereEntity) TplName() string{
			return "sql.server.gen.sql.PaginateWhere"
		}

		func (t *SQLServerGenSQLPaginateWhereEntity) TplOutput() (out string,err error){
			out,err=gqttpl.TplOutput(t,t.TplName())
			return 
		}

	

		type SQLServerGenSQLUpdateEntity struct{
			
				
					Description string 
				
				
			
				
					ExtensionIds string 
				
				
			
				
					Proxy string 
				
				
			
				
					ServerID string 
				
				
			
				
					ServiceID string 
				
				
			
				
					URL string 
				
				
			
			*gqttpl.DataVolumeMap
		}

		func (t *SQLServerGenSQLUpdateEntity) TplName() string{
			return "sql.server.gen.sql.Update"
		}

		func (t *SQLServerGenSQLUpdateEntity) TplOutput() (out string,err error){
			out,err=gqttpl.TplOutput(t,t.TplName())
			return 
		}

	

		type SQLServiceGenSQLGetAllByServiceIDListEntity struct{
			
				
					ServiceIDList []string 
				
				
			
			*gqttpl.DataVolumeMap
		}

		func (t *SQLServiceGenSQLGetAllByServiceIDListEntity) TplName() string{
			return "sql.service.gen.sql.GetAllByServiceIDList"
		}

		func (t *SQLServiceGenSQLGetAllByServiceIDListEntity) TplOutput() (out string,err error){
			out,err=gqttpl.TplOutput(t,t.TplName())
			return 
		}

	

		type SQLServiceGenSQLGetByServiceIDEntity struct{
			
				
					ServiceID string 
				
				
			
			*gqttpl.DataVolumeMap
		}

		func (t *SQLServiceGenSQLGetByServiceIDEntity) TplName() string{
			return "sql.service.gen.sql.GetByServiceID"
		}

		func (t *SQLServiceGenSQLGetByServiceIDEntity) TplOutput() (out string,err error){
			out,err=gqttpl.TplOutput(t,t.TplName())
			return 
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
				
				
			
			*gqttpl.DataVolumeMap
		}

		func (t *SQLServiceGenSQLInsertEntity) TplName() string{
			return "sql.service.gen.sql.Insert"
		}

		func (t *SQLServiceGenSQLInsertEntity) TplOutput() (out string,err error){
			out,err=gqttpl.TplOutput(t,t.TplName())
			return 
		}

	

		type SQLServiceGenSQLPaginateEntity struct{
			
				
					Limit int 
				
				
			
				
					Offset int 
				
				
			
				
					*SQLServiceGenSQLPaginateWhereEntity
				
				
			
			*gqttpl.DataVolumeMap
		}

		func (t *SQLServiceGenSQLPaginateEntity) TplName() string{
			return "sql.service.gen.sql.Paginate"
		}

		func (t *SQLServiceGenSQLPaginateEntity) TplOutput() (out string,err error){
			out,err=gqttpl.TplOutput(t,t.TplName())
			return 
		}

	

		type SQLServiceGenSQLPaginateTotalEntity struct{
			
				
					*SQLServiceGenSQLPaginateWhereEntity
				
				
			
			*gqttpl.DataVolumeMap
		}

		func (t *SQLServiceGenSQLPaginateTotalEntity) TplName() string{
			return "sql.service.gen.sql.PaginateTotal"
		}

		func (t *SQLServiceGenSQLPaginateTotalEntity) TplOutput() (out string,err error){
			out,err=gqttpl.TplOutput(t,t.TplName())
			return 
		}

	

		type SQLServiceGenSQLPaginateWhereEntity struct{
			
			*gqttpl.DataVolumeMap
		}

		func (t *SQLServiceGenSQLPaginateWhereEntity) TplName() string{
			return "sql.service.gen.sql.PaginateWhere"
		}

		func (t *SQLServiceGenSQLPaginateWhereEntity) TplOutput() (out string,err error){
			out,err=gqttpl.TplOutput(t,t.TplName())
			return 
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
				
				
			
			*gqttpl.DataVolumeMap
		}

		func (t *SQLServiceGenSQLUpdateEntity) TplName() string{
			return "sql.service.gen.sql.Update"
		}

		func (t *SQLServiceGenSQLUpdateEntity) TplOutput() (out string,err error){
			out,err=gqttpl.TplOutput(t,t.TplName())
			return 
		}

	

		type SQLValidateSchemaGenSQLGetAllByValidateSchemaIDListEntity struct{
			
				
					ValidateSchemaIDList []string 
				
				
			
			*gqttpl.DataVolumeMap
		}

		func (t *SQLValidateSchemaGenSQLGetAllByValidateSchemaIDListEntity) TplName() string{
			return "sql.validate_schema.gen.sql.GetAllByValidateSchemaIDList"
		}

		func (t *SQLValidateSchemaGenSQLGetAllByValidateSchemaIDListEntity) TplOutput() (out string,err error){
			out,err=gqttpl.TplOutput(t,t.TplName())
			return 
		}

	

		type SQLValidateSchemaGenSQLGetByValidateSchemaIDEntity struct{
			
				
					ValidateSchemaID string 
				
				
			
			*gqttpl.DataVolumeMap
		}

		func (t *SQLValidateSchemaGenSQLGetByValidateSchemaIDEntity) TplName() string{
			return "sql.validate_schema.gen.sql.GetByValidateSchemaID"
		}

		func (t *SQLValidateSchemaGenSQLGetByValidateSchemaIDEntity) TplOutput() (out string,err error){
			out,err=gqttpl.TplOutput(t,t.TplName())
			return 
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
				
				
			
			*gqttpl.DataVolumeMap
		}

		func (t *SQLValidateSchemaGenSQLInsertEntity) TplName() string{
			return "sql.validate_schema.gen.sql.Insert"
		}

		func (t *SQLValidateSchemaGenSQLInsertEntity) TplOutput() (out string,err error){
			out,err=gqttpl.TplOutput(t,t.TplName())
			return 
		}

	

		type SQLValidateSchemaGenSQLPaginateEntity struct{
			
				
					Limit int 
				
				
			
				
					Offset int 
				
				
			
				
					*SQLValidateSchemaGenSQLPaginateWhereEntity
				
				
			
			*gqttpl.DataVolumeMap
		}

		func (t *SQLValidateSchemaGenSQLPaginateEntity) TplName() string{
			return "sql.validate_schema.gen.sql.Paginate"
		}

		func (t *SQLValidateSchemaGenSQLPaginateEntity) TplOutput() (out string,err error){
			out,err=gqttpl.TplOutput(t,t.TplName())
			return 
		}

	

		type SQLValidateSchemaGenSQLPaginateTotalEntity struct{
			
				
					*SQLValidateSchemaGenSQLPaginateWhereEntity
				
				
			
			*gqttpl.DataVolumeMap
		}

		func (t *SQLValidateSchemaGenSQLPaginateTotalEntity) TplName() string{
			return "sql.validate_schema.gen.sql.PaginateTotal"
		}

		func (t *SQLValidateSchemaGenSQLPaginateTotalEntity) TplOutput() (out string,err error){
			out,err=gqttpl.TplOutput(t,t.TplName())
			return 
		}

	

		type SQLValidateSchemaGenSQLPaginateWhereEntity struct{
			
			*gqttpl.DataVolumeMap
		}

		func (t *SQLValidateSchemaGenSQLPaginateWhereEntity) TplName() string{
			return "sql.validate_schema.gen.sql.PaginateWhere"
		}

		func (t *SQLValidateSchemaGenSQLPaginateWhereEntity) TplOutput() (out string,err error){
			out,err=gqttpl.TplOutput(t,t.TplName())
			return 
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
				
				
			
			*gqttpl.DataVolumeMap
		}

		func (t *SQLValidateSchemaGenSQLUpdateEntity) TplName() string{
			return "sql.validate_schema.gen.sql.Update"
		}

		func (t *SQLValidateSchemaGenSQLUpdateEntity) TplOutput() (out string,err error){
			out,err=gqttpl.TplOutput(t,t.TplName())
			return 
		}

	