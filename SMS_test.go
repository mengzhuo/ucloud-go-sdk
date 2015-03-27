// Package ucloud provides ...
package ucloud

import (
	"os"
	"testing"
)

func TestSendSms(t *testing.T) {

	number := os.Getenv("UCLOUD_SENDSMS_PHONE")
	if number == "" {
		t.Skip("UCLOUD_SENDSMS_PHONE not set")
	}
	rsp, err := u.Do(&SendSms{"Hello world", []string{number}})

	if err != nil || !rsp.OK() {
		t.Fatal(err, rsp)
	}
}
