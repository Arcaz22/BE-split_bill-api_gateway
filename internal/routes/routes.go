package routes

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"api-gateway/pkg/response"
)

func forwardAuthService(c *gin.Context) {
	authServiceURL := "http://localhost:3000/v1"

	targetURL := authServiceURL + c.Request.URL.Path

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Error reading request body")
	}

	req, err := http.NewRequest(c.Request.Method, targetURL, bytes.NewBuffer(body))
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Error creating request")
        return
	}

	req.Header = c.Request.Header

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		response.Error(c, http.StatusServiceUnavailable, "Error forwarding request")
        return
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
        response.Error(c, http.StatusInternalServerError, "Error reading response")
        return
	}

	for k, v := range resp.Header {
		c.Writer.Header()[k] = v
	}

	c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), respBody)
}

func SetupRoutes(r *gin.Engine) {
	auth := r.Group("/auth")
	{
		auth.POST("/signup", forwardAuthService)
		auth.POST("/signin", forwardAuthService)
		auth.POST("/refresh-token", forwardAuthService)
		auth.POST("/logout", forwardAuthService)
		auth.POST("/google", forwardAuthService)
		auth.GET("/google/callback", forwardAuthService)
	}

	user := r.Group("/user")
	{
		user.GET("/current-user", forwardAuthService)
		user.GET("/get-all-user", forwardAuthService)
		user.POST("/add-profile", forwardAuthService)
		user.POST("/add-avatar-profile", forwardAuthService)
		user.PUT("/update-profile", forwardAuthService)
		user.PUT("/update-avatar-profile", forwardAuthService)
	}

	role := r.Group("/role")
	{
		role.GET("/get-all-role", forwardAuthService)
		role.POST("/create-role", forwardAuthService)
		role.POST("/add-user-role", forwardAuthService)
	}
}
