package domain

type Balance struct {
	LoyalTokens     int `gorm:"colum:tokens; type:int"`
	UsedLoyalTokens int `gorm:"colum:used_tokens; type:int"`
}
