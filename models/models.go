package models

import (
	"fmt"
	"time"

	"github.com/rohit-ambre/go-auth-sql/database"
	"gorm.io/gorm"
)

type User struct {
	UserID    uint64 `gorm:"primaryKey; column:user_id"`
	EmailID   string `gorm:"column:email_id"`
	Password  string `gorm:"column:password"`
	FirstName string `gorm:"column:first_name"`
	LastName  string `gorm:"column:last_name"`
	Active    bool   `gorm:"column:active"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type SignUpReq struct {
	EmailID   string
	Password  string
	FirstName string
	LastName  string
}

type LoginReq struct {
	EmailID  string
	Password string
}

func GetUserByEmail(EmailID string) (User, error) {
	var users User
	db := database.InitDB()

	// db.Where("active=1 and email_id=?", reqBody.EmailID).Find(&users)
	r := db.Model(&User{}).Where("active=1 and email_id=?", EmailID).First(&users)

	if r.Error == gorm.ErrRecordNotFound {
		fmt.Println("User not found", r.Error)
		// log.Fatal(r.Error)
		return User{}, nil
	}

	if r.Error != nil {
		fmt.Println("Error", r.Error)
		return User{}, r.Error
	}

	return users, nil
}

func GetUserByID(UserID uint64) User {
	var user User
	db := database.InitDB()

	// db.Where("active=1 and email_id=?", reqBody.EmailID).Find(&users)
	db.Model(&User{}).Where("active=1 and user_id=?", UserID).First(&user)
	fmt.Println(user)

	return user
}

func CreateUser(userReq User) (User, error) {
	db := database.InitDB()

	r := db.Create(&userReq)

	if r.Error != nil {
		fmt.Println("error creating user")
		return userReq, r.Error
	}
	// fmt.Println("R", r)
	fmt.Println("userReq", userReq.UserID)

	return userReq, nil
}

func GetAllUsers() ([]User, error) {
	db := database.InitDB()

	var users []User

	r := db.Where("active = 1").Find(&users)

	if r.Error != nil {
		if r.Error == gorm.ErrRecordNotFound {
			return []User{}, nil
		} else {
			return []User{}, r.Error
		}
	}

	return users, nil
}
