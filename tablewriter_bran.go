package main

import (
    "fmt"
    "bytes"
    "github.com/olekukonko/tablewriter"
)

func formatUser(user userModel) (string, error) {
    var b bytes.Buffer
    keys := []string{"ID", "Name"}

    table := tablewriter.NewWriter(&b)
    table.SetHeader(keys)

    row := []string{user.id, user.name}
    table.Append(row)
    table.Render()

    return b.String(), nil
}

type userModel struct {
    id string
    name string
}

func main() {
    user := userModel{}
    user.id = "123456789"
    user.name = "chen_chang_sheng"
    res, _ := formatUser(user)
    fmt.Printf(res)
}

//OUTPUT:
/*
+-----------+-------------------+
|    ID     |       NAME        |
+-----------+-------------------+
| 123456789 | chen_chang_sheng |
+-----------+-------------------+
*/
