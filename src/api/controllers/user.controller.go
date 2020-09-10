package controllers

import (
	"context"
	"my-portfolio-api/services"
	"my-portfolio-api/repositories"
	"my-portfolio-api/database"
	"my-portfolio-api/utils/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignIn(c *gin.Context) {
	dbContext, err := database.ConnectDb()
	if err != nil {
		responses.ERROR(c, http.StatusInternalServerError, err.Error()
	}

	defer dbContext.Close()

	email,_ := c.GetPostForm("email")
	password,_ := c.GetPostForm("password")
	
	repo := repositories.NewUserRepository(dbContext.GetDbContext())
	service := services.NewUserService(repo)

	token, err := service.SignIn(email, password)
	if err != nil {
		responses.ERROR(c, http.StatusInternalServerError, err.Error())
	}

	tokens := map[string]string{
		"access_token": token
	}

	responses.OK(c, tokens) 
}

func SignOut(c *gin.Context) {
	dbContext, err := database.ConnectDb()
	if err != nil {
		responses.ERROR(c, http.StatusInternalServerError, err.Error()
	}

	defer dbContext.Close()
	
	repo := repositories.NewUserRepository(dbContext.GetDbContext())
	service := services.NewUserService(repo)

	token, err := service.SignOut(c)
	if err != nil {
		responses.ERROR(c, http.StatusInternalServerError, err.Error())
	}

	
	responses.OK(c, token) 
}
