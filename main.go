package main

import (
	"context"
	"log"

	v1 "github.com/turao/topics/api/v1"
	userRepository "github.com/turao/topics/users/repository/user"
	userService "github.com/turao/topics/users/service/user"
)

func main() {
	userRepo, err := userRepository.NewRepository()
	if err != nil {
		log.Fatal(err)
	}

	userSvc, err := userService.NewService(userRepo)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	res, err := userSvc.RegisterUser(ctx, v1.RegisteUserRequest{
		Email:     "john@doe.com",
		FirstName: "john",
		LastName:  "doe",
		Tenancy:   "tenancy/test",
	})
	if err != nil {
		log.Fatal(err)
	}

	_, err = userSvc.DeleteUser(ctx, v1.DeleteUserRequest{
		ID: "fake-user-id",
	})
	if err != nil {
		log.Fatal(err)
	}

	userinfo, err := userSvc.GetUserInfo(ctx, v1.GetUserInfoRequest{
		ID: res.ID,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("userinfo", userinfo)
}
