package models

type DLModel struct {
	ID    int64  `gorm:"primaryKey;autoIncrement"`
	Code  string `gorm:"column:code;uniqueIndex"`
	Title string `gorm:"column:title;uniqueIndex"`
}
