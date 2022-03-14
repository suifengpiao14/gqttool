
{{define "GetByServiceID"}}
select * from `t_service`  where `service_id`=:ServiceID  and `deleted_at` is null;
{{end}}


{{define "PaginateWhere"}}
    and `deleted_at` is null;
{{end}}



{{define "PaginateTotal"}}
select count(*) as `count` from `t_service`  where 1=1 {{template "PaginateWhere" "."}} ;
{{end}}

  {{define "Paginate"}}
select * from `t_service`  where 1=1 {{template "PaginateWhere" "."}}  order by `updated_at` desc  limit :Offset,:Limit ;
{{end}}


{{define "Insert"}}
insert into `t_service` (`service_id`,`title`,`description`,`version`,`contact_ids`,`license`,`security`,`proxy`,`variables`)values
 (:ServiceID,:Title,:Description,:Version,:ContactIds,:License,:Security,:Proxy,:Variables);
{{end}}

{{define "Update"}}
{{$preComma:=newPreComma}}
 update `t_service` set {{if .ServiceID}} {{$preComma.PreComma}} `service_id`=:ServiceID {{end}} 
{{if .Title}} {{$preComma.PreComma}} `title`=:Title {{end}} 
{{if .Description}} {{$preComma.PreComma}} `description`=:Description {{end}} 
{{if .Version}} {{$preComma.PreComma}} `version`=:Version {{end}} 
{{if .ContactIds}} {{$preComma.PreComma}} `contact_ids`=:ContactIds {{end}} 
{{if .License}} {{$preComma.PreComma}} `license`=:License {{end}} 
{{if .Security}} {{$preComma.PreComma}} `security`=:Security {{end}} 
{{if .Proxy}} {{$preComma.PreComma}} `proxy`=:Proxy {{end}} 
{{if .Variables}} {{$preComma.PreComma}} `variables`=:Variables {{end}}  where `service_id`=:ServiceID;
{{end}}


{{define "Del"}}
update `t_service` set `deleted_at`={{currentTime .}},`operator_id`=:OperatorID,`operator`=:Operator where `service_id`=:ServiceID;
{{end}}
