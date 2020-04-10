package main

import (
	"fmt"
	"os"

	wechatclient "github.com/rdxsl/go-wechat-work/client"
	"gitlab.com/rdxsl/go-wechat-work/workerpool"
)

func main() {
	wechatclient.Init()
	err := wechatclient.GetAccessTocken()

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	wechatclient.Test1()
	workerpool.Test()
}
