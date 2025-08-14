package repository

import (
    "github.com/anuraaaa/threadgo/domain"
    "gorm.io/gorm"
)

type CommentGorm struct{ db *gorm.DB }

func NewCommentGorm(db *gorm.DB) *CommentGorm { return &CommentGorm{db: db} }

func (r *CommentGorm) Create(cmt *domain.Comment) error { return r.db.Create(cmt).Error }

func (r *CommentGorm) ListByPost(postID uint, page, limit int) ([]domain.Comment, int64, error) {
    if page < 1 { page = 1 }
    if limit < 1 { limit = 10 }
    var items []domain.Comment
    var total int64
    r.db.Model(&domain.Comment{}).Where("post_id = ?", postID).Count(&total)
    err := r.db.Where("post_id = ?", postID).Order("created_at asc").
        Offset((page-1)*limit).Limit(limit).Find(&items).Error
    return items, total, err
}
