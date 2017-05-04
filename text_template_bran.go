package main

import (
    "fmt"
    "bytes"
    "text/template"
)

func parsetemplate(templateFile string) string {
    tmpl, err := template.ParseFiles(templateFile)
    if err != nil {
        fmt.Printf("Error parsing template %s\n", templateFile, err)
        return ""
    }

    var buf bytes.Buffer
    err = tmpl.Execute(&buf, paraMap())
    if err != nil {
        fmt.Printf("Error executing template %s\n", templateFile, err)
        return ""
    }
    return buf.String()
}

func paraMap() map[string]string {
    params := make(map[string]string)
    params["hostname"] = "bran wang"
    return params
}


func main() {
    res := parsetemplate("/Users/branw/Documents/code/go_moudle/config.config")
    fmt.Println(res)
}

//input
/*
name: steve
hostname: {{.hostname}}

*/

//output
/*

name: steve
hostname: bran wang
*/


