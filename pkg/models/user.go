package models

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
	Password []byte `json:"-"`
	Phone    string `json:"phone" gorm:"unique"`
	IsAdmin  bool	`json:"admin"`	

}

