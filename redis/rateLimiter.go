package cache

import (
	"context"
	"fmt"
	"log/slog"
	"time"
	"weather-aggregator/weather/api"

	"github.com/redis/go-redis/v9"
)

// There is no API key currently for the free to use National Weather Service API.
// They mention it will use an API key in the future.
// Instead authentication currently requires a User-Agent key-value in the header.
// The User-Agent header value is used apart of the Redis key.
func NWSSetRateLimit(ctx context.Context, rdb *redis.Client, nwsClient *api.NWSClient, logger *slog.Logger) string {
	key := fmt.Sprintf("nwsApiRate:%s:%d", nwsClient.UserAgent, time.Now().Minute())
	logger.DebugContext(ctx, "setting rate limit key", "key", key)
	rdb.HSet(ctx, fmt.Sprintf("nwsApiRate:%s:%d", nwsClient.UserAgent, time.Now().Minute()))
	return key
}

func NWSIncrRateLimit(ctx context.Context, rdb *redis.Client, key string, nwsClient *api.NWSClient, logger *slog.Logger) error {
	count, err := rdb.Incr(ctx, key).Result()
	if err != nil {
		logger.ErrorContext(ctx, "internal redis error incrementing rate limit", "error", err, "key", key)
		return fmt.Errorf("internal error with National Weather Service rate limiter")
	}
	if count == 1 {
		rdb.Expire(ctx, key, time.Minute)
	}
	if count > nwsClient.RateLimitPerMinute {
		logger.ErrorContext(ctx, "rate limit exceeded for National Weather Service", "rate_limit_per_minute", nwsClient.RateLimitPerMinute)
		return fmt.Errorf("error rate limit exceeded for National Weather Service, please wait and try again later")
	}
	return nil
}
