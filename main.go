package main

import (
  "fmt"
  "flag"
  "reflect"
  "errors"
  "encoding/json"
  "io/ioutil"
  "github.com/c9s/inflect"
)

func main() {
  flag.Parse()
  if  1 > len(flag.Args()) {
    panic(errors.New("USAGE: j2gs <json file>"))
  }

  jsonBuff, err := ioutil.ReadFile(flag.Arg(0))
  if nil != err {
    panic(err)
  }

  structName := flag.Arg(1)
  if "" == structName {
    structName = "Hoge"
  }

  var data map[string]interface{}
  err = json.Unmarshal(jsonBuff, &data)
  if nil != err {
    panic(err)
  }

  var out string
  out += fmt.Sprintf("type %s struct {\n", structName)

  for k, v := range data {
    varName := k
    varNameCamel := inflect.Camelize(varName)

    var typeName string

    tempType := reflect.TypeOf(v)
    if nil == tempType {
      typeName = "UNKNOWN" // support null type
    } else {
      switch tempType.Name() {
      case "float64":
        typeName = "int"
      case "bool":
        typeName = "bool"
      default: //case "string":
        typeName = "string"
      }
    }

    out += fmt.Sprintf("  %s %s `json:\"%s:%s\"`\n", varNameCamel, typeName, varName, typeName)
  }

  out += fmt.Sprintf("}")

  fmt.Println(out)
}

