package main

import (
	"fmt"
	"net"
	"runtime"

	"github.com/with-insomnia/tcp-chat-golang/service"
)

func main() {
	port, err := service.Port()
	if err != nil {
		service.Wrap("error when get port\n", err)
		return
	}
	listen, err := net.Listen("tcp", "localhost:"+port)
	if err != nil {
		service.Wrap(("error when listening port: " + port + "\n"), err)
		return
	}
	fmt.Println("Listening on the port :" + port)
	go service.Broadcaster()
	for {
		conn, err := listen.Accept()
		if err != nil {
			service.Wrap(conn.RemoteAddr().String()+": ", err)
			conn.Close()
		}
		if runtime.NumGoroutine()-1 > 10 {
			fmt.Fprint(conn, "user limit exceeded")
			conn.Close()
		}
		go service.Handle(conn)
	}
}
