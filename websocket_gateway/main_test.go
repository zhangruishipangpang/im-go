package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/changan/websocket_gateway/message"
	"github.com/changan/websocket_gateway/register_center"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"testing"
	"unsafe"
)

func TestNullPath(t *testing.T) {
	//point()
	//p1()
	//p2()
	tNacos()
}

func tNacos() {
	nacos := register_center.Nacos{}
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
