package main

import (
	"fmt"
	"log"

	"github.com/turao/topics/config"

	userRepository "github.com/turao/topics/users/repository/user"
	userService "github.com/turao/topics/users/service/user"

	redis "github.com/redis/go-redis/v9"
)

func main() {
	userscfg := config.Users{
		RedisClient: config.RedisConfig{
			Host: "localhost",
			Port: 6379,
		},
	}

	redis := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%v", userscfg.RedisClient.Host, userscfg.RedisClient.Port),
		Password: userscfg.RedisClient.Password,
	})
	defer redis.Close()

	userRepo, err := userRepository.NewRepository(redis)
	if err != nil {
		log.Fatal(err)
	}

	userSvc, err := userService.NewService(userRepo)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(userSvc)
}
