package service

import (
	"EQA/backend/app/constant"
	"EQA/backend/app/domain/dto"
	"EQA/backend/app/helper"
	"EQA/backend/app/pkg"
	"EQA/backend/app/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(c *gin.Context)
}

type AuthServiceImpl struct {
	userRepo repository.UserRepository
}

func AuthServiceInit(userRepo repository.UserRepository) *AuthServiceImpl {
	return &AuthServiceImpl{
		userRepo: userRepo,
	}
}

func (s AuthServiceImpl) Login(c *gin.Context) {
	defer pkg.PanicHandler(c)

	var req dto.LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		pkg.PanicException(constant.BadRequest, "")
	}

	user, err := s.userRepo.FindByUsername(req.Username)
	if err != nil {
		pkg.PanicException(constant.DataNotFound, "User does not exist")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		pkg.PanicException(constant.Unauthorized, "Invalid username or password")
	}

	token, err := helper.CreateToken(user.ID, user.Username)

	if err != nil {
		pkg.PanicException(constant.UnknownError, "Error when generate JWT token")
	}

	resp := dto.LoginResp{
		AccessToken: token,
	}
	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, resp))
}
