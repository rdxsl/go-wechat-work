package wechatclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
	"time"
)

const wechatWorkAPI = "https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s"

type AccessToken struct {
	ErrorCode   int    `json:"errorcode"`
	ErrorMsg    string `json:"errmsg"`
	EccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	mu          sync.RWMutex
	timeStamp   int64
}

var accessToken AccessToken

func GetAccessTocken(force bool) (err error) {

	// never get the access token yet, call the wechat api
	if accessToken.timeStamp == 0 {
		return getAccessToken()
	}

	// token expired, refresh
	if accessToken.timeStamp+int64(accessToken.ExpiresIn) < time.Now().Unix() {
		return getAccessToken()
	}

	if force == true {
		return getAccessToken()
	}

	return nil
}

func getAccessToken() (err error) {
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

	accessToken.mu.Lock()
	err = json.Unmarshal(body, &accessToken)
	if err != nil {
		return
	}

	accessToken.timeStamp = time.Now().Unix()
	accessToken.mu.Unlock()
	return
}
