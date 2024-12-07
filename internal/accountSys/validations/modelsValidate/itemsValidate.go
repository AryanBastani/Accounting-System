package modelsValidate

import (
	"Final/internal/accountSys/consts"
	"Final/internal/accountSys/models"
	"Final/internal/accountSys/utils"
	"errors"

	"gorm.io/gorm"
)

func isItemsLenValid(len int) bool {
	return len >= 2 && len <= 500
}

func isCreditAndDebitValid(credit, debit int32) bool {
	return (credit > 0 && debit == 0) ||
		(credit == 0 && debit > 0)
}

func validateSlAndDlInItem(slId, dlId int64, db *gorm.DB) error {
	var sl models.SLModel
	var dl models.DLModel

	if !utils.IsItInDb(slId, &sl, db) {
		return errors.New(consts.MODEL_NOT_FOUND)
	}
	utils.GetModelFromDb(slId, &sl, db)

	if !sl.HasDl {
		if dlId != 0 {
			return errors.New(consts.SL_HASNT_DL_BUT_DL_ID_ISNT_NIL)
		}
	} else {
		if dlId == 0 {
			return errors.New(consts.SL_HAS_DL_BUT_DL_ID_IS_NIL)
		}
		if !utils.IsItInDb(dlId, &dl, db) {
			return errors.New(consts.MODEL_NOT_FOUND)
		}
	}

	return nil
}

func ValidateItems(items *[]*models.VoucherItem, db *gorm.DB) error {
	if !isItemsLenValid(len(*items)) {
		return errors.New(consts.INVALID_NUM_OF_ITEMS)
	}

	balanceVal := int32(0)
	for _, item := range *items {
		if item.SL_ID == 0 {
			return errors.New(consts.EMPTY_SL_ID)
		}
		if !isCreditAndDebitValid(item.CREDIT, item.DEBIT) {
			return errors.New(consts.INVALID_CREDIT_DEBIT)
		}

		if item.CREDIT > 0 {
			balanceVal += item.CREDIT
		} else {
			balanceVal -= item.DEBIT
		}

		if err := validateSlAndDlInItem(item.SL_ID, item.DL_ID, db); err != nil {
			return err
		}
	}

	if balanceVal != 0 {
		return errors.New(consts.NOT_BALANCED)
	}

	return nil
}
