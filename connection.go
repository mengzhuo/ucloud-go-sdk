// Package ucloud provides ...
package ucloud

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	URL "net/url"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

type UResponse interface {
	OK() bool
	Msg() string
	Data() interface{}
	SetCode(int)
}

type URequest interface {
	// This should return new instance for json Unmarshal
	R() UResponse
}

type BaseResponse struct {
	RetCode int
	Action  string `json:",omitempty"`
	Message string `json:",omitempty"`
}

func (b *BaseResponse) OK() bool {
	return (b.RetCode == 0) // XXX ReturnCode is string for abnormal
}
func (b *BaseResponse) SetCode(i int) {
	b.RetCode = i // This is a patch for Ucloud API return with string
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
	return u.conn.Get(u.MakeURI(url, params))
}

func (u *UcloudApiClient) MakeURI(url string, params map[string]string) string {
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
	//fmt.Println(uri)
	return uri
}

func (u *UcloudApiClient) Get(params map[string]string, rsp UResponse) error {

	rsp.SetCode(500) // Patch broken API
	r, err := u.RawGet("/", params)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	//fmt.Printf("%s", body)
	json.Unmarshal(body, &rsp)
	if !rsp.OK() {
		return fmt.Errorf("%s", body)
	}
	return nil
}

func (u *UcloudApiClient) MakeParams(request URequest) (params map[string]string, err error) {

	v := reflect.ValueOf(request)
	typ := reflect.TypeOf(request)

	if typ.Kind() == reflect.Ptr {
		v = v.Elem()
		typ = typ.Elem()
	}

	params = make(map[string]string, 0) // it's 0 for optional might skip

	for i := 0; i < typ.NumField(); i++ {

		name := typ.Field(i).Name
		tag := typ.Field(i).Tag.Get("ucloud")
		field := v.Field(i)

		switch field.Kind() {
		case reflect.Slice:
			if field.IsNil() {
				if tag == "optional" {
					continue
				} else {
					return params, fmt.Errorf("Lack of parameter:%s", name)
				}
			}
			for j := 0; j < field.Len(); j++ {
				// Must be string in Slice
				params[fmt.Sprintf("%s.%d", name, j)] = field.Index(j).String()
			}
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			num := field.Int()
			if num == -1 && tag == "optional" {
				continue
			}
			/*
				XXX How can we tell that if user JUST need 0 value?
					if num == 0 {
						if tag == "optional" {
							continue
						} else {
							return nil, fmt.Errorf("Lack of parameter:%s", name)
						}
					}*/
			params[name] = strconv.FormatInt(num, 10)
		case reflect.Float32, reflect.Float64:
			num := field.Float()
			if num == 0.0 && tag == "optional" {
				continue
			}
			params[name] = fmt.Sprintf("%v", num)
		case reflect.Bool:
			b := field.Bool()
			if !b && tag == "optional" {
				continue // XXX Boolean act like should be skip?
			}
			bt := "False"
			if b {
				bt = "True"
			}
			params[name] = bt
		case reflect.String:
			str := field.String()

			if str == "" {
				if tag == "optional" {
					continue
				} else {
					return nil, fmt.Errorf("Lack of parameter:%s", name)
				}
			}

			params[name] = str
		}
	}
	//XXX Set Action Name should be more elegent
	typ_name_list := strings.Split(typ.String(), ".")
	typ_name := typ_name_list[len(typ_name_list)-1]
	params["Action"] = typ_name
	return params, nil
}

func (u *UcloudApiClient) Do(request URequest) (UResponse, error) {

	rsp := request.R()
	params, err := u.MakeParams(request)
	if err != nil {
		return rsp, err
	}
	// OK, now we had params
	err = u.Get(params, rsp)
	return rsp, err
}
