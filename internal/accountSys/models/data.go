package models

import "Final/internal/accountSys/utils"

type VoucherInsertData struct {
	*Voucher
	Items []*VoucherItem
}

type VoucherItemUpdateData struct {
	ToInserts []*VoucherItem
	ToUpdates []*VoucherItem
	ToDeletes []*int64
}

type VoucherUpdateData struct {
	*Voucher
	Items *VoucherItemUpdateData
}

func (v1 *VoucherInsertData) Equals(v2 *VoucherInsertData) bool {
	if *v1.Voucher != *v2.Voucher {
		return false
	}

	if len(v1.Items) != len(v2.Items) {
		return false
	}

	v2ItemsMap := utils.ArrayToMap(v2.Items, func(v *VoucherItem) int64 {
		return v.ID
	})

	for _, item := range v1.Items {
		if _, ok := v2ItemsMap[item.ID]; !ok {
			return false
		}
		if *item != *v2ItemsMap[item.ID] {
			return false
		}
	}

	return true
}
