package database

import (
	"time"
)

type File struct {
	ID         uint      `gorm:"primaryKey"`
	Name       string    `json:"name"`
	Path       string    `json:"url" gorm:"unique"` //change this later
	CreatedAt  time.Time `json:"createdAt"`
	ModifiedAt time.Time `json:"modifiedAt"`
}
