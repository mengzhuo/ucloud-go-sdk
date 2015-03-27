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
}

type GetMetricResponse struct {
	BaseResponse
	DataSets map[string]*GetMetricItem
}

func (r *GetMetric) R() UResponse {
	return &GetMetricResponse{}
}
