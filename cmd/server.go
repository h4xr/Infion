package main

import (
	"fmt"
	"infion/broker/server"
)

func main() {
	s, err := server.NewServer("127.0.0.1", 10001)
	if err != nil {
		fmt.Println(err)
	}
	s.Listen()
}