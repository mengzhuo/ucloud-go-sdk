// Package ucloud provides ...
package ucloud

import (
	"os"
	"testing"

	"fmt"
)

var publicKey = os.Getenv("UCLOUD_PUBLIC_KEY")
var privateKey = os.Getenv("UCLOUD_PRIVATE_KEY")
var u = NewUcloudApiClient("https://api.ucloud.cn", publicKey, privateKey, "1", "1")

func TestSendSms(t *testing.T) {
	rsp, err := u.Do(&SendSms{"Hello world", []string{"18600754601"}})
	fmt.Println(rsp, err)
}
