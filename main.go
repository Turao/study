package main

import (
	"context"
	"log"

	v1 "github.com/turao/topics/api/v1"
	userRepository "github.com/turao/topics/users/repository/user"
	userService "github.com/turao/topics/users/service/user"

	movieRepository "github.com/turao/topics/movies/repository/movie"
	movieService "github.com/turao/topics/movies/service/movie"
)

func main() {
	testMovieService()
}

func testMovieService() {
	movieRepo, err := movieRepository.NewRepository()
	if err != nil {
		log.Fatal(err)
	}

	movieSvc, err := movieService.NewService(movieRepo)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	res, err := movieSvc.RegisterMovie(ctx, v1.RegisterMovieRequest{
		Title:   "John Wick",
		URI:     "uri-example",
		Tenancy: "tenancy/test",
	})
	if err != nil {
		log.Fatal(err)
	}

	_, err = movieSvc.DeleteMovie(ctx, v1.DeleteMovieRequest{
		ID: res.ID,
	})
	if err != nil {
		log.Fatal(err)
	}

	movieinfo, err := movieSvc.GetMovie(ctx, v1.GetMovieRequest{
		ID: res.ID,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(movieinfo)

	_, err = movieSvc.DownloadMovie(ctx, v1.DownloadMovieRequest{
		ID: res.ID,
	})
	if err != nil {
		log.Fatal(err)
	}

	movieinfos, err := movieSvc.ListMovies(ctx, v1.ListMoviesRequest{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(movieinfos)
}

func testUserService() {
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
		ID: res.ID,
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
