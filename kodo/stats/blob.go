package stats

import (
	"encoding/json"
	"net/url"
	"time"

	"github.com/yzchan/qiniu-sdk-patch/lib"
)

type BlobResp struct {
	Time   time.Time `json:"time"`
	Values struct {
		Flow int64 `json:"flow"`
		Hits int64 `json:"hits"`
		Size int64 `json:"size"`
	} `json:"values"`
}

// BlobIO https://developer.qiniu.com/kodo/3820/blob-io
func (m *StatsManager) BlobIO(beginDate, endDate, granularity, _select, bucket, domain, region, ftype string, src []string) (ret []BlobResp, err error) {
	var uri url.URL
	query := uri.Query()
	query.Add("begin", lib.FromDate(beginDate))
	query.Add("end", lib.ToDate(endDate))
	query.Add("g", granularity)  // 时间聚合粒度(5min hour day month)
	query.Add("select", _select) // 值字段    flow 流量 (Byte)       hits GET 请求次数
	// 下面是非必填
	if bucket != "" {
		query.Add("$bucket", bucket)
	}
	if region != "" {
		query.Add("$region", region)
	}
	if ftype != "" {
		query.Add("$ftype", ftype) // 存储类型 0-标准存储  1-低频存储  2-归档存储
	}
	if domain != "" {
		query.Add("$domain", domain)
	}
	if src != nil && len(src) > 0 {
		for _, item := range src {
			query.Add("$src", item)
		}
	}

	var resp []byte
	if resp, err = sendGetRequest(m.mac, "/v6/blob_io", query.Encode()); err != nil {
		return
	}
	if err = json.Unmarshal(resp, &ret); err != nil {
		return
	}
	return
}

// BlobTransfer https://developer.qiniu.com/kodo/3911/blob-transfer
func (m *StatsManager) BlobTransfer(beginDate, endDate, isOversea, taskId string) (ret []BlobResp, err error) {
	var uri url.URL
	query := uri.Query()
	query.Add("begin", lib.FromDate(beginDate))
	query.Add("end", lib.ToDate(endDate))
	query.Add("select", "size")
	query.Add("g", "day") // 时间聚合粒度(day month)
	// 以下为非必选项
	if isOversea != "" {
		query.Add("$is_oversea", isOversea)
	}
	if taskId != "" {
		query.Add("$taskid", taskId)
	}

	var resp []byte
	if resp, err = sendGetRequest(m.mac, "/v6/blob_transfer", query.Encode()); err != nil {
		return
	}
	if err = json.Unmarshal(resp, &ret); err != nil {
		return
	}
	return
}

// RsChType https://developer.qiniu.com/kodo/3913/rs-chtype
func (m *StatsManager) RsChType(beginDate, endDate, granularity, bucket, region string) (ret []BlobResp, err error) {
	var uri url.URL
	query := uri.Query()
	query.Add("begin", lib.FromDate(beginDate))
	query.Add("end", lib.ToDate(endDate))
	query.Add("select", "hits") // 值固定为hits
	query.Add("g", granularity) // 时间聚合粒度(5min hour day month)
	if bucket != "" {
		query.Add("$bucket", bucket)
	}
	if region != "" {
		query.Add("$region", region)
	}
	var resp []byte
	if resp, err = sendGetRequest(m.mac, "/v6/rs_chtype", query.Encode()); err != nil {
		return
	}
	if err = json.Unmarshal(resp, &ret); err != nil {
		return
	}
	return
}

// RsPut https://developer.qiniu.com/kodo/3912/rs-put
func (m *StatsManager) RsPut(beginDate, endDate, granularity, bucket, region string, ftype string) (ret []BlobResp, err error) {
	var uri url.URL
	query := uri.Query()
	query.Add("begin", lib.FromDate(beginDate))
	query.Add("end", lib.ToDate(endDate))
	query.Add("select", "hits")
	query.Add("g", granularity) // 时间聚合粒度(5min hour day month)
	if bucket != "" {
		query.Add("$bucket", bucket)
	}
	if region != "" {
		query.Add("$region", region)
	}
	if ftype != "" {
		query.Add("$ftype", ftype) // 存储类型 0-标准存储  1-低频存储  2-归档存储
	}
	var resp []byte
	if resp, err = sendGetRequest(m.mac, "/v6/rs_put", query.Encode()); err != nil {
		return
	}
	if err = json.Unmarshal(resp, &ret); err != nil {
		return
	}
	return
}
