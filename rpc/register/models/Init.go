package models

import (
	"github.com/go-redis/redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"rj97807_work_serve/rpc/register/internal/config"
)

// InitGorm 初始化gorm
func InitGorm(dataSource string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dataSource), &gorm.Config{
		PrepareStmt: true, //创建缓存
	})
	if err != nil {
		log.Panicf("网页库新建错误:%v", err)
		return nil
	}
	return db
}

// InitRedis 初始化gorm缓存库
func InitRedis(c config.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     c.Redis.Host,
		Password: c.Redis.Password, // no password set
		DB:       0,                // use default DB
	})
}
