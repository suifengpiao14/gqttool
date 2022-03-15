
{{define "Del"}}
update `[[.TableName]]` set `deleted_at`={{currentTime .}},`operator_id`=:OperatorID,`operator`=:Operator where `[[.PrimaryKey]]`=:[[.PrimaryKeyCamel]];
{{end}}
