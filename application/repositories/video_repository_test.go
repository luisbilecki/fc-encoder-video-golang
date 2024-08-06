package repositories_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/luisbilecki/fc-encoder-video-golang/application/repositories"
	"github.com/luisbilecki/fc-encoder-video-golang/domain"
	"github.com/luisbilecki/fc-encoder-video-golang/framework/database"
	"github.com/stretchr/testify/require"
)

func TestVideoRepositoryDbInsert(t *testing.T) {
	db := database.NewDbTest()
	video := domain.NewVideo()
	video.ID = uuid.NewString()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	repo := repositories.VideoRepositoryDb{Db: db}
	repo.Insert(video)

	v, err := repo.Find(video.ID)

	require.NotEmpty(t, v.ID)
	require.Nil(t, err)
	require.Equal(t, v.ID, video.ID)
}
