package service

import (
	"bufio"
	"fmt"
	"net"
	"sync"
	"time"
)

var (
	clients  = make(map[message]net.Conn)
	leaving  = make(chan message)
	messages = make(chan message)
	history  []string
	mu       sync.Mutex
)

type message struct {
	username string
	text     string
	time     string
	address  string
	conn     net.Conn
}

func Handle(conn net.Conn) {
	var statusChatHistory bool
	err := welcome(conn)
	if err != nil {
		WrapConn(conn, err)
		conn.Close()
		return
	}
	username, err := nickname(conn)
	if err != nil {
		return
	}

	if !statusChatHistory {
		mu.Lock()
		for _, v := range history {
			fmt.Fprintln(conn, string(v))
		}
		mu.Unlock()
		statusChatHistory = true
	}
	t := time.Now().Format("2006-01-02 15:04:05")
	user := message{
		username: username,
		address:  conn.RemoteAddr().String(),
		conn:     conn,
		text:     username + " has joined our chat",
		time:     t,
	}
	mu.Lock()
	clients[user] = conn
	messages <- user
	mu.Unlock()
	input := bufio.NewScanner(conn)
	for input.Scan() {
		str := input.Text()
		if !checkText(str) {
			str = ""
		}
		time := time.Now().Format("2006-01-02 15:04:05")

		mu.Lock()

		messages <- newMessage(str, username, conn, time)
		mu.Unlock()
	}
	mu.Lock()
	delete(clients, user)
	time := time.Now().Format("2006-01-02 15:04:05")
	leaving <- newMessage(" has left.", username, conn, time)
	mu.Unlock()
	conn.Close()
}

func checkText(s string) bool {
	for _, v := range s {
		if v < 32 {
			return false
		}
	}
	return true
}
