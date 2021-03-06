package dcdn

import (
	"os"
	"testing"
	"time"

	"github.com/qiniu/go-sdk/v7/auth"
)

//global variables

var (
	ak     = os.Getenv("accessKey")
	sk     = os.Getenv("secretKey")
	domain = os.Getenv("QINIU_TEST_DOMAIN")

	layout    = "2006-01-02"
	now       = time.Now()
	startDate = now.AddDate(0, 0, -2).Format(layout)
	endDate   = now.AddDate(0, 0, -1).Format(layout)
)

var mac *auth.Credentials
var dcdnManager *DcdnManager

func init() {
	if ak == "" || sk == "" {
		panic("ak/sk should not be empty")
	}
	mac = auth.New(ak, sk)
	dcdnManager = NewDcdnManager(mac)
}

// TestGetDynFluxData
func TestGetDynFluxData(t *testing.T) {
	type args struct {
		startDate   string
		endDate     string
		granularity string
		domainList  []string
	}

	testCases := []struct {
		name     string
		args     args
		wantCode int
	}{
		{
			name: "DcdnManager_TestGetDynFluxData",
			args: args{
				startDate,
				endDate,
				"day",
				[]string{domain},
			},
			wantCode: 200,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ret, err := dcdnManager.GetDynFluxData(tc.args.startDate, tc.args.endDate, tc.args.granularity, tc.args.domainList)
			t.Log(ret.Data)
			if err != nil || ret.Code != tc.wantCode {
				t.Errorf("GetDynFluxData() error = %v, %v", err, ret.Error)
				return
			}
		})
	}
}

// TestGetStaticFluxData
func TestGetStaticFluxData(t *testing.T) {
	type args struct {
		startDate   string
		endDate     string
		granularity string
		domainList  []string
	}

	testCases := []struct {
		name     string
		args     args
		wantCode int
	}{
		{
			name: "DcdnManager_TestGetFluxData",
			args: args{
				startDate,
				endDate,
				"day",
				[]string{domain},
			},
			wantCode: 200,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ret, err := dcdnManager.GetFluxData(tc.args.startDate, tc.args.endDate, tc.args.granularity, tc.args.domainList)
			t.Log(ret.Data)
			if err != nil || ret.Code != tc.wantCode {
				t.Errorf("GetFluxData() error = %v, %v", err, ret.Error)
				return
			}
		})
	}
}

// TestGetDynReqCount
func TestGetDynReqCount(t *testing.T) {
	type args struct {
		startDate   string
		endDate     string
		granularity string
		domainList  []string
	}

	testCases := []struct {
		name     string
		args     args
		wantCode int
	}{
		{
			name: "DcdnManager_TestGetDynReqCount",
			args: args{
				startDate,
				endDate,
				"day",
				[]string{domain},
			},
			wantCode: 200,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ret, err := dcdnManager.GetDynReqCount(tc.args.startDate, tc.args.endDate, tc.args.granularity, tc.args.domainList)
			t.Log(ret.Data)
			if err != nil || ret.Code != tc.wantCode {
				t.Errorf("GetDynReqCount() error = %v, %v", err, ret.Error)
				return
			}
		})
	}
}
