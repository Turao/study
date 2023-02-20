package main

import (
	"context"
	"log"

	v1 "github.com/turao/topics/api/v1"
	userRepository "github.com/turao/topics/users/repository/user"
	userService "github.com/turao/topics/users/service/user"

	movieRepository "github.com/turao/topics/movies/repository/movie"
	movieService "github.com/turao/topics/movies/service/movie"

	"github.com/turao/topics/movies/entity/chunk"
	"github.com/turao/topics/movies/entity/movie"
	chunkRepository "github.com/turao/topics/movies/repository/chunk"
)

func main() {
	testMovieService()
}

func testChunks() {
	chunkRepo, err := chunkRepository.NewRepository()
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()

	chunkcfg, errs := chunk.NewConfig(
		chunk.WithMovieID("movie"),
		chunk.WithURI("example-uri"),
	)
	if len(errs) > 0 {
		log.Fatal(errs)
	}

	chunk := chunk.NewChunk(chunkcfg)
	err = chunkRepo.Save(ctx, chunk)
	if err != nil {
		log.Fatal(err)
	}

	chunks, err := chunkRepo.FindByMovieID(ctx, movie.ID("movie"))
	if err != nil {
		log.Fatal(err)
	}

	for _, chunk := range chunks {
		log.Println(chunk)
	}
}

func testMovieService() {
	movieRepo, err := movieRepository.NewRepository()
	if err != nil {
		log.Fatal(err)
	}

	chunkRepo, err := chunkRepository.NewRepository()
	if err != nil {
		log.Fatal(err)
	}

	movieSvc, err := movieService.NewService(movieRepo, chunkRepo)
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

	_, err = movieSvc.SplitIntoChunks(ctx, v1.SplitIntoChunksRequest{
		MovieID: res.ID,
		Chunks:  10,
	})
	if err != nil {
		log.Fatal(err)
	}

	chunks, err := chunkRepo.FindByMovieID(ctx, movie.ID(res.ID))
	if err != nil {
		log.Fatal(err)
	}

	for _, chunk := range chunks {
		log.Println(chunk)
	}
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
