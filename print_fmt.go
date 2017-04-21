package main


import "fmt"


func main() {
    bs := "hello string"
    fmt.Printf("this is a test, %s\n", bs)

    bbs := "end happy"
    endofhappy := fmt.Sprintf("this is %s, and there is %s", bs, bbs)
    fmt.Println(endofhappy)
}
