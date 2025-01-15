package models

type User struct {
	ID       uint    `gorm:"primaryKey" json:"id"`
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Password string  `json:"-"`
	Profile  Profile `gorm:"constraint:OnDelete:CASCADE;" json:"profile"`
}

type Profile struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	UserID   uint   `gorm:"unique;not null" json:"user_id"`
	Bio      string `json:"bio"`
	Location string `json:"location"`
}
