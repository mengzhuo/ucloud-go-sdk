package ucloud

import (
	"fmt"
	"testing"
)

func TestGetMetric(t *testing.T) {

	g := GetMetric{Region: "cn-east-01",
		MetricName:   []string{"NetworkOut", "TotalNetworkOut"},
		ResourceId:   "uhost-ahdvfh",
		ResourceType: "uhost"}
	rsp, err := u.Do(&g)
	if err != nil || !rsp.OK() {
		t.Fatal(err, rsp)
	}
	data := rsp.Data().(map[string]*GetMetricItem)
	fmt.Println(data["NetworkOut"])
}
