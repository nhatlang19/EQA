package controller

import (
	"EQA/backend/app/service"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Login(c *gin.Context)
}

type AuthControllerImpl struct {
	svc service.AuthService
}

// @Summary Login
// @Tags Auth
// @Description Login for user
// @Accept json
// @Produce json
// @Param user body dto.LoginReq true "Login Data"
// @Success 200 {object} dto.LoginResp
// @Router /auth/login [post]
func (u AuthControllerImpl) Login(c *gin.Context) {
	u.svc.Login(c)
}

func AuthControllerInit(authService service.AuthService) *AuthControllerImpl {
	return &AuthControllerImpl{
		svc: authService,
	}
}
