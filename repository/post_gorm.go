package repository

import (
    "github.com/anuraaaa/threadgo/domain"
    "gorm.io/gorm"
)

type PostGorm struct{ db *gorm.DB }

func NewPostGorm(db *gorm.DB) *PostGorm { return &PostGorm{db: db} }

func (r *PostGorm) Create(p *domain.Post) error { return r.db.Create(p).Error }

func (r *PostGorm) GetByID(id uint) (*domain.Post, error) {
    var p domain.Post
    if err := r.db.First(&p, id).Error; err != nil { return nil, err }
    return &p, nil
}

func (r *PostGorm) List(page, limit int) ([]domain.Post, int64, error) {
    if page < 1 { page = 1 }
    if limit < 1 { limit = 10 }
    var posts []domain.Post
    var total int64
    r.db.Model(&domain.Post{}).Count(&total)
    err := r.db.Order("created_at desc").
        Offset((page-1)*limit).Limit(limit).Find(&posts).Error
    return posts, total, err
}