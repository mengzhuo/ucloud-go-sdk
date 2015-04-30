package ucloud

import (
	"fmt"
	"testing"
)

// ---------------- TestAllocateEIP ------------------
func TestAllocateEIP(t *testing.T) {
	fmt.Println("UNet....")
	r := &AllocateEIP{Region: "cn-north-01",
		OperatorName: "Bgp",
		Bandwidth:    4,
	}
	cmp := `https://api.ucloud.cn/?Action=AllocateEIP&Region=cn-north-01&OperatorName=Bgp&Bandwidth=4`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestDescribeEIP ------------------
func TestDescribeEIP(t *testing.T) {
	r := &DescribeEIP{Region: "cn-north-01"}
	cmp := `https://api.ucloud.cn/?Action=DescribeEIP&Region=cn-north-01`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestUpdateEIPAttribute ------------------
func TestUpdateEIPAttribute(t *testing.T) {
	r := &UpdateEIPAttribute{Remark: "test",
		Name:   "test",
		Region: "cn-north-01",
		EIPId:  "eip-w2pew1",
		Tag:    "test",
	}
	cmp := `https://api.ucloud.cn/?Action=UpdateEIPAttribute&Region=cn-north-01&EIPId=eip-w2pew1&Name=test&Tag=test&Remark=test`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestReleaseEIP ------------------
func TestReleaseEIP(t *testing.T) {
	r := &ReleaseEIP{Region: "cn-north-01",
		EIPId: "eip-wintta",
	}
	cmp := `https://api.ucloud.cn/?Action=ReleaseEIP&Region=cn-north-01&EIPId=eip-wintta`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestBindEIP ------------------
func TestBindEIP(t *testing.T) {
	r := &BindEIP{ResourceType: "uhost",
		Region:     "cn-north-01",
		EIPId:      "eip-1inlb2",
		ResourceId: "uhost-0ttesd",
	}
	cmp := `https://api.ucloud.cn/?Action=BindEIP&Region=cn-north-01&EIPId=eip-1inlb2&ResourceType=uhost&ResourceId=uhost-0ttesd`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestUnbindEIP ------------------
func TestUnbindEIP(t *testing.T) {
	r := &UnBindEIP{ResourceType: "uhost",
		Region:     "cn-north-01",
		EIPId:      "eip-1inlb2",
		ResourceId: "uhost-0ttesd",
	}
	cmp := `https://api.ucloud.cn/?Action=UnBindEIP&Region=cn-north-01&EIPId=eip-1inlb2&ResourceType=uhost&ResourceId=uhost-0ttesd`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestModifyEIPBandwidth ------------------
func TestModifyEIPBandwidth(t *testing.T) {
	r := &ModifyEIPBandwidth{Region: "cn-north-01",
		EIPId:     "eip-dr1e2n",
		Bandwidth: 4,
	}
	cmp := `https://api.ucloud.cn/?Action=ModifyEIPBandwidth&Region=cn-north-01&EIPId=eip-dr1e2n&Bandwidth=4`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestModifyEIPWeight ------------------
func TestModifyEIPWeight(t *testing.T) {
	r := &ModifyEIPWeight{Region: "cn-north-01",
		EIPId:  "eip-dr1e2n",
		Weight: 4,
	}
	cmp := `https://api.ucloud.cn/?Action=ModifyEIPWeight&Region=cn-north-01&EIPId=eip-dr1e2n&Weight=4`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestGetEIPPrice ------------------
func TestGetEIPPrice(t *testing.T) {
	r := &GetEIPPrice{Region: "cn-north-01",
		OperatorName: "Bgp",
		Bandwidth:    4,
	}
	cmp := `https://api.ucloud.cn/?Action=GetEIPPrice&Region=cn-north-01&OperatorName=Bgp&Bandwidth=4`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestAllocateVIP ------------------
func TestAllocateVIP(t *testing.T) {
	r := &AllocateVIP{Region: "cn-north-01"}
	cmp := `https://api.ucloud.cn/?Action=AllocateVIP&Region=cn-north-01`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestDescribeVIP ------------------
func TestDescribeVIP(t *testing.T) {
	r := &DescribeVIP{Region: "cn-north-01"}
	cmp := `https://api.ucloud.cn/?Action=DescribeVIP&Region=cn-north-01`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestReleaseVIP ------------------
func TestReleaseVIP(t *testing.T) {
	r := &ReleaseVIP{Region: "cn-north-01",
		VIP: "10.10.3.13",
	}
	cmp := `https://api.ucloud.cn/?Action=ReleaseVIP&Region=cn-north-01&VIP=10.10.3.13`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestDescribeSecurityGroup ------------------
func TestDescribeSecurityGroup(t *testing.T) {
	r := &DescribeSecurityGroup{Region: "cn-north-01",
		GroupId: 6583,
	}
	cmp := `https://api.ucloud.cn/?Action=DescribeSecurityGroup&Region=cn-north-01&GroupId=6583`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestDescribeSecurityGroupResource ------------------
func TestDescribeSecurityGroupResource(t *testing.T) {
	r := &DescribeSecurityGroupResource{Region: "cn-north-01",
		GroupId: 6583,
	}
	cmp := `https://api.ucloud.cn/?Action=DescribeSecurityGroupResource&Region=cn-north-01&GroupId=6583`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestCreateSecurityGroup ------------------
func TestCreateSecurityGroup(t *testing.T) {
	r := &CreateSecurityGroup{Region: "cn-north-01",
		Rule:      []string{"TCP|3306|0.0.0.0/0|DROP|50", "UDP|53|0.0.0.0/0|ACCEPT|50"},
		GroupName: "NewSecurityGroup",
	}
	cmp := `https://api.ucloud.cn/?Action=CreateSecurityGroup&Region=cn-north-01&GroupName=NewSecurityGroup&Rule.1=UDP|53|0.0.0.0/0|ACCEPT|50&Rule.0=TCP|3306|0.0.0.0/0|DROP|50`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestUpdateSecurityGroup ------------------
func TestUpdateSecurityGroup(t *testing.T) {
	r := &UpdateSecurityGroup{Region: "cn-north-01",
		GroupId: "6583",
		Rule:    []string{"TCP|3306|0.0.0.0/0|DROP|50", "UDP|53|0.0.0.0/0|ACCEPT|50"},
	}
	cmp := `https://api.ucloud.cn/?Action=UpdateSecurityGroup&Region=cn-north-01&GroupId=6583&Rule.1=UDP|53|0.0.0.0/0|ACCEPT|50&Rule.0=TCP|3306|0.0.0.0/0|DROP|50`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestGrantSecurityGroup ------------------
func TestGrantSecurityGroup(t *testing.T) {
	r := &GrantSecurityGroup{ResourceType: "UHost",
		Region:     "cn-north-01",
		GroupId:    "6583",
		ResourceId: "uhost-w4d53b",
	}
	cmp := `https://api.ucloud.cn/?Action=GrantSecurityGroup&Region=cn-north-01&GroupId=6583&ResourceType=UHost&ResourceId=uhost-w4d53b`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestDeleteSecurityGroup ------------------
func TestDeleteSecurityGroup(t *testing.T) {
	r := &DeleteSecurityGroup{Region: "cn-north-01",
		GroupId: "6583",
	}
	cmp := `https://api.ucloud.cn/?Action=DeleteSecurityGroup&Region=cn-north-01&GroupId=6583`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}
