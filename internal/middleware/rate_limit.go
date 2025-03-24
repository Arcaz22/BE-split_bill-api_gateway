package middleware

import (
	"api-gateway/pkg/constant"
	"api-gateway/pkg/response"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

var limiter = rate.NewLimiter(rate.Limit(constant.RateLimitPerMinute/60), constant.RateLimitBurst)

func RateLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !limiter.Allow() {
			response.Error(c, http.StatusTooManyRequests, constant.ErrCodeRateLimitExceeded, "Rate limit exceeded")
			c.Abort()
			return
		}

		// Calculate remaining tokens
		remaining := limiter.Tokens()
		remainingStr := fmt.Sprintf("%.0f", remaining)

		// Convert Unix timestamp to string
		resetTime := time.Now().Add(time.Minute).Unix()
		resetTimeStr := strconv.FormatInt(resetTime, 10)

		c.Header(constant.HeaderRateLimit, strconv.Itoa(constant.RateLimitPerMinute))
		c.Header(constant.HeaderRateRemaining, remainingStr)
		c.Header(constant.HeaderRateReset, resetTimeStr)

		c.Next()
	}
}
