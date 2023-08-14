package utils

import (
	"github.com/valyala/fasthttp"
	"xx/global"
)

type IFastHttp interface {
	Http() string
}

type FastHttpManager struct {
	url    string
	method string
	params string
}

func FastHttp(url string, method string, params string) IFastHttp {
	return &FastHttpManager{url: url, method: method, params: params}
}

func (h *FastHttpManager) Http() string {

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.Header.SetContentType("application/json")
	req.Header.SetMethod(h.method)

	req.SetRequestURI(h.url)
	requestBody := []byte(h.params)
	req.SetBody(requestBody)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)
	if err := fasthttp.Do(req, resp); err != nil {
		global.SugarLogger.Error("请求失败:" + err.Error())
		return ""
	}

	return string(resp.Body())
}
