package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/anuraaaa/threadgo/config"
	httpDelivery "github.com/anuraaaa/threadgo/delivery/http"
	"github.com/anuraaaa/threadgo/middleware"
	"github.com/anuraaaa/threadgo/repository"
	"github.com/anuraaaa/threadgo/storage"
	"github.com/anuraaaa/threadgo/usecase"
	"github.com/joho/godotenv"
)

func main() {
	// init DB
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	db := config.MustOpenDB()

	// storage (local uploads)
	fs := storage.NewLocalStorage("./uploads")

	// repositories
	userRepo := repository.NewUserGorm(db)
	postRepo := repository.NewPostGorm(db)
	commentRepo := repository.NewCommentGorm(db)
	likeRepo := repository.NewLikeGorm(db)

	// usecases
	authUC := usecase.NewAuthUseCase(userRepo)
	postUC := usecase.NewPostUseCase(postRepo, fs)
	commentUC := usecase.NewCommentUseCase(commentRepo, postRepo)
	likeUC := usecase.NewLikeUseCase(likeRepo, postRepo)

	r := gin.Default()

	// static serving for uploads
	r.Static("/uploads", "./uploads")

	// middlewares
	r.Use(gin.Recovery())

	// router wiring
	api := r.Group("/api/v1")
	httpDelivery.RegisterAuthRoutes(api, authUC)

	// protected group
	auth := api.Group("")
	auth.Use(middleware.JWTAuth())
	httpDelivery.RegisterPostRoutes(auth, postUC)
	httpDelivery.RegisterCommentRoutes(auth, commentUC)
	httpDelivery.RegisterLikeRoutes(auth, likeUC)

	// public read routes
	httpDelivery.RegisterPublicPostRoutes(api, postUC)

	log.Println("listening thread go on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
