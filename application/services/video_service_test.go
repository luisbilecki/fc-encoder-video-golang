package services_test

import (
	"log"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/luisbilecki/fc-encoder-video-golang/application/repositories"
	"github.com/luisbilecki/fc-encoder-video-golang/application/services"
	"github.com/luisbilecki/fc-encoder-video-golang/domain"
	"github.com/luisbilecki/fc-encoder-video-golang/framework/database"

	"github.com/stretchr/testify/require"
)

func init() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func prepare() (*domain.Video, repositories.VideoRepositoryDb) {
	db := database.NewDbTest()

	video := domain.NewVideo()
	video.ID = uuid.NewString()
	video.FilePath = "exemplo.mp4"
	video.CreatedAt = time.Now()

	repo := repositories.VideoRepositoryDb{Db: db}

	return video, repo
}

func TestVideoServiceDownload(t *testing.T) {
	video, repo := prepare()

	videoService := services.NewVideoService()
	videoService.Video = video
	videoService.VideoRepository = repo

	err := videoService.Download("fc-video-encoder-example")
	require.Nil(t, err)

	err = videoService.Fragment()
	require.Nil(t, err)

	err = videoService.Encode()
	require.Nil(t, err)

	err = videoService.Finish()
	require.Nil(t, err)
}
