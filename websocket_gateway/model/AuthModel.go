package model

// UserMsg 从auth服务获取的用户信息
type UserMsg struct {
	UserEmail  string `json:"user_email"`
	UserMobile string `json:"user_mobile"`
	UserName   string `json:"user_name"`
	ClientId   string `json:"client_id"`
	UserNumber string `json:"user_number"`
}

type TokenBody struct {
	Token string `json:"token"`
}

func ConvertToUserMsg(u map[string]interface{}) UserMsg {
	userMsg := UserMsg{
		UserEmail:  u["user_email"].(string),
		UserMobile: u["user_mobile"].(string),
		UserName:   u["user_name"].(string),
		ClientId:   u["client_id"].(string),
		UserNumber: u["user_number"].(string),
	}
	return userMsg
}
