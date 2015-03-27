package ucloud

type DescribeUHostInstanceResponse struct{ *BaseResponse }
type TerminateUHostInstanceResponse struct{ *BaseResponse }
type ResizeUHostInstanceResponse struct{ *BaseResponse }
type ReinstallUHostInstanceResponse struct{ *BaseResponse }
type StartUHostInstanceResponse struct{ *BaseResponse }
type StopUHostInstanceResponse struct{ *BaseResponse }
type RebootUHostInstanceResponse struct{ *BaseResponse }
type ResetUHostInstancePasswordResponse struct{ *BaseResponse }
type ModifyUHostInstanceNameResponse struct{ *BaseResponse }
type ModifyUHostInstanceTagResponse struct{ *BaseResponse }
type ModifyUHostInstanceRemarkResponse struct{ *BaseResponse }
type GetUHostInstancePriceResponse struct{ *BaseResponse }
type GetUHostInstanceVncInfoResponse struct{ *BaseResponse }
type DescribeImageResponse struct{ *BaseResponse }
type CreateCustomImageResponse struct{ *BaseResponse }
type TerminateCustomImageResponse struct{ *BaseResponse }
type AttachUdiskResponse struct{ *BaseResponse }
type DetachUdiskResponse struct{ *BaseResponse }
type CreateUHostInstanceSnapshotResponse struct{ *BaseResponse }
type DescribeUHostInstanceSnapshotResponse struct{ *BaseResponse }

type CreateUHostInstanceResponse struct {
	*BaseResponse
	UHostIds []string `json:",omitempty"`
}

type CreateUHostInstance struct {
	Region          string
	ImageId         string
	LoginMode       string
	Password        string
	CPU             int
	Memory          int
	DiskSpace       int
	Name            string
	NetworkId       string
	SecurityGroupId string
	ChargeType      string
	Quantity        int
}

func (r *CreateUHostInstance) R() (rsp *CreateUHostInstanceResponse) {
	return
}

func (u *UcloudApiClient) DescribeUHostInstance() (rsp *DescribeUHostInstanceResponse, err error) {
	return
}
func (u *UcloudApiClient) TerminateUHostInstance() (rsp *TerminateUHostInstanceResponse, err error) {
	return
}
func (u *UcloudApiClient) ResizeUHostInstance() (rsp *ResizeUHostInstanceResponse, err error) {
	return
}
func (u *UcloudApiClient) ReinstallUHostInstance() (rsp *ReinstallUHostInstanceResponse, err error) {
	return
}
func (u *UcloudApiClient) StartUHostInstance() (rsp *StartUHostInstanceResponse, err error) {
	return
}
func (u *UcloudApiClient) StopUHostInstance() (rsp *StopUHostInstanceResponse, err error) {
	return
}
func (u *UcloudApiClient) RebootUHostInstance() (rsp *RebootUHostInstanceResponse, err error) {
	return
}
func (u *UcloudApiClient) ResetUHostInstancePassword() (rsp *ResetUHostInstancePasswordResponse, err error) {
	return
}
func (u *UcloudApiClient) ModifyUHostInstanceName() (rsp *ModifyUHostInstanceNameResponse, err error) {
	return
}
func (u *UcloudApiClient) ModifyUHostInstanceTag() (rsp *ModifyUHostInstanceTagResponse, err error) {
	return
}
func (u *UcloudApiClient) ModifyUHostInstanceRemark() (rsp *ModifyUHostInstanceRemarkResponse, err error) {
	return
}
func (u *UcloudApiClient) GetUHostInstancePrice() (rsp *GetUHostInstancePriceResponse, err error) {
	return
}
func (u *UcloudApiClient) GetUHostInstanceVncInfo() (rsp *GetUHostInstanceVncInfoResponse, err error) {
	return
}
func (u *UcloudApiClient) DescribeImage() (rsp *DescribeImageResponse, err error) {
	return
}
func (u *UcloudApiClient) CreateCustomImage() (rsp *CreateCustomImageResponse, err error) {
	return
}
func (u *UcloudApiClient) TerminateCustomImage() (rsp *TerminateCustomImageResponse, err error) {
	return
}
func (u *UcloudApiClient) AttachUdisk() (rsp *AttachUdiskResponse, err error) {
	return
}
func (u *UcloudApiClient) DetachUdisk() (rsp *DetachUdiskResponse, err error) {
	return
}
func (u *UcloudApiClient) CreateUHostInstanceSnapshot() (rsp *CreateUHostInstanceSnapshotResponse, err error) {
	return
}
func (u *UcloudApiClient) DescribeUHostInstanceSnapshot() (rsp *DescribeUHostInstanceSnapshotResponse, err error) {
	return
}
