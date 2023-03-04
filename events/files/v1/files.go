package files

import "github.com/turao/topics/files/entity/file"

type FileDownloaded struct {
	ID      string `json:"fileId"`
	MovieID string `json:"movieId"`
	URI     string `json:"uri"`
}

func NewFileDownloaded(file file.File) *FileDownloaded {
	return &FileDownloaded{
		ID:      file.ID().String(),
		MovieID: file.Movie().String(),
		URI:     file.URI(),
	}
}
