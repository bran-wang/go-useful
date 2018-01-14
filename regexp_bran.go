package main


import "fmt"
import "regexp"


func main() {

    var name = "bran"
    res,_ := regexp.Match("aemu()|01", []byte("qemu()001"))
    fmt.Println(res)
    if ok,_ := regexp.Match("qemum", []byte("qemu()001")); ok || name=="bran" {
        fmt.Println(ok)
        fmt.Println("ok, let's do it")
    }

    var datastore = "ran"

    var option = "branshare2T"
    res, _ = regexp.MatchString(datastore, option)
    fmt.Println(res)

    re_com,_ := regexp.Compile(datastore)
    loc := re_com.FindStringIndex(option)
    if loc != nil && loc[0] == 0 {
        fmt.Println("you are really match from the start")
    }
    fmt.Println(loc)
}
