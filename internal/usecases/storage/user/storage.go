package user

import (
	"backend/internal/entities"
	"backend/logger"
	"fmt"
)

type repository interface {
	InsertUser(user entities.User) error
	FindUserByEmail(email string) (entities.User, error)
	QueryUser(userID int) (entities.User, error)
}

type Controller interface {
	CreateUser(user entities.User) error
	Login(credentials Credentials) (Cookies, error)
	GetProfile(userID int) (entities.User, error)
}

type Cookies struct {
	UserID int    `json:"user_id"`
	Token  string `json:"token"`
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type application struct {
	repo repository
}

func New(repo repository) *application {
	return &application{
		repo: repo,
	}
}

func (app *application) CreateUser(user entities.User) error {
	return app.repo.InsertUser(user)
}

func (app *application) Login(credentials Credentials) (Cookies, error) {
	user, err := app.repo.FindUserByEmail(credentials.Email)
	if err != nil {
		logger.Error.Println(err)
		return Cookies{}, fmt.Errorf("Account didn't found try again")
	}

	if credentials.Password != user.Password {
		return Cookies{}, fmt.Errorf("Wrong email or password. Try again")
	}

	return Cookies{UserID: user.ID, Token: "random"}, nil
}

func (app *application) GetProfile(userID int) (entities.User, error) {
	return app.repo.QueryUser(userID)
}
