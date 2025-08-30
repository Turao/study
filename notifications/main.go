package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/turao/topics/config"
	"github.com/turao/topics/notifications/api/v1"
	notificationbuilder "github.com/turao/topics/notifications/builder"
	"github.com/turao/topics/notifications/builder/confirmation"
	notificationrepository "github.com/turao/topics/notifications/repository/notification"
	notificationsender "github.com/turao/topics/notifications/sender"
	notificationservice "github.com/turao/topics/notifications/service/notification"
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
	repository, err := notificationrepository.NewRepository(database)

	sender := notificationsender.NewSender()
	builder := notificationbuilder.NewBuilder(
		confirmation.NewBuilder(),
	)
	if err != nil {
		log.Fatalln(err)
	}

	service := notificationservice.NewService(builder, repository, sender)
	response, err := service.SendNotification(context.Background(), api.SendNotificationRequest{
		NotificationType: "confirmation",
	})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("notification sent", response)
}
