package main

import (
    "bytes"
    "encoding/json"
    "fmt"
)

type ColorGroup struct{
    ID int
    Name string
    Colors []string
}

func main(){
    group := ColorGroup{
        ID: 1,
        Name: "Reds",
        Colors: []string{"Crimson", "Red", "Maroon"},
    }

    b, err := json.Marshal(group)
    if err != nil {
        fmt.Println("error:", err)
    }
    fmt.Println(string(b[:]))

    //--------Unmarshal
    var jsonBlog = []byte(`[
        {"ID":1, "Name":"Reds1", "Colors":["Crimson", "Red1", "Ruby1", "Maroon1"]},
        {"ID":2, "Name":"Reds2", "Colors":["Crimson", "Red2", "Ruby2", "Maroon2"]},
        {"ID":3, "Name":"Reds3", "Colors":["Crimson", "Red3", "Ruby3", "Maroon3"]}
    ]`)

    var animals []ColorGroup
    error := json.Unmarshal(jsonBlog, &animals)
    if error != nil{
        fmt.Println("error:", error)
    }

    for i,x := range animals{
        fmt.Println(i, x.Name)
    }

    //---------Indent
    dst := new(bytes.Buffer)
    json.Indent(dst, jsonBlog, "", "    ")
    fmt.Println(dst)
}

/*

{"ID":1,"Name":"Reds","Colors":["Crimson","Red","Maroon"]}
0 Reds1
1 Reds2
2 Reds3
[
    {
        "ID": 1,
        "Name": "Reds1",
        "Colors": [
            "Crimson",
            "Red1",
            "Ruby1",
            "Maroon1"
        ]
    },
    {
        "ID": 2,
        "Name": "Reds2",
        "Colors": [
            "Crimson",
            "Red2",
            "Ruby2",
            "Maroon2"
        ]
    },
    {
        "ID": 3,
        "Name": "Reds3",
        "Colors": [
            "Crimson",
            "Red3",
            "Ruby3",
            "Maroon3"
        ]
    }
]
*/
