package main

import (
	"flag"
	"github.com/changan/websocket_gateway/connection"
	"github.com/changan/websocket_gateway/user"
	"html/template"
	"log"
	"net/http"
)

/*
	WebSocket 网关程序入口
*/

var (
	addr = flag.String("addr", "localhost:8080", "http service address")
)

func main() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/", NullPath)
	http.HandleFunc("/wg", WebsocketGateway)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

func WebsocketGateway(writer http.ResponseWriter, r *http.Request) {
	newConnection := connection.NewConnection(writer, r, user.User{})
	// 开启协程处理链接
	go newConnection.Start()
}

// NullPath 404
func NullPath(writer http.ResponseWriter, r *http.Request) {
	errTemplate.Execute(writer, "ws : //"+r.Host+"/wg")
}

var errTemplate = template.Must(template.New("").Parse(`
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
</head>
<body>
	404 ERROR PAGE!
</body>
</html>
`))
