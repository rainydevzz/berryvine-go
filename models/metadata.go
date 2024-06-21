package models

type Metadata struct {
	File string `gorm:"primaryKey"`
	Name string
}
