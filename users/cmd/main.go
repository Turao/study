package main

import (
	"fmt"
	"log"
	"net"
	_ "net/http/pprof"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/turao/topics/config"
	"github.com/turao/topics/lib/grpc/interceptor"

	userspb "github.com/turao/topics/proto/users"
	grouprepository "github.com/turao/topics/users/repository/group"
	userrepository "github.com/turao/topics/users/repository/user"
	userstreamrepository "github.com/turao/topics/users/repository/userstream"
	usersgrpcserver "github.com/turao/topics/users/server/grpc"
	userswebserver "github.com/turao/topics/users/server/web"
	groupservice "github.com/turao/topics/users/service/group"
	userservice "github.com/turao/topics/users/service/user"
	userstreamservice "github.com/turao/topics/users/service/userstream"
)

func main() {
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

	userRepository, err := userrepository.NewRepository(database)
	if err != nil {
		log.Fatal(err)
	}

	userService, err := userservice.NewService(userRepository)
	if err != nil {
		log.Fatal(err)
	}

	groupRepository, err := grouprepository.NewRepository(database)
	if err != nil {
		log.Fatal(err)
	}

	groupService, err := groupservice.NewService(groupRepository)
	if err != nil {
		log.Fatal(err)
	}

	subscriber, err := kafka.NewSubscriber(
		kafka.SubscriberConfig{
			Brokers:     []string{"localhost:9093"},
			Unmarshaler: kafka.DefaultMarshaler{},
		},
		watermill.NewStdLogger(false, false),
	)
	if err != nil {
		log.Fatalln(err)
	}

	userStreamRepository, err := userstreamrepository.NewRepository(subscriber)
	if err != nil {
		log.Fatalln(err)
	}

	userStreamService, err := userstreamservice.New(userStreamRepository)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		userserver := userswebserver.NewServer(userService, userStreamService, config.HTTPServerConfig{Port: 7070})
		if err := userserver.ListenAndServe(); err != nil {
			log.Fatalln(err)
		}
	}()

	server, err := usersgrpcserver.NewServer(
		userService,
		groupService,
	)
	if err != nil {
		log.Fatalln(err)
	}

	registrar := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			interceptor.WithTenancyInterceptor(),
		),
	)

	userspb.RegisterUsersServer(registrar, server)
	userspb.RegisterGroupsServer(registrar, server)
	reflection.Register(registrar)

	// start TCP listener for GRPC server
	listener, err := net.Listen("tcp", "localhost:9090")
	if err != nil {
		log.Fatalln(err)
	}

	if err := registrar.Serve(listener); err != nil {
		log.Fatalln(err)
	}
}
