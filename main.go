package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/turao/topics/config"
	"github.com/turao/topics/lib/grpc/interceptor"

	userspb "github.com/turao/topics/proto/users"
	usersV1 "github.com/turao/topics/users/api/v1"
	userrepository "github.com/turao/topics/users/repository/user"
	usersserver "github.com/turao/topics/users/server"
	userservice "github.com/turao/topics/users/service/user"
)

func main() {
	users()
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

	listener, err := net.Listen("tcp", "localhost:9090")
	if err != nil {
		log.Fatalln(err)
	}

	go func() {
		if err := http.ListenAndServe("localhost:8081", nil); err != nil {
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
