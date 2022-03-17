
{{define "GetByMarkdownID"}}
select * from `t_markdown`  where `markdown_id`=:MarkdownID  and `deleted_at` is null;
{{end}}

{{define "GetAllByMarkdownIDList"}}
select * from `t_markdown`  where `markdown_id` in ({{in . .MarkdownIDList}})  and `deleted_at` is null;
{{end}}


{{define "PaginateWhere"}}
  
{{end}}

{{define "PaginateTotal"}}
select count(*) as `count` from `t_markdown`  where 1=1 {{template "PaginateWhere" "."}}   and `deleted_at` is null;
{{end}}

  {{define "Paginate"}}
select * from `t_markdown`  where 1=1 {{template "PaginateWhere" "."}}   and `deleted_at` is null order by `updated_at` desc  limit :Offset,:Limit ;
{{end}}


{{define "Insert"}}
insert into `t_markdown` (`markdown_id`,`service_id`,`api_id`,`name`,`title`,`markdown`,`content`,`owner_id`,`owner_name`)values
 (:MarkdownID,:ServiceID,:APIID,:Name,:Title,:Markdown,:Content,:OwnerID,:OwnerName);
{{end}}

{{define "Update"}}
{{$preComma:=newPreComma}}
 update `t_markdown` set {{if .MarkdownID}} {{$preComma.PreComma}} `markdown_id`=:MarkdownID {{end}} 
{{if .ServiceID}} {{$preComma.PreComma}} `service_id`=:ServiceID {{end}} 
{{if .APIID}} {{$preComma.PreComma}} `api_id`=:APIID {{end}} 
{{if .Name}} {{$preComma.PreComma}} `name`=:Name {{end}} 
{{if .Title}} {{$preComma.PreComma}} `title`=:Title {{end}} 
{{if .Markdown}} {{$preComma.PreComma}} `markdown`=:Markdown {{end}} 
{{if .Content}} {{$preComma.PreComma}} `content`=:Content {{end}} 
{{if .OwnerID}} {{$preComma.PreComma}} `owner_id`=:OwnerID {{end}} 
{{if .OwnerName}} {{$preComma.PreComma}} `owner_name`=:OwnerName {{end}}  where `markdown_id`=:MarkdownID;
{{end}}


