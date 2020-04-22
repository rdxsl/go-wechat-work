package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	wechatclient "github.com/rdxsl/go-wechat-work/client"
)

func Test() {
	fmt.Println("sending")
	toUser := "2"
	agentID := 1000002
	var text1 wechatclient.WechatMsg
	for i := 0; i < 3; i++ {
		stringI := "this is a from RDX-sl " + strconv.Itoa(i)
		text1 = wechatclient.WechatMsg{
			ToTag:   toUser,
			MsgType: "markdown",
			AgentID: agentID,
			MarkDownBody: wechatclient.WechatMsgMarkDown{
				Content: stringI,
			},
			Safe: 0,
		}
		enc := json.NewEncoder(os.Stdout)
		enc.Encode(text1)
		wechatclient.SendText(text1, "", "")
	}
}

const WORKER_COUNT = 5

func main() {
	Test()
	time.Sleep(5 * time.Second)
}
