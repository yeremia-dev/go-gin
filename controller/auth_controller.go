package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yeremia-dev/go-gin/dto"
	"github.com/yeremia-dev/go-gin/entity"
	"github.com/yeremia-dev/go-gin/helper"
	"github.com/yeremia-dev/go-gin/service"
)

//AuthController interface is a contract for auth controller
type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
	authService service.AuthService
	jwtService  service.JWTService
}

func NewAuthController(authService service.AuthService, jwtService service.JWTService) AuthController {
	return &authController{
		authService: authService,
		jwtService:  jwtService,
	}

}

func (c *authController) Login(ctx *gin.Context) {
	var loginDto dto.LoginDTO
	errDto := ctx.ShouldBind(&loginDto)
	if errDto != nil {
		response := helper.BuildErrorResponse("Failde to process request", errDto.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	authResult := c.authService.VerifyCredential(loginDto.Email, loginDto.Password)
	if v, ok := authResult.(entity.User); ok {
		generatedToken := c.jwtService.GenerateToken(strconv.FormatUint(v.ID, 10))
		v.Token = generatedToken
		response := helper.BuildResponse(true, "OK!", v)
		ctx.JSON(http.StatusOK, response)
		return
	}
	response := helper.BuildErrorResponse("Please check your credential again", "Invalid credential", helper.EmptyObject{})
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)

}

func (c *authController) Register(ctx *gin.Context) {

	var registerDto dto.RegisterDTO
	errDto := ctx.ShouldBind(&registerDto)
	if errDto != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDto.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if !c.authService.IsDuplicateEmail(registerDto.Email) {
		response := helper.BuildErrorResponse("Failed to process request", "Duplicate email", helper.EmptyObject{})
		ctx.JSON(http.StatusConflict, response)
	} else {
		createdUser := c.authService.CreateUser(registerDto)
		token := c.jwtService.GenerateToken(strconv.FormatUint(createdUser.ID, 10))
		createdUser.Token = token
		response := helper.BuildResponse(true, "OK!", createdUser)
		ctx.JSON(http.StatusCreated, response)
	}

}
