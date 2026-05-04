package repository

import (
	"encoding/json"
	"os"
	"path/filepath"
	"robin/internal/model"
)

type SongRepository struct {
	DirPath string
}

func NewSongRepository(path string) *SongRepository {
	return &SongRepository{DirPath: path}
}

func (r *SongRepository) FetchAllSongs() ([]model.Song, error) {
	var songs []model.Song
	files, err := os.ReadDir(r.DirPath)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".json" {
			data, _ := os.ReadFile(filepath.Join(r.DirPath, file.Name()))
			var song model.Song
			if err := json.Unmarshal(data, &song); err == nil {
				songs = append(songs, song)
			}
		}
	}
	return songs, nil
}