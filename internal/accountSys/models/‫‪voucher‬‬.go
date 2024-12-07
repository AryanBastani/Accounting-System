package models

type Voucher struct {
	ID   int64  `gorm:"primaryKey;autoIncrement"`
	Code string `gorm:"column:code;uniqueIndex"`
}
