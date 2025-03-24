package handler

import (
	"api-gateway/config"
	"api-gateway/internal/proxy"
	"time"

	"github.com/gin-gonic/gin"
)

func TransactionHandler() gin.HandlerFunc {
	cfg := config.GetConfig()
	serviceURL := cfg.Services.Transaction.URL
	timeout := cfg.Services.Transaction.Timeout

	if timeout == 0 {
		timeout = 10 * time.Second
	}

	return proxy.ServiceForwarder(serviceURL, timeout)
}
