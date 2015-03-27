package ucloud

type GetMetric struct {
	Region       string
	MetricName   []string
	ResourceId   string
	ResourceType string
	TimeRange    int `uapi:"optional"`
	BeginTime    int `uapi:"optional"`
	EndTime      int `uapi:"optional"`
}

type GetMetricItem struct {
	Timestamp int
	Value     float64
}

type GetMetricResponse struct {
	*BaseResponse
	DataSets map[string]*GetMetricItem
}

func (r *GetMetric) R() *GetMetricResponse {
	return &GetMetricResponse{}
}

/*

rsp, err := u.Do(&GetMetric{})
if err != nil {
	//bla bla
}
*/
