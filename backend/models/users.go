package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

var (
	// ErrNotFound is returned when a resource cannot be found // in the database.
	ErrNotFound = errors.New("Models: Resource not found")
)

// User defines the table for users
type User struct {
	gorm.Model
	Name  string
	Email string `gorm:"not null; unique index"`
}

// NewUserService creates a new database connection to serve our users
func NewUserService(connectionInfo string) (*UserService, error) {
	db, err := gorm.Open("postgres", connectionInfo)
	if err != nil {
		return nil, err
	}
	db.LogMode(true)
	return &UserService{
		db: db,
	}, nil
}

// UserService handles user related data
type UserService struct {
	db *gorm.DB
}

// ByID will return a user of matching id
func (us *UserService) ByID(id uint) (*User, error) {
	var user User
	db := us.db.Where("id = ?", id)
	err := first(db, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// ByEmail will return the first user with matching emai
func (us *UserService) ByEmail(email string) (*User, error) {
	var user User
	db := us.db.Where("email = ?", email)
	err := first(db, &user)
	return &user, err
}

// Create creates a user and inserts into database
func (us *UserService) Create(user *User) error {
	return us.db.Create(user).Error
}

// Close closes the UserService database connection
func (us *UserService) Close() error {
	return us.db.Close()
}

// DestructiveReset drops our tables and rebuilds them
func (us *UserService) DestructiveReset() {
	us.db.DropTableIfExists(&User{})
	us.db.AutoMigrate(&User{})
}

func first(db *gorm.DB, dst interface{}) error {
	err := db.First(dst).Error
	if err == gorm.ErrRecordNotFound {
		return ErrNotFound
	}
	return err
}
