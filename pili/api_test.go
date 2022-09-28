package pili

import (
	"os"
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
var piliManager *PiliManager

func init() {
	if ak == "" || sk == "" {
		panic("ak/sk should not be empty")
	}
	mac = auth.New(ak, sk)
	piliManager = NewPiliManager(mac)
}
