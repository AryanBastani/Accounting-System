package models

type SLModel struct {
	ID    int64  `gorm:"primaryKey;autoIncrement"`
	Code  string `gorm:"column:code;uniqueIndex"`
	Title string `gorm:"column:title;uniqueIndex"`
	HasDl bool   `gorm:"column:has_dl;type:bool"`
}
