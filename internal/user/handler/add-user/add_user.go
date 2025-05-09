package handler_adduser

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	user_repository "github.com/williamkoller/divine-beast/internal/user/repository"
	adduser_usecase "github.com/williamkoller/divine-beast/internal/user/usecases/add-user"
)

type UserRequest struct {
	Email string `json:"email" binding:"required,email"`
	Age   int    `json:"age" binding:"required,min=18"`
}

type UserResponse struct {
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func (ur UserRequest) ToUserResponse() *UserResponse {
	return &UserResponse{
		Email: ur.Email,
		Age:   ur.Age,
	}
}

func handleError(c *gin.Context, err error) {
	switch err {
	case adduser_usecase.ErrUserAlreadyExists:
		c.JSON(http.StatusConflict, gin.H{
			"error":      err.Error(),
			"statusCode": http.StatusConflict,
		})
	case adduser_usecase.ErrInvalidEmail, adduser_usecase.ErrInvalidAge:
		c.JSON(http.StatusBadRequest, gin.H{
			"error":      err.Error(),
			"statusCode": http.StatusBadRequest,
		})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":      err.Error(),
			"statusCode": http.StatusInternalServerError,
		})
	}
}

func AddUser(c *gin.Context) {
	var userRequest UserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			var errorMessages []string
			var errorFields []string

			for _, fieldError := range validationErrors {
				errorMessages = append(errorMessages, fieldError.Error())
				errorFields = append(errorFields, fieldError.Field())
			}

			c.JSON(http.StatusBadRequest, gin.H{
				"message":    errorMessages,
				"fields":     errorFields,
				"statusCode": http.StatusBadRequest,
			})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	usecase := adduser_usecase.NewAddUserUseCase(user_repository.NewInMemoryUserRepository())
	response := userRequest.ToUserResponse()

	if err := usecase.Execute(response.Email, response.Age); err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response)
}
