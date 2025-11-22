package weather

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"weather-aggregator/weather/api"
	"weather-aggregator/weather/schemas"
)

// First, obtain the grid forecast(s) for a point location.
// Then, obtain the specified grid forecase. Hourly or 12 hour periods. Raw grid data is also available.
// https://www.weather.gov/documentation/services-web-api
func WeatherServerHandler(ctx context.Context, logger *slog.Logger) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/weather/point", func(w http.ResponseWriter, r *http.Request) {
		// handle GET request for weather by lat, long
		// validate query parameters
		values := r.URL.Query()
		latitude := values.Get("latitude")
		lat64, err := strconv.ParseFloat(latitude, 64)
		if err != nil {
			logger.ErrorContext(ctx, "invalid latitude in weather point GET request", "error", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		longitude := values.Get("longitude")
		long64, err := strconv.ParseFloat(longitude, 64)
		if err != nil {
			logger.ErrorContext(ctx, "invalid longitude in weather point GET request", "error", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// formulate request to National Weather Service API
		// TODOs
		// first look in cache, then db, then make external network request
		// rate limiting
		// timeout for fetching data
		req, err := http.NewRequestWithContext(
			ctx,
			http.MethodGet,
			fmt.Sprintf("%s/%s/%f,%f",
				api.NationalWeatherServiceApiHost,
				api.PointsResource,
				lat64,
				long64,
			),
			nil,
		)
		if err != nil {
			logger.ErrorContext(ctx, "malformed internal GET request", "error", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		req.Header.Set("Content-Type", "application/geo+json")
		req.Header.Add("User-Agent", fmt.Sprintf("(%s, %s)", os.Getenv("WEBSITE_CONTACT"), os.Getenv("EMAIL_CONTACT")))
		// TODO add request to queue and implment producer consumer pattern
		// TODO prometheus metric summary or histogram capturing latency
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			logger.ErrorContext(ctx, "error sending GET request for weather points data to National Weather Service", "error", err)
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			logger.ErrorContext(ctx, "error reading response from National Weather Service for GET points request", "error", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		var result schemas.NWSApiPointResponse
		err = json.Unmarshal(body, &result)
		if err != nil {
			logger.ErrorContext(ctx, "error reading response from National Weather Service for GET points request", "error", err)
			logger.InfoContext(ctx, "raw response", "response", string(body))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		// optional TODO, custom validate() method for NWSApiPointResponse

		logger.Debug("expected national weather service api point response")
		// TODO save off relevant response data to cache and eventual consistency to db

		// TODO use forecast url to make subsequent request for forecast data

		w.WriteHeader(http.StatusOK)
		w.Write(body)
	})
	return mux
}
