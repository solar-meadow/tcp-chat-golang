package service

import (
	"bufio"
	"net"
	"strings"
)

func nickname(connection net.Conn) (string, error) {
	for {

		connection.Write([]byte("[ENTER YOUR NAME]: "))
		duplicateStatus := false
		nick, err := text(connection)
		if err != nil {
			return "", err
		}
		mu.Lock()
		for u := range clients {
			if u.username == nick {
				duplicateStatus = true
			}
		}
		mu.Unlock()

		if nick != "" && !duplicateStatus && checkText(nick) {
			return nick, nil
		}
		if duplicateStatus {
			mu.Lock()
			connection.Write([]byte("Username is already used. Please use other nickname\n"))
			mu.Unlock()
		}
	}
}

func text(connection net.Conn) (string, error) {
	netData, err := bufio.NewReader(connection).ReadString('\n')
	if err != nil {
		return "", err
	}
	mes := strings.TrimSpace(netData)
	return mes, nil
}
