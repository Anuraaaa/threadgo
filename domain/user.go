package domain

import "time"

type User struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    Name      string    `json:"name"`
    Email     string    `gorm:"uniqueIndex" json:"email"`
    Password  string    `json:"-"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

type UserRepository interface {
    Create(u *User) error
    GetByEmail(email string) (*User, error)
    GetByID(id uint) (*User, error)
}