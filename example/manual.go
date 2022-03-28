package example

import (
	"net/http"

	"github.com/pkg/errors"
)

type ResponseStatus struct {
	HttpStatus      int    `json:"httpstatus"`
	BusinessCode    string `josn:"businessCode" njson:"_data._retcode"`
	BusinessMessage string `josn:"businessMessage" njson:"_data._retinfo"`
	request         *http.Request
}

func (r *ResponseStatus) GetHttpStatus() int {
	return r.HttpStatus
}
func (r *ResponseStatus) SetHttpStatus(httpstatus int) {
	r.HttpStatus = httpstatus
}

func (r *ResponseStatus) GetBusinessMessage() string {
	return r.BusinessMessage
}

func (r *ResponseStatus) GetBusinessCode() string {
	return r.BusinessCode
}

func (r *ResponseStatus) SetRequest(req *http.Request) {
	r.request = req
	return
}

func (r *ResponseStatus) GetBusinessError() (err error) {
	businessCode := r.GetBusinessCode()
	if !(businessCode == "" || businessCode == "0") {
		url := ""
		if r.request != nil {
			url = r.request.URL.String()
		}
		err = errors.Errorf("api :%s response business code:%s,business message:%s", url, businessCode, r.GetBusinessMessage())
		return err
	}
	return
}

type ResponseInterface interface {
	GetBusinessCode() string
	GetBusinessMessage() string
	GetHttpStatus() int
	SetHttpStatus(httpstatus int)
	GetBusinessError() error
	SetRequest(r *http.Request)
}
