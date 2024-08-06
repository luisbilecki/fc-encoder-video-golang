package domain

import (
	"time"

	valid "github.com/asaskevich/govalidator"
)

type Video struct {
	ID         string    `valid:"uuidv4"`
	ResourceID string    `valid:"notnull"`
	FilePath   string    `valid:"notnull"`
	CreatedAt  time.Time `valid:"-"`
	Jobs       []*Job    `valid:"-"`
}

func init() {
	valid.SetFieldsRequiredByDefault(true)
}

func NewVideo() *Video {
	return &Video{}
}

func (video *Video) Validate() error {
	_, err := valid.ValidateStruct(video)
	if err != nil {
		return err
	}
	return nil
}
