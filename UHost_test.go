package ucloud

import (
	"fmt"
	"testing"
)

func TestCreateUHostInstance(t *testing.T) {
	rsp, err := u.CreateUHostInstance("cn-north-01", "0b560b1d-f3b1-43de-a492-08875d79211b",
		"Password", "UGFzc3dvcmQxCg==", 1, 2048, 10, "UHostTest", "基础网络", "Web防火墙", "Trial", 1)

	fmt.Println(rsp, err)
}
