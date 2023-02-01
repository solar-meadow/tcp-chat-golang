package service

import (
	"crypto/sha256"
	"errors"
	"io/ioutil"
	"net"
	"os"
)

func welcome(connection net.Conn) error {
	file, err := os.Open("logo.txt")
	defer file.Close()
	if err != nil {
		return errors.New("error i can't find logo file")
	}
	res := checkHush("logo.txt")
	if !res {
		return errors.New("error file was changed\n")
	}
	data := make([]byte, 332)
	_, err = file.Read(data)
	if err != nil {
		return errors.New("i can't read from logo file")
	}

	data[331] = '\n'
	connection.Write([]byte("Welcome to TCP-Chat!\n"))
	connection.Write([]byte(data))
	return nil
}

func checkHush(filename string) bool {
	hasher := sha256.New()
	s, err := ioutil.ReadFile(filename)
	if err != nil {
		return false
	}
	hasher.Write(s)
	l := hasher.Sum(nil)
	hashlogo := []byte{209, 85, 238, 199, 175, 168, 176, 190, 183, 14, 61, 200, 78, 156, 31, 52, 207, 227, 20, 3, 79, 119, 183, 87, 136, 185, 6, 205, 196, 106, 250, 159}

	return string(l) == string(hashlogo)
}
