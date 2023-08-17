package v1

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kalleriakronos24/mygoapp2nd/dto"
	"github.com/kalleriakronos24/mygoapp2nd/services"
)

func POSTLogin(c *gin.Context) {
	var err error
	var user dto.UserLogin
	if err = c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Error: err.Error()})
		return
	}
	var token string
	if token, err = services.Handler.AuthenticateUser(user); err != nil {
		c.JSON(http.StatusUnauthorized, dto.Response{Error: "incorrect username or password"})
		return
	}
	c.JSON(http.StatusOK, dto.Response{Data: fmt.Sprintf("Bearer %s", token)})
}

func POSTRegister(c *gin.Context) {
	var err error
	var user dto.UserSignup
	if err = c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Error: err.Error()})
		return
	}
	if err = services.Handler.RegisterUser(user); err != nil {
		c.JSON(http.StatusUnauthorized, dto.Response{Error: "failed registering user"})
		return
	}
	c.JSON(http.StatusCreated, dto.Response{Data: "user created"})
}
