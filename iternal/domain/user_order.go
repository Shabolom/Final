package domain

import "github.com/gofrs/uuid"

type UserOrder struct {
	Base
	UserID    uuid.UUID `gorm:"colum:user_id"`
	UserOrder int64     `gorm:"colum:user_order; type:int"`
}
