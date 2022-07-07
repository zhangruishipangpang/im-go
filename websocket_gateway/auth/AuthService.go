package auth

import "github.com/changan/websocket_gateway/http"

type AuthService struct {
	//Token string
}

type TokenBody struct {
	Token string `json:"token"`
}

func (a *AuthService) CheckToken(token string) (bool, string) {
	body := TokenBody{Token: token}
	responseBody := http.InvokeRequestFromServiceName("auth-serve", "/auth/check_token", body, nil)
	return responseBody != (http.ResponseBody{}) && responseBody.Code == 200, responseBody.Msg
}
