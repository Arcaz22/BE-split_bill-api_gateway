package proxy

import (
	"api-gateway/pkg/constant"
	"api-gateway/pkg/response"
	"api-gateway/pkg/utils"
	"bytes"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func ServiceForwarder(serviceURL string, timeout time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		targetURL := serviceURL + c.Request.URL.Path

		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			response.InternalError(c, "Error reading request body")
			return
		}

		req, err := http.NewRequest(c.Request.Method, targetURL, bytes.NewBuffer(body))
		if err != nil {
			response.InternalError(c, "Error creating request")
			return
		}

		req.Header = c.Request.Header

		requestID := utils.ExtractRequestID(req.Header)
		if requestID == "" {
			requestID = "req-" + time.Now().Format("20060102150405")
			utils.SetRequestID(req.Header, requestID)
		}

		client := utils.CreateClientWithTimeout(timeout)

		resp, err := client.Do(req)
		if err != nil {
			response.ServiceUnavailable(c, "Service unavailable")
			return
		}
		defer resp.Body.Close()

		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			response.InternalError(c, "Error reading response")
			return
		}

		for k, v := range resp.Header {
			c.Writer.Header()[k] = v
		}

		c.Writer.Header().Set(constant.HeaderXRequestID, requestID)

		c.Data(resp.StatusCode, utils.GetContentType(resp.Header), respBody)
	}
}
