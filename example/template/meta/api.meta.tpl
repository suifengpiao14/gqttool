
[[define "Del"]]
{{define "Del"}}
update `[[.TableName]]` set `deleted_at`={{currentTime .}},`operator_id`=:OperatorIDInt,`operator`=:OperatorStr where `[[.PrimaryKey]]`=:[[.PrimaryKeyCamel]];
{{end}}
[[end]]
