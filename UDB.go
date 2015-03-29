package ucloud

// ---------------- BackupUdbInstance ------------------

type BackupUdbInstanceResponse struct {
	BaseResponse
}

func (r *BackupUdbInstanceResponse) Data() interface{} {
	return r.RetCode
}

type BackupUdbInstance struct {
	Region       string // 数据中心，请参见数据中心RegionList
	DBId         string // DB实例Id
	BackupName   string // 备份名称
	UseBlacklist bool   `ucloud:"optional"` // 是否使用黑名单备份，默认false

}

func (r *BackupUdbInstance) R() UResponse {
	return &BackupUdbInstanceResponse{}
}

// ---------------- ClearUdbLog ------------------

type ClearUdbLogResponse struct {
	BaseResponse
}

func (r *ClearUdbLogResponse) Data() interface{} {
	return r.RetCode
}

type ClearUdbLog struct {
	Region     string // 数据中心，请参见数据中心RegionList
	DBId       string // DB实例的id
	LogType    int    // 日志类型，10-error（暂不支持）、20-slow（暂不支持）、30-binlog
	BeforeTime int    `ucloud:"optional"` // 删除时间点(至少前一天)之前log，采用时间戳(秒)，默认当前时间点前一天

}

func (r *ClearUdbLog) R() UResponse {
	return &ClearUdbLogResponse{}
}

// ---------------- CreateUdbInstance ------------------

type CreateUdbInstanceResponse struct {
	BaseResponse
	DBId string `json:",omitempty"` // db实例id

}

func (r *CreateUdbInstanceResponse) Data() interface{} {
	return r.DBId
}

type CreateUdbInstance struct {
	Region         string // 数据中心，请参见数据中心RegionList
	DBTypeId       string // DB类型id，mysql/mongodb按版本细分各有一个id1：mysql-5.5，2：mysql-5.1，3：percona-5.54：mongodb-2.4，5：mongodb-2.6，6：mysql-5.6，7：percona-5.6
	ChargeType     string `ucloud:"optional"` // Year， Month， Dynamic，Trial，默认: Dynamic
	CouponId       string `ucloud:"optional"` // 使用的代金券id
	Quantity       int    `ucloud:"optional"` // 购买时长，默认值1
	Name           string // 实例名称，至少6位
	AdminUser      string `ucloud:"optional"` // 管理员帐户名，默认root
	AdminPassword  string // 管理员密码
	Port           int    `ucloud:"optional"` // 端口号，mysql默认3306，mongodb默认27017
	ParamGroupId   int    // DB实例使用的配置参数组id
	MemoryLimit    int    // 内存限制(MB)，目前支持以下几档600M/1500M/3000M/6000M/15000M/30000M
	DiskSpace      int    // 磁盘空间(GB), 暂时支持20G - 500G
	BackupCount    int    `ucloud:"optional"` // 备份策略，每周备份数量，默认7次
	BackupTime     int    `ucloud:"optional"` // 备份策略，备份开始时间，单位小时计，默认3点
	BackupDuration int    `ucloud:"optional"` // 备份策略，备份时间间隔，单位小时计，默认24小时
	BackupId       int    `ucloud:"optional"` // 备份id，如果指定，则表明从备份恢复实例
	UseSSD         bool   `ucloud:"optional"` // 是否使用SSD，默认为false

}

func (r *CreateUdbInstance) R() UResponse {
	return &CreateUdbInstanceResponse{}
}

// ---------------- CreateUdbParamGroup ------------------

type CreateUdbParamGroupResponse struct {
	BaseResponse
	GroupId int `json:",omitempty"` // 新配置参数组id

}

func (r *CreateUdbParamGroupResponse) Data() interface{} {
	return r.GroupId
}

type CreateUdbParamGroup struct {
	Region      string // 数据中心，请参见数据中心RegionList
	GroupName   string // 新配置参数组名称
	Description string // 参数组描述
	SrcGroupId  int    // 源参数组id
	DBTypeId    string // DB类型id，mysql/mongodb按版本细分各有一个id1：mysql-5.5，2：mysql-5.1，3：percona-5.54：mongodb-2.4，5：mongodb-2.6，6：mysql-5.6，7：percona-5.6

}

func (r *CreateUdbParamGroup) R() UResponse {
	return &CreateUdbParamGroupResponse{}
}

// ---------------- CreateUdbReplicationInstance ------------------

type CreateUdbReplicationInstanceResponse struct {
	BaseResponse
	DBId string `json:",omitempty"` // 创建从节点的DBId

}

func (r *CreateUdbReplicationInstanceResponse) Data() interface{} {
	return r.DBId
}

type CreateUdbReplicationInstance struct {
	Region    string // 数据中心，请参见数据中心RegionList
	SrcId     string // primary节点的DBId
	Name      string // 实例名称，至少6位
	Port      int    `ucloud:"optional"` // 端口号，默认27017
	IsArbiter bool   `ucloud:"optional"` // 是否是仲裁节点，默认false，仲裁节点按最小机型创建
	CouponId  string `ucloud:"optional"` // 使用的代金券id

}

func (r *CreateUdbReplicationInstance) R() UResponse {
	return &CreateUdbReplicationInstanceResponse{}
}

// ---------------- CreateUdbSlave ------------------

type CreateUdbSlaveResponse struct {
	BaseResponse
	DBId string `json:",omitempty"` // 创建slave的DBId

}

func (r *CreateUdbSlaveResponse) Data() interface{} {
	return r.DBId
}

type CreateUdbSlave struct {
	Region   string // 数据中心，请参见数据中心RegionList
	SrcId    string // master实例的DBId
	Name     string // 实例名称，至少6位
	Port     int    `ucloud:"optional"` // 端口号，mysql默认3306
	UseSSD   bool   `ucloud:"optional"` // 是否使用SSD，默认为false
	CouponId string `ucloud:"optional"` // 使用的代金券id
	IsLock   bool   `ucloud:"optional"` // 是否锁主库，默认为true

}

func (r *CreateUdbSlave) R() UResponse {
	return &CreateUdbSlaveResponse{}
}

// ---------------- DeleteUdbBackup ------------------

type DeleteUdbBackupResponse struct {
	BaseResponse
}

func (r *DeleteUdbBackupResponse) Data() interface{} {
	return r.RetCode
}

type DeleteUdbBackup struct {
	Region   string // 数据中心，请参见数据中心RegionList
	BackupId int    // 备份id，可通过DescribeUDBBackup获得

}

func (r *DeleteUdbBackup) R() UResponse {
	return &DeleteUdbBackupResponse{}
}

// ---------------- DeleteUdbInstance ------------------

type DeleteUdbInstanceResponse struct {
	BaseResponse
}

func (r *DeleteUdbInstanceResponse) Data() interface{} {
	return r.RetCode
}

type DeleteUdbInstance struct {
	Region string // 数据中心，请参见数据中心RegionList
	DBId   string // DB实例的id

}

func (r *DeleteUdbInstance) R() UResponse {
	return &DeleteUdbInstanceResponse{}
}

// ---------------- DeleteUdbParamGroup ------------------

type DeleteUdbParamGroupResponse struct {
	BaseResponse
}

func (r *DeleteUdbParamGroupResponse) Data() interface{} {
	return r.RetCode
}

type DeleteUdbParamGroup struct {
	Region  string // 数据中心，请参见数据中心RegionList
	GroupId int    // 参数组id

}

func (r *DeleteUdbParamGroup) R() UResponse {
	return &DeleteUdbParamGroupResponse{}
}

// ---------------- DescribeUdbBackupBlacklist ------------------

type DescribeUdbBackupBlacklistResponse struct {
	BaseResponse
	Blacklist string `json:",omitempty"` // 黑名单

}

func (r *DescribeUdbBackupBlacklistResponse) Data() interface{} {
	return r.Blacklist
}

type DescribeUdbBackupBlacklist struct {
	Region string // 数据中心，请参见数据中心RegionList
	DBId   string // DB实例Id

}

func (r *DescribeUdbBackupBlacklist) R() UResponse {
	return &DescribeUdbBackupBlacklistResponse{}
}

// ---------------- DescribeUdbBackup ------------------
type DescribeUdbBackupDataSet struct {
	BackupId   int    // 备份id
	BackupName string // 备份名称
	BackupTime int    // 备份时间
	BackupSize int    // 备份文件大小
	BackupType int    // 备份类型，包括0-自动，1-手动
	State      string // 备份状态Backuping          // 备份中Success            // 备份成功Failed             // 备份失败Expired            // 备份过期
	DBId       string // dbid
	DBName     string // 对应的db名称
}

type DescribeUdbBackupResponse struct {
	BaseResponse
	DataSet    []*DescribeUdbBackupDataSet `json:",omitempty"` // 备份信息
	TotalCount int                         `json:",omitempty"` // 备份总数，如果指定dbid，则是该db备份总数

}

func (r *DescribeUdbBackupResponse) Data() interface{} {
	return r.DataSet
}

type DescribeUdbBackup struct {
	Region     string // 数据中心，请参见数据中心RegionList
	Offset     int    // 分页显示的起始偏移，列表操作则指定
	Limit      int    // 分页显示的条目数，列表操作则指定
	DBId       string `ucloud:"optional"` // DB实例Id，如果指定，则只获取该db的备份信息
	BackupType int    `ucloud:"optional"` // 备份类型，包括0-自动，1-手动
	BeginTime  int    `ucloud:"optional"` // 过滤条件:起始时间(时间戳)
	EndTime    int    `ucloud:"optional"` // 过滤条件:结束时间(时间戳)

}

func (r *DescribeUdbBackup) R() UResponse {
	return &DescribeUdbBackupResponse{}
}

// ---------------- DescribeUdbInstancePrice ------------------
type DescribeUdbInstancePriceDataSet struct {
	ChargeType string  // Year， Month， Dynamic，Trial
	Price      float64 // 价格，单位分
}

type DescribeUdbInstancePriceResponse struct {
	BaseResponse
	DataSet []*DescribeUdbInstancePriceDataSet `json:",omitempty"` // 价格

}

func (r *DescribeUdbInstancePriceResponse) Data() interface{} {
	return r.DataSet
}

type DescribeUdbInstancePrice struct {
	Region      string // 数据中心，请参见数据中心RegionList
	Count       int    // 购买DB实例数量
	ChargeType  string `ucloud:"optional"` // Year， Month， Dynamic，Trial，默认: Dynamic如果不指定，则一次性获取三种计费
	Quantity    int    `ucloud:"optional"` // 购买DB的时长，默认值为1
	MemoryLimit int    // 内存限制(MB)，目前支持以下几档600M/1500M/3000M/6000M/15000M/30000M
	DiskSpace   int    // 磁盘空间(GB), 暂时支持20G - 500G
	UseSSD      bool   `ucloud:"optional"` // 是否使用SSD，默认为false

}

func (r *DescribeUdbInstancePrice) R() UResponse {
	return &DescribeUdbInstancePriceResponse{}
}

// ---------------- DescribeUdbInstance ------------------
type DescribeUdbInstanceDataSet struct {
	DBId            string                        // DB实例id
	Name            string                        // 实例名称，至少6位
	DBTypeId        string                        // DB类型id，mysql/mongodb按版本细分各有一个id1：mysql-5.5，2：mysql-5.1，3：percona-5.54：mongodb-2.4，5：mongodb-2.6，6：mysql-5.6，7：percona-5.6
	ParamGroupId    int                           // DB实例使用的配置参数组id
	AdminUser       string                        // 管理员帐户名，默认root
	VirtualIP       string                        // DB实例虚ip
	VirtualIPMac    string                        // DB实例虚ip的mac地址
	Port            int                           // 端口号，mysql默认3306，mongodb默认27017
	SrcDBId         string                        // 对mysql的slave而言是master的DBId，对master则为空，对mongodb则是副本集id
	BackupCount     int                           // 备份策略，每周备份数量，默认7次
	BackupBeginTime int                           // 备份策略，备份开始时间，单位小时计，默认3点
	BackupDuration  int                           // 备份策略，备份时间间隔，单位小时计，默认24小时
	BackupBlacklist string                        // 备份策略，备份黑名单，mongodb则不适用
	State           string                        // DB状态标记Init                           // 初始化中Fail                           // 安装失败Starting                       // 启动中Running                        // 运行Shutdown                       // 关闭中Shutoff                        // 已关闭Delete                         // 已删除Upgrading                      // 升级中Promoting                      // 提升为独库进行中Recovering                     // 恢复中Recover fail                   // 恢复失败
	CreateTime      int                           // DB实例创建时间，采用UTC计时时间戳
	ModifyTime      int                           // DB实例修改时间，采用UTC计时时间戳
	ExpiredTime     int                           // DB实例过期时间，采用UTC计时时间戳
	ChargeType      string                        // Year， Month， Dynamic，Trial，默认: Dynamic
	MemoryLimit     int                           // 内存限制(MB)，默认根据配置机型
	DiskSpace       int                           // 磁盘空间(GB), 默认根据配置机型
	UseSSD          bool                          // 是否使用SSD
	Role            string                        // DB实例角色，mysql区分master/slave，mongodb多种角色
	DiskUsedSize    int                           // DB实例磁盘已使用空间，单位GB
	DataFileSize    int                           // DB实例数据文件大小，单位GB
	SystemFileSize  int                           // DB实例系统文件大小，单位GB
	LogFileSize     int                           // DB实例日志文件大小，单位GB
	DataSet         []*DescribeUdbInstanceDataSet // 如果列表操作，则有从DB实例信息列表
}

type DescribeUdbInstanceResponse struct {
	BaseResponse
	DataSet    []*DescribeUdbInstanceDataSet `json:",omitempty"` // DB实例信息列表
	TotalCount int                           `json:",omitempty"` // 用户db组的数量，对于mysql: 主从结对数量，没有slave，则只有master mongodb: 副本集数量

}

func (r *DescribeUdbInstanceResponse) Data() interface{} {
	return r.DataSet
}

type DescribeUdbInstance struct {
	Region    string // 数据中心，请参见数据中心RegionList
	DBId      string `ucloud:"optional"` // DB实例id，如果指定则获取描述，否则为列表操作，指定Offset/Limit
	ClassType string `ucloud:"optional"` // DB种类，分为SQL和NOSQL，如果是别表操作，则需要制定
	Offset    int    `ucloud:"optional"` // 分页显示起始偏移位置，列表操作则指定
	Limit     int    `ucloud:"optional"` // 分页显示数量，列表操作则指定

}

func (r *DescribeUdbInstance) R() UResponse {
	return &DescribeUdbInstanceResponse{}
}

// ---------------- DescribeUdbInstanceState ------------------

type DescribeUdbInstanceStateResponse struct {
	BaseResponse
	State string `json:",omitempty"` // DB状态标记 Init          // 初始化中 Fail              // 安装失败 Starting          // 启动中 Running           // 运行 Shutdown          // 关闭中 Shutoff           // 已关闭 Delete            // 已删除 Upgrading         // 升级中 Promoting         // 提升为独库进行中 Recovering        // 恢复中 Recover    fail   // 恢复失败
}

func (r *DescribeUdbInstanceStateResponse) Data() interface{} {
	return r.State
}

type DescribeUdbInstanceState struct {
	Region string // 数据中心，请参见数据中心RegionList
	DBId   string // 实例的Id

}

func (r *DescribeUdbInstanceState) R() UResponse {
	return &DescribeUdbInstanceStateResponse{}
}

// ---------------- DescribeUdbParamGroup ------------------
type DescribeUdbParamGroupDataSet struct {
	GroupId     int            // 参数组id
	GroupName   string         // 参数组名称
	DBTypeId    string         // DB类型id，mysql/mongodb按版本细分各有一个id1：mysql-5.5，2：mysql-5.1，3：percona-5.54：mongodb-2.4，5：mongodb-2.6，6：mysql-5.67：percona-5.6
	Description string         // 参数组描述
	Modifiable  bool           // 参数组是否可修改，默认值为true
	ParamMember []*ParamMember `json:",omitempty"` // 参数的键值对
}
type ParamMember struct {
	Key        string // 参数名称
	Value      string // 参数值
	ValueType  int    // 参数值应用类型，包括0-unknown、10-int、20-string、30-bool
	AllowedVal string // 允许的值(根据参数类型，用分隔符表示)
	ApplyType  int    // 参数值应用类型，包括0-unknown、10-static、20-dynamic
	Modifiable bool   // 是否可更改，默认为false
	FormatType int    // 允许值得格式类型，包括PVFT_UNKOWN=0,PVFT_RANGE=10, PVFT_ENUM=20
}

type DescribeUdbParamGroupResponse struct {
	BaseResponse
	DataSet    []*DescribeUdbParamGroupDataSet `json:",omitempty"` // 参数组
	TotalCount int                             `json:",omitempty"` // 参数组总数，列表操作时才会有该参数

}

func (r *DescribeUdbParamGroupResponse) Data() interface{} {
	return r.DataSet
}

type DescribeUdbParamGroup struct {
	Region  string // 数据中心，请参见数据中心RegionList
	GroupId int    `ucloud:"optional"` // 参数组id，如果指定则获取描述，否则是列表操作，需要指定Offset/Limit
	Offset  int    `ucloud:"optional"` // 分页显示的起始偏移，列表操作则指定
	Limit   int    `ucloud:"optional"` // 分页显示的条目数，列表操作则指定

}

func (r *DescribeUdbParamGroup) R() UResponse {
	return &DescribeUdbParamGroupResponse{}
}

// ---------------- DescribeUdbType ------------------
type DescribeUdbTypeDataSet struct {
	DBTypeId string // DB类型id，mysql/mongodb按版本细分各有一个id1：mysql-5.5，2：mysql-5.1，3：percona-5.54：mongodb-2.4，5：mongodb-2.6，6：mysql-5.6，7：percona-5.6
}

type DescribeUdbTypeResponse struct {
	BaseResponse
	DataSet []*DescribeUdbTypeDataSet `json:",omitempty"` // DB类型列表

}

func (r *DescribeUdbTypeResponse) Data() interface{} {
	return r.DataSet
}

type DescribeUdbType struct {
	Region string // 数据中心，请参见数据中心RegionList

}

func (r *DescribeUdbType) R() UResponse {
	return &DescribeUdbTypeResponse{}
}

// ---------------- EditUdbBackupBlacklist ------------------

type EditUdbBackupBlacklistResponse struct {
	BaseResponse
}

func (r *EditUdbBackupBlacklistResponse) Data() interface{} {
	return r.RetCode
}

type EditUdbBackupBlacklist struct {
	Region    string // 数据中心，请参见数据中心RegionList
	DBId      string // DB实例Id
	Blacklist string // 黑名单，规范示例abc.%;user.%;city.address;

}

func (r *EditUdbBackupBlacklist) R() UResponse {
	return &EditUdbBackupBlacklistResponse{}
}

// ---------------- ModifyUdbInstanceName ------------------

type ModifyUdbInstanceNameResponse struct {
	BaseResponse
}

func (r *ModifyUdbInstanceNameResponse) Data() interface{} {
	return r.RetCode
}

type ModifyUdbInstanceName struct {
	Region string // 数据中心，请参见数据中心RegionList
	DBId   string // 实例的Id
	Name   string // 实例的新名字

}

func (r *ModifyUdbInstanceName) R() UResponse {
	return &ModifyUdbInstanceNameResponse{}
}

// ---------------- PromoteUdbSlave ------------------

type PromoteUdbSlaveResponse struct {
	BaseResponse
}

func (r *PromoteUdbSlaveResponse) Data() interface{} {
	return r.RetCode
}

type PromoteUdbSlave struct {
	Region  string // 数据中心，请参见数据中心RegionList
	DBId    string // 实例的Id
	IsForce bool   `ucloud:"optional"` // 是否强制(如果从库落后可能会禁止提升)，默认false如果落后情况下，强制提升丢失数据

}

func (r *PromoteUdbSlave) R() UResponse {
	return &PromoteUdbSlaveResponse{}
}

// ---------------- RestartUdbInstance ------------------

type RestartUdbInstanceResponse struct {
	BaseResponse
}

func (r *RestartUdbInstanceResponse) Data() interface{} {
	return r.RetCode
}

type RestartUdbInstance struct {
	Region string // 数据中心，请参见数据中心RegionList
	DBId   string // 实例的Id

}

func (r *RestartUdbInstance) R() UResponse {
	return &RestartUdbInstanceResponse{}
}

// ---------------- StartUdbInstance ------------------

type StartUdbInstanceResponse struct {
	BaseResponse
}

func (r *StartUdbInstanceResponse) Data() interface{} {
	return r.RetCode
}

type StartUdbInstance struct {
	Region string // 数据中心，请参见数据中心RegionList
	DBId   string // 实例的Id

}

func (r *StartUdbInstance) R() UResponse {
	return &StartUdbInstanceResponse{}
}

// ---------------- StopUdbInstance ------------------

type StopUdbInstanceResponse struct {
	BaseResponse
}

func (r *StopUdbInstanceResponse) Data() interface{} {
	return r.RetCode
}

type StopUdbInstance struct {
	Region string // 数据中心，请参见数据中心RegionList
	DBId   string // 实例的Id

}

func (r *StopUdbInstance) R() UResponse {
	return &StopUdbInstanceResponse{}
}

// ---------------- UpdateUdbParamGroup ------------------

type UpdateUdbParamGroupResponse struct {
	BaseResponse
}

func (r *UpdateUdbParamGroupResponse) Data() interface{} {
	return r.RetCode
}

type UpdateUdbParamGroup struct {
	Region  string // 数据中心，请参见数据中心RegionList
	GroupId int    // 配置参数组id，使用DescribeUDBParamGroup获得
	Key     string // 参数名称
	Value   string // 参数值

}

func (r *UpdateUdbParamGroup) R() UResponse {
	return &UpdateUdbParamGroupResponse{}
}

// ---------------- UploadUdbParamGroup ------------------

type UploadUdbParamGroupResponse struct {
	BaseResponse
	GroupId int `json:",omitempty"` // 配置参数组id

}

func (r *UploadUdbParamGroupResponse) Data() interface{} {
	return r.GroupId
}

type UploadUdbParamGroup struct {
	Region      string // 数据中心，请参见数据中心RegionList
	DBTypeId    string // DB类型id，mysql/mongodb按版本细分各有一个id1：mysql-5.5，2：mysql-5.1，3：percona-5.54：mongodb-2.4，5：mongodb-2.6，6：mysql-5.6，7：percona-5.6
	GroupName   string // 配置参数组名称
	Description string // 参数组描述
	Content     string // 配置内容，导入的配置内容采用base64编码。mysql只支持[mysqld]段，如：[mysqld]back_log=102character_set_server=utf8......mongodb则不需要带段，如：auth=truemaxConns=2000......

}

func (r *UploadUdbParamGroup) R() UResponse {
	return &UploadUdbParamGroupResponse{}
}
