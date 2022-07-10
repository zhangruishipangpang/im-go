package iface

import (
	"github.com/changan/websocket_gateway/message"
)

/*
	链接的持有者，用户<----->wg
	链接的持有者需要有如下的方法提供：
		* 1、发送消息
		* 2、读取消息
		* 3、关闭链接
*/

type IConnection interface {
	Start() // 开始链接

	SendMsg(*message.Message)          // 发送消息
	ReadMsg() (message.Message, error) // 读取消息

	Close() error // 关闭链接
}
