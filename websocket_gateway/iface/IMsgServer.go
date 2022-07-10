package iface

import (
	"github.com/changan/websocket_gateway/model"
)

/*
	消息转发服务接口
*/

type IMsgServer interface {
	Connect() (IMsgServer, error)
	SendMsg(message model.ImMessage) (bool, error)
	SendMsgId(id uint32, message model.ImMessage) (bool, error)
	Close() error
}
