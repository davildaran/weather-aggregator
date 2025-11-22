# About

Weather aggregator service.

Provides a REST API for weather data requests.

Examples:
1. GET provides current weather data based on location and/or time.

## Implementation

Proxy calls to free [National Weather Service API](https://www.weather.gov/documentation/services-web-api) when applicable.

Uses caching, rate-limiting, local data.

Future TODO: use local IoT sensors and devices.
