package http

import (
	"net/http"
	"strconv"

	"github.com/anuraaaa/threadgo/usecase"
	"github.com/gin-gonic/gin"
)

type CommentHandlerDeps interface {
	Create(usecase.CreateCommentInput) error
	List(postID uint, page, limit int) (interface{}, int64, error)
}

type commentHandler struct{ uc *usecase.CommentUseCase }

func NewCommentHandler(rg *gin.RouterGroup, uc *usecase.CommentUseCase) {
	h := &commentHandler{uc: uc}
	rg.POST("/posts/:id/comments", h.create)
	rg.GET("/posts/:id/comments", h.list)
}

func (h *commentHandler) create(c *gin.Context) {
	uid := c.MustGet("user_id").(uint)
	id, _ := strconv.Atoi(c.Param("id"))
	var req struct {
		Content string `json:"content"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.uc.Create(usecase.CreateCommentInput{PostID: uint(id), UserID: uid, Content: req.Content}); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "commented"})
}

func (h *commentHandler) list(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	items, total, err := h.uc.List(uint(id), page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": items, "total": total, "page": page, "limit": limit})
}
