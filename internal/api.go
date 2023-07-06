package internal

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func sayHi(c *gin.Context) {
	c.String(http.StatusOK, "Hi, there!")
}
