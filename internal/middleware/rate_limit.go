package middleware

import (
    "github.com/gin-gonic/gin"
    "golang.org/x/time/rate"
    "time"
)

var limiter = rate.NewLimiter(rate.Every(time.Second), 100)

func RateLimit() gin.HandlerFunc {
    return func(c *gin.Context) {
        if !limiter.Allow() {
            c.JSON(429, gin.H{"error": "Too many requests"})
            c.Abort()
            return
        }
        c.Next()
    }
}
