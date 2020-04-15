package wechatclient

import (
	"testing"
)

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
			GetAccessTocken(tc.force)
		})
	}
}
