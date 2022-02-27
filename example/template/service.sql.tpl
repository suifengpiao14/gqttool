
{{define "getByServiceId"}}
select * from `service` where `service_id`=:ServiceID and  `deleted_at` is null;
{{end}}

{{define "total"}}
   select count(*) as `count` from `service` where 1=1 
   {{if .Title}} and `title` like "%:Title%" {{end}};
{{end}}

{{define "list"}}
   select * from `service` where 1=1 
   {{if .Title}} and `title` like "%:Title%" {{end}}
   limit :Offset,:Limit;
{{end}}

{{define "insert"}}
  insert into `service` (`service_id`,`title`,`description`,`version`,`contact_ids`,`license`,`security`,`proxy`,`variables`)
  values(:ServiceID,:Title,:Description,:Version,:ContactIds,:License,:Security,:Proxy,:Variables);
{{end}}

{{define "update"}}
{{$preComma:=newPreComma}}
    update `service` set
     {{if .Title}} {{$preComma.PreComma}} `title`=:Title{{end}}
     {{if .Description}} {{$preComma.PreComma}} `description`=:Description{{end}}
     {{if .Version}}  {{$preComma.PreComma}}`version`=:Version{{end}}
     {{if .ContactIds}} {{$preComma.PreComma}} `contact_ids`=:ContactIds{{end}}
     {{if .License}}  {{$preComma.PreComma}}`license`=:License{{end}}
     {{if .Security}} {{$preComma.PreComma}} `security`=:Security{{end}}
     {{if .Proxy}} {{$preComma.PreComma}} `proxy`=:Proxy{{end}}
     {{if .Variables}}  {{$preComma.PreComma}}`variables` =:Variables {{end}}
       where `service_id`=:ServiceID;
 {{end}}

  {{define "del"}}
 update `service` set `deleted_at`={{currentTime . }} where `service_id`=:ServiceID;
 {{end}}