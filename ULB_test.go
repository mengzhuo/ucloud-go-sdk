package ucloud

import (
	"fmt"
	"testing"
)

// ---------------- TestCreateULB ------------------
func TestCreateULB(t *testing.T) {
	fmt.Println("ULB......")
	r := &CreateULB{Region: "cn-north-01"}
	cmp := `https://api.ucloud.cn/?Action=CreateULB&Region=cn-north-01`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestDeleteULB ------------------
func TestDeleteULB(t *testing.T) {
	r := &DeleteULB{Region: "cn-north-01",
		ULBId: "Abcdefg",
	}
	cmp := `https://api.ucloud.cn/?Action=DeleteULB&Region=cn-north-01&ULBId=Abcdefg`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestDescribeULB ------------------
func TestDescribeULB(t *testing.T) {
	r := &DescribeULB{
		Region: "cn-south-01",
	}
	cmp := `http://api.spark.ucloud.cn/?Action=DescribeULB&Region=cn-south-01`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestUpdateULBAttribute ------------------
func TestUpdateULBAttribute(t *testing.T) {
	r := &UpdateULBAttribute{
		Remark: "test",
		Name:   "test",
		ULBId:  "ulb-xjqquy",
		Region: "cn-north-01",
		Tag:    "test",
	}
	cmp := `http://api.ucloud.cn/?Action=UpdateULBAttribute&Region=cn-north-01&ULBId=ulb-xjqquy&Name=test&Tag=test&Remark=test`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestCreateVServer ------------------
func TestCreateVServer(t *testing.T) {
	r := &CreateVServer{PersistenceType: "None",
		Protocol:     "TCP",
		ULBId:        "ulb-b1jhgw",
		Region:       "cn-north-01",
		VServerName:  "testvserver1",
		Method:       "Source",
		FrontendPort: 80,
	}
	cmp := `https://api.ucloud.cn/?Action=CreateVServer&Region=cn-north-01&Action=CreateVServer&ULBId=ulb-b1jhgw&PersistenceType=None&Protocol=TCP&VServerName=testvserver1&Method=Source&FrontendPort=80`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestDeleteVServer ------------------
func TestDeleteVServer(t *testing.T) {
	r := &DeleteVServer{Region: "cn-north-01",
		VServerId: "19ac0869-d0f2-4372-bd13-eeb1968d2951",
		ULBId:     "ulb-igw0gd",
	}
	cmp := `https://api.ucloud.cn/?Action=DeleteVServer&Region=cn-north-01&ULBId=ulb-igw0gd&VServerId=19ac0869-d0f2-4372-bd13-eeb1968d2951`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestUpdateVServerAttribute ------------------
func TestUpdateVServerAttribute(t *testing.T) {
	r := &UpdateVServerAttribute{Region: "cn-north-03",
		VServerId:   "8aae44ba-f3fd-4162-9e32-2e64b89b646b",
		VServerName: "Testapi",
		ULBId:       "ulb-kix4tp",
	}
	cmp := `https://api.ucloud.cn/?Action=UpdateVServerAttribute&Region=cn-north-03&ULBId=ulb-kix4tp&VServerId=8aae44ba-f3fd-4162-9e32-2e64b89b646b&VServerName=Testapi`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestAllocateBackend ------------------
func TestAllocateBackend(t *testing.T) {
	r := &AllocateBackend{ULBId: "ulb-a3qq0o",
		ResourceType: "UHost",
		ResourceId:   "uhost-hicy2y",
		Enabled:      1,
		VServerId:    "d132dec3-0340-4a04-b129-be08074fbf9d",
		Port:         80,
		Region:       "cn-north-01",
	}
	cmp := `https://api.ucloud.cn/?Action=AllocateBackend&ULBId=ulb-a3qq0o&VServerId=d132dec3-0340-4a04-b129-be08074fbf9d&ResourceType=UHost&ResourceId=uhost-hicy2y&Port=80&Enabled=1&Region=cn-north-01`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestReleaseBackend ------------------
func TestReleaseBackend(t *testing.T) {
	r := &ReleaseBackend{Region: "cn-north-01",
		BackendId: "6a049334-77dc-49b9-a55d-fce8d3d0d9e8",
		ULBId:     "ulb-a3qq0o",
	}
	cmp := `https://api.ucloud.cn/?Action=ReleaseBackend&Region=cn-north-01&ULBId=ulb-a3qq0o&BackendId=6a049334-77dc-49b9-a55d-fce8d3d0d9e8`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestUpdateBackendAttribute ------------------
func TestUpdateBackendAttribute(t *testing.T) {
	r := &UpdateBackendAttribute{ULBId: "ulb-a3qq0o",
		Region:    "cn-north-01",
		Enabled:   0,
		BackendId: "6a049334-77dc-49b9-a55d-fce8d3d0d9e8",
		Port:      8080,
	}
	cmp := `https://api.ucloud.cn/?Action=UpdateBackendAttribute&Region=cn-north-01&ULBId=ulb-a3qq0o&BackendId=6a049334-77dc-49b9-a55d-fce8d3d0d9e8&Port=8080`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestCreateSSL ------------------
func TestCreateSSL(t *testing.T) {
	r := &CreateSSL{SSLName: "new-ssl",
		Region:     "cn-north-01",
		SSLContent: "-----BEGIN",
	}
	cmp := `https://api.ucloud.cn/?Action=CreateSSL&Region=cn-north-01&SSLName=new-ssl&SSLContent=-----BEGIN&nbspRSA&nbspPRIVATE&nbspKEY-----....`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestDeleteSSL ------------------
func TestDeleteSSL(t *testing.T) {
	r := &DeleteSSL{Region: "cn-north-01",
		SSLId: "ssl-s1v2e0",
	}
	cmp := `https://api.ucloud.cn/?Action=DeleteSSL&Region=cn-north-01&SSLId=ssl-s1v2e0`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestBindSSL ------------------
func TestBindSSL(t *testing.T) {
	r := &BindSSL{VServerId: "8aae44ba-f3fd-4162-9e32-2e64b89b646b",
		Region: "cn-north-01",
		SSLId:  "ssl-kskl39s",
		ULBId:  "ulb-kix4tp",
	}
	cmp := `https://api.ucloud.cn/?Action=BindSSL&Region=cn-north-01&SSLId=ssl-kskl39s&ULBId=ulb-kix4tp&VServerId=8aae44ba-f3fd-4162-9e32-2e64b89b646b`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestDescribeSSL ------------------
func TestDescribeSSL(t *testing.T) {
	r := &DescribeSSL{Region: "cn-north-01"}
	cmp := `https://api.ucloud.cn/?Action=DescribeSSL&Region=cn-north-01`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestCreatePolicyGroup ------------------
func TestCreatePolicyGroup(t *testing.T) {
	r := &CreatePolicyGroup{Region: "cn-north-01",
		GroupName: "new-policy",
	}
	cmp := `https://api.ucloud.cn/?Action=CreatePolicyGroup&Region=cn-north-01&GroupName=new-policy`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestDeletePolicyGroup ------------------
func TestDeletePolicyGroup(t *testing.T) {
	r := &DeletePolicyGroup{Region: "cn-north-01",
		GroupId: "ulb-fr-2axjbg",
	}
	cmp := `https://api.ucloud.cn/?Action=DeletePolicyGroup&Region=cn-north-01&GroupId=ulb-fr-2axjbg`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestDescribePolicyGroup ------------------
func TestDescribePolicyGroup(t *testing.T) {
	r := &DescribePolicyGroup{Region: "cn-north-01"}
	cmp := `https://api.ucloud.cn/?Action=DescribePolicyGroup&Region=cn-north-01`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestUpdatePolicyGroupAttribute ------------------
func TestUpdatePolicyGroupAttribute(t *testing.T) {
	r := &UpdatePolicyGroupAttribute{Region: "cn-north-01",
		GroupId:   "ulb-fr-2axjbg",
		GroupName: "new-policy-name",
	}
	cmp := `https://api.ucloud.cn/?Action=UpdatePolicyGroupAttribute&Region=cn-north-01&GroupId=ulb-fr-2axjbg&GroupName=new-policy-name`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestCreatePolicy ------------------
func TestCreatePolicy(t *testing.T) {
	r := &CreatePolicy{VServerId: "8aae44ba-f3fd-4162-9e32-2e64b89b646b",
		Region:    "cn-north-01",
		ULBId:     "ulb-kix54p",
		BackendId: []string{"6277e964-83de-4548-b598-6b4d5a67a442"},
		GroupId:   "ulb-fr-2axjbg",
		Match:     "ok",
	}
	cmp := `https://api.ucloud.cn/?Action=CreatePolicy&Region=cn-north-01&GroupId=ulb-fr-2axjbg&Match=ok&ULBId=ulb-kix54p&VServerId=8aae44ba-f3fd-4162-9e32-2e64b89b646b&BackendId.0=6277e964-83de-4548-b598-6b4d5a67a442`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestDeletePolicy ------------------
func TestDeletePolicy(t *testing.T) {
	r := &DeletePolicy{Region: "cn-north-01",
		GroupId:  "ulb-fr-2axjbg",
		PolicyId: "0a074d02-b7c9-4a5d-9acf-414081f1b85f",
	}
	cmp := `https://api.ucloud.cn/?Action=DeletePolicy&Region=cn-north-01&GroupId=ulb-fr-2axjbg&PolicyId=0a074d02-b7c9-4a5d-9acf-414081f1b85f`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}
