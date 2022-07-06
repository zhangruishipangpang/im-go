package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/changan/websocket_gateway/message"
	"testing"
	"unsafe"
)

func TestNullPath(t *testing.T) {
	//point()
	//p1()
	p2()
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
	msg := message.NewMessage(1, 2, " PING ")
	buf := &bytes.Buffer{}
	err := binary.Write(buf, binary.BigEndian, msg)
	if err != nil {
		panic(err)
	}
	fmt.Println(buf.Bytes())
	fmt.Println(" data : ", string(buf.Bytes()))

}
