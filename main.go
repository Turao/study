package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	v1 "github.com/turao/topics/api/users/v1"
	"github.com/turao/topics/config"

	userRepository "github.com/turao/topics/users/repository/user"
	userService "github.com/turao/topics/users/service/user"
)

func main() {
	userscfg := config.Users{
		DatabaseConfig: config.PostgresConfig{
			Host:     "localhost",
			Port:     5432,
			Database: "database",
			User:     "pguser",
			Password: "pwd",
		},
	}

	database, err := sql.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			userscfg.DatabaseConfig.Host,
			userscfg.DatabaseConfig.Port,
			userscfg.DatabaseConfig.User,
			userscfg.DatabaseConfig.Password,
			userscfg.DatabaseConfig.Database,
		),
	)
	if err != nil {
		log.Fatalln(err)
	}
	defer database.Close()

	// sql connections are lazy loaded. call ping to make sure our database connects
	err = database.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	userRepo, err := userRepository.NewRepository(database)
	if err != nil {
		log.Fatal(err)
	}

	userSvc, err := userService.NewService(userRepo)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(userSvc)
	_, err = userSvc.RegisterUser(
		context.Background(),
		v1.RegisteUserRequest{
			Email:     "example@domain.com",
			FirstName: "john",
			LastName:  "cleese",
			Tenancy:   "tenancy/test",
		},
	)
	if err != nil {
		log.Fatalln(err)
	}
}
