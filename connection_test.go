package ucloud

import (
	"os"
)

var publicKey = os.Getenv("UCLOUD_PUBLIC_KEY")
var privateKey = os.Getenv("UCLOUD_PRIVATE_KEY")
var u = NewUcloudApiClient("https://api.ucloud.cn", publicKey, privateKey, "1", "1")
