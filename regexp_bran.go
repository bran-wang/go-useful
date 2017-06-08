package main


import "fmt"
import "regexp"


func main() {

    var name = "bran"
    res,_ := regexp.Match("qemu()", []byte("qemu()001"))
    fmt.Println(res)
    if ok,_ := regexp.Match("qemum", []byte("qemu()001")); ok || name=="bran" {
        fmt.Println(ok)
        fmt.Println("ok, let's do it")
    }
}
