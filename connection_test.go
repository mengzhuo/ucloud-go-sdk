package ucloud

import (
	"fmt"
	"net/url"
	"os"
)

var publicKey = os.Getenv("UCLOUD_PUBLIC_KEY")
var privateKey = os.Getenv("UCLOUD_PRIVATE_KEY")
var u = NewUcloudApiClient("https://api.ucloud.cn", publicKey, privateKey, "1", "1")

func FakeGetAndCmp(req URequest, cmp string) error {
	params, err := u.MakeParams(req)
	if err != nil {
		return fmt.Errorf("MakeParams %v FAILED %s", req, err)
	}
	uri := u.MakeURI("/", params)

	gen_uri, err := url.Parse(uri)
	cmp_uri, err := url.Parse(cmp)
	if err != nil {
		return err
	}

	for k, v := range cmp_uri.Query() {
		gen := gen_uri.Query().Get(k)
		cmp := v[0]
		if cmp == "true" {
			cmp = "True"
		}
		if cmp != gen {
			return fmt.Errorf("%v FAILED at %s \n Generated:%#v\n  Original:%#v", req, k, gen, cmp)
		}
	}
	return nil
}
