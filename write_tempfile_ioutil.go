package main

import (
    "fmt"
    "io/ioutil"
)

func main() {
    s := "hi man, this is a test, how are you\n"
    content := []byte(s)

    file, err := ioutil.TempFile("/Users/branw/Documents/code/go_moudle", "ioutil.config")
    if err != nil {
        fmt.Println("Error writing to configuratin file")
    }
    defer file.Close()
    filename := file.Name()
    ioutil.WriteFile(filename, content, 0644)
    fmt.Println("filename is : ", filename)

}
