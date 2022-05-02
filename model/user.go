package model

type User struct {
	BaseModel
	Mobile   string `gorm:"type:varchar(11);not null" json:"mobile"`
	Username string `gorm:"type:varchar(11);not null" json:"username"`
	Password string `gorm:"type:varchar(100);not null" json:"password"`
}
