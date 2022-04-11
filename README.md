# qiniu-sdk-patch

对 [七牛开发者中心](https://developer.qiniu.com/) 中包含但是 [七牛SDK(V7)](https://github.com/qiniu/go-sdk) 中并未封装的部分方法进行补充。

## 补充的方法列表

### dcdn 全站加速

[动态加速升级为全站加速](https://developer.qiniu.com/dcdn/10489/dynamic-acceleration-upgraded-to-total-station) [官方文档](https://developer.qiniu.com/dcdn/10753/dcdn-traffic-bandwidth) 尚未更新（2022-04-11）

- 全站加速静态流量
- 全站加速动态流量
- 全站加速动态请求次数

### 对象存储 数据统计接口

[数据统计接口](https://developer.qiniu.com/kodo/3906/statistic-interface)

- [space](https://developer.qiniu.com/kodo/3908/statistic-space) 获取标准存储的存储量统计
- [count](https://developer.qiniu.com/kodo/3914/count) 获取标准存储的文件数量统计
- [space_line](https://developer.qiniu.com/kodo/3910/space-line) 获取低频存储的存储量统计
- [count_line](https://developer.qiniu.com/kodo/3915/count-line) 获取低频存储的文件数量统计
- [space_archive](https://developer.qiniu.com/kodo/6462/space-archive) 获取归档存储的存储量统计
- [count_archive](https://developer.qiniu.com/kodo/6463/count-archive) 获取归档存储的文件数量统计
- [blob_transfer](https://developer.qiniu.com/kodo/3911/blob-transfer) 获取跨区域同步流量统计
- [rs_chtype](https://developer.qiniu.com/kodo/3913/rs-chtype) 获取存储类型请求次数统计
- [blob_io](https://developer.qiniu.com/kodo/3820/blob-io) 获取外网流出流量统计，CDN回源流出流量统计，数据读取统计， GET请求次数统计
- [rs_put](https://developer.qiniu.com/kodo/3912/rs-put) 获取PUT请求次数统计

### Dora 计量接口

[Dora计量接口文档](Dora.md)这个接口在开发者中心没有找到。