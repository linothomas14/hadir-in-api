package repository

import (
	"log"

	"github.com/linothomas14/hadir-in-api/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

//UserRepository is contract what userRepository can do to db
type UserRepository interface {
	InsertUser(user model.User) (model.User, error)
	UpdateUser(user model.User) (model.User, error)
	VerifyCredential(email string, password string) interface{}
	IsDuplicateEmail(email string) model.User
	FindByEmail(email string) model.User
	GetUser(userID int) (model.User, error)
}

type userConnection struct {
	connection *gorm.DB
}

//NewUserRepository is creates a new instance of UserRepository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{
		connection: db,
	}
}

func (db *userConnection) InsertUser(user model.User) (model.User, error) {

	user.Password = hashAndSalt([]byte(user.Password))
	err := db.connection.Save(&user).Error
	return user, err
}

func (db *userConnection) UpdateUser(user model.User) (model.User, error) {

	if user.Password != "" {
		user.Password = hashAndSalt([]byte(user.Password))
	} else {
		var tempUser model.User
		db.connection.Find(&tempUser, user.ID)
		user.Password = tempUser.Password
	}

	err := db.connection.Save(&user).Error
	return user, err
}

func (db *userConnection) VerifyCredential(email string, password string) interface{} {
	var user model.User
	res := db.connection.Where("email = ?", email).Take(&user)
	if res.Error == nil {
		return user
	}
	return nil
}

func (db *userConnection) IsDuplicateEmail(email string) model.User {
	var user model.User
	db.connection.Where("email = ?", email).Take(&user)
	return user
}

func (db *userConnection) FindByEmail(email string) model.User {
	var user model.User
	db.connection.Where("email = ?", email).Take(&user)
	return user
}

func (db *userConnection) GetUser(userId int) (model.User, error) {
	var user model.User
	err := db.connection.Find(&user, userId).Error
	return user, err
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash a password")
	}
	return string(hash)
}
