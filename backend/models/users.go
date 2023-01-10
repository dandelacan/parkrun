package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	_ "github.com/lib/pq"
)

type User struct {
	gorm.Model
	Name      string
	BarcodeID uint
}

type UserService struct {
	db *gorm.DB
}

func NewUserService(conectionInfo string) (*UserService, error) {
	db, err := gorm.Open("postgres", conectionInfo)
	if err != nil {
		return nil, err
	}
	db.LogMode(true)
	return &UserService{db: db}, nil
}

func (us *UserService) AutoMigrate() error {
	if err := us.db.AutoMigrate(&User{}).Error; err != nil {
		return err
	}
	return nil
}

func (us *UserService) Close() error {
	return us.db.Close()
}

func (us *UserService) Create(user *User) error {

	return us.db.Create(user).Error
}

func (us *UserService) AllUsersse() ([]User, error) {
	var users []User

	db := us.db.Find(&users)

	if db.Error != nil {
		return nil, db.Error
	}

	return users, nil
}

func (us *UserService) DestructiveReset() error {
	if err := us.db.DropTableIfExists(&User{}).Error; err != nil {
		return err
	}
	return us.AutoMigrate()
}
