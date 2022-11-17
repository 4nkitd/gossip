package controller

import (
	"net/http"

	"github.com/dagar-in/try-catch"
	"github.com/gin-gonic/gin"
)

func Response(c *gin.Context, data interface{}, err try.E) {
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	} else {
		c.JSON(http.StatusOK, data)
	}
}
