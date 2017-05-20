package main

import (
    "crypto/md5"
    "encoding/hex"
    "fmt"
)
func main() {
    hash := md5.New()
    hash.Write([]byte("www.branfun.com")) // 需要加密的字符串为 www.branfun.com
    fmt.Printf("%s\n", hex.EncodeToString(hash.Sum(nil))) // 输出加密结果
}
