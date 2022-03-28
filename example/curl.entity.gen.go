package example

import "github.com/suifengpiao14/gqt/v2/gqttpl"

type CurlOrderCurlGetOrderByOrderNumberBodyEntity struct {
	OrderNumber interface{}

	ServiceID interface{}

	gqttpl.TplEmptyEntity
}

func (t *CurlOrderCurlGetOrderByOrderNumberBodyEntity) TplName() string {
	return "curl.order.curl.GetOrderByOrderNumberBody"
}

func (t *CurlOrderCurlGetOrderByOrderNumberBodyEntity) TplType() string {
	return "text"
}

type CurlOrderCurlGetOrderByOrderNumberEntity struct {
	GetOrderByOrderNumberBody *CurlOrderCurlGetOrderByOrderNumberBodyEntity

	CurlOrderCurlGetOrderByOrderNumberBodyEntity

	SecretKey interface{}

	ServiceID interface{}

	gqttpl.TplEmptyEntity
}

func (t *CurlOrderCurlGetOrderByOrderNumberEntity) TplName() string {
	return "curl.order.curl.GetOrderByOrderNumber"
}

func (t *CurlOrderCurlGetOrderByOrderNumberEntity) TplType() string {
	return "curl_request"
}

type CurlOrderCurlGetOrderByOrderNumberResponseEntity struct {
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
	Quotation       string  `json:"price" njson:"_data._data.basic.quotation"`      // 用户预估价
	Address         string  `json:"address" njson:"_data._data.customData.address"` //上门地址
	BrandId         string  `json:"brandId" njson:"_data._data.good.brandId"`
	Business        string  `json:"business" njson:"_data._data.good.brandId"`
	ResponseStatus

	gqttpl.TplEmptyEntity
}

func (t *CurlOrderCurlGetOrderByOrderNumberResponseEntity) TplName() string {
	return "curl.order.curl.GetOrderByOrderNumberResponse"
}

func (t *CurlOrderCurlGetOrderByOrderNumberResponseEntity) TplType() string {
	return "curl_response"
}
