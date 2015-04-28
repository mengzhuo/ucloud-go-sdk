// Package ucloud provides ...
package ucloud

import (
	"testing"
)

func TestCreateUhostInstance(t *testing.T) {
	r := &CreateUHostInstance{Region: "cn-north-01",
		ImageId:    "f43736e1-65a5-4bea-ad2e-8a46e18883c2",
		CPU:        2,
		Memory:     2048,
		DiskSpace:  10,
		LoginMode:  "Password",
		Password:   "UGFzc3dvcmQxCg==",
		Name:       "UCloudExample01",
		ChargeType: "Month",
		Quantity:   1,
	}

	cmp := `https://api.ucloud.cn/?Action=CreateUHostInstance&Region=cn-north-01&ImageId=f43736e1-65a5-4bea-ad2e-8a46e18883c2&CPU=2&Memory=2048&DiskSpace=10&LoginMode=Password&Password=UGFzc3dvcmQxCg%3D%3D&Name=UCloudExample01&ChargeType=Month&Quantity=1`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}
