package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"

	"github.com/gocql/gocql"
	_ "github.com/lib/pq"
	"github.com/scylladb/gocqlx/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/turao/topics/config"

	v1 "github.com/turao/topics/api/channels/v1"
	usersV1 "github.com/turao/topics/api/users/v1"
	userrepository "github.com/turao/topics/users/repository/user"
	userservice "github.com/turao/topics/users/service/user"

	messagerepository "github.com/turao/topics/messages/repository/message"
	messagesserver "github.com/turao/topics/messages/server"
	messageservice "github.com/turao/topics/messages/service/message"
	messagespb "github.com/turao/topics/proto/messages"

	channelrepository "github.com/turao/topics/channels/repository/channel"
	channelservice "github.com/turao/topics/channels/service/channel"
)

func main() {
	// users()
	messages()
	// channels()
}

func messages() {
	cfg := config.CassandraConfig{
		Host:     "localhost",
		Port:     9042,
		Keyspace: "messages",
	}

	cluster := gocql.NewCluster(fmt.Sprintf("%s:%v", cfg.Host, cfg.Port))
	cluster.Keyspace = cfg.Keyspace
	session, err := gocqlx.WrapSession(cluster.CreateSession())
	if err != nil {
		log.Fatalln(err)
	}
	defer session.Close()

	repository, err := messagerepository.NewRepository(session)
	if err != nil {
		log.Fatalln(err)
	}

	service, err := messageservice.NewService(repository)
	if err != nil {
		log.Fatalln(err)
	}

	registrar := grpc.NewServer()
	server, err := messagesserver.NewServer(service)
	if err != nil {
		log.Fatalln(err)
	}
	messagespb.RegisterMessagesServer(registrar, server)
	reflection.Register(registrar)

	listener, err := net.Listen("tcp", "localhost:8001")
	if err != nil {
		log.Fatalln(err)
	}

	if err := registrar.Serve(listener); err != nil {
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
	cfg := config.CassandraConfig{
		Host:     "localhost",
		Port:     9042,
		Keyspace: "channels",
	}

	cluster := gocql.NewCluster(fmt.Sprintf("%s:%v", cfg.Host, cfg.Port))
	cluster.Keyspace = cfg.Keyspace
	session, err := gocqlx.WrapSession(cluster.CreateSession())
	if err != nil {
		log.Fatalln(err)
	}
	defer session.Close()

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
