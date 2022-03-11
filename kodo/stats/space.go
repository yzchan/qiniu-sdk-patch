package stats

import (
	"encoding/json"
	"errors"

	"github.com/google/go-querystring/query"
	"github.com/yzchan/qiniu-sdk-patch/kodo"
	"github.com/yzchan/qiniu-sdk-patch/lib"
)

type SpaceReq struct {
	Begin       string `url:"begin"`
	End         string `url:"end"`
	Granularity string `url:"g"`
	Bucket      string `url:"bucket,omitempty"`
	Region      string `url:"region,omitempty"`
	NoPreDel    int    `url:"no_predel,omitempty"`
	OnlyPreDel  int    `url:"only_predel,omitempty"`
}

type SpaceResp struct {
	Times []int64 `json:"times"`
	Datas []int64 `json:"datas"`
}

// Space https://developer.qiniu.com/kodo/3908/statistic-space
func (m *StatsManager) Space(beginDate, endDate, granularity, bucket, region string) (ret SpaceResp, err error) {
	return m.SpaceRequest("/v6/space", SpaceReq{beginDate, endDate, granularity, bucket, region, 0, 0})
}

// SpaceLine https://developer.qiniu.com/kodo/3910/space-line
func (m *StatsManager) SpaceLine(beginDate, endDate, granularity, bucket, region string) (ret SpaceResp, err error) {
	return m.SpaceRequest("/v6/space_line", SpaceReq{beginDate, endDate, granularity, bucket, region, 0, 0})
}

// SpaceArchive https://developer.qiniu.com/kodo/6462/space-archive
func (m *StatsManager) SpaceArchive(beginDate, endDate, granularity, bucket, region string) (ret SpaceResp, err error) {
	return m.SpaceRequest("/v6/space", SpaceReq{beginDate, endDate, granularity, bucket, region, 0, 0})
}

// Count https://developer.qiniu.com/kodo/3914/count
func (m *StatsManager) Count(beginDate, endDate, granularity, bucket, region string) (ret SpaceResp, err error) {
	return m.SpaceRequest("/v6/count", SpaceReq{beginDate, endDate, granularity, bucket, region, 0, 0})
}

// CountLine https://developer.qiniu.com/kodo/3915/count-line
func (m *StatsManager) CountLine(beginDate, endDate, granularity, bucket, region string) (ret SpaceResp, err error) {
	return m.SpaceRequest("/v6/count_line", SpaceReq{beginDate, endDate, granularity, bucket, region, 0, 0})
}

// CountArchive https://developer.qiniu.com/kodo/6463/count-archive
func (m *StatsManager) CountArchive(beginDate, endDate, granularity, bucket, region string) (ret SpaceResp, err error) {
	return m.SpaceRequest("/v6/count_archive", SpaceReq{beginDate, endDate, granularity, bucket, region, 0, 0})
}

func (m *StatsManager) SpaceRequest(path string, options SpaceReq) (ret SpaceResp, err error) {
	options.Begin = lib.FromDate(options.Begin)
	options.End = lib.ToDate(options.End)
	v, _ := query.Values(options)
	var resp []byte
	if resp, err = sendGetRequest(m.mac, path, v.Encode()); err != nil {
		return
	}
	if err = json.Unmarshal(resp, &ret); err != nil {
		return
	}
	return
}

func (m *StatsManager) GetStorageUsage(fType int, beginDate, endDate, granularity, bucket, region string) (ret SpaceResp, err error) {
	if fType == kodo.Standard {
		return m.Space(beginDate, endDate, granularity, bucket, region)
	} else if fType == kodo.LowUsage {
		return m.SpaceLine(beginDate, endDate, granularity, bucket, region)
	} else if fType == kodo.Archive {
		return m.SpaceArchive(beginDate, endDate, granularity, bucket, region)
	}
	return SpaceResp{}, errors.New("invalid fType")
}

func (m *StatsManager) GetFileCount(fType int, beginDate, endDate, granularity, bucket, region string) (ret SpaceResp, err error) {
	if fType == kodo.Standard {
		return m.Count(beginDate, endDate, granularity, bucket, region)
	} else if fType == kodo.LowUsage {
		return m.CountLine(beginDate, endDate, granularity, bucket, region)
	} else if fType == kodo.Archive {
		return m.CountArchive(beginDate, endDate, granularity, bucket, region)
	}
	return SpaceResp{}, errors.New("invalid fType")
}
