package model

type Room struct {
	Rid      int64 `gorm:"AUTO_INCREMENT;primary_key"`
	Lord     User
	Loyalist User
	Quisling User
	Thief    User
}
