package schemas

type Elevation struct {
	Value          float64 `json:"value"`
	MaxValue       int     `json:"maxValue"`
	MinValue       int     `json:"minValue"`
	UnitCode       string  `json:"unitCode"`
	QualityControl string  `json:"qualityControl"`
}
