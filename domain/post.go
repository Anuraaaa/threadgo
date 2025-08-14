package domain

import "time"

type Post struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    UserID    uint      `json:"user_id"`
    Content   string    `json:"content"`
    ImageURL  *string   `json:"image_url"`
    Likes     []Like    `json:"-"`
    Comments  []Comment `json:"-"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

type PostRepository interface {
    Create(p *Post) error
    GetByID(id uint) (*Post, error)
    List(page, limit int) ([]Post, int64, error)
}