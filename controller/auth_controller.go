package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//AuthController interface is a contract for auth controller
type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
}

func NewAuthController() AuthController {
	return &authController{}

}

func (c *authController) Login(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Login",
	})

}

func (c *authController) Register(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Register",
	})

}
