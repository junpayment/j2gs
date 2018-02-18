package main

import (
  "fmt"
  "flag"
  "reflect"
  "errors"
  "encoding/json"
  "github.com/c9s/inflect"
)

// 標準入力で受けたjsonをstructにする
func main() {
  flag.Parse()
  if  1 > len(flag.Args()) {
    panic(errors.New("USAGE: j2gs <json string>"))
  }

  jsonStr := flag.Arg(0)

  var data map[string]interface{}
  err := json.Unmarshal([]byte(jsonStr), &data)
  if nil != err {
    panic(err)
  }

  var out string
  out += fmt.Sprintf("type Hoge struct {\n")

  for k, v := range data {
    varName := k
    varNameCamel := inflect.Camelize(varName)

    var typeName string
    switch reflect.TypeOf(v).Name() {
    case "float64":
      typeName = "int"
    case "bool":
      typeName = "bool"
    default: //case "string":
      typeName = "string"
    }

    out += fmt.Sprintf("  %s %s `json:\"%s:%s\"`\n", varNameCamel, typeName, varName, typeName)
  }

  out += fmt.Sprintf("}")

  fmt.Println(out)
}
