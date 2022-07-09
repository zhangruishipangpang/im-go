package auth

import (
	"fmt"
	"github.com/changan/websocket_gateway/model"
	"github.com/changan/websocket_gateway/register"
	"testing"
)

var token = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOlsiYWxsIl0sInVzZXJfZW1haWwiOiIxODg0NjQzOTk1MkAxNjMuY29tIiwidXNlcl9tb2JpbGUiOiIxODg0NjQzOTk1MiIsInVzZXJfbmFtZSI6ImNoYW5nYW4iLCJzY29wZSI6WyJhbGwiXSwiZXhwIjoxNjU3MzgxNDE1LCJhdXRob3JpdGllcyI6WyJlbXBsb3llZSJdLCJqdGkiOiIzMjNkZmE1Yy1kYjJkLTQ5YjctODc1OS1lNzQzMjk3OWFjYTYiLCJjbGllbnRfaWQiOiJnYXRld2F5IiwidXNlcl9udW1iZXIiOiJjMzVkNDcxZjNlYjhjYjYyMTYwNTRmYzVhMmIwOGQ2NyJ9.C0Z6z_eUIoY0u3hXW3maz3ZAlCtAE20A6l07jAuITdCxyUfErH7iLOipOWCWir73FBciMx4TZrP4zMgfXLVEQw6s93FrheikZ_k_-tv8ixiKbDZcscdEy2LqDKYNT_zr9NGcSWG-q5lpGIo3mFZ0cb-QOOOsIBr281_rVtyAtpPOtsAWAmQCYBWl5tTtQv7unytGsteVDIOOH1Sb_2mEqijRjr-7x0NW-OXPF6BTNo98j3N4IUWKSMihBie38UOsjv-vzVyyTVMZS7GCaLk4hUZ-yqxaYJ7f-T5e_UeXcUWpsggeUAW1Hn2uvWxo93S-BOHOssTmh4M-_mQwIwkbQw"

func TestAuthService_CheckToken(t *testing.T) {

	nacos := register.GetNacosClient()
	nacos.RegisterInstance()
	authService := AuthService{}
	_, _ = authService.CheckToken(token)
}

func TestAuthService_CheckToken2(t *testing.T) {
	msg := model.UserMsg{}
	service := AuthService{}
	fmt.Println(msg)

	fmt.Println(service.UserMsg == msg)
	fmt.Println(service.UserMsg == model.UserMsg{})
	fmt.Println(&service.UserMsg == nil)
}
