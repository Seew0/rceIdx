package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/seew0/rceIdx/domain/models"
)

type MongoStorer interface {
	InsertUser(models.User) error
	GetUser(string) (string, error)
}

type RedisStorer interface {
	Get(string) (models.RespStore, error)
	Store(string, models.RespStore) error
}

type Handlers interface {
	GetHealth(*gin.Context)
}
