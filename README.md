# go-qiwi-wallet-api
[Qiwi wallet api](https://developer.qiwi.com/en/qiwicom/index.html) implementation in Go

### Installation  
`$ go get -u github.com/Barrokgl/go-qiwi-wallet-api`

### Usage  
```go
package main

import (
        "github.com/Barrokgl/go-qiwi-wallet-api"
        "log"
)

const token = "MyMagicToken"

func main() {
        api := goqiwi.NewQiwiApi(token, nil)
        balance, err := api.GetBalance()
        if err != nil {
                log.Fatal(err)
        }
        log.Print(balance) // see Balance struct
}
```