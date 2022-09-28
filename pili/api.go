package pili

import (
	"encoding/json"
	"fmt"
	"github.com/qiniu/go-sdk/v7/storage"
	"github.com/yzchan/qiniu-sdk-patch/lib"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/conf"
)

var (
	FusionHost = "pili.qiniuapi.com"
)

type PiliManager struct {
	mac   *auth.Credentials
	Debug bool
}

func NewPiliManager(mac *qbox.Mac) *PiliManager {
	return &PiliManager{mac, false}
}

type FluxResp struct {
	Time   time.Time `json:"time"`
	Values struct {
		Flow int64 `json:"flow"`
	} `json:"values"`
}

// TrafficData 为带宽/流量数据
type TrafficData struct {
	DomainChina   []int `json:"china"`
	DomainOversea []int `json:"oversea"`
}

// GetUpFluxData
/**
 *  参数
 *	StartDate	string		必须	开始日期，例如：20160701
 *	EndDate		string		必须	结束日期，例如：20160703
 *	Granularity	string		必须	时间粒度，可取值为 5min、hour、day、month
 *	Area		[]string	必须	域名列表
 */
func (m *PiliManager) GetUpFluxData(startDate, endDate, granularity string, area []string) (fluxData []FluxResp, err error) {
	return m.getData("/statd/upflow", startDate, endDate, granularity, area)
}

// GetDownFluxData
/**
 *  参数
 *	StartDate	string		必须	开始日期，例如：20160701
 *	EndDate		string		必须	结束日期，例如：20160703
 *	Granularity	string		必须	时间粒度，可取值为 5min、hour、day、month
 *	Area		[]string	必须	域名列表
 */
func (m *PiliManager) GetDownFluxData(startDate, endDate, granularity string, area []string) (fluxData []FluxResp, err error) {
	return m.getData("/statd/downflow", startDate, endDate, granularity, area)
}

func (m *PiliManager) getData(path, startDate, endDate, granularity string, area []string) (fluxData []FluxResp, err error) {
	var uri url.URL
	query := uri.Query()
	query.Add("begin", lib.FromDate(startDate))
	query.Add("end", lib.ToDate(endDate))
	query.Add("select", "flow")
	query.Add("g", granularity) // 时间聚合粒度(5min hour day month)
	if area != nil {
		for _, a := range area {
			query.Add("$area", a)
		}
	}
	fmt.Println(query.Encode())
	resData, reqErr := sendGetRequest(m.mac, path, query.Encode())
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

func sendGetRequest(mac *qbox.Mac, path string, query string) (resp []byte, err error) {
	u := url.URL{
		Scheme:   "http",
		Host:     FusionHost,
		Path:     path,
		RawQuery: query,
	}
	//fmt.Println(u.String())
	client := &http.Client{}
	var request *http.Request
	if request, err = http.NewRequest("GET", u.String(), nil); err != nil {
		return
	}
	request.Header.Set("Host", storage.DefaultAPIHost)
	request.Header.Set("Content-Type", conf.CONTENT_TYPE_FORM)
	if _, err = mac.SignRequest(request); err != nil {
		return
	}

	if err = mac.AddToken(auth.TokenQBox, request); err != nil {
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
	fmt.Println(string(resp))
	return
}
