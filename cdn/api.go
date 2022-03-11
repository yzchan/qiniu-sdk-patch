package cdn

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/qiniu/go-sdk/v7/conf"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/cdn"
)

type CdnManager struct {
	mac *auth.Credentials
	*cdn.CdnManager
	Debug bool
}

func NewCdnManager(mac *qbox.Mac) *CdnManager {
	return &CdnManager{mac, cdn.NewCdnManager(mac), false}
}

type TrafficReq struct {
	StartDate   string `json:"startDate"`
	EndDate     string `json:"endDate"`
	Granularity string `json:"granularity"`
	Domains     string `json:"domains"`
	Type        string `json:"type"`
}

// GetDynFluxData 方法用来批量查询动态加速计费流量
// 接口文档 https://developer.qiniu.com/fusion/1230/traffic-bandwidth#4
//	StartDate	string		必须	开始日期，例如：2016-07-01
//	EndDate		string		必须	结束日期，例如：2016-07-03
//	Granularity	string		必须	粒度，取值：5min ／ hour ／day
//	Domains		[]string	必须	域名列表
func (m *CdnManager) GetDynFluxData(startDate, endDate, granularity string, domainList []string) (fluxData cdn.TrafficResp, err error) {
	domains := strings.Join(domainList, ";")
	reqBody := TrafficReq{
		StartDate:   startDate,
		EndDate:     endDate,
		Granularity: granularity,
		Domains:     domains,
		Type:        "dynflux",
	}

	resData, reqErr := postRequest(m.mac, "/v2/tune/flux", reqBody)
	if reqErr != nil {
		err = reqErr
		return
	}

	umErr := json.Unmarshal(resData, &fluxData)
	if umErr != nil {
		err = umErr
		return
	}

	return
}

// GetDynReqCount 方法用来批量查询动态加速之动态请求数
// 接口文档 https://developer.qiniu.com/fusion/1230/traffic-bandwidth#5
// StartDate	string		必须	开始日期，例如：2016-07-01
// EndDate		string		必须	结束日期，例如：2016-07-03
// Granularity	string		必须	粒度，取值：5min/hour/day
// Domains		[]string	必须	域名列表
func (m *CdnManager) GetDynReqCount(startDate, endDate, granularity string, domainList []string) (fluxData cdn.TrafficResp, err error) {
	domains := strings.Join(domainList, ";")
	reqBody := TrafficReq{
		StartDate:   startDate,
		EndDate:     endDate,
		Granularity: granularity,
		Domains:     domains,
		Type:        "dynreqcount",
	}

	resData, reqErr := postRequest(m.mac, "/v2/tune/dynreqcount", reqBody)
	if reqErr != nil {
		err = reqErr
		return
	}

	umErr := json.Unmarshal(resData, &fluxData)
	if umErr != nil {
		err = umErr
		return
	}

	return
}

// RequestWithBody 带body对api发出请求并且返回response body
func postRequest(mac *auth.Credentials, path string, body interface{}) (resData []byte, err error) {
	urlStr := fmt.Sprintf("%s%s", cdn.FusionHost, path)
	reqData, _ := json.Marshal(body)
	req, reqErr := http.NewRequest("POST", urlStr, bytes.NewReader(reqData))
	if reqErr != nil {
		err = reqErr
		return
	}

	accessToken, signErr := mac.SignRequest(req)
	if signErr != nil {
		err = signErr
		return
	}

	req.Header.Add("Authorization", "QBox "+accessToken)
	req.Header.Add("Content-Type", conf.CONTENT_TYPE_JSON)

	resp, respErr := http.DefaultClient.Do(req)
	if respErr != nil {
		err = respErr
		return
	}
	defer resp.Body.Close()

	resData, ioErr := ioutil.ReadAll(resp.Body)
	if ioErr != nil {
		err = ioErr
		return
	}

	return
}
