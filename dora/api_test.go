package dora

import (
	"os"
	"testing"
	"time"

	"github.com/qiniu/go-sdk/v7/auth"
)

//global variables

var (
	ak = os.Getenv("accessKey")
	sk = os.Getenv("secretKey")

	layout    = "2006-01-02"
	now       = time.Now()
	startDate = now.AddDate(0, 0, -2).Format(layout)
	endDate   = now.AddDate(0, 0, -1).Format(layout)
)

var mac *auth.Credentials
var doraManager *DoraManager

func init() {
	if ak == "" || sk == "" {
		panic("ak/sk should not be empty")
	}
	mac = auth.New(ak, sk)
	doraManager = NewDoraManager(mac)
}

// TestGetCount
func TestGetCount(t *testing.T) {
	type args struct {
		startDate string
		endDate   string
	}

	testCases := []struct {
		name      string
		args      args
		wantCount int
	}{
		{
			name: "doraManager_TestGetCount",
			args: args{
				startDate,
				endDate,
			},
			wantCount: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ret, err := doraManager.GetCount("imageslim-auto", tc.args.startDate, tc.args.endDate)
			t.Log(ret)
			if err != nil || len(ret) != tc.wantCount {
				t.Errorf("GetCount() error = %v", err)
				return
			}
		})
	}
}
