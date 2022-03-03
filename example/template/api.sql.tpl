{{define "GetByAPIID"}}
select * from `api`  where `api_id`=:APIID  and `deleted_at` is null;
{{end}}

{{define "Total"}}
select count(*) as `count` from `api`  where 1=1   and `deleted_at` is null;
{{end}}

 {{define "Paginate"}}
select * from `api`  where 1=1   and `deleted_at` is null  order by `updated_at` desc  limit :Offset,:Limit ;
{{end}}

{{define "Insert"}}
insert into `api` (`api_id`,`service_id`,`name`,`title`,`tags`,`uri`,`summary`,`description`)values
 (:APIID,:ServiceID,:Name,:Title,:Tags,:URI,:Summary,:Description);
{{end}}

{{define "Update"}}
{{$preComma:=newPreComma}}
 update `api` set {{if .APIID}} {{$preComma.PreComma}} `api_id`=:APIID {{end}} 
{{if .ServiceID}} {{$preComma.PreComma}} `service_id`=:ServiceID {{end}} 
{{if .Name}} {{$preComma.PreComma}} `name`=:Name {{end}} 
{{if .Title}} {{$preComma.PreComma}} `title`=:Title {{end}} 
{{if .Tags}} {{$preComma.PreComma}} `tags`=:Tags {{end}} 
{{if .URI}} {{$preComma.PreComma}} `uri`=:URI {{end}} 
{{if .Summary}} {{$preComma.PreComma}} `summary`=:Summary {{end}} 
{{if .Description}} {{$preComma.PreComma}} `description`=:Description {{end}}  where `api_id`=:APIID;
{{end}}

{{define "Del"}}
update `api` set `deleted_at`={{currentTime}} where `api_id`=:APIID;
{{end}}
