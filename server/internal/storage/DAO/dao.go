package dao

import (
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/seew0/rceIdx/domain/interfaces"
	"github.com/seew0/rceIdx/domain/models"
	"github.com/seew0/rceIdx/internal/storage/db"
	redisDB "github.com/seew0/rceIdx/internal/storage/redis"
)

type DAO struct {
	MongoStorer interfaces.MongoStorer
	RedisStorer interfaces.RedisStorer
}

func NewDAO(mongoStorer interfaces.MongoStorer) *DAO {
	MongoDbClient, err := db.NewMongoClient(os.Getenv("MONGO_URI"), 10)
	if err != nil {
		log.Fatalf("error occured while connecting to MongoDB: %v", err)
	}

	storage := db.NewStorage(MongoDbClient, os.Getenv("MONGO_DB_NAME"), 10)

	redisClient := redisDB.NewRedisStore(os.Getenv("REDIS_ADDR"), os.Getenv("REDIS_PASSWD"))

	return &DAO{
		MongoStorer: storage,
		RedisStorer: redisClient,
	}
}

func (d *DAO) MongoInsertUser(user models.User) error {

	user.UserID = uuid.New().String()
	user.CreatedAt = time.Now()
	return d.MongoStorer.InsertUser(user)
}

func (d *DAO) RedisGetResp(uuid string) (models.RespStore, error) {
	return d.RedisStorer.Get(uuid)
}

func (d *DAO) RedisStoreResp(uuid string, resp models.RespStore) error {
	return d.RedisStorer.Store(uuid, resp)
}
