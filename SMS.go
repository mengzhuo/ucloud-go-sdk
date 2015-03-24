// Package main provides ...
package ucloud

import (
	"fmt"
)

func (c *UcloudApiClient) SendSms(content string, phones []string) {

	params := make(map[string]string)
	params["Action"] = "SendSms"
	params["Content"] = content
	for i, p := range phones {
		params[fmt.Sprintf("Phone.%d", i)] = p
	}
	r := c.Get(params)
	fmt.Println(r)
}
