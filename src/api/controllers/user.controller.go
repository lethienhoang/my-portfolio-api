package controllers

import (
	"my-portfolio-api/database"
	"my-portfolio-api/repositories"
	"my-portfolio-api/services"
	"my-portfolio-api/utils/responses"
	"net/http"

	"github.com/gin-gonic/gin/binding"

	uuid "github.com/satori/go.uuid"

	"github.com/gin-gonic/gin"
)

// SignIn godoc
// @Description login to system
// @Accept  json
// @Produce  json
// @Param email body string true "Email"
// @Param password body string true "Password"
// @Success 200 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Router /users/signin [post]
func SignIn(c *gin.Context) {
	dbContext, err := database.ConnectDb()
	if err != nil {
		responses.ERROR(c, http.StatusInternalServerError, err)
		return
	}

	email, _ := c.GetPostForm("email")
	password, _ := c.GetPostForm("password")

	repo := repositories.NewUserRepository(dbContext.GetDbContext())
	service := services.NewUserService(repo)

	token, err := service.SignIn(email, password)
	if err != nil {
		responses.ERROR(c, http.StatusInternalServerError, err)
		return
	}

	res := map[string]string{
		"access_token": token,
	}

	responses.OK(c, res)
}

// SignUp godoc
// @Description login to system
// @Accept  json
// @Produce  json
// @Param email body string true "Email"
// @Param password body string true "Password"
// @Success 200 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Router /users/signup [post]
func SignUp(c *gin.Context) {
	dbContext, err := database.ConnectDb()
	if err != nil {
		responses.ERROR(c, http.StatusInternalServerError, err)
		return
	}

	var req map[string]string
	if err := c.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		responses.ERROR(c, http.StatusBadRequest, err)
		return
	}

	email := req["email"]
	password := req["password"]

	repo := repositories.NewUserRepository(dbContext.GetDbContext())
	service := services.NewUserService(repo)

	err = service.SignUp(email, password)
	if err != nil {
		responses.ERROR(c, http.StatusInternalServerError, err)
		return
	}

	res := map[string]string{
		"message": "Account is created successfully",
	}

	responses.OK(c, res)
}

// SignOut godoc
// @Description login to system
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Router /users/signout [get]
func SignOut(c *gin.Context) {
	dbContext, err := database.ConnectDb()
	if err != nil {
		responses.ERROR(c, http.StatusInternalServerError, err)
		return
	}

	repo := repositories.NewUserRepository(dbContext.GetDbContext())
	service := services.NewUserService(repo)

	err = service.SignOut(c)
	if err != nil {
		responses.ERROR(c, http.StatusInternalServerError, err)
		return
	}

	res := map[string]string{
		"message": "Logout is successfully",
	}

	responses.OK(c, res)
}

// UpdatePassword godoc
// @Description login to system
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param id path string true "UserId"
// @Param password body string true "Password"
// @Success 200 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Router /users/update-password [put]
func UpdatePassword(c *gin.Context) {
	dbContext, err := database.ConnectDb()
	if err != nil {
		responses.ERROR(c, http.StatusInternalServerError, err)
		return
	}

	repo := repositories.NewUserRepository(dbContext.GetDbContext())
	service := services.NewUserService(repo)

	id, _ := uuid.FromString(c.Param("id"))
	pass := c.PostForm("password")

	err = service.UpdateUserPassword(id, pass)
	if err != nil {
		responses.ERROR(c, http.StatusInternalServerError, err)
		return
	}

	res := map[string]string{
		"message": "Logout is successfully",
	}

	responses.OK(c, res)
}
