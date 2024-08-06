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

func TestJobRepositoryDbInsert(t *testing.T) {
	db := database.NewDbTest()

	video := domain.NewVideo()
	video.ID = uuid.NewString()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	repo := repositories.VideoRepositoryDb{Db: db}
	repo.Insert(video)

	job, err := domain.NewJob("output_path", "Pending", video)
	require.Nil(t, err)

	repoJob := repositories.JobRepositoryDb{Db: db}
	repoJob.Insert(job)

	j, err := repoJob.Find(job.ID)
	require.NotEmpty(t, j.ID)
	require.Nil(t, err)
	require.Equal(t, j.ID, job.ID)
	require.Equal(t, j.VideoID, video.ID)
}

func TestJobRepositoryDbUpdate(t *testing.T) {
	db := database.NewDbTest()

	video := domain.NewVideo()
	video.ID = uuid.NewString()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	repo := repositories.VideoRepositoryDb{Db: db}
	repo.Insert(video)

	job, err := domain.NewJob("output_path", "Pending", video)
	require.Nil(t, err)

	repoJob := repositories.JobRepositoryDb{Db: db}
	repoJob.Insert(job)

	job.Status = "Complete"

	repoJob.Update(job)

	j, err := repoJob.Find(job.ID)
	require.NotEmpty(t, j.ID)
	require.Nil(t, err)
	require.Equal(t, j.Status, job.Status)
}
