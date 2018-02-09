package main

import (
  "log"
  "fmt"
  "net"
)

func handleConn(c net.Conn) {
  defer c.Close()
  buf := make([]byte, 1024)

  for {
    _, err := c.Read(buf)
    if err != nil {
      log.Println(err)
      return
    }
    log.Println(string(buf))
  }
}

func main() {
  l, err := net.Listen("tcp", ":8888")
  if err != nil {
    fmt.Println("Listen error:", err)
    return
  }

  for {
    c, err := l.Accept()
    if err != nil {
      fmt.Println("Accept error: ", err)
        break
    }

    go handleConn(c)
  }
}