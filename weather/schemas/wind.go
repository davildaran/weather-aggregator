package schemas

type WindSpeed struct {
	Value          int    `json:"value"`
	MaxValue       int    `json:"maxValue"`
	MinValue       int    `json:"minValue"`
	UnitCode       string `json:"unitCode"`
	QualityControl string `json:"qualityControl"`
}

type WindGust struct {
	Value          int    `json:"value"`
	MaxValue       int    `json:"maxValue"`
	MinValue       int    `json:"minValue"`
	UnitCode       string `json:"unitCode"`
	QualityControl string `json:"qualityControl"`
}
