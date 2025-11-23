# About

Weather aggregator service.

A REST API for weather data requests.

## Usage

### Standalone Go binary
1. Build locally, `make build-bin`
1. Run locally, `./weatheragg`

### Docker 
1. Build locally, `make build`
1. Run locally, `make up`


Then make a request for a point geometry in the United States. E.g. Central Park in Manhattan, NY
```
  curl -X GET "http://localhost:8080/weather/point?longitude=-73.973192&latitude=40.772596"
```

This returns a fair amount of (JSON) data, hourly weather forecast for a given location.

The original Point geometry ends up representing a 2.5 by 2.5 kilometer square.

If you are running locally you can find the output log file, e.g. `log/11-23-weather-log.json`.

If running in a container you can inspect the container filesystem with `docker exec -it <container_id> sh` and copy to the host filesystem with `docker cp <container_id>:/log/11-23-weather-log.json log`.

## Implementation

Proxy calls to free [National Weather Service API](https://www.weather.gov/documentation/services-web-api) when applicable.

## Roadmap

- [ ] alerts API
  - [ ] subscription and opt-in notifications
- [ ] caching
- [ ] rate-limiting
- [ ] local data
- [ ] incorporate IoT sensors and devices