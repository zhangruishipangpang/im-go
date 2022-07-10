package mserver

import (
	"encoding/json"
	"fmt"
	"github.com/aceld/zinx/znet"
	"github.com/changan/websocket_gateway/iface"
	"github.com/changan/websocket_gateway/model"
	"net"
)

/*
	使用Zinx 框架作为消息转发服务
*/

type ZinxServer struct {
	Network, Address string
	conn             net.Conn
}

func (server *ZinxServer) Connect() (iface.IMsgServer, error) {
	conn, err := net.Dial(server.Network, server.Address)
	if err != nil {
		fmt.Println("client start err, exit!", err)
		return nil, err
	}
	server.conn = conn
	return server, nil
}

func (server ZinxServer) SendMsg(message model.ImMessage) (bool, error) {
	//发封包message消息
	dp := znet.NewDataPack()
	bytes, err := json.Marshal(message)
	if err != nil {
		return false, err
	}
	msg, _ := dp.Pack(znet.NewMsgPackage(10, bytes))
	_, errW := server.conn.Write(msg)
	if errW != nil {
		fmt.Println("write error err ", errW)
		return false, errW
	}
	return true, nil
}

func (server ZinxServer) SendMsgId(id uint32, message model.ImMessage) (bool, error) {
	//发封包message消息
	dp := znet.NewDataPack()
	bytes, err := json.Marshal(message)
	if err != nil {
		return false, err
	}
	msg, _ := dp.Pack(znet.NewMsgPackage(id, bytes))
	_, errW := server.conn.Write(msg)
	if errW != nil {
		fmt.Println("write error err ", errW)
		return false, errW
	}
	return true, nil
}

func (server ZinxServer) Close() error {
	err := server.conn.Close()
	return err
}
