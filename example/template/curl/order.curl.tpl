{{define "GetOrderByOrderNumberBody"}}
{
    "_head":{
        "_interface":"getOrderInfo",
        "_msgType":"request",
        "_remark":"",
        "_version":"0.01",
        "_timestamps":"{{timestampSecond}}",
        "_invokeId":"dispatch_order_{{xid}}",
        "_callerServiceId":"{{.ServiceID}}",
        "_groupNo":"1"
    },
    "_param":{
        "orderNum":"{{.OrderNumber}}",
        "containInfo":[
            "basic",
            "good",
            "logistics"
        ]
    }
}
{{end}}

{{define "GetOrderByOrderNumber"}}
{{- $body:=jsonCompact ( .GetOrderByOrderNumberBody.TplOutput *) -}}
POST http://ordserver.huishoubao.com/order_center/getOrderInfo HTTP/1.1
Content-Type: application/json
HSB-OPENAPI-CALLERSERVICEID: {{.ServiceID}}
HSB-OPENAPI-SIGNATURE: {{getMD5LOWER  $body "_" .SecretKey}}

{{$body}}
{{end}}