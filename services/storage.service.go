package services

import (
	"context"
	"errors"
	"log"
	"sync"
	"time"

	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
	"github.com/kamva/mgm/v3"
	models "github.com/kerem-ozt/GoodBlast_API/models/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongoDB() {
	// Setup the mgm default config
	err := mgm.SetDefaultConfig(nil, Config.MongodbDatabase, options.Client().ApplyURI(Config.MongodbUri))
	if err != nil {
		panic(err)
	}

	log.Println("Connected to MongoDB!")
}

var redisDefaultClient *redis.Client
var redisDefaultOnce sync.Once

var redisCache *cache.Cache
var redisCacheOnce sync.Once

func GetRedisDefaultClient() *redis.Client {
	redisDefaultOnce.Do(func() {
		redisDefaultClient = redis.NewClient(&redis.Options{
			Addr: Config.RedisDefaultAddr,
		})
	})

	return redisDefaultClient
}

func GetRedisCache() *cache.Cache {
	redisCacheOnce.Do(func() {
		redisCache = cache.New(&cache.Options{
			Redis:      GetRedisDefaultClient(),
			LocalCache: cache.NewTinyLFU(1000, time.Minute),
		})
	})

	return redisCache
}

func CheckRedisConnection() {
	redisClient := GetRedisDefaultClient()
	err := redisClient.Ping(context.Background()).Err()
	if err != nil {
		panic(err)
	}

	log.Println("Connected to Redis!")
}

func getTournamentCacheKey(userId primitive.ObjectID, tournamentId primitive.ObjectID) string {
	return "req:cache:tournament:" + userId.Hex() + ":" + tournamentId.Hex()
}

func CacheOneTournament(userId primitive.ObjectID, tournament *models.Tournament) {
	if !Config.UseRedis {
		return
	}

	tournamentCacheKey := getTournamentCacheKey(userId, tournament.ID)

	_ = GetRedisCache().Set(&cache.Item{
		Ctx:   context.TODO(),
		Key:   tournamentCacheKey,
		Value: tournament,
		TTL:   time.Minute,
	})
}

func GetTournamentFromCache(userId primitive.ObjectID, tournamentId primitive.ObjectID) (*models.Tournament, error) {
	if !Config.UseRedis {
		return nil, errors.New("no redis client, set USE_REDIS in .env")
	}

	tournament := &models.Tournament{}
	tournamentCacheKey := getTournamentCacheKey(userId, tournamentId)
	err := GetRedisCache().Get(context.TODO(), tournamentCacheKey, tournament)
	return tournament, err
}
