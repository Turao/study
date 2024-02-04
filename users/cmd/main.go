package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/cobra"
	"github.com/turao/topics/config"
	"github.com/turao/topics/metadata"

	// usersV1 "github.com/turao/topics/users/api/v1"
	v1 "github.com/turao/topics/users/api/v1"
	userrepository "github.com/turao/topics/users/repository/user"
	userservice "github.com/turao/topics/users/service/user"
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

	repository, err := userrepository.NewRepository(database)
	if err != nil {
		log.Fatal(err)
	}

	service, err := userservice.NewService(repository)
	if err != nil {
		log.Fatal(err)
	}

	// prepare commands
	RootCommand := cobra.Command{
		Use: "users",
	}

	// flags
	var userID string
	var email string
	var firstName string
	var lastName string
	var tenancy string

	RegisterUserCommand := cobra.Command{
		Use:   "register",
		Short: "registers a new user",
		Run: func(cmd *cobra.Command, args []string) {
			_, err := service.RegisterUser(
				context.Background(),
				v1.RegisteUserRequest{
					Email:     email,
					FirstName: firstName,
					LastName:  lastName,
					Tenancy:   tenancy,
				},
			)
			if err != nil {
				log.Fatalln(err)
			}
		},
	}
	RegisterUserCommand.PersistentFlags().StringVar(&email, "email", "", "user email")
	RegisterUserCommand.PersistentFlags().StringVar(&firstName, "first-name", "", "user first name")
	RegisterUserCommand.PersistentFlags().StringVar(&lastName, "last-name", "", "user last name")
	RegisterUserCommand.PersistentFlags().StringVar(&tenancy, "tenancy", metadata.TenancyTesting.String(), "user tenancy")
	RootCommand.AddCommand(&RegisterUserCommand)

	DeleteUserCommand := cobra.Command{
		Use:   "delete",
		Short: "deletes a new user",
		Run: func(cmd *cobra.Command, args []string) {
			_, err := service.DeleteUser(
				context.Background(),
				v1.DeleteUserRequest{
					ID: userID,
				},
			)
			if err != nil {
				log.Fatalln(err)
			}
		},
	}
	DeleteUserCommand.PersistentFlags().StringVar(&userID, "user-id", "", "user id")
	RootCommand.AddCommand(&DeleteUserCommand)

}
