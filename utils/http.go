package utils

import (
	"errors"
	"github.com/wujiangweiphp/go-curl"
	"log"
)

func HttpPost(url string, queries map[string]string, postData map[string]interface{}) (string, error) {
	headers := map[string]string{
		//"User-Agent": "Sublime",
		"Authorization": "Bearer YWMtuGb8shZdEe62bvMpz9tRAMMRjWxrrzJpkpy6JDac4BcIRaIyoFpHz5XZkEfP6cRvAgMAAAGJBnEsyzeeSAAs6umKMBMofAOmA6e9eXbeyro5QNdWzpSdlylyJ6RpNg",
		"Content-Type":  "application/json",
		"Accept":        "application/json",
	}
	cookies := map[string]string{
		//"userId": "12",
		//"loginTime": "15045682199",
	}
	// 链式操作
	req := curl.NewRequest()
	resp, err := req.
		SetUrl(url).
		SetHeaders(headers).
		SetCookies(cookies).
		SetQueries(queries).
		SetPostData(postData).
		Post()
	if err != nil {
		return "", err
	} else {
		if resp.IsOk() {
			return resp.Body, nil
		} else {
			log.Printf("%v\n", resp.Raw)
			return "", errors.New("请求失败")
		}
	}
}
