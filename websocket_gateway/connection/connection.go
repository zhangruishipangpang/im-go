package connection

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/changan/websocket_gateway/message"
	"github.com/changan/websocket_gateway/user"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

var upgrader = websocket.Upgrader{
	// 这里先配置所有过来的链接都放行
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Connection struct {
	conn *websocket.Conn // webSocket链接

	ctx context.Context // 全局控制

	connId uint32 // 链接id

	user user.User // 用户

	property map[string]interface{}
}

func NewConnection(rw http.ResponseWriter, r *http.Request, user user.User) *Connection {
	connection := &Connection{
		ctx:  context.Background(),
		user: user,
	}

	fmt.Println(" 来链接了... ")
	// 创建链接
	connW, err := upgrader.Upgrade(rw, r, nil)
	if err != nil {
		fmt.Println(" 链接异常： ", err)
		return nil
	}

	connection.conn = connW
	return connection
}

// Start 启动一个go协程来处理消息
func (conn *Connection) Start() {
	log.Printf(" 获取一个新链接，remote : %s ", conn.conn.RemoteAddr())

	for {
		_, err := conn.ReadMsg()
		if err != nil {
			log.Print(" 读取异常1：", err)
			return
		}

		time.Sleep(5 * time.Second)
		// 链接成功发送一个ping通知
		sendMsg := message.NewMessage(1, 1, " PING ")
		conn.SendMsg(sendMsg)
	}
}

// {"from":1, "dest":2}

// SendMsg 发送消息
func (conn *Connection) SendMsg(msg *message.Message) {
	msgJ, err := json.Marshal(msg)
	if err != nil {
		return
	}
	fmt.Println("json : ", string(msgJ))

	err = conn.conn.WriteMessage(websocket.TextMessage, []byte(msgJ))
	if err != nil {
		log.Printf(" 发送消息异常 : %s ", err)
		return
	}
}

// ReadMsg 读取消息
func (conn *Connection) ReadMsg() (message.Message, error) {
	mt, msg, err := conn.conn.ReadMessage()
	if err != nil {
		log.Printf(" 读取数据异常：{%s} ", err.Error())
		return message.Message{}, nil
	}
	if mt != websocket.TextMessage {
		log.Printf(" 消息类型只支持文本 ")
		return message.Message{}, errors.New(" 消息类型只支持文本, 本次消息类型：" + string(mt))
	}
	log.Printf(" 读取数据：{%s} ", string(msg))

	// 读取消息
	var msgStruct message.Message
	err = json.Unmarshal(msg, &msgStruct)
	if err != nil {
		return message.Message{}, err
	}

	return msgStruct, nil
}

//

// Close 关闭链接
func (conn *Connection) Close() error {
	err := conn.conn.Close()
	if err != nil {
		log.Printf(" 关闭链接异常 : %s ", err)
		return err
	}
	return nil
}
