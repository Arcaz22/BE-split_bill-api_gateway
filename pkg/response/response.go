package response

import (
	"api-gateway/pkg/constant"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthResponse represents the health check response
// @Description Health check response format
type HealthResponse struct {
    Status    string `json:"status" example:"OK"`
    Timestamp int64  `json:"timestamp" example:"1648243759"`
    Message   string `json:"message" example:"API Gateway is running"`
}

// ErrorResponse represents the error response structure
// @Description Error response format
type ErrorResponse struct {
    Code    string      `json:"code" example:"BAD_REQUEST" description:"Error code"`
    Message string      `json:"message" example:"Invalid input parameters" description:"Error message"`
    Details interface{} `json:"details,omitempty" description:"Additional error details"`
}

type SuccessResponse struct {
    Data interface{} `json:"data" description:"Response data"`
    Meta interface{} `json:"meta,omitempty" description:"Metadata information"`
}

func JSON(c *gin.Context, status int, data interface{}) {
	c.Header(constant.HeaderContentType, constant.ContentTypeJSON)
	c.JSON(status, data)
}

func Success(c *gin.Context, status int, data interface{}, meta interface{}) {
	response := SuccessResponse{
		Data: data,
		Meta: meta,
	}

	JSON(c, status, response)
}

func Error(c *gin.Context, status int, code string, message string, details ...interface{}) {
	var detailsData interface{}
	if len(details) > 0 {
		detailsData = details[0]
	}

	response := ErrorResponse{
		Code:    code,
		Message: message,
		Details: detailsData,
	}

	JSON(c, status, response)
}

// Helper functions for common errors
func BadRequest(c *gin.Context, message string, details ...interface{}) {
	Error(c, http.StatusBadRequest, constant.ErrCodeBadRequest, message, details...)
}

func Unauthorized(c *gin.Context, message string, details ...interface{}) {
	Error(c, http.StatusUnauthorized, constant.ErrCodeUnauthorized, message, details...)
}

func Forbidden(c *gin.Context, message string, details ...interface{}) {
	Error(c, http.StatusForbidden, constant.ErrCodeForbidden, message, details...)
}

func NotFound(c *gin.Context, message string, details ...interface{}) {
	Error(c, http.StatusNotFound, constant.ErrCodeNotFound, message, details...)
}

func InternalError(c *gin.Context, message string, details ...interface{}) {
	Error(c, http.StatusInternalServerError, constant.ErrCodeInternalError, message, details...)
}

func ServiceUnavailable(c *gin.Context, message string, details ...interface{}) {
	Error(c, http.StatusServiceUnavailable, constant.ErrCodeServiceUnavailable, message, details...)
}
