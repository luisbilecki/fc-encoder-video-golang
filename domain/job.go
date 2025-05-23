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
	ID               string    `json:"job_id" valid:"uuidv4" gorm:"type:uuid;primary_key"`
	OutputBucketPath string    `json:"output_bucket_path" valid:"notnull"`
	Status           string    `json:"status" valid:"notnull"`
	Video            *Video    `json:"video" valid:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	VideoID          string    `json:"-" valid:"-" gorm:"column:video_id;type:uuid;notnull"`
	Error            string    `valid:"-"`
	CreatedAt        time.Time `json:"created_at" valid:"-"`
	UpdatedAt        time.Time `json:"updated_at" valid:"-"`
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
