package gqttool

import (
	"fmt"
	"testing"

	"github.com/suifengpiao14/gqt/v2"
	"github.com/suifengpiao14/gqt/v2/gqttpl"
	"github.com/suifengpiao14/gqttool/example"
)

func TestGetTemplateNames(t *testing.T) {
	sqtTpl := "select count(*) as `count` from `t_example`  where 1=1 {{template \"PaginateWhere\" \".\"}} {{template \"aab\" \".\"}} ;"
	nameList := GetTemplateNames(sqtTpl)
	fmt.Println(nameList)
}

func TestParsSqlTplVariable(t *testing.T) {
	sqlTpl := `
{{$preComma:=newPreComma}}
    update service set
     {{if .Title}} {{$preComma.PreComma}} title=:Title{{end}}
     {{if .Description}} {{$preComma.PreComma}} description=:Description{{end}}
     {{if .Version}}  {{$preComma.PreComma}}version=:Version{{end}}
     {{if .ContactIds}} {{$preComma.PreComma}} contact_ids=:ContactIds{{end}}
     {{if .License}}  {{$preComma.PreComma}}license=:License{{end}}
     {{if .Security}} {{$preComma.PreComma}} security=:Security{{end}}
     {{if .Proxy}} {{$preComma.PreComma}} proxy=:Proxy{{end}}
     {{if .Variables}}  {{$preComma.PreComma}}variables =:Variables {{end}}
	 {{if .Variables}} {{in . .List}} {{end}}
       where {{func .Func}} service_id=:ServiceID limit :Offset,:limit;
  `
	tplDefine := &TPLDefineText{
		Text:      sqlTpl,
		Namespace: "a",
	}
	veriableList := ParsSqlTplVariable(tplDefine)
	fmt.Println(veriableList)
}

func TestGetDefineName(t *testing.T) {
	sqlTpl := `
	{{define "list"}}
   select * from service where 1=1 
   {{if .Title}} and title like "%:Title%" {{end}}
   limit :Offset,:Limit;
{{end}}
	`
	defineName, err := GetDefineName(sqlTpl)
	if err != nil {
		panic(err)
	}
	fmt.Println(defineName)
}

func TestParseDefine(t *testing.T) {
	sqlTpl := `
	{{define "list"}}
   select * from service where 1=1 
   {{if .Title}} and title like "%:Title%" {{end}}
   limit :Offset,:Limit;
{{end}}

{{define "insert"}}
  insert into service (service_id,title,description,version,contact_ids,license,security,proxy,variables)
  values(:ServiceID,:Title,:Description,:Version,:ContactIds,:License,:Security,:Proxy,:Variables);
{{end}}
	{{define "update"}}
{{$preComma:=newPreComma}}
    update service set
     {{if .Title}} {{$preComma.PreComma}} title=:Title{{end}}
     {{if .Description}} {{$preComma.PreComma}} description=:Description{{end}}
     {{if .Version}}  {{$preComma.PreComma}}version=:Version{{end}}
     {{if .ContactIds}} {{$preComma.PreComma}} contact_ids=:ContactIds{{end}}
     {{if .License}}  {{$preComma.PreComma}}license=:License{{end}}
     {{if .Security}} {{$preComma.PreComma}} security=:Security{{end}}
     {{if .Proxy}} {{$preComma.PreComma}} proxy=:Proxy{{end}}
     {{if .Variables}}  {{$preComma.PreComma}}variables =:Variables {{end}}
       where service_id=:ServiceID;
 {{end}}
	`
	defineList := ManualParseDefine(sqlTpl, "", gqttpl.LeftDelim, gqttpl.RightDelim)
	fmt.Println(defineList[2])

}

func TestParseSQLTPLTableName(t *testing.T) {
	sqlTpl := " aa  update `a` b	"
	tableNameList, err := ParseSQLTPLTableName(sqlTpl)
	if err != nil {
		panic(err)
	}
	fmt.Println(tableNameList)
}

func TestRepositoryEntity(t *testing.T) {
	ddlSqlTpl := ddlSqlTplData()
	ddlList := []string{ddlSqlTpl.DDL}
	cfg := &DatabaseConfig{}
	tableList, err := GenerateTable(ddlList, cfg)
	if err != nil {
		panic(err)
	}
	sqlTplDefine := &TPLDefineText{
		Text:      ddlSqlTpl.SQLTpl,
		Name:      "test",
		Namespace: "ddl",
	}
	entityStruct, err := SQLEntity(sqlTplDefine, tableList)
	if err != nil {
		panic(err)
	}
	fmt.Println(entityStruct)
}

type DDLSqlTpl struct {
	DDL    string
	SQLTpl string
}

func ddlSqlTplData() (ddlSqlTpl *DDLSqlTpl) {
	ddlService := `CREATE TABLE if not exists service (
		service_id varchar(64) NOT NULL DEFAULT '' COMMENT '????????????',
		title varchar(255) NOT NULL DEFAULT '' COMMENT '??????',
		description varchar(255) NOT NULL DEFAULT '' COMMENT '??????',
		version varchar(20) NOT NULL DEFAULT '' COMMENT '??????',
		contact_ids varchar(100) NOT NULL DEFAULT '' COMMENT '?????????',
		license char(255) NOT NULL DEFAULT '' COMMENT '??????',
		security char(255) NOT NULL DEFAULT '' COMMENT '??????',
		proxy char(60) NOT NULL DEFAULT '' COMMENT '????????????',
		variables text NOT NULL COMMENT 'json?????????',
		created_at datetime DEFAULT CURRENT_TIMESTAMP COMMENT '????????????',
		updated_at datetime DEFAULT CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP COMMENT '????????????',
		deleted_at datetime  DEFAULT null COMMENT '????????????',
		PRIMARY KEY (service_id)
	  ) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT="??????/????????????";
	  `
	sqlTplService := `

{{define "getByServiceId"}}
select * from service where service_id=:ServiceID and  deleted_at is null;
{{end}}

{{define "total"}}
   select count(*) as count from service where 1=1 
   {{if .Title}} and title like "%:Title%" {{end}};
{{end}}

{{define "list"}}
   select * from service where 1=1 
   {{if .Title}} and title like "%:Title%" {{end}}
   limit :Offset,:Limit;
{{end}}

{{define "insert"}}
  insert into service (service_id,title,description,version,contact_ids,license,security,proxy,variables)
  values(:ServiceID,:Title,:Description,:Version,:ContactIds,:License,:Security,:Proxy,:Variables);
{{end}}

{{define "update"}}
{{$preComma:=newPreComma}}
    update service set
     {{if .Title}} {{$preComma.PreComma}} title=:Title{{end}}
     {{if .Description}} {{$preComma.PreComma}} description=:Description{{end}}
     {{if .Version}}  {{$preComma.PreComma}}version=:Version{{end}}
     {{if .ContactIds}} {{$preComma.PreComma}} contact_ids=:ContactIds{{end}}
     {{if .License}}  {{$preComma.PreComma}}license=:License{{end}}
     {{if .Security}} {{$preComma.PreComma}} security=:Security{{end}}
     {{if .Proxy}} {{$preComma.PreComma}} proxy=:Proxy{{end}}
     {{if .Variables}}  {{$preComma.PreComma}}variables =:Variables {{end}}
       where service_id=:ServiceID;
 {{end}}

  {{define "del"}}
 update service set deleted_at={{currentTime . }} where service_id=:ServiceID;
 {{end}}
	`
	ddlSqlTpl = &DDLSqlTpl{
		DDL:    ddlService,
		SQLTpl: sqlTplService,
	}
	return
}

func TestWhereConditon(t *testing.T) {
	entity := example.SQLAPISQLPaginateEntity{
		Limit:  10,
		Offset: 20,
	}
	// entity.SQLAPIGenSQLPaginateWhereEntity = example.SQLAPIGenSQLPaginateWhereEntity{
	// 	APIIDList: []int{1, 3},
	// }
	repo := gqt.NewRepositorySQL()
	err := repo.AddByDir("example/template", gqt.TemplatefuncMap)
	if err != nil {
		panic(err)
	}
	sqlRow, err := repo.GetSQL(&entity)
	if err != nil {
		panic(err)
	}
	fmt.Println(sqlRow.SQL)
}

// func TestCurlEntity(t *testing.T) {
// 	repo := gqtcurl.NewRepositoryCURL()
// 	err := repo.AddByDir("example/template", gqtcurl.TemplatefuncMap)
// 	if err != nil {
// 		panic(err)
// 	}
// 	CURLEntity := example.CurlOrderCurlGetOrderByOrderNumberEntity{
// 		SecretKey: "mafera15478515",
// 		ServiceID: "100354",
// 	}
// 	CURLEntity.OrderNumber = "15454"
// 	curlRow, err := repo.GetCURLByTplEntity(&CURLEntity)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(curlRow.RequestData)

// }
