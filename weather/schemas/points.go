package schemas

import "fmt"

func (n NWSApiPointsResponse) Validate() error {
	if n.Properties.GridID == "" {
		return fmt.Errorf("unexpected empty string for gridId")
	}

	if n.Properties.GridX <= 0 {
		return fmt.Errorf("unexpected value for gridX")
	}

	if n.Properties.GridY <= 0 {
		return fmt.Errorf("unexpected value for gridY")
	}

	return nil
}

type NWSApiPointsResponse struct {
	Context    []interface{}    `json:"@context"`
	ID         string           `json:"id"`
	Type       string           `json:"type"`
	Geometry   Geometry         `json:"geometry"`
	Properties PointsProperties `json:"properties"`
}

type PointsProperties struct {
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
