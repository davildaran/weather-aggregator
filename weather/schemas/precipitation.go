package schemas

type ProbabilityOfPrecipitation struct {
	Value          int    `json:"value"`
	MaxValue       int    `json:"maxValue"`
	MinValue       int    `json:"minValue"`
	UnitCode       string `json:"unitCode"`
	QualityControl string `json:"qualityControl"`
}

type Dewpoint struct {
	Value          float64 `json:"value"`
	MaxValue       int     `json:"maxValue"`
	MinValue       int     `json:"minValue"`
	UnitCode       string  `json:"unitCode"`
	QualityControl string  `json:"qualityControl"`
}

type RelativeHumidity struct {
	Value          int    `json:"value"`
	MaxValue       int    `json:"maxValue"`
	MinValue       int    `json:"minValue"`
	UnitCode       string `json:"unitCode"`
	QualityControl string `json:"qualityControl"`
}
