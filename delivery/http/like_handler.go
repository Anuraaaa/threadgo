package http

import (
	"net/http"
	"strconv"

	"github.com/anuraaaa/threadgo/usecase"
	"github.com/gin-gonic/gin"
)

type LikeHandlerDeps interface {
	Like(userID, postID uint) error
	Unlike(userID, postID uint) error
}

type likeHandler struct{ uc *usecase.LikeUseCase }

func NewLikeHandler(rg *gin.RouterGroup, uc *usecase.LikeUseCase) {
	h := &likeHandler{uc: uc}
	rg.POST("/posts/:id/like", h.like)
	rg.DELETE("/posts/:id/like", h.unlike)
}

func (h *likeHandler) like(c *gin.Context) {
	uid := c.MustGet("user_id").(uint)
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.uc.Like(uid, uint(id)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "liked"})
}

func (h *likeHandler) unlike(c *gin.Context) {
	uid := c.MustGet("user_id").(uint)
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.uc.Unlike(uid, uint(id)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "unliked"})
}
