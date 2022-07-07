package message

/*
	消息体
*/

type Message struct {
	From     uint32 `json:"from"`      // 消息来源
	Dest     uint32 `json:"dest"`      // 消息目标
	DataType uint32 `json:"data_type"` // 消息类型
	Data     string `json:"data"`      // 消息体
}

const (
	TEXT    = 1 // 文本消息
	PICTURE = 2 // 图片消息
	TOKEN   = 3 // token
)

func NewMessage(from uint32, dest uint32, data string) *Message {
	msg := &Message{
		From: from,
		Dest: dest,
		Data: data,
	}
	return msg
}

func NewMessageContainsType(from uint32, dest uint32, DataType uint32, data string) *Message {
	msg := &Message{
		From:     from,
		Dest:     dest,
		Data:     data,
		DataType: DataType,
	}
	return msg
}
