
{{define "GetByExampleID"}}
select * from `t_example`  where `example_id`=:ExampleID  and `deleted_at` is null;
{{end}}

{{define "GetAllByExampleIDList"}}
select * from `t_example`  where `example_id` in ({{in . .ExampleIDList}})  and `deleted_at` is null;
{{end}}


{{define "PaginateWhere"}}
  
{{end}}

{{define "PaginateTotal"}}
select count(*) as `count` from `t_example`  where 1=1 {{template "PaginateWhere" "."}}   and `deleted_at` is null;
{{end}}

  {{define "Paginate"}}
select * from `t_example`  where 1=1 {{template "PaginateWhere" "."}}   and `deleted_at` is null order by `updated_at` desc  limit :Offset,:Limit ;
{{end}}


{{define "Insert"}}
insert into `t_example` (`example_id`,`service_id`,`api_id`,`tag`,`title`,`summary`,`request`,`response`)values
 (:ExampleID,:ServiceID,:APIID,:Tag,:Title,:Summary,:Request,:Response);
{{end}}

{{define "Update"}}
{{$preComma:=newPreComma}}
 update `t_example` set {{if .ExampleID}} {{$preComma.PreComma}} `example_id`=:ExampleID {{end}} 
{{if .ServiceID}} {{$preComma.PreComma}} `service_id`=:ServiceID {{end}} 
{{if .APIID}} {{$preComma.PreComma}} `api_id`=:APIID {{end}} 
{{if .Tag}} {{$preComma.PreComma}} `tag`=:Tag {{end}} 
{{if .Title}} {{$preComma.PreComma}} `title`=:Title {{end}} 
{{if .Summary}} {{$preComma.PreComma}} `summary`=:Summary {{end}} 
{{if .Request}} {{$preComma.PreComma}} `request`=:Request {{end}} 
{{if .Response}} {{$preComma.PreComma}} `response`=:Response {{end}}  where `example_id`=:ExampleID;
{{end}}


