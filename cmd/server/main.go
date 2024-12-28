package main

import (
	"shorten-url-be/internal/config"
	"shorten-url-be/internal/domain/models"
	"shorten-url-be/internal/handler"
	"shorten-url-be/internal/middleware"
	"shorten-url-be/internal/repository"
	"shorten-url-be/internal/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Initialize Database Connection
	db, err := gorm.Open(postgres.Open(config.DBURL), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Link{})

	// Initialize Repository and UseCase
	linkRepo := repository.NewLinkRepositoryGorm(db)
	linkUseCase := usecase.NewLinkUseCase(linkRepo)
	linkHandler := handler.NewLinkHandler(linkUseCase)

	authRepo := repository.NewAuthRepositoryGorm(db)
	authUseCase := usecase.NewAuthUseCase(authRepo)
	authHandler := handler.NewAuthHandler(authUseCase)

	// Set up Gin router
	r := gin.Default()

	r.POST("/login", authHandler.Login)
	r.POST("/sign-up", authHandler.SignUp)

	// Apply authentication middleware protected router
	r.GET("/links/:short_url", linkHandler.GetLinkByShortURL)

	protectGroup := r.Group("/protected")
	protectGroup.Use(middleware.AuthMiddleware)
	{
		// Routes for managing links
		protectGroup.GET("/links", linkHandler.UpdateLink)
		protectGroup.POST("/links", linkHandler.CreateLink)
		protectGroup.PUT("/links/:id", linkHandler.UpdateLink)
		protectGroup.DELETE("/links/:id", linkHandler.DeleteLink)
	}

	// Start the server
	r.Run(":8080")
}
