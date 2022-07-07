package http

import (
	"strings"
	"testing"
)

func TestInvokeRequest(t *testing.T) {

	header := make(map[string]string, 3)
	header["token"] = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c111112VyX25hbWUiOiJ6aGFuZ3NhbiIsInNjb3BlIjpbImFsbCJdLCJTVXNlciI6eyJwYXNzd29yZCI6bnVsbCwidXNlcm5hbWUiOiJ6aGFuZ3NhbiIsImF1dGhvcml0aWVzIjpbeyJhdXRob3JpdHkiOiJhZG1pbiJ9XSwiYWNjb3VudE5vbkV4cGlyZWQiOnRydWUsImFjY291bnROb25Mb2NrZWQiOnRydWUsImNyZWRlbnRpYWxzTm9uRXhwaXJlZCI6dHJ1ZSwiZW5hYmxlZCI6dHJ1ZX0sImV4cCI6MTY0ODY5MTY5NSwiYXV0aG9yaXRpZXMiOlsiYWRtaW4iXSwianRpIjoiYTExNDE5ODgtODcyYy00NTE2LTk5YWMtMjQ3MDVlZjRmNWI0IiwiY2xpZW50X2lkIjoiYmFpZHUifQ.G32br6PTm_EEW6R02l_Ws8SRqt_o8fl59ATwGqPdd5tSSJeWKMVdebuQWvMHy5HzzUO7cl6rwtYoflh9UgwhQnNY_ivYGM_SZVGLss1625WS87fySlEn17hNnkU461bxdJKSBIzKrDE-zfJ-Ody_RggaTcNPGpZVeE2ectHy9ZUUPQB_aUznaWZxolkN2eZ1wgb2Ruq7APsqqOr70IFs55c8mCsmhlbd8yBenEPLwJP2igyGcPRsmKQqdEnsqk1NjIBXh_UPVwlXbPVOa0Y-Q5GzTwPAYtoMIDue-G3U1kBNVwDNUiPVGRN7ZlG92Ivk5CFe5td_WjHGXLj0WWLzKw"
	param := "token=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c111112VyX25hbWUiOiJ6aGFuZ3NhbiIsInNjb3BlIjpbImFsbCJdLCJTVXNlciI6eyJwYXNzd29yZCI6bnVsbCwidXNlcm5hbWUiOiJ6aGFuZ3NhbiIsImF1dGhvcml0aWVzIjpbeyJhdXRob3JpdHkiOiJhZG1pbiJ9XSwiYWNjb3VudE5vbkV4cGlyZWQiOnRydWUsImFjY291bnROb25Mb2NrZWQiOnRydWUsImNyZWRlbnRpYWxzTm9uRXhwaXJlZCI6dHJ1ZSwiZW5hYmxlZCI6dHJ1ZX0sImV4cCI6MTY0ODY5MTY5NSwiYXV0aG9yaXRpZXMiOlsiYWRtaW4iXSwianRpIjoiYTExNDE5ODgtODcyYy00NTE2LTk5YWMtMjQ3MDVlZjRmNWI0IiwiY2xpZW50X2lkIjoiYmFpZHUifQ.G32br6PTm_EEW6R02l_Ws8SRqt_o8fl59ATwGqPdd5tSSJeWKMVdebuQWvMHy5HzzUO7cl6rwtYoflh9UgwhQnNY_ivYGM_SZVGLss1625WS87fySlEn17hNnkU461bxdJKSBIzKrDE-zfJ-Ody_RggaTcNPGpZVeE2ectHy9ZUUPQB_aUznaWZxolkN2eZ1wgb2Ruq7APsqqOr70IFs55c8mCsmhlbd8yBenEPLwJP2igyGcPRsmKQqdEnsqk1NjIBXh_UPVwlXbPVOa0Y-Q5GzTwPAYtoMIDue-G3U1kBNVwDNUiPVGRN7ZlG92Ivk5CFe5td_WjHGXLj0WWLzKw"
	InvokeRequest(
		"http://127.0.0.1:8081/auth/check_token?"+param,
		"POST",
		strings.NewReader(param),
		header,
	)

}
