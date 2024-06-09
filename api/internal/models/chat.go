package models

type Chat struct {
	ID       int    `json:"id" gorm:"primaryKey;autoIncrement:false"`
	Username string `json:"username"`
	Age      int    `json:"age"`
	Location string `json:"location"`
	Bio      string `json:"bio"`
}
