package api

import (
	"fmt"
	"os"
	"strconv"
)

const (
	NationalWeatherServiceApiHost = "https://api.weather.gov"
	PointsResource                = "points"
	GridpointsResource            = "gridpoints"
)

type NWSClient struct {
	UserAgent          string
	RateLimitPerMinute int64
}

func (c *NWSClient) Init() {
	c.UserAgent = fmt.Sprintf("(%s, %s)", os.Getenv("WEBSITE_CONTACT"), os.Getenv("EMAIL_CONTACT"))
	rateLimitPerMinute, ok := os.LookupEnv("RATE_LIMIT_PER_MINUTE")
	var rateInt int
	var err error
	if ok {
		rateInt, err = strconv.Atoi(rateLimitPerMinute)
		if err != nil {
			rateInt = 5
		}
	} else {
		rateInt = 5
	}
	c.RateLimitPerMinute = int64(rateInt)
}
