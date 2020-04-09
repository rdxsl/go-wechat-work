package wechatclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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

func SendText(message string, toUser string, agentID int, accessToken AccessToken) (err error) {
	url := fmt.Sprintf(wechatSendURL, accessToken.EccessToken)

	wechatMsg := &WechatMsg{
		ToUser:  toUser,
		MsgType: "text",
		AgentID: agentID,
		TextBody: WechatMsgText{
			Content: message,
		},
		Safe: 0,
	}

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
