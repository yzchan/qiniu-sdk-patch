package dcdn

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/conf"
)

var (
	FusionHost = "http://fusion.qiniuapi.com"
)

type DcdnManager struct {
	mac   *auth.Credentials
	Debug bool
}

func NewDcdnManager(mac *qbox.Mac) *DcdnManager {
	return &DcdnManager{mac, false}
}

type TrafficReq struct {
	StartDate   string `json:"startDate"`
	EndDate     string `json:"endDate"`
	Granularity string `json:"granularity"`
	Domains     string `json:"domains"`
	Type        string `json:"type"`
}

// TrafficResp 为带宽/流量查询响应内容
type TrafficResp struct {
	Code  int                    `json:"code"`
	Error string                 `json:"error"`
	Time  []string               `json:"time,omitempty"`
	Data  map[string]TrafficData `json:"data,omitempty"`
}

// TrafficData 为带宽/流量数据
type TrafficData struct {
	DomainChina   []int `json:"china"`
	DomainOversea []int `json:"oversea"`
}

// GetStaticFluxData 纯静态CDN
//	StartDate	string		必须	开始日期，例如：2016-07-01
//	EndDate		string		必须	结束日期，例如：2016-07-03
//	Granularity	string		必须	粒度，取值：5min ／ hour ／day
//	Domains		[]string	必须	域名列表
func (m *DcdnManager) GetStaticFluxData(startDate, endDate, granularity string, domainList []string) (fluxData TrafficResp, err error) {
	return m.getData("/v2/dcdn/flux", "flux", startDate, endDate, granularity, domainList)
}

// GetDynFluxData 纯动态CDN
//	StartDate	string		必须	开始日期，例如：2016-07-01
//	EndDate		string		必须	结束日期，例如：2016-07-03
//	Granularity	string		必须	粒度，取值：5min ／ hour ／day
//	Domains		[]string	必须	域名列表
func (m *DcdnManager) GetDynFluxData(startDate, endDate, granularity string, domainList []string) (fluxData TrafficResp, err error) {
	return m.getData("/v2/dcdn/flux", "dynflux", startDate, endDate, granularity, domainList)
}

// GetDynReqCount 方法用来批量查询动态加速之动态请求数
// StartDate	string		必须	开始日期，例如：2016-07-01
// EndDate		string		必须	结束日期，例如：2016-07-03
// Granularity	string		必须	粒度，取值：5min/hour/day
// Domains		[]string	必须	域名列表
func (m *DcdnManager) GetDynReqCount(startDate, endDate, granularity string, domainList []string) (fluxData TrafficResp, err error) {
	return m.getData("/v2/dcdn/dynreqcount", "dynreqcount", startDate, endDate, granularity, domainList)
}

func (m *DcdnManager) getData(path, reqType, startDate, endDate, granularity string, domainList []string) (fluxData TrafficResp, err error) {
	domains := strings.Join(domainList, ";")
	reqBody := TrafficReq{
		StartDate:   startDate,
		EndDate:     endDate,
		Granularity: granularity,
		Domains:     domains,
		Type:        reqType,
	}

	resData, reqErr := postRequest(m.mac, path, reqBody)
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
	urlStr := fmt.Sprintf("%s%s", FusionHost, path)
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
