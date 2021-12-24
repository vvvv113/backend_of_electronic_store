package user

import (
	"backend/internal/entities"
	"backend/logger"
)

type driver interface {
	Create(obj interface{}) error
	FindAll(obj interface{}) error
	FindByParameters(searchObj interface{}, obj interface{}, isAll bool) error
	FindByID(ID int, obj interface{}) error
}

type database struct {
	d driver
}

func New(dbHandler driver) *database {
	return &database{
		d: dbHandler,
	}
}

func (db *database) InsertUser(user entities.User) error {
	user.Role = entities.NoAdmin
	return db.d.Create(&user)
}

func (db *database) QueryUser(userID int) (entities.User, error) {
	var result entities.User
	err := db.d.FindByID(userID, &result)
	if err != nil {
		logger.Error.Printf("Error during getting product: %s", err.Error())
		return entities.User{}, err
	}
	return result, nil
}

func (db *database) FindUserByEmail(email string) (entities.User, error) {
	type Query struct {
		Email string
	}
	query := Query{
		Email: email,
	}

	var result entities.User
	err := db.d.FindByParameters(&query, &result, false)
	if err != nil {
		logger.Error.Printf("Error during getting product: %s", err.Error())
		return entities.User{}, err
	}
	return result, nil
}
