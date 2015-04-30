package ucloud

import (
	"fmt"
	"testing"
)

// ---------------- TestCreateUDiskSnapshot ------------------
func TestCreateUDiskSnapshot(t *testing.T) {
	fmt.Println("UDisk.....")
	r := &CreateUDiskSnapshot{
		Comment: "test",
		Region:  "1",
		Name:    "test_snapshot",
		UDiskId: "2AFCD36A-2A47-4D26-8514-CAA3FC2DC6BF",
	}
	cmp := `https://api.ucloud.cn/?Action=CreateUDiskSnapshot&Region=1&UDiskId=2AFCD36A-2A47-4D26-8514-CAA3FC2DC6BF&Name=test_snapshot&Comment=test`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestCreateUDisk ------------------
func TestCreateUDisk(t *testing.T) {
	r := &CreateUDisk{Region: "1",
		Name: "test_udisk",
		Size: 1,
	}
	cmp := `https://api.ucloud.cn/?Action=CreateUDisk&Region=1&Size=1&Name=test_udisk`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestCloneUDisk ------------------
func TestCloneUDisk(t *testing.T) {
	r := &CloneUDisk{
		SourceId: "F207B59C-BA54-4455-9CF1-996DB5A7CD22",
		Region:   "1",
	}
	cmp := `https://api.ucloud.cn/?Action=CloneUDisk&Region=1&SourceId=F207B59C-BA54-4455-9CF1-996DB5A7CD22`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestCloneUDiskSnapshot ------------------
func TestCloneUDiskSnapshot(t *testing.T) {
	r := &CloneUDiskSnapshot{SourceId: "F207B59C-BA54-4455-9CF1-996DB5A7CD22",
		Region: "1",
		Size:   20,
	}
	cmp := `https://api.ucloud.cn/?Action=CloneUDiskSnapshot&Region=1&Size=20&SourceId=F207B59C-BA54-4455-9CF1-996DB5A7CD22`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestDeleteUDiskSnapshot ------------------
func TestDeleteUDiskSnapshot(t *testing.T) {
	r := &DeleteUDiskSnapshot{Region: "1",
		SnapshotId: "736aac0a-0fcc-4ba7-b2d6-30661387e02c",
	}
	cmp := `https://api.ucloud.cn/?Action=DeleteUDiskSnapshot&Region=1&SnapshotId=736aac0a-0fcc-4ba7-b2d6-30661387e02c`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestDeleteUDisk ------------------
func TestDeleteUDisk(t *testing.T) {
	r := &DeleteUDisk{
		Region:  "1",
		UDiskId: "2AFCD36A-2A47-4D26-8514-CAA3FC2DC6BF",
	}
	cmp := `https://api.ucloud.cn/?Action=DeleteUDisk&Region=1&UDiskId=2AFCD36A-2A47-4D26-8514-CAA3FC2DC6BF`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestDescribeUDiskSnapshot ------------------
func TestDescribeUDiskSnapshot(t *testing.T) {
	r := &DescribeUDiskSnapshot{Region: "1",
		SnapshotId: "736aac0a-0fcc-4ba7-b2d6-30661387e02c",
		Offset:     0,
	}
	cmp := `https://api.ucloud.cn/?Action=DescribeUDiskSnapshot&Region=1&SnapshotId=736aac0a-0fcc-4ba7-b2d6-30661387e02c`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestDescribeUDisk ------------------
func TestDescribeUDisk(t *testing.T) {
	r := &DescribeUDisk{Region: "1",
		Limit:   20,
		UDiskId: []string{"B40421B6-7EBC-4F79-B7C3-8BD78D51A4B0"},
		Offset:  0,
	}
	cmp := `https://api.ucloud.cn/?Action=DescribeUDisk&Region=1&UDiskId.0=B40421B6-7EBC-4F79-B7C3-8BD78D51A4B0&Offset=0&Limit=20`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestDescribeUDiskPrice ------------------
func TestDescribeUDiskPrice(t *testing.T) {
	r := &DescribeUDiskPrice{
		Region:     "cn-east-01",
		Quantity:   12,
		ChargeType: "Month",
		Size:       1,
	}
	cmp := `https://api.ucloud.cn/udisk/?Action=DescribeUDiskPrice&Region=cn-east-01&Size=1&ChargeType=Month&Quantity=12`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestDescribeUDiskUpgradePrice ------------------
func TestDescribeUDiskUpgradePrice(t *testing.T) {
	r := &DescribeUDiskUpgradePrice{
		Region:   "cn-east-01",
		Size:     1,
		SourceId: "ubs_abdc2",
	}
	cmp := `https://api.ucloud.cn/udisk/?Action=DescribeUDiskUpgradePrice&Region=cn-east-01&Size=1&SourceId=ubs_abdc2`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestRenameUDisk ------------------
func TestRenameUDisk(t *testing.T) {
	r := &RenameUDisk{Region: "1",
		UDiskName: "udisk_2",
		UDiskId:   "2AFCD36A-2A47-4D26-8514-CAA3FC2DC6BF",
	}
	cmp := `https://api.ucloud.cn/?Action=RenameUDisk&Region=1&UDiskId=2AFCD36A-2A47-4D26-8514-CAA3FC2DC6BF&UDiskName=udisk_2`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestResizeUDisk ------------------
func TestResizeUDisk(t *testing.T) {
	r := &ResizeUDisk{Region: "1",
		UDiskId: "4b196682-e45b-4ac4-a431-31f231c68313",
		Size:    20,
	}
	cmp := `https://api.ucloud.cn/?Action=ResizeUDisk&Region=1&UDiskId=4b196682-e45b-4ac4-a431-31f231c68313&Size=20`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}
