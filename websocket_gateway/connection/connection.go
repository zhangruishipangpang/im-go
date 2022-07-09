package connection

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/changan/websocket_gateway/auth"
	"github.com/changan/websocket_gateway/message"
	"github.com/changan/websocket_gateway/redis"
	"github.com/changan/websocket_gateway/user"
	"github.com/gorilla/websocket"
	jsoniter "github.com/json-iterator/go"
	"log"
	"net/http"
	"sync"
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

	isLogin bool
}

var count int32 = 0
var lock sync.RWMutex = sync.RWMutex{}

func increase() {
	lock.Lock()
	defer lock.Unlock()
	count++
}

// NewConnection 创建一个链接结构体
func NewConnection(rw http.ResponseWriter, r *http.Request, user user.User) *Connection {
	log.Printf(" 请求创建ws链接! ")

	connection := &Connection{
		ctx:  context.Background(),
		user: user,
	}

	connW, err := upgrader.Upgrade(rw, r, nil)
	if err != nil {
		fmt.Println(" 链接异常： ", err)
		return nil
	}

	increase()
	fmt.Println("当前连接数：", count)
	connection.conn = connW
	return connection
}

// Start 启动一个go协程来处理消息
func (conn *Connection) Start() {
	log.Printf(" 获取一个新链接，remote : %s ", conn.conn.RemoteAddr())

	defer conn.Close()
	for {
		msg, err := conn.ReadMsg()
		if err != nil {
			log.Print(" 读取异常1：", err)
			return
		}

		if msg.DataType == message.TOKEN {
			if conn.isLogin {
				sendMsg := message.NewMessage(msg.Dest, msg.From, "已经登录无需重复发送认证信息！")
				conn.SendMsg(sendMsg)
				continue
			}
			token := msg.Data
			authService := auth.AuthService{}
			checkToken, errMsg := authService.CheckToken(token)
			if !checkToken {
				sendMsg := message.NewMessage(msg.Dest, msg.From, errMsg)
				conn.SendMsg(sendMsg)
				panic(errMsg)
			}
			conn.isLogin = true
			redisClient := redis.GetRedisClient()
			authUserMsg := authService.AuthUserMsg("")
			exit := redisClient.Get("USER:MSG:" + authUserMsg.UserNumber)
			if exit != "" {
				log.Println("用户已经成功认证，无需在此认证")
				continue
			}
			userJson, err := jsoniter.MarshalToString(authUserMsg)
			if err != nil {
				return
			}
			redisClient.Set("USER:MSG:"+authUserMsg.UserNumber, string(userJson))
		}

		time.Sleep(2 * time.Second)
		// 链接成功发送一个ping通知
		sendMsg := message.NewMessage(1, 1, " PING ")
		conn.SendMsg(sendMsg)
	}
}

// {"from":1, "dest":2, "data_type":3, "data":"eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOlsiYWxsIl0sInVzZXJfZW1haWwiOiIxODg0NjQzOTk1MkAxNjMuY29tIiwidXNlcl9tb2JpbGUiOiIxODg0NjQzOTk1MiIsInVzZXJfbmFtZSI6ImNoYW5nYW4iLCJzY29wZSI6WyJhbGwiXSwiZXhwIjoxNjU3MzgxNDE1LCJhdXRob3JpdGllcyI6WyJlbXBsb3llZSJdLCJqdGkiOiIzMjNkZmE1Yy1kYjJkLTQ5YjctODc1OS1lNzQzMjk3OWFjYTYiLCJjbGllbnRfaWQiOiJnYXRld2F5IiwidXNlcl9udW1iZXIiOiJjMzVkNDcxZjNlYjhjYjYyMTYwNTRmYzVhMmIwOGQ2NyJ9.C0Z6z_eUIoY0u3hXW3maz3ZAlCtAE20A6l07jAuITdCxyUfErH7iLOipOWCWir73FBciMx4TZrP4zMgfXLVEQw6s93FrheikZ_k_-tv8ixiKbDZcscdEy2LqDKYNT_zr9NGcSWG-q5lpGIo3mFZ0cb-QOOOsIBr281_rVtyAtpPOtsAWAmQCYBWl5tTtQv7unytGsteVDIOOH1Sb_2mEqijRjr-7x0NW-OXPF6BTNo98j3N4IUWKSMihBie38UOsjv-vzVyyTVMZS7GCaLk4hUZ-yqxaYJ7f-T5e_UeXcUWpsggeUAW1Hn2uvWxo93S-BOHOssTmh4M-_mQwIwkbQw"}

// SendMsg 发送消息
func (conn *Connection) SendMsg(msg *message.Message) {
	msgJ, err := json.Marshal(msg)
	if err != nil {
		return
	}

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
		return message.Message{}, err
	}
	if mt != websocket.TextMessage {
		log.Printf(" 消息类型只支持文本 ")
		return message.Message{}, errors.New(" 消息类型只支持文本, 本次消息类型：" + string(mt))
	}
	//log.Printf(" 读取数据：{%s} ", string(msg))

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
	log.Printf(" 关闭链接,链接数量： " + string(count))
	err := conn.conn.Close()
	if err != nil {
		log.Printf(" 关闭链接异常 : %s ", err)
		return err
	}
	return nil
}
