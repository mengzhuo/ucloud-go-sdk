package ucloud

// CreateUHostInstance ----------------------
type CreateUHostInstanceResponse struct {
	BaseResponse
	UHostIds []string `json:",omitempty"`
}

func (r *CreateUHostInstanceResponse) Data() interface{} {
	return r.UHostIds
}

type CreateUHostInstance struct {
	Region          string
	ImageId         string
	LoginMode       string
	Password        string `ucloud:"optional"`
	KeyPair         string `ucloud:"optional"`
	CPU             int    `ucloud:"optional"`
	Memory          int    `ucloud:"optional"`
	DiskSpace       int    `ucloud:"optional"`
	Name            string `ucloud:"optional"`
	NetworkId       string `ucloud:"optional"`
	SecurityGroupId string `ucloud:"optional"`
	ChargeType      string `ucloud:"optional"`
	Quantity        int    `ucloud:"optional"`
}

func (r *CreateUHostInstance) R() UResponse {

	return &CreateUHostInstanceResponse{}
}

// DescribeUHostInstance ----------------------

type DiskSet struct {
	Type   string
	DiskId string
	Size   int
}
type IPSet struct {
	Type string
	IP   string
}
type UHostSet struct {
	UHostId        string     //UHost实例ID
	UHostType      string     //UHost类型，枚举为：Normal：普通机型；SSD：SSD高性能机型；BD：大数据机型
	ImageId        string     //镜像ID
	BasicImageId   string     `json:",omitempty"` //基础镜像ID
	BasicImageName string     `json:",omitempty"` //基础镜像名称
	Tag            string     `json:",omitempty"` //业务组名称
	Remark         string     `json:",omitempty"` //备注
	Name           string     //UHost实例名称
	State          string     //实例状态， 初始化: Initializing; 启动中: Starting; 运行中: Running; 关机中: Stopping; 关机: Stopped 安装失败: Install Fail; 重启中: Rebooting
	CreateTime     int        //创建时间
	ChargeType     string     //计费模式，枚举值为： Year，按年付费； Month，按月付费； Dynamic，按需付费（需开启权限）； Trial，试用（需开启权限） 默认为月付
	ExpireTime     int        //到期时间
	CPU            int        //虚拟CPU核数，单位: 个
	Memory         int        //内存大小，单位: MB
	DiskSet        []*DiskSet `json:",omitempty"`
	IPSet          []*IPSet   `json:",omitempty"`
}

type DescribeUHostInstanceResponse struct {
	BaseResponse
	TotalCount int         `json:",omitempty"`
	UHostSet   []*UHostSet `json:",omitempty"`
}

func (r *DescribeUHostInstanceResponse) Data() interface{} {
	return r.UHostSet
}

func (u *DescribeUHostInstance) R() UResponse {
	return &DescribeUHostInstanceResponse{}
}

type DescribeUHostInstance struct {
	Region   string
	UHostIds []string `ucloud:"optional"`
	Offset   int      `ucloud:"optional"`
	Limit    int      `ucloud:"optional"`
}

// TerminateUHostInstance -----------------

func (r *TerminateUHostInstanceResponse) Data() interface{} {
	return r.UHostIds
}

type TerminateUHostInstanceResponse struct {
	BaseResponse
	UHostIds []string `json:",omitempty"`
}

func (r *TerminateUHostInstance) R() UResponse {
	return &TerminateUHostInstanceResponse{}
}

type TerminateUHostInstance struct {
	Region  string //数据中心，
	UHostId string // UHost资源Id
}

// ResizeUHostInstance ---------------

type ResizeUHostInstanceResponse struct {
	BaseResponse
	UHostId string `json:",omitempty"`
}

func (r *ResizeUHostInstanceResponse) Data() interface{} {
	return r.UHostId
}

type ResizeUHostInstance struct {
	Region    string // 数据中心，参见 数据中心列表	Yes
	UHostId   string // UHost实例ID	Yes
	CPU       int    `ucloud:"optional"` // 虚拟CPU核数，单位：个，范围：[1,16]，最小值为1，其他值是2的倍数，默认值为当前实例的CPU核数	no
	Memory    int    `ucloud:"optional"` //	内存大小，单位：MB，范围[2048,65536]，步长：2048，默认值为当前实例的内存大小	no
	DiskSpace int    `ucloud:"optional"` // 数据盘大小，单位：GB，范围[10,1000]，步长：10，默认值为当前实例的数据盘大小，数据盘不支持缩容，因此不允许输入比当前实例数据盘大小的值
}

func (r *ResizeUHostInstance) R() UResponse {
	return &ResizeUHostInstanceResponse{}
}

// ReinstallUHostInstance ------------------
type ReinstallUHostInstanceResponse struct {
	BaseResponse
	UHostId string
}

func (r *ReinstallUHostInstanceResponse) Data() interface{} {
	return r.UHostId
}

type ReinstallUHostInstance struct {
	Region      string //	数据中心，参见 数据中心列表	Yes
	UHostId     string //	UHost实例资源ID	Yes
	Password    string `ucloud:"optional"` //	如果创建UHost实例时LoginMode为Password，则必须填写，如果LoginMode为KeyPair，不需要填写 （密码格式使用BASE64编码）	no
	ImageId     string `ucloud:"optional"` //	镜像Id，参考镜像列表，默认使用原镜像	no
	ReserveDisk string `ucloud:"optional"` //	是否保留数据盘，保留：Yes，不报留：No， 默认：Yes	no

}

func (r *ReinstallUHostInstance) R() UResponse {
	return &ReinstallUHostInstanceResponse{}
}

// StartUHostInstance ------------------
type StartUHostInstanceResponse struct {
	BaseResponse
	UHostId string `json:",omitempty"`
}

func (r *StartUHostInstanceResponse) Data() interface{} {
	return r.UHostId
}

type StartUHostInstance struct {
	Region  string //数据中心，参见 数据中心列表	Yes
	UHostId string //UHost实例ID	Yes
}

func (r *StartUHostInstance) R() UResponse {
	return &StartUHostInstanceResponse{}
}

// StopUHostInstanceResponse ----------
type StopUHostInstanceResponse struct {
	BaseResponse
	UHostId string `json:",omitempty"`
}

func (r *StopUHostInstanceResponse) Data() interface{} {
	return r.UHostId
}

type StopUHostInstance struct {
	Region  string //数据中心，参见 数据中心列表	Yes
	UHostId string //UHost实例ID	Yes
}

func (r *StopUHostInstance) R() UResponse {
	return &StartUHostInstanceResponse{}
}

// RebootUHostInstance -------------

type RebootUHostInstanceResponse struct {
	BaseResponse
	UHostId string `json:",omitempty"`
}

func (r *RebootUHostInstanceResponse) Data() interface{} {
	return r.UHostId
}

type RebootUHostInstance struct {
	Region  string //数据中心，参见 数据中心列表	Yes
	UHostId string //UHost实例ID	Yes
}

func (r *RebootUHostInstance) R() UResponse {
	return &RebootUHostInstanceResponse{}
}

// ResetUHostInstancePassword -----------

type ResetUHostInstancePasswordResponse struct {
	BaseResponse
	UHostId string `json:",omitempty"`
}

func (r *ResetUHostInstancePasswordResponse) Data() interface{} {
	return r.UHostId
}

type ResetUHostInstancePassword struct {
	Region   string //数据中心，参见 数据中心列表	Yes
	UHostId  string //UHost实例ID	Yes
	Password string
}

func (r *ResetUHostInstancePassword) R() UResponse {
	return &ResetUHostInstancePasswordResponse{}
}

// ModifyUHostInstanceName ------------
type ModifyUHostInstanceNameResponse struct {
	BaseResponse
	UHostId string `json:",omitempty"`
}

func (r *ModifyUHostInstanceNameResponse) Data() interface{} {
	return r.UHostId
}

type ModifyUHostInstanceName struct {
	Region  string //数据中心，参见 数据中心列表	Yes
	UHostId string //UHost实例ID	Yes
	Name    string //UHost实例名称	Yes
}

func (r *ModifyUHostInstanceName) R() UResponse {
	return &ModifyUHostInstanceNameResponse{}
}

// ModifyUHostInstanceTag ---------
type ModifyUHostInstanceTagResponse struct {
	BaseResponse
	UHostId string `json:",omitempty"`
}

func (r *ModifyUHostInstanceTagResponse) Data() interface{} {
	return r.UHostId
}

type ModifyUHostInstanceTag struct {
	Region  string //数据中心，参见 数据中心列表	Yes
	UHostId string //UHost实例ID	Yes
	Tag     string //UHost实例名称	Yes
}

func (r *ModifyUHostInstanceTag) R() UResponse {
	return &ModifyUHostInstanceTagResponse{}
}

// ModifyUHostInstanceRemark ----------
type ModifyUHostInstanceRemarkResponse struct {
	BaseResponse
	UHostId string `json:",omitempty"`
}

func (r *ModifyUHostInstanceRemarkResponse) Data() interface{} {
	return r.UHostId
}

type ModifyUHostInstanceRemark struct {
	Region  string //数据中心，参见 数据中心列表	Yes
	UHostId string //UHost实例ID	Yes
	Remark  string
}

func (r *ModifyUHostInstanceRemark) R() UResponse {
	return &ModifyUHostInstanceTagResponse{}
}

// GetUHostInstancePrice -------------
type PriceSet struct {
	ChargeType string
	Price      float64
}

type GetUHostInstancePriceResponse struct {
	BaseResponse
	PriceSet []*PriceSet `json:",omitempty"`
}

func (r *GetUHostInstancePriceResponse) Data() interface{} {
	return r.PriceSet
}

type GetUHostInstancePrice struct {
	Region     string //	数据中心，参见 数据中心列表	Yes
	ImageId    string //	镜像Id，参见 镜像列表	Yes
	CPU        int    //	虚拟CPU核心数，单位: 个，范围: [1,16]，最小值为1，其他值是2的整数倍	Yes
	Memory     int    //	内存容量大小，单位: MB，范围: [2048,65536]，步长: 2048	Yes
	Count      int    //购买台数，范围[1,5]	Yes
	ChargeType string `ucloud:"optional"` //Year，Month，Dynamic，默认返回全部计费方式对应的价格	No
	DiskSpace  int    `ucloud:"optional"` //数据盘大小，单位: GB，范围[0,1000]，步长: 10，默认值: 0	No
}

func (r *GetUHostInstancePrice) R() UResponse {
	return &GetUHostInstancePriceResponse{}
}

// GetUHostInstanceVncInfo -----------------

type GetUHostInstanceVncInfoResponse struct {
	BaseResponse
	UHostId     string `json:",omitempty"` //UHost实例ID
	VncIP       string `json:",omitempty"` //Vnc登录IP
	VncPort     int    `json:",omitempty"` //Vnc登录端口
	VncPassword string `json:",omitempty"` //Vnc 登录密码
}

func (r *GetUHostInstanceVncInfoResponse) Data() interface{} {
	return struct {
		UHostId     string
		VncIP       string
		VncPort     int
		VncPassword string
	}{
		r.UHostId,
		r.VncIP,
		r.VncPort,
		r.VncPassword,
	}
}

type GetUHostInstanceVncInfo struct {
	Region  string //数据中心，参见 数据中心列表	Yes
	UHostId string //UHost实例ID	Yes
}

func (r *GetUHostInstanceVncInfo) R() UResponse {
	return &GetUHostInstanceVncInfoResponse{}
}

// DescribeImage  ----------------
type ImageSet struct {
	ImageId          string //	镜像ID
	ImageName        string //	镜像名称
	OsType           string //	操作系统类型：Liunx，Windows
	OsName           string //	操作系统名称
	ImageType        string //	标准镜像：Base，行业镜像：Business， 自定义镜像：Custom
	State            string //	镜像状态， 可用：Available，制作中：Making， 不可用：Unavailable
	ImageDescription string //	镜像描述
	CreateTime       int    //	创建时间
}

type DescribeImageResponse struct {
	BaseResponse
	TotalCount int
	ImageSet   []*ImageSet `json:",omitempty"`
}

func (r *DescribeImageResponse) Data() interface{} {
	return r.ImageSet
}

type DescribeImage struct {
	Region    string //	数据中心，参见 数据中心列表	Yes
	ImageType string `ucloud:"optional"` //	标准镜像：Base，行业镜像：Business， 自定义镜像：Custom，默认返回所有类型	No
	OsType    string `ucloud:"optional"` //	操作系统类型：Linux， Windows 默认返回所有类型	No
	ImageId   string `ucloud:"optional"` //	镜像Id	No
	Offset    int    `ucloud:"optional"` //	数据偏移量，默认为0	No
	Limit     int    `ucloud:"optional"` //	返回数据长度，默认为20	No
}

func (r *DescribeImage) R() UResponse {
	return &DescribeImageResponse{}
}

// CreateCustomImage-------------
type CreateCustomImageResponse struct {
	BaseResponse
	ImageId string
}

func (r *CreateCustomImageResponse) Data() interface{} {
	return r.ImageId
}

type CreateCustomImage struct {
	Region           string //	数据中心，参见 数据中心列表	Yes
	UHostId          string //	UHost实例ID	Yes
	ImageName        string //	镜像名称	Yes
	ImageDescription string `ucloud:"optional"` //	镜像描述	No
}

func (r *CreateCustomImage) R() UResponse {
	return &CreateCustomImageResponse{}
}

// TerminateCustomImage--------------
type TerminateCustomImageResponse struct {
	BaseResponse
	ImageId string `json:",omitempty"`
}

func (r *TerminateCustomImageResponse) Data() interface{} {
	return r.ImageId
}

type TerminateCustomImage struct {
	Region  string //	数据中心，参见 数据中心列表	Yes
	ImageId string //  自制镜像ID
}

func (r *TerminateCustomImage) R() UResponse {
	return &TerminateCustomImageResponse{}
}

// ---------------- CreateUHostInstanceSnapshot ------------------

type CreateUHostInstanceSnapshotResponse struct {
	BaseResponse
	UHostId      string `json:",omitempty"` // UHost实例ID
	SnapshotName string `json:",omitempty"` // 快照名称

}

func (r *CreateUHostInstanceSnapshotResponse) Data() interface{} {
	return struct {
		UHostId      string `json:",omitempty"` // UHost实例ID
		SnapshotName string `json:",omitempty"` // 快照名称
	}{r.UHostId, r.SnapshotName}
}

type CreateUHostInstanceSnapshot struct {
	Region  string // 数据中心，参见 数据中心列表
	UHostId string // UHost实例ID
}

func (r *CreateUHostInstanceSnapshot) R() UResponse {
	return &CreateUHostInstanceSnapshotResponse{}
}

// ---------------- DescribeUHostInstanceSnapshot ------------------
type SnapshotSet struct {
	SnapshotName string // 快照名称
	SnapshotTime string // 快照制作时间
}

type DescribeUHostInstanceSnapshotResponse struct {
	BaseResponse
	UHostId     string         `json:",omitempty"` // UHost实例ID
	SnapshotSet []*SnapshotSet `json:",omitempty"` // UHost快照列表，详细参数可见下面SnapshotSet

}

func (r *DescribeUHostInstanceSnapshotResponse) Data() interface{} {
	return struct {
		UHostId     string         `json:",omitempty"` // UHost实例ID
		SnapshotSet []*SnapshotSet `json:",omitempty"` // UHost快照列表，详细参数可见下面SnapshotSet
	}{r.UHostId, r.SnapshotSet}
}

type DescribeUHostInstanceSnapshot struct {
	Region  string // 数据中心，参见 数据中心列表
	UHostId string // UHost实例ID
}

func (r *DescribeUHostInstanceSnapshot) R() UResponse {
	return &DescribeUHostInstanceSnapshotResponse{}
}

// ---------------- AttachUDisk ------------------

type AttachUDiskResponse struct {
	BaseResponse
	UHostId string `json:"omitempty"` // 挂载的UHost实例ID
	UDiskId string `json:"omitempty"` // 挂载的UDisk实例ID

}

func (r *AttachUDiskResponse) Data() interface{} {
	return struct {
		UHostId string
		UDiskId string
	}{r.UHostId, r.UHostId}
}

type AttachUDisk struct {
	Region  string // 数据中心，参见 数据中心列表
	UHostId string // UHost实例ID
	UDiskId string //需要挂载的UDisk实例ID
}

func (r *AttachUDisk) R() UResponse {
	return &AttachUDiskResponse{}
}

// ---------------- DetachUDisk ------------------

type DetachUDiskResponse struct {
	BaseResponse
	UHostId string `json:"omitempty"` // 卸载的UHost实例ID
	UDiskId string `json:"omitempty"` // 卸载的UDisk实例ID

}

func (r *DetachUDiskResponse) Data() interface{} {
	return struct {
		UHostId string
		UDiskId string
	}{
		r.UHostId,
		r.UDiskId}
}

type DetachUDisk struct {
	Region  string // 数据中心，参见 数据中心列表
	UHostId string // UHost实例ID
	UDiskId string //需要挂载的UDisk实例ID
}

func (r *DetachUDisk) R() UResponse {
	return &DetachUDiskResponse{}
}
