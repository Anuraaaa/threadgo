package domain

import "time"

type Comment struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    PostID    uint      `json:"post_id"`
    UserID    uint      `json:"user_id"`
    Content   string    `json:"content"`
    CreatedAt time.Time `json:"created_at"`
}

type CommentRepository interface {
    Create(c *Comment) error
    ListByPost(postID uint, page, limit int) ([]Comment, int64, error)
}