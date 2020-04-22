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

const httpTimeOut = 10

// AccessToken wechat work access token
type AccessToken struct {
	ErrorCode   int    `json:"errorcode"`
	ErrorMsg    string `json:"errmsg"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	mu          sync.RWMutex
	timeStamp   int64
}

var accessToken AccessToken

// GetAccessTocken get wechat work access token from api
func GetAccessTocken(force bool, wechatCorpID string, wechatCorpSecret string) (err error) {

	// never get the access token yet, call the wechat api
	if accessToken.timeStamp == 0 {
		return getAccessToken(wechatCorpID, wechatCorpSecret)
	}

	// token expired, refresh
	if accessToken.timeStamp+int64(accessToken.ExpiresIn) < time.Now().Unix() {
		return getAccessToken(wechatCorpID, wechatCorpSecret)
	}

	if force == true {
		return getAccessToken(wechatCorpID, wechatCorpSecret)
	}

	return nil
}

func getAccessToken(wechatCorpID string, wechatCorpSecret string) (err error) {
	var netClient = &http.Client{
		Timeout: time.Second * httpTimeOut,
	}
	if wechatCorpID == "" {
		wechatCorpID = os.Getenv("WECHAT_CORPID")
	}
	if wechatCorpSecret == "" {
		wechatCorpSecret = os.Getenv("WECHAT_CORPSECRET")
	}
	if wechatCorpID == "" || wechatCorpSecret == "" {
		err = fmt.Errorf("env variable WECHAT_CORPID or WECHAT_CORPSECRET is empty, and wechatCorpID wechatCorpSecret in getAccessToken are not set")
		return
	}

	wechatURL := fmt.Sprintf(wechatWorkAPI, wechatCorpID, wechatCorpSecret)
	resp, err := netClient.Get(wechatURL)
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
