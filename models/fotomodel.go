package models

type Foto struct {
	ID        int `gorm:"primaryKey"`
	Title     string
	Caption   string
	PhotoUrl  string
	UserID    int
	UserRefer Users `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}