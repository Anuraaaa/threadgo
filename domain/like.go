package domain

import "time"

type Like struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    PostID    uint      `gorm:"uniqueIndex:uniq_user_post" json:"post_id"`
    UserID    uint      `gorm:"uniqueIndex:uniq_user_post" json:"user_id"`
    CreatedAt time.Time `json:"created_at"`
}

type LikeRepository interface {
    Create(l *Like) error
    Delete(userID, postID uint) error
    Count(postID uint) (int64, error)
    IsLiked(userID, postID uint) (bool, error)
}