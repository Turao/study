package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gocql/gocql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/scylladb/gocqlx/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/turao/topics/config"
	"github.com/turao/topics/lib/grpc/interceptor"

	channelsV1 "github.com/turao/topics/channels/api/v1"
	userspb "github.com/turao/topics/proto/users"
	usersV1 "github.com/turao/topics/users/api/v1"
	userrepository "github.com/turao/topics/users/repository/user"
	usersserver "github.com/turao/topics/users/server"
	userservice "github.com/turao/topics/users/service/user"

	messagerepository "github.com/turao/topics/messages/repository/message"
	messagesserver "github.com/turao/topics/messages/server"
	messageservice "github.com/turao/topics/messages/service/message"
	messagespb "github.com/turao/topics/proto/messages"

	channelrepository "github.com/turao/topics/channels/repository/channel"
	membershiprepository "github.com/turao/topics/channels/repository/membership"
	channelservice "github.com/turao/topics/channels/service/channel"
)

func main() {
	// users()
	// messages()
	channels()
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
	cfg := config.PostgresConfig{
		Host:     "localhost",
		Port:     5432,
		Database: "database",
		User:     "pguser",
		Password: "pwd",
	}

	database, err := sqlx.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			cfg.Host,
			cfg.Port,
			cfg.User,
			cfg.Password,
			cfg.Database,
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

	registrar := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			interceptor.WithTenancyInterceptor(),
		),
	)

	listener, err := net.Listen("tcp", "localhost:8001")
	if err != nil {
		log.Fatalln(err)
	}

	go func() {
		if err := http.ListenAndServe("localhost:8002", nil); err != nil {
			log.Fatalln(err)
		}
	}()

	server, err := usersserver.NewServer(service)
	if err != nil {
		log.Fatalln(err)
	}
	userspb.RegisterUsersServer(registrar, server)
	reflection.Register(registrar)

	if err := registrar.Serve(listener); err != nil {
		log.Fatalln(err)
	}
}

func channels() {
	cfg := config.MySQLConfig{
		Host:     "localhost",
		Port:     3306,
		Database: "channels",
		User:     "mysqluser",
		Password: "pwd",
	}

	database, err := sqlx.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%v)/%s?parseTime=true",
			cfg.User,
			cfg.Password,
			cfg.Host,
			cfg.Port,
			cfg.Database,
		),
	)
	if err != nil {
		log.Fatalln(err)
	}
	defer database.Close()

	channelRepository, err := channelrepository.NewRepository(database)
	if err != nil {
		log.Fatalln(err)
	}

	membershipRepository, err := membershiprepository.NewRepository(database)
	if err != nil {
		log.Fatalln(err)
	}

	service, err := channelservice.NewService(
		channelRepository,
		membershipRepository,
	)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = service.CreateChannel(
		context.Background(),
		channelsV1.CreateChannelRequest{
			Name:    "tech-support",
			Tenancy: "tenancy/test",
		},
	)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = service.DeleteChannel(
		context.Background(),
		channelsV1.DeleteChannelRequest{
			ID: "66dd5b74-c54a-11ee-aab2-0242ac110002",
		},
	)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = service.JoinChannel(
		context.Background(),
		channelsV1.JoinChannelRequest{
			ChannelID: "7527542f-7bb9-46cc-b702-7b71047bbf78",
			UserID:    "7b0ba219-a020-4821-a494-9233143866f0",
		},
	)
	if err != nil {
		log.Fatalln(err)
	}
}
