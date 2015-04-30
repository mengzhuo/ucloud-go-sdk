package ucloud

import (
	"fmt"
	"testing"
)

// ---------------- TestCreateUhostInstance ------------------
func TestCreateUhostInstance(t *testing.T) {
	fmt.Println("UHost.....")
	r := &CreateUHostInstance{Name: "UCloudExample01",
		DiskSpace:  10,
		LoginMode:  "Password",
		Region:     "cn-north-01",
		ImageId:    "f43736e1-65a5-4bea-ad2e-8a46e18883c2",
		ChargeType: "Month",
		Memory:     2048,
		Password:   "UGFzc3dvcmQxCg==",
		CPU:        2,
		Quantity:   1,
	}
	cmp := `https://api.ucloud.cn/?Action=CreateUHostInstance&Region=cn-north-01&ImageId=f43736e1-65a5-4bea-ad2e-8a46e18883c2&CPU=2&Memory=2048&DiskSpace=10&LoginMode=Password&Password=UGFzc3dvcmQxCg%3D%3D&Name=UCloudExample01&ChargeType=Month&Quantity=1`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestDescribeUhostInstance ------------------
func TestDescribeUHostInstance(t *testing.T) {
	r := &DescribeUHostInstance{Region: "cn-north-01",
		UHostIds: []string{"uhost-qs20fr"},
		Limit:    20,
		Offset:   0,
	}
	cmp := `https://api.ucloud.cn/?Action=DescribeUHostInstance&Region=cn-north-01&UHostIds.0=uhost-qs20fr&Limit=20`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestTerminateUhostInstance ------------------
func TestTerminateUHostInstance(t *testing.T) {
	r := &TerminateUHostInstance{Region: "cn-north-01",
		UHostId: "uhost-qs20fr",
	}
	cmp := `https://api.ucloud.cn/?Action=TerminateUHostInstance&Region=cn-north-01&UHostId=uhost-qs20fr`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestResizeUhostInstance ------------------
func TestResizeUHostInstance(t *testing.T) {
	r := &ResizeUHostInstance{Region: "cn-north-01",
		UHostId: "uhost-qs20fr",
		CPU:     4,
	}
	cmp := `https://api.ucloud.cn/?Action=ResizeUHostInstance&Region=cn-north-01&UHostId=uhost-qs20fr&CPU=4`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestReinstallUhostInstance ------------------
func TestReinstallUHostInstance(t *testing.T) {
	r := &ReinstallUHostInstance{Region: "cn-north-01",
		UHostId:  "uhost-qs20fr",
		Password: "UCloud123",
	}
	cmp := `https://api.ucloud.cn/?Action=ReinstallUHostInstance&Region=cn-north-01&UHostId=uhost-qs20fr&Password=UCloud123`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestStartUhostInstance ------------------
func TestStartUHostInstance(t *testing.T) {
	r := &StartUHostInstance{Region: "cn-north-01",
		UHostId: "uhost-qs20fr",
	}
	cmp := `https://api.ucloud.cn/?Action=StartUHostInstance&Region=cn-north-01&UHostId=uhost-qs20fr`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestStopUhostInstance ------------------
func TestStopUHostInstance(t *testing.T) {
	r := &StopUHostInstance{Region: "cn-north-01",
		UHostId: "uhost-qs20fr",
	}
	cmp := `https://api.ucloud.cn/?Action=StopUHostInstance&Region=cn-north-01&UHostId=uhost-qs20fr`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestRebootUhostInstance ------------------
func TestRebootUHostInstance(t *testing.T) {
	r := &RebootUHostInstance{Region: "cn-north-01",
		UHostId: "uhost-qs20fr",
	}
	cmp := `https://api.ucloud.cn/?Action=RebootUHostInstance&Region=cn-north-01&UHostId=uhost-qs20fr`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestResetUhostInstancePassword ------------------
func TestResetUHostInstancePassword(t *testing.T) {
	r := &ResetUHostInstancePassword{Region: "cn-north-01",
		UHostId:  "uhost-qs20fr",
		Password: "UCloud123",
	}
	cmp := `https://api.ucloud.cn/?Action=ResetUHostInstancePassword&Region=cn-north-01&UHostId=uhost-qs20fr&Password=UCloud123`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestModifyUhostInstanceName ------------------
func TestModifyUHostInstanceName(t *testing.T) {
	r := &ModifyUHostInstanceName{Region: "cn-north-01",
		UHostId: "uhost-qs20fr",
		Name:    "Test",
	}
	cmp := `https://api.ucloud.cn/?Action=ModifyUHostInstanceName&Region=cn-north-01&UHostId=uhost-qs20fr&Name=Test`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestModifyUhostInstanceTag ------------------
func TestModifyUHostInstanceTag(t *testing.T) {
	r := &ModifyUHostInstanceTag{Region: "cn-north-01",
		UHostId: "uhost-qs20fr",
		Tag:     "Test",
	}
	cmp := `https://api.ucloud.cn/?Action=ModifyUHostInstanceTag&Region=cn-north-01&UHostId=uhost-qs20fr&Tag=Test`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestModifyUhostInstanceRemark ------------------
func TestModifyUHostInstanceRemark(t *testing.T) {
	r := &ModifyUHostInstanceRemark{Region: "cn-north-01",
		UHostId: "uhost-qs20fr",
		Remark:  "Test",
	}
	cmp := `https://api.ucloud.cn/?Action=ModifyUHostInstanceRemark&Region=cn-north-01&UHostId=uhost-qs20fr&Remark=Test`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestGetUhostInstancePrice ------------------
func TestGetUHostInstancePrice(t *testing.T) {
	r := &GetUHostInstancePrice{Count: 1,
		DiskSpace:  10,
		Region:     "cn-north-01",
		ImageId:    "f43736e1-65a5-4bea-ad2e-8a46e18883c2",
		ChargeType: "Dynamic",
		Memory:     4096,
		CPU:        2,
	}
	cmp := `https://api.ucloud.cn/?Action=GetUHostInstancePrice&Region=cn-north-01&ImageId=f43736e1-65a5-4bea-ad2e-8a46e18883c2&CPU=2&Memory=4096&Count=1&ChargeType=Dynamic&DiskSpace=10`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestGetUhostInstanceVncInfo ------------------
func TestGetUHostInstanceVncInfo(t *testing.T) {
	r := &GetUHostInstanceVncInfo{Region: "cn-north-01",
		UHostId: "uhost-qs20fr",
	}
	cmp := `https://api.ucloud.cn/?Action=GetUHostInstanceVncInfo&Region=cn-north-01&UHostId=uhost-qs20fr`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestDescribeImage ------------------
func TestDescribeImage(t *testing.T) {
	r := &DescribeImage{OsType: "Windows",
		Region:    "cn-north-01",
		ImageType: "Base",
	}
	cmp := `https://api.ucloud.cn/?Action=DescribeImage&Region=cn-north-01&ImageType=Base&OsType=Windows`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestCreateCustomImage ------------------
func TestCreateCustomImage(t *testing.T) {
	r := &CreateCustomImage{Region: "cn-north-01",
		UHostId:   "uhost-qs20fr",
		ImageName: "Test",
	}
	cmp := `https://api.ucloud.cn/?Action=CreateCustomImage&Region=cn-north-01&UHostId=uhost-qs20fr&ImageName=Test`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestTerminateCustomImage ------------------
func TestTerminateCustomImage(t *testing.T) {
	r := &TerminateCustomImage{Region: "cn-north-01",
		ImageId: "6fa91ff8-663c-4a7c-b054-1824aa41a1a3",
	}
	cmp := `https://api.ucloud.cn/?Action=TerminateCustomImage&Region=cn-north-01&ImageId=6fa91ff8-663c-4a7c-b054-1824aa41a1a3`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestAttachUdisk ------------------
func TestAttachUDisk(t *testing.T) {
	r := &AttachUDisk{Region: "cn-north-01",
		UHostId: "uhost-qs20fr",
		UDiskId: "0b560b1d-f3b1-43de-a492-08875d79211b",
	}
	cmp := `https://api.ucloud.cn/?Action=AttachUDisk&Region=cn-north-01&UHostId=uhost-qs20fr&UDiskId=0b560b1d-f3b1-43de-a492-08875d79211b`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestDetachUdisk ------------------
func TestDetachUDisk(t *testing.T) {
	r := &DetachUDisk{Region: "cn-north-01",
		UHostId: "uhost-qs20fr",
		UDiskId: "0b560b1d-f3b1-43de-a492-08875d79211b",
	}
	cmp := `https://api.ucloud.cn/?Action=DetachUDisk&Region=cn-north-01&UHostId=uhost-qs20fr&UDiskId=0b560b1d-f3b1-43de-a492-08875d79211b`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestCreateUHostInstanceSnapshot ------------------
func TestCreateUHostInstanceSnapshot(t *testing.T) {
	r := &CreateUHostInstanceSnapshot{Region: "cn-north-01",
		UHostId: "uhost-qs20fr",
	}
	cmp := `https://api.ucloud.cn/?Action=CreateUHostInstanceSnapshot&Region=cn-north-01&UHostId=uhost-qs20fr`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestDescribeUhostInstanceSnapshot ------------------
func TestDescribeUHostInstanceSnapshot(t *testing.T) {
	r := &DescribeUHostInstanceSnapshot{Region: "cn-north-01",
		UHostId: "uhost-qs20fr",
	}
	cmp := `https://api.ucloud.cn/?Action=DescribeUHostInstanceSnapshot&Region=cn-north-01&UHostId=uhost-qs20fr`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}
