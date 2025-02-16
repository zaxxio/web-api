package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"net/http"
	"web/model"
	"web/service"
)

type UserController struct {
	UserService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{UserService: userService}
}

// @Summary Create a new user
// @Description Creates a new user with name, email, and password
// @Tags users
// @Accept json
// @Produce json
// @Security    BearerAuth
// @Success     200  {object} map[string]string
// @Failure     401  {object} map[string]string
// @Param user body model.User true "User Data"
// @Success 201 {object} model.User
// @Failure 400 {string} string "Bad Request"
// @Router /users [post]
func (userController *UserController) CreateUser(ctx *gin.Context) {
	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := userController.UserService.CreateUser(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusCreated, user)
}

// @Summary Get all users
// @Description Fetches all users from the database
// @Tags users
// @Produce json
// @Security    BearerAuth
// @Success     200  {object} map[string]string
// @Failure     401  {object} map[string]string
// @Success 200 {array} model.User
// @Router /users [get]
func (userController *UserController) GetUsers(ctx *gin.Context) {
	users, err := userController.UserService.GetUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	ctx.JSON(http.StatusOK, users)
}

var ControllerModule = fx.Module("user_controller", fx.Provide(NewUserController))
