# TCP Chat

TCP Chat is an application built using the Golang programming language. It is designed to allow users to send and receive messages on a single chat room over a TCP connection.

The application is built using goroutines to ensure that multiple users can connect and send messages concurrently. All messages are stored in a single chat room, allowing all users to read and send messages.

To use the application, users must first connect to the specified address and port with netcat. Once connected, users can send messages to the chat room and read messages sent by others.

# How to use

1. we need start server with command:
```
 go run . 8080
 ```
2. client need to connect to our server with netcat command:
```
 nc  localhost 8080
 ```
this command will give us connection to our server and then we can to send and receive data from server