package wechatclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

const wechatSendURL = "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=%s"

type WechatMsg struct {
	ToUser   string        `json:"touser"`
	MsgType  string        `json:"msgtype"`
	AgentID  int           `json:"agentid"`
	TextBody WechatMsgText `json:"text"`
	Safe     int           `json:"safe"`
}

type WechatMsgText struct {
	Content string `json:"content"`
}

var TextChannel chan WechatMsg

func Init() {
	TextChannel = make(chan WechatMsg, 1)
}

func Test1() {

	fmt.Println("sending")
	toUser := "jackxie"
	agentID := 1000002
	var text1 WechatMsg
	for i := 0; i < 2; i++ {
		stringI := "paipai xie " + strconv.Itoa(i)
		text1 = WechatMsg{
			ToUser:  toUser,
			MsgType: "text",
			AgentID: agentID,
			TextBody: WechatMsgText{
				Content: stringI,
			},
			Safe: 0,
		}
		fmt.Println(text1)
		TextChannel <- text1
		SendText()
	}
}

func SendText() (err error) {
	url := fmt.Sprintf(wechatSendURL, accessToken.EccessToken)

	wechatMsg := <-TextChannel

	reqBody, err := json.Marshal(wechatMsg)
	if err != nil {
		return
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	fmt.Println(string(body))
	return
}
