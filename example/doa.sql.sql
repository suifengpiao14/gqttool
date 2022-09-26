insert ignore into `source` (`source_id`,`source_type`,`config`) values('database_name','SQL','{"logLevel":"debug","dsn":"root:123456@tcp(mysql_address:3306)/database_name?charset=utf8&timeout=1s&readTimeout=5s&writeTimeout=5s&parseTime=False&loc=Local&multiStatements=true","timeout":30}');
insert ignore into `template` (`template_id`,`type`,`title`,`description`,`source_id`,`tpl`) values('DatabaseNameParameterGetByParameterID','SQL','database name 请求/响应参数 获取 通过 ID','database name 请求/响应参数 获取 通过 ID','database_name','{{define "GetByParameterID"}} select * from `t_parameter` where `parameter_id`=:ParameterID and `deleted_at` is null; {{end}}');
insert ignore into `api` (`api_id`,`title`,`description`,`method`,`route`,`template_ids`,`exec`,`input`,`output`,`input_tpl`,`output_tpl`) values('database_name-t_parameter-GetByParameterID','database name 请求/响应参数 获取 通过 ID','database name 请求/响应参数 获取 通过 ID','POST','/api/database_name/v1/parameter/get_by_parameter_id','DatabaseNameParameterGetByParameterID','{{define "main"}}
{{execSQLTpl . "GetByParameterID"}}
{{end}}','version=http://json-schema.org/draft-07/schema,id=input,direction=in
fullname=parameterId,dst=ParameterID,required','version=http://json-schema.org/draft-07/schema,id=output,direction=out
fullname=allowEmptyValue,src=GetByParameterIDOut.#.allow_empty_value,required
fullname=allowReserved,src=GetByParameterIDOut.#.allow_reserved,required
fullname=apiId,src=GetByParameterIDOut.#.api_id,required
fullname=createdAt,src=GetByParameterIDOut.#.created_at,required
fullname=deletedAt,src=GetByParameterIDOut.#.deleted_at,required
fullname=deprecated,src=GetByParameterIDOut.#.deprecated,required
fullname=description,src=GetByParameterIDOut.#.description,required
fullname=example,src=GetByParameterIDOut.#.example,required
fullname=explode,src=GetByParameterIDOut.#.explode,required
fullname=fullName,src=GetByParameterIDOut.#.full_name,required
fullname=httpStatus,src=GetByParameterIDOut.#.http_status,required
fullname=method,src=GetByParameterIDOut.#.method,required
fullname=name,src=GetByParameterIDOut.#.name,required
fullname=parameterId,src=GetByParameterIDOut.#.parameter_id,required
fullname=position,src=GetByParameterIDOut.#.position,required
fullname=required,src=GetByParameterIDOut.#.required,required
fullname=serialize,src=GetByParameterIDOut.#.serialize,required
fullname=serviceId,src=GetByParameterIDOut.#.service_id,required
fullname=tag,src=GetByParameterIDOut.#.tag,required
fullname=title,src=GetByParameterIDOut.#.title,required
fullname=type,src=GetByParameterIDOut.#.type,required
fullname=updatedAt,src=GetByParameterIDOut.#.updated_at,required
fullname=validateSchemaId,src=GetByParameterIDOut.#.validate_schema_id,required','{{define "input"}}
{{getSetValue . "ParameterID" "input.parameterId"}}
{{end}}','{{define "output"}}
{{getSetValue . "output.allowEmptyValue" "GetByParameterIDOut.#.allow_empty_value"}}
{{getSetValue . "output.allowReserved" "GetByParameterIDOut.#.allow_reserved"}}
{{getSetValue . "output.apiId" "GetByParameterIDOut.#.api_id"}}
{{getSetValue . "output.createdAt" "GetByParameterIDOut.#.created_at"}}
{{getSetValue . "output.deletedAt" "GetByParameterIDOut.#.deleted_at"}}
{{getSetValue . "output.deprecated" "GetByParameterIDOut.#.deprecated"}}
{{getSetValue . "output.description" "GetByParameterIDOut.#.description"}}
{{getSetValue . "output.example" "GetByParameterIDOut.#.example"}}
{{getSetValue . "output.explode" "GetByParameterIDOut.#.explode"}}
{{getSetValue . "output.fullName" "GetByParameterIDOut.#.full_name"}}
{{getSetValue . "output.httpStatus" "GetByParameterIDOut.#.http_status"}}
{{getSetValue . "output.method" "GetByParameterIDOut.#.method"}}
{{getSetValue . "output.name" "GetByParameterIDOut.#.name"}}
{{getSetValue . "output.parameterId" "GetByParameterIDOut.#.parameter_id"}}
{{getSetValue . "output.position" "GetByParameterIDOut.#.position"}}
{{getSetValue . "output.required" "GetByParameterIDOut.#.required"}}
{{getSetValue . "output.serialize" "GetByParameterIDOut.#.serialize"}}
{{getSetValue . "output.serviceId" "GetByParameterIDOut.#.service_id"}}
{{getSetValue . "output.tag" "GetByParameterIDOut.#.tag"}}
{{getSetValue . "output.title" "GetByParameterIDOut.#.title"}}
{{getSetValue . "output.type" "GetByParameterIDOut.#.type"}}
{{getSetValue . "output.updatedAt" "GetByParameterIDOut.#.updated_at"}}
{{getSetValue . "output.validateSchemaId" "GetByParameterIDOut.#.validate_schema_id"}}
{{end}}');
insert ignore into `template` (`template_id`,`type`,`title`,`description`,`source_id`,`tpl`) values('DatabaseNameParameterGetAllByParameterIDList','SQL','database name 请求/响应参数 获取 所有 通过 ID 列表','database name 请求/响应参数 获取 所有 通过 ID 列表','database_name','{{define "GetAllByParameterIDList"}} select * from `t_parameter` where `parameter_id` in ({{in . .ParameterIDList}}) and `deleted_at` is null; {{end}}');
insert ignore into `api` (`api_id`,`title`,`description`,`method`,`route`,`template_ids`,`exec`,`input`,`output`,`input_tpl`,`output_tpl`) values('database_name-t_parameter-GetAllByParameterIDList','database name 请求/响应参数 获取 所有 通过 ID 列表','database name 请求/响应参数 获取 所有 通过 ID 列表','POST','/api/database_name/v1/parameter/get_all_by_parameter_id_list','DatabaseNameParameterGetAllByParameterIDList','{{define "main"}}
{{execSQLTpl . "GetAllByParameterIDList"}}
{{end}}','version=http://json-schema.org/draft-07/schema,id=input,direction=in
fullname=parameterIdList,dst=ParameterIDList,required','version=http://json-schema.org/draft-07/schema,id=output,direction=out
fullname=allowEmptyValue,src=GetAllByParameterIDListOut.#.allow_empty_value,required
fullname=allowReserved,src=GetAllByParameterIDListOut.#.allow_reserved,required
fullname=apiId,src=GetAllByParameterIDListOut.#.api_id,required
fullname=createdAt,src=GetAllByParameterIDListOut.#.created_at,required
fullname=deletedAt,src=GetAllByParameterIDListOut.#.deleted_at,required
fullname=deprecated,src=GetAllByParameterIDListOut.#.deprecated,required
fullname=description,src=GetAllByParameterIDListOut.#.description,required
fullname=example,src=GetAllByParameterIDListOut.#.example,required
fullname=explode,src=GetAllByParameterIDListOut.#.explode,required
fullname=fullName,src=GetAllByParameterIDListOut.#.full_name,required
fullname=httpStatus,src=GetAllByParameterIDListOut.#.http_status,required
fullname=method,src=GetAllByParameterIDListOut.#.method,required
fullname=name,src=GetAllByParameterIDListOut.#.name,required
fullname=parameterId,src=GetAllByParameterIDListOut.#.parameter_id,required
fullname=position,src=GetAllByParameterIDListOut.#.position,required
fullname=required,src=GetAllByParameterIDListOut.#.required,required
fullname=serialize,src=GetAllByParameterIDListOut.#.serialize,required
fullname=serviceId,src=GetAllByParameterIDListOut.#.service_id,required
fullname=tag,src=GetAllByParameterIDListOut.#.tag,required
fullname=title,src=GetAllByParameterIDListOut.#.title,required
fullname=type,src=GetAllByParameterIDListOut.#.type,required
fullname=updatedAt,src=GetAllByParameterIDListOut.#.updated_at,required
fullname=validateSchemaId,src=GetAllByParameterIDListOut.#.validate_schema_id,required','{{define "input"}}
{{getSetValue . "ParameterIDList" "input.parameterIdList"}}
{{end}}','{{define "output"}}
{{getSetValue . "output.allowEmptyValue" "GetAllByParameterIDListOut.#.allow_empty_value"}}
{{getSetValue . "output.allowReserved" "GetAllByParameterIDListOut.#.allow_reserved"}}
{{getSetValue . "output.apiId" "GetAllByParameterIDListOut.#.api_id"}}
{{getSetValue . "output.createdAt" "GetAllByParameterIDListOut.#.created_at"}}
{{getSetValue . "output.deletedAt" "GetAllByParameterIDListOut.#.deleted_at"}}
{{getSetValue . "output.deprecated" "GetAllByParameterIDListOut.#.deprecated"}}
{{getSetValue . "output.description" "GetAllByParameterIDListOut.#.description"}}
{{getSetValue . "output.example" "GetAllByParameterIDListOut.#.example"}}
{{getSetValue . "output.explode" "GetAllByParameterIDListOut.#.explode"}}
{{getSetValue . "output.fullName" "GetAllByParameterIDListOut.#.full_name"}}
{{getSetValue . "output.httpStatus" "GetAllByParameterIDListOut.#.http_status"}}
{{getSetValue . "output.method" "GetAllByParameterIDListOut.#.method"}}
{{getSetValue . "output.name" "GetAllByParameterIDListOut.#.name"}}
{{getSetValue . "output.parameterId" "GetAllByParameterIDListOut.#.parameter_id"}}
{{getSetValue . "output.position" "GetAllByParameterIDListOut.#.position"}}
{{getSetValue . "output.required" "GetAllByParameterIDListOut.#.required"}}
{{getSetValue . "output.serialize" "GetAllByParameterIDListOut.#.serialize"}}
{{getSetValue . "output.serviceId" "GetAllByParameterIDListOut.#.service_id"}}
{{getSetValue . "output.tag" "GetAllByParameterIDListOut.#.tag"}}
{{getSetValue . "output.title" "GetAllByParameterIDListOut.#.title"}}
{{getSetValue . "output.type" "GetAllByParameterIDListOut.#.type"}}
{{getSetValue . "output.updatedAt" "GetAllByParameterIDListOut.#.updated_at"}}
{{getSetValue . "output.validateSchemaId" "GetAllByParameterIDListOut.#.validate_schema_id"}}
{{end}}');
insert ignore into `template` (`template_id`,`type`,`title`,`description`,`source_id`,`tpl`) values('DatabaseNameParameterPaginateWhere','SQL','database name 请求/响应参数 分页列表 where','database name 请求/响应参数 分页列表 where','database_name','{{define "PaginateWhere"}} {{end}}');
insert ignore into `template` (`template_id`,`type`,`title`,`description`,`source_id`,`tpl`) values('DatabaseNameParameterPaginateTotal','SQL','database name 请求/响应参数 分页列表 总数','database name 请求/响应参数 分页列表 总数','database_name','{{define "PaginateTotal"}} select count(*) as `count` from `t_parameter` where 1=1 {{template "PaginateWhere" .}} and `deleted_at` is null; {{end}}');
insert ignore into `api` (`api_id`,`title`,`description`,`method`,`route`,`template_ids`,`exec`,`input`,`output`,`input_tpl`,`output_tpl`) values('database_name-t_parameter-PaginateTotal','database name 请求/响应参数 分页列表 总数','database name 请求/响应参数 分页列表 总数','POST','/api/database_name/v1/parameter/paginate_total','DatabaseNameParameterPaginateTotal','{{define "main"}}
{{execSQLTpl . "PaginateTotal"}}
{{end}}','version=http://json-schema.org/draft-07/schema,id=input,direction=in','version=http://json-schema.org/draft-07/schema,id=output,direction=out','{{define "input"}}
{{end}}','{{define "output"}}
{{end}}');
insert ignore into `template` (`template_id`,`type`,`title`,`description`,`source_id`,`tpl`) values('DatabaseNameParameterPaginate','SQL','database name 请求/响应参数 分页列表','database name 请求/响应参数 分页列表','database_name','{{define "Paginate"}} select * from `t_parameter` where 1=1 {{template "PaginateWhere" .}} and `deleted_at` is null order by `updated_at` desc limit :Offset,:Limit ; {{end}}');
insert ignore into `api` (`api_id`,`title`,`description`,`method`,`route`,`template_ids`,`exec`,`input`,`output`,`input_tpl`,`output_tpl`) values('database_name-t_parameter-Paginate','database name 请求/响应参数 分页列表','database name 请求/响应参数 分页列表','POST','/api/database_name/v1/parameter/paginate','DatabaseNameParameterPaginateWhere,DatabaseNameParameterPaginateTotal,DatabaseNameParameterPaginate','{{define "main"}}
{{execSQLTpl . "PaginateTotal"}}
{{execSQLTpl . "Paginate"}}
{{end}}','version=http://json-schema.org/draft-07/schema,id=input,direction=in
fullname=limit,dst=Limit,format=number,required
fullname=offset,dst=Offset,format=number,required','version=http://json-schema.org/draft-07/schema,id=output,direction=out
fullname=allowEmptyValue,src=PaginateOut.#.allow_empty_value,required
fullname=allowReserved,src=PaginateOut.#.allow_reserved,required
fullname=apiId,src=PaginateOut.#.api_id,required
fullname=createdAt,src=PaginateOut.#.created_at,required
fullname=deletedAt,src=PaginateOut.#.deleted_at,required
fullname=deprecated,src=PaginateOut.#.deprecated,required
fullname=description,src=PaginateOut.#.description,required
fullname=example,src=PaginateOut.#.example,required
fullname=explode,src=PaginateOut.#.explode,required
fullname=fullName,src=PaginateOut.#.full_name,required
fullname=httpStatus,src=PaginateOut.#.http_status,required
fullname=method,src=PaginateOut.#.method,required
fullname=name,src=PaginateOut.#.name,required
fullname=parameterId,src=PaginateOut.#.parameter_id,required
fullname=position,src=PaginateOut.#.position,required
fullname=required,src=PaginateOut.#.required,required
fullname=serialize,src=PaginateOut.#.serialize,required
fullname=serviceId,src=PaginateOut.#.service_id,required
fullname=tag,src=PaginateOut.#.tag,required
fullname=title,src=PaginateOut.#.title,required
fullname=type,src=PaginateOut.#.type,required
fullname=updatedAt,src=PaginateOut.#.updated_at,required
fullname=validateSchemaId,src=PaginateOut.#.validate_schema_id,required','{{define "input"}}
{{getSetNumber . "Limit" "input.limit"}}
{{getSetNumber . "Offset" "input.offset"}}
{{end}}','{{define "output"}}
{{getSetValue . "output.allowEmptyValue" "PaginateOut.#.allow_empty_value"}}
{{getSetValue . "output.allowReserved" "PaginateOut.#.allow_reserved"}}
{{getSetValue . "output.apiId" "PaginateOut.#.api_id"}}
{{getSetValue . "output.createdAt" "PaginateOut.#.created_at"}}
{{getSetValue . "output.deletedAt" "PaginateOut.#.deleted_at"}}
{{getSetValue . "output.deprecated" "PaginateOut.#.deprecated"}}
{{getSetValue . "output.description" "PaginateOut.#.description"}}
{{getSetValue . "output.example" "PaginateOut.#.example"}}
{{getSetValue . "output.explode" "PaginateOut.#.explode"}}
{{getSetValue . "output.fullName" "PaginateOut.#.full_name"}}
{{getSetValue . "output.httpStatus" "PaginateOut.#.http_status"}}
{{getSetValue . "output.method" "PaginateOut.#.method"}}
{{getSetValue . "output.name" "PaginateOut.#.name"}}
{{getSetValue . "output.parameterId" "PaginateOut.#.parameter_id"}}
{{getSetValue . "output.position" "PaginateOut.#.position"}}
{{getSetValue . "output.required" "PaginateOut.#.required"}}
{{getSetValue . "output.serialize" "PaginateOut.#.serialize"}}
{{getSetValue . "output.serviceId" "PaginateOut.#.service_id"}}
{{getSetValue . "output.tag" "PaginateOut.#.tag"}}
{{getSetValue . "output.title" "PaginateOut.#.title"}}
{{getSetValue . "output.type" "PaginateOut.#.type"}}
{{getSetValue . "output.updatedAt" "PaginateOut.#.updated_at"}}
{{getSetValue . "output.validateSchemaId" "PaginateOut.#.validate_schema_id"}}
{{end}}');
insert ignore into `template` (`template_id`,`type`,`title`,`description`,`source_id`,`tpl`) values('DatabaseNameParameterInsert','SQL','database name 请求/响应参数 新增','database name 请求/响应参数 新增','database_name','{{define "Insert"}} insert into `t_parameter` (`parameter_id`,`service_id`,`api_id`,`validate_schema_id`,`full_name`,`name`,`title`,`type`,`tag`,`method`,`http_status`,`position`,`example`,`deprecated`,`required`,`serialize`,`explode`,`allow_empty_value`,`allow_reserved`,`description`)values (:ParameterID,:ServiceID,:APIID,:ValidateSchemaID,:FullName,:Name,:Title,:Type,:Tag,:Method,:HTTPStatus,:Position,:Example,:Deprecated,:Required,:Serialize,:Explode,:AllowEmptyValue,:AllowReserved,:Description); {{end}}');
insert ignore into `api` (`api_id`,`title`,`description`,`method`,`route`,`template_ids`,`exec`,`input`,`output`,`input_tpl`,`output_tpl`) values('database_name-t_parameter-Insert','database name 请求/响应参数 新增','database name 请求/响应参数 新增','POST','/api/database_name/v1/parameter/insert','DatabaseNameParameterInsert','{{define "main"}}
{{execSQLTpl . "Insert"}}
{{end}}','version=http://json-schema.org/draft-07/schema,id=input,direction=in
fullname=apiid,dst=APIID,required
fullname=allowEmptyValue,dst=AllowEmptyValue,required
fullname=allowReserved,dst=AllowReserved,required
fullname=deprecated,dst=Deprecated,required
fullname=description,dst=Description,required
fullname=example,dst=Example,required
fullname=explode,dst=Explode,required
fullname=fullName,dst=FullName,required
fullname=httpStatus,dst=HTTPStatus,required
fullname=method,dst=Method,required
fullname=name,dst=Name,required
fullname=parameterId,dst=ParameterID,required
fullname=position,dst=Position,required
fullname=required,dst=Required,required
fullname=serialize,dst=Serialize,required
fullname=serviceId,dst=ServiceID,required
fullname=tag,dst=Tag,required
fullname=title,dst=Title,required
fullname=type,dst=Type,required
fullname=validateSchemaId,dst=ValidateSchemaID,required','version=http://json-schema.org/draft-07/schema,id=output,direction=out','{{define "input"}}
{{getSetValue . "APIID" "input.apiid"}}
{{getSetValue . "AllowEmptyValue" "input.allowEmptyValue"}}
{{getSetValue . "AllowReserved" "input.allowReserved"}}
{{getSetValue . "Deprecated" "input.deprecated"}}
{{getSetValue . "Description" "input.description"}}
{{getSetValue . "Example" "input.example"}}
{{getSetValue . "Explode" "input.explode"}}
{{getSetValue . "FullName" "input.fullName"}}
{{getSetValue . "HTTPStatus" "input.httpStatus"}}
{{getSetValue . "Method" "input.method"}}
{{getSetValue . "Name" "input.name"}}
{{getSetValue . "ParameterID" "input.parameterId"}}
{{getSetValue . "Position" "input.position"}}
{{getSetValue . "Required" "input.required"}}
{{getSetValue . "Serialize" "input.serialize"}}
{{getSetValue . "ServiceID" "input.serviceId"}}
{{getSetValue . "Tag" "input.tag"}}
{{getSetValue . "Title" "input.title"}}
{{getSetValue . "Type" "input.type"}}
{{getSetValue . "ValidateSchemaID" "input.validateSchemaId"}}
{{end}}','{{define "output"}}
{{end}}');
insert ignore into `template` (`template_id`,`type`,`title`,`description`,`source_id`,`tpl`) values('DatabaseNameParameterUpdate','SQL','database name 请求/响应参数 修改','database name 请求/响应参数 修改','database_name','{{define "Update"}} {{$preComma:=newPreComma}} update `t_parameter` set {{if .ParameterID}} {{$preComma.PreComma}} `parameter_id`=:ParameterID {{end}} {{if .ServiceID}} {{$preComma.PreComma}} `service_id`=:ServiceID {{end}} {{if .APIID}} {{$preComma.PreComma}} `api_id`=:APIID {{end}} {{if .ValidateSchemaID}} {{$preComma.PreComma}} `validate_schema_id`=:ValidateSchemaID {{end}} {{if .FullName}} {{$preComma.PreComma}} `full_name`=:FullName {{end}} {{if .Name}} {{$preComma.PreComma}} `name`=:Name {{end}} {{if .Title}} {{$preComma.PreComma}} `title`=:Title {{end}} {{if .Type}} {{$preComma.PreComma}} `type`=:Type {{end}} {{if .Tag}} {{$preComma.PreComma}} `tag`=:Tag {{end}} {{if .Method}} {{$preComma.PreComma}} `method`=:Method {{end}} {{if .HTTPStatus}} {{$preComma.PreComma}} `http_status`=:HTTPStatus {{end}} {{if .Position}} {{$preComma.PreComma}} `position`=:Position {{end}} {{if .Example}} {{$preComma.PreComma}} `example`=:Example {{end}} {{if .Deprecated}} {{$preComma.PreComma}} `deprecated`=:Deprecated {{end}} {{if .Required}} {{$preComma.PreComma}} `required`=:Required {{end}} {{if .Serialize}} {{$preComma.PreComma}} `serialize`=:Serialize {{end}} {{if .Explode}} {{$preComma.PreComma}} `explode`=:Explode {{end}} {{if .AllowEmptyValue}} {{$preComma.PreComma}} `allow_empty_value`=:AllowEmptyValue {{end}} {{if .AllowReserved}} {{$preComma.PreComma}} `allow_reserved`=:AllowReserved {{end}} {{if .Description}} {{$preComma.PreComma}} `description`=:Description {{end}} where `parameter_id`=:ParameterID; {{end}}');
insert ignore into `api` (`api_id`,`title`,`description`,`method`,`route`,`template_ids`,`exec`,`input`,`output`,`input_tpl`,`output_tpl`) values('database_name-t_parameter-Update','database name 请求/响应参数 修改','database name 请求/响应参数 修改','POST','/api/database_name/v1/parameter/update','DatabaseNameParameterUpdate','{{define "main"}}
{{execSQLTpl . "Update"}}
{{end}}','version=http://json-schema.org/draft-07/schema,id=input,direction=in
fullname=apiid,dst=APIID,required
fullname=allowEmptyValue,dst=AllowEmptyValue,required
fullname=allowReserved,dst=AllowReserved,required
fullname=deprecated,dst=Deprecated,required
fullname=description,dst=Description,required
fullname=example,dst=Example,required
fullname=explode,dst=Explode,required
fullname=fullName,dst=FullName,required
fullname=httpStatus,dst=HTTPStatus,required
fullname=method,dst=Method,required
fullname=name,dst=Name,required
fullname=parameterId,dst=ParameterID,required
fullname=position,dst=Position,required
fullname=required,dst=Required,required
fullname=serialize,dst=Serialize,required
fullname=serviceId,dst=ServiceID,required
fullname=tag,dst=Tag,required
fullname=title,dst=Title,required
fullname=type,dst=Type,required
fullname=validateSchemaId,dst=ValidateSchemaID,required','version=http://json-schema.org/draft-07/schema,id=output,direction=out','{{define "input"}}
{{getSetValue . "APIID" "input.apiid"}}
{{getSetValue . "AllowEmptyValue" "input.allowEmptyValue"}}
{{getSetValue . "AllowReserved" "input.allowReserved"}}
{{getSetValue . "Deprecated" "input.deprecated"}}
{{getSetValue . "Description" "input.description"}}
{{getSetValue . "Example" "input.example"}}
{{getSetValue . "Explode" "input.explode"}}
{{getSetValue . "FullName" "input.fullName"}}
{{getSetValue . "HTTPStatus" "input.httpStatus"}}
{{getSetValue . "Method" "input.method"}}
{{getSetValue . "Name" "input.name"}}
{{getSetValue . "ParameterID" "input.parameterId"}}
{{getSetValue . "Position" "input.position"}}
{{getSetValue . "Required" "input.required"}}
{{getSetValue . "Serialize" "input.serialize"}}
{{getSetValue . "ServiceID" "input.serviceId"}}
{{getSetValue . "Tag" "input.tag"}}
{{getSetValue . "Title" "input.title"}}
{{getSetValue . "Type" "input.type"}}
{{getSetValue . "ValidateSchemaID" "input.validateSchemaId"}}
{{end}}','{{define "output"}}
{{end}}');
insert ignore into `template` (`template_id`,`type`,`title`,`description`,`source_id`,`tpl`) values('DatabaseNameParameterDel','SQL','database name 请求/响应参数 删除','database name 请求/响应参数 删除','database_name','{{define "Del"}} update `t_parameter` set `deleted_at`={{currentTime .}} where `parameter_id`=:ParameterID; {{end}}');
insert ignore into `api` (`api_id`,`title`,`description`,`method`,`route`,`template_ids`,`exec`,`input`,`output`,`input_tpl`,`output_tpl`) values('database_name-t_parameter-Del','database name 请求/响应参数 删除','database name 请求/响应参数 删除','POST','/api/database_name/v1/parameter/del','DatabaseNameParameterDel','{{define "main"}}
{{execSQLTpl . "Del"}}
{{end}}','version=http://json-schema.org/draft-07/schema,id=input,direction=in
fullname=parameterId,dst=ParameterID,required','version=http://json-schema.org/draft-07/schema,id=output,direction=out','{{define "input"}}
{{getSetValue . "ParameterID" "input.parameterId"}}
{{end}}','{{define "output"}}
{{end}}');
insert ignore into `template` (`template_id`,`type`,`title`,`description`,`source_id`,`tpl`) values('DatabaseNameValidateSchemaGetByValidateSchemaID','SQL','database name validate schema 获取 通过 ID','database name validate schema 获取 通过 ID','database_name','{{define "GetByValidateSchemaID"}} select * from `t_validate_schema` where `validate_schema_id`=:ValidateSchemaID and `deleted_at` is null; {{end}}');
insert ignore into `api` (`api_id`,`title`,`description`,`method`,`route`,`template_ids`,`exec`,`input`,`output`,`input_tpl`,`output_tpl`) values('database_name-t_validate_schema-GetByValidateSchemaID','database name validate schema 获取 通过 ID','database name validate schema 获取 通过 ID','POST','/api/database_name/v1/validate_schema/get_by_validate_schema_id','DatabaseNameValidateSchemaGetByValidateSchemaID','{{define "main"}}
{{execSQLTpl . "GetByValidateSchemaID"}}
{{end}}','version=http://json-schema.org/draft-07/schema,id=input,direction=in
fullname=validateSchemaId,dst=ValidateSchemaID,required','version=http://json-schema.org/draft-07/schema,id=output,direction=out
fullname=additionalProperties,src=GetByValidateSchemaIDOut.#.additional_properties,required
fullname=allOf,src=GetByValidateSchemaIDOut.#.all_of,required
fullname=allowEmptyValue,src=GetByValidateSchemaIDOut.#.allow_empty_value,required
fullname=allowReserved,src=GetByValidateSchemaIDOut.#.allow_reserved,required
fullname=anyOf,src=GetByValidateSchemaIDOut.#.any_of,required
fullname=createdAt,src=GetByValidateSchemaIDOut.#.created_at,required
fullname=default,src=GetByValidateSchemaIDOut.#.default,required
fullname=deletedAt,src=GetByValidateSchemaIDOut.#.deleted_at,required
fullname=deprecated,src=GetByValidateSchemaIDOut.#.deprecated,required
fullname=description,src=GetByValidateSchemaIDOut.#.description,required
fullname=discriminator,src=GetByValidateSchemaIDOut.#.discriminator,required
fullname=enum,src=GetByValidateSchemaIDOut.#.enum,required
fullname=enumNames,src=GetByValidateSchemaIDOut.#.enum_names,required
fullname=enumTitles,src=GetByValidateSchemaIDOut.#.enum_titles,required
fullname=example,src=GetByValidateSchemaIDOut.#.example,required
fullname=exclusiveMaximum,src=GetByValidateSchemaIDOut.#.exclusive_maximum,required
fullname=exclusiveMinimum,src=GetByValidateSchemaIDOut.#.exclusive_minimum,required
fullname=extensions,src=GetByValidateSchemaIDOut.#.extensions,required
fullname=externalDocs,src=GetByValidateSchemaIDOut.#.external_docs,required
fullname=externalPros,src=GetByValidateSchemaIDOut.#.external_pros,required
fullname=format,src=GetByValidateSchemaIDOut.#.format,required
fullname=maxItems,src=GetByValidateSchemaIDOut.#.max_items,required
fullname=maxLength,src=GetByValidateSchemaIDOut.#.max_length,required
fullname=maxProperties,src=GetByValidateSchemaIDOut.#.max_properties,required
fullname=maxnum,src=GetByValidateSchemaIDOut.#.maxnum,required
fullname=minItems,src=GetByValidateSchemaIDOut.#.min_items,required
fullname=minLength,src=GetByValidateSchemaIDOut.#.min_length,required
fullname=minProperties,src=GetByValidateSchemaIDOut.#.min_properties,required
fullname=minimum,src=GetByValidateSchemaIDOut.#.minimum,required
fullname=multipleOf,src=GetByValidateSchemaIDOut.#.multiple_of,required
fullname=not,src=GetByValidateSchemaIDOut.#.not,required
fullname=nullable,src=GetByValidateSchemaIDOut.#.nullable,required
fullname=oneOf,src=GetByValidateSchemaIDOut.#.one_of,required
fullname=pattern,src=GetByValidateSchemaIDOut.#.pattern,required
fullname=readOnly,src=GetByValidateSchemaIDOut.#.read_only,required
fullname=remark,src=GetByValidateSchemaIDOut.#.remark,required
fullname=required,src=GetByValidateSchemaIDOut.#.required,required
fullname=serviceId,src=GetByValidateSchemaIDOut.#.service_id,required
fullname=summary,src=GetByValidateSchemaIDOut.#.summary,required
fullname=type,src=GetByValidateSchemaIDOut.#.type,required
fullname=uniqueItems,src=GetByValidateSchemaIDOut.#.unique_items,required
fullname=updatedAt,src=GetByValidateSchemaIDOut.#.updated_at,required
fullname=validateSchemaId,src=GetByValidateSchemaIDOut.#.validate_schema_id,required
fullname=writeOnly,src=GetByValidateSchemaIDOut.#.write_only,required
fullname=xml,src=GetByValidateSchemaIDOut.#.xml,required','{{define "input"}}
{{getSetValue . "ValidateSchemaID" "input.validateSchemaId"}}
{{end}}','{{define "output"}}
{{getSetValue . "output.additionalProperties" "GetByValidateSchemaIDOut.#.additional_properties"}}
{{getSetValue . "output.allOf" "GetByValidateSchemaIDOut.#.all_of"}}
{{getSetValue . "output.allowEmptyValue" "GetByValidateSchemaIDOut.#.allow_empty_value"}}
{{getSetValue . "output.allowReserved" "GetByValidateSchemaIDOut.#.allow_reserved"}}
{{getSetValue . "output.anyOf" "GetByValidateSchemaIDOut.#.any_of"}}
{{getSetValue . "output.createdAt" "GetByValidateSchemaIDOut.#.created_at"}}
{{getSetValue . "output.default" "GetByValidateSchemaIDOut.#.default"}}
{{getSetValue . "output.deletedAt" "GetByValidateSchemaIDOut.#.deleted_at"}}
{{getSetValue . "output.deprecated" "GetByValidateSchemaIDOut.#.deprecated"}}
{{getSetValue . "output.description" "GetByValidateSchemaIDOut.#.description"}}
{{getSetValue . "output.discriminator" "GetByValidateSchemaIDOut.#.discriminator"}}
{{getSetValue . "output.enum" "GetByValidateSchemaIDOut.#.enum"}}
{{getSetValue . "output.enumNames" "GetByValidateSchemaIDOut.#.enum_names"}}
{{getSetValue . "output.enumTitles" "GetByValidateSchemaIDOut.#.enum_titles"}}
{{getSetValue . "output.example" "GetByValidateSchemaIDOut.#.example"}}
{{getSetValue . "output.exclusiveMaximum" "GetByValidateSchemaIDOut.#.exclusive_maximum"}}
{{getSetValue . "output.exclusiveMinimum" "GetByValidateSchemaIDOut.#.exclusive_minimum"}}
{{getSetValue . "output.extensions" "GetByValidateSchemaIDOut.#.extensions"}}
{{getSetValue . "output.externalDocs" "GetByValidateSchemaIDOut.#.external_docs"}}
{{getSetValue . "output.externalPros" "GetByValidateSchemaIDOut.#.external_pros"}}
{{getSetValue . "output.format" "GetByValidateSchemaIDOut.#.format"}}
{{getSetValue . "output.maxItems" "GetByValidateSchemaIDOut.#.max_items"}}
{{getSetValue . "output.maxLength" "GetByValidateSchemaIDOut.#.max_length"}}
{{getSetValue . "output.maxProperties" "GetByValidateSchemaIDOut.#.max_properties"}}
{{getSetValue . "output.maxnum" "GetByValidateSchemaIDOut.#.maxnum"}}
{{getSetValue . "output.minItems" "GetByValidateSchemaIDOut.#.min_items"}}
{{getSetValue . "output.minLength" "GetByValidateSchemaIDOut.#.min_length"}}
{{getSetValue . "output.minProperties" "GetByValidateSchemaIDOut.#.min_properties"}}
{{getSetValue . "output.minimum" "GetByValidateSchemaIDOut.#.minimum"}}
{{getSetValue . "output.multipleOf" "GetByValidateSchemaIDOut.#.multiple_of"}}
{{getSetValue . "output.not" "GetByValidateSchemaIDOut.#.not"}}
{{getSetValue . "output.nullable" "GetByValidateSchemaIDOut.#.nullable"}}
{{getSetValue . "output.oneOf" "GetByValidateSchemaIDOut.#.one_of"}}
{{getSetValue . "output.pattern" "GetByValidateSchemaIDOut.#.pattern"}}
{{getSetValue . "output.readOnly" "GetByValidateSchemaIDOut.#.read_only"}}
{{getSetValue . "output.remark" "GetByValidateSchemaIDOut.#.remark"}}
{{getSetValue . "output.required" "GetByValidateSchemaIDOut.#.required"}}
{{getSetValue . "output.serviceId" "GetByValidateSchemaIDOut.#.service_id"}}
{{getSetValue . "output.summary" "GetByValidateSchemaIDOut.#.summary"}}
{{getSetValue . "output.type" "GetByValidateSchemaIDOut.#.type"}}
{{getSetValue . "output.uniqueItems" "GetByValidateSchemaIDOut.#.unique_items"}}
{{getSetValue . "output.updatedAt" "GetByValidateSchemaIDOut.#.updated_at"}}
{{getSetValue . "output.validateSchemaId" "GetByValidateSchemaIDOut.#.validate_schema_id"}}
{{getSetValue . "output.writeOnly" "GetByValidateSchemaIDOut.#.write_only"}}
{{getSetValue . "output.xml" "GetByValidateSchemaIDOut.#.xml"}}
{{end}}');
insert ignore into `template` (`template_id`,`type`,`title`,`description`,`source_id`,`tpl`) values('DatabaseNameValidateSchemaGetAllByValidateSchemaIDList','SQL','database name validate schema 获取 所有 通过 ID 列表','database name validate schema 获取 所有 通过 ID 列表','database_name','{{define "GetAllByValidateSchemaIDList"}} select * from `t_validate_schema` where `validate_schema_id` in ({{in . .ValidateSchemaIDList}}) and `deleted_at` is null; {{end}}');
insert ignore into `api` (`api_id`,`title`,`description`,`method`,`route`,`template_ids`,`exec`,`input`,`output`,`input_tpl`,`output_tpl`) values('database_name-t_validate_schema-GetAllByValidateSchemaIDList','database name validate schema 获取 所有 通过 ID 列表','database name validate schema 获取 所有 通过 ID 列表','POST','/api/database_name/v1/validate_schema/get_all_by_validate_schema_id_list','DatabaseNameValidateSchemaGetAllByValidateSchemaIDList','{{define "main"}}
{{execSQLTpl . "GetAllByValidateSchemaIDList"}}
{{end}}','version=http://json-schema.org/draft-07/schema,id=input,direction=in
fullname=validateSchemaIdList,dst=ValidateSchemaIDList,required','version=http://json-schema.org/draft-07/schema,id=output,direction=out
fullname=additionalProperties,src=GetAllByValidateSchemaIDListOut.#.additional_properties,required
fullname=allOf,src=GetAllByValidateSchemaIDListOut.#.all_of,required
fullname=allowEmptyValue,src=GetAllByValidateSchemaIDListOut.#.allow_empty_value,required
fullname=allowReserved,src=GetAllByValidateSchemaIDListOut.#.allow_reserved,required
fullname=anyOf,src=GetAllByValidateSchemaIDListOut.#.any_of,required
fullname=createdAt,src=GetAllByValidateSchemaIDListOut.#.created_at,required
fullname=default,src=GetAllByValidateSchemaIDListOut.#.default,required
fullname=deletedAt,src=GetAllByValidateSchemaIDListOut.#.deleted_at,required
fullname=deprecated,src=GetAllByValidateSchemaIDListOut.#.deprecated,required
fullname=description,src=GetAllByValidateSchemaIDListOut.#.description,required
fullname=discriminator,src=GetAllByValidateSchemaIDListOut.#.discriminator,required
fullname=enum,src=GetAllByValidateSchemaIDListOut.#.enum,required
fullname=enumNames,src=GetAllByValidateSchemaIDListOut.#.enum_names,required
fullname=enumTitles,src=GetAllByValidateSchemaIDListOut.#.enum_titles,required
fullname=example,src=GetAllByValidateSchemaIDListOut.#.example,required
fullname=exclusiveMaximum,src=GetAllByValidateSchemaIDListOut.#.exclusive_maximum,required
fullname=exclusiveMinimum,src=GetAllByValidateSchemaIDListOut.#.exclusive_minimum,required
fullname=extensions,src=GetAllByValidateSchemaIDListOut.#.extensions,required
fullname=externalDocs,src=GetAllByValidateSchemaIDListOut.#.external_docs,required
fullname=externalPros,src=GetAllByValidateSchemaIDListOut.#.external_pros,required
fullname=format,src=GetAllByValidateSchemaIDListOut.#.format,required
fullname=maxItems,src=GetAllByValidateSchemaIDListOut.#.max_items,required
fullname=maxLength,src=GetAllByValidateSchemaIDListOut.#.max_length,required
fullname=maxProperties,src=GetAllByValidateSchemaIDListOut.#.max_properties,required
fullname=maxnum,src=GetAllByValidateSchemaIDListOut.#.maxnum,required
fullname=minItems,src=GetAllByValidateSchemaIDListOut.#.min_items,required
fullname=minLength,src=GetAllByValidateSchemaIDListOut.#.min_length,required
fullname=minProperties,src=GetAllByValidateSchemaIDListOut.#.min_properties,required
fullname=minimum,src=GetAllByValidateSchemaIDListOut.#.minimum,required
fullname=multipleOf,src=GetAllByValidateSchemaIDListOut.#.multiple_of,required
fullname=not,src=GetAllByValidateSchemaIDListOut.#.not,required
fullname=nullable,src=GetAllByValidateSchemaIDListOut.#.nullable,required
fullname=oneOf,src=GetAllByValidateSchemaIDListOut.#.one_of,required
fullname=pattern,src=GetAllByValidateSchemaIDListOut.#.pattern,required
fullname=readOnly,src=GetAllByValidateSchemaIDListOut.#.read_only,required
fullname=remark,src=GetAllByValidateSchemaIDListOut.#.remark,required
fullname=required,src=GetAllByValidateSchemaIDListOut.#.required,required
fullname=serviceId,src=GetAllByValidateSchemaIDListOut.#.service_id,required
fullname=summary,src=GetAllByValidateSchemaIDListOut.#.summary,required
fullname=type,src=GetAllByValidateSchemaIDListOut.#.type,required
fullname=uniqueItems,src=GetAllByValidateSchemaIDListOut.#.unique_items,required
fullname=updatedAt,src=GetAllByValidateSchemaIDListOut.#.updated_at,required
fullname=validateSchemaId,src=GetAllByValidateSchemaIDListOut.#.validate_schema_id,required
fullname=writeOnly,src=GetAllByValidateSchemaIDListOut.#.write_only,required
fullname=xml,src=GetAllByValidateSchemaIDListOut.#.xml,required','{{define "input"}}
{{getSetValue . "ValidateSchemaIDList" "input.validateSchemaIdList"}}
{{end}}','{{define "output"}}
{{getSetValue . "output.additionalProperties" "GetAllByValidateSchemaIDListOut.#.additional_properties"}}
{{getSetValue . "output.allOf" "GetAllByValidateSchemaIDListOut.#.all_of"}}
{{getSetValue . "output.allowEmptyValue" "GetAllByValidateSchemaIDListOut.#.allow_empty_value"}}
{{getSetValue . "output.allowReserved" "GetAllByValidateSchemaIDListOut.#.allow_reserved"}}
{{getSetValue . "output.anyOf" "GetAllByValidateSchemaIDListOut.#.any_of"}}
{{getSetValue . "output.createdAt" "GetAllByValidateSchemaIDListOut.#.created_at"}}
{{getSetValue . "output.default" "GetAllByValidateSchemaIDListOut.#.default"}}
{{getSetValue . "output.deletedAt" "GetAllByValidateSchemaIDListOut.#.deleted_at"}}
{{getSetValue . "output.deprecated" "GetAllByValidateSchemaIDListOut.#.deprecated"}}
{{getSetValue . "output.description" "GetAllByValidateSchemaIDListOut.#.description"}}
{{getSetValue . "output.discriminator" "GetAllByValidateSchemaIDListOut.#.discriminator"}}
{{getSetValue . "output.enum" "GetAllByValidateSchemaIDListOut.#.enum"}}
{{getSetValue . "output.enumNames" "GetAllByValidateSchemaIDListOut.#.enum_names"}}
{{getSetValue . "output.enumTitles" "GetAllByValidateSchemaIDListOut.#.enum_titles"}}
{{getSetValue . "output.example" "GetAllByValidateSchemaIDListOut.#.example"}}
{{getSetValue . "output.exclusiveMaximum" "GetAllByValidateSchemaIDListOut.#.exclusive_maximum"}}
{{getSetValue . "output.exclusiveMinimum" "GetAllByValidateSchemaIDListOut.#.exclusive_minimum"}}
{{getSetValue . "output.extensions" "GetAllByValidateSchemaIDListOut.#.extensions"}}
{{getSetValue . "output.externalDocs" "GetAllByValidateSchemaIDListOut.#.external_docs"}}
{{getSetValue . "output.externalPros" "GetAllByValidateSchemaIDListOut.#.external_pros"}}
{{getSetValue . "output.format" "GetAllByValidateSchemaIDListOut.#.format"}}
{{getSetValue . "output.maxItems" "GetAllByValidateSchemaIDListOut.#.max_items"}}
{{getSetValue . "output.maxLength" "GetAllByValidateSchemaIDListOut.#.max_length"}}
{{getSetValue . "output.maxProperties" "GetAllByValidateSchemaIDListOut.#.max_properties"}}
{{getSetValue . "output.maxnum" "GetAllByValidateSchemaIDListOut.#.maxnum"}}
{{getSetValue . "output.minItems" "GetAllByValidateSchemaIDListOut.#.min_items"}}
{{getSetValue . "output.minLength" "GetAllByValidateSchemaIDListOut.#.min_length"}}
{{getSetValue . "output.minProperties" "GetAllByValidateSchemaIDListOut.#.min_properties"}}
{{getSetValue . "output.minimum" "GetAllByValidateSchemaIDListOut.#.minimum"}}
{{getSetValue . "output.multipleOf" "GetAllByValidateSchemaIDListOut.#.multiple_of"}}
{{getSetValue . "output.not" "GetAllByValidateSchemaIDListOut.#.not"}}
{{getSetValue . "output.nullable" "GetAllByValidateSchemaIDListOut.#.nullable"}}
{{getSetValue . "output.oneOf" "GetAllByValidateSchemaIDListOut.#.one_of"}}
{{getSetValue . "output.pattern" "GetAllByValidateSchemaIDListOut.#.pattern"}}
{{getSetValue . "output.readOnly" "GetAllByValidateSchemaIDListOut.#.read_only"}}
{{getSetValue . "output.remark" "GetAllByValidateSchemaIDListOut.#.remark"}}
{{getSetValue . "output.required" "GetAllByValidateSchemaIDListOut.#.required"}}
{{getSetValue . "output.serviceId" "GetAllByValidateSchemaIDListOut.#.service_id"}}
{{getSetValue . "output.summary" "GetAllByValidateSchemaIDListOut.#.summary"}}
{{getSetValue . "output.type" "GetAllByValidateSchemaIDListOut.#.type"}}
{{getSetValue . "output.uniqueItems" "GetAllByValidateSchemaIDListOut.#.unique_items"}}
{{getSetValue . "output.updatedAt" "GetAllByValidateSchemaIDListOut.#.updated_at"}}
{{getSetValue . "output.validateSchemaId" "GetAllByValidateSchemaIDListOut.#.validate_schema_id"}}
{{getSetValue . "output.writeOnly" "GetAllByValidateSchemaIDListOut.#.write_only"}}
{{getSetValue . "output.xml" "GetAllByValidateSchemaIDListOut.#.xml"}}
{{end}}');
insert ignore into `template` (`template_id`,`type`,`title`,`description`,`source_id`,`tpl`) values('DatabaseNameValidateSchemaPaginateWhere','SQL','database name validate schema 分页列表 where','database name validate schema 分页列表 where','database_name','{{define "PaginateWhere"}} {{end}}');
insert ignore into `template` (`template_id`,`type`,`title`,`description`,`source_id`,`tpl`) values('DatabaseNameValidateSchemaPaginateTotal','SQL','database name validate schema 分页列表 总数','database name validate schema 分页列表 总数','database_name','{{define "PaginateTotal"}} select count(*) as `count` from `t_validate_schema` where 1=1 {{template "PaginateWhere" .}} and `deleted_at` is null; {{end}}');
insert ignore into `api` (`api_id`,`title`,`description`,`method`,`route`,`template_ids`,`exec`,`input`,`output`,`input_tpl`,`output_tpl`) values('database_name-t_validate_schema-PaginateTotal','database name validate schema 分页列表 总数','database name validate schema 分页列表 总数','POST','/api/database_name/v1/validate_schema/paginate_total','DatabaseNameValidateSchemaPaginateTotal','{{define "main"}}
{{execSQLTpl . "PaginateTotal"}}
{{end}}','version=http://json-schema.org/draft-07/schema,id=input,direction=in','version=http://json-schema.org/draft-07/schema,id=output,direction=out','{{define "input"}}
{{end}}','{{define "output"}}
{{end}}');
insert ignore into `template` (`template_id`,`type`,`title`,`description`,`source_id`,`tpl`) values('DatabaseNameValidateSchemaPaginate','SQL','database name validate schema 分页列表','database name validate schema 分页列表','database_name','{{define "Paginate"}} select * from `t_validate_schema` where 1=1 {{template "PaginateWhere" .}} and `deleted_at` is null order by `updated_at` desc limit :Offset,:Limit ; {{end}}');
insert ignore into `api` (`api_id`,`title`,`description`,`method`,`route`,`template_ids`,`exec`,`input`,`output`,`input_tpl`,`output_tpl`) values('database_name-t_validate_schema-Paginate','database name validate schema 分页列表','database name validate schema 分页列表','POST','/api/database_name/v1/validate_schema/paginate','DatabaseNameValidateSchemaPaginateWhere,DatabaseNameValidateSchemaPaginateTotal,DatabaseNameValidateSchemaPaginate','{{define "main"}}
{{execSQLTpl . "PaginateTotal"}}
{{execSQLTpl . "Paginate"}}
{{end}}','version=http://json-schema.org/draft-07/schema,id=input,direction=in
fullname=limit,dst=Limit,format=number,required
fullname=offset,dst=Offset,format=number,required','version=http://json-schema.org/draft-07/schema,id=output,direction=out
fullname=additionalProperties,src=PaginateOut.#.additional_properties,required
fullname=allOf,src=PaginateOut.#.all_of,required
fullname=allowEmptyValue,src=PaginateOut.#.allow_empty_value,required
fullname=allowReserved,src=PaginateOut.#.allow_reserved,required
fullname=anyOf,src=PaginateOut.#.any_of,required
fullname=createdAt,src=PaginateOut.#.created_at,required
fullname=default,src=PaginateOut.#.default,required
fullname=deletedAt,src=PaginateOut.#.deleted_at,required
fullname=deprecated,src=PaginateOut.#.deprecated,required
fullname=description,src=PaginateOut.#.description,required
fullname=discriminator,src=PaginateOut.#.discriminator,required
fullname=enum,src=PaginateOut.#.enum,required
fullname=enumNames,src=PaginateOut.#.enum_names,required
fullname=enumTitles,src=PaginateOut.#.enum_titles,required
fullname=example,src=PaginateOut.#.example,required
fullname=exclusiveMaximum,src=PaginateOut.#.exclusive_maximum,required
fullname=exclusiveMinimum,src=PaginateOut.#.exclusive_minimum,required
fullname=extensions,src=PaginateOut.#.extensions,required
fullname=externalDocs,src=PaginateOut.#.external_docs,required
fullname=externalPros,src=PaginateOut.#.external_pros,required
fullname=format,src=PaginateOut.#.format,required
fullname=maxItems,src=PaginateOut.#.max_items,required
fullname=maxLength,src=PaginateOut.#.max_length,required
fullname=maxProperties,src=PaginateOut.#.max_properties,required
fullname=maxnum,src=PaginateOut.#.maxnum,required
fullname=minItems,src=PaginateOut.#.min_items,required
fullname=minLength,src=PaginateOut.#.min_length,required
fullname=minProperties,src=PaginateOut.#.min_properties,required
fullname=minimum,src=PaginateOut.#.minimum,required
fullname=multipleOf,src=PaginateOut.#.multiple_of,required
fullname=not,src=PaginateOut.#.not,required
fullname=nullable,src=PaginateOut.#.nullable,required
fullname=oneOf,src=PaginateOut.#.one_of,required
fullname=pattern,src=PaginateOut.#.pattern,required
fullname=readOnly,src=PaginateOut.#.read_only,required
fullname=remark,src=PaginateOut.#.remark,required
fullname=required,src=PaginateOut.#.required,required
fullname=serviceId,src=PaginateOut.#.service_id,required
fullname=summary,src=PaginateOut.#.summary,required
fullname=type,src=PaginateOut.#.type,required
fullname=uniqueItems,src=PaginateOut.#.unique_items,required
fullname=updatedAt,src=PaginateOut.#.updated_at,required
fullname=validateSchemaId,src=PaginateOut.#.validate_schema_id,required
fullname=writeOnly,src=PaginateOut.#.write_only,required
fullname=xml,src=PaginateOut.#.xml,required','{{define "input"}}
{{getSetNumber . "Limit" "input.limit"}}
{{getSetNumber . "Offset" "input.offset"}}
{{end}}','{{define "output"}}
{{getSetValue . "output.additionalProperties" "PaginateOut.#.additional_properties"}}
{{getSetValue . "output.allOf" "PaginateOut.#.all_of"}}
{{getSetValue . "output.allowEmptyValue" "PaginateOut.#.allow_empty_value"}}
{{getSetValue . "output.allowReserved" "PaginateOut.#.allow_reserved"}}
{{getSetValue . "output.anyOf" "PaginateOut.#.any_of"}}
{{getSetValue . "output.createdAt" "PaginateOut.#.created_at"}}
{{getSetValue . "output.default" "PaginateOut.#.default"}}
{{getSetValue . "output.deletedAt" "PaginateOut.#.deleted_at"}}
{{getSetValue . "output.deprecated" "PaginateOut.#.deprecated"}}
{{getSetValue . "output.description" "PaginateOut.#.description"}}
{{getSetValue . "output.discriminator" "PaginateOut.#.discriminator"}}
{{getSetValue . "output.enum" "PaginateOut.#.enum"}}
{{getSetValue . "output.enumNames" "PaginateOut.#.enum_names"}}
{{getSetValue . "output.enumTitles" "PaginateOut.#.enum_titles"}}
{{getSetValue . "output.example" "PaginateOut.#.example"}}
{{getSetValue . "output.exclusiveMaximum" "PaginateOut.#.exclusive_maximum"}}
{{getSetValue . "output.exclusiveMinimum" "PaginateOut.#.exclusive_minimum"}}
{{getSetValue . "output.extensions" "PaginateOut.#.extensions"}}
{{getSetValue . "output.externalDocs" "PaginateOut.#.external_docs"}}
{{getSetValue . "output.externalPros" "PaginateOut.#.external_pros"}}
{{getSetValue . "output.format" "PaginateOut.#.format"}}
{{getSetValue . "output.maxItems" "PaginateOut.#.max_items"}}
{{getSetValue . "output.maxLength" "PaginateOut.#.max_length"}}
{{getSetValue . "output.maxProperties" "PaginateOut.#.max_properties"}}
{{getSetValue . "output.maxnum" "PaginateOut.#.maxnum"}}
{{getSetValue . "output.minItems" "PaginateOut.#.min_items"}}
{{getSetValue . "output.minLength" "PaginateOut.#.min_length"}}
{{getSetValue . "output.minProperties" "PaginateOut.#.min_properties"}}
{{getSetValue . "output.minimum" "PaginateOut.#.minimum"}}
{{getSetValue . "output.multipleOf" "PaginateOut.#.multiple_of"}}
{{getSetValue . "output.not" "PaginateOut.#.not"}}
{{getSetValue . "output.nullable" "PaginateOut.#.nullable"}}
{{getSetValue . "output.oneOf" "PaginateOut.#.one_of"}}
{{getSetValue . "output.pattern" "PaginateOut.#.pattern"}}
{{getSetValue . "output.readOnly" "PaginateOut.#.read_only"}}
{{getSetValue . "output.remark" "PaginateOut.#.remark"}}
{{getSetValue . "output.required" "PaginateOut.#.required"}}
{{getSetValue . "output.serviceId" "PaginateOut.#.service_id"}}
{{getSetValue . "output.summary" "PaginateOut.#.summary"}}
{{getSetValue . "output.type" "PaginateOut.#.type"}}
{{getSetValue . "output.uniqueItems" "PaginateOut.#.unique_items"}}
{{getSetValue . "output.updatedAt" "PaginateOut.#.updated_at"}}
{{getSetValue . "output.validateSchemaId" "PaginateOut.#.validate_schema_id"}}
{{getSetValue . "output.writeOnly" "PaginateOut.#.write_only"}}
{{getSetValue . "output.xml" "PaginateOut.#.xml"}}
{{end}}');
insert ignore into `template` (`template_id`,`type`,`title`,`description`,`source_id`,`tpl`) values('DatabaseNameValidateSchemaInsert','SQL','database name validate schema 新增','database name validate schema 新增','database_name','{{define "Insert"}} insert into `t_validate_schema` (`validate_schema_id`,`service_id`,`description`,`remark`,`type`,`example`,`deprecated`,`required`,`enum`,`enum_names`,`enum_titles`,`format`,`default`,`nullable`,`multiple_of`,`maxnum`,`exclusive_maximum`,`minimum`,`exclusive_minimum`,`max_length`,`min_length`,`pattern`,`max_items`,`min_items`,`unique_items`,`max_properties`,`min_properties`,`all_of`,`one_of`,`any_of`,`allow_empty_value`,`allow_reserved`,`not`,`additional_properties`,`discriminator`,`read_only`,`write_only`,`xml`,`external_docs`,`external_pros`,`extensions`,`summary`)values (:ValidateSchemaID,:ServiceID,:Description,:Remark,:Type,:Example,:Deprecated,:Required,:Enum,:EnumNames,:EnumTitles,:Format,:Default,:Nullable,:MultipleOf,:Maxnum,:ExclusiveMaximum,:Minimum,:ExclusiveMinimum,:MaxLength,:MinLength,:Pattern,:MaxItems,:MinItems,:UniqueItems,:MaxProperties,:MinProperties,:AllOf,:OneOf,:AnyOf,:AllowEmptyValue,:AllowReserved,:Not,:AdditionalProperties,:Discriminator,:ReadOnly,:WriteOnly,:XML,:ExternalDocs,:ExternalPros,:Extensions,:Summary); {{end}}');
insert ignore into `api` (`api_id`,`title`,`description`,`method`,`route`,`template_ids`,`exec`,`input`,`output`,`input_tpl`,`output_tpl`) values('database_name-t_validate_schema-Insert','database name validate schema 新增','database name validate schema 新增','POST','/api/database_name/v1/validate_schema/insert','DatabaseNameValidateSchemaInsert','{{define "main"}}
{{execSQLTpl . "Insert"}}
{{end}}','version=http://json-schema.org/draft-07/schema,id=input,direction=in
fullname=additionalProperties,dst=AdditionalProperties,required
fullname=allOf,dst=AllOf,required
fullname=allowEmptyValue,dst=AllowEmptyValue,required
fullname=allowReserved,dst=AllowReserved,required
fullname=anyOf,dst=AnyOf,required
fullname=default,dst=Default,required
fullname=deprecated,dst=Deprecated,required
fullname=description,dst=Description,required
fullname=discriminator,dst=Discriminator,required
fullname=enum,dst=Enum,required
fullname=enumNames,dst=EnumNames,required
fullname=enumTitles,dst=EnumTitles,required
fullname=example,dst=Example,required
fullname=exclusiveMaximum,dst=ExclusiveMaximum,required
fullname=exclusiveMinimum,dst=ExclusiveMinimum,required
fullname=extensions,dst=Extensions,required
fullname=externalDocs,dst=ExternalDocs,required
fullname=externalPros,dst=ExternalPros,required
fullname=format,dst=Format,required
fullname=maxItems,dst=MaxItems,format=number,required
fullname=maxLength,dst=MaxLength,format=number,required
fullname=maxProperties,dst=MaxProperties,format=number,required
fullname=maxnum,dst=Maxnum,format=number,required
fullname=minItems,dst=MinItems,format=number,required
fullname=minLength,dst=MinLength,format=number,required
fullname=minProperties,dst=MinProperties,format=number,required
fullname=minimum,dst=Minimum,format=number,required
fullname=multipleOf,dst=MultipleOf,format=number,required
fullname=not,dst=Not,required
fullname=nullable,dst=Nullable,required
fullname=oneOf,dst=OneOf,required
fullname=pattern,dst=Pattern,required
fullname=readOnly,dst=ReadOnly,required
fullname=remark,dst=Remark,required
fullname=required,dst=Required,required
fullname=serviceId,dst=ServiceID,required
fullname=summary,dst=Summary,required
fullname=type,dst=Type,required
fullname=uniqueItems,dst=UniqueItems,required
fullname=validateSchemaId,dst=ValidateSchemaID,required
fullname=writeOnly,dst=WriteOnly,required
fullname=xml,dst=XML,required','version=http://json-schema.org/draft-07/schema,id=output,direction=out','{{define "input"}}
{{getSetValue . "AdditionalProperties" "input.additionalProperties"}}
{{getSetValue . "AllOf" "input.allOf"}}
{{getSetValue . "AllowEmptyValue" "input.allowEmptyValue"}}
{{getSetValue . "AllowReserved" "input.allowReserved"}}
{{getSetValue . "AnyOf" "input.anyOf"}}
{{getSetValue . "Default" "input.default"}}
{{getSetValue . "Deprecated" "input.deprecated"}}
{{getSetValue . "Description" "input.description"}}
{{getSetValue . "Discriminator" "input.discriminator"}}
{{getSetValue . "Enum" "input.enum"}}
{{getSetValue . "EnumNames" "input.enumNames"}}
{{getSetValue . "EnumTitles" "input.enumTitles"}}
{{getSetValue . "Example" "input.example"}}
{{getSetValue . "ExclusiveMaximum" "input.exclusiveMaximum"}}
{{getSetValue . "ExclusiveMinimum" "input.exclusiveMinimum"}}
{{getSetValue . "Extensions" "input.extensions"}}
{{getSetValue . "ExternalDocs" "input.externalDocs"}}
{{getSetValue . "ExternalPros" "input.externalPros"}}
{{getSetValue . "Format" "input.format"}}
{{getSetNumber . "MaxItems" "input.maxItems"}}
{{getSetNumber . "MaxLength" "input.maxLength"}}
{{getSetNumber . "MaxProperties" "input.maxProperties"}}
{{getSetNumber . "Maxnum" "input.maxnum"}}
{{getSetNumber . "MinItems" "input.minItems"}}
{{getSetNumber . "MinLength" "input.minLength"}}
{{getSetNumber . "MinProperties" "input.minProperties"}}
{{getSetNumber . "Minimum" "input.minimum"}}
{{getSetNumber . "MultipleOf" "input.multipleOf"}}
{{getSetValue . "Not" "input.not"}}
{{getSetValue . "Nullable" "input.nullable"}}
{{getSetValue . "OneOf" "input.oneOf"}}
{{getSetValue . "Pattern" "input.pattern"}}
{{getSetValue . "ReadOnly" "input.readOnly"}}
{{getSetValue . "Remark" "input.remark"}}
{{getSetValue . "Required" "input.required"}}
{{getSetValue . "ServiceID" "input.serviceId"}}
{{getSetValue . "Summary" "input.summary"}}
{{getSetValue . "Type" "input.type"}}
{{getSetValue . "UniqueItems" "input.uniqueItems"}}
{{getSetValue . "ValidateSchemaID" "input.validateSchemaId"}}
{{getSetValue . "WriteOnly" "input.writeOnly"}}
{{getSetValue . "XML" "input.xml"}}
{{end}}','{{define "output"}}
{{end}}');
insert ignore into `template` (`template_id`,`type`,`title`,`description`,`source_id`,`tpl`) values('DatabaseNameValidateSchemaUpdate','SQL','database name validate schema 修改','database name validate schema 修改','database_name','{{define "Update"}} {{$preComma:=newPreComma}} update `t_validate_schema` set {{if .ValidateSchemaID}} {{$preComma.PreComma}} `validate_schema_id`=:ValidateSchemaID {{end}} {{if .ServiceID}} {{$preComma.PreComma}} `service_id`=:ServiceID {{end}} {{if .Description}} {{$preComma.PreComma}} `description`=:Description {{end}} {{if .Remark}} {{$preComma.PreComma}} `remark`=:Remark {{end}} {{if .Type}} {{$preComma.PreComma}} `type`=:Type {{end}} {{if .Example}} {{$preComma.PreComma}} `example`=:Example {{end}} {{if .Deprecated}} {{$preComma.PreComma}} `deprecated`=:Deprecated {{end}} {{if .Required}} {{$preComma.PreComma}} `required`=:Required {{end}} {{if .Enum}} {{$preComma.PreComma}} `enum`=:Enum {{end}} {{if .EnumNames}} {{$preComma.PreComma}} `enum_names`=:EnumNames {{end}} {{if .EnumTitles}} {{$preComma.PreComma}} `enum_titles`=:EnumTitles {{end}} {{if .Format}} {{$preComma.PreComma}} `format`=:Format {{end}} {{if .Default}} {{$preComma.PreComma}} `default`=:Default {{end}} {{if .Nullable}} {{$preComma.PreComma}} `nullable`=:Nullable {{end}} {{if .MultipleOf}} {{$preComma.PreComma}} `multiple_of`=:MultipleOf {{end}} {{if .Maxnum}} {{$preComma.PreComma}} `maxnum`=:Maxnum {{end}} {{if .ExclusiveMaximum}} {{$preComma.PreComma}} `exclusive_maximum`=:ExclusiveMaximum {{end}} {{if .Minimum}} {{$preComma.PreComma}} `minimum`=:Minimum {{end}} {{if .ExclusiveMinimum}} {{$preComma.PreComma}} `exclusive_minimum`=:ExclusiveMinimum {{end}} {{if .MaxLength}} {{$preComma.PreComma}} `max_length`=:MaxLength {{end}} {{if .MinLength}} {{$preComma.PreComma}} `min_length`=:MinLength {{end}} {{if .Pattern}} {{$preComma.PreComma}} `pattern`=:Pattern {{end}} {{if .MaxItems}} {{$preComma.PreComma}} `max_items`=:MaxItems {{end}} {{if .MinItems}} {{$preComma.PreComma}} `min_items`=:MinItems {{end}} {{if .UniqueItems}} {{$preComma.PreComma}} `unique_items`=:UniqueItems {{end}} {{if .MaxProperties}} {{$preComma.PreComma}} `max_properties`=:MaxProperties {{end}} {{if .MinProperties}} {{$preComma.PreComma}} `min_properties`=:MinProperties {{end}} {{if .AllOf}} {{$preComma.PreComma}} `all_of`=:AllOf {{end}} {{if .OneOf}} {{$preComma.PreComma}} `one_of`=:OneOf {{end}} {{if .AnyOf}} {{$preComma.PreComma}} `any_of`=:AnyOf {{end}} {{if .AllowEmptyValue}} {{$preComma.PreComma}} `allow_empty_value`=:AllowEmptyValue {{end}} {{if .AllowReserved}} {{$preComma.PreComma}} `allow_reserved`=:AllowReserved {{end}} {{if .Not}} {{$preComma.PreComma}} `not`=:Not {{end}} {{if .AdditionalProperties}} {{$preComma.PreComma}} `additional_properties`=:AdditionalProperties {{end}} {{if .Discriminator}} {{$preComma.PreComma}} `discriminator`=:Discriminator {{end}} {{if .ReadOnly}} {{$preComma.PreComma}} `read_only`=:ReadOnly {{end}} {{if .WriteOnly}} {{$preComma.PreComma}} `write_only`=:WriteOnly {{end}} {{if .XML}} {{$preComma.PreComma}} `xml`=:XML {{end}} {{if .ExternalDocs}} {{$preComma.PreComma}} `external_docs`=:ExternalDocs {{end}} {{if .ExternalPros}} {{$preComma.PreComma}} `external_pros`=:ExternalPros {{end}} {{if .Extensions}} {{$preComma.PreComma}} `extensions`=:Extensions {{end}} {{if .Summary}} {{$preComma.PreComma}} `summary`=:Summary {{end}} where `validate_schema_id`=:ValidateSchemaID; {{end}}');
insert ignore into `api` (`api_id`,`title`,`description`,`method`,`route`,`template_ids`,`exec`,`input`,`output`,`input_tpl`,`output_tpl`) values('database_name-t_validate_schema-Update','database name validate schema 修改','database name validate schema 修改','POST','/api/database_name/v1/validate_schema/update','DatabaseNameValidateSchemaUpdate','{{define "main"}}
{{execSQLTpl . "Update"}}
{{end}}','version=http://json-schema.org/draft-07/schema,id=input,direction=in
fullname=additionalProperties,dst=AdditionalProperties,required
fullname=allOf,dst=AllOf,required
fullname=allowEmptyValue,dst=AllowEmptyValue,required
fullname=allowReserved,dst=AllowReserved,required
fullname=anyOf,dst=AnyOf,required
fullname=default,dst=Default,required
fullname=deprecated,dst=Deprecated,required
fullname=description,dst=Description,required
fullname=discriminator,dst=Discriminator,required
fullname=enum,dst=Enum,required
fullname=enumNames,dst=EnumNames,required
fullname=enumTitles,dst=EnumTitles,required
fullname=example,dst=Example,required
fullname=exclusiveMaximum,dst=ExclusiveMaximum,required
fullname=exclusiveMinimum,dst=ExclusiveMinimum,required
fullname=extensions,dst=Extensions,required
fullname=externalDocs,dst=ExternalDocs,required
fullname=externalPros,dst=ExternalPros,required
fullname=format,dst=Format,required
fullname=maxItems,dst=MaxItems,format=number,required
fullname=maxLength,dst=MaxLength,format=number,required
fullname=maxProperties,dst=MaxProperties,format=number,required
fullname=maxnum,dst=Maxnum,format=number,required
fullname=minItems,dst=MinItems,format=number,required
fullname=minLength,dst=MinLength,format=number,required
fullname=minProperties,dst=MinProperties,format=number,required
fullname=minimum,dst=Minimum,format=number,required
fullname=multipleOf,dst=MultipleOf,format=number,required
fullname=not,dst=Not,required
fullname=nullable,dst=Nullable,required
fullname=oneOf,dst=OneOf,required
fullname=pattern,dst=Pattern,required
fullname=readOnly,dst=ReadOnly,required
fullname=remark,dst=Remark,required
fullname=required,dst=Required,required
fullname=serviceId,dst=ServiceID,required
fullname=summary,dst=Summary,required
fullname=type,dst=Type,required
fullname=uniqueItems,dst=UniqueItems,required
fullname=validateSchemaId,dst=ValidateSchemaID,required
fullname=writeOnly,dst=WriteOnly,required
fullname=xml,dst=XML,required','version=http://json-schema.org/draft-07/schema,id=output,direction=out','{{define "input"}}
{{getSetValue . "AdditionalProperties" "input.additionalProperties"}}
{{getSetValue . "AllOf" "input.allOf"}}
{{getSetValue . "AllowEmptyValue" "input.allowEmptyValue"}}
{{getSetValue . "AllowReserved" "input.allowReserved"}}
{{getSetValue . "AnyOf" "input.anyOf"}}
{{getSetValue . "Default" "input.default"}}
{{getSetValue . "Deprecated" "input.deprecated"}}
{{getSetValue . "Description" "input.description"}}
{{getSetValue . "Discriminator" "input.discriminator"}}
{{getSetValue . "Enum" "input.enum"}}
{{getSetValue . "EnumNames" "input.enumNames"}}
{{getSetValue . "EnumTitles" "input.enumTitles"}}
{{getSetValue . "Example" "input.example"}}
{{getSetValue . "ExclusiveMaximum" "input.exclusiveMaximum"}}
{{getSetValue . "ExclusiveMinimum" "input.exclusiveMinimum"}}
{{getSetValue . "Extensions" "input.extensions"}}
{{getSetValue . "ExternalDocs" "input.externalDocs"}}
{{getSetValue . "ExternalPros" "input.externalPros"}}
{{getSetValue . "Format" "input.format"}}
{{getSetNumber . "MaxItems" "input.maxItems"}}
{{getSetNumber . "MaxLength" "input.maxLength"}}
{{getSetNumber . "MaxProperties" "input.maxProperties"}}
{{getSetNumber . "Maxnum" "input.maxnum"}}
{{getSetNumber . "MinItems" "input.minItems"}}
{{getSetNumber . "MinLength" "input.minLength"}}
{{getSetNumber . "MinProperties" "input.minProperties"}}
{{getSetNumber . "Minimum" "input.minimum"}}
{{getSetNumber . "MultipleOf" "input.multipleOf"}}
{{getSetValue . "Not" "input.not"}}
{{getSetValue . "Nullable" "input.nullable"}}
{{getSetValue . "OneOf" "input.oneOf"}}
{{getSetValue . "Pattern" "input.pattern"}}
{{getSetValue . "ReadOnly" "input.readOnly"}}
{{getSetValue . "Remark" "input.remark"}}
{{getSetValue . "Required" "input.required"}}
{{getSetValue . "ServiceID" "input.serviceId"}}
{{getSetValue . "Summary" "input.summary"}}
{{getSetValue . "Type" "input.type"}}
{{getSetValue . "UniqueItems" "input.uniqueItems"}}
{{getSetValue . "ValidateSchemaID" "input.validateSchemaId"}}
{{getSetValue . "WriteOnly" "input.writeOnly"}}
{{getSetValue . "XML" "input.xml"}}
{{end}}','{{define "output"}}
{{end}}');
insert ignore into `template` (`template_id`,`type`,`title`,`description`,`source_id`,`tpl`) values('DatabaseNameValidateSchemaDel','SQL','database name validate schema 删除','database name validate schema 删除','database_name','{{define "Del"}} update `t_validate_schema` set `deleted_at`={{currentTime .}} where `validate_schema_id`=:ValidateSchemaID; {{end}}');
insert ignore into `api` (`api_id`,`title`,`description`,`method`,`route`,`template_ids`,`exec`,`input`,`output`,`input_tpl`,`output_tpl`) values('database_name-t_validate_schema-Del','database name validate schema 删除','database name validate schema 删除','POST','/api/database_name/v1/validate_schema/del','DatabaseNameValidateSchemaDel','{{define "main"}}
{{execSQLTpl . "Del"}}
{{end}}','version=http://json-schema.org/draft-07/schema,id=input,direction=in
fullname=validateSchemaId,dst=ValidateSchemaID,required','version=http://json-schema.org/draft-07/schema,id=output,direction=out','{{define "input"}}
{{getSetValue . "ValidateSchemaID" "input.validateSchemaId"}}
{{end}}','{{define "output"}}
{{end}}');
insert ignore into `template` (`template_id`,`type`,`title`,`description`,`source_id`,`tpl`) values('DatabaseNameExampleGetByExampleID','SQL','database name 案例 获取 通过 ID','database name 案例 获取 通过 ID','database_name','{{define "GetByExampleID"}} select * from `t_example` where `example_id`=:ExampleID and `deleted_at` is null; {{end}}');
insert ignore into `api` (`api_id`,`title`,`description`,`method`,`route`,`template_ids`,`exec`,`input`,`output`,`input_tpl`,`output_tpl`) values('database_name-t_example-GetByExampleID','database name 案例 获取 通过 ID','database name 案例 获取 通过 ID','POST','/api/database_name/v1/example/get_by_example_id','DatabaseNameExampleGetByExampleID','{{define "main"}}
{{execSQLTpl . "GetByExampleID"}}
{{end}}','version=http://json-schema.org/draft-07/schema,id=input,direction=in
fullname=exampleId,dst=ExampleID,required','version=http://json-schema.org/draft-07/schema,id=output,direction=out
fullname=apiId,src=GetByExampleIDOut.#.api_id,required
fullname=createdAt,src=GetByExampleIDOut.#.created_at,required
fullname=deletedAt,src=GetByExampleIDOut.#.deleted_at,required
fullname=exampleId,src=GetByExampleIDOut.#.example_id,required
fullname=request,src=GetByExampleIDOut.#.request,required
fullname=response,src=GetByExampleIDOut.#.response,required
fullname=serviceId,src=GetByExampleIDOut.#.service_id,required
fullname=summary,src=GetByExampleIDOut.#.summary,required
fullname=tag,src=GetByExampleIDOut.#.tag,required
fullname=title,src=GetByExampleIDOut.#.title,required
fullname=updatedAt,src=GetByExampleIDOut.#.updated_at,required','{{define "input"}}
{{getSetValue . "ExampleID" "input.exampleId"}}
{{end}}','{{define "output"}}
{{getSetValue . "output.apiId" "GetByExampleIDOut.#.api_id"}}
{{getSetValue . "output.createdAt" "GetByExampleIDOut.#.created_at"}}
{{getSetValue . "output.deletedAt" "GetByExampleIDOut.#.deleted_at"}}
{{getSetValue . "output.exampleId" "GetByExampleIDOut.#.example_id"}}
{{getSetValue . "output.request" "GetByExampleIDOut.#.request"}}
{{getSetValue . "output.response" "GetByExampleIDOut.#.response"}}
{{getSetValue . "output.serviceId" "GetByExampleIDOut.#.service_id"}}
{{getSetValue . "output.summary" "GetByExampleIDOut.#.summary"}}
{{getSetValue . "output.tag" "GetByExampleIDOut.#.tag"}}
{{getSetValue . "output.title" "GetByExampleIDOut.#.title"}}
{{getSetValue . "output.updatedAt" "GetByExampleIDOut.#.updated_at"}}
{{end}}');
insert ignore into `template` (`template_id`,`type`,`title`,`description`,`source_id`,`tpl`) values('DatabaseNameExampleGetAllByExampleIDList','SQL','database name 案例 获取 所有 通过 ID 列表','database name 案例 获取 所有 通过 ID 列表','database_name','{{define "GetAllByExampleIDList"}} select * from `t_example` where `example_id` in ({{in . .ExampleIDList}}) and `deleted_at` is null; {{end}}');
insert ignore into `api` (`api_id`,`title`,`description`,`method`,`route`,`template_ids`,`exec`,`input`,`output`,`input_tpl`,`output_tpl`) values('database_name-t_example-GetAllByExampleIDList','database name 案例 获取 所有 通过 ID 列表','database name 案例 获取 所有 通过 ID 列表','POST','/api/database_name/v1/example/get_all_by_example_id_list','DatabaseNameExampleGetAllByExampleIDList','{{define "main"}}
{{execSQLTpl . "GetAllByExampleIDList"}}
{{end}}','version=http://json-schema.org/draft-07/schema,id=input,direction=in
fullname=exampleIdList,dst=ExampleIDList,required','version=http://json-schema.org/draft-07/schema,id=output,direction=out
fullname=apiId,src=GetAllByExampleIDListOut.#.api_id,required
fullname=createdAt,src=GetAllByExampleIDListOut.#.created_at,required
fullname=deletedAt,src=GetAllByExampleIDListOut.#.deleted_at,required
fullname=exampleId,src=GetAllByExampleIDListOut.#.example_id,required
fullname=request,src=GetAllByExampleIDListOut.#.request,required
fullname=response,src=GetAllByExampleIDListOut.#.response,required
fullname=serviceId,src=GetAllByExampleIDListOut.#.service_id,required
fullname=summary,src=GetAllByExampleIDListOut.#.summary,required
fullname=tag,src=GetAllByExampleIDListOut.#.tag,required
fullname=title,src=GetAllByExampleIDListOut.#.title,required
fullname=updatedAt,src=GetAllByExampleIDListOut.#.updated_at,required','{{define "input"}}
{{getSetValue . "ExampleIDList" "input.exampleIdList"}}
{{end}}','{{define "output"}}
{{getSetValue . "output.apiId" "GetAllByExampleIDListOut.#.api_id"}}
{{getSetValue . "output.createdAt" "GetAllByExampleIDListOut.#.created_at"}}
{{getSetValue . "output.deletedAt" "GetAllByExampleIDListOut.#.deleted_at"}}
{{getSetValue . "output.exampleId" "GetAllByExampleIDListOut.#.example_id"}}
{{getSetValue . "output.request" "GetAllByExampleIDListOut.#.request"}}
{{getSetValue . "output.response" "GetAllByExampleIDListOut.#.response"}}
{{getSetValue . "output.serviceId" "GetAllByExampleIDListOut.#.service_id"}}
{{getSetValue . "output.summary" "GetAllByExampleIDListOut.#.summary"}}
{{getSetValue . "output.tag" "GetAllByExampleIDListOut.#.tag"}}
{{getSetValue . "output.title" "GetAllByExampleIDListOut.#.title"}}
{{getSetValue . "output.updatedAt" "GetAllByExampleIDListOut.#.updated_at"}}
{{end}}');
insert ignore into `template` (`template_id`,`type`,`title`,`description`,`source_id`,`tpl`) values('DatabaseNameExamplePaginateWhere','SQL','database name 案例 分页列表 where','database name 案例 分页列表 where','database_name','{{define "PaginateWhere"}} {{end}}');
insert ignore into `template` (`template_id`,`type`,`title`,`description`,`source_id`,`tpl`) values('DatabaseNameExamplePaginateTotal','SQL','database name 案例 分页列表 总数','database name 案例 分页列表 总数','database_name','{{define "PaginateTotal"}} select count(*) as `count` from `t_example` where 1=1 {{template "PaginateWhere" .}} and `deleted_at` is null; {{end}}');
insert ignore into `api` (`api_id`,`title`,`description`,`method`,`route`,`template_ids`,`exec`,`input`,`output`,`input_tpl`,`output_tpl`) values('database_name-t_example-PaginateTotal','database name 案例 分页列表 总数','database name 案例 分页列表 总数','POST','/api/database_name/v1/example/paginate_total','DatabaseNameExamplePaginateTotal','{{define "main"}}
{{execSQLTpl . "PaginateTotal"}}
{{end}}','version=http://json-schema.org/draft-07/schema,id=input,direction=in','version=http://json-schema.org/draft-07/schema,id=output,direction=out','{{define "input"}}
{{end}}','{{define "output"}}
{{end}}');
insert ignore into `template` (`template_id`,`type`,`title`,`description`,`source_id`,`tpl`) values('DatabaseNameExamplePaginate','SQL','database name 案例 分页列表','database name 案例 分页列表','database_name','{{define "Paginate"}} select * from `t_example` where 1=1 {{template "PaginateWhere" .}} and `deleted_at` is null order by `updated_at` desc limit :Offset,:Limit ; {{end}}');
insert ignore into `api` (`api_id`,`title`,`description`,`method`,`route`,`template_ids`,`exec`,`input`,`output`,`input_tpl`,`output_tpl`) values('database_name-t_example-Paginate','database name 案例 分页列表','database name 案例 分页列表','POST','/api/database_name/v1/example/paginate','DatabaseNameExamplePaginateWhere,DatabaseNameExamplePaginateTotal,DatabaseNameExamplePaginate','{{define "main"}}
{{execSQLTpl . "PaginateTotal"}}
{{execSQLTpl . "Paginate"}}
{{end}}','version=http://json-schema.org/draft-07/schema,id=input,direction=in
fullname=limit,dst=Limit,format=number,required
fullname=offset,dst=Offset,format=number,required','version=http://json-schema.org/draft-07/schema,id=output,direction=out
fullname=apiId,src=PaginateOut.#.api_id,required
fullname=createdAt,src=PaginateOut.#.created_at,required
fullname=deletedAt,src=PaginateOut.#.deleted_at,required
fullname=exampleId,src=PaginateOut.#.example_id,required
fullname=request,src=PaginateOut.#.request,required
fullname=response,src=PaginateOut.#.response,required
fullname=serviceId,src=PaginateOut.#.service_id,required
fullname=summary,src=PaginateOut.#.summary,required
fullname=tag,src=PaginateOut.#.tag,required
fullname=title,src=PaginateOut.#.title,required
fullname=updatedAt,src=PaginateOut.#.updated_at,required','{{define "input"}}
{{getSetNumber . "Limit" "input.limit"}}
{{getSetNumber . "Offset" "input.offset"}}
{{end}}','{{define "output"}}
{{getSetValue . "output.apiId" "PaginateOut.#.api_id"}}
{{getSetValue . "output.createdAt" "PaginateOut.#.created_at"}}
{{getSetValue . "output.deletedAt" "PaginateOut.#.deleted_at"}}
{{getSetValue . "output.exampleId" "PaginateOut.#.example_id"}}
{{getSetValue . "output.request" "PaginateOut.#.request"}}
{{getSetValue . "output.response" "PaginateOut.#.response"}}
{{getSetValue . "output.serviceId" "PaginateOut.#.service_id"}}
{{getSetValue . "output.summary" "PaginateOut.#.summary"}}
{{getSetValue . "output.tag" "PaginateOut.#.tag"}}
{{getSetValue . "output.title" "PaginateOut.#.title"}}
{{getSetValue . "output.updatedAt" "PaginateOut.#.updated_at"}}
{{end}}');
insert ignore into `template` (`template_id`,`type`,`title`,`description`,`source_id`,`tpl`) values('DatabaseNameExampleInsert','SQL','database name 案例 新增','database name 案例 新增','database_name','{{define "Insert"}} insert into `t_example` (`example_id`,`service_id`,`api_id`,`tag`,`title`,`summary`,`request`,`response`)values (:ExampleID,:ServiceID,:APIID,:Tag,:Title,:Summary,:Request,:Response); {{end}}');
insert ignore into `api` (`api_id`,`title`,`description`,`method`,`route`,`template_ids`,`exec`,`input`,`output`,`input_tpl`,`output_tpl`) values('database_name-t_example-Insert','database name 案例 新增','database name 案例 新增','POST','/api/database_name/v1/example/insert','DatabaseNameExampleInsert','{{define "main"}}
{{execSQLTpl . "Insert"}}
{{end}}','version=http://json-schema.org/draft-07/schema,id=input,direction=in
fullname=apiid,dst=APIID,required
fullname=exampleId,dst=ExampleID,required
fullname=request,dst=Request,required
fullname=response,dst=Response,required
fullname=serviceId,dst=ServiceID,required
fullname=summary,dst=Summary,required
fullname=tag,dst=Tag,required
fullname=title,dst=Title,required','version=http://json-schema.org/draft-07/schema,id=output,direction=out','{{define "input"}}
{{getSetValue . "APIID" "input.apiid"}}
{{getSetValue . "ExampleID" "input.exampleId"}}
{{getSetValue . "Request" "input.request"}}
{{getSetValue . "Response" "input.response"}}
{{getSetValue . "ServiceID" "input.serviceId"}}
{{getSetValue . "Summary" "input.summary"}}
{{getSetValue . "Tag" "input.tag"}}
{{getSetValue . "Title" "input.title"}}
{{end}}','{{define "output"}}
{{end}}');
insert ignore into `template` (`template_id`,`type`,`title`,`description`,`source_id`,`tpl`) values('DatabaseNameExampleUpdate','SQL','database name 案例 修改','database name 案例 修改','database_name','{{define "Update"}} {{$preComma:=newPreComma}} update `t_example` set {{if .ExampleID}} {{$preComma.PreComma}} `example_id`=:ExampleID {{end}} {{if .ServiceID}} {{$preComma.PreComma}} `service_id`=:ServiceID {{end}} {{if .APIID}} {{$preComma.PreComma}} `api_id`=:APIID {{end}} {{if .Tag}} {{$preComma.PreComma}} `tag`=:Tag {{end}} {{if .Title}} {{$preComma.PreComma}} `title`=:Title {{end}} {{if .Summary}} {{$preComma.PreComma}} `summary`=:Summary {{end}} {{if .Request}} {{$preComma.PreComma}} `request`=:Request {{end}} {{if .Response}} {{$preComma.PreComma}} `response`=:Response {{end}} where `example_id`=:ExampleID; {{end}}');
insert ignore into `api` (`api_id`,`title`,`description`,`method`,`route`,`template_ids`,`exec`,`input`,`output`,`input_tpl`,`output_tpl`) values('database_name-t_example-Update','database name 案例 修改','database name 案例 修改','POST','/api/database_name/v1/example/update','DatabaseNameExampleUpdate','{{define "main"}}
{{execSQLTpl . "Update"}}
{{end}}','version=http://json-schema.org/draft-07/schema,id=input,direction=in
fullname=apiid,dst=APIID,required
fullname=exampleId,dst=ExampleID,required
fullname=request,dst=Request,required
fullname=response,dst=Response,required
fullname=serviceId,dst=ServiceID,required
fullname=summary,dst=Summary,required
fullname=tag,dst=Tag,required
fullname=title,dst=Title,required','version=http://json-schema.org/draft-07/schema,id=output,direction=out','{{define "input"}}
{{getSetValue . "APIID" "input.apiid"}}
{{getSetValue . "ExampleID" "input.exampleId"}}
{{getSetValue . "Request" "input.request"}}
{{getSetValue . "Response" "input.response"}}
{{getSetValue . "ServiceID" "input.serviceId"}}
{{getSetValue . "Summary" "input.summary"}}
{{getSetValue . "Tag" "input.tag"}}
{{getSetValue . "Title" "input.title"}}
{{end}}','{{define "output"}}
{{end}}');
insert ignore into `template` (`template_id`,`type`,`title`,`description`,`source_id`,`tpl`) values('DatabaseNameExampleDel','SQL','database name 案例 删除','database name 案例 删除','database_name','{{define "Del"}} update `t_example` set `deleted_at`={{currentTime .}} where `example_id`=:ExampleID; {{end}}');
insert ignore into `api` (`api_id`,`title`,`description`,`method`,`route`,`template_ids`,`exec`,`input`,`output`,`input_tpl`,`output_tpl`) values('database_name-t_example-Del','database name 案例 删除','database name 案例 删除','POST','/api/database_name/v1/example/del','DatabaseNameExampleDel','{{define "main"}}
{{execSQLTpl . "Del"}}
{{end}}','version=http://json-schema.org/draft-07/schema,id=input,direction=in
fullname=exampleId,dst=ExampleID,required','version=http://json-schema.org/draft-07/schema,id=output,direction=out','{{define "input"}}
{{getSetValue . "ExampleID" "input.exampleId"}}
{{end}}','{{define "output"}}
{{end}}');
insert ignore into `template` (`template_id`,`type`,`title`,`description`,`source_id`,`tpl`) values('DatabaseNameMarkdownGetByMarkdownID','SQL','database name markdown 文档 获取 通过 ID','database name markdown 文档 获取 通过 ID','database_name','{{define "GetByMarkdownID"}} select * from `t_markdown` where `markdown_id`=:MarkdownID and `deleted_at` is null; {{end}}');
insert ignore into `api` (`api_id`,`title`,`description`,`method`,`route`,`template_ids`,`exec`,`input`,`output`,`input_tpl`,`output_tpl`) values('database_name-t_markdown-GetByMarkdownID','database name markdown 文档 获取 通过 ID','database name markdown 文档 获取 通过 ID','POST','/api/database_name/v1/markdown/get_by_markdown_id','DatabaseNameMarkdownGetByMarkdownID','{{define "main"}}
{{execSQLTpl . "GetByMarkdownID"}}
{{end}}','version=http://json-schema.org/draft-07/schema,id=input,direction=in
fullname=markdownId,dst=MarkdownID,required','version=http://json-schema.org/draft-07/schema,id=output,direction=out
fullname=apiId,src=GetByMarkdownIDOut.#.api_id,required
fullname=content,src=GetByMarkdownIDOut.#.content,required
fullname=createdAt,src=GetByMarkdownIDOut.#.created_at,required
fullname=deletedAt,src=GetByMarkdownIDOut.#.deleted_at,required
fullname=markdown,src=GetByMarkdownIDOut.#.markdown,required
fullname=markdownId,src=GetByMarkdownIDOut.#.markdown_id,required
fullname=name,src=GetByMarkdownIDOut.#.name,required
fullname=ownerId,src=GetByMarkdownIDOut.#.owner_id,required
fullname=ownerName,src=GetByMarkdownIDOut.#.owner_name,required
fullname=serviceId,src=GetByMarkdownIDOut.#.service_id,required
fullname=title,src=GetByMarkdownIDOut.#.title,required
fullname=updatedAt,src=GetByMarkdownIDOut.#.updated_at,required','{{define "input"}}
{{getSetValue . "MarkdownID" "input.markdownId"}}
{{end}}','{{define "output"}}
{{getSetValue . "output.apiId" "GetByMarkdownIDOut.#.api_id"}}
{{getSetValue . "output.content" "GetByMarkdownIDOut.#.content"}}
{{getSetValue . "output.createdAt" "GetByMarkdownIDOut.#.created_at"}}
{{getSetValue . "output.deletedAt" "GetByMarkdownIDOut.#.deleted_at"}}
{{getSetValue . "output.markdown" "GetByMarkdownIDOut.#.markdown"}}
{{getSetValue . "output.markdownId" "GetByMarkdownIDOut.#.markdown_id"}}
{{getSetValue . "output.name" "GetByMarkdownIDOut.#.name"}}
{{getSetValue . "output.ownerId" "GetByMarkdownIDOut.#.owner_id"}}
{{getSetValue . "output.ownerName" "GetByMarkdownIDOut.#.owner_name"}}
{{getSetValue . "output.serviceId" "GetByMarkdownIDOut.#.service_id"}}
{{getSetValue . "output.title" "GetByMarkdownIDOut.#.title"}}
{{getSetValue . "output.updatedAt" "GetByMarkdownIDOut.#.updated_at"}}
{{end}}');
insert ignore into `template` (`template_id`,`type`,`title`,`description`,`source_id`,`tpl`) values('DatabaseNameMarkdownGetAllByMarkdownIDList','SQL','database name markdown 文档 获取 所有 通过 ID 列表','database name markdown 文档 获取 所有 通过 ID 列表','database_name','{{define "GetAllByMarkdownIDList"}} select * from `t_markdown` where `markdown_id` in ({{in . .MarkdownIDList}}) and `deleted_at` is null; {{end}}');
insert ignore into `api` (`api_id`,`title`,`description`,`method`,`route`,`template_ids`,`exec`,`input`,`output`,`input_tpl`,`output_tpl`) values('database_name-t_markdown-GetAllByMarkdownIDList','database name markdown 文档 获取 所有 通过 ID 列表','database name markdown 文档 获取 所有 通过 ID 列表','POST','/api/database_name/v1/markdown/get_all_by_markdown_id_list','DatabaseNameMarkdownGetAllByMarkdownIDList','{{define "main"}}
{{execSQLTpl . "GetAllByMarkdownIDList"}}
{{end}}','version=http://json-schema.org/draft-07/schema,id=input,direction=in
fullname=markdownIdList,dst=MarkdownIDList,required','version=http://json-schema.org/draft-07/schema,id=output,direction=out
fullname=apiId,src=GetAllByMarkdownIDListOut.#.api_id,required
fullname=content,src=GetAllByMarkdownIDListOut.#.content,required
fullname=createdAt,src=GetAllByMarkdownIDListOut.#.created_at,required
fullname=deletedAt,src=GetAllByMarkdownIDListOut.#.deleted_at,required
fullname=markdown,src=GetAllByMarkdownIDListOut.#.markdown,required
fullname=markdownId,src=GetAllByMarkdownIDListOut.#.markdown_id,required
fullname=name,src=GetAllByMarkdownIDListOut.#.name,required
fullname=ownerId,src=GetAllByMarkdownIDListOut.#.owner_id,required
fullname=ownerName,src=GetAllByMarkdownIDListOut.#.owner_name,required
fullname=serviceId,src=GetAllByMarkdownIDListOut.#.service_id,required
fullname=title,src=GetAllByMarkdownIDListOut.#.title,required
fullname=updatedAt,src=GetAllByMarkdownIDListOut.#.updated_at,required','{{define "input"}}
{{getSetValue . "MarkdownIDList" "input.markdownIdList"}}
{{end}}','{{define "output"}}
{{getSetValue . "output.apiId" "GetAllByMarkdownIDListOut.#.api_id"}}
{{getSetValue . "output.content" "GetAllByMarkdownIDListOut.#.content"}}
{{getSetValue . "output.createdAt" "GetAllByMarkdownIDListOut.#.created_at"}}
{{getSetValue . "output.deletedAt" "GetAllByMarkdownIDListOut.#.deleted_at"}}
{{getSetValue . "output.markdown" "GetAllByMarkdownIDListOut.#.markdown"}}
{{getSetValue . "output.markdownId" "GetAllByMarkdownIDListOut.#.markdown_id"}}
{{getSetValue . "output.name" "GetAllByMarkdownIDListOut.#.name"}}
{{getSetValue . "output.ownerId" "GetAllByMarkdownIDListOut.#.owner_id"}}
{{getSetValue . "output.ownerName" "GetAllByMarkdownIDListOut.#.owner_name"}}
{{getSetValue . "output.serviceId" "GetAllByMarkdownIDListOut.#.service_id"}}
{{getSetValue . "output.title" "GetAllByMarkdownIDListOut.#.title"}}
{{getSetValue . "output.updatedAt" "GetAllByMarkdownIDListOut.#.updated_at"}}
{{end}}');
insert ignore into `template` (`template_id`,`type`,`title`,`description`,`source_id`,`tpl`) values('DatabaseNameMarkdownPaginateWhere','SQL','database name markdown 文档 分页列表 where','database name markdown 文档 分页列表 where','database_name','{{define "PaginateWhere"}} {{end}}');
insert ignore into `template` (`template_id`,`type`,`title`,`description`,`source_id`,`tpl`) values('DatabaseNameMarkdownPaginateTotal','SQL','database name markdown 文档 分页列表 总数','database name markdown 文档 分页列表 总数','database_name','{{define "PaginateTotal"}} select count(*) as `count` from `t_markdown` where 1=1 {{template "PaginateWhere" .}} and `deleted_at` is null; {{end}}');
insert ignore into `api` (`api_id`,`title`,`description`,`method`,`route`,`template_ids`,`exec`,`input`,`output`,`input_tpl`,`output_tpl`) values('database_name-t_markdown-PaginateTotal','database name markdown 文档 分页列表 总数','database name markdown 文档 分页列表 总数','POST','/api/database_name/v1/markdown/paginate_total','DatabaseNameMarkdownPaginateTotal','{{define "main"}}
{{execSQLTpl . "PaginateTotal"}}
{{end}}','version=http://json-schema.org/draft-07/schema,id=input,direction=in','version=http://json-schema.org/draft-07/schema,id=output,direction=out','{{define "input"}}
{{end}}','{{define "output"}}
{{end}}');
insert ignore into `template` (`template_id`,`type`,`title`,`description`,`source_id`,`tpl`) values('DatabaseNameMarkdownPaginate','SQL','database name markdown 文档 分页列表','database name markdown 文档 分页列表','database_name','{{define "Paginate"}} select * from `t_markdown` where 1=1 {{template "PaginateWhere" .}} and `deleted_at` is null order by `updated_at` desc limit :Offset,:Limit ; {{end}}');
insert ignore into `api` (`api_id`,`title`,`description`,`method`,`route`,`template_ids`,`exec`,`input`,`output`,`input_tpl`,`output_tpl`) values('database_name-t_markdown-Paginate','database name markdown 文档 分页列表','database name markdown 文档 分页列表','POST','/api/database_name/v1/markdown/paginate','DatabaseNameMarkdownPaginateWhere,DatabaseNameMarkdownPaginateTotal,DatabaseNameMarkdownPaginate','{{define "main"}}
{{execSQLTpl . "PaginateTotal"}}
{{execSQLTpl . "Paginate"}}
{{end}}','version=http://json-schema.org/draft-07/schema,id=input,direction=in
fullname=limit,dst=Limit,format=number,required
fullname=offset,dst=Offset,format=number,required','version=http://json-schema.org/draft-07/schema,id=output,direction=out
fullname=apiId,src=PaginateOut.#.api_id,required
fullname=content,src=PaginateOut.#.content,required
fullname=createdAt,src=PaginateOut.#.created_at,required
fullname=deletedAt,src=PaginateOut.#.deleted_at,required
fullname=markdown,src=PaginateOut.#.markdown,required
fullname=markdownId,src=PaginateOut.#.markdown_id,required
fullname=name,src=PaginateOut.#.name,required
fullname=ownerId,src=PaginateOut.#.owner_id,required
fullname=ownerName,src=PaginateOut.#.owner_name,required
fullname=serviceId,src=PaginateOut.#.service_id,required
fullname=title,src=PaginateOut.#.title,required
fullname=updatedAt,src=PaginateOut.#.updated_at,required','{{define "input"}}
{{getSetNumber . "Limit" "input.limit"}}
{{getSetNumber . "Offset" "input.offset"}}
{{end}}','{{define "output"}}
{{getSetValue . "output.apiId" "PaginateOut.#.api_id"}}
{{getSetValue . "output.content" "PaginateOut.#.content"}}
{{getSetValue . "output.createdAt" "PaginateOut.#.created_at"}}
{{getSetValue . "output.deletedAt" "PaginateOut.#.deleted_at"}}
{{getSetValue . "output.markdown" "PaginateOut.#.markdown"}}
{{getSetValue . "output.markdownId" "PaginateOut.#.markdown_id"}}
{{getSetValue . "output.name" "PaginateOut.#.name"}}
{{getSetValue . "output.ownerId" "PaginateOut.#.owner_id"}}
{{getSetValue . "output.ownerName" "PaginateOut.#.owner_name"}}
{{getSetValue . "output.serviceId" "PaginateOut.#.service_id"}}
{{getSetValue . "output.title" "PaginateOut.#.title"}}
{{getSetValue . "output.updatedAt" "PaginateOut.#.updated_at"}}
{{end}}');
insert ignore into `template` (`template_id`,`type`,`title`,`description`,`source_id`,`tpl`) values('DatabaseNameMarkdownInsert','SQL','database name markdown 文档 新增','database name markdown 文档 新增','database_name','{{define "Insert"}} insert into `t_markdown` (`markdown_id`,`service_id`,`api_id`,`name`,`title`,`markdown`,`content`,`owner_id`,`owner_name`)values (:MarkdownID,:ServiceID,:APIID,:Name,:Title,:Markdown,:Content,:OwnerID,:OwnerName); {{end}}');
insert ignore into `api` (`api_id`,`title`,`description`,`method`,`route`,`template_ids`,`exec`,`input`,`output`,`input_tpl`,`output_tpl`) values('database_name-t_markdown-Insert','database name markdown 文档 新增','database name markdown 文档 新增','POST','/api/database_name/v1/markdown/insert','DatabaseNameMarkdownInsert','{{define "main"}}
{{execSQLTpl . "Insert"}}
{{end}}','version=http://json-schema.org/draft-07/schema,id=input,direction=in
fullname=apiid,dst=APIID,required
fullname=content,dst=Content,required
fullname=markdown,dst=Markdown,required
fullname=markdownId,dst=MarkdownID,required
fullname=name,dst=Name,required
fullname=ownerId,dst=OwnerID,format=number,required
fullname=ownerName,dst=OwnerName,required
fullname=serviceId,dst=ServiceID,required
fullname=title,dst=Title,required','version=http://json-schema.org/draft-07/schema,id=output,direction=out','{{define "input"}}
{{getSetValue . "APIID" "input.apiid"}}
{{getSetValue . "Content" "input.content"}}
{{getSetValue . "Markdown" "input.markdown"}}
{{getSetValue . "MarkdownID" "input.markdownId"}}
{{getSetValue . "Name" "input.name"}}
{{getSetNumber . "OwnerID" "input.ownerId"}}
{{getSetValue . "OwnerName" "input.ownerName"}}
{{getSetValue . "ServiceID" "input.serviceId"}}
{{getSetValue . "Title" "input.title"}}
{{end}}','{{define "output"}}
{{end}}');
insert ignore into `template` (`template_id`,`type`,`title`,`description`,`source_id`,`tpl`) values('DatabaseNameMarkdownUpdate','SQL','database name markdown 文档 修改','database name markdown 文档 修改','database_name','{{define "Update"}} {{$preComma:=newPreComma}} update `t_markdown` set {{if .MarkdownID}} {{$preComma.PreComma}} `markdown_id`=:MarkdownID {{end}} {{if .ServiceID}} {{$preComma.PreComma}} `service_id`=:ServiceID {{end}} {{if .APIID}} {{$preComma.PreComma}} `api_id`=:APIID {{end}} {{if .Name}} {{$preComma.PreComma}} `name`=:Name {{end}} {{if .Title}} {{$preComma.PreComma}} `title`=:Title {{end}} {{if .Markdown}} {{$preComma.PreComma}} `markdown`=:Markdown {{end}} {{if .Content}} {{$preComma.PreComma}} `content`=:Content {{end}} {{if .OwnerID}} {{$preComma.PreComma}} `owner_id`=:OwnerID {{end}} {{if .OwnerName}} {{$preComma.PreComma}} `owner_name`=:OwnerName {{end}} where `markdown_id`=:MarkdownID; {{end}}');
insert ignore into `api` (`api_id`,`title`,`description`,`method`,`route`,`template_ids`,`exec`,`input`,`output`,`input_tpl`,`output_tpl`) values('database_name-t_markdown-Update','database name markdown 文档 修改','database name markdown 文档 修改','POST','/api/database_name/v1/markdown/update','DatabaseNameMarkdownUpdate','{{define "main"}}
{{execSQLTpl . "Update"}}
{{end}}','version=http://json-schema.org/draft-07/schema,id=input,direction=in
fullname=apiid,dst=APIID,required
fullname=content,dst=Content,required
fullname=markdown,dst=Markdown,required
fullname=markdownId,dst=MarkdownID,required
fullname=name,dst=Name,required
fullname=ownerId,dst=OwnerID,format=number,required
fullname=ownerName,dst=OwnerName,required
fullname=serviceId,dst=ServiceID,required
fullname=title,dst=Title,required','version=http://json-schema.org/draft-07/schema,id=output,direction=out','{{define "input"}}
{{getSetValue . "APIID" "input.apiid"}}
{{getSetValue . "Content" "input.content"}}
{{getSetValue . "Markdown" "input.markdown"}}
{{getSetValue . "MarkdownID" "input.markdownId"}}
{{getSetValue . "Name" "input.name"}}
{{getSetNumber . "OwnerID" "input.ownerId"}}
{{getSetValue . "OwnerName" "input.ownerName"}}
{{getSetValue . "ServiceID" "input.serviceId"}}
{{getSetValue . "Title" "input.title"}}
{{end}}','{{define "output"}}
{{end}}');
insert ignore into `template` (`template_id`,`type`,`title`,`description`,`source_id`,`tpl`) values('DatabaseNameMarkdownDel','SQL','database name markdown 文档 删除','database name markdown 文档 删除','database_name','{{define "Del"}} update `t_markdown` set `deleted_at`={{currentTime .}} where `markdown_id`=:MarkdownID; {{end}}');
insert ignore into `api` (`api_id`,`title`,`description`,`method`,`route`,`template_ids`,`exec`,`input`,`output`,`input_tpl`,`output_tpl`) values('database_name-t_markdown-Del','database name markdown 文档 删除','database name markdown 文档 删除','POST','/api/database_name/v1/markdown/del','DatabaseNameMarkdownDel','{{define "main"}}
{{execSQLTpl . "Del"}}
{{end}}','version=http://json-schema.org/draft-07/schema,id=input,direction=in
fullname=markdownId,dst=MarkdownID,required','version=http://json-schema.org/draft-07/schema,id=output,direction=out','{{define "input"}}
{{getSetValue . "MarkdownID" "input.markdownId"}}
{{end}}','{{define "output"}}
{{end}}');
insert ignore into `template` (`template_id`,`type`,`title`,`description`,`source_id`,`tpl`) values('DatabaseNameServerGetByServerID','SQL','database name 服务器部署 获取 通过 ID','database name 服务器部署 获取 通过 ID','database_name','{{define "GetByServerID"}} select * from `t_server` where `server_id`=:ServerID and `deleted_at` is null; {{end}}');
insert ignore into `api` (`api_id`,`title`,`description`,`method`,`route`,`template_ids`,`exec`,`input`,`output`,`input_tpl`,`output_tpl`) values('database_name-t_server-GetByServerID','database name 服务器部署 获取 通过 ID','database name 服务器部署 获取 通过 ID','POST','/api/database_name/v1/server/get_by_server_id','DatabaseNameServerGetByServerID','{{define "main"}}
{{execSQLTpl . "GetByServerID"}}
{{end}}','version=http://json-schema.org/draft-07/schema,id=input,direction=in
fullname=serverId,dst=ServerID,required','version=http://json-schema.org/draft-07/schema,id=output,direction=out
fullname=createdAt,src=GetByServerIDOut.#.created_at,required
fullname=deletedAt,src=GetByServerIDOut.#.deleted_at,required
fullname=description,src=GetByServerIDOut.#.description,required
fullname=extensionIds,src=GetByServerIDOut.#.extension_ids,required
fullname=proxy,src=GetByServerIDOut.#.proxy,required
fullname=serverId,src=GetByServerIDOut.#.server_id,required
fullname=serviceId,src=GetByServerIDOut.#.service_id,required
fullname=updatedAt,src=GetByServerIDOut.#.updated_at,required
fullname=url,src=GetByServerIDOut.#.url,required','{{define "input"}}
{{getSetValue . "ServerID" "input.serverId"}}
{{end}}','{{define "output"}}
{{getSetValue . "output.createdAt" "GetByServerIDOut.#.created_at"}}
{{getSetValue . "output.deletedAt" "GetByServerIDOut.#.deleted_at"}}
{{getSetValue . "output.description" "GetByServerIDOut.#.description"}}
{{getSetValue . "output.extensionIds" "GetByServerIDOut.#.extension_ids"}}
{{getSetValue . "output.proxy" "GetByServerIDOut.#.proxy"}}
{{getSetValue . "output.serverId" "GetByServerIDOut.#.server_id"}}
{{getSetValue . "output.serviceId" "GetByServerIDOut.#.service_id"}}
{{getSetValue . "output.updatedAt" "GetByServerIDOut.#.updated_at"}}
{{getSetValue . "output.url" "GetByServerIDOut.#.url"}}
{{end}}');
insert ignore into `template` (`template_id`,`type`,`title`,`description`,`source_id`,`tpl`) values('DatabaseNameServerGetAllByServerIDList','SQL','database name 服务器部署 获取 所有 通过 ID 列表','database name 服务器部署 获取 所有 通过 ID 列表','database_name','{{define "GetAllByServerIDList"}} select * from `t_server` where `server_id` in ({{in . .ServerIDList}}) and `deleted_at` is null; {{end}}');
insert ignore into `api` (`api_id`,`title`,`description`,`method`,`route`,`template_ids`,`exec`,`input`,`output`,`input_tpl`,`output_tpl`) values('database_name-t_server-GetAllByServerIDList','database name 服务器部署 获取 所有 通过 ID 列表','database name 服务器部署 获取 所有 通过 ID 列表','POST','/api/database_name/v1/server/get_all_by_server_id_list','DatabaseNameServerGetAllByServerIDList','{{define "main"}}
{{execSQLTpl . "GetAllByServerIDList"}}
{{end}}','version=http://json-schema.org/draft-07/schema,id=input,direction=in
fullname=serverIdList,dst=ServerIDList,required','version=http://json-schema.org/draft-07/schema,id=output,direction=out
fullname=createdAt,src=GetAllByServerIDListOut.#.created_at,required
fullname=deletedAt,src=GetAllByServerIDListOut.#.deleted_at,required
fullname=description,src=GetAllByServerIDListOut.#.description,required
fullname=extensionIds,src=GetAllByServerIDListOut.#.extension_ids,required
fullname=proxy,src=GetAllByServerIDListOut.#.proxy,required
fullname=serverId,src=GetAllByServerIDListOut.#.server_id,required
fullname=serviceId,src=GetAllByServerIDListOut.#.service_id,required
fullname=updatedAt,src=GetAllByServerIDListOut.#.updated_at,required
fullname=url,src=GetAllByServerIDListOut.#.url,required','{{define "input"}}
{{getSetValue . "ServerIDList" "input.serverIdList"}}
{{end}}','{{define "output"}}
{{getSetValue . "output.createdAt" "GetAllByServerIDListOut.#.created_at"}}
{{getSetValue . "output.deletedAt" "GetAllByServerIDListOut.#.deleted_at"}}
{{getSetValue . "output.description" "GetAllByServerIDListOut.#.description"}}
{{getSetValue . "output.extensionIds" "GetAllByServerIDListOut.#.extension_ids"}}
{{getSetValue . "output.proxy" "GetAllByServerIDListOut.#.proxy"}}
{{getSetValue . "output.serverId" "GetAllByServerIDListOut.#.server_id"}}
{{getSetValue . "output.serviceId" "GetAllByServerIDListOut.#.service_id"}}
{{getSetValue . "output.updatedAt" "GetAllByServerIDListOut.#.updated_at"}}
{{getSetValue . "output.url" "GetAllByServerIDListOut.#.url"}}
{{end}}');
insert ignore into `template` (`template_id`,`type`,`title`,`description`,`source_id`,`tpl`) values('DatabaseNameServerPaginateWhere','SQL','database name 服务器部署 分页列表 where','database name 服务器部署 分页列表 where','database_name','{{define "PaginateWhere"}} {{end}}');
insert ignore into `template` (`template_id`,`type`,`title`,`description`,`source_id`,`tpl`) values('DatabaseNameServerPaginateTotal','SQL','database name 服务器部署 分页列表 总数','database name 服务器部署 分页列表 总数','database_name','{{define "PaginateTotal"}} select count(*) as `count` from `t_server` where 1=1 {{template "PaginateWhere" .}} and `deleted_at` is null; {{end}}');
insert ignore into `api` (`api_id`,`title`,`description`,`method`,`route`,`template_ids`,`exec`,`input`,`output`,`input_tpl`,`output_tpl`) values('database_name-t_server-PaginateTotal','database name 服务器部署 分页列表 总数','database name 服务器部署 分页列表 总数','POST','/api/database_name/v1/server/paginate_total','DatabaseNameServerPaginateTotal','{{define "main"}}
{{execSQLTpl . "PaginateTotal"}}
{{end}}','version=http://json-schema.org/draft-07/schema,id=input,direction=in','version=http://json-schema.org/draft-07/schema,id=output,direction=out','{{define "input"}}
{{end}}','{{define "output"}}
{{end}}');
insert ignore into `template` (`template_id`,`type`,`title`,`description`,`source_id`,`tpl`) values('DatabaseNameServerPaginate','SQL','database name 服务器部署 分页列表','database name 服务器部署 分页列表','database_name','{{define "Paginate"}} select * from `t_server` where 1=1 {{template "PaginateWhere" .}} and `deleted_at` is null order by `updated_at` desc limit :Offset,:Limit ; {{end}}');
insert ignore into `api` (`api_id`,`title`,`description`,`method`,`route`,`template_ids`,`exec`,`input`,`output`,`input_tpl`,`output_tpl`) values('database_name-t_server-Paginate','database name 服务器部署 分页列表','database name 服务器部署 分页列表','POST','/api/database_name/v1/server/paginate','DatabaseNameServerPaginateWhere,DatabaseNameServerPaginateTotal,DatabaseNameServerPaginate','{{define "main"}}
{{execSQLTpl . "PaginateTotal"}}
{{execSQLTpl . "Paginate"}}
{{end}}','version=http://json-schema.org/draft-07/schema,id=input,direction=in
fullname=limit,dst=Limit,format=number,required
fullname=offset,dst=Offset,format=number,required','version=http://json-schema.org/draft-07/schema,id=output,direction=out
fullname=createdAt,src=PaginateOut.#.created_at,required
fullname=deletedAt,src=PaginateOut.#.deleted_at,required
fullname=description,src=PaginateOut.#.description,required
fullname=extensionIds,src=PaginateOut.#.extension_ids,required
fullname=proxy,src=PaginateOut.#.proxy,required
fullname=serverId,src=PaginateOut.#.server_id,required
fullname=serviceId,src=PaginateOut.#.service_id,required
fullname=updatedAt,src=PaginateOut.#.updated_at,required
fullname=url,src=PaginateOut.#.url,required','{{define "input"}}
{{getSetNumber . "Limit" "input.limit"}}
{{getSetNumber . "Offset" "input.offset"}}
{{end}}','{{define "output"}}
{{getSetValue . "output.createdAt" "PaginateOut.#.created_at"}}
{{getSetValue . "output.deletedAt" "PaginateOut.#.deleted_at"}}
{{getSetValue . "output.description" "PaginateOut.#.description"}}
{{getSetValue . "output.extensionIds" "PaginateOut.#.extension_ids"}}
{{getSetValue . "output.proxy" "PaginateOut.#.proxy"}}
{{getSetValue . "output.serverId" "PaginateOut.#.server_id"}}
{{getSetValue . "output.serviceId" "PaginateOut.#.service_id"}}
{{getSetValue . "output.updatedAt" "PaginateOut.#.updated_at"}}
{{getSetValue . "output.url" "PaginateOut.#.url"}}
{{end}}');
insert ignore into `template` (`template_id`,`type`,`title`,`description`,`source_id`,`tpl`) values('DatabaseNameServerInsert','SQL','database name 服务器部署 新增','database name 服务器部署 新增','database_name','{{define "Insert"}} insert into `t_server` (`server_id`,`service_id`,`url`,`description`,`proxy`,`extension_ids`)values (:ServerID,:ServiceID,:URL,:Description,:Proxy,:ExtensionIds); {{end}}');
insert ignore into `api` (`api_id`,`title`,`description`,`method`,`route`,`template_ids`,`exec`,`input`,`output`,`input_tpl`,`output_tpl`) values('database_name-t_server-Insert','database name 服务器部署 新增','database name 服务器部署 新增','POST','/api/database_name/v1/server/insert','DatabaseNameServerInsert','{{define "main"}}
{{execSQLTpl . "Insert"}}
{{end}}','version=http://json-schema.org/draft-07/schema,id=input,direction=in
fullname=description,dst=Description,required
fullname=extensionIds,dst=ExtensionIds,required
fullname=proxy,dst=Proxy,required
fullname=serverId,dst=ServerID,required
fullname=serviceId,dst=ServiceID,required
fullname=url,dst=URL,required','version=http://json-schema.org/draft-07/schema,id=output,direction=out','{{define "input"}}
{{getSetValue . "Description" "input.description"}}
{{getSetValue . "ExtensionIds" "input.extensionIds"}}
{{getSetValue . "Proxy" "input.proxy"}}
{{getSetValue . "ServerID" "input.serverId"}}
{{getSetValue . "ServiceID" "input.serviceId"}}
{{getSetValue . "URL" "input.url"}}
{{end}}','{{define "output"}}
{{end}}');
insert ignore into `template` (`template_id`,`type`,`title`,`description`,`source_id`,`tpl`) values('DatabaseNameServerUpdate','SQL','database name 服务器部署 修改','database name 服务器部署 修改','database_name','{{define "Update"}} {{$preComma:=newPreComma}} update `t_server` set {{if .ServerID}} {{$preComma.PreComma}} `server_id`=:ServerID {{end}} {{if .ServiceID}} {{$preComma.PreComma}} `service_id`=:ServiceID {{end}} {{if .URL}} {{$preComma.PreComma}} `url`=:URL {{end}} {{if .Description}} {{$preComma.PreComma}} `description`=:Description {{end}} {{if .Proxy}} {{$preComma.PreComma}} `proxy`=:Proxy {{end}} {{if .ExtensionIds}} {{$preComma.PreComma}} `extension_ids`=:ExtensionIds {{end}} where `server_id`=:ServerID; {{end}}');
insert ignore into `api` (`api_id`,`title`,`description`,`method`,`route`,`template_ids`,`exec`,`input`,`output`,`input_tpl`,`output_tpl`) values('database_name-t_server-Update','database name 服务器部署 修改','database name 服务器部署 修改','POST','/api/database_name/v1/server/update','DatabaseNameServerUpdate','{{define "main"}}
{{execSQLTpl . "Update"}}
{{end}}','version=http://json-schema.org/draft-07/schema,id=input,direction=in
fullname=description,dst=Description,required
fullname=extensionIds,dst=ExtensionIds,required
fullname=proxy,dst=Proxy,required
fullname=serverId,dst=ServerID,required
fullname=serviceId,dst=ServiceID,required
fullname=url,dst=URL,required','version=http://json-schema.org/draft-07/schema,id=output,direction=out','{{define "input"}}
{{getSetValue . "Description" "input.description"}}
{{getSetValue . "ExtensionIds" "input.extensionIds"}}
{{getSetValue . "Proxy" "input.proxy"}}
{{getSetValue . "ServerID" "input.serverId"}}
{{getSetValue . "ServiceID" "input.serviceId"}}
{{getSetValue . "URL" "input.url"}}
{{end}}','{{define "output"}}
{{end}}');
insert ignore into `template` (`template_id`,`type`,`title`,`description`,`source_id`,`tpl`) values('DatabaseNameServerDel','SQL','database name 服务器部署 删除','database name 服务器部署 删除','database_name','{{define "Del"}} update `t_server` set `deleted_at`={{currentTime .}} where `server_id`=:ServerID; {{end}}');
insert ignore into `api` (`api_id`,`title`,`description`,`method`,`route`,`template_ids`,`exec`,`input`,`output`,`input_tpl`,`output_tpl`) values('database_name-t_server-Del','database name 服务器部署 删除','database name 服务器部署 删除','POST','/api/database_name/v1/server/del','DatabaseNameServerDel','{{define "main"}}
{{execSQLTpl . "Del"}}
{{end}}','version=http://json-schema.org/draft-07/schema,id=input,direction=in
fullname=serverId,dst=ServerID,required','version=http://json-schema.org/draft-07/schema,id=output,direction=out','{{define "input"}}
{{getSetValue . "ServerID" "input.serverId"}}
{{end}}','{{define "output"}}
{{end}}');
insert ignore into `template` (`template_id`,`type`,`title`,`description`,`source_id`,`tpl`) values('DatabaseNameServiceGetByServiceID','SQL','database name 项目/服务组件 获取 通过 ID','database name 项目/服务组件 获取 通过 ID','database_name','{{define "GetByServiceID"}} select * from `t_service` where `service_id`=:ServiceID and `deleted_at` is null; {{end}}');
insert ignore into `api` (`api_id`,`title`,`description`,`method`,`route`,`template_ids`,`exec`,`input`,`output`,`input_tpl`,`output_tpl`) values('database_name-t_service-GetByServiceID','database name 项目/服务组件 获取 通过 ID','database name 项目/服务组件 获取 通过 ID','POST','/api/database_name/v1/service/get_by_service_id','DatabaseNameServiceGetByServiceID','{{define "main"}}
{{execSQLTpl . "GetByServiceID"}}
{{end}}','version=http://json-schema.org/draft-07/schema,id=input,direction=in
fullname=serviceId,dst=ServiceID,required','version=http://json-schema.org/draft-07/schema,id=output,direction=out
fullname=contactIds,src=GetByServiceIDOut.#.contact_ids,required
fullname=createdAt,src=GetByServiceIDOut.#.created_at,required
fullname=deletedAt,src=GetByServiceIDOut.#.deleted_at,required
fullname=description,src=GetByServiceIDOut.#.description,required
fullname=license,src=GetByServiceIDOut.#.license,required
fullname=proxy,src=GetByServiceIDOut.#.proxy,required
fullname=security,src=GetByServiceIDOut.#.security,required
fullname=serviceId,src=GetByServiceIDOut.#.service_id,required
fullname=title,src=GetByServiceIDOut.#.title,required
fullname=updatedAt,src=GetByServiceIDOut.#.updated_at,required
fullname=variables,src=GetByServiceIDOut.#.variables,required
fullname=version,src=GetByServiceIDOut.#.version,required','{{define "input"}}
{{getSetValue . "ServiceID" "input.serviceId"}}
{{end}}','{{define "output"}}
{{getSetValue . "output.contactIds" "GetByServiceIDOut.#.contact_ids"}}
{{getSetValue . "output.createdAt" "GetByServiceIDOut.#.created_at"}}
{{getSetValue . "output.deletedAt" "GetByServiceIDOut.#.deleted_at"}}
{{getSetValue . "output.description" "GetByServiceIDOut.#.description"}}
{{getSetValue . "output.license" "GetByServiceIDOut.#.license"}}
{{getSetValue . "output.proxy" "GetByServiceIDOut.#.proxy"}}
{{getSetValue . "output.security" "GetByServiceIDOut.#.security"}}
{{getSetValue . "output.serviceId" "GetByServiceIDOut.#.service_id"}}
{{getSetValue . "output.title" "GetByServiceIDOut.#.title"}}
{{getSetValue . "output.updatedAt" "GetByServiceIDOut.#.updated_at"}}
{{getSetValue . "output.variables" "GetByServiceIDOut.#.variables"}}
{{getSetValue . "output.version" "GetByServiceIDOut.#.version"}}
{{end}}');
insert ignore into `template` (`template_id`,`type`,`title`,`description`,`source_id`,`tpl`) values('DatabaseNameServiceGetAllByServiceIDList','SQL','database name 项目/服务组件 获取 所有 通过 ID 列表','database name 项目/服务组件 获取 所有 通过 ID 列表','database_name','{{define "GetAllByServiceIDList"}} select * from `t_service` where `service_id` in ({{in . .ServiceIDList}}) and `deleted_at` is null; {{end}}');
insert ignore into `api` (`api_id`,`title`,`description`,`method`,`route`,`template_ids`,`exec`,`input`,`output`,`input_tpl`,`output_tpl`) values('database_name-t_service-GetAllByServiceIDList','database name 项目/服务组件 获取 所有 通过 ID 列表','database name 项目/服务组件 获取 所有 通过 ID 列表','POST','/api/database_name/v1/service/get_all_by_service_id_list','DatabaseNameServiceGetAllByServiceIDList','{{define "main"}}
{{execSQLTpl . "GetAllByServiceIDList"}}
{{end}}','version=http://json-schema.org/draft-07/schema,id=input,direction=in
fullname=serviceIdList,dst=ServiceIDList,required','version=http://json-schema.org/draft-07/schema,id=output,direction=out
fullname=contactIds,src=GetAllByServiceIDListOut.#.contact_ids,required
fullname=createdAt,src=GetAllByServiceIDListOut.#.created_at,required
fullname=deletedAt,src=GetAllByServiceIDListOut.#.deleted_at,required
fullname=description,src=GetAllByServiceIDListOut.#.description,required
fullname=license,src=GetAllByServiceIDListOut.#.license,required
fullname=proxy,src=GetAllByServiceIDListOut.#.proxy,required
fullname=security,src=GetAllByServiceIDListOut.#.security,required
fullname=serviceId,src=GetAllByServiceIDListOut.#.service_id,required
fullname=title,src=GetAllByServiceIDListOut.#.title,required
fullname=updatedAt,src=GetAllByServiceIDListOut.#.updated_at,required
fullname=variables,src=GetAllByServiceIDListOut.#.variables,required
fullname=version,src=GetAllByServiceIDListOut.#.version,required','{{define "input"}}
{{getSetValue . "ServiceIDList" "input.serviceIdList"}}
{{end}}','{{define "output"}}
{{getSetValue . "output.contactIds" "GetAllByServiceIDListOut.#.contact_ids"}}
{{getSetValue . "output.createdAt" "GetAllByServiceIDListOut.#.created_at"}}
{{getSetValue . "output.deletedAt" "GetAllByServiceIDListOut.#.deleted_at"}}
{{getSetValue . "output.description" "GetAllByServiceIDListOut.#.description"}}
{{getSetValue . "output.license" "GetAllByServiceIDListOut.#.license"}}
{{getSetValue . "output.proxy" "GetAllByServiceIDListOut.#.proxy"}}
{{getSetValue . "output.security" "GetAllByServiceIDListOut.#.security"}}
{{getSetValue . "output.serviceId" "GetAllByServiceIDListOut.#.service_id"}}
{{getSetValue . "output.title" "GetAllByServiceIDListOut.#.title"}}
{{getSetValue . "output.updatedAt" "GetAllByServiceIDListOut.#.updated_at"}}
{{getSetValue . "output.variables" "GetAllByServiceIDListOut.#.variables"}}
{{getSetValue . "output.version" "GetAllByServiceIDListOut.#.version"}}
{{end}}');
insert ignore into `template` (`template_id`,`type`,`title`,`description`,`source_id`,`tpl`) values('DatabaseNameServicePaginateWhere','SQL','database name 项目/服务组件 分页列表 where','database name 项目/服务组件 分页列表 where','database_name','{{define "PaginateWhere"}} {{end}}');
insert ignore into `template` (`template_id`,`type`,`title`,`description`,`source_id`,`tpl`) values('DatabaseNameServicePaginateTotal','SQL','database name 项目/服务组件 分页列表 总数','database name 项目/服务组件 分页列表 总数','database_name','{{define "PaginateTotal"}} select count(*) as `count` from `t_service` where 1=1 {{template "PaginateWhere" .}} and `deleted_at` is null; {{end}}');
insert ignore into `api` (`api_id`,`title`,`description`,`method`,`route`,`template_ids`,`exec`,`input`,`output`,`input_tpl`,`output_tpl`) values('database_name-t_service-PaginateTotal','database name 项目/服务组件 分页列表 总数','database name 项目/服务组件 分页列表 总数','POST','/api/database_name/v1/service/paginate_total','DatabaseNameServicePaginateTotal','{{define "main"}}
{{execSQLTpl . "PaginateTotal"}}
{{end}}','version=http://json-schema.org/draft-07/schema,id=input,direction=in','version=http://json-schema.org/draft-07/schema,id=output,direction=out','{{define "input"}}
{{end}}','{{define "output"}}
{{end}}');
insert ignore into `template` (`template_id`,`type`,`title`,`description`,`source_id`,`tpl`) values('DatabaseNameServicePaginate','SQL','database name 项目/服务组件 分页列表','database name 项目/服务组件 分页列表','database_name','{{define "Paginate"}} select * from `t_service` where 1=1 {{template "PaginateWhere" .}} and `deleted_at` is null order by `updated_at` desc limit :Offset,:Limit ; {{end}}');
insert ignore into `api` (`api_id`,`title`,`description`,`method`,`route`,`template_ids`,`exec`,`input`,`output`,`input_tpl`,`output_tpl`) values('database_name-t_service-Paginate','database name 项目/服务组件 分页列表','database name 项目/服务组件 分页列表','POST','/api/database_name/v1/service/paginate','DatabaseNameServicePaginateWhere,DatabaseNameServicePaginateTotal,DatabaseNameServicePaginate','{{define "main"}}
{{execSQLTpl . "PaginateTotal"}}
{{execSQLTpl . "Paginate"}}
{{end}}','version=http://json-schema.org/draft-07/schema,id=input,direction=in
fullname=limit,dst=Limit,format=number,required
fullname=offset,dst=Offset,format=number,required','version=http://json-schema.org/draft-07/schema,id=output,direction=out
fullname=contactIds,src=PaginateOut.#.contact_ids,required
fullname=createdAt,src=PaginateOut.#.created_at,required
fullname=deletedAt,src=PaginateOut.#.deleted_at,required
fullname=description,src=PaginateOut.#.description,required
fullname=license,src=PaginateOut.#.license,required
fullname=proxy,src=PaginateOut.#.proxy,required
fullname=security,src=PaginateOut.#.security,required
fullname=serviceId,src=PaginateOut.#.service_id,required
fullname=title,src=PaginateOut.#.title,required
fullname=updatedAt,src=PaginateOut.#.updated_at,required
fullname=variables,src=PaginateOut.#.variables,required
fullname=version,src=PaginateOut.#.version,required','{{define "input"}}
{{getSetNumber . "Limit" "input.limit"}}
{{getSetNumber . "Offset" "input.offset"}}
{{end}}','{{define "output"}}
{{getSetValue . "output.contactIds" "PaginateOut.#.contact_ids"}}
{{getSetValue . "output.createdAt" "PaginateOut.#.created_at"}}
{{getSetValue . "output.deletedAt" "PaginateOut.#.deleted_at"}}
{{getSetValue . "output.description" "PaginateOut.#.description"}}
{{getSetValue . "output.license" "PaginateOut.#.license"}}
{{getSetValue . "output.proxy" "PaginateOut.#.proxy"}}
{{getSetValue . "output.security" "PaginateOut.#.security"}}
{{getSetValue . "output.serviceId" "PaginateOut.#.service_id"}}
{{getSetValue . "output.title" "PaginateOut.#.title"}}
{{getSetValue . "output.updatedAt" "PaginateOut.#.updated_at"}}
{{getSetValue . "output.variables" "PaginateOut.#.variables"}}
{{getSetValue . "output.version" "PaginateOut.#.version"}}
{{end}}');
insert ignore into `template` (`template_id`,`type`,`title`,`description`,`source_id`,`tpl`) values('DatabaseNameServiceInsert','SQL','database name 项目/服务组件 新增','database name 项目/服务组件 新增','database_name','{{define "Insert"}} insert into `t_service` (`service_id`,`title`,`description`,`version`,`contact_ids`,`license`,`security`,`proxy`,`variables`)values (:ServiceID,:Title,:Description,:Version,:ContactIds,:License,:Security,:Proxy,:Variables); {{end}}');
insert ignore into `api` (`api_id`,`title`,`description`,`method`,`route`,`template_ids`,`exec`,`input`,`output`,`input_tpl`,`output_tpl`) values('database_name-t_service-Insert','database name 项目/服务组件 新增','database name 项目/服务组件 新增','POST','/api/database_name/v1/service/insert','DatabaseNameServiceInsert','{{define "main"}}
{{execSQLTpl . "Insert"}}
{{end}}','version=http://json-schema.org/draft-07/schema,id=input,direction=in
fullname=contactIds,dst=ContactIds,required
fullname=description,dst=Description,required
fullname=license,dst=License,required
fullname=proxy,dst=Proxy,required
fullname=security,dst=Security,required
fullname=serviceId,dst=ServiceID,required
fullname=title,dst=Title,required
fullname=variables,dst=Variables,required
fullname=version,dst=Version,required','version=http://json-schema.org/draft-07/schema,id=output,direction=out','{{define "input"}}
{{getSetValue . "ContactIds" "input.contactIds"}}
{{getSetValue . "Description" "input.description"}}
{{getSetValue . "License" "input.license"}}
{{getSetValue . "Proxy" "input.proxy"}}
{{getSetValue . "Security" "input.security"}}
{{getSetValue . "ServiceID" "input.serviceId"}}
{{getSetValue . "Title" "input.title"}}
{{getSetValue . "Variables" "input.variables"}}
{{getSetValue . "Version" "input.version"}}
{{end}}','{{define "output"}}
{{end}}');
insert ignore into `template` (`template_id`,`type`,`title`,`description`,`source_id`,`tpl`) values('DatabaseNameServiceUpdate','SQL','database name 项目/服务组件 修改','database name 项目/服务组件 修改','database_name','{{define "Update"}} {{$preComma:=newPreComma}} update `t_service` set {{if .ServiceID}} {{$preComma.PreComma}} `service_id`=:ServiceID {{end}} {{if .Title}} {{$preComma.PreComma}} `title`=:Title {{end}} {{if .Description}} {{$preComma.PreComma}} `description`=:Description {{end}} {{if .Version}} {{$preComma.PreComma}} `version`=:Version {{end}} {{if .ContactIds}} {{$preComma.PreComma}} `contact_ids`=:ContactIds {{end}} {{if .License}} {{$preComma.PreComma}} `license`=:License {{end}} {{if .Security}} {{$preComma.PreComma}} `security`=:Security {{end}} {{if .Proxy}} {{$preComma.PreComma}} `proxy`=:Proxy {{end}} {{if .Variables}} {{$preComma.PreComma}} `variables`=:Variables {{end}} where `service_id`=:ServiceID; {{end}}');
insert ignore into `api` (`api_id`,`title`,`description`,`method`,`route`,`template_ids`,`exec`,`input`,`output`,`input_tpl`,`output_tpl`) values('database_name-t_service-Update','database name 项目/服务组件 修改','database name 项目/服务组件 修改','POST','/api/database_name/v1/service/update','DatabaseNameServiceUpdate','{{define "main"}}
{{execSQLTpl . "Update"}}
{{end}}','version=http://json-schema.org/draft-07/schema,id=input,direction=in
fullname=contactIds,dst=ContactIds,required
fullname=description,dst=Description,required
fullname=license,dst=License,required
fullname=proxy,dst=Proxy,required
fullname=security,dst=Security,required
fullname=serviceId,dst=ServiceID,required
fullname=title,dst=Title,required
fullname=variables,dst=Variables,required
fullname=version,dst=Version,required','version=http://json-schema.org/draft-07/schema,id=output,direction=out','{{define "input"}}
{{getSetValue . "ContactIds" "input.contactIds"}}
{{getSetValue . "Description" "input.description"}}
{{getSetValue . "License" "input.license"}}
{{getSetValue . "Proxy" "input.proxy"}}
{{getSetValue . "Security" "input.security"}}
{{getSetValue . "ServiceID" "input.serviceId"}}
{{getSetValue . "Title" "input.title"}}
{{getSetValue . "Variables" "input.variables"}}
{{getSetValue . "Version" "input.version"}}
{{end}}','{{define "output"}}
{{end}}');
insert ignore into `template` (`template_id`,`type`,`title`,`description`,`source_id`,`tpl`) values('DatabaseNameServiceDel','SQL','database name 项目/服务组件 删除','database name 项目/服务组件 删除','database_name','{{define "Del"}} update `t_service` set `deleted_at`={{currentTime .}} where `service_id`=:ServiceID; {{end}}');
insert ignore into `api` (`api_id`,`title`,`description`,`method`,`route`,`template_ids`,`exec`,`input`,`output`,`input_tpl`,`output_tpl`) values('database_name-t_service-Del','database name 项目/服务组件 删除','database name 项目/服务组件 删除','POST','/api/database_name/v1/service/del','DatabaseNameServiceDel','{{define "main"}}
{{execSQLTpl . "Del"}}
{{end}}','version=http://json-schema.org/draft-07/schema,id=input,direction=in
fullname=serviceId,dst=ServiceID,required','version=http://json-schema.org/draft-07/schema,id=output,direction=out','{{define "input"}}
{{getSetValue . "ServiceID" "input.serviceId"}}
{{end}}','{{define "output"}}
{{end}}');
insert ignore into `template` (`template_id`,`type`,`title`,`description`,`source_id`,`tpl`) values('DatabaseNameAPIDel','SQL','database name 接口 删除','database name 接口 删除','database_name','{{define "Del"}} update `t_api` set `deleted_at`={{currentTime .}},`operator_id`=:OperatorIDInt,`operator`=:OperatorStr where `api_id`=:APIID; {{end}}');
insert ignore into `api` (`api_id`,`title`,`description`,`method`,`route`,`template_ids`,`exec`,`input`,`output`,`input_tpl`,`output_tpl`) values('database_name-t_api-Del','database name 接口 删除','database name 接口 删除','POST','/api/database_name/v1/api/del','DatabaseNameAPIDel,DatabaseNameAPIDel','{{define "main"}}
{{execSQLTpl . "Del"}}
{{execSQLTpl . "Del"}}
{{end}}','version=http://json-schema.org/draft-07/schema,id=input,direction=in
fullname=apiid,dst=APIID,required
fullname=operatorIdInt,dst=OperatorIDInt,format=number,required
fullname=operatorStr,dst=OperatorStr,required','version=http://json-schema.org/draft-07/schema,id=output,direction=out','{{define "input"}}
{{getSetValue . "APIID" "input.apiid"}}
{{getSetNumber . "OperatorIDInt" "input.operatorIdInt"}}
{{getSetValue . "OperatorStr" "input.operatorStr"}}
{{end}}','{{define "output"}}
{{end}}');
insert ignore into `template` (`template_id`,`type`,`title`,`description`,`source_id`,`tpl`) values('DatabaseNameAPIGetByAPIID','SQL','database name 接口 获取 通过 apiid','database name 接口 获取 通过 apiid','database_name','{{define "GetByAPIID"}} select * from `t_api` where `api_id`=:APIID and `deleted_at` is null; {{end}}');
insert ignore into `api` (`api_id`,`title`,`description`,`method`,`route`,`template_ids`,`exec`,`input`,`output`,`input_tpl`,`output_tpl`) values('database_name-t_api-GetByAPIID','database name 接口 获取 通过 apiid','database name 接口 获取 通过 apiid','POST','/api/database_name/v1/api/get_by_apiid','DatabaseNameAPIGetByAPIID','{{define "main"}}
{{execSQLTpl . "GetByAPIID"}}
{{end}}','version=http://json-schema.org/draft-07/schema,id=input,direction=in
fullname=apiid,dst=APIID,required','version=http://json-schema.org/draft-07/schema,id=output,direction=out
fullname=apiId,src=GetByAPIIDOut.#.api_id,required
fullname=createdAt,src=GetByAPIIDOut.#.created_at,required
fullname=deletedAt,src=GetByAPIIDOut.#.deleted_at,required
fullname=description,src=GetByAPIIDOut.#.description,required
fullname=name,src=GetByAPIIDOut.#.name,required
fullname=serviceId,src=GetByAPIIDOut.#.service_id,required
fullname=summary,src=GetByAPIIDOut.#.summary,required
fullname=tags,src=GetByAPIIDOut.#.tags,required
fullname=title,src=GetByAPIIDOut.#.title,required
fullname=updatedAt,src=GetByAPIIDOut.#.updated_at,required
fullname=uri,src=GetByAPIIDOut.#.uri,required','{{define "input"}}
{{getSetValue . "APIID" "input.apiid"}}
{{end}}','{{define "output"}}
{{getSetValue . "output.apiId" "GetByAPIIDOut.#.api_id"}}
{{getSetValue . "output.createdAt" "GetByAPIIDOut.#.created_at"}}
{{getSetValue . "output.deletedAt" "GetByAPIIDOut.#.deleted_at"}}
{{getSetValue . "output.description" "GetByAPIIDOut.#.description"}}
{{getSetValue . "output.name" "GetByAPIIDOut.#.name"}}
{{getSetValue . "output.serviceId" "GetByAPIIDOut.#.service_id"}}
{{getSetValue . "output.summary" "GetByAPIIDOut.#.summary"}}
{{getSetValue . "output.tags" "GetByAPIIDOut.#.tags"}}
{{getSetValue . "output.title" "GetByAPIIDOut.#.title"}}
{{getSetValue . "output.updatedAt" "GetByAPIIDOut.#.updated_at"}}
{{getSetValue . "output.uri" "GetByAPIIDOut.#.uri"}}
{{end}}');
insert ignore into `template` (`template_id`,`type`,`title`,`description`,`source_id`,`tpl`) values('DatabaseNameAPIGetAllByAPIIDList','SQL','database name 接口 获取 所有 通过 apiid 列表','database name 接口 获取 所有 通过 apiid 列表','database_name','{{define "GetAllByAPIIDList"}} select * from `t_api` where `api_id` in ({{in . .APIIDList}}) and `deleted_at` is null; {{end}}');
insert ignore into `api` (`api_id`,`title`,`description`,`method`,`route`,`template_ids`,`exec`,`input`,`output`,`input_tpl`,`output_tpl`) values('database_name-t_api-GetAllByAPIIDList','database name 接口 获取 所有 通过 apiid 列表','database name 接口 获取 所有 通过 apiid 列表','POST','/api/database_name/v1/api/get_all_by_apiid_list','DatabaseNameAPIGetAllByAPIIDList','{{define "main"}}
{{execSQLTpl . "GetAllByAPIIDList"}}
{{end}}','version=http://json-schema.org/draft-07/schema,id=input,direction=in
fullname=apiidList,dst=APIIDList,required','version=http://json-schema.org/draft-07/schema,id=output,direction=out
fullname=apiId,src=GetAllByAPIIDListOut.#.api_id,required
fullname=createdAt,src=GetAllByAPIIDListOut.#.created_at,required
fullname=deletedAt,src=GetAllByAPIIDListOut.#.deleted_at,required
fullname=description,src=GetAllByAPIIDListOut.#.description,required
fullname=name,src=GetAllByAPIIDListOut.#.name,required
fullname=serviceId,src=GetAllByAPIIDListOut.#.service_id,required
fullname=summary,src=GetAllByAPIIDListOut.#.summary,required
fullname=tags,src=GetAllByAPIIDListOut.#.tags,required
fullname=title,src=GetAllByAPIIDListOut.#.title,required
fullname=updatedAt,src=GetAllByAPIIDListOut.#.updated_at,required
fullname=uri,src=GetAllByAPIIDListOut.#.uri,required','{{define "input"}}
{{getSetValue . "APIIDList" "input.apiidList"}}
{{end}}','{{define "output"}}
{{getSetValue . "output.apiId" "GetAllByAPIIDListOut.#.api_id"}}
{{getSetValue . "output.createdAt" "GetAllByAPIIDListOut.#.created_at"}}
{{getSetValue . "output.deletedAt" "GetAllByAPIIDListOut.#.deleted_at"}}
{{getSetValue . "output.description" "GetAllByAPIIDListOut.#.description"}}
{{getSetValue . "output.name" "GetAllByAPIIDListOut.#.name"}}
{{getSetValue . "output.serviceId" "GetAllByAPIIDListOut.#.service_id"}}
{{getSetValue . "output.summary" "GetAllByAPIIDListOut.#.summary"}}
{{getSetValue . "output.tags" "GetAllByAPIIDListOut.#.tags"}}
{{getSetValue . "output.title" "GetAllByAPIIDListOut.#.title"}}
{{getSetValue . "output.updatedAt" "GetAllByAPIIDListOut.#.updated_at"}}
{{getSetValue . "output.uri" "GetAllByAPIIDListOut.#.uri"}}
{{end}}');
insert ignore into `template` (`template_id`,`type`,`title`,`description`,`source_id`,`tpl`) values('DatabaseNameAPIPaginateWhere','SQL','database name 接口 分页列表 where','database name 接口 分页列表 where','database_name','{{define "PaginateWhere"}} {{end}}');
insert ignore into `template` (`template_id`,`type`,`title`,`description`,`source_id`,`tpl`) values('DatabaseNameAPIPaginateTotal','SQL','database name 接口 分页列表 总数','database name 接口 分页列表 总数','database_name','{{define "PaginateTotal"}} select count(*) as `count` from `t_api` where 1=1 {{template "PaginateWhere" .}} and `deleted_at` is null; {{end}}');
insert ignore into `api` (`api_id`,`title`,`description`,`method`,`route`,`template_ids`,`exec`,`input`,`output`,`input_tpl`,`output_tpl`) values('database_name-t_api-PaginateTotal','database name 接口 分页列表 总数','database name 接口 分页列表 总数','POST','/api/database_name/v1/api/paginate_total','DatabaseNameAPIPaginateTotal','{{define "main"}}
{{execSQLTpl . "PaginateTotal"}}
{{end}}','version=http://json-schema.org/draft-07/schema,id=input,direction=in','version=http://json-schema.org/draft-07/schema,id=output,direction=out','{{define "input"}}
{{end}}','{{define "output"}}
{{end}}');
insert ignore into `template` (`template_id`,`type`,`title`,`description`,`source_id`,`tpl`) values('DatabaseNameAPIPaginate','SQL','database name 接口 分页列表','database name 接口 分页列表','database_name','{{define "Paginate"}} select * from `t_api` where 1=1 {{template "PaginateWhere" .}} and `deleted_at` is null order by `updated_at` desc limit :Offset,:Limit ; {{end}}');
insert ignore into `api` (`api_id`,`title`,`description`,`method`,`route`,`template_ids`,`exec`,`input`,`output`,`input_tpl`,`output_tpl`) values('database_name-t_api-Paginate','database name 接口 分页列表','database name 接口 分页列表','POST','/api/database_name/v1/api/paginate','DatabaseNameAPIPaginateWhere,DatabaseNameAPIPaginateTotal,DatabaseNameAPIPaginate','{{define "main"}}
{{execSQLTpl . "PaginateTotal"}}
{{execSQLTpl . "Paginate"}}
{{end}}','version=http://json-schema.org/draft-07/schema,id=input,direction=in
fullname=limit,dst=Limit,format=number,required
fullname=offset,dst=Offset,format=number,required','version=http://json-schema.org/draft-07/schema,id=output,direction=out
fullname=apiId,src=PaginateOut.#.api_id,required
fullname=createdAt,src=PaginateOut.#.created_at,required
fullname=deletedAt,src=PaginateOut.#.deleted_at,required
fullname=description,src=PaginateOut.#.description,required
fullname=name,src=PaginateOut.#.name,required
fullname=serviceId,src=PaginateOut.#.service_id,required
fullname=summary,src=PaginateOut.#.summary,required
fullname=tags,src=PaginateOut.#.tags,required
fullname=title,src=PaginateOut.#.title,required
fullname=updatedAt,src=PaginateOut.#.updated_at,required
fullname=uri,src=PaginateOut.#.uri,required','{{define "input"}}
{{getSetNumber . "Limit" "input.limit"}}
{{getSetNumber . "Offset" "input.offset"}}
{{end}}','{{define "output"}}
{{getSetValue . "output.apiId" "PaginateOut.#.api_id"}}
{{getSetValue . "output.createdAt" "PaginateOut.#.created_at"}}
{{getSetValue . "output.deletedAt" "PaginateOut.#.deleted_at"}}
{{getSetValue . "output.description" "PaginateOut.#.description"}}
{{getSetValue . "output.name" "PaginateOut.#.name"}}
{{getSetValue . "output.serviceId" "PaginateOut.#.service_id"}}
{{getSetValue . "output.summary" "PaginateOut.#.summary"}}
{{getSetValue . "output.tags" "PaginateOut.#.tags"}}
{{getSetValue . "output.title" "PaginateOut.#.title"}}
{{getSetValue . "output.updatedAt" "PaginateOut.#.updated_at"}}
{{getSetValue . "output.uri" "PaginateOut.#.uri"}}
{{end}}');
insert ignore into `template` (`template_id`,`type`,`title`,`description`,`source_id`,`tpl`) values('DatabaseNameAPIInsert','SQL','database name 接口 新增','database name 接口 新增','database_name','{{define "Insert"}} insert into `t_api` (`api_id`,`service_id`,`name`,`title`,`tags`,`uri`,`summary`,`description`)values (:APIID,:ServiceID,:Name,:Title,:Tags,:URI,:Summary,:Description); {{end}}');
insert ignore into `api` (`api_id`,`title`,`description`,`method`,`route`,`template_ids`,`exec`,`input`,`output`,`input_tpl`,`output_tpl`) values('database_name-t_api-Insert','database name 接口 新增','database name 接口 新增','POST','/api/database_name/v1/api/insert','DatabaseNameAPIInsert','{{define "main"}}
{{execSQLTpl . "Insert"}}
{{end}}','version=http://json-schema.org/draft-07/schema,id=input,direction=in
fullname=apiid,dst=APIID,required
fullname=description,dst=Description,required
fullname=name,dst=Name,required
fullname=serviceId,dst=ServiceID,required
fullname=summary,dst=Summary,required
fullname=tags,dst=Tags,required
fullname=title,dst=Title,required
fullname=uri,dst=URI,required','version=http://json-schema.org/draft-07/schema,id=output,direction=out','{{define "input"}}
{{getSetValue . "APIID" "input.apiid"}}
{{getSetValue . "Description" "input.description"}}
{{getSetValue . "Name" "input.name"}}
{{getSetValue . "ServiceID" "input.serviceId"}}
{{getSetValue . "Summary" "input.summary"}}
{{getSetValue . "Tags" "input.tags"}}
{{getSetValue . "Title" "input.title"}}
{{getSetValue . "URI" "input.uri"}}
{{end}}','{{define "output"}}
{{end}}');
insert ignore into `template` (`template_id`,`type`,`title`,`description`,`source_id`,`tpl`) values('DatabaseNameAPIUpdate','SQL','database name 接口 修改','database name 接口 修改','database_name','{{define "Update"}} {{$preComma:=newPreComma}} update `t_api` set {{if .APIID}} {{$preComma.PreComma}} `api_id`=:APIID {{end}} {{if .ServiceID}} {{$preComma.PreComma}} `service_id`=:ServiceID {{end}} {{if .Name}} {{$preComma.PreComma}} `name`=:Name {{end}} {{if .Title}} {{$preComma.PreComma}} `title`=:Title {{end}} {{if .Tags}} {{$preComma.PreComma}} `tags`=:Tags {{end}} {{if .URI}} {{$preComma.PreComma}} `uri`=:URI {{end}} {{if .Summary}} {{$preComma.PreComma}} `summary`=:Summary {{end}} {{if .Description}} {{$preComma.PreComma}} `description`=:Description {{end}} where `api_id`=:APIID; {{end}}');
insert ignore into `api` (`api_id`,`title`,`description`,`method`,`route`,`template_ids`,`exec`,`input`,`output`,`input_tpl`,`output_tpl`) values('database_name-t_api-Update','database name 接口 修改','database name 接口 修改','POST','/api/database_name/v1/api/update','DatabaseNameAPIUpdate','{{define "main"}}
{{execSQLTpl . "Update"}}
{{end}}','version=http://json-schema.org/draft-07/schema,id=input,direction=in
fullname=apiid,dst=APIID,required
fullname=description,dst=Description,required
fullname=name,dst=Name,required
fullname=serviceId,dst=ServiceID,required
fullname=summary,dst=Summary,required
fullname=tags,dst=Tags,required
fullname=title,dst=Title,required
fullname=uri,dst=URI,required','version=http://json-schema.org/draft-07/schema,id=output,direction=out','{{define "input"}}
{{getSetValue . "APIID" "input.apiid"}}
{{getSetValue . "Description" "input.description"}}
{{getSetValue . "Name" "input.name"}}
{{getSetValue . "ServiceID" "input.serviceId"}}
{{getSetValue . "Summary" "input.summary"}}
{{getSetValue . "Tags" "input.tags"}}
{{getSetValue . "Title" "input.title"}}
{{getSetValue . "URI" "input.uri"}}
{{end}}','{{define "output"}}
{{end}}');
insert ignore into `template` (`template_id`,`type`,`title`,`description`,`source_id`,`tpl`) values('DatabaseNameAPIDel','SQL','database name 接口 删除','database name 接口 删除','database_name','{{define "Del"}} update `t_api` set `deleted_at`={{currentTime .}} where `api_id`=:APIID; {{end}}');
insert ignore into `api` (`api_id`,`title`,`description`,`method`,`route`,`template_ids`,`exec`,`input`,`output`,`input_tpl`,`output_tpl`) values('database_name-t_api-Del','database name 接口 删除','database name 接口 删除','POST','/api/database_name/v1/api/del','DatabaseNameAPIDel,DatabaseNameAPIDel','{{define "main"}}
{{execSQLTpl . "Del"}}
{{execSQLTpl . "Del"}}
{{end}}','version=http://json-schema.org/draft-07/schema,id=input,direction=in
fullname=apiid,dst=APIID,required
fullname=operatorIdInt,dst=OperatorIDInt,format=number,required
fullname=operatorStr,dst=OperatorStr,required','version=http://json-schema.org/draft-07/schema,id=output,direction=out','{{define "input"}}
{{getSetValue . "APIID" "input.apiid"}}
{{getSetNumber . "OperatorIDInt" "input.operatorIdInt"}}
{{getSetValue . "OperatorStr" "input.operatorStr"}}
{{end}}','{{define "output"}}
{{end}}');