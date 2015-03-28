// Package ucloud provides ...
package ucloud

import (
	"fmt"
	"testing"
)

func TestCreateUHostInstance(t *testing.T) {

	rsp, err := u.Do(&CreateUHostInstance{Region: "cn-east-01", ImageId: "UImage", LoginMode: "Password",
		Password: "UGFzc3dvcmQxCg==",
		Name:     "Uhost11222"})
	fmt.Println(rsp.Data().([]string))
	if !rsp.OK() || err != nil {
		t.Fatal(rsp, err)
	}
}

func TestDescribeUHostInstance(t *testing.T) {
	rsp, err := u.Do(&DescribeUHostInstance{Region: "cn-east-01"})
	fmt.Println(rsp.Data().([]*UHostSet)[0])
	if !rsp.OK() || err != nil {
		t.Fatal(rsp, err)
	}
}
