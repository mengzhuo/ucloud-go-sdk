package ucloud

import (
	"fmt"
	"testing"
)

// ---------------- TestGetMetric ------------------
func TestGetMetric(t *testing.T) {
	fmt.Println("UMon....")
	r := &GetMetric{ResourceType: "ulb",
		ResourceId: "ulb-kix4tp",
		Region:     "1",
		TimeRange:  300,
		MetricName: []string{"NetworkOut", "TotalNetworkOut"},
	}
	cmp := `https://api.ucloud.cn/?Action=GetMetric&Region=1&MetricName.0=NetworkOut&MetricName.1=TotalNetworkOut&ResourceType=ulb&ResourceId=ulb-kix4tp&TimeRange=300`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}
