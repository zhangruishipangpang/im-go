package interface_definition

/*
	用户认证接口
	- 1、检查token是否有效
*/

type IAuth interface {
	ck(token string) bool
}

type IAuthService interface {
	checkToken(token string) bool
}

type AuthService struct {
}

func (a AuthService) checkToken(token string) bool {
	//TODO implement me
	panic("implement me")
}
