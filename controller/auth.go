package controller

import (
	"log"
	"net/http"

	"github.com/linothomas14/hadir-in-api/helper"
	"github.com/linothomas14/hadir-in-api/helper/param"
	"github.com/linothomas14/hadir-in-api/helper/response"
	"github.com/linothomas14/hadir-in-api/service"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Login(context *gin.Context)
	Register(context *gin.Context)
	GetProfile(context *gin.Context)
}

type authController struct {
	authService service.AuthService
}

func NewAuthController(authService service.AuthService) AuthController {
	return &authController{
		authService: authService,
	}
}

func (c *authController) Login(ctx *gin.Context) {
	var loginparam param.Login
	var res response.TokenResponse
	err := ctx.ShouldBind(&loginparam)

	if err != nil {
		response := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	err = helper.ValidateStruct(loginparam)
	if err != nil {
		response := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	token, err := c.authService.Login(loginparam.Email, loginparam.Password)

	if token == "" {
		response := helper.BuildResponse("invalid credential", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	res.Token = token
	response := helper.BuildResponse("OK", res)
	ctx.JSON(http.StatusOK, response)
	return
}

func (c *authController) Register(ctx *gin.Context) {
	var registerParam param.Register
	var registerResponse response.RegisterResponse

	err := ctx.ShouldBind(&registerParam)
	if err != nil {
		response := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	err = helper.ValidateStruct(registerParam)

	if err != nil {
		response := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if c.authService.IsDuplicateEmail(registerParam.Email) {
		response := helper.BuildResponse("Duplicate email", helper.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
		return
	}

	createdUser, err := c.authService.CreateUser(registerParam)
	if err != nil {
		response := helper.BuildResponse(err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
		return
	}

	registerResponse.ID = createdUser.ID
	registerResponse.Email = createdUser.Email
	registerResponse.Name = createdUser.Name
	response := helper.BuildResponse("OK", registerResponse)
	ctx.JSON(http.StatusCreated, response)
}

func (c *authController) GetProfile(ctx *gin.Context) {
	type responseStruct struct {
		UserID int `json:"user_id"`
	}

	res := responseStruct{}

	userAuth, ok := ctx.Get("user_id")

	log.Println(userAuth)

	if !ok {
		response := helper.BuildResponse("Error1", helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	userId, ok := userAuth.(float64)
	log.Println(userId)

	if !ok {
		response := helper.BuildResponse("Error2", helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	res.UserID = int(userId)
	response := helper.BuildResponse("OK", res)
	ctx.JSON(http.StatusCreated, response)

}
