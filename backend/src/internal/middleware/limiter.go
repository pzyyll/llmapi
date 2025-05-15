package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	"github.com/ulule/limiter/v3/drivers/store/memory"
)

func IpLimiterMiddleware() gin.HandlerFunc {
	rate := limiter.Rate{
		Limit:  10,
		Period: 1 * time.Second,
	}

	store := memory.NewStore()

	limiterInstance := limiter.New(store, rate)

	return mgin.NewMiddleware(limiterInstance)
}
