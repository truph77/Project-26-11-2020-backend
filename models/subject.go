package models

//User ...
type User struct {
	ID       uint   `gorm:"primary_key" json:"id"`
	Fullname string `gorm:"size:100; not null" json:"fullname"`
	Email    string `gorm:"size:100; not null;unique" json:"email"`
	Password string `gorm:"size:50; not null" json:"password"`
	Phone    string `gorm:"size:10; not null" json:"phone"`
	Country  string `gorm:"size:100; not null" json:"country"`
	Sponsor  string `gorm:"size:100; not null" json:"sponsor"`
}

//ResponseToken ...
type ResponseToken struct {
	Token string `json: "token"`
}
