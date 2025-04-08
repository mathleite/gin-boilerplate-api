package database

import (
	"time"

	"github.com/google/uuid"
)

type BaseTable struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at" gorm:"default:now()"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:now()"`
}
