package domain

import "github.com/gofrs/uuid"

type UserOrder struct {
	Base
	UserID uuid.UUID `gorm:"colum:user_id"`
	Order  int64     `gorm:"colum:order; type:int"`
}
