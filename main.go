package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/gocql/gocql"
	"github.com/gofrs/uuid"
	_ "github.com/lib/pq"

	"github.com/turao/topics/config"

	usersV1 "github.com/turao/topics/api/users/v1"
	userRepository "github.com/turao/topics/users/repository/user"
	userService "github.com/turao/topics/users/service/user"

	messagessV1 "github.com/turao/topics/api/messages/v1"
	messageRepository "github.com/turao/topics/messages/repository/message"
	messageService "github.com/turao/topics/messages/service/message"
)

func main() {
	// users()
	messages()
}

func messages() {
	cluster := gocql.NewCluster("localhost:9042")
	cluster.Keyspace = "messages"
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatalln(err)
	}

	messageRepo, err := messageRepository.NewRepository(session)
	if err != nil {
		log.Fatalln(err)
	}

	messageSvc, err := messageService.NewService(messageRepo)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = messageSvc.SendMessage(
		context.Background(),
		messagessV1.SendMessageRequest{
			Author:  uuid.Must(uuid.NewV4()).String(),
			Content: "this is my content",
			Channel: "outages",
		},
	)
	if err != nil {
		log.Fatalln(err)
	}
}

func users() {
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
		usersV1.RegisteUserRequest{
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
