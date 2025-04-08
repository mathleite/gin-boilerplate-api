package database

import (
	"fmt"
	"strings"
)

type User struct {
	BaseTable
	Name     string `gorm:"not null" json:"name"`
	Email    string `gorm:"not null;uniqueIndex" json:"email"`
	Password string `gorm:"not null" json:"password"`
}

func (u *User) Save() error {
	if err := Db.Create(&u).Error; err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return fmt.Errorf("e-mail already exists: %s", u.Email)
		}
		return fmt.Errorf("fail to save user: %v", err)
	}
	return nil
}

func (u *User) FindByEmail(email string) (*User, error) {
	var user User
	if err := Db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, fmt.Errorf("user not found: %v", err)
	}
	return &user, nil
}
