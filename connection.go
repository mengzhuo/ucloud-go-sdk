// Package ucloud provides ...
package ucloud

import (
	"bytes"
	"crypto/sha1"
	"encoding/json"
	"net/http"
	_ "reflect"
)

type UcloudApiClient struct {
	baseURL    string
	publicKey  string
	privateKey string
	regionId   string
	zoneId     string
	conn       *http.Client
}

func NewUcloudApiClient(baseURL, publicKey, privateKey, regionId, zoneId string) *UcloudApiClient {

	conn := &http.Client{}
	return &UcloudApiClient{baseURL, publicKey, privateKey, regionId, zoneId, conn}
}

func (u *UcloudApiClient) verify_ac(params interface{}) ([]byte, error) {

	var buf bytes.Buffer
	b, err := json.Marshal(params)
	if err != nil {
		return []byte{}, err
	}
	buf.Write(b)
	buf.WriteString(u.privateKey)

	h := sha1.New()

	return h.Sum(buf.Bytes()), nil
}

func (u *UcloudApiClient) MakeReqParams(origin_req []byte) []byte {

	return []byte{}

}

func (u *UcloudApiClient) Get(url string, params interface{}) {

}

func (u *UcloudApiClient) Post(url string, params interface{}) {
}

func (u *UcloudApiClient) Put(url string, params interface{}) {

}
func (u *UcloudApiClient) Delete(url string, params interface{}) {

}
