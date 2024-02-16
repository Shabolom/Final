package models

import "time"

type ReqOrders struct {
	Order int64     `json:"order"`
	When  time.Time `json:"when"`
}
