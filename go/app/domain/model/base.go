package model

import (
	"github.com/google/uuid"
	"time"
)

type BaseModel struct {
	// ID        string `gorm:"primary_key"`
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
