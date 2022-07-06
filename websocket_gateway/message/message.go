package message

/*
	消息体
*/

type Message struct {
	From uint32 `json:"from"` // 消息来源
	Dest uint32 `json:"dest"` // 消息目标
	Data string `json:"data"` // 消息体
}

func NewMessage(from uint32, dest uint32, data string) *Message {
	msg := &Message{
		From: from,
		Dest: dest,
		Data: data,
	}
	return msg
}
