# j2gs
Convert a json file to Go struct.

```bash
go get -u github.com/junpayment/j2gs
j2gs sample.json
```

sample.json

```json
{"hoge_hoge": 1, "fuga": "chome"}
```

You will get go struct below.

```go
type Hoge struct {
  Fuga string `json:"fuga:string"`
  HogeHoge int `json:"hoge_hoge:int"`
}
```