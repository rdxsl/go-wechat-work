package wechatclient

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//Test requires you to set the correct WECHAT_CORPID & WECHAT_CORPSECRET env variable
func TestGetAccessTocken(t *testing.T) {
	testcases := map[string]struct {
		force bool
	}{
		"force is true": {
			force: true,
		},
	}
	for testName, tc := range testcases {
		t.Run(testName, func(t *testing.T) {
			a := accessToken.AccessToken
			GetAccessTocken(tc.force, "", "")
			assert.NotEqual(t, a, accessToken.AccessToken)
		})
	}
}
