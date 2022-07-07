package iface

/*
	用户认证接口
	- 1、检查token是否有效
*/

type IAuthService interface {
	CheckToken(token string) (bool, string)
}
