package domain

import (
	"time"

	valid "github.com/asaskevich/govalidator"
	"github.com/google/uuid"
)

func init() {
	valid.SetFieldsRequiredByDefault(true)
}

type Job struct {
	ID               string    `valid:"uuidv4"`
	OutputBucketPath string    `valid:"notnull"`
	Status           string    `valid:"notnull"`
	Video            *Video    `valid:"-"`
	VideoID          string    `valid:"-"`
	Error            string    `valid:"-"`
	CreatedAt        time.Time `valid:"-"`
	UpdatedAt        time.Time `valid:"-"`
}

func NewJob(output string, status string, video *Video) (*Job, error) {
	job := Job{
		OutputBucketPath: output,
		Status:           status,
		Video:            video,
	}
	job.prepare()
	err := job.Validate()
	if err != nil {
		return nil, err
	}
	return &job, nil

}

func (job *Job) prepare() {
	job.ID = uuid.NewString()
	job.CreatedAt = time.Now()
	job.UpdatedAt = time.Now()
}

func (job *Job) Validate() error {
	_, err := valid.ValidateStruct(job)
	if err != nil {
		return err
	}
	return nil
}
