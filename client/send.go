package wechatclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const wechatSendURL = "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=%s"

const wechatSendRetry = 3

type WechatMsg struct {
	ToTag    string        `json:"totag"`
	ToParty  string        `json:"toparty"`
	ToUser   string        `json:"touser"`
	MsgType  string        `json:"msgtype"`
	AgentID  int           `json:"agentid"`
	TextBody WechatMsgText `json:"text"`
	Safe     int           `json:"safe"`
}

type WechatMsgText struct {
	Content string `json:"content"`
}

type WechatMsgSendReturn struct {
	ErrCode int64  `json:"errcdoe"`
	ErrMsg  string `json.:"errmsg"`
}

func SendText(wechatMsg WechatMsg, wechatCorpID string, wechatCorpSecret string) (err error) {

	err = GetAccessTocken(false, wechatCorpID, wechatCorpSecret)
	if err != nil {
		return
	}
	err = send(wechatMsg)
	return
}

func send(wechatMsg WechatMsg) (err error) {
	accessToken.mu.RLock()
	defer accessToken.mu.RUnlock()
	url := fmt.Sprintf(wechatSendURL, accessToken.EccessToken)

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
	var wmr WechatMsgSendReturn
	json.Unmarshal(body, wmr)
	if wmr.ErrCode != 0 {
		return fmt.Errorf("wechat send return with error code %d, error msg %s", wmr.ErrCode, wmr.ErrMsg)
	}
	return
}
