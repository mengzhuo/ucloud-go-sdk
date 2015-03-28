// Package main provides ...
package ucloud

type SendSmsResponse struct {
	BaseResponse
}

type SendSms struct {
	Content string
	Phone   []string
}

func (r *SendSms) R() UResponse {
	return &SendSmsResponse{}
}

func (r *SendSmsResponse) Data() interface{} {
	return r.OK()
}
