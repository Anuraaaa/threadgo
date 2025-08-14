package repository

import (
	"github.com/anuraaaa/threadgo/domain"
	"gorm.io/gorm"
)

type LikeGorm struct{ db *gorm.DB }

func NewLikeGorm(db *gorm.DB) *LikeGorm { return &LikeGorm{db: db} }

func (r *LikeGorm) Create(l *domain.Like) error { return r.db.Create(l).Error }

func (r *LikeGorm) Delete(userID, postID uint) error {
	return r.db.Where("user_id = ? AND post_id = ?", userID, postID).Delete(&domain.Like{}).Error
}

func (r *LikeGorm) Count(postID uint) (int64, error) {
	var n int64
	err := r.db.Model(&domain.Like{}).Where("post_id = ?", postID).Count(&n).Error
	return n, err
}

func (r *LikeGorm) IsLiked(userID, postID uint) (bool, error) {
	var n int64
	if err := r.db.Model(&domain.Like{}).Where("user_id = ? AND post_id = ?", userID, postID).Count(&n).Error; err != nil {
		return false, err
	}
	return n > 0, nil
}
