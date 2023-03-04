package main

import (
	"log"

	userRepository "github.com/turao/topics/users/repository/user"
	userService "github.com/turao/topics/users/service/user"

	fileRepository "github.com/turao/topics/files/repository/file"
	fileService "github.com/turao/topics/files/service/file"

	movieRepository "github.com/turao/topics/movies/repository/movie"
	movieService "github.com/turao/topics/movies/service/movie"
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

	movieRepo, err := movieRepository.NewRepository()
	if err != nil {
		log.Fatal(err)
	}

	movieSvc, err := movieService.NewService(movieRepo)
	if err != nil {
		log.Fatal(err)
	}

	fileRepo, err := fileRepository.NewRepository()
	if err != nil {
		log.Fatal(err)
	}

	fileSvc, err := fileService.NewService(fileRepo)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(userSvc, movieSvc, fileSvc)
}
