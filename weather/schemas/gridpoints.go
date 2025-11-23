package schemas

import "time"

func (g NWSApiGridpointsResponse) Validate() error {
	// add custom validation

	return nil
}

type NWSApiGridpointsResponse struct {
	Context    []interface{}        `json:"@context"`
	Type       string               `json:"type"`
	Geometry   PolygonGeometry      `json:"geometry"`
	Properties GridpointsProperties `json:"properties"`
}

type GridpointsProperties struct {
	Units             string              `json:"units"`
	ForecastGenerator string              `json:"forecastGenerator"`
	GeneratedAt       time.Time           `json:"generatedAt"`
	UpdateTime        time.Time           `json:"updateTime"`
	ValidTimes        string              `json:"validTimes"` // e.g. "2007-03-01T13:00:00Z/2008-05-11T15:30:00Z"
	Elevation         Elevation           `json:"elevation"`
	Periods           []GridpointsPeriods `json:"periods"`
}

type GridpointsPeriods struct {
	Number                     int                        `json:"number"`
	Name                       string                     `json:"name"`
	StartTime                  time.Time                  `json:"startTime"`
	EndTime                    time.Time                  `json:"endTime"`
	IsDaytime                  bool                       `json:"isDaytime"`
	Temperature                int                        `json:"temperature"`
	TemperatureUnit            string                     `json:"temperatureUnit"`
	TemperatureTrend           string                     `json:"temperatureTrend"`
	ProbabilityOfPrecipitation ProbabilityOfPrecipitation `json:"probabilityOfPrecipitation"`
	Dewpoint                   Dewpoint                   `json:"dewpoint"`
	RelativeHumidity           RelativeHumidity           `json:"relativeHumidity"`
	WindSpeed                  string                     `json:"windSpeed"`
	WindGust                   WindGust                   `json:"windGust"`
	WindDirection              string                     `json:"windDirection"`
	Icon                       string                     `json:"icon"`
	ShortForecast              string                     `json:"shortForecast"`
	DetailedForecast           string                     `json:"detailedForecast"`
}
