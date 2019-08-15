package store

import (
	"github.com/gomodule/redigo/redis"
	"github.com/thinkoner/thinkgo/cache"
)

var memRepository *cache.Repository
var redisRepository *cache.Repository

func init() {
	memRepository, _ = cache.Cache(cache.NewMemoryStore("store"))

	address := "127.0.0.1:6379"
	//password := "123"
	pool := &redis.Pool{
		MaxIdle:   3,
		MaxActive: 5,
		Dial: func() (conn redis.Conn, e error) {
			con, err := redis.Dial("tcp", address)
			if err != nil {
				return nil, err
			}
			/*if _, err := con.Do("AUTH",password); err != nil {//没设密码，不需认证
				con.Close()
				return nil,err
			}*/
			return con, err
		},
	}
	redisRepository, _ = cache.Cache(cache.NewRedisStore(pool, "rstore"))
}

func MemRepository() *cache.Repository {
	return memRepository
}
func RedisReposotory() *cache.Repository {
	return redisRepository
}
