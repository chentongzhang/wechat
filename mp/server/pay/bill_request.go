// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/magicshui/wechat for the canonical source repository
// @license     https://github.com/magicshui/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay

import (
	"encoding/xml"
	"errors"
	"github.com/magicshui/wechat/mp/pay/native"
	"net/http"
)

// native api 请求订单详情的 Handler
type BillRequestHandler struct {
	paySignKey     string
	invalidHandler InvalidRequestHandlerFunc
	requestHandler BillRequestHandlerFunc
}

// NOTE: 所有参数必须有效
func NewBillRequestHandler(
	paySignKey string,
	invalidHandler InvalidRequestHandlerFunc,
	requestHandler BillRequestHandlerFunc,

) (handler *BillRequestHandler) {

	if paySignKey == "" {
		panic(`paySignKey == ""`)
	}
	if invalidHandler == nil {
		panic("invalidHandler == nil")
	}
	if requestHandler == nil {
		panic("requestHandler == nil")
	}

	handler = &BillRequestHandler{
		paySignKey:     paySignKey,
		invalidHandler: invalidHandler,
		requestHandler: requestHandler,
	}

	return
}

// BillRequestHandler 实现 http.Handler 接口
func (handler *BillRequestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		err := errors.New("request method is not POST")
		handler.invalidHandler(w, r, err)
		return
	}

	var billReq native.BillRequest
	if err := xml.NewDecoder(r.Body).Decode(&billReq); err != nil {
		handler.invalidHandler(w, r, err)
		return
	}

	if err := billReq.Check(handler.paySignKey); err != nil {
		handler.invalidHandler(w, r, err)
		return
	}

	handler.requestHandler(w, r, &billReq)
}
