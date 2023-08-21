package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func jsonResponse(c *gin.Context, status int, message string, data interface{}) {
	c.JSON(status, gin.H{"status": status, "message": message, "data": data})
}

// Success https://httpstatuses.com/200.
// The request has succeeded.
func Success(c *gin.Context, message string, data interface{}) {
	jsonResponse(c, http.StatusOK, message, data)
}

// InternalServerError https://httpstatuses.com/500
func InternalServerError(c *gin.Context, message string, data interface{}) {
	jsonResponse(c, http.StatusInternalServerError, message, data)
}

// NotFound https://httpstatuses.com/404
// The origin server did not find a current representation for the target resource.
func NotFound(c *gin.Context, message string, data interface{}) {
	jsonResponse(c, http.StatusNotFound, message, data)
}

// BadRequest https://httpstatuses.com/400
// The server cannot or will not process the request due to something that is perceived to be a client error
// (e.g., malformed request syntax, invalid request message framing, or deceptive request routing)
func BadRequest(c *gin.Context, message string, data interface{}) {
	jsonResponse(c, http.StatusBadRequest, message, data)
}

// Created https://httpstatuses.com/201.
// The request has been fulfilled and has resulted in one or more new resources being created.
func Created(c *gin.Context, message string, data interface{}) {
	jsonResponse(c, http.StatusCreated, message, data)
}

// Error build common error response. .
func Error(c *gin.Context, statusCode int, message string, data interface{}) {
	jsonResponse(c, statusCode, message, data)
}
