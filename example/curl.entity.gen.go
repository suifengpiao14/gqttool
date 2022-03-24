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

		func (t *CurlOrderCurlGetOrderByOrderNumberBodyEntity) TplOutput(entity gqttpl.TplEntityInterface ) (out string,err error){
			out,err=gqttpl.TplOutput(entity,t.TplName())
			return 
		}

	

		type CurlOrderCurlGetOrderByOrderNumberEntity struct{
			
				GetOrderByOrderNumberBody CurlOrderCurlGetOrderByOrderNumberBodyEntity 
			
				 CurlOrderCurlGetOrderByOrderNumberBodyEntity 
			
				SecretKey interface{} 
			
				ServiceID interface{} 
			
			gqttpl.TplEmptyEntity
		}

		func (t *CurlOrderCurlGetOrderByOrderNumberEntity) TplName() string{
			return "curl.order.curl.GetOrderByOrderNumber"
		}

		func (t *CurlOrderCurlGetOrderByOrderNumberEntity) TplOutput(entity gqttpl.TplEntityInterface ) (out string,err error){
			out,err=gqttpl.TplOutput(entity,t.TplName())
			return 
		}

	