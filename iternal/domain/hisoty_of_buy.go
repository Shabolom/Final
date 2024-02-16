package domain

import (
	"github.com/gofrs/uuid"
	"time"
)

type History struct {
	UserID      uuid.UUID `gorm:"colum:user-id"`
	CreatedAt   time.Time `gorm:"colum:created-at"`
	Sum         int       `gorm:"colum:sum; type:int"`
	OrderNumber int       `gorm:"colum:order-number; type:int"`
}
