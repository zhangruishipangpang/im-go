package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/changan/websocket_gateway/message"
	"github.com/changan/websocket_gateway/register"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"strconv"
	"testing"
	"unsafe"
)

func TestNullPath(t *testing.T) {
	//point()
	//p1()
	//p2()
	//tNacos()
}

func tNacos() {
	nacos := register.Nacos{}
	nacos.RegisterInstance()
	instance, err := nacos.Client.SelectOneHealthyInstance(vo.SelectOneHealthInstanceParam{
		//Clusters:    nil,
		ServiceName: "nacos.test.3",
		//GroupName:   "",
	})
	if err != nil {
		return
	}

	fmt.Println(" 获取到的实例信息： ", instance)
	fmt.Println(instance.Ip + ":" + strconv.Itoa(int(instance.Port)))
}

func p1() {
	a := 1
	e := unsafe.Pointer(&a)
	fmt.Println(" e ", e)
}

func p2() {
	msg := message.NewMessage(1, 2, " PING ")
	marshal, err := json.Marshal(msg)
	if err != nil {
		return
	}
	fmt.Println("json : ", string(marshal))
}

func point() {
	msg := M1{
		From: 1,
		Dest: 2,
		Data: 3,
	}
	buf := &bytes.Buffer{}
	err := binary.Write(buf, binary.BigEndian, msg)
	if err != nil {
		panic(err)
	}
	fmt.Println(buf.Bytes())
	fmt.Println(" data : ", string(buf.Bytes()))
}

func point2() {
	msg := M1{
		From: 1,
		Dest: 2,
		Data: 3,
	}
	buf := &bytes.Buffer{}
	err := binary.Write(buf, binary.BigEndian, msg)
	if err != nil {
		panic(err)
	}
	fmt.Println(buf.Bytes())
	fmt.Println(" data : ", string(buf.Bytes()))
}

type M1 struct {
	From uint32 `json:"from"` // 消息来源
	Dest uint32 `json:"dest"` // 消息目标
	Data uint32 `json:"data"` // 消息体
}

type M2 struct {
	From uint32 `json:"from"` // 消息来源
	Dest uint32 `json:"dest"` // 消息目标
	Data string `json:"data"` // 消息体
}

/*

func (auth *Auth) checkToken(token string) bool {
	body := TokenBody{token: "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c111112VyX25hbWUiOiJ6aGFuZ3NhbiIsInNjb3BlIjpbImFsbCJdLCJTVXNlciI6eyJwYXNzd29yZCI6bnVsbCwidXNlcm5hbWUiOiJ6aGFuZ3NhbiIsImF1dGhvcml0aWVzIjpbeyJhdXRob3JpdHkiOiJhZG1pbiJ9XSwiYWNjb3VudE5vbkV4cGlyZWQiOnRydWUsImFjY291bnROb25Mb2NrZWQiOnRydWUsImNyZWRlbnRpYWxzTm9uRXhwaXJlZCI6dHJ1ZSwiZW5hYmxlZCI6dHJ1ZX0sImV4cCI6MTY0ODY5MTY5NSwiYXV0aG9yaXRpZXMiOlsiYWRtaW4iXSwianRpIjoiYTExNDE5ODgtODcyYy00NTE2LTk5YWMtMjQ3MDVlZjRmNWI0IiwiY2xpZW50X2lkIjoiYmFpZHUifQ.G32br6PTm_EEW6R02l_Ws8SRqt_o8fl59ATwGqPdd5tSSJeWKMVdebuQWvMHy5HzzUO7cl6rwtYoflh9UgwhQnNY_ivYGM_SZVGLss1625WS87fySlEn17hNnkU461bxdJKSBIzKrDE-zfJ-Ody_RggaTcNPGpZVeE2ectHy9ZUUPQB_aUznaWZxolkN2eZ1wgb2Ruq7APsqqOr70IFs55c8mCsmhlbd8yBenEPLwJP2igyGcPRsmKQqdEnsqk1NjIBXh_UPVwlXbPVOa0Y-Q5GzTwPAYtoMIDue-G3U1kBNVwDNUiPVGRN7ZlG92Ivk5CFe5td_WjHGXLj0WWLzKw"}
	responseBody := http.InvokeRequestFromServiceName("auth-serve", "/auth/check_token", body, nil)
	return responseBody != (http.ResponseBody{}) && responseBody.Code == 200
}
*/
