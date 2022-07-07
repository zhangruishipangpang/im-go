package http

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

/*
	Http请求工具类
*/

type HttpProxy struct {
	Http http.Client
}

func InvokeRequest(url, method string, body io.Reader, header map[string]string) {

	request, err := http.NewRequest(method, url, body)
	if err != nil {
		panic(" 创建HTTP Request失败，错误信息： " + err.Error())
		return
	}
	for key, value := range header {
		if len(key) > 0 {
			request.Header.Set(key, value)
		}
	}

	c := http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := c.Do(request)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	RB, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(" 请求结果 ： ", string(RB))
}
