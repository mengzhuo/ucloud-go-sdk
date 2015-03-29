package ucloud

// ---------------- AllocateEip ------------------
type EIPSet struct {
	EIPId   string     // 申请到的EIP资源ID
	EIPAddr []*EIPAddr // 申请到的IPv4地址（双线拥有两个IP地址）
}
type EIPAddr struct {
	OperatorName string // 运营商信息
	IP           string // IP地址
}

type AllocateEipResponse struct {
	BaseResponse
	EIPSet []*EIPSet `json:"omitempty"` // 申请到的EIP资源详情

}

func (r *AllocateEipResponse) Data() interface{} {
	return r.EIPSet
}

type AllocateEip struct {
	Region       string // 数据中心，参见 数据中心列表
	OperatorName string // 弹性IP的线路如下:电信: Telecom联通: Unicom国际: InternationalBGP: Bgp双线: Duplet各数据中心允许的线路参数如下：cn-east-01: Telecom, Unicom, Dupletcn-south-01: Telecom, Unicom, Dupletcn-north-01: Bgpcn-north-02: Bgpcn-north-03: Bgphk-01: Internationalus-west-01: International
	Bandwidth    int    // 弹性IP的外网带宽，单位为Mbps，范围 [0-800]；共享带宽模式必须指定0M带宽，非共享带宽模式必须指定非0M带宽
	ChargeType   string `ucloud:"optional"` // 计费模式，枚举值为：Year，按年付费；Month，按月付费；Dynamic，按需付费(需开启权限)；Trial，试用(需开启权限)默认为按月付费
	Quantity     int    `ucloud:"optional"` // 购买时长，默认: 1

}

func (r *AllocateEip) R() UResponse {
	return &AllocateEipResponse{}
}

// ---------------- DescribeEip ------------------
type DescribeEIPSet struct {
	EIPId         string             // 弹性IP的资源ID
	Weight        int                // 外网出口权重，默认为50，范围[0-100]
	BandwidthType int                // 带宽模式，枚举值为：0：非共享带宽模式，1：共享带宽模式
	Bandwidth     int                // 弹性IP的带宽，单位为Mbps，当BandwidthType=1时，该处显示为共享带宽值。
	Status        string             // 弹性IP的资源绑定状态，枚举值为：used：已绑定，free：未绑定
	ChargeType    string             // 计费模式，枚举值为：Year，按年付费；Month，按月付费；Dynamic，按小时付费；Trial，试用
	CreateTime    int                // 弹性IP的创建时间，格式为Unix Timestamp
	ExpireTime    int                // 弹性IP的到期时间，格式为Unix Timestamp
	Resource      []*Resource        // 弹性IP的详细信息列表，具体结构见下方Resource
	EIPAddr       []*DescribeEIPAddr // 弹性IP的详细信息列表，具体结构见下方EIPAddr
}
type Resource struct {
	ResourceType string // 已绑定的资源类型
	ResourceName string // 已绑定的资源名称
	ResourceId   string // 已绑定资源的资源ID
}
type DescribeEIPAddr struct {
	OperatorName string // 弹性IP的运营商信息，枚举值为：Telecom：电信IPUnicom：联通IPDuplet：双线IP（电信+联通）Bgp：BGP IPInternational：国际IP
	IP           string // 弹性IP地址
}

type DescribeEipResponse struct {
	BaseResponse
	TotalCount int               `json:"omitempty"` // 满足条件的弹性IP总数
	EIPSet     []*DescribeEIPSet `json:"omitempty"` // 弹性IP列表，每项参数详见DataSet

}

func (r *DescribeEipResponse) Data() interface{} {
	return r.EIPSet
}

type DescribeEip struct {
	Region string   // 数据中心，参见 数据中心列表
	EIPIds []string `ucloud:"optional"` // 弹性IP的资源ID如果为空，则返回当前Region中符合条件的的所有EIP
	Offset int      `ucloud:"optional"` // 数据偏移量，默认为0
	Limit  int      `ucloud:"optional"` // 数据分页值，默认为20

}

func (r *DescribeEip) R() UResponse {
	return &DescribeEipResponse{}
}

// ---------------- ReleaseEip ------------------

type ReleaseEipResponse struct {
	BaseResponse
}

func (r *ReleaseEipResponse) Data() interface{} {
	return r.RetCode
}

type ReleaseEip struct {
	Region string // 数据中心，参见 数据中心列表
	EIPId  string // 弹性IP的资源ID

}

func (r *ReleaseEip) R() UResponse {
	return &ReleaseEipResponse{}
}

// ---------------- BindEip ------------------

type BindEipResponse struct {
	BaseResponse
}

func (r *BindEipResponse) Data() interface{} {
	return r.RetCode
}

type BindEip struct {
	Region       string // 数据中心，参见 数据中心列表
	EIPId        string // 弹性IP的资源Id
	ResourceType string // 弹性IP请求绑定的资源类型，枚举值为：uhost：云主机；vrouter：虚拟路由器；ulb，负载均衡器
	ResourceId   string // 弹性IP请求绑定的资源ID

}

func (r *BindEip) R() UResponse {
	return &BindEipResponse{}
}

// ---------------- UnbindEip ------------------

type UnbindEipResponse struct {
	BaseResponse
}

func (r *UnbindEipResponse) Data() interface{} {
	return r.RetCode
}

type UnbindEip struct {
	Region       string // 数据中心，参见 数据中心列表
	EIPId        string // 弹性IP的资源Id
	ResourceType string // 弹性IP请求解绑的资源类型，枚举值为：uhost：云主机；vrouter：虚拟路由器；ulb，负载均衡器
	ResourceId   string // 弹性IP请求解绑的资源ID

}

func (r *UnbindEip) R() UResponse {
	return &UnbindEipResponse{}
}

// ---------------- ModifyEipBandwidth ------------------

type ModifyEipBandwidthResponse struct {
	BaseResponse
}

func (r *ModifyEipBandwidthResponse) Data() interface{} {
	return r.RetCode
}

type ModifyEipBandwidth struct {
	Region    string // 数据中心，参见 数据中心列表
	EIPId     string // 弹性IP的资源ID
	Bandwidth int    // 弹性IP的外网带宽，单位为Mbps，范围 [0-800]

}

func (r *ModifyEipBandwidth) R() UResponse {
	return &ModifyEipBandwidthResponse{}
}

// ---------------- ModifyEipWeight ------------------

type ModifyEipWeightResponse struct {
	BaseResponse
}

func (r *ModifyEipWeightResponse) Data() interface{} {
	return r.RetCode
}

type ModifyEipWeight struct {
	Region string // 数据中心，参见 数据中心列表
	EIPId  string // 弹性IP的资源ID
	Weight int    // 外网出口权重，范围[0-100]

}

func (r *ModifyEipWeight) R() UResponse {
	return &ModifyEipWeightResponse{}
}

// ---------------- GetEipPrice ------------------
type GetEipPriceSet struct {
	ChargeType    string  // 弹性IP计费类型
	Price         float64 // 弹性IP价格
	PurchaseValue int     // 资源有效期，以Unix Timestamp表示
}

type GetEipPriceResponse struct {
	BaseResponse
	PriceSet []*GetEipPriceSet `json:"omitempty"` // 弹性IP价格详情

}

func (r *GetEipPriceResponse) Data() interface{} {
	return r.PriceSet
}

type GetEipPrice struct {
	Region       string // 数据中心，参见 数据中心列表
	OperatorName string // 弹性IP的线路如下:电信: Telecom联通: Unicom国际: InternationalBGP: Bgp双线: Duplet各数据中心允许的线路参数如下：cn-east-01: Telecom, Unicom, Dupletcn-south-01: Telecom, Unicom, Dupletcn-north-01: Bgpcn-north-02: Bgpcn-north-03: Bgphk-01: Internationalus-west-01: International
	Bandwidth    int    // 弹性IP的外网带宽，单位为Mbps，范围 [0-800]；
	ChargeType   string `ucloud:"optional"` // 计费模式，枚举值为：Year，按年付费；Month，按月付费；Dynamic，按需付费(需开启权限)；默认为获取三种价格

}

func (r *GetEipPrice) R() UResponse {
	return &GetEipPriceResponse{}
}

// ---------------- AllocateVip ------------------

type AllocateVipResponse struct {
	BaseResponse
	DataSet []string `json:"omitempty"` // 申请到的VIP资源地址

}

func (r *AllocateVipResponse) Data() interface{} {
	return r.DataSet
}

type AllocateVip struct {
	Region string // 数据中心，参见 数据中心列表
	Count  int    `ucloud:"optional"` // 申请数量，默认: 1

}

func (r *AllocateVip) R() UResponse {
	return &AllocateVipResponse{}
}

// ---------------- DescribeVip ------------------

type DescribeVipResponse struct {
	BaseResponse
	DataSet []string `json:"omitempty"` // 内网VIP地址列表

}

func (r *DescribeVipResponse) Data() interface{} {
	return
}

type DescribeVip struct {
	Region string // 数据中心，参见 数据中心列表

}

func (r *DescribeVip) R() UResponse {
	return &DescribeVipResponse{}
}

// ---------------- ReleaseVip ------------------

type ReleaseVipResponse struct {
	BaseResponse
}

func (r *ReleaseVipResponse) Data() interface{} {
	return
}

type ReleaseVip struct {
	Region string // 数据中心，参见 数据中心列表
	VIP    string // 内网VIP的IP地址

}

func (r *ReleaseVip) R() UResponse {
	return &ReleaseVipResponse{}
}

// ---------------- DescribeSecurityGroup ------------------
type DescribeSecurityGroupDataSet struct {
	GroupId    string  // 防火墙组的资源ID
	GroupName  string  // 防火墙组的名称
	CreateTime string  // 防火墙组创建时间，格式为Unix Timestamp
	Type       string  // 防火墙组类型，枚举值为：0：用户自定义防火墙；1：默认Web防火墙；2：默认非Web防火墙
	Rule       []*Rule // 防火墙组中的规则列表，结构如下：
}
type Rule struct {
	SrcIP        string // 源地址
	Priority     int    // 优先级
	ProtocolType string // 协议类型
	DstPort      string // 目标端口
	RuleAction   string // 防火墙动作
}

type DescribeSecurityGroupResponse struct {
	BaseResponse
	DataSet []*DescribeSecurityGroupDataSet `json:"omitempty"` // 获取的防火墙组详细信息

}

func (r *DescribeSecurityGroupResponse) Data() interface{} {
	return
}

type DescribeSecurityGroup struct {
	Region       string // 数据中心，参见 数据中心列表
	ResourceType string `ucloud:"optional"` // 绑定防火墙组的资源类型，如uhost
	ResourceId   int    `ucloud:"optional"` // 绑定防火墙组的资源ID
	GroupId      int    `ucloud:"optional"` // 防火墙ID

}

func (r *DescribeSecurityGroup) R() UResponse {
	return &DescribeSecurityGroupResponse{}
}

// ---------------- DescribeSecurityGroupResource ------------------

type DescribeSecurityGroupResourceResponse struct {
	BaseResponse
	DataSet []string `json:"omitempty"` // IP列表数组 如 [“10.10.10.10”, “10.10.10.11”]

}

func (r *DescribeSecurityGroupResourceResponse) Data() interface{} {
	return
}

type DescribeSecurityGroupResource struct {
	Region  string // 数据中心，参见 数据中心列表
	GroupId int    `ucloud:"optional"` // 防火墙ID

}

func (r *DescribeSecurityGroupResource) R() UResponse {
	return &DescribeSecurityGroupResourceResponse{}
}

// ---------------- CreateSecurityGroup ------------------

type CreateSecurityGroupResponse struct {
	BaseResponse
}

func (r *CreateSecurityGroupResponse) Data() interface{} {
	return
}

type CreateSecurityGroup struct {
	Region      string   // 数据中心，参见 数据中心列表
	GroupName   string   // 防火墙组名称
	Description string   `ucloud:"optional"` // 防火墙组描述
	Rule        []string // 防火墙规则，格式如下：
	Proto       string   // 网络协议，枚举值为：TCP，UDP，ICMP，GRE
	Dst_port    string   // 目标端口，范围：[1-65535]，可指定单个端口（如80），或指定端口段（1-1024）
	Src_ip      string   // 源地址，格式为’x.x.x.x/x 或 x.x.x.x’的有效IP地址。
	Action      string   // 防火墙动作，枚举值为：ACCEPT：允许通过防火墙；DROP：禁止通过防火墙并不给任何返回信息
	Priority    int      // 规则优先级，枚举值为：0（高），50（中），100（低）

}

func (r *CreateSecurityGroup) R() UResponse {
	return &CreateSecurityGroupResponse{}
}

// ---------------- UpdateSecurityGroup ------------------

type UpdateSecurityGroupResponse struct {
	BaseResponse
}

func (r *UpdateSecurityGroupResponse) Data() interface{} {
	return
}

type UpdateSecurityGroup struct {
	Region   string   // 数据中心，参见 数据中心列表
	GroupId  string   // 防火墙资源ID
	Rule     []string // 防火墙规则，格式如下
	Proto    string   // 网络协议，枚举值为：TCP，UDP，ICMP，GRE
	Dst_port string   // 目标端口，范围：[1-65535]，可指定单个端口（如80），或指定端口段（1-1024）
	Src_ip   string   // 源地址，格式为’x.x.x.x/x 或 x.x.x.x’的有效IP地址。
	Action   string   // 防火墙动作，枚举值为：ACCEPT：允许通过防火墙；DROP：禁止通过防火墙并不给任何返回信息
	Priority int      // 规则优先级，枚举值为：0（高），50（中），100（低）

}

func (r *UpdateSecurityGroup) R() UResponse {
	return &UpdateSecurityGroupResponse{}
}

// ---------------- GrantSecurityGroup ------------------

type GrantSecurityGroupResponse struct {
	BaseResponse
}

func (r *GrantSecurityGroupResponse) Data() interface{} {
	return r.RetCode
}

type GrantSecurityGroup struct {
	Region       string // 数据中心，参见 数据中心列表
	GroupId      string // 防火墙资源ID
	ResourceType string // 所应用资源类型，如UHost
	ResourceId   string // 所应用资源ID

}

func (r *GrantSecurityGroup) R() UResponse {
	return &GrantSecurityGroupResponse{}
}

// ---------------- DeleteSecurityGroup ------------------

type DeleteSecurityGroupResponse struct {
	BaseResponse
}

func (r *DeleteSecurityGroupResponse) Data() interface{} {
	return r.RetCode
}

type DeleteSecurityGroup struct {
	Region  string // 数据中心，参见 数据中心列表
	GroupId string // 防火墙资源ID

}

func (r *DeleteSecurityGroup) R() UResponse {
	return &DeleteSecurityGroupResponse{}
}
