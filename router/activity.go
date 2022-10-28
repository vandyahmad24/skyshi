package router

import (
	"vandyahmad/skyshi/activity"
	"vandyahmad/skyshi/config"
	"vandyahmad/skyshi/handler"

	"github.com/gin-gonic/gin"
)

func ActivityRouter(app *gin.Engine) {

	activityRepository := activity.NewRepository(config.DB)
	activityService := activity.NewService(activityRepository)
	activityHandler := handler.NewActivityHandler(activityService)

	app.GET("/activity-groups", activityHandler.ListActivity)
	app.GET("/activity-groups/:activityId", activityHandler.DetailActivity)
	app.POST("/activity-groups", activityHandler.CreateActivity)
	app.PATCH("/activity-groups/:activityId", activityHandler.UpdateActivity)
	app.DELETE("/activity-groups/:activityId", activityHandler.DeleteActivity)

}
