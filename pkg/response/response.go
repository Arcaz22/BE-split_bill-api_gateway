package response

import "github.com/gin-gonic/gin"

type Response struct {
    Status  int         `json:"status"`
    Message string      `json:"message"`
    Data    interface{} `json:"data,omitempty"`
    Error   string      `json:"error,omitempty"`
}

func JSON(c *gin.Context, statusCode int, data interface{}) {
    c.JSON(statusCode, Response{
        Status:  statusCode,
        Message: "success",
        Data:    data,
    })
}

func Error(c *gin.Context, statusCode int, message string) {
    c.JSON(statusCode, Response{
        Status:  statusCode,
        Message: "error",
        Error:   message,
    })
}
