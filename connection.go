// Package ucloud provides ...
package ucloud

import (
	"bytes"
	"crypto/sha1"
	_ "encoding/json"
	"net/http"
	URL "net/url"
	"sort"

	"github.com/mengzhuo/gopystr"
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

func (u *UcloudApiClient) verify_ac(params map[string]interface{}) []byte {

	var buf bytes.Buffer

	keys := make([]string, len(params))
	i := 0
	for k, _ := range params {
		keys[i] = k
	}
	sort.Strings(keys)

	for _, k := range keys {
		buf.WriteString(k)
		buf.WriteString(gopystr.Str(params[k]))
	}
	buf.WriteString(u.privateKey)

	h := sha1.New()

	return h.Sum(buf.Bytes())
}

func (u *UcloudApiClient) Get(url string, params map[string]interface{}) (*http.Response, error) {

	data := URL.Values{}
	for k, v := range params {
		data.Set(k, gopystr.Str(v))
	}
	data.Set("Signature", string(u.verify_ac(params)))
	r, _ := http.NewRequest("GET", url, bytes.NewBufferString(data.Encode()))
	return u.conn.Do(r)
}
