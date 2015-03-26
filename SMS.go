// Package main provides ...
package ucloud

import (
	"fmt"
)

type SendSmsResponse struct {
	*BaseResponse
}

func (c *UcloudApiClient) SendSms(content string, phones []string) (*SendSmsResponse, error) {

	params := make(map[string]string)
	params["Action"] = "SendSms"
	params["Content"] = content
	for i, p := range phones {
		params[fmt.Sprintf("Phone.%d", i)] = p
	}
	rsp := &SendSmsResponse{}
	err := c.Get(params, rsp)
	return rsp, err
}
