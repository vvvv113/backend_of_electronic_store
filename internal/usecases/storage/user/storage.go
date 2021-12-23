package user

import (
	"backend/internal/entities"
)

type repository interface {
	InsertUser(user entities.User) error
	FindUserByEmail(email string) (entities.User, error)
	QueryUser(userID int) (entities.User, error)
}

type Controller interface {
	CreateUser(user entities.User) error
	Login(email string, password string) (Credentials, error)
	GetProfile(userID int) (entities.User, error)
}

type Credentials struct {
	UserID int    `json:"user_id"`
	Token  string `json:"token"`
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

func (app *application) Login(email string, password string) (Credentials, error) {
	user, err := app.repo.FindUserByEmail(email)
	if err != nil {
		return Credentials{}, err
	}

	if password != user.Password {
		return Credentials{}, err
	}

	return Credentials{UserID: user.ID, Token: "random"}, nil
}

func (app *application) GetProfile(userID int) (entities.User, error) {
	return app.repo.QueryUser(userID)
}
