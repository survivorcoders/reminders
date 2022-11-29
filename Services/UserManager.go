package Services

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"survivorcoders.com/reminders/entity"
	"survivorcoders.com/reminders/repository"
)

type UserManager struct {
	UserRepository *repository.UserProviderRepository
}

func NewUserManager(userRepository *repository.UserProviderRepository) *UserManager {
	return &UserManager{UserRepository: userRepository}
}

func (u *UserManager) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (u *UserManager) checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (u *UserManager) SignIn(username, password string) (*entity.User, error) {
	var user entity.User
	e := "invalid username or password"

	err := u.UserRepository.GetUserByUsername(username, &user)
	if err != nil {
		return nil, errors.New(e)
	}

	//username found
	if u.checkPasswordHash(password, user.Password) == false {
		return nil, errors.New(e)
	}

	return &user, nil
}

func (u *UserManager) SignUp(user *entity.User) error {
	//password hash
	hashedPassword, err := u.hashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	//user creation
	err = u.UserRepository.CreateUser(user)
	if err != nil {
		return err
	}
	//user Created
	return nil
}
