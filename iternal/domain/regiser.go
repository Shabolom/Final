package domain

type Register struct {
	Base
	Login           string `gorm:"colum:login; type:text"`
	Password        string `gorm:"colum:password; type:text"`
	LoyalTokens     int    `gorm:"colum:tokens; type:int"`
	UsedLoyalTokens int    `gorm:"colum:used_tokens; type:int"`
}
