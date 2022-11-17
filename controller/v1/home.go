package v1

import (
	rest "fs/api/controller"
	"fs/api/utils"

	"github.com/dagar-in/try-catch"
	"github.com/gin-gonic/gin"
)

var err try.E

func Index(c *gin.Context) {

	try.This(func() {

		panic("my panic")

	}).Finally(func() {

		rest.Response(c, gin.H{"message": "Hello World"}, err)

	}).Catch(func(e try.E) {

		// loggin
		utils.At().Error().Msg(e.(string))

	})

}
