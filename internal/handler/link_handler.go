package handler

import (
	"net/http"
	"strconv"
	"strings"

	"shorten-url-be/internal/usecase"

	"github.com/gin-gonic/gin"
)

// LinkHandler handles HTTP requests related to links
type LinkHandler struct {
	usecase *usecase.LinkUseCase
}

// NewLinkHandler creates a new instance of LinkHandler
func NewLinkHandler(usecase *usecase.LinkUseCase) *LinkHandler {
	return &LinkHandler{usecase: usecase}
}

// CreateLink handles the creation of a new shortened link
func (h *LinkHandler) CreateLink(c *gin.Context) {
	var request struct {
		OriginalURL string `json:"original_url"`
	}

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if strings.TrimSpace(request.OriginalURL) == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "OriginalURL must not be empty or null"})
		return
	}

	userID, _ := c.Get("userID")

	link, err := h.usecase.CreateLink(request.OriginalURL, userID.(uint))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating link"})
		return
	}

	c.JSON(http.StatusOK, link)
}

// GetLinkByShortURL returns a link by its short URL
func (h *LinkHandler) GetLinkByShortURL(c *gin.Context) {
	shortURL := c.Param("short_url")
	link, err := h.usecase.GetLinkByShortURL(shortURL)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Link not found"})
		return
	}

	c.Redirect(http.StatusFound, link.OriginalURL)
}

// UpdateLink handles updating the original URL of a link
func (h *LinkHandler) UpdateLink(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var request struct {
		OriginalURL string `json:"original_url"`
	}

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	link, err := h.usecase.UpdateLink(uint(id), request.OriginalURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating link"})
		return
	}

	c.JSON(http.StatusOK, link)
}

// DeleteLink handles deleting a link by its ID
func (h *LinkHandler) DeleteLink(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.usecase.DeleteLink(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting link"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
