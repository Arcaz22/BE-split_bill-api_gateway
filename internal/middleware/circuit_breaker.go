package middleware

import (
    "github.com/gin-gonic/gin"
    "github.com/sony/gobreaker"
    "time"
)

var cb *gobreaker.CircuitBreaker

func InitCircuitBreaker() {
    cb = gobreaker.NewCircuitBreaker(gobreaker.Settings{
        Name:        "auth-service",
        MaxRequests: 3,
        Interval:    10 * time.Second,
        Timeout:     60 * time.Second,
    })
}

func CircuitBreaker() gin.HandlerFunc {
    return func(c *gin.Context) {
        _, err := cb.Execute(func() (interface{}, error) {
            c.Next()
            return nil, nil
        })

        if err != nil {
            c.JSON(503, gin.H{"error": "Service temporarily unavailable"})
            c.Abort()
            return
        }
    }
}
