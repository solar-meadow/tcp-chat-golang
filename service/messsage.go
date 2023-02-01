package service

import (
	"fmt"
	"net"
	"time"
)

func newMessage(msg string, username string, conn net.Conn, time string) message {
	addr := conn.RemoteAddr().String()
	if msg == " has left." {
		return message{
			text:    "\n" + username + msg,
			address: addr,
		}
	}
	if msg != "" {
		return message{
			text:     "[" + time + "][" + username + "]: " + msg,
			username: username,
			address:  addr,
			time:     time,
		}
	}
	return message{
		username: username,
		address:  addr,
		time:     time,
	}
}

func Broadcaster() {
	for {
		select {
		case msg := <-messages:
			mu.Lock()
			if msg.text != msg.username+" has joined our chat" && msg.text != "" {
				history = append(history, msg.text)
			}
			for _, conn := range clients {
				msg.time = time.Now().Format("2006-01-02 15:04:05")
				if msg.text != "" {
					if msg.address == conn.RemoteAddr().String() {
						text := "[" + msg.time + "][" + msg.username + "]: "
						fmt.Fprint(conn, text)
						continue
					}
					fmt.Fprint(conn, "\n"+msg.text)
					time := time.Now().Format("2006-01-02 15:04:05")
					for key := range clients {
						if key.conn == conn {
							fmt.Fprint(conn, "\n"+"["+time+"]["+key.username+"]: ")
						}
					}
				} else {
					if msg.address == conn.RemoteAddr().String() {
						msg.time = time.Now().Format("2006-01-02 15:04:05")
						text := "[" + msg.time + "][" + msg.username + "]: "
						fmt.Fprint(conn, text)
						continue
					}
				}
			}
			mu.Unlock()

		case msg := <-leaving:
			mu.Lock()
			for u, conn := range clients {
				fmt.Fprintln(conn, msg.text)

				text := "[" + u.time + "][" + u.username + "]: "
				fmt.Fprint(conn, text)

			}
			// history = append(history, msg.text[1:])

			mu.Unlock()

		}
	}
}
