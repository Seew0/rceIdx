package logic

import (
	"errors"
	"fmt"

	"github.com/seew0/rceIdx/domain/models"
	dao "github.com/seew0/rceIdx/internal/storage/DAO"
)

type Logic struct {
	dao *dao.DAO
}

func NewLogic(dao *dao.DAO) *Logic {
	return &Logic{
		dao: dao,
	}
}

func (l *Logic) InsertUser(user models.User) error {
	fmt.Println("sdahjhdjas")
	if user.Username == "" {
		return errors.New("username cannot be empty")
	}

	if !user.Password.IsValid() {
		return errors.New("password is not valid")
	}

	if !user.Email.IsValid() {
		return errors.New("email is not valid")
	}

	err := l.dao.MongoInsertUser(user)
	if err != nil {
		return errors.New("error while inserting user to database: " + err.Error())
	}

	return nil
}
