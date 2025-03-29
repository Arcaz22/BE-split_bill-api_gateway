package middleware

import (
	"api-gateway/pkg/constant"
	"api-gateway/pkg/response"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sony/gobreaker"
)

var (
	circuitBreakers     = make(map[string]*gobreaker.CircuitBreaker)
	circuitBreakerMutex = &sync.RWMutex{}
)

func InitCircuitBreaker() {
	services := []string{
		constant.ServiceAuth,
		constant.ServiceTransaction,
		constant.ServiceNotification,
	}

	for _, service := range services {
		circuitBreakers[service] = gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:        service,
			MaxRequests: constant.CircuitBreakerMaxRequests,
			Interval:    time.Duration(constant.CircuitBreakerInterval) * time.Second,
			Timeout:     time.Duration(constant.CircuitBreakerTimeout) * time.Second,
			ReadyToTrip: func(counts gobreaker.Counts) bool {
				failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
				return counts.Requests >= constant.CircuitBreakerThreshold && failureRatio >= 0.6
			},
			OnStateChange: func(name string, from gobreaker.State, to gobreaker.State) {
				fmt.Printf("Circuit breaker %s state changed from %v to %v\n", name, from, to)
			},
		})
	}
}

// getCircuitBreaker returns the circuit breaker for the given service
func getCircuitBreaker(service string) *gobreaker.CircuitBreaker {
	circuitBreakerMutex.RLock()
	defer circuitBreakerMutex.RUnlock()

	cb, exists := circuitBreakers[service]
	if !exists {
		return circuitBreakers[constant.ServiceAuth]
	}
	return cb
}

func CircuitBreaker() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Determine which service this request is for based on path
		service := determineService(c.Request.URL.Path)

		// Get appropriate circuit breaker
		cb := getCircuitBreaker(service)

		// Execute request through circuit breaker
		_, err := cb.Execute(func() (interface{}, error) {
			// Store original handlers so we can restore them
			c.Next()

			// Check if the response was successful
			if c.Writer.Status() >= 500 {
				return nil, fmt.Errorf("service error: %d", c.Writer.Status())
			}

			return nil, nil
		})

		// If circuit breaker is open or half-open and fails
		if err != nil {
			if cb.State() == gobreaker.StateOpen {
				// Circuit is open, return error immediately
				response.Error(c, http.StatusServiceUnavailable, constant.ErrCodeServiceUnavailable, "Service temporarily unavailable")
				c.Abort()
				return
			}
		}
	}
}

// determineService maps a request path to a service name
func determineService(path string) string {
	if len(path) <= 1 {
		return constant.ServiceAuth
	}

	// Extract first segment of path
	var segment string
	for i := 1; i < len(path); i++ {
		if path[i] == '/' {
			segment = path[1:i]
			break
		}
		if i == len(path)-1 {
			segment = path[1:]
		}
	}

	switch segment {
	case "auth":
		return constant.ServiceAuth
	case "transaction":
		return constant.ServiceTransaction
	case "notifications":
        return constant.ServiceNotification
	default:
		return constant.ServiceAuth
	}
}
