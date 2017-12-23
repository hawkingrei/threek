package model

type Room struct {
	Rid      int64 `gorm:"AUTO_INCREMENT;primary_key"`
	Lord     string
	Loyalist string
	Quisling string
	Thief    string
}
