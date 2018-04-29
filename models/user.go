package models

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// User ...
type User struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	Name      string     `json:"name"`
	Email     string     `json:"email" gorm:"type:varchar(100);unique_index"`
	Address   string     `json:"address"`
	Password  string     `json:"password" gorm:"size:60"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

// HashPassword ...
func (u *User) HashPassword(password string) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		fmt.Println("Hash password error", err)
	}

	u.Password = string(hash)
}

// CheckPassword ...
func (u *User) CheckPassword(password string) error {

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		fmt.Println("Password not match")
		return err
	}

	return nil
}
