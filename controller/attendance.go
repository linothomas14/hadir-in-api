package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/linothomas14/hadir-in-api/helper"
	"github.com/linothomas14/hadir-in-api/helper/response"
	"github.com/linothomas14/hadir-in-api/model"
	"github.com/linothomas14/hadir-in-api/service"
)

type AttendanceController interface {
	Attend(context *gin.Context)
}

type attendanceController struct {
	attendanceService service.AttendanceService
}

func NewAttendanceController(attendanceService service.AttendanceService) AttendanceController {
	return &attendanceController{
		attendanceService: attendanceService,
	}
}

func (c *attendanceController) Attend(ctx *gin.Context) {
	var attendance model.Attendance
	var res response.PresentResponse
	type PresentParam struct {
		Token string `json:"token"`
	}

	var tokenEvent PresentParam

	UserID := helper.GetUserIdFromClaims(ctx)

	if UserID == 0 {
		response := helper.BuildResponse("User id = 0, there is error occur", helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	attendance.UserID = uint32(UserID)

	err := ctx.ShouldBind(&tokenEvent)

	if err != nil {
		response := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	res, err = c.attendanceService.Attend(attendance, tokenEvent.Token)

	if err != nil {
		response := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := helper.BuildResponse("OK", res)
	ctx.JSON(http.StatusCreated, response)
}
