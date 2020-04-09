package main

import (
	"fmt"
	"os"

	wechatclient "github.com/rdxsl/go-wechat-work/client"
)

func main() {
	a, err := wechatclient.GetAccessTocken()

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(a.ExpiresIn)
	err = wechatclient.SendText("你好", "jackxie", 1000002, a)
	if err != nil {
		fmt.Println("we have error sending msg")
	}
}
