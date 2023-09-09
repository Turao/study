package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/gocql/gocql"
	"github.com/gofrs/uuid"
	_ "github.com/lib/pq"
	"github.com/scylladb/gocqlx/v2"

	"github.com/turao/topics/config"

	v1 "github.com/turao/topics/api/channels/v1"
	usersV1 "github.com/turao/topics/api/users/v1"
	userrepository "github.com/turao/topics/users/repository/user"
	userservice "github.com/turao/topics/users/service/user"

	messagesV1 "github.com/turao/topics/api/messages/v1"
	messagerepository "github.com/turao/topics/messages/repository/message"
	messageservice "github.com/turao/topics/messages/service/message"

	channelrepository "github.com/turao/topics/channels/repository/channel"
	channelservice "github.com/turao/topics/channels/service/channel"
)

func main() {
	// users()
	messages()
	// channels()
}

func messages() {
	cluster := gocql.NewCluster("localhost:9042")
	cluster.Keyspace = "messages"
	session, err := gocqlx.WrapSession(cluster.CreateSession())
	if err != nil {
		log.Fatalln(err)
	}

	repository, err := messagerepository.NewRepository(session)
	if err != nil {
		log.Fatalln(err)
	}

	service, err := messageservice.NewService(repository)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = service.SendMessage(
		context.Background(),
		messagesV1.SendMessageRequest{
			AuthorID:  uuid.Must(uuid.NewV4()).String(),
			Content:   "this is my content",
			ChannelID: "outages",
		},
	)
	if err != nil {
		log.Fatalln(err)
	}

	res, err := service.GetMessages(
		context.Background(),
		messagesV1.GetMessagesRequest{
			ChannelID: "outages",
		},
	)
	if err != nil {
		log.Fatalln(err)
	}

	encoded, err := json.MarshalIndent(res, "", " ")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(encoded))
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

	repository, err := userrepository.NewRepository(database)
	if err != nil {
		log.Fatal(err)
	}

	service, err := userservice.NewService(repository)
	if err != nil {
		log.Fatal(err)
	}

	_, err = service.RegisterUser(
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

func channels() {
	cluster := gocql.NewCluster("localhost:9042")
	cluster.Keyspace = "channels"
	session, err := gocqlx.WrapSession(cluster.CreateSession())
	if err != nil {
		log.Fatalln(err)
	}

	repository, err := channelrepository.NewRepository(session)
	if err != nil {
		log.Fatalln(err)
	}

	service, err := channelservice.NewService(repository)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = service.CreateChannel(
		context.Background(),
		v1.CreateChannelRequest{
			Name:    "tech-support",
			Tenancy: "tenancy/test",
		},
	)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = service.DeleteChannel(
		context.Background(),
		v1.DeleteChannelRequest{
			ID: "969388f9-6199-402d-b550-55e87013f85a",
		},
	)
	if err != nil {
		log.Fatalln(err)
	}
}
