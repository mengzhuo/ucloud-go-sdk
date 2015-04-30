package ucloud

import (
	"fmt"
	"testing"
)

// ---------------- TestBuyUcdnTraffic ------------------
func TestBuyUcdnTraffic(t *testing.T) {
	fmt.Println("UCDN......")
	r := &BuyUcdnTraffic{Areacode: "cn",
		Traffic: 10,
	}
	cmp := `https://api.ucloud.cn/?Action=BuyUcdnTraffic&Traffic=10&Areacode=cn`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestGetUcdnTraffic ------------------
func TestGetUcdnTraffic(t *testing.T) {
	r := &GetUcdnTraffic{}
	cmp := `https://api.ucloud.cn/?Action=GetUcdnTraffic`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestCreateUcdnDomain ------------------
func TestCreateUcdnDomain(t *testing.T) {
	r := &CreateUcdnDomain{
		Areacode: []string{"cn"},
		Domain:   "testcdn.ucloud.cn",
		SourceIp: []string{"1.1.1.1"},
		CdnType:  "web",
	}
	cmp := `https://api.ucloud.cn/?Action=CreateUcdnDomain&Areacode.0=cn&Domain=testcdn.ucloud.cn&CdnType=web&SourceIp.0=1.1.1.1`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestUpdateUcdnDomain ------------------
func TestUpdateUcdnDomain(t *testing.T) {
	r := &UpdateUcdnDomain{
		CacheFileType: []string{"zip"},
		DomainId:      "domain-0331qd",
		SourceIp:      []string{"1.1.1.1"},
		CacheTtl:      86400,
		Areacode:      []string{"cn"},
		NoCacheUrl:    []string{"http://static.ucloud.cn/test/*"},
		TestUrl:       "http://static.ucloud.cn/abc/abc.png",
		CacheUrl:      []string{"http://static.ucloud.cn/abc/*"},
	}
	cmp := `https://api.ucloud.cn/?Action=UpdateUcdnDomain&DomainId=domain-0331qd&SourceIp.0=1.1.1.1&TestUrl=http://static.ucloud.cn/abc/abc.png&Areacode.0=cn&CacheFileType.0=zip&CacheUrl.0=http://static.ucloud.cn/abc/*&CacheTtl=86400&NoCacheUrl.0=http://static.ucloud.cn/test/*`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestDescribeUcdnDomain ------------------
func TestDescribeUcdnDomain(t *testing.T) {
	r := &DescribeUcdnDomain{}
	cmp := `https://api.ucloud.cn/?Action=DescribeUcdnDomain`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestGetUcdnDomainBandwidth ------------------
func TestGetUcdnDomainBandwidth(t *testing.T) {
	r := &GetUcdnDomainBandwidth{
		Areacode:  []string{"cn"},
		EndTime:   1421166064,
		BeginTime: 1420992000,
		DomainId:  []string{"ucdn-iw4hun"},
	}
	cmp := `https://api.ucloud.cn/?Action=GetUcdnDomainBandwidth&DomainId.0=ucdn-iw4hun&BeginTime=1420992000&EndTime=1421166064&Areacode.0=cn`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestGetUcdnDomainTraffic ------------------
func TestGetUcdnDomainTraffic(t *testing.T) {
	r := &GetUcdnDomainTraffic{
		Areacode:  []string{"cn"},
		EndTime:   1421166064,
		BeginTime: 1420992000,
		DomainId:  []string{"ucdn-iw4hun"},
	}
	cmp := `https://api.ucloud.cn/?Action=GetUcdnDomainTraffic&DomainId.0=ucdn-iw4hun&BeginTime=1420992000&EndTime=1421166064&Areacode.0=cn`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestGetUcdnDomainLog ------------------
func TestGetUcdnDomainLog(t *testing.T) {
	r := &GetUcdnDomainLog{EndTime: 1421166064,
		BeginTime: 1420992000,
		DomainId:  []string{"ucdn-iw4hun"},
	}
	cmp := `https://api.ucloud.cn/?Action=GetUcdnDomainLog&DomainId.0=ucdn-iw4hun&BeginTime=1420992000&EndTime=1421166064`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestRefreshUcdnDomainCache ------------------
func TestRefreshUcdnDomainCache(t *testing.T) {
	r := &RefreshUcdnDomainCache{UrlList: []string{"http://ucloud.cn/images/"},
		DomainId: "ucdn-0331qd",
		Type:     "1",
	}
	cmp := `https://api.ucloud.cn/?Action=RefreshUcdnDomainCache&DomainId=ucdn-0331qd&Type=1&UrlList.0=http://ucloud.cn/images/`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestDescribeRefreshCacheTask ------------------
func TestDescribeRefreshCacheTask(t *testing.T) {
	r := &DescribeRefreshCacheTask{EndTime: 1398873600,
		DomainId:  "ucdn-0331qd",
		BeginTime: 1398700800,
		Status:    "success",
	}
	cmp := `https://api.ucloud.cn/?Action=DescribeRefreshCacheTask&DomainId=ucdn-0331qd&BeginTime=1398700800&EndTime=1398873600&Status=success`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestPrefetchDomainCache ------------------
func TestPrefetchDomainCache(t *testing.T) {
	r := &PrefetchDomainCache{UrlList: []string{"http://ucloud.cn/images/"},
		DomainId: "ucdn-iw4hun",
	}
	cmp := `https://api.ucloud.cn/?Action=PrefetchDomainCache&DomainId=ucdn-iw4hun&UrlList.0=http://ucloud.cn/images/`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestDescribePrefetchCacheTask ------------------
func TestDescribePrefetchCacheTask(t *testing.T) {
	r := &DescribePrefetchCacheTask{EndTime: 1398873600,
		DomainId:  "ucdn-iw4hun",
		BeginTime: 1398700800,
		Status:    "success",
	}
	cmp := `https://api.ucloud.cn/?Action=DescribePrefetchCacheTask&DomainId=ucdn-iw4hun&BeginTime=1398700800&EndTime=1398873600&Status=success`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestUpdateUcdnDomainStatus ------------------
func TestUpdateUcdnDomainStatus(t *testing.T) {
	r := &UpdateUcdnDomainStatus{Status: "disable",
		DomainId: "domain-0331qd",
	}
	cmp := `https://api.ucloud.cn/?Action=UpdateUcdnDomainStatus&DomainId=domain-0331qd&Status=disable`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestGetUcdnDomainPrefetchEnable ------------------
func TestGetUcdnDomainPrefetchEnable(t *testing.T) {
	r := &GetUcdnDomainPrefetchEnable{DomainId: "domain-0331qd"}
	cmp := `https://api.ucloud.cn/?Action=GetUcdnDomainPrefetchEnable&DomainId=domain-0331qd`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}
