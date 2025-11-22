package schemas

type NWSApiPointResponse struct {
	Context    []interface{} `json:"@context"`
	ID         string        `json:"id"`
	Type       string        `json:"type"`
	Geometry   Geometry      `json:"geometry"`
	Properties Properties    `json:"properties"`
}

type Geometry struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
	BBox        []float64 `json:"bbox"`
}

type Properties struct {
	Context             []interface{}    `json:"@context"`
	Geometry            string           `json:"geometry"`
	ID                  string           `json:"@id"`
	Type                string           `json:"@type"`
	CWA                 string           `json:"cwa"`
	ForecastOffice      string           `json:"forecastOffice"`
	GridID              string           `json:"gridId"`
	GridX               int              `json:"gridX"`
	GridY               int              `json:"gridY"`
	Forecast            string           `json:"forecast"`
	ForecastHourly      string           `json:"forecastHourly"`
	ForecastGridData    string           `json:"forecastGridData"`
	ObservationStations string           `json:"observationStations"`
	RelativeLocation    RelativeLocation `json:"relativeLocation"`
	ForecastZone        string           `json:"forecastZone"`
	County              string           `json:"county"`
	FireWeatherZone     string           `json:"fireWeatherZone"`
	TimeZone            string           `json:"timeZone"`
	RadarStation        string           `json:"radarStation"`
}

type RelativeLocation struct {
	Context    []interface{}      `json:"@context"`
	ID         string             `json:"id"`
	Type       string             `json:"type"`
	Geometry   Geometry           `json:"geometry"`
	Properties RelativeProperties `json:"properties"`
}

type RelativeProperties struct {
	City     string   `json:"city"`
	State    string   `json:"state"`
	Distance Distance `json:"distance"`
	Bearing  Bearing  `json:"bearing"`
}

type Distance struct {
	Value          float64 `json:"value"`
	MaxValue       float64 `json:"maxValue"`
	MinValue       float64 `json:"minValue"`
	UnitCode       string  `json:"unitCode"`
	QualityControl string  `json:"qualityControl"`
}

type Bearing struct {
	Value          float64 `json:"value"`
	MaxValue       float64 `json:"maxValue"`
	MinValue       float64 `json:"minValue"`
	UnitCode       string  `json:"unitCode"`
	QualityControl string  `json:"qualityControl"`
}
