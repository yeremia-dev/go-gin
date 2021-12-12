package repository

import (
	"log"

	"github.com/yeremia-dev/go-gin/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository interface {
	InsertUser(user entity.User) entity.User
	UpdateUser(user entity.User) entity.User
	VerifyCredential(email string, password string) interface{}
	IsDuplicateEmail(email string) (tx *gorm.DB)
	FindByEmail(email string) entity.User
	ProfileUser(userId string) entity.User
}

type userRepository struct {
	connection *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		connection: db,
	}
}

func (rep *userRepository) InsertUser(user entity.User) entity.User {

	user.Password = hashAndSalt([]byte(user.Password))
	rep.connection.Save(&user)
	return user

}

func (rep *userRepository) UpdateUser(user entity.User) entity.User {
	if user.Password != "" {
		user.Password = hashAndSalt([]byte(user.Password))
	} else {
		var tempUser entity.User
		rep.connection.Find(&tempUser, user.ID)
		user.Password = tempUser.Password
	}

	rep.connection.Save(&user)
	return user
}

func (rep *userRepository) VerifyCredential(email string, password string) interface{} {
	var user entity.User
	ress := rep.connection.Where("email = ?", email).Take(&user)
	if ress.Error == nil {
		return user
	}
	return nil
}

func (rep *userRepository) IsDuplicateEmail(email string) (tx *gorm.DB) {
	var user entity.User
	return rep.connection.Where("email = ?", email).Take(&user)
}

func (rep *userRepository) FindByEmail(email string) entity.User {
	var user entity.User
	rep.connection.Where("email = ?", email).Take(&user)
	return user
}

func (rep *userRepository) ProfileUser(userId string) entity.User {
	var user entity.User
	rep.connection.Find(&user, userId)
	return user
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash a password")
	}
	return string(hash)
}
