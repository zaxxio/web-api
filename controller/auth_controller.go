package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"

	"web/common"
	"web/model"
	"web/service"
)

// AuthController handles user sign-up and sign-in.
type AuthController struct {
	UserService *service.UserService
}

// NewAuthController is an FX constructor function.
func NewAuthController(userService *service.UserService) *AuthController {
	return &AuthController{UserService: userService}
}

// SignUp godoc
// @Summary     Register a new user
// @Description Creates a new user account in the database
// @Tags        Auth Controller
// @Accept      json
// @Produce     json
// @Param       user body model.SignUpRequest true "SignUp data"
// @Success     201  {object} model.SignUpRequest
// @Failure     400  {object} map[string]string
// @Failure     500  {object} map[string]string
// @Router      /auth/signup [post]
func (ac *AuthController) SignUp(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// In real-world code, hash your password before saving:
	// user.Password = hashPassword(user.Password)

	if err := ac.UserService.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// SignIn godoc
// @Summary     Log in a user
// @Description Authenticates user credentials and returns a JWT token
// @Tags        auth
// @Accept      json
// @Produce     json
// @Param       credentials body model.LoginRequest true "SignIn credentials"
// @Success     200  {object} map[string]string
// @Failure     401  {object} map[string]string
// @Failure     500  {object} map[string]string
// @Router      /auth/signin [post]
func (ac *AuthController) SignIn(c *gin.Context) {
	var credentials model.LoginRequest
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 1. Verify user exists
	user, err := ac.UserService.GetUserByEmail(credentials.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// 2. Check password (in real code, compare hashed password)
	if user.Password != credentials.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// 3. Generate a JWT token
	token, err := common.GenerateJWT(user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// AuthControllerModule is the Fx module that provides AuthController.
var AuthControllerModule = fx.Module(
	"auth_controller",
	fx.Provide(NewAuthController),
)
