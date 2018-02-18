# j2gs
Convert json strings to Go struct.

```bash
go get -u github.com/junpayment/j2gs
j2gs '{"hoge_hoge": 1, "fuga": "chome"}'
```

You will get go struct below.

```go
type Hoge struct {
  Fuga string `json:"fuga:string"`
  HogeHoge int `json:"hoge_hoge:int"`
}
```