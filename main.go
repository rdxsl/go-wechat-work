package main

import (
	"fmt"
	"strconv"
	"time"

	wechatclient "github.com/rdxsl/go-wechat-work/client"
	pool "github.com/rdxsl/go-wechat-work/workerpool"
)

func Test(collector pool.Collector) {
	fmt.Println("sending")
	toUser := "jackxie"
	agentID := 1000002
	var text1 wechatclient.WechatMsg
	for i := 0; i < 3; i++ {
		stringI := "this is a from RDX-sl " + strconv.Itoa(i)
		text1 = wechatclient.WechatMsg{
			ToUser:  toUser,
			MsgType: "text",
			AgentID: agentID,
			TextBody: wechatclient.WechatMsgText{
				Content: stringI,
			},
			Safe: 0,
		}
		wechatclient.SendText(text1, "", "")
		// collector.Work <- text1
	}
}

const WORKER_COUNT = 5

func main() {

	// err := wechatclient.GetAccessTocken(false, "", "")

	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "error: %v\n", err)
	// 	os.Exit(1)
	// }

	collector := pool.StartDispatcher(WORKER_COUNT)
	Test(collector)
	time.Sleep(5 * time.Second)
}
