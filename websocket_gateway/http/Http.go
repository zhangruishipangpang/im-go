package http

import (
	"bytes"
	"encoding/json"
	"github.com/changan/websocket_gateway/model"
	"github.com/changan/websocket_gateway/register"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

/*
	Http请求工具类
*/

type HttpProxy struct {
	Http http.Client
}

//// ResponseBody 默认的返回类型
//type ResponseBody struct {
//	Code uint32      `json:"code"`
//	Msg  string      `json:"msg"`
//	E    interface{} `json:"e"`
//}

// InvokeRequest 发起HTTP请求，这里只支持post json请求
func InvokeRequest(url, method string, body interface{}, header map[string]string) []byte {
	if method != "POST" {
		panic(" 目前只支持POST请求 ")
	}
	bodyJson, errJson := json.Marshal(body)
	if errJson != nil {
		panic(" 请求参数体转换JSON异常，异常信息:" + errJson.Error())
	}

	request, err := http.NewRequest(method, url, bytes.NewReader(bodyJson))
	if err != nil {
		panic(" 创建HTTP Request失败，错误信息： " + err.Error())
	}
	// 添加参数Json请求头
	request.Header.Add("content-type", "application/json;charset=UTF-8")
	for key, value := range header {
		if len(key) > 0 {
			request.Header.Add(key, value)
		}
	}

	c := http.Client{
		Timeout: 100 * time.Second,
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
	//fmt.Println(" 请求结果 ： ", string(RB))
	//fmt.Println(" = ", resp.Request.URL.String())
	return RB
}

func InvokeRequestFromServiceName(serviceName, path string, body interface{}, header map[string]string) model.ResponseBody {
	client := register.GetNacosClient()
	instance, err := client.SelectOneHealthyInstance(vo.SelectOneHealthInstanceParam{ServiceName: serviceName})
	if err != nil {
		panic(err)
	}
	host := "http://" + instance.Ip + ":" + strconv.Itoa(int(instance.Port)) + path
	resp := InvokeRequest(host, http.MethodPost, body, header)
	rb := &model.ResponseBody{}
	errJson := json.Unmarshal(resp, rb)
	if errJson != nil {
		panic(errJson)
	}
	return *rb
}
