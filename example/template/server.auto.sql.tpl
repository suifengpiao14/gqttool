
{{define "GetByServerID"}}
select * from `t_server`  where `server_id`=:ServerID  and `deleted_at` is null;
{{end}}


{{define "PaginateWhere"}}
    and `deleted_at` is null;
{{end}}



{{define "PaginateTotal"}}
select count(*) as `count` from `t_server`  where 1=1 {{template "PaginateWhere" "."}} ;
{{end}}

  {{define "Paginate"}}
select * from `t_server`  where 1=1 {{template "PaginateWhere" "."}}  order by `updated_at` desc  limit :Offset,:Limit ;
{{end}}


{{define "Insert"}}
insert into `t_server` (`server_id`,`service_id`,`url`,`description`,`proxy`,`extension_ids`)values
 (:ServerID,:ServiceID,:URL,:Description,:Proxy,:ExtensionIds);
{{end}}

{{define "Update"}}
{{$preComma:=newPreComma}}
 update `t_server` set {{if .ServerID}} {{$preComma.PreComma}} `server_id`=:ServerID {{end}} 
{{if .ServiceID}} {{$preComma.PreComma}} `service_id`=:ServiceID {{end}} 
{{if .URL}} {{$preComma.PreComma}} `url`=:URL {{end}} 
{{if .Description}} {{$preComma.PreComma}} `description`=:Description {{end}} 
{{if .Proxy}} {{$preComma.PreComma}} `proxy`=:Proxy {{end}} 
{{if .ExtensionIds}} {{$preComma.PreComma}} `extension_ids`=:ExtensionIds {{end}}  where `server_id`=:ServerID;
{{end}}


{{define "Del"}}
update `t_server` set `deleted_at`={{currentTime .}},`operator_id`=:OperatorID,`operator`=:Operator where `server_id`=:ServerID;
{{end}}
