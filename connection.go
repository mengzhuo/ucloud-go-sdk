// Package ucloud provides ...
package ucloud

import (
	"bytes"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	URL "net/url"
	"sort"
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

func (u *UcloudApiClient) verify_ac(params map[string]string) []byte {

	var buf bytes.Buffer
	keys := make([]string, len(params))
	i := 0
	for k, _ := range params {
		keys[i] = k
		i++
	}
	sort.Strings(keys)

	for _, k := range keys {
		buf.WriteString(k)
		buf.WriteString(params[k])
	}
	buf.WriteString(u.privateKey)

	h := sha1.New()
	h.Write(buf.Bytes())
	return h.Sum(nil)
}

func (u *UcloudApiClient) RawGet(url string, params map[string]string) (*http.Response, error) {

	_params := make(map[string]string, len(params)+1)
	for k, v := range params {
		_params[k] = v
	}
	_params["PublicKey"] = u.publicKey

	data := URL.Values{}
	for k, v := range _params {
		data.Set(k, v)
	}

	data.Set("Signature", fmt.Sprintf("%x", u.verify_ac(_params)))
	return u.conn.Get(u.baseURL + url + "?" + data.Encode())
}

func (u *UcloudApiClient) Get(url string, params map[string]string) map[string]*json.RawMessage {

	r, err := u.RawGet(url, params)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)

	var rsp map[string]*json.RawMessage
	json.Unmarshal(body, &rsp)
	return rsp
}
