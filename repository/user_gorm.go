package repository

import (
    "gorm.io/gorm"
    "github.com/anuraaaa/threadgo/domain"
	"errors"
)

type UserGorm struct{ db *gorm.DB }

func NewUserGorm(db *gorm.DB) *UserGorm { return &UserGorm{db: db} }

func (r *UserGorm) Create(u *domain.User) error { return r.db.Create(u).Error }

func (r *UserGorm) GetByEmail(email string) (*domain.User, error) {
    var u domain.User
    if err := r.db.Where("email = ?", email).First(&u).Error; err != nil {
        return nil, err
    }
    return &u, nil
}

func (r *UserGorm) GetByID(id uint) (*domain.User, error) {
    var u domain.User
    if err := r.db.First(&u, id).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) { return nil, err }
        return nil, err
    }
    return &u, nil
}
