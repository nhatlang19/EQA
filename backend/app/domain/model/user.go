package model

type User struct {
	ID       int    `gorm:"column:id; primary_key; not null" json:"id"`
	Active   bool   `gorm:"column:active" json:"active"`
	Password string `gorm:"column:password" json:"-"`
	Username string `gorm:"column:username;uniqueIndex" json:"username"`
	Name     string `gorm:"column:name;" json:"name"`
	Email    string `gorm:"column:email;" json:"email"`
	BaseModel
}
