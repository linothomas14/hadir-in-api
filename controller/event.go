package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/linothomas14/hadir-in-api/helper"
	"github.com/linothomas14/hadir-in-api/helper/param"
	"github.com/linothomas14/hadir-in-api/service"
)

type EventController interface {
	CreateEvent(context *gin.Context)
}

type eventController struct {
	eventService service.EventService
}

func NewEventController(eventService service.EventService) EventController {
	return &eventController{
		eventService: eventService,
	}
}

func (c *eventController) CreateEvent(ctx *gin.Context) {
	var eventParam param.CreateEvent

	err := ctx.ShouldBind(&eventParam)
	UserID := helper.GetUserIdFromClaims(ctx)

	if err != nil {
		response := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	err = helper.ValidateStruct(eventParam)

	if err != nil {
		response := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	event, err := c.eventService.CreateEvent(eventParam, uint32(UserID))

	if err != nil {
		response := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := helper.BuildResponse("OK", event)
	ctx.JSON(http.StatusCreated, response)
}
