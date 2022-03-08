package dora

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/conf"
	"github.com/yzchan/qiniu-sdk-patch/lib"
)

const Host = "stats-dora.qiniuapi.com"

type DoraManager struct {
	mac *auth.Credentials
}

func NewDoraManager(mac *qbox.Mac) *DoraManager {
	return &DoraManager{mac: mac}
}

type CountResp struct {
	Time   time.Time `json:"time"`
	Values struct {
		Value int64 `json:"value"`
	} `json:"values"`
}

func (m *DoraManager) GetCount(item string, beginDate string, endDate string) (data []CountResp, err error) {
	var (
		content []byte
		resp    []CountResp
	)

	query := fmt.Sprintf("start=%s&end=%s", lib.TransDate(beginDate), lib.TransDate(endDate))
	content, err = sendRequest(m.mac, item, query)

	if err = json.Unmarshal(content, &resp); err != nil {
		return
	}

	return resp, nil
}

func sendRequest(mac *qbox.Mac, item string, query string) (resp []byte, err error) {
	u := url.URL{
		Scheme:   "https",
		Host:     Host,
		Path:     "/v3/statistic/" + item,
		RawQuery: query,
	}

	client := &http.Client{}
	var request *http.Request

	if request, err = http.NewRequest("GET", u.String(), nil); err != nil {
		return
	}
	request.Header.Set("Host", Host)
	request.Header.Set("Content-Type", conf.CONTENT_TYPE_FORM)
	if _, err = mac.SignRequest(request); err != nil {
		return
	}

	if err = mac.AddToken(auth.TokenQiniu, request); err != nil {
		return
	}

	var response *http.Response
	if response, err = client.Do(request); err != nil {
		return
	}

	defer response.Body.Close()
	if resp, err = ioutil.ReadAll(response.Body); err != nil {
		return
	}

	return
}
