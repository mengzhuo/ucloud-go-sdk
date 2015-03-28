package ucloud

type GetMetric struct {
	Region       string
	MetricName   []string
	ResourceId   string
	ResourceType string
	TimeRange    int `ucloud:"optional"`
	BeginTime    int `ucloud:"optional"`
	EndTime      int `ucloud:"optional"`
}

type GetMetricItem struct {
	Timestamp int
	Value     float64
	IP        string `json:",omitempty"`
}

type GetMetricResponse struct {
	BaseResponse
	DataSets map[string]*GetMetricItem
}

func (r *GetMetricResponse) Data() interface{} {
	return r.DataSets
}

func (r *GetMetric) R() UResponse {
	return &GetMetricResponse{}
}
