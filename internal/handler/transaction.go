package handler

import (
	"api-gateway/config"
	"api-gateway/internal/proxy"
	"api-gateway/pkg/response"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	serviceURL string
	timeout    time.Duration
}

func NewTransactionHandler() *TransactionHandler {
	cfg := config.GetConfig()
	timeout := cfg.Services.Transaction.Timeout

	if timeout == 0 {
		timeout = 10 * time.Second
	}

	return &TransactionHandler{
        serviceURL: cfg.Services.Transaction.URL,
        timeout:    timeout,
    }
}

func (h *TransactionHandler) HealthCheck(c *gin.Context) {
    client := &http.Client{
        Timeout: h.timeout,
    }

    resp, err := client.Get(h.serviceURL + "/health")
    if err != nil {
        response.JSON(c, http.StatusServiceUnavailable, gin.H{
            "status":    "DOWN",
            "service":   "transaction-service",
            "error":     err.Error(),
            "timestamp": time.Now().Format(time.RFC3339),
        })
        return
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        response.JSON(c, http.StatusServiceUnavailable, gin.H{
            "status":    "DOWN",
            "service":   "transaction-service",
            "code":      resp.StatusCode,
            "timestamp": time.Now().Format(time.RFC3339),
        })
        return
    }

    response.JSON(c, http.StatusOK, gin.H{
        "status":    "UP",
        "service":   "transaction-service",
        "timestamp": time.Now().Format(time.RFC3339),
    })
}

func (h *TransactionHandler) Forward() gin.HandlerFunc {
    return proxy.ServiceForwarder(h.serviceURL, h.timeout)
}
