package models

type User struct {
	ID       uint    `gorm:"primarykey" json:"id"`
	Name     string  `gorm:"size:100;not null" json:"name"`
	Email    string  `gorm:"size:100;not null" json:"email"`
	Password string  `gorm:"size:100;not null" json:"password"`
	Profile  Profile `gorm:"constraint:OnDelete:CASCADE;" json:"profile"`
}

type Profile struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	UserID   uint   `gorm:"unique;not null" json:"user_id"`
	Bio      string `gorm:"size:255" json:"bio"`
	Location string `gorm:"size:100" json:"location"`
}
