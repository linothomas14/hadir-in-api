package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/linothomas14/hadir-in-api/helper"
	"github.com/linothomas14/hadir-in-api/helper/response"
	"github.com/linothomas14/hadir-in-api/model"
	"github.com/linothomas14/hadir-in-api/service"
)

type UserController interface {
	GetProfile(context *gin.Context)
	Update(context *gin.Context)
}

type userController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &userController{
		userService: userService,
	}
}

func (c *userController) GetProfile(ctx *gin.Context) {
	var user response.UserResponse

	userID := helper.GetUserIdFromClaims(ctx)

	if userID == 0 {
		response := helper.BuildResponse("User id = 0, there is error occur", helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	user, err := c.userService.GetProfile(userID)

	if err != nil {
		response := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.BuildResponse("OK", user)
	ctx.JSON(http.StatusOK, response)

}

func (c *userController) Update(ctx *gin.Context) {

	var userParam model.User

	userID := helper.GetUserIdFromClaims(ctx)

	userParam.ID = uint32(userID)

	err := ctx.ShouldBind(&userParam)
	if err != nil {
		response := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	err = helper.ValidateStruct(userParam)

	if err != nil {
		response := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	user, err := c.userService.Update(userParam)

	if err != nil {
		response := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.BuildResponse("Updated", user)
	ctx.JSON(http.StatusOK, response)
}
