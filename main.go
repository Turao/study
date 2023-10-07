package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gocql/gocql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/scylladb/gocqlx/v2"
	"github.com/surrealdb/surrealdb.go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/turao/topics/config"

	channelsV1 "github.com/turao/topics/channels/api/v1"
	usersV1 "github.com/turao/topics/users/api/v1"
	userrepository "github.com/turao/topics/users/repository/user"
	userservice "github.com/turao/topics/users/service/user"

	messagerepository "github.com/turao/topics/messages/repository/message"
	messagesserver "github.com/turao/topics/messages/server"
	messageservice "github.com/turao/topics/messages/service/message"
	messagespb "github.com/turao/topics/proto/messages"

	cassandrachannelrepository "github.com/turao/topics/channels/repository/channel/cassandra"
	mysqlchannelrepository "github.com/turao/topics/channels/repository/channel/mysql"
	psotgreschannelrepository "github.com/turao/topics/channels/repository/channel/postgres"
	surrealdbchannelrepository "github.com/turao/topics/channels/repository/channel/surrealdb"
	channelservice "github.com/turao/topics/channels/service/channel"
)

func main() {
	// users()
	// messages()
	// channelsCassandra()
	// channelsMySQL()
	// channelsSurrealDB()
	channelsPostgres()
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

func channelsCassandra() {
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

	repository, err := cassandrachannelrepository.NewRepository(session)
	if err != nil {
		log.Fatalln(err)
	}

	service, err := channelservice.NewService(repository)
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
			ID: "969388f9-6199-402d-b550-55e87013f85a",
		},
	)
	if err != nil {
		log.Fatalln(err)
	}
}

func channelsMySQL() {
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

	repository, err := mysqlchannelrepository.NewRepository(database)
	if err != nil {
		log.Fatalln(err)
	}

	service, err := channelservice.NewService(repository)
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
			ID: "9aea047d-b456-4d30-ba5d-3141f02cc4f2",
		},
	)
	if err != nil {
		log.Fatalln(err)
	}
}

func channelsSurrealDB() {
	cfg := config.SurrealDBConfig{
		Host:      "localhost",
		Port:      8000,
		Namespace: "channels",
		Database:  "channels",
		User:      "root",
		Password:  "root",
	}

	database, err := surrealdb.New(
		fmt.Sprintf("ws://%s:%v/rpc", cfg.Host, cfg.Port),
	)
	if err != nil {
		log.Fatalln(err)
	}
	defer database.Close()

	_, err = database.Signin(map[string]string{
		"user": cfg.User,
		"pass": cfg.Password,
	})
	if err != nil {
		log.Fatalln(err)
	}

	_, err = database.Use("channels", cfg.Database)
	if err != nil {
		log.Fatalln(err)
	}

	repository, err := surrealdbchannelrepository.NewRepository(database)
	if err != nil {
		log.Fatalln(err)
	}

	service, err := channelservice.NewService(repository)
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
			ID: "eafde1c1-67e1-43a1-b19e-c2f42e213733",
		},
	)
	if err != nil {
		log.Fatalln(err)
	}
}

func channelsPostgres() {
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

	repository, err := psotgreschannelrepository.NewRepository(database)
	if err != nil {
		log.Fatalln(err)
	}

	service, err := channelservice.NewService(repository)
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
			ID: "3eecad40-a879-4848-831b-8685f2d284fc",
		},
	)
	if err != nil {
		log.Fatalln(err)
	}
}
