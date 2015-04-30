package ucloud

// ---------------- CreateUdiskSnapshot ------------------

type CreateUdiskSnapshotResponse struct {
	BaseResponse
	SnapshotId string `json:"omitempty"` // 快照Id

}

func (r *CreateUdiskSnapshotResponse) Data() interface{} {
	return r.SnapshotId
}

type CreateUdiskSnapshot struct {
}

func (r *CreateUdiskSnapshot) R() UResponse {
	return &CreateUdiskSnapshotResponse{}
}

// ---------------- CreateUdisk ------------------

type CreateUdiskResponse struct {
	BaseResponse
	UDiskId string `json:"omitempty"` // UDisk实例Id

}

func (r *CreateUdiskResponse) Data() interface{} {
	return r.UDiskId
}

type CreateUdisk struct {
}

func (r *CreateUdisk) R() UResponse {
	return &CreateUdiskResponse{}
}

// ---------------- CloneUdisk ------------------

type CloneUdiskResponse struct {
	BaseResponse
	UDiskId string `json:"omitempty"` // 创建UDisk Id

}

func (r *CloneUdiskResponse) Data() interface{} {
	return r.UDiskId
}

type CloneUdisk struct {
}

func (r *CloneUdisk) R() UResponse {
	return &CloneUdiskResponse{}
}

// ---------------- CloneUdiskSnapshot ------------------

type CloneUdiskSnapshotResponse struct {
	BaseResponse
	UDiskId string `json:"omitempty"` // 创建UDisk Id

}

func (r *CloneUdiskSnapshotResponse) Data() interface{} {
	return r.UDiskId
}

type CloneUdiskSnapshot struct {
	ProjectId string // 项目编号 None  // 购买时长 默认: 1

}

func (r *CloneUdiskSnapshot) R() UResponse {
	return &CloneUdiskSnapshotResponse{}
}

// ---------------- DeleteUdiskSnapshot ------------------

type DeleteUdiskSnapshotResponse struct {
	BaseResponse
}

func (r *DeleteUdiskSnapshotResponse) Data() interface{} {
	return r.RetCode
}

type DeleteUdiskSnapshot struct {
}

func (r *DeleteUdiskSnapshot) R() UResponse {
	return &DeleteUdiskSnapshotResponse{}
}

// ---------------- DeleteUdisk ------------------

type DeleteUdiskResponse struct {
	BaseResponse
}

func (r *DeleteUdiskResponse) Data() interface{} {
	return r.RetCode
}

type DeleteUdisk struct {
}

func (r *DeleteUdisk) R() UResponse {
	return &DeleteUdiskResponse{}
}

// ---------------- DescribeUdiskSnapshot ------------------
type DescribeUdiskSnapshotItem struct {
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

type DescribeUdiskSnapshotResponse struct {
	BaseResponse
	DataSet    []*DescribeUdiskSnapshotItem `json:"omitempty"` // JSON 格式的Snapshot列表, 详细参见ResponseItem
	TotalCount int                          `json:"omitempty"` // 根据过滤条件得到的总数

}

func (r *DescribeUdiskSnapshotResponse) Data() interface{} {
	return r.DataSet
}

type DescribeUdiskSnapshot struct {
}

func (r *DescribeUdiskSnapshot) R() UResponse {
	return &DescribeUdiskSnapshotResponse{}
}

// ---------------- DescribeUdisk ------------------
type DescribeUdiskItem struct {
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

type DescribeUdiskResponse struct {
	BaseResponse
	DataSet    []*DescribeUdiskItem `json:"omitempty"` // JSON 格式的UDisk数据列表, 每项参数可见下面 ResponseItem
	TotalCount int                  `json:"omitempty"` // 根据过滤条件得到的总数

}

func (r *DescribeUdiskResponse) Data() interface{} {
	return r.DataSet
}

type DescribeUdisk struct {
}

func (r *DescribeUdisk) R() UResponse {
	return &DescribeUdiskResponse{}
}

// ---------------- DescribeUdiskPrice ------------------
type DescribeUdiskPriceItem struct {
	ChargeType string  // Year， Month， Dynamic，Trial
	Price      float64 // 价格
}

type DescribeUdiskPriceResponse struct {
	BaseResponse
	DataSet []*DescribeUdiskPriceItem `json:"omitempty"` // 价格

}

func (r *DescribeUdiskPriceResponse) Data() interface{} {
	return r.DataSet
}

type DescribeUdiskPrice struct {
}

func (r *DescribeUdiskPrice) R() UResponse {
	return &DescribeUdiskPriceResponse{}
}

// ---------------- DescribeUdiskUpgradePrice ------------------

type DescribeUdiskUpgradePriceResponse struct {
	BaseResponse
	Price float64 `json:"omitempty"` // 价格

}

func (r *DescribeUdiskUpgradePriceResponse) Data() interface{} {
	return r.Price
}

type DescribeUdiskUpgradePrice struct {
}

func (r *DescribeUdiskUpgradePrice) R() UResponse {
	return &DescribeUdiskUpgradePriceResponse{}
}

// ---------------- RenameUdisk ------------------

type RenameUdiskResponse struct {
	BaseResponse
}

func (r *RenameUdiskResponse) Data() interface{} {
	return r.RetCode
}

type RenameUdisk struct {
}

func (r *RenameUdisk) R() UResponse {
	return &RenameUdiskResponse{}
}

// ---------------- ResizeUdisk ------------------

type ResizeUdiskResponse struct {
	BaseResponse
}

func (r *ResizeUdiskResponse) Data() interface{} {
	return r.RetCode
}

type ResizeUdisk struct {
}

func (r *ResizeUdisk) R() UResponse {
	return &ResizeUdiskResponse{}
}
