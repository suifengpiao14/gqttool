package example
import "github.com/suifengpiao14/gqt/v2/gqttpl"

		type CurlOrderCurlGetOrderByOrderNumberBodyEntity struct{
			
				OrderNumber interface{} 
			
				ServiceID interface{} 
			
			gqttpl.TplEmptyEntity
		}

		func (t *CurlOrderCurlGetOrderByOrderNumberBodyEntity) TplName() string{
			return "curl.order.curl.GetOrderByOrderNumberBody"
		}
	

		type CurlOrderCurlGetOrderByOrderNumberEntity struct{
			
				GetOrderByOrderNumberBody *CurlOrderCurlGetOrderByOrderNumberBodyEntity 
			
				 CurlOrderCurlGetOrderByOrderNumberBodyEntity 
			
				SecretKey interface{} 
			
				ServiceID interface{} 
			
			gqttpl.TplEmptyEntity
		}

		func (t *CurlOrderCurlGetOrderByOrderNumberEntity) TplName() string{
			return "curl.order.curl.GetOrderByOrderNumber"
		}
	