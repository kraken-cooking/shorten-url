package handler

import (
	"net/http"

	"shorten-url-be/internal/usecase"
	"shorten-url-be/internal/utils"

	"github.com/gin-gonic/gin"
)

// LinkHandler handles HTTP requests related to links
type AuthHandler struct {
	usecase *usecase.AuthUseCase
}

// NewLinkHandler creates a new instance of LinkHandler
func NewAuthHandler(usecase *usecase.AuthUseCase) *AuthHandler {
	return &AuthHandler{usecase: usecase}
}

// CreateLink handles the creation of a new shortened link
func (h *AuthHandler) SignUp(c *gin.Context) {
	var request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	user, err := h.usecase.SignUp(request.Username, request.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error Sign Up"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// GetLinkByShortURL returns a link by its short URL
func (h *AuthHandler) Login(c *gin.Context) {
	var request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var err error

	err = c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	user, err := h.usecase.Login(request.Username, request.Password)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not correct username or password"})
		return
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
