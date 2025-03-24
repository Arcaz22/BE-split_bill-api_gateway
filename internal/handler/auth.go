package handler

import (
	"api-gateway/config"
	"api-gateway/internal/proxy"
	"time"

	"github.com/gin-gonic/gin"
)

func AuthHandler() gin.HandlerFunc {
	cfg := config.GetConfig()
	serviceURL := cfg.Services.Auth.URL
	timeout := cfg.Services.Auth.Timeout

	if timeout == 0 {
		timeout = 10 * time.Second
	}

	return proxy.ServiceForwarder(serviceURL, timeout)
}
