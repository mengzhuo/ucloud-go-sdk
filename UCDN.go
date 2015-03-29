package ucloud

// ---------------- BuyUcdnTraffic ------------------

type BuyUcdnTrafficResponse struct {
	BaseResponse
}

func (r *BuyUcdnTrafficResponse) Data() interface{} {
	return r.RetCode
}

type BuyUcdnTraffic struct {
	Traffic  int    // 所购买的流量, 单位GB
	Areacode string `ucloud:"optional"` // 购买流量的区域, 枚举值为:cn: 国内; abroad: 海外

}

func (r *BuyUcdnTraffic) R() UResponse {
	return &BuyUcdnTrafficResponse{}
}

// ---------------- GetUcdnTraffic ------------------
type TrafficSet struct {
	Areacode     string  // 购买流量的区域, cn: 国内; aboard: 国外
	TrafficTotal float64 // Areacode区域内总购买流量, 单位GB
	TrafficLeft  float64 // Areacode区域内总剩余流量, 单位GB
	TrafficUsed  float64 // Areacode区域内总使用流量, 单位GB
}

type GetUcdnTrafficResponse struct {
	BaseResponse
	TrafficSet []*TrafficSet `json:"omitempty"` // 用户不同区域的流量信息, 具体结构参见TrafficSet部分

}

func (r *GetUcdnTrafficResponse) Data() interface{} {
	return r.TrafficSet
}

type GetUcdnTraffic struct {
}

func (r *GetUcdnTraffic) R() UResponse {
	return &GetUcdnTrafficResponse{}
}

// ---------------- CreateUcdnDomain ------------------

type CreateUcdnDomainResponse struct {
	BaseResponse
	DomainId string `json:"omitempty"` // 创建域名对应的域名ID，后续获取域名相关数据的操作均需要使用该ID

}

func (r *CreateUcdnDomainResponse) Data() interface{} {
	return r.DomainId
}

type CreateUcdnDomain struct {
	Domain          string   // 用于加速的域名
	SourceIps       []string `ucloud:"optional"` // 源站IP，即cdn服务器回源访问的IP地址。支持多个源站IP。多个源站IP可以表述为：SourceIps.0=1.1.1.1，SourceIps.1=2.2.2.2（如果CdnType为live，则该字段非必须，否则该字段为必须字段）
	TestUrl         string   `ucloud:"optional"` // 测试url，用于域名创建加速时的测试。（如果CdnType为live，则该字段非必须，否则该字段为必须字段）
	Areacodes       []string // CDN加速区域，目前区域代表有：cn：国内；abroad：国外。可选择多个区域，表述为：Areacodes.0=cn, Areacodes.1=aboard。表示同时使用国内和海外节点
	CdnType         string   // 加速域名的业务类型，web代表网站，stream代表视频，download代表下载，Live代表直播
	ChargeType      string   `ucloud:"optional"` // 计费方式。默认使用流量包计费。枚举值为：traffic：按流量包计费；bandwidth：按带宽计费；trafficused：代表流量后付费。（目前仅支持按流量包计费）
	LiveStreamCount int      `ucloud:"optional"` // 直播流数（CdnType为live时，该字段为必须字段）
	LiveSrcType     string   `ucloud:"optional"` // 直播类型，枚举值为：rtmppush；rtmppull；hls（CdnType为live时，该字段为必须字段）
	LifeSrcUrl      string   `ucloud:"optional"` // 用于获取流的源URL。（LiveSrcType为rtmppull/hls时，该字段为必须字段）

}

func (r *CreateUcdnDomain) R() UResponse {
	return &CreateUcdnDomainResponse{}
}

// ---------------- UpdateUcdnDomain ------------------

type UpdateUcdnDomainResponse struct {
	BaseResponse
}

func (r *UpdateUcdnDomainResponse) Data() interface{} {
	return r.RetCode
}

type UpdateUcdnDomain struct {
	DomainId       string   // 域名ID，创建加速域名时生成。
	SourceIps      []string `ucloud:"optional"` // 源站IP，即cdn服务器回源访问的IP地址。支持多个源站IP。多个源站IP可以表述为：SourceIps.0=1.1.1.1，SourceIps.1=2.2.2.2
	TestUrl        string   `ucloud:"optional"` // 测试url，用于域名创建加速时的测试。
	Areacodes      []string `ucloud:"optional"` // CDN加速区域，目前区域代表有：cn：国内；abroad：国外。可选择多个区域，表述为：”Areacodes.0=cn, Areacodes.1=aboard”，表示同时使用国内和海外节点
	CacheFileTypes []string `ucloud:"optional"` // 加速成功后需要缓存在节点服务器的静态文件类型，动态文件不支持缓存。多个文件类型，请使用：”CacheFileTypes.0=zip, CacheFileTypes.1=txt”，依赖于CacheTel参数
	CacheUrls      []string `ucloud:"optional"` // 需要缓存的文件或路径的URL。URL支持模糊匹配，不支持正则表达式。
	CacheTtl       int      `ucloud:"optional"` // 缓存文件或路径需要缓存的时间。单位：秒
	NoCacheUrl     []string `ucloud:"optional"` // 不需要缓存的文件或路径。格式同CacheUrls

}

func (r *UpdateUcdnDomain) R() UResponse {
	return &UpdateUcdnDomainResponse{}
}

// ---------------- DescribeUcdnDomain ------------------
type DomainSet struct {
	Domain          string   // 加速域名
	DomainId        string   // 加速域名ID，创建加速域名时生成
	SourceIp        []string // 源站IP，即cdn服务器回源访问的IP地址。多个IP地址可以表述为：[“1.1.1.1”, “2.2.2.2”]
	CdnType         string   // 加速域名的业务类型，枚举值：web：网站；Stream：视频；download：下载；live：直播
	ChargeType      string   // 计费方式，枚举值为：traffic：按流量包计费；bandwidth：按带宽付费；trafficused：流量后付费（目前仅支持流量包计费）
	Status          string   // 加速域名的当前状态：check：审核中；checkSuccess：审核通过；checkFaile：审核失败；enable(ing)：(正在)加速中；disable(ing)：(正在)停止加速；delete(ing)：(正在)删除加速
	Cname           string   // cdn域名，创建加速域名时生成的cdn域名，用于设置CNAME记录
	CreateTime      string   // 域名创建的时间。格式为Unix Timestamp
	VaildTime       string   // 开始分配Cname时间。格式为Unix Timestamp
	TestUrl         string   // 测试URL，用于测试加速域名
	Areaode         []string // CDN加速区域，cn：国内；abroad：国外。[ “cn”, “abroad” ]表示两者皆启用
	CacheFileType   []string // 缓存文件类型，多个文件表述为: [“zip”, “txt”]
	CacheUrl        []string // 需要缓存的文件或者路径，多个文件表述为：[ “http://Domain/*.jpg”, “http://Domain/*.png” ]
	CacheTtl        int      // 缓存文件生命周期，单位秒
	NoCacheUrl      []string // 不需要缓存的文件，格式同CacheUrl
	LiveStreamCount int      // 直播流数
	LiveSrcType     string   // 直播接入类型，rtmppush/rtmppull/hls
	LiveSrcUrl      string   // 拉去流url
	LiveStreamIds   string   // 直播流ID
	ApplyMessage    string   // 流程申请失败原因
}

type DescribeUcdnDomainResponse struct {
	BaseResponse
	TotalCount int          `json:"omitempty"` // 满足条件的条目数
	DomainSet  []*DomainSet `json:"omitempty"` // 获取的域名信息，具体结构见 DomainSet

}

func (r *DescribeUcdnDomainResponse) Data() interface{} {
	return r.DomainSet
}

type DescribeUcdnDomain struct {
	DomainIds []string `ucloud:"optional"` // 域名ID，创建加速域名时生成。默认获取账户下所有域名。
	OffSet    int      `ucloud:"optional"` // 返回数据的偏移量，默认0。
	Limit     string   `ucloud:"optional"` // 返回数据的长度，默认20。

}

func (r *DescribeUcdnDomain) R() UResponse {
	return &DescribeUcdnDomainResponse{}
}

// ---------------- GetUcdnDomainBandwidth ------------------
type BandwidthSet struct {
	Time  int     // 带宽获取的时间点，格式为Unix Timestamp
	Value float64 // 具体带宽值，单位为Mbps。其中查询区间为1天，value的值表示5分钟内的最大带宽值；查询区间为2-4天，value的值表示15分钟内的最大带宽值；查询区间为5-14天，value的值表示3-分钟内的最大带宽值；查询区间为15-30天，value的值为1小时内的带宽最大值。（如果请求参数中Daily为1，则value的值表示1天内的带宽最大值）
}

type GetUcdnDomainBandwidthResponse struct {
	BaseResponse
	Traffic      float64         `json:"omitempty"` // 从起始时间到结束时间内所使用的带宽总量，单位GB
	BandWidthSet []*BandwidthSet `json:"omitempty"` // 带宽流量实例表，具体结构见 BandwidthSet

}

func (r *GetUcdnDomainBandwidthResponse) Data() interface{} {
	return r.BandWidthSet
}

type GetUcdnDomainBandwidth struct {
	DomainId  []string `ucloud:"optional"` // 域名ID，创建加速域名时生成。
	Areacode  []string `ucloud:"optional"` // CDN加速区域，目前区域代表有：cn：国内；abroad：国外。可选择多个区域，表述为：”Areacodes.0=cn, Areacodes.1=aboard”，表示同时使用国内和海外节点
	BeginTime int      `ucloud:"optional"` // 查询的起始时间，格式为Unix Timestamp。如果有EndTime，BeginTime必须赋值。
	EndTime   int      `ucloud:"optional"` // 查询的结束时间，格式为Unix Timestamp。EndTime默认为当前时间，BeginTime默认为当前时间前一天时间。
	Daily     int      `ucloud:"optional"` // 是否按天展示带宽峰值，枚举值：0：否；1：是；默认为0

}

func (r *GetUcdnDomainBandwidth) R() UResponse {
	return &GetUcdnDomainBandwidthResponse{}
}

// ---------------- GetUcdnDomainTraffic ------------------
type GetUcdnDomainTrafficSet struct {
	Time  int     // 流量获取的时间点，格式为Unix Timestamp
	Value float64 // 查询每日流量总值，单位：GB
}

type GetUcdnDomainTrafficResponse struct {
	BaseResponse
	TrafficSet []*GetUcdnDomainTrafficSet `json:"omitempty"` // 流量实例表，具体结构见 TrafficSet

}

func (r *GetUcdnDomainTrafficResponse) Data() interface{} {
	return r.TrafficSet
}

type GetUcdnDomainTraffic struct {
	DomainId  []string `ucloud:"optional"` // 域名ID，创建加速域名时生成。
	Areacode  []string `ucloud:"optional"` // CDN加速区域，目前区域代表有：cn：国内；abroad：国外。可选择多个区域，表述为：”Areacodes.0=cn, Areacodes.1=aboard”，表示同时使用国内和海外节点
	BeginTime int      `ucloud:"optional"` // 查询的起始时间，格式为Unix Timestamp。如果有EndTime，BeginTime必须赋值。
	EndTime   int      `ucloud:"optional"` // 查询的结束时间，格式为Unix Timestamp。EndTime默认为当前时间，BeginTime默认为当前时间前一天时间。

}

func (r *GetUcdnDomainTraffic) R() UResponse {
	return &GetUcdnDomainTrafficResponse{}
}

// ---------------- GetUcdnDomainLog ------------------
type LogSet struct {
	Time  int    // 流量获取的时间点，格式为Unix Timestamp
	Value string // 具体获取日志连的URL，其中获取最近7天内的下载日志。下载日志URL中的域名后带有cn表示中国大陆，hk表示香港，ov表示国外。创建加速域名时，aboard包含hk和ov。
}

type GetUcdnDomainLogResponse struct {
	BaseResponse
	LogSet []*LogSet `json:"omitempty"` // 流量实例表，具体结构见 LogSet

}

func (r *GetUcdnDomainLogResponse) Data() interface{} {
	return r.LogSet
}

type GetUcdnDomainLog struct {
	DomainId  []string `ucloud:"optional"` // 域名ID，创建加速域名时生成。
	BeginTime int      `ucloud:"optional"` // 查询的起始时间，格式为Unix Timestamp。如果有EndTime，BeginTime必须赋值。
	EndTime   int      `ucloud:"optional"` // 查询的结束时间，格式为Unix Timestamp。EndTime默认为当前时间，BeginTime默认为当前时间前一天时间。

}

func (r *GetUcdnDomainLog) R() UResponse {
	return &GetUcdnDomainLogResponse{}
}

// ---------------- RefreshUcdnDomainCache ------------------

type RefreshUcdnDomainCacheResponse struct {
	BaseResponse
}

func (r *RefreshUcdnDomainCacheResponse) Data() interface{} {
	return r.RetCode
}

type RefreshUcdnDomainCache struct {
	DomainId string   // 域名ID，创建加速域名时生成。
	Type     string   // 刷新类型，file代表文件刷新，dir代表路径刷新
	UrlList  []string // 刷新多个URL列表时，一次最多提交30个。必须以”http://域名/”开始。目录要以”/”结尾，如刷新目录a下所有文件，格式为：http://abc.ucloud.cn/a/；如刷新文件目录a下面所有img.png文件，格式为http://abc.ucloud.cn/a/img.png。请正确提交需要刷新的域名

}

func (r *RefreshUcdnDomainCache) R() UResponse {
	return &RefreshUcdnDomainCacheResponse{}
}

// ---------------- DescribeRefreshCacheTask ------------------
type TaskSet struct {
	Domain     string   // 刷新域名
	CreateTime int      // 刷新任务创建的时间。格式为Unix Timestamp
	CheckTime  string   // 任务状态查询时间，成功后不再查询。格式为Unix Timestamp
	UrlLists   []string // 刷新提交的多条URL
	Status     string   // 刷新任务的当前状态，枚举值：success：成功；wait：等待处理；process：正在处理；failure：失败；unkonw：未知
	Percent    int      // 刷新任务执行的百分比。1代表1%，100代表100%
}

type DescribeRefreshCacheTaskResponse struct {
	BaseResponse
	TotalCount string     `json:"omitempty"` // 返回总条目数
	TaskSet    []*TaskSet `json:"omitempty"` // 刷新任务信息，具体结构见 TaskSet

}

func (r *DescribeRefreshCacheTaskResponse) Data() interface{} {
	return r.TaskSet
}

type DescribeRefreshCacheTask struct {
	DomainId  string // 域名ID，创建加速域名时生成。
	BeginTime int    `ucloud:"optional"` // 查询的起始时间，格式为Unix Timestamp。如果有EndTime，BeginTime必须赋值。
	EndTime   int    `ucloud:"optional"` // 查询的结束时间，格式为Unix Timestamp。EndTime默认为当前时间，BeginTime默认为当前时间前一天时间。
	Status    string `ucloud:"optional"` // 需要获取的内容刷新的状态，枚举值：success：成功；wait：等待处理；process：正在处理；failure：失败；unkonw：未知，默认选择所有状态
	Offset    int    `ucloud:"optional"` // 数据偏移量，默认为0
	Limit     int    `ucloud:"optional"` // 返回数据的长度，默认20。

}

func (r *DescribeRefreshCacheTask) R() UResponse {
	return &DescribeRefreshCacheTaskResponse{}
}

// ---------------- PrefetchDomainCache ------------------

type PrefetchDomainCacheResponse struct {
	BaseResponse
}

func (r *PrefetchDomainCacheResponse) Data() interface{} {
	return r.RetCode
}

type PrefetchDomainCache struct {
	DomainId string   // 域名ID，创建加速域名时生成。
	UrlList  []string // 需要刷新的多个URL列表，一次最多提交10个，必须以”http://域名/”开始。目录要以”/”结尾，如刷新目录a下所有文件，格式为：http://abc.ucloud.cn/a/；如刷新文件目录a下面所有img.png文件，格式为http://abc.ucloud.cn/a/img.png。请正确提交需要刷新的域名（点播，下载只支持每次预取一个）
	Md5      string   `ucloud:"optional"` // 大文件下载、点播支持文件的md5校验

}

func (r *PrefetchDomainCache) R() UResponse {
	return &PrefetchDomainCacheResponse{}
}

// ---------------- DescribePrefetchCacheTask ------------------
type DescribePrefetchCacheTaskSet struct {
	Domain     string   // 预取域名
	CreateTime int      // 预取任务创建的时间。格式为Unix Timestamp
	CheckTime  string   // 任务状态查询时间，成功后不再查询。格式为Unix Timestamp
	UrlLists   []string // 预取提交的多条URL
	Status     string   // 预取任务的当前状态，枚举值：success：成功；wait：等待处理；process：正在处理；failure：失败；unkonw：未知
	Percent    int      // 预取任务执行的百分比。1代表1%，100代表100%。（大文件下载、点播不支持展示精确进度，只有成功或失败）
}

type DescribePrefetchCacheTaskResponse struct {
	BaseResponse
	TotalCount string                          `json:"omitempty"` // 返回总条目数
	TaskSet    []*DescribePrefetchCacheTaskSet `json:"omitempty"` // 预取任务信息，具体结构见 TaskSet

}

func (r *DescribePrefetchCacheTaskResponse) Data() interface{} {
	return r.TaskSet
}

type DescribePrefetchCacheTask struct {
	DomainId  string // 域名ID，创建加速域名时生成。
	BeginTime int    `ucloud:"optional"` // 查询的起始时间，格式为Unix Timestamp。如果有EndTime，BeginTime必须赋值。
	EndTime   int    `ucloud:"optional"` // 查询的结束时间，格式为Unix Timestamp。EndTime默认为当前时间，BeginTime默认为当前时间前一天时间。
	Status    string `ucloud:"optional"` // 需要获取的内容刷新的状态，枚举值：success：成功；wait：等待处理；process：正在处理；failure：失败；unkonw：未知，默认选择所有状态
	Offset    int    `ucloud:"optional"` // 数据偏移量，默认为0
	Limit     int    `ucloud:"optional"` // 返回数据的长度，默认20。

}

func (r *DescribePrefetchCacheTask) R() UResponse {
	return &DescribePrefetchCacheTaskResponse{}
}

// ---------------- UpdateUcdnDomainStatus ------------------

type UpdateUcdnDomainStatusResponse struct {
	BaseResponse
}

func (r *UpdateUcdnDomainStatusResponse) Data() interface{} {
	return r.RetCode
}

type UpdateUcdnDomainStatus struct {
	DomainId string // 域名ID，创建加速域名时生成。
	Status   string // 域名状态，enable代表加速中，disable代表停止加速，delete代表删除加速。（目前仅支持删除审核失败的域名）

}

func (r *UpdateUcdnDomainStatus) R() UResponse {
	return &UpdateUcdnDomainStatusResponse{}
}

// ---------------- GetUcdnDomainPrefetchEnable ------------------

type GetUcdnDomainPrefetchEnableResponse struct {
	BaseResponse
	Enable int `json:"omitempty"` // 0表示该域名未开启预取，1表示该域名已开启预取

}

func (r *GetUcdnDomainPrefetchEnableResponse) Data() interface{} {
	return r.Enable
}

type GetUcdnDomainPrefetchEnable struct {
	DomainId string // 域名ID，创建加速域名时生成。

}

func (r *GetUcdnDomainPrefetchEnable) R() UResponse {
	return &GetUcdnDomainPrefetchEnableResponse{}
}
