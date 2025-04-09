package handler

import (
	"api-gateway/config"
	"api-gateway/internal/proxy"
	"api-gateway/pkg/response"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	serviceURL string
	timeout    time.Duration
}

func NewAuthHandler() *AuthHandler {
    cfg := config.GetConfig()
	timeout := cfg.Services.Auth.Timeout

	if timeout == 0 {
		timeout = 10 * time.Second
	}

	return &AuthHandler{
        serviceURL: cfg.Services.Auth.URL,
        timeout:    timeout,
    }
}

func (h *AuthHandler) HealthCheck(c *gin.Context) {
    client := &http.Client{
		Timeout: h.timeout,
	}
	resp, err := client.Get(h.serviceURL + "/health")
	if err != nil {
		response.JSON(c, http.StatusServiceUnavailable, gin.H{
			"status":    "DOWN",
			"service":   "auth-service",
			"error":     err.Error(),
			"timestamp": time.Now().Format(time.RFC3339),
		})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		response.JSON(c, http.StatusServiceUnavailable, gin.H{
			"status":    "DOWN",
			"service":   "auth-service",
			"code":      resp.StatusCode,
			"timestamp": time.Now().Format(time.RFC3339),
		})
		return
	}
	response.JSON(c, http.StatusOK, gin.H{
        "status":    "UP",
        "service":   "auth-service",
        "timestamp": time.Now().Format(time.RFC3339),
    })
}

func (h *AuthHandler) Forward() gin.HandlerFunc {
    return proxy.ServiceForwarder(h.serviceURL, h.timeout)
}
