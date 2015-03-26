// Package ucloud provides ...
package ucloud

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	URL "net/url"
	_ "reflect"
	"sort"
)

type UResponse interface {
	OK() bool
	Msg() string
}

type BaseResponse struct {
	RetCode int
	Action  string `json:",omitempty"`
	Message string `json:",omitempty"`
}

func (b *BaseResponse) OK() bool {
	fmt.Println(b.RetCode)
	return (b.RetCode == 0)
}
func (b *BaseResponse) Msg() string {
	return b.Message
}

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

func (u *UcloudApiClient) verify_ac(encoded_params string) []byte {

	h := sha1.New()
	h.Write([]byte(encoded_params))
	return h.Sum(nil)
}

func (u *UcloudApiClient) RawGet(url string, params map[string]string) (*http.Response, error) {

	// Copy params
	_params := make(map[string]string, len(params)+1)
	for k, v := range params {
		_params[k] = v
	}
	_params["PublicKey"] = u.publicKey

	// URL Encode will sort Keys :)
	data := URL.Values{}
	keys := make([]string, 0)
	for k, _ := range _params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	s := ""
	for _, k := range keys {
		s += k + _params[k]
		data.Set(k, _params[k])
	}
	s += u.privateKey
	sig := fmt.Sprintf("%x", u.verify_ac(s))

	data.Set("Signature", sig)
	uri := u.baseURL + url + "?" + data.Encode()
	return u.conn.Get(uri)
}

func (u *UcloudApiClient) Get(params map[string]string, rsp interface{}) error {

	r, err := u.RawGet("/", params)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("%s", body)
	json.Unmarshal(body, rsp)
	return nil
}
