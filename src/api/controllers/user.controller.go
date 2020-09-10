package controllers

import (
	"my-portfolio-api/database"
	"my-portfolio-api/repositories"
	"my-portfolio-api/services"
	"my-portfolio-api/utils/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SignIn login to system
func SignIn(c *gin.Context) {
	dbContext, err := database.ConnectDb()
	if err != nil {
		responses.ERROR(c, http.StatusInternalServerError, err)
	}

	defer dbContext.Close()

	email, _ := c.GetPostForm("email")
	password, _ := c.GetPostForm("password")

	repo := repositories.NewUserRepository(dbContext.GetDbContext())
	service := services.NewUserService(repo)

	token, err := service.SignIn(email, password)
	if err != nil {
		responses.ERROR(c, http.StatusInternalServerError, err)
	}

	tokens := map[string]string{
		"access_token": token,
	}

	responses.OK(c, tokens)
}

// SignOut logout system
func SignOut(c *gin.Context) {
	dbContext, err := database.ConnectDb()
	if err != nil {
		responses.ERROR(c, http.StatusInternalServerError, err)
	}

	defer dbContext.Close()

	repo := repositories.NewUserRepository(dbContext.GetDbContext())
	service := services.NewUserService(repo)

	err = service.SignOut(c)
	if err != nil {
		responses.ERROR(c, http.StatusInternalServerError, err)
	}

	responses.OK(c, "Logout is successfully")
}
