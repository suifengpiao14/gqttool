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
{{- $body:=jsonCompact ( tplOutput . .GetOrderByOrderNumberBody ) -}}
POST http://ordserver.huishoubao.com/order_center/getOrderInfo HTTP/1.1
Content-Type: application/json
HSB-OPENAPI-CALLERSERVICEID: {{.ServiceID}}
HSB-OPENAPI-SIGNATURE: {{getMD5LOWER  $body "_" .SecretKey}}

{{$body}}
{{end}}

{{define "GetOrderByOrderNumberResponse"}}
HTTP/1.1 200 OK
Content-Type: text/json; charset=UTF-8

OrderID         string  `json:"orderId" njson:"_data._data.basic.orderId" `
OrderNumber     string  `json:"orderNumber" njson:"_data._data.basic.orderNum" `
ArrivalType     string  `json:"arrivalType" njson:"_data._data.basic.arrivalType" `
AccessBusiness  string  `json:"accessBusiness" njson:"_data._data.basic.accessBusiness" `
AppointmentTime string  `json:"appointmentTime" njson:"fullname:_data._data.basic.customData.appointmentTime;format:json;formatPath:_data._data.basic.customData"`
CityID          string  `json:"cityId" njson:"_data._data.basic.cityId"`
City            string  `json:"city" njson:"_data._data.basic.city"`
Latitude        float64 `json:"latitude" njson:"_data._data.basic.customData.latitude"`
Longitude       float64 `json:"longitude" njson:"_data._data.basic.customData.longitude"`
ChannelID       string  `json:"channelId" njson:"_data._data.channel.ChannelID"`
ChannelName     string  `json:"channelName" njson:"_data._data.channel.channelName"`
ProductID       string  `json:"productId" njson:"_data._data.good.productId"`
ProductName     string  `json:"productName" njson:"_data._data.good.productName"`
Picture         string  `json:"productPic" njson:"_data._data.good.productPic"`
Quotation       string  `json:"price" njson:"_data._data.basic.quotation"`      {{// 用户预估价}}
Address         string  `json:"address" njson:"_data._data.customData.address"` {{//上门地址}}
BrandId         string  `json:"brandId" njson:"_data._data.good.brandId"`
Business        string  `json:"business" njson:"_data._data.good.brandId"`
{{end}}

