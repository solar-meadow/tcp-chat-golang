package service

import (
	"fmt"
	"net"
)

func Wrap(msg string, err error) {
	fmt.Println(msg + err.Error())
}

func WrapConn(conn net.Conn, err error) {
	fmt.Fprint(conn, err.Error())
}
