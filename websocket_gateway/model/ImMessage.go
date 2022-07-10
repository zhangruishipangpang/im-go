package model

/*
	IM 消息结构体
*/

type ImMessage struct {
	From    string      `json:"from,omitempty"`
	Dest    string      `json:"dest,omitempty"`
	ImType  uint32      `json:"im_type,omitempty"`
	Msg     interface{} `json:"msg,omitempty"`
	MsgType uint32      `json:"msg_type,omitempty"`
}

const (
	IM_TYPE_NOTIFY = 1
	IM_TYPE_1_1    = 2
	IM_TYPE_1_N    = 3

	MSG_TYPE_TEXT = 1
	MSG_TYPE_FILE = 2
)
