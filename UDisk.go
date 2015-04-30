package ucloud

// ---------------- CreateUDiskSnapshot ------------------

type CreateUDiskSnapshotResponse struct {
	BaseResponse
	SnapshotId string `json:"omitempty"` // 快照Id

}

func (r *CreateUDiskSnapshotResponse) Data() interface{} {
	return r.SnapshotId
}

type CreateUDiskSnapshot struct {
	Region     string //	数据中心, 参见 数据中心列表	Yes
	UDiskId    string //	快照的UDisk的Id	Yes
	Name       string //	快照名称	Yes
	ChargeType string `ucloud:"optional"` //Year , Month, Dynamic 默认: Dynamic	No
	Quantity   int    `ucloud:"optional"` //	购买时长 默认: 1	No
	ProjectId  string `ucloud:"optional"` //	项目编号	No
	Comment    string `ucloud:"optional"` //快照描述	No

}

func (r *CreateUDiskSnapshot) R() UResponse {
	return &CreateUDiskSnapshotResponse{}
}

// ---------------- CreateUDisk ------------------

type CreateUDiskResponse struct {
	BaseResponse
	UDiskId string `json:"omitempty"` // UDisk实例Id

}

func (r *CreateUDiskResponse) Data() interface{} {
	return r.UDiskId
}

type CreateUDisk struct {
	Region     string //	数数据中心, 参见 数据中心列表	Yes
	Size       int    // 磁盘大小, 单位:GB, 范围[1~1000]	Yes
	Name       string //	实例名称	Yes
	CouponId   string `ucloud:"optional"` //	使用的代金券id	No
	ChargeType string `ucloud:"optional"` //	Year , Month, Dynamic, Trial 默认: Dynamic	No
	Quantity   int    `ucloud:"optional"` //  购买时长 默认: 1	No
	ProjectId  string `ucloud:"optional"` //	项目编号	No

}

func (r *CreateUDisk) R() UResponse {
	return &CreateUDiskResponse{}
}

// ---------------- CloneUDisk ------------------

type CloneUDiskResponse struct {
	BaseResponse
	UDiskId string `json:"omitempty"` // 创建UDisk Id

}

func (r *CloneUDiskResponse) Data() interface{} {
	return r.UDiskId
}

type CloneUDisk struct {
	Region     string // 	数据中心, 参见 数据中心列表	Yes
	Name       string `ucloud:"optional"` //  实例名称	No
	SourceId   string //	克隆父Disk的Id	Yes
	CouponId   string `ucloud:"optional"` //	使用的代金券id	No
	Comment    string `ucloud:"optional"` //  Disk注释	No
	ChargeType string `ucloud:"optional"` //	Year , Month, Dynamic 默认: Dynamic	No
	Quantity   int    `ucloud:"optional"` //r	购买时长 默认: 1	No
	ProjectId  string `ucloud:"optional"` //	项目编号	No
}

func (r *CloneUDisk) R() UResponse {
	return &CloneUDiskResponse{}
}

// ---------------- CloneUDiskSnapshot ------------------

type CloneUDiskSnapshotResponse struct {
	BaseResponse
	UDiskId string `json:"omitempty"` // 创建UDisk Id

}

func (r *CloneUDiskSnapshotResponse) Data() interface{} {
	return r.UDiskId
}

type CloneUDiskSnapshot struct {
	Region     string //	数据中心, 参见 数据中心列表	Yes
	Name       string `ucloud:"optional"` //实例名称	No
	SourceId   string //克隆父Snapshot的Id	Yes
	Size       int    //磁盘大小, 单位:GB, 范围[1~1000]	Yes
	CouponId   string `ucloud:"optional"` //使用的代金券id	No
	Comment    string `ucloud:"optional"` //Disk注释	No
	ChargeType string `ucloud:"optional"` //Year , Month, Dynamic 默认: Dynamic	No
	Quantity   int    `ucloud:"optional"` //购买时长 默认: 1	No
	ProjectId  string `ucloud:"optional"` // 项目编号 None  // 购买时长 默认: 1
}

func (r *CloneUDiskSnapshot) R() UResponse {
	return &CloneUDiskSnapshotResponse{}
}

// ---------------- DeleteUDiskSnapshot ------------------

type DeleteUDiskSnapshotResponse struct {
	BaseResponse
}

func (r *DeleteUDiskSnapshotResponse) Data() interface{} {
	return r.RetCode
}

type DeleteUDiskSnapshot struct {
	Region     string //	数据中心, 参见 数据中心列表	Yes
	SnapshotId string //	快照Id	Yes
	ProjectId  string `ucloud:"optional"` //项目编号	No

}

func (r *DeleteUDiskSnapshot) R() UResponse {
	return &DeleteUDiskSnapshotResponse{}
}

// ---------------- DeleteUDisk ------------------

type DeleteUDiskResponse struct {
	BaseResponse
}

func (r *DeleteUDiskResponse) Data() interface{} {
	return r.RetCode
}

type DeleteUDisk struct {
	Region    string //	数据中心, 参见 数据中心列表	Yes
	UDiskId   string //	要删除的UDisk的Id	Yes
	ProjectId string `ucloud:"optional"` //	项目编号	No
}

func (r *DeleteUDisk) R() UResponse {
	return &DeleteUDiskResponse{}
}

// ---------------- DescribeUDiskSnapshot ------------------
type DescribeUDiskSnapshotItem struct {
	SnapshotId  string // 快照Id
	Name        string // 快照名称
	UDiskId     string // 快照的源UDisk的Id
	UDiskName   string // 快照的源UDisk的Name
	CreateTime  int    // 创建时间
	ExpiredTime int    // 过期时间
	ChargeType  string // Year,Month,Dynamic,Trial
	Size        int    // 容量单位GB
	Comment     string // 快照描述
}

type DescribeUDiskSnapshotResponse struct {
	BaseResponse
	DataSet    []*DescribeUDiskSnapshotItem `json:"omitempty"` // JSON 格式的Snapshot列表, 详细参见ResponseItem
	TotalCount int                          `json:"omitempty"` // 根据过滤条件得到的总数

}

func (r *DescribeUDiskSnapshotResponse) Data() interface{} {
	return r.DataSet
}

type DescribeUDiskSnapshot struct {
	Region     string //数据中心, 参见 数据中心列表	Yes
	SnapshotId string `ucloud:"optional"` //UDisk快照Id(留空返回全部)	No
	Offset     int    `ucloud:"optional"` //数据偏移量, 默认为0	No
	Limit      int    `ucloud:"optional"` //返回数据长度, 默认为20	No
	ProjectId  string `ucloud:"optional"` //项目编号	No
}

func (r *DescribeUDiskSnapshot) R() UResponse {
	return &DescribeUDiskSnapshotResponse{}
}

// ---------------- DescribeUDisk ------------------
type DescribeUDiskItem struct {
	UDiskId     string // UDisk实例Id
	Name        string // 实例名称
	Size        string // 容量单位GB
	Status      string // 状态：Available(可用),Attaching(挂载中),InUse(已挂载),Detaching(卸载中),
	CreateTime  int    // 创建时间
	ExpiredTime int    // 过期时间
	UHostId     string // 挂载的UHost的Id
	UHostName   string // 挂载的UHost的Name
	UHostIP     string // 挂载的UHost的IP
	DeviceName  string // 挂载的设备名称
	ChargeType  string // Year,Month,Dynamic,Trial
}

type DescribeUDiskResponse struct {
	BaseResponse
	DataSet    []*DescribeUDiskItem `json:"omitempty"` // JSON 格式的UDisk数据列表, 每项参数可见下面 ResponseItem
	TotalCount int                  `json:"omitempty"` // 根据过滤条件得到的总数

}

func (r *DescribeUDiskResponse) Data() interface{} {
	return r.DataSet
}

type DescribeUDisk struct {
	Region    string   //	数据中心, 参见 数据中心列表	Yes
	UDiskId   []string `ucloud:"optional"` //	UDisk Id(留空返回全部)	No
	Offset    int      `ucloud:"optional"` //	数据偏移量, 默认为0	No
	Limit     int      `ucloud:"optional"` //    返回数据长度, 默认为20	No
	ProjectId string   `ucloud:"optional"` //	项目编号	No
}

func (r *DescribeUDisk) R() UResponse {
	return &DescribeUDiskResponse{}
}

// ---------------- DescribeUDiskPrice ------------------
type DescribeUDiskPriceItem struct {
	ChargeType string  // Year， Month， Dynamic，Trial
	Price      float64 // 价格
}

type DescribeUDiskPriceResponse struct {
	BaseResponse
	DataSet []*DescribeUDiskPriceItem `json:"omitempty"` // 价格

}

func (r *DescribeUDiskPriceResponse) Data() interface{} {
	return r.DataSet
}

type DescribeUDiskPrice struct {
	Region     string //	数据中心, 参见 数据中心列表	Yes
	Size       int    //	购买UDisk大小,单位:GB,范围[1~1000]	Yes
	ChargeType string `ucloud:"optional"` //	Year， Month， Dynamic，Trial，默认: Dynamic 如果不指定，则一次性获取三种计费	No
	Quantity   int    `ucloud:"optional"` //	购买UDisk的时长，默认值为1	No
	ProjectId  string `ucloud:"optional"` //	项目编号	No

}

func (r *DescribeUDiskPrice) R() UResponse {
	return &DescribeUDiskPriceResponse{}
}

// ---------------- DescribeUDiskUpgradePrice ------------------

type DescribeUDiskUpgradePriceResponse struct {
	BaseResponse
	Price float64 `json:"omitempty"` // 价格

}

func (r *DescribeUDiskUpgradePriceResponse) Data() interface{} {
	return r.Price
}

type DescribeUDiskUpgradePrice struct {
	Region    string //	数据中心, 参见 数据中心列表	Yes
	Size      int    //	购买UDisk大小,单位:GB,范围[1~1000]	Yes
	SourceId  string //	升级目标UDisk ID	Yes
	ProjectId string `ucloud:"optional"` //	项目编号	No
}

func (r *DescribeUDiskUpgradePrice) R() UResponse {
	return &DescribeUDiskUpgradePriceResponse{}
}

// ---------------- RenameUDisk ------------------

type RenameUDiskResponse struct {
	BaseResponse
}

func (r *RenameUDiskResponse) Data() interface{} {
	return r.RetCode
}

type RenameUDisk struct {
	Region    string //	数据中心, 参见 数据中心列表	Yes
	UDiskId   string //	重命名的UDisk的Id	Yes
	UDiskName string //	重命名UDisk的name	Yes
	ProjectId string `ucloud:"optional"` //	项目编号	No

}

func (r *RenameUDisk) R() UResponse {
	return &RenameUDiskResponse{}
}

// ---------------- ResizeUDisk ------------------

type ResizeUDiskResponse struct {
	BaseResponse
}

func (r *ResizeUDiskResponse) Data() interface{} {
	return r.RetCode
}

type ResizeUDisk struct {
	Region    string //	数据中心, 参见 数据中心列表	Yes
	UDiskId   string //	UDisk Id	Yes
	Size      int    //	调整后大小, 单位:GB, 范围[1~1000],只能增大	Yes
	CouponId  string `ucloud:"optional"` //	使用的代金券id	No
	ProjectId string `ucloud:"optional"` //	项目编号	No

}

func (r *ResizeUDisk) R() UResponse {
	return &ResizeUDiskResponse{}
}
