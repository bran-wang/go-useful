package main

import (
    "os/exec"
    "fmt"
)


func main() {
    cmd := exec.Command("ls", "-l")
    out, err := cmd.Output()
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println(string(out))
}

