package dora

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/qiniu/x/log"
	"github.com/yzchan/qiniu-sdk-patch/lib"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/conf"
)

const Host = "stats-dora.qiniuapi.com"

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Llevel)
}

type DoraManager struct {
	mac   *auth.Credentials
	Debug bool
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

	query := fmt.Sprintf("start=%s&end=%s", lib.FromDate(beginDate), lib.ToDate(endDate))
	if content, err = m.sendRequest(item, query); err != nil {
		return
	}

	if err = json.Unmarshal(content, &resp); err != nil {
		err = errors.New(err.Error() + ". response body:" + string(content))
		return
	}

	return resp, nil
}

func (m *DoraManager) sendRequest(item string, query string) (resp []byte, err error) {
	u := url.URL{
		Scheme:   "https",
		Host:     Host,
		Path:     "/v3/statistic/" + item,
		RawQuery: query,
	}

	client := &http.Client{}
	var request *http.Request

	if m.Debug {
		log.Infof(" request: GET %s", u.String())
	}
	if request, err = http.NewRequest("GET", u.String(), nil); err != nil {
		return
	}
	request.Header.Set("Host", Host)
	request.Header.Set("Content-Type", conf.CONTENT_TYPE_FORM)
	if _, err = m.mac.SignRequest(request); err != nil {
		return
	}

	if err = m.mac.AddToken(auth.TokenQiniu, request); err != nil {
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

	if m.Debug {
		log.Infof(" response: %s", string(resp))
	}
	return
}
