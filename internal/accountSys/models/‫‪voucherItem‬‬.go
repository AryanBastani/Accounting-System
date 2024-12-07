package models

type VoucherItem struct {
	ID       int64 `gorm:"primaryKey;autoIncrement"`
	DL_ID    int64 `gorm:"column:dl_id"`
	SL_ID    int64 `gorm:"column:sl_id"`
	DEBIT    int32 `gorm:"column:debit"`
	CREDIT   int32 `gorm:"column:credit"`
	Vouch_ID int64 `gorm:"column:vouch_id"`
}
