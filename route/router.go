package route

import (
	"hng_stage_0/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getApiResponse(context *gin.Context) {
	response, err := model.NewApiResponse()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch response. Try again later.", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, response)
}

func RegisterRoutes(router *gin.Engine) {
	router.GET("/hngstage0", getApiResponse)
}
