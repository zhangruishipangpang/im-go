package auth

import (
	"github.com/changan/websocket_gateway/http"
	"github.com/changan/websocket_gateway/model"
)

type AuthService struct {
	//Token string
	UserMsg model.UserMsg
}

func (a *AuthService) CheckToken(token string) (bool, string) {
	body := model.TokenBody{Token: token}
	responseBody := http.InvokeRequestFromServiceName("auth-serve", "/auth/check_token", body, nil)
	isValid := responseBody != (model.ResponseBody{}) && responseBody.Code == 200
	if isValid && a.UserMsg == (model.UserMsg{}) {
		user := responseBody.E
		var u = user.(map[string]interface{})
		a.UserMsg = model.ConvertToUserMsg(u)
	}
	return isValid, responseBody.Msg
}

func (a AuthService) AuthUserMsg(token string) model.UserMsg {
	return a.UserMsg
}
