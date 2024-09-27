package model

type Provider struct {
	ID   int    `gorm:"column:id; primary_key; not null" json:"id"`
	Name string `gorm:"column:name;" json:"name"`
	BaseModel
}
