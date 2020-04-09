package wechatclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const wechatWorkAPI = "https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s"

type AccessToken struct {
	ErrorCode   int    `json:"errorcode"`
	ErrorMsg    string `json:"errmsg"`
	EccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

func GetAccessTocken() (accessToken AccessToken, err error) {
	wechatCorpID := os.Getenv("WECHAT_CORPID")
	wechatCorpSecret := os.Getenv("WECHAT_CORPSECRET")
	if wechatCorpID == "" || wechatCorpSecret == "" {
		err = fmt.Errorf("env variable WECHAT_CORPID or WECHAT_CORPSECRET is empty")
		return
	}

	wechatURL := fmt.Sprintf(wechatWorkAPI, wechatCorpID, wechatCorpSecret)
	resp, err := http.Get(wechatURL)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &accessToken)
	if err != nil {
		return
	}
	return
}
