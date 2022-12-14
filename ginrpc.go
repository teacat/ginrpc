package ginrpc

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func New[T any, M any](handler func(*gin.Context, T) (M, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req T
		if err := c.BindJSON(&req); err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		resp, err := handler(c, req)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, resp)
	}
}

func NewForm[T any, M any](handler func(*gin.Context, T) (M, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req T
		if err := c.BindWith(&req, binding.FormMultipart); err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		resp, err := handler(c, req)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, resp)
	}
}
