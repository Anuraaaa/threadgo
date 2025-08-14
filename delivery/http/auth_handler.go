package http

import (
	"net/http"
	"time"

	"github.com/anuraaaa/threadgo/usecase"
	"github.com/gin-gonic/gin"
)

type AuthHandlerDeps interface {
	Register(usecase.AuthRegisterInput) error
	Login(usecase.AuthLoginInput) (string, interface{}, error)
}

type authHandler struct{ uc *usecase.AuthUseCase }

func NewAuthHandler(rg *gin.RouterGroup, uc *usecase.AuthUseCase) {
	h := &authHandler{uc: uc}
	auth := rg.Group("/auth")
	auth.POST("/register", h.register)
	auth.POST("/login", h.login)
}

func (h *authHandler) register(c *gin.Context) {
	var req struct{ Name, Email, Password string }
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.uc.Register(usecase.AuthRegisterInput{Name: req.Name, Email: req.Email, Password: req.Password}); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "registered"})
}

func (h *authHandler) login(c *gin.Context) {
	var req struct{ Email, Password string }
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, user, err := h.uc.Login(usecase.AuthLoginInput{Email: req.Email, Password: req.Password})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token, "user": user, "expires_in": int((24 * time.Hour).Seconds())})
}
