package ucloud

// ---------------- CreateUlb ------------------

type CreateUlbResponse struct {
	BaseResponse
	ULBId string `json:",omitempty"` // 负载均衡实例的Id

}

func (r *CreateUlbResponse) Data() interface{} {
	return r.ULBId
}

type CreateUlb struct {
	Region  string // 数据中心，参见 数据中心列表
	ULBName string `ucloud:"optional"` // 负载均衡的名字，默认值为空

}

func (r *CreateUlb) R() UResponse {
	return &CreateUlbResponse{}
}

// ---------------- DeleteUlb ------------------

type DeleteUlbResponse struct {
	BaseResponse
}

func (r *DeleteUlbResponse) Data() interface{} {
	return r.RetCode
}

type DeleteUlb struct {
	Region string // 数据中心，参见 数据中心列表
	ULBId  string // 负载均衡实例的ID

}

func (r *DeleteUlb) R() UResponse {
	return &DeleteUlbResponse{}
}

// ---------------- DescribeUlb ------------------
type DescribUlbDataSet struct {
	ULBId         string              // 负载均衡实例的Id
	ULBName       string              // 负载均衡实例的名字
	BandwidthType int                 // 带宽模式，枚举值为：0，该负载均衡实例非共享带宽；1，该负载均衡实例使用共享带宽
	CreateTime    int                 // 负载均衡的创建时间，格式为Unix Timestamp
	IPSet         []*DescribeUlbIPSet // 负载均衡实例的外网IP信息，具体结构见下方IPSet
	VserverSet    []*VserverSet       // 负载均衡实例中存在的VServer实例列表，具体结构件下方VserverSet
}
type DescribeUlbIPSet struct {
	OperatorName string // 弹性IP的运营商信息，枚举值为：Telecom：电信IP；Unicom：联通IP；Bgp：BGP IP；International：国际IP
	EIPId        string // 弹性IP的资源ID
	EIP          string // 弹性IP的地址
}
type VserverSet struct {
	VServerId       string        // VServer实例的ID
	VServerName     string        // VServer实例的名字
	Protocol        string        // VServer实例的网络协议，枚举值为：HTTP, TCP, UDP, HTTPS
	FrontendPort    int           // VServer服务前端端口
	Method          string        // VServer的负载均衡模式，枚举值为：Roundrobin，轮询模式；Source，源地址模式
	PersistenceType string        // VServer会话保持模式，枚举值为：None，关闭会话保持；ServerInsert，自动生成；UserDefined，用户自定义
	PersistenceInfo string        // 由PersistenceType确定：None或ServerInsert时，此字段为空；UserDefined时，此字段展示用户自定义会话保持String。
	ClientTimeout   int           // 空闲连接的回收时间，单位：秒
	Status          int           // VServer的运行状态，枚举值为：0：运行正常；1：运行异常
	SSLSet          []*SSLSet     // VServer绑定的SSL证书信息，具体结构见下方SSLSet
	BackendSet      []*BackendSet // VServer后端资源信息列表，具体结构见下方BackendSet
}
type SSLSet struct {
	SSLId   string // SSL证书的ID
	SSLName string // SSL证书的名称
}
type BackendSet struct {
	BackendId    string // VServer后端的负载均衡实例ID
	ResourceType string // VServer后端的资源类型，如uhost
	ResourceId   string // VServer后端的资源ID
	ResourceName string // VServer后端的资源名称
	PrivateIP    string // VServer后端的资源内网IP
	Port         int    // VServer后端的服务端口
	Enabled      int    // 后端资源的启用状态，枚举值为：0：禁用；1：启用
	Status       int    // 后端资源的运行状态，枚举值为：0：运行正常；1：运行异常
}

type DescribeUlbResponse struct {
	BaseResponse
	TotalCount int                  `json:",omitempty"` // 满足条件的负载均衡实例的总数
	DataSet    []*DescribUlbDataSet `json:",omitempty"` // 负载均衡实例列表，具体结构见下方DataSet

}

func (r *DescribeUlbResponse) Data() interface{} {
	return r.DataSet
}

type DescribeUlb struct {
	Region string // 数据中心，参见 数据中心列表
	ULBId  string `ucloud:"optional"` // 负载均衡实例的ID，若指定则返回指定的负载均衡实例信息；若不指定则返回用户当前数据中心中所有负载均衡实例的信息
	Offset int    // 数据偏移量，默认为0
	Limit  int    // 数据分页值，默认为20

}

func (r *DescribeUlb) R() UResponse {
	return &DescribeUlbResponse{}
}

// ---------------- UpdateUlbAttribute ------------------

type UpdateUlbAttributeResponse struct {
	BaseResponse
}

func (r *UpdateUlbAttributeResponse) Data() interface{} {
	return r.RetCode
}

type UpdateUlbAttribute struct {
	Region  string // 数据中心，参见 数据中心列表
	ULBId   string // 负载均衡实例ID
	ULBName string `ucloud:"optional"` // 负载均衡的名字，若无此字段则不做修改

}

func (r *UpdateUlbAttribute) R() UResponse {
	return &UpdateUlbAttributeResponse{}
}

// ---------------- CreateVserver ------------------

type CreateVserverResponse struct {
	BaseResponse
	VServerId string `json:",omitempty"` // VServer实例的Id

}

func (r *CreateVserverResponse) Data() interface{} {
	return r.VServerId
}

type CreateVserver struct {
	Region          string // 数据中心，参见 数据中心列表
	ULBId           string // 负载均衡实例ID
	VServerName     string `ucloud:"optional"` // VServer实例名称，默认为空
	Protocol        string `ucloud:"optional"` // VServer实例的协议，枚举值为：HTTP，TCP，UDP；默认为HTTP协议
	FrontendPort    int    `ucloud:"optional"` // VServer后端端口，取值范围[1-65535]；默认值为80
	Method          string `ucloud:"optional"` // VServer负载均衡模式，枚举值为：Roundrobin，轮询；Source，源地址；默认为轮询模式
	PersistenceType string `ucloud:"optional"` // VServer会话保持方式，枚举值为：None，关闭会话保持；ServerInsert，自动生成；UserDefined，用户自定义；默认关闭会话保持。
	PersistenceInfo string `ucloud:"optional"` // 根据PersistenceType确认；None和ServerInsert：此字段无意义；UserDefined：此字段传入自定义会话保持String
	ClientTimeout   int    `ucloud:"optional"` // 空闲连接的回收时间，单位：秒；取值范围：(0，86400]，默认值为60

}

func (r *CreateVserver) R() UResponse {
	return &CreateVserverResponse{}
}

// ---------------- DeleteVserver ------------------

type DeleteVserverResponse struct {
	BaseResponse
}

func (r *DeleteVserverResponse) Data() interface{} {
	return r.RetCode
}

type DeleteVserver struct {
	Region    string // 数据中心，参见 数据中心列表
	ULBId     string // 负载均衡实例的ID
	VServerId string // VServer实例的ID

}

func (r *DeleteVserver) R() UResponse {
	return &DeleteVserverResponse{}
}

// ---------------- UpdateVserverAttribute ------------------

type UpdateVserverAttributeResponse struct {
	BaseResponse
}

func (r *UpdateVserverAttributeResponse) Data() interface{} {
	return r.RetCode
}

type UpdateVserverAttribute struct {
	Region          string // 数据中心，参见 数据中心列表
	ULBId           string // 负载均衡实例ID
	VServerId       string // VServer实例ID
	VServerName     string // VServer实例名称，若无此字段则不做修改
	Protocol        string // VServer协议，枚举值为: HTTP, TCP, UDP. 若无此字段则不做修改
	Method          string // VServer负载均衡算法, 枚举值为: Roundrobin, 轮询模式; Source,源IP模式. 若无此字段则不做修改
	PersistenceType string // VServer会话保持模式, 枚举值为:None, 关闭会话保持;ServerInsert, 自动生成;UserDefined, 用户自定义.若无此字段则不做修改
	PersistenceInfo string // 根据PersistenceType确定:None或ServerInsert, 此字段无意义;UserDefined, 则此字段传入用户自定义会话保持String.若无此字段则不做修改
	ClientTimeout   string // 空闲连接的回收时间, 单位: 秒.取值范围: (0, 86400], 默认值: 60秒.若无此字段,则不做修改

}

func (r *UpdateVserverAttribute) R() UResponse {
	return &UpdateVserverAttributeResponse{}
}

// ---------------- AllocateBackend ------------------

type AllocateBackendResponse struct {
	BaseResponse
	BackendId string `json:",omitempty"` // 所添加的后端资源ID，（为ULB系统中使用，与资源自身ID无关）

}

func (r *AllocateBackendResponse) Data() interface{} {
	return r.BackendId
}

type AllocateBackend struct {
	Region       string // 数据中心，参见 数据中心列表
	ULBId        string // 负载均衡实例的ID
	VServerId    string // VServer实例的ID
	ResourceType string // 所添加的后端资源的类型，如UHost
	ResourceId   string // 所添加的后端资源的资源ID
	Port         int    `ucloud:"optional"` // 所添加的后端资源服务端口，取值范围[1-65535]，默认80
	Enabled      int    `ucloud:"optional"` // 后端实例状态开关，枚举值：1：启动；0：禁用默认为启用

}

func (r *AllocateBackend) R() UResponse {
	return &AllocateBackendResponse{}
}

// ---------------- ReleaseBackend ------------------

type ReleaseBackendResponse struct {
	BaseResponse
}

func (r *ReleaseBackendResponse) Data() interface{} {
	return r.RetCode
}

type ReleaseBackend struct {
	Region    string // 数据中心，参见 数据中心列表
	ULBId     string // 负载均衡实例的ID
	BackendId string // 后端资源实例的ID(ULB后端ID，非资源自身ID)

}

func (r *ReleaseBackend) R() UResponse {
	return &ReleaseBackendResponse{}
}

// ---------------- UpdateBackendAttribute ------------------

type UpdateBackendAttributeResponse struct {
	BaseResponse
}

func (r *UpdateBackendAttributeResponse) Data() interface{} {
	return r.RetCode
}

type UpdateBackendAttribute struct {
	Region    string // 数据中心，参见 数据中心列表
	ULBId     string // 负载均衡资源ID
	BackendId string // 后端资源实例的ID(ULB后端ID，非资源自身ID)
	Port      int    `ucloud:"optional"` // 后端资源服务端口，取值范围[1-65535]
	Enabled   int    `ucloud:"optional"` // 后端实例状态开关，枚举值：1：启动；0：禁用若无此字段，则不做修改

}

func (r *UpdateBackendAttribute) R() UResponse {
	return &UpdateBackendAttributeResponse{}
}

// ---------------- CreateSsl ------------------

type CreateSslResponse struct {
	BaseResponse
	ULBId string `json:",omitempty"` // SSL证书的Id

}

func (r *CreateSslResponse) Data() interface{} {
	return r.ULBId
}

type CreateSsl struct {
	Region     string // 数据中心，参见 数据中心列表
	SSLContent string // SSL证书的内容
	SSLName    string `ucloud:"optional"` // SSL证书的名字，默认值为空
	SSLType    int    `ucloud:"optional"` // 所添加的SSL证书类型，目前只支持0：Pem格式

}

func (r *CreateSsl) R() UResponse {
	return &CreateSslResponse{}
}

// ---------------- DeleteSsl ------------------

type DeleteSslResponse struct {
	BaseResponse
}

func (r *DeleteSslResponse) Data() interface{} {
	return r.RetCode
}

type DeleteSsl struct {
	Region string // 数据中心，参见 数据中心列表
	SSLId  string // SSL证书的ID

}

func (r *DeleteSsl) R() UResponse {
	return &DeleteSslResponse{}
}

// ---------------- BindSsl ------------------

type BindSslResponse struct {
	BaseResponse
	ULBId string `json:",omitempty"` // SSL证书的Id

}

func (r *BindSslResponse) Data() interface{} {
	return r.ULBId
}

type BindSsl struct {
	Region    string // 数据中心，参见 数据中心列表
	SSLId     string // SSL证书的Id
	ULBId     string // 所绑定ULB实例ID
	VServerId string // 所绑定VServer实例ID

}

func (r *BindSsl) R() UResponse {
	return &BindSslResponse{}
}

// ---------------- DescribeSsl ------------------
type DescribeSslDataSet struct {
	SSLId      string // 证书的ID
	SSLName    string // 证书的名字
	SSLType    string // 证书的类型，枚举值为：0：Pem格式
	SSLContent string // 证书的内容
	HashValue  string // 证书的HASH值
	CreateTime int    // 创建时间，格式为Unix Timestamp
	VServerId  string // 所绑定VServer的资源ID
}

type DescribeSslResponse struct {
	BaseResponse
	TotalCOunt string                `json:",omitempty"` // 满足条件的SSL证书总数
	DataSet    []*DescribeSslDataSet `json:",omitempty"` // SSL证书详细信息，具体结构见DataSet

}

func (r *DescribeSslResponse) Data() interface{} {
	return r.DataSet
}

type DescribeSsl struct {
	Region string // 数据中心，参见 数据中心列表
	SSLId  string `ucloud:"optional"` // SSL证书的Id
	Limit  int    `ucloud:"optional"` // 数据分页值，默认为20
	Offset int    `ucloud:"optional"` // 数据偏移量，默认值为0

}

func (r *DescribeSsl) R() UResponse {
	return &DescribeSslResponse{}
}

// ---------------- CreatePolicyGroup ------------------

type CreatePolicyGroupResponse struct {
	BaseResponse
	GroupId string `json:",omitempty"` // 内容转发策略组的Id

}

func (r *CreatePolicyGroupResponse) Data() interface{} {
	return r.GroupId
}

type CreatePolicyGroup struct {
	Region    string // 数据中心，参见 数据中心列表
	GroupName string `ucloud:"optional"` // 内容转发策略组名称，默认为空

}

func (r *CreatePolicyGroup) R() UResponse {
	return &CreatePolicyGroupResponse{}
}

// ---------------- DeletePolicyGroup ------------------

type DeletePolicyGroupResponse struct {
	BaseResponse
}

func (r *DeletePolicyGroupResponse) Data() interface{} {
	return r.RetCode
}

type DeletePolicyGroup struct {
	Region  string // 数据中心，参见 数据中心列表
	GroupId string // 内容转发策略组ID

}

func (r *DeletePolicyGroup) R() UResponse {
	return &DeletePolicyGroupResponse{}
}

// ---------------- DescribePolicyGroup ------------------
type DescribePolicyGroupDataSet struct {
	GroupId   string       // 内容转发策略组ID
	GroupName string       // 内容转发策略组名称
	PolicySet []*PolicySet // 内容转发策略组详细信息，具体结构见PolicySet
}
type PolicySet struct {
	PolicyId   string                           // 内容转发策略组ID
	Type       string                           // 内容转发匹配字段的类型，当前只支持按域名转发。枚举值为：Domain，按域名转发
	Match      string                           // 内容转发匹配字段
	VServerId  string                           // 内容转发策略组ID应用的VServer实例的ID
	BackendSet []*DescribePolicyGroupBackendSet // 内容转发策略组ID所应用的后端资源列表，具体结构见”BackendSet”
}
type DescribePolicyGroupBackendSet struct {
	BackendId string // 后端资源实例的ID
	PrivateIP string // 后端资源实例的内网IP
	Port      int    // 后端资源实例的服务端口
}

type DescribePolicyGroupResponse struct {
	BaseResponse
	DataSet []*DescribePolicyGroupDataSet `json:",omitempty"` // 内容转发策略组列表，具体结构见DataSet

}

func (r *DescribePolicyGroupResponse) Data() interface{} {
	return r.DataSet
}

type DescribePolicyGroup struct {
	Region  string // 数据中心，参见 数据中心列表
	GroupId string `ucloud:"optional"` // 内容转发策略组ID
	Offset  int    `ucloud:"optional"` // 数据偏移量，默认值为0
	Limit   int    `ucloud:"optional"` // 数据分页值，默认为20

}

func (r *DescribePolicyGroup) R() UResponse {
	return &DescribePolicyGroupResponse{}
}

// ---------------- UpdatePolicyGroupAttribute ------------------

type UpdatePolicyGroupAttributeResponse struct {
	BaseResponse
}

func (r *UpdatePolicyGroupAttributeResponse) Data() interface{} {
	return r.RetCode
}

type UpdatePolicyGroupAttribute struct {
	Region    string // 数据中心，参见 数据中心列表
	GroupId   string // 内容转发策略组ID
	GroupName string `ucloud:"optional"` // 修改策略转发组名称

}

func (r *UpdatePolicyGroupAttribute) R() UResponse {
	return &UpdatePolicyGroupAttributeResponse{}
}

// ---------------- CreatePolicy ------------------

type CreatePolicyResponse struct {
	BaseResponse
	PolicyId string `json:",omitempty"` // 内容转发策略ID

}

func (r *CreatePolicyResponse) Data() interface{} {
	return r.PolicyId
}

type CreatePolicy struct {
	Region    string   // 数据中心，参见 数据中心列表
	GroupId   string   // 内容转发策略组ID
	Match     string   // 内容转发匹配字段
	ULBId     string   // 需要添加内容转发策略的负载均衡实例ID
	VServerId string   // 需要添加内容转发策略的VServer实例ID
	BackendId []string // 内容转发策略应用的后端资源实例的ID
	Type      string   `ucloud:"optional"` // 内容转发匹配字段的类型，当前仅支持按域名转发。枚举值：Domain

}

func (r *CreatePolicy) R() UResponse {
	return &CreatePolicyResponse{}
}

// ---------------- DeletePolicy ------------------

type DeletePolicyResponse struct {
	BaseResponse
}

func (r *DeletePolicyResponse) Data() interface{} {
	return r.RetCode
}

type DeletePolicy struct {
	Region   string // 数据中心，参见 数据中心列表
	GroupId  string // 内容转发策略组ID
	PolicyId string `ucloud:"optional"` // 内容转发策略ID

}

func (r *DeletePolicy) R() UResponse {
	return &DeletePolicyResponse{}
}
