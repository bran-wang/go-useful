package main

import (
	"fmt"
	"net"
	"log"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("usage: go run client.go YOUR_MESSAGE")
		return
	}

	log.Println("begin dial....")
	conn, err := net.Dial("tcp", "10.161.145.37:8888")
	if err != nil {
		log.Println("dial error:", err)
		return
	}

	defer conn.Close()
	log.Println("dial ok")

	data := os.Args[1]
	conn.Write([]byte(data))
}