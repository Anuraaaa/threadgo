package http

import (
	"github.com/anuraaaa/threadgo/usecase"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(rg *gin.RouterGroup, uc *usecase.AuthUseCase) {
	NewAuthHandler(rg, uc)
}

func RegisterPostRoutes(rg *gin.RouterGroup, uc *usecase.PostUseCase) {
	NewPostHandler(rg, uc)
}

func RegisterPublicPostRoutes(rg *gin.RouterGroup, uc *usecase.PostUseCase) {
	NewPublicPostHandler(rg, uc)
}

func RegisterCommentRoutes(rg *gin.RouterGroup, uc *usecase.CommentUseCase) {
	NewCommentHandler(rg, uc)
}

func RegisterLikeRoutes(rg *gin.RouterGroup, uc *usecase.LikeUseCase) {
	NewLikeHandler(rg, uc)
}
