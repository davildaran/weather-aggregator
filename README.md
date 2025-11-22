# About

Weather aggregator service.

Provides a REST API for weather data requests.

## Usage

1. Build locally, `make build-bin`
1. Run locally, `./weatheragg`
1. Make a request for a point geometry in the United States. E.g. Central Park in Manhattan, NYk
```
  curl -X GET "http://localhost:8080/weather/point?longitude=-73.973192&latitude=40.772596"
```

## Implementation

Proxy calls to free [National Weather Service API](https://www.weather.gov/documentation/services-web-api) when applicable.

## Roadmap

- [ ] alerts API
  - [ ] subscription and opt-in notifications
- [ ] caching
- [ ] rate-limiting
- [ ] local data
- [ ] incorporate IoT sensors and devices