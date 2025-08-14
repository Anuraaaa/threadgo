package http

import (
	"net/http"
	"strconv"

	"github.com/anuraaaa/threadgo/usecase"
	"github.com/gin-gonic/gin"
)

type PostHandlerDeps interface {
	Create(usecase.CreatePostInput) (*interface { /*unused*/
	}, error)
	List(page, limit int) (interface{}, int64, error)
}

// To avoid interface gymnastics in deps, we implement with concrete type below

type postHandler struct{ uc *usecase.PostUseCase }

func NewPostHandler(rg *gin.RouterGroup, uc *usecase.PostUseCase) {
	h := &postHandler{uc: uc}
	rg.POST("/posts", h.create)
}

func NewPublicPostHandler(rg *gin.RouterGroup, uc *usecase.PostUseCase) {
	h := &postHandler{uc: uc}
	rg.GET("/posts", h.list)
}

func (h *postHandler) create(c *gin.Context) {
	uid := c.MustGet("user_id").(uint)
	content := c.PostForm("content")
	file, _ := c.FormFile("file") // optional

	in := usecase.CreatePostInput{UserID: uid, Content: content, File: file}
	p, err := h.uc.Create(in)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, p)
}

func (h *postHandler) list(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	items, total, err := h.uc.List(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": items, "total": total, "page": page, "limit": limit})
}
