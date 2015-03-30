// Package ucloud provides ...
package ucloud

import (
	"fmt"
	"testing"
)

const REGION = "cn-east-01"

func TestFromImageToUHost(t *testing.T) {

	var imageid string

	rsp, err := u.Do(&DescribeImage{Region: REGION})
	if err != nil {
		t.Fatal(rsp, err)
	}
	images := rsp.Data().([]*ImageSet)
	for _, v := range images {
		imageid = v.ImageId
		fmt.Println(v.ImageId, v.OsType)
		break
	}

	rsp, err = u.Do(&CreateUHostInstance{Region: REGION,
		ImageId:   imageid,
		LoginMode: "Password",
		Password:  "cGlHVWhlbkRBCg==", // piGUhenDA
		Name:      "Uhost11222"})
	if err != nil {
		t.Fatal(rsp, err)
	}

}

func TestDescribeUHostInstance(t *testing.T) {
	rsp, err := u.Do(&DescribeUHostInstance{Region: REGION})
	if err != nil {
		t.Fatal(rsp, err)
	}
}
