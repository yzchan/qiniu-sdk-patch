## Dora 计量接口

### HOST

```http
https://stats-dora.qiniuapi.com
```

### 单个计费项按天的计量接口

请求包:

```http
GET /v3/statistic/<item>?start=<start>&end=<end>
Authorization: Qiniu <MacToken>
```

返回包:

```http
200 OK

[
    {
        "time": "<day>",
        "values": {
            "value": <value>
        }
    },
    {
        "time": "<day>",
        "values": {
            "value": <value>
        }
    },

	...
]
```

#### 查询参数

*  `item` 计费项参数
*  `start` `end` 分别为查询的开始日期和结束日志，格式为 `20160102150405`, 左闭右开

#### 返回量格式

*  `day` RFC3339
*  `value` 长整数


#### 计费项

| 账单名称                             | 用于查询的计费项（item）               |
|----------------------------------|------------------------------|
| 文件HASH值                          | qhash-fsize                  |
| 文本文件合并                           | concat-fsize                 |
| 多文件压缩                            | mkzip-fsize                  |
| MD转HTML                          | md2html                      |
| 资源下载二维码                          | qrcode                       |
|                                  |                              |
| 图片 exif 处理量                      | exif-fsize                   |
| 图片 imageInfo 处理量                 | imageinfo-fsize              |
| 图片 imageMogr2 处理量                | imagemogr2-fsize             |
| 图片 imageMogr 处理量                 | imagemogr-fsize              |
| 图片 imageView2 处理量                | imageview2-fsize             |
| 图片 imageView 处理量                 | imageview-fsize              |
| 图片 watermark 处理量                 | watermark-fsize              |
| 图片 mgjThumb 处理量                  | mgjthumb-fsize               |
| 图片 roundPic 处理量                  | roundpic-fsize               |
|                                  |                              |
| 音频转码                             | audio                        |
| 视频转码4K                           | 4k                           |
| 视频转码2K                           | 2k                           |
| 视频转码hd                           | hd                           |
| 视频转码sd                           | sd                           |
| 视频转码sd480                        | sd480                        |
| 视频转码sd240                        | sd240                        |
|                                  |                              |
| H265视频转码4K                       | h265-4k                      |
| H265视频转码2K                       | h265-2k                      |
| H265视频转码HD                       | h265-hd                      |
| H265视频转码SD                       | h265-sd                      |
| H265视频转码SD480                    | h265-sd480                   |
| H265视频转码SD240                    | h265-sd240                   |
|                                  |                              |
| 高帧率视频转码2K                        | 2k-hfr                       |
| 高帧率视频转码4K                        | 4k-hfr                       |
| 高帧率视频转码HD                        | hd-hfr                       |
| 高帧率视频转码SD                        | sd-hfr                       |
| 高帧率视频转码SD480                     | sd480-hfr                    |
| 高帧率视频转码SD240                     | sd240-hfr                    |
|                                  |                              |
| 高帧率H265视频转码4K                    | h265-4k-hfr                  |
| 高帧率H265视频转码2K                    | h265-2k-hfr                  |
| 高帧率H265视频转码HD                    | h265-hd-hfr                  |
| 高帧率H265视频转码SD                    | h265-sd-hfr                  |
| 高帧率H265视频转码SD480                 | h265-sd480-hfr               |
| 高帧率H265视频转码SD240                 | h265-sd240-hfr               |
|                                  |                              |
| 倍速转码-视频转码SD240                   | avfast-sd240                 |
| 倍速转码-视频转码SD480                   | avfast-sd480                 |
| 倍速转码-视频转码SD                      | avfast-sd                    |
| 倍速转码-视频转码HD                      | avfast-hd                    |
| 倍速转码-视频转码2K                      | avfast-2k                    |
| 倍速转码-视频转码4K                      | avfast-4k                    |
| 倍速转码-高帧率视频转码SD240                | avfast-sd240-hfr             |
| 倍速转码-高帧率视频转码SD480                | avfast-sd480-hfr             |
| 倍速转码-高帧率视频转码SD                   | avfast-sd-hfr                |
| 倍速转码-高帧率视频转码HD                   | avfast-hd-hfr                |
| 倍速转码-高帧率视频转码2K                   | avfast-2k-hfr                |
| 倍速转码-高帧率视频转码4K                   | avfast-4k-hfr                |
| 倍速转码-H265视频转码SD240               | avfast-h265-sd240            |
| 倍速转码-H265视频转码SD480               | avfast-h265-sd480            |
| 倍速转码-H265视频转码SD                  | avfast-h265-sd               |
| 倍速转码-H265视频转码HD                  | avfast-h265-hd               |
| 倍速转码-H265视频转码2K                  | avfast-h265-2k               |
| 倍速转码-H265视频转码4K                  | avfast-h265-4k               |
| 倍速转码-高帧率H265视频转码SD240            | avfast-h265-sd240-hfr        |
| 倍速转码-高帧率H265视频转码SD480            | avfast-h265-sd480-hfr        |
| 倍速转码-高帧率H265视频转码SD               | avfast-h265-sd-hfr           |
| 倍速转码-高帧率H265视频转码HD               | avfast-h265-hd-hfr           |
| 倍速转码-高帧率H265视频转码2K               | avfast-h265-2k-hfr           |
| 倍速转码-高帧率H265视频转码4K               | avfast-h265-4k-hfr           |
|                                  |                              |
| 音视频转封装                           | av-copy                      |
| 实时音视频转封装                         | avvod-copy                   |
|                                  |                              |
| 实时音频转码                           | avvod-audio                  |
| 实时视频转码4K                         | avvod-4k                     |
| 实时视频转码2K                         | avvod-2k                     |
| 实时视频转码HD                         | avvod-hd                     |
| 实时视频转码SD                         | avvod-sd                     |
| 实时视频转码SD480                      | avvod-sd480                  |
| 实时视频转码SD240                      | avvod-sd240                  |
|                                  |                              |
| 锐智视频转码4K                         | avsmart-4k                   |
| 锐智视频转码2K                         | avsmart-2k                   |
| 锐智视频转码HD                         | avsmart-hd                   |
| 锐智视频转码SD                         | avsmart-sd                   |
|                                  |                              |
| API avinfo                       | avinfo                       |
| imageAve请求                       | imageave                     |
| 主动调用图片瘦身                         | imageslim                    |
| CDN 自动图片瘦身                       | imageslim-auto               |
| API vframe(目前合并vframe和vsample计量) | snap                         |
| 盲水印添加                            | bwm-encode                   |
| 盲水印提取                            | bwm-decode                   |
| 动图合成                             | animate                      |
|                                  |                              |
| 图片鉴暴恐调用量 - 机器智能审核                | terror                       |
| 图片政治人物调用量 - 机器智能审核               | politician                   |
| 图片鉴黄调用量 - 机器智能审核                 | pulp                         |
| 图片广告识别调用量 - 机器智能审核               | ads                          |
|                                  |                              |
| 直播音频反垃圾                          | live-audio-antispam          |
| 直播视频鉴黄                           | live-pulp                    |
| 直播视频鉴暴恐                          | live-terror                  |
| 直播视频敏感人物识别                       | live-politician              |
| 直播视频图文违规识别                       | live-ads                     |
|                                  |                              |
| 人脸检测                             | face-detect                  |
| 1:1人脸比对                          | face-sim                     |
| 1:N人脸比对                          | face-group-search            |
|                                  |                              |
| 以图搜图                             | image-group-search           |
| OCR身份证识别                         | ocr-idcard                   |
|                                  |                              |
| 图普广告请求确定部分                       | tupu-ad-certain              |
| 图普广告请求不确定部分                      | tupu-ad-depend               |
| 图普广告增强版请求确定部分                    | tupu-ad-plus-certain         |
| 图普广告增强版请求不确定部分                   | tupu-ad-plus-depend          |
| 图普鉴黄请求确定部分                       | tupu-nrop-certain            |
| 图普鉴黄请求不确定部分                      | tupu-nrop-depend             |
| 图普鉴暴恐请求确定部分                      | tupu-terror-certain          |
| 图普鉴暴恐请求不确定部分                     | tupu-terror-depend           |
| 视频鉴黄请求                           | tupu-video                   |
| 阿塔科技图片鉴黄不确定部分                    | atar-pulp-certain            |
| 阿塔科技图片鉴黄确定部分                     | atar-pulp-depend             |
| 数美垃圾文本识别                         | sm-text-spam                 |
| 深图图片鉴黄确定部分                       | deepirugc-certain            |
| 深图图片鉴黄不确定部分                      | deepirugc-depend             |
| TuSDK 人脸特征点标识                    | tusdk-face-landmark          |
| TuSDK 人脸检测                       | tusdk-face-detection         |
| 网易易盾文本反垃圾                        | netease-ydtext               |
| 网易广告过滤                           | netease-image-ad             |
| 网易易盾鉴黄                           | netease-image-porn           |
| 阅面人脸识别比对                         | ym-face-analyze-verification |
| 阅面人脸属性识别                         | ym-face-analyze-attributes   |
| 阅面人脸关键点定位                        | ym-face-analyze-landmarks    |
| 蒲公英APP分析                         | pgy-app-parse                |
| 达观垃圾评论过滤                         | dg-spam-filter               |
| 达观文本鉴黄鉴政                         | dg-content-audit             |
| 阿里文本反垃圾                          | ali-textscan                 |
| 阿里音频审核                           | ali-audio                    |
| 阿里广告识别                           | ali-ad                       |
|                                  |                              |
| Versa 风格迁移服务                     | mkr-style-trans              |
| Versa 人像分割服务                     | mkr-seg-human                |
| Versa 智能填充服务                     | mkr-inpainting               |
| Versa 实例分割服务                     | mkr-seg-ins                  |
|                                  |                              |
| 光线活体检测                           | face-flashlive               |
| 防翻拍活体检测                          | face-piclive                 |
| 身份证二要素                           | idcard-auth                  |
| 动作活体检测                           | face-actlive                 |
| 人脸比对                             | face-compare                 |
| 公安核验                             | face-hdphotoauth             |