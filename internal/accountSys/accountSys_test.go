package accountSys

import (
	"Final/internal/accountSys/consts"
	"Final/internal/accountSys/migration"
	"Final/internal/accountSys/models"
	"Final/internal/accountSys/utils"
	"testing"
	"time"

	"golang.org/x/exp/rand"
	"gorm.io/gorm"
)

var testDB *gorm.DB

func init() {
	testDB = migration.MigrateDatabase()
}

func createdRandomDl() (createdDl *models.DLModel) {
	createdDl = &models.DLModel{
		Code:  utils.RandomString(consts.STRINGS_LEN),
		Title: utils.RandomString(consts.STRINGS_LEN),
	}
	CreateDLModel(createdDl, testDB)

	return createdDl
}

func createdRandomSl(hasDl bool) (createdSl *models.SLModel) {
	createdSl = &models.SLModel{
		Code:  utils.RandomString(consts.STRINGS_LEN),
		Title: utils.RandomString(consts.STRINGS_LEN),
		HasDl: hasDl,
	}
	CreateSLModel(createdSl, testDB)

	return createdSl
}

func createRandomVoucher(slId, itemId1, itemId2 int64) (createdV *models.VoucherInsertData) {
	creditVal := utils.RandomInt32()
	createdV = &models.VoucherInsertData{
		Voucher: &models.Voucher{
			Code: utils.RandomString(consts.STRINGS_LEN),
		},
		Items: []*models.VoucherItem{
			{
				ID:     itemId1,
				DL_ID:  0,
				SL_ID:  slId,
				DEBIT:  0,
				CREDIT: creditVal,
			},
			{
				ID:     itemId2,
				DL_ID:  0,
				SL_ID:  slId,
				DEBIT:  creditVal,
				CREDIT: 0,
			}},
	}
	CreateVoucher(createdV, testDB)

	return createdV
}

func TestCreateDLModel_ValidArgs(t *testing.T) {
	dl := &models.DLModel{
		Code:  utils.RandomString(consts.STRINGS_LEN),
		Title: utils.RandomString(consts.STRINGS_LEN),
	}

	if _, err := CreateDLModel(dl, testDB); err != nil {
		t.Fatalf(consts.EXPECTED_NO_ERR, err)
	}
}

func TestCreateDLModel_EmptyCode(t *testing.T) {
	dl := &models.DLModel{
		Code:  "",
		Title: utils.RandomString(consts.STRINGS_LEN),
	}
	_, err := CreateDLModel(dl, testDB)
	if err == nil || err.Error() != consts.EMPTY_CODE {
		t.Fatalf(consts.EXPECTED_EMPTY_CODE_ERR, err)
	}
}

func TestCreateDLModel_EmptyTitle(t *testing.T) {
	dl := &models.DLModel{
		Code:  utils.RandomString(consts.STRINGS_LEN),
		Title: "",
	}
	_, err := CreateDLModel(dl, testDB)
	if err == nil || err.Error() != consts.EMPTY_TITLE {
		t.Fatalf(consts.EXPECTED_EMPTY_TITLE_ERR, err)
	}
}

func TestCreateDLModel_LongCode(t *testing.T) {
	dl := &models.DLModel{
		Code:  utils.RandomString(consts.STRINGS_LEN + 1),
		Title: utils.RandomString(consts.STRINGS_LEN),
	}
	_, err := CreateDLModel(dl, testDB)
	if err == nil || err.Error() != consts.LONG_CODE {
		t.Fatalf(consts.EXPECTED_LONG_CODE_ERR, err)
	}
}

func TestCreateDLModel_LongTitle(t *testing.T) {
	dl := &models.DLModel{
		Code:  utils.RandomString(consts.STRINGS_LEN),
		Title: utils.RandomString(consts.STRINGS_LEN + 1),
	}
	_, err := CreateDLModel(dl, testDB)
	if err == nil || err.Error() != consts.LONG_TITLE {
		t.Fatalf(consts.EXPECTED_LONG_TITLE_ERR, err)
	}
}

func TestUpdateDLModel_ValidArgs(t *testing.T) {
	createdDl := createdRandomDl()

	dl := &models.DLModel{
		ID:    createdDl.ID,
		Code:  utils.RandomString(consts.STRINGS_LEN),
		Title: utils.RandomString(consts.STRINGS_LEN),
	}

	if err := updateDLModel(dl, testDB); err != nil {
		t.Fatalf(consts.EXPECTED_NO_ERR, err)
	}
}

func TestUpdateDLModel_EmptyCode(t *testing.T) {
	createdDl := createdRandomDl()

	newDl := &models.DLModel{
		ID:    createdDl.ID,
		Code:  "",
		Title: utils.RandomString(consts.STRINGS_LEN),
	}
	err := updateDLModel(newDl, testDB)
	if err == nil || err.Error() != consts.EMPTY_CODE {
		t.Fatalf(consts.EXPECTED_EMPTY_CODE_ERR, err)
	}
}

func TestUpdateDLModel_EmptyTitle(t *testing.T) {
	createdDl := createdRandomDl()

	newDl := &models.DLModel{
		ID:    createdDl.ID,
		Code:  utils.RandomString(consts.STRINGS_LEN),
		Title: "",
	}
	err := updateDLModel(newDl, testDB)
	if err == nil || err.Error() != consts.EMPTY_TITLE {
		t.Fatalf(consts.EXPECTED_EMPTY_TITLE_ERR, err)
	}
}

func TestUpdateDLModel_LongCode(t *testing.T) {
	createdDl := createdRandomDl()

	newDl := &models.DLModel{
		ID:    createdDl.ID,
		Code:  utils.RandomString(consts.STRINGS_LEN + 1),
		Title: utils.RandomString(consts.STRINGS_LEN),
	}
	err := updateDLModel(newDl, testDB)
	if err == nil || err.Error() != consts.LONG_CODE {
		t.Fatalf(consts.EXPECTED_LONG_CODE_ERR, err)
	}
}

func TestUpdateDLModel_LongTitle(t *testing.T) {
	createdDl := createdRandomDl()

	newDl := &models.DLModel{
		ID:    createdDl.ID,
		Code:  utils.RandomString(consts.STRINGS_LEN),
		Title: utils.RandomString(consts.STRINGS_LEN + 1),
	}
	err := updateDLModel(newDl, testDB)
	if err == nil || err.Error() != consts.LONG_TITLE {
		t.Fatalf(consts.EXPECTED_LONG_TITLE_ERR, err)
	}
}

func TestUpdateDLModel_NotExistModel(t *testing.T) {
	dl := &models.DLModel{
		ID:    utils.RandomInt64(),
		Code:  utils.RandomString(consts.STRINGS_LEN),
		Title: utils.RandomString(consts.STRINGS_LEN),
	}

	err := updateDLModel(dl, testDB)
	if err == nil || err.Error() != consts.MODEL_NOT_FOUND {
		t.Fatalf(consts.EXPECTED_NOT_FOUND_ERR, err)
	}
}

func TestDeleteDLModel_ValidArgs(t *testing.T) {
	createdDl := createdRandomDl()

	if err := deleteDLModel(createdDl.ID, testDB); err != nil {
		t.Fatalf(consts.EXPECTED_NO_ERR, err)
	}
}

func TestDeleteDLModel_NotExistId(t *testing.T) {
	id := utils.RandomInt64()

	err := deleteDLModel(id, testDB)
	if err == nil || err.Error() != consts.MODEL_NOT_FOUND {
		t.Fatalf(consts.EXPECTED_NOT_FOUND_ERR, err)
	}
}

func TestDeleteDLModel_ReferencedDl(t *testing.T) {
	sl := createdRandomSl(true)

	dl := createdRandomDl()

	creditVal := utils.AbsInt32(utils.RandomInt32())
	data := models.VoucherInsertData{
		Voucher: &models.Voucher{
			Code: utils.RandomString(consts.STRINGS_LEN),
		},
		Items: []*models.VoucherItem{
			{
				DL_ID:  dl.ID,
				SL_ID:  sl.ID,
				DEBIT:  0,
				CREDIT: creditVal,
			},
			{
				DL_ID:  dl.ID,
				SL_ID:  sl.ID,
				DEBIT:  creditVal,
				CREDIT: 0,
			}},
	}
	if _, err := CreateVoucher(&data, testDB); err != nil {
		t.Fatalf(consts.EXPECTED_NO_ERR, err)
	}

	err := deleteDLModel(dl.ID, testDB)
	if err == nil || err.Error() != consts.THIS_DL_IS_REFERENCED {
		t.Fatalf(consts.EXPECTED_REFERENCED_ERR, err)
	}
}

func TestGetDLModel_ValidArgs(t *testing.T) {
	createdDl := createdRandomDl()

	returnedDl, err := getDLModel(createdDl.ID, testDB)
	if err != nil {
		t.Fatalf(consts.EXPECTED_NO_ERR, err)
	}
	if *returnedDl != *createdDl {
		t.Fatalf(consts.EXPECTED_DIFF_RETURNED_MODEL_ERR, *createdDl, *returnedDl)
	}
}

func TestGetDLModel_NotExistId(t *testing.T) {
	id := utils.RandomInt64()

	_, err := getDLModel(id, testDB)
	if err == nil {
		t.Fatalf(consts.EXPECTED_NOT_FOUND_ERR, err)
	}
}

func TestGetDLModel_AfterUpdate(t *testing.T) {
	createdDl := createdRandomDl()

	updatedDl := &models.DLModel{
		ID:    createdDl.ID,
		Code:  utils.RandomString(consts.STRINGS_LEN),
		Title: utils.RandomString(consts.STRINGS_LEN),
	}

	if err := updateDLModel(updatedDl, testDB); err != nil {
		t.Fatalf(consts.EXPECTED_NO_ERR, err)
	}

	returnedDl, err := getDLModel(createdDl.ID, testDB)
	if err != nil {
		t.Fatalf(consts.EXPECTED_NO_ERR, err)
	}
	if *returnedDl != *updatedDl {
		t.Fatalf(consts.EXPECTED_DIFF_RETURNED_MODEL_ERR, *createdDl, *returnedDl)
	}
}

func TestCreateSLModel_ValidArgs(t *testing.T) {
	sl := &models.SLModel{
		Code:  utils.RandomString(consts.STRINGS_LEN),
		Title: utils.RandomString(consts.STRINGS_LEN),
		HasDl: utils.RandomBool(),
	}

	if _, err := CreateSLModel(sl, testDB); err != nil {
		t.Fatalf(consts.EXPECTED_NO_ERR, err)
	}
}

func TestCreateSLModel_EmptyCode(t *testing.T) {
	sl := &models.SLModel{
		Code:  "",
		Title: utils.RandomString(consts.STRINGS_LEN),
		HasDl: utils.RandomBool(),
	}

	_, err := CreateSLModel(sl, testDB)
	if err == nil || err.Error() != consts.EMPTY_CODE {
		t.Fatalf(consts.EXPECTED_EMPTY_CODE_ERR, err)
	}
}

func TestCreateSLModel_EmptyTitle(t *testing.T) {
	sl := &models.SLModel{
		Code:  utils.RandomString(consts.STRINGS_LEN),
		Title: "",
		HasDl: utils.RandomBool(),
	}
	_, err := CreateSLModel(sl, testDB)
	if err == nil || err.Error() != consts.EMPTY_TITLE {
		t.Fatalf(consts.EXPECTED_EMPTY_TITLE_ERR, err)
	}
}

func TestCreateSLModel_LongCode(t *testing.T) {
	sl := &models.SLModel{
		Code:  utils.RandomString(consts.STRINGS_LEN + 1),
		Title: utils.RandomString(consts.STRINGS_LEN),
		HasDl: utils.RandomBool(),
	}
	_, err := CreateSLModel(sl, testDB)
	if err == nil || err.Error() != consts.LONG_CODE {
		t.Fatalf(consts.EXPECTED_LONG_CODE_ERR, err)
	}
}

func TestCreateSLModel_LongTitle(t *testing.T) {
	sl := &models.SLModel{
		Code:  utils.RandomString(consts.STRINGS_LEN),
		Title: utils.RandomString(consts.STRINGS_LEN + 1),
	}
	_, err := CreateSLModel(sl, testDB)
	if err == nil || err.Error() != consts.LONG_TITLE {
		t.Fatalf(consts.EXPECTED_LONG_TITLE_ERR, err)
	}
}

func TestUpdateSLModel_ValidArgs(t *testing.T) {
	createdSl := createdRandomSl(utils.RandomBool())

	newSl := &models.SLModel{
		ID:    createdSl.ID,
		Code:  utils.RandomString(consts.STRINGS_LEN),
		Title: utils.RandomString(consts.STRINGS_LEN),
		HasDl: utils.RandomBool(),
	}

	if err := updateSLModeldl(newSl, testDB); err != nil {
		t.Fatalf(consts.EXPECTED_NO_ERR, err)
	}
}

func TestUpdateSLModel_EmptyCode(t *testing.T) {
	createdSl := createdRandomSl(utils.RandomBool())

	newSl := &models.SLModel{
		ID:    createdSl.ID,
		Code:  "",
		Title: utils.RandomString(consts.STRINGS_LEN),
		HasDl: utils.RandomBool(),
	}
	err := updateSLModeldl(newSl, testDB)
	if err == nil || err.Error() != consts.EMPTY_CODE {
		t.Fatalf(consts.EXPECTED_EMPTY_CODE_ERR, err)
	}
}

func TestUpdateSLModel_EmptyTitle(t *testing.T) {
	createdSl := createdRandomSl(utils.RandomBool())

	newSl := &models.SLModel{
		ID:    createdSl.ID,
		Code:  utils.RandomString(consts.STRINGS_LEN),
		Title: "",
		HasDl: utils.RandomBool(),
	}
	err := updateSLModeldl(newSl, testDB)
	if err == nil || err.Error() != consts.EMPTY_TITLE {
		t.Fatalf(consts.EXPECTED_EMPTY_TITLE_ERR, err)
	}
}

func TestUpdateSLModel_LongCode(t *testing.T) {
	createdSl := createdRandomSl(utils.RandomBool())

	newSl := &models.SLModel{
		ID:    createdSl.ID,
		Code:  utils.RandomString(consts.STRINGS_LEN + 1),
		Title: utils.RandomString(consts.STRINGS_LEN),
		HasDl: utils.RandomBool(),
	}
	err := updateSLModeldl(newSl, testDB)
	if err == nil || err.Error() != consts.LONG_CODE {
		t.Fatalf(consts.EXPECTED_LONG_CODE_ERR, err)
	}
}

func TestUpdateSLModel_LongTitle(t *testing.T) {
	createdSl := createdRandomSl(utils.RandomBool())

	newSl := &models.SLModel{
		ID:    createdSl.ID,
		Code:  utils.RandomString(consts.STRINGS_LEN),
		Title: utils.RandomString(consts.STRINGS_LEN + 1),
		HasDl: utils.RandomBool(),
	}
	err := updateSLModeldl(newSl, testDB)
	if err == nil || err.Error() != consts.LONG_TITLE {
		t.Fatalf(consts.EXPECTED_LONG_TITLE_ERR, err)
	}
}

func TestUpdateSLModel_NotExistModel(t *testing.T) {
	sl := &models.SLModel{
		ID:    utils.RandomInt64(),
		Code:  utils.RandomString(consts.STRINGS_LEN),
		Title: utils.RandomString(consts.STRINGS_LEN),
		HasDl: utils.RandomBool(),
	}

	err := updateSLModeldl(sl, testDB)
	if err == nil || err.Error() != consts.MODEL_NOT_FOUND {
		t.Fatalf(consts.EXPECTED_NOT_FOUND_ERR, err)
	}
}

func TestUpdateSLModel_ReferencedSl(t *testing.T) {
	sl := createdRandomSl(false)

	creditVal := utils.AbsInt32(utils.RandomInt32())
	data := models.VoucherInsertData{
		Voucher: &models.Voucher{
			Code: utils.RandomString(consts.STRINGS_LEN),
		},
		Items: []*models.VoucherItem{
			{
				DL_ID:  0,
				SL_ID:  sl.ID,
				DEBIT:  0,
				CREDIT: creditVal,
			},
			{
				DL_ID:  0,
				SL_ID:  sl.ID,
				DEBIT:  creditVal,
				CREDIT: 0,
			}},
	}
	if _, err := CreateVoucher(&data, testDB); err != nil {
		t.Fatalf(consts.EXPECTED_NO_ERR, err)
	}

	err := updateSLModeldl(sl, testDB)
	if err == nil || err.Error() != consts.THIS_SL_IS_REFERENCED {
		t.Fatalf(consts.EXPECTED_REFERENCED_ERR, err)
	}
}

func TestDeleteSLModel_NotExistId(t *testing.T) {
	ID := utils.RandomInt64()

	err := deleteSLModel(ID, testDB)
	if err == nil || err.Error() != consts.MODEL_NOT_FOUND {
		t.Fatalf(consts.EXPECTED_NOT_FOUND_ERR, err)
	}
}

func TestDeleteSLModel_ValidArgs(t *testing.T) {
	createdSl := createdRandomSl(utils.RandomBool())

	if err := deleteSLModel(createdSl.ID, testDB); err != nil {
		t.Fatalf(consts.EXPECTED_NO_ERR, err)
	}
}

func TestDeleteSLModel_ReferencedSl(t *testing.T) {
	sl := createdRandomSl(false)

	creditVal := utils.AbsInt32(utils.RandomInt32())
	data := models.VoucherInsertData{
		Voucher: &models.Voucher{
			Code: utils.RandomString(consts.STRINGS_LEN),
		},
		Items: []*models.VoucherItem{
			{
				DL_ID:  0,
				SL_ID:  sl.ID,
				DEBIT:  0,
				CREDIT: creditVal,
			},
			{
				DL_ID:  0,
				SL_ID:  sl.ID,
				DEBIT:  creditVal,
				CREDIT: 0,
			}},
	}
	if _, err := CreateVoucher(&data, testDB); err != nil {
		t.Fatalf(consts.EXPECTED_NO_ERR, err)
	}

	err := deleteSLModel(sl.ID, testDB)
	if err == nil || err.Error() != consts.THIS_SL_IS_REFERENCED {
		t.Fatalf(consts.EXPECTED_REFERENCED_ERR, err)
	}
}

func TestGetSLModel_ValidArgs(t *testing.T) {
	createdSl := createdRandomSl(utils.RandomBool())

	returnedSl, err := getSLModel(createdSl.ID, testDB)
	if err != nil {
		t.Fatalf(consts.EXPECTED_NO_ERR, err)
	}
	if *returnedSl != *createdSl {
		t.Fatalf(consts.EXPECTED_DIFF_RETURNED_MODEL_ERR, *createdSl, *returnedSl)
	}
}

func TestGetSLModel_NotExistId(t *testing.T) {
	ID := utils.RandomInt64()

	_, err := getSLModel(ID, testDB)
	if err == nil {
		t.Fatalf(consts.EXPECTED_NOT_FOUND_ERR, err)
	}
}

func TestCreateVoucher_ValidArgsWithoutDl(t *testing.T) {
	sl := createdRandomSl(false)

	creditVal := utils.AbsInt32(utils.RandomInt32())
	data := models.VoucherInsertData{
		Voucher: &models.Voucher{
			Code: utils.RandomString(consts.STRINGS_LEN),
		},
		Items: []*models.VoucherItem{
			{
				DL_ID:  0,
				SL_ID:  sl.ID,
				DEBIT:  0,
				CREDIT: creditVal,
			},
			{
				DL_ID:  0,
				SL_ID:  sl.ID,
				DEBIT:  creditVal,
				CREDIT: 0,
			}},
	}
	if _, err := CreateVoucher(&data, testDB); err != nil {
		t.Fatalf(consts.EXPECTED_NO_ERR, err)
	}
}

func TestCreateVoucher_ValidArgsWithDl(t *testing.T) {
	sl := createdRandomSl(true)

	dl := createdRandomDl()

	creditVal := utils.AbsInt32(utils.RandomInt32())
	data := models.VoucherInsertData{
		Voucher: &models.Voucher{
			Code: utils.RandomString(consts.STRINGS_LEN),
		},
		Items: []*models.VoucherItem{
			{
				DL_ID:  dl.ID,
				SL_ID:  sl.ID,
				DEBIT:  0,
				CREDIT: creditVal,
			},
			{
				DL_ID:  dl.ID,
				SL_ID:  sl.ID,
				DEBIT:  creditVal,
				CREDIT: 0,
			}},
	}
	if _, err := CreateVoucher(&data, testDB); err != nil {
		t.Fatalf(consts.EXPECTED_NO_ERR, err)
	}
}

func TestCreateVoucher_NilSlId(t *testing.T) {
	sl := createdRandomSl(false)

	creditVal := utils.AbsInt32(utils.RandomInt32())
	data := models.VoucherInsertData{
		Voucher: &models.Voucher{
			Code: utils.RandomString(consts.STRINGS_LEN),
		},
		Items: []*models.VoucherItem{
			{
				DL_ID:  0,
				SL_ID:  sl.ID,
				DEBIT:  0,
				CREDIT: creditVal,
			},
			{
				DL_ID:  0,
				SL_ID:  0,
				DEBIT:  creditVal,
				CREDIT: 0,
			}},
	}
	_, err := CreateVoucher(&data, testDB)
	if err == nil || err.Error() != consts.EMPTY_SL_ID {
		t.Fatalf(consts.EXPECTED_EMPTY_SL_ID_ERR, err)
	}
}

func TestCreateVoucher_NotExistSl(t *testing.T) {
	sl := createdRandomSl(false)
	creditVal := utils.AbsInt32(utils.RandomInt32())
	data := models.VoucherInsertData{
		Voucher: &models.Voucher{
			Code: utils.RandomString(consts.STRINGS_LEN),
		},
		Items: []*models.VoucherItem{
			{
				DL_ID:  0,
				SL_ID:  utils.RandomInt64(),
				DEBIT:  0,
				CREDIT: creditVal,
			},
			{
				DL_ID:  0,
				SL_ID:  sl.ID,
				DEBIT:  creditVal,
				CREDIT: 0,
			}},
	}
	_, err := CreateVoucher(&data, testDB)
	if err == nil || err.Error() != consts.MODEL_NOT_FOUND {
		t.Fatalf(consts.EXPECTED_NOT_EXIST_SL_ID_ERR, err)
	}
}

func TestCreateVoucher_SlHasDlButDlIdIsNil(t *testing.T) {
	sl := createdRandomSl(true)

	dl := createdRandomDl()

	creditVal := utils.AbsInt32(utils.RandomInt32())
	data := models.VoucherInsertData{
		Voucher: &models.Voucher{
			Code: utils.RandomString(consts.STRINGS_LEN),
		},
		Items: []*models.VoucherItem{
			{
				DL_ID:  dl.ID,
				SL_ID:  sl.ID,
				DEBIT:  0,
				CREDIT: creditVal,
			},
			{
				DL_ID:  0,
				SL_ID:  sl.ID,
				DEBIT:  creditVal,
				CREDIT: 0,
			}},
	}
	_, err := CreateVoucher(&data, testDB)
	if err == nil || err.Error() != consts.SL_HAS_DL_BUT_DL_ID_IS_NIL {
		t.Fatalf(consts.EXPECTED_EMPTY_DL_ID_ERR, err)
	}
}

func TestCreateVoucher_SlWithDlAndNotExistDl(t *testing.T) {
	sl := createdRandomSl(true)

	dl := createdRandomDl()

	creditVal := utils.AbsInt32(utils.RandomInt32())
	data := models.VoucherInsertData{
		Voucher: &models.Voucher{
			Code: utils.RandomString(consts.STRINGS_LEN),
		},
		Items: []*models.VoucherItem{
			{
				DL_ID:  utils.RandomInt64(),
				SL_ID:  sl.ID,
				DEBIT:  0,
				CREDIT: creditVal,
			},
			{
				DL_ID:  dl.ID,
				SL_ID:  sl.ID,
				DEBIT:  creditVal,
				CREDIT: 0,
			}},
	}
	_, err := CreateVoucher(&data, testDB)
	if err == nil || err.Error() != consts.MODEL_NOT_FOUND {
		t.Fatalf(consts.EXPECTED_NOT_EXIST_DL_ID_ERR, err)
	}
}

func TestCreateVoucher_SlWithoutDlAndNotNilDlId(t *testing.T) {
	sl := createdRandomSl(false)

	creditVal := utils.AbsInt32(utils.RandomInt32())
	data := models.VoucherInsertData{
		Voucher: &models.Voucher{
			Code: utils.RandomString(consts.STRINGS_LEN),
		},
		Items: []*models.VoucherItem{
			{
				DL_ID:  0,
				SL_ID:  sl.ID,
				DEBIT:  0,
				CREDIT: creditVal,
			},
			{
				DL_ID:  utils.RandomInt64(),
				SL_ID:  sl.ID,
				DEBIT:  creditVal,
				CREDIT: 0,
			}},
	}
	_, err := CreateVoucher(&data, testDB)
	if err == nil || err.Error() != consts.SL_HASNT_DL_BUT_DL_ID_ISNT_NIL {
		t.Fatalf(consts.EXPECTED_NOT_EMPTY_DL_ID_ERR, err)
	}
}

func TestCreateVoucher_NegativeCredit(t *testing.T) {
	sl := createdRandomSl(false)

	creditVal := utils.AbsInt32(utils.RandomInt32())
	data := models.VoucherInsertData{
		Voucher: &models.Voucher{
			Code: utils.RandomString(consts.STRINGS_LEN),
		},
		Items: []*models.VoucherItem{
			{
				DL_ID:  0,
				SL_ID:  sl.ID,
				DEBIT:  0,
				CREDIT: -creditVal,
			},
			{
				DL_ID:  0,
				SL_ID:  sl.ID,
				DEBIT:  creditVal,
				CREDIT: 0,
			}},
	}
	_, err := CreateVoucher(&data, testDB)
	if err == nil || err.Error() != consts.INVALID_CREDIT_DEBIT {
		t.Fatalf(consts.EXPECTED_NEG_CREDIT_ERR, err)
	}
}

func TestCreateVoucher_NegativeDebit(t *testing.T) {
	sl := createdRandomSl(false)

	creditVal := utils.AbsInt32(utils.RandomInt32())
	data := models.VoucherInsertData{
		Voucher: &models.Voucher{
			Code: utils.RandomString(consts.STRINGS_LEN),
		},
		Items: []*models.VoucherItem{
			{
				DL_ID:  0,
				SL_ID:  sl.ID,
				DEBIT:  0,
				CREDIT: creditVal,
			},
			{
				DL_ID:  0,
				SL_ID:  sl.ID,
				DEBIT:  -creditVal,
				CREDIT: 0,
			}},
	}
	_, err := CreateVoucher(&data, testDB)
	if err == nil || err.Error() != consts.INVALID_CREDIT_DEBIT {
		t.Fatalf(consts.EXPECTED_NEG_DEBIT_ERR, err)
	}
}

func TestCreateVoucher_NegativeCreditAndDebit(t *testing.T) {
	sl := createdRandomSl(false)

	creditVal := utils.AbsInt32(utils.RandomInt32())
	data := models.VoucherInsertData{
		Voucher: &models.Voucher{
			Code: utils.RandomString(consts.STRINGS_LEN),
		},
		Items: []*models.VoucherItem{
			{
				DL_ID:  0,
				SL_ID:  sl.ID,
				DEBIT:  -creditVal,
				CREDIT: -creditVal,
			},
			{
				DL_ID:  0,
				SL_ID:  sl.ID,
				DEBIT:  creditVal,
				CREDIT: 0,
			}},
	}
	_, err := CreateVoucher(&data, testDB)
	if err == nil || err.Error() != consts.INVALID_CREDIT_DEBIT {
		t.Fatalf(consts.EXPECTED_NEG_DEB_AND_CRED_ERR, err)
	}
}

func TestCreateVoucher_ZeroCreditAndDebit(t *testing.T) {
	sl := createdRandomSl(false)

	creditVal := utils.AbsInt32(utils.RandomInt32())
	data := models.VoucherInsertData{
		Voucher: &models.Voucher{
			Code: utils.RandomString(consts.STRINGS_LEN),
		},
		Items: []*models.VoucherItem{
			{
				DL_ID:  0,
				SL_ID:  sl.ID,
				DEBIT:  creditVal,
				CREDIT: 0,
			},
			{
				DL_ID:  0,
				SL_ID:  sl.ID,
				DEBIT:  0,
				CREDIT: 0,
			}},
	}
	_, err := CreateVoucher(&data, testDB)
	if err == nil || err.Error() != consts.INVALID_CREDIT_DEBIT {
		t.Fatalf(consts.EXPECTED_ZERO_DEB_AND_CRED_ERR, err)
	}
}

func TestCreateVoucher_PositiveCreditAndDebit(t *testing.T) {
	sl := createdRandomSl(false)

	creditVal := utils.AbsInt32(utils.RandomInt32())
	data := models.VoucherInsertData{
		Voucher: &models.Voucher{
			Code: utils.RandomString(consts.STRINGS_LEN),
		},
		Items: []*models.VoucherItem{
			{
				DL_ID:  0,
				SL_ID:  sl.ID,
				DEBIT:  creditVal,
				CREDIT: creditVal,
			},
			{
				DL_ID:  0,
				SL_ID:  sl.ID,
				DEBIT:  creditVal,
				CREDIT: 0,
			}},
	}
	_, err := CreateVoucher(&data, testDB)
	if err == nil || err.Error() != consts.INVALID_CREDIT_DEBIT {
		t.Fatalf(consts.EXPECTED_POS_DEB_AND_CRED_ERR, err)
	}
}

func TestCreateVoucher_EmptyItems(t *testing.T) {
	data := models.VoucherInsertData{
		Voucher: &models.Voucher{
			Code: utils.RandomString(consts.STRINGS_LEN),
		},
		Items: []*models.VoucherItem{},
	}
	_, err := CreateVoucher(&data, testDB)
	if err == nil || err.Error() != consts.INVALID_NUM_OF_ITEMS {
		t.Fatalf(consts.EXPECTED_OUT_OF_LIMIT_ITEMS_ERR, err)
	}
}

func TestCreateVoucher_OneItem(t *testing.T) {
	sl := createdRandomSl(false)

	creditVal := utils.AbsInt32(utils.RandomInt32())
	data := models.VoucherInsertData{
		Voucher: &models.Voucher{
			Code: utils.RandomString(consts.STRINGS_LEN),
		},
		Items: []*models.VoucherItem{
			{
				DL_ID:  0,
				SL_ID:  sl.ID,
				DEBIT:  creditVal,
				CREDIT: 0,
			},
		},
	}
	_, err := CreateVoucher(&data, testDB)
	if err == nil || err.Error() != consts.INVALID_NUM_OF_ITEMS {
		t.Fatalf(consts.EXPECTED_OUT_OF_LIMIT_ITEMS_ERR, err)
	}
}

func TestCreateVoucher_AboveRangeNumOfItems(t *testing.T) {
	sl := createdRandomSl(false)

	numOfItems := rand.Intn(1000) + 500
	creditVal := utils.AbsInt32(utils.RandomInt32())
	data := models.VoucherInsertData{
		Voucher: &models.Voucher{
			Code: utils.RandomString(consts.STRINGS_LEN),
		},
		Items: make([]*models.VoucherItem, numOfItems),
	}

	for i := 0; i < numOfItems; i++ {
		item := &models.VoucherItem{
			DL_ID:  0,
			SL_ID:  sl.ID,
			DEBIT:  0,
			CREDIT: creditVal,
		}
		if i%2 == 1 {
			item.DEBIT = creditVal
			item.CREDIT = 0
		}
		if i == 1 {
			item.DEBIT *= 2
		}
		data.Items[i] = item
	}

	_, err := CreateVoucher(&data, testDB)
	if err == nil || err.Error() != consts.INVALID_NUM_OF_ITEMS {
		t.Fatalf(consts.EXPECTED_OUT_OF_LIMIT_ITEMS_ERR, err)
	}
}

func TestCreateVoucher_NotBalanced(t *testing.T) {
	sl := createdRandomSl(false)

	numOfItems := rand.Intn(499) + 2

	data := models.VoucherInsertData{
		Voucher: &models.Voucher{
			Code: utils.RandomString(consts.STRINGS_LEN),
		},
		Items: make([]*models.VoucherItem, numOfItems),
	}

	for i := 0; i < numOfItems; i++ {
		rand.Seed(uint64(time.Now().UnixNano()))
		creditVal := utils.AbsInt32(utils.RandomInt32())
		positCredit := utils.RandomBool()
		item := &models.VoucherItem{
			DL_ID:  0,
			SL_ID:  sl.ID,
			DEBIT:  creditVal * int32(utils.BoolToInt(!positCredit)),
			CREDIT: creditVal * int32(utils.BoolToInt(positCredit)),
		}
		data.Items[i] = item
	}

	_, err := CreateVoucher(&data, testDB)
	if err == nil || err.Error() != consts.NOT_BALANCED {
		t.Fatalf(consts.EXPECTED_NOT_BALANCED_ITEMS_ERR, err)
	}
}

func TestUpdateVoucher_ValidArgs(t *testing.T) {
	sl1 := createdRandomSl(false)

	id1 := utils.RandomInt64()
	id2 := utils.RandomInt64()
	creditVal := utils.AbsInt32(utils.RandomInt32())
	createdV := createRandomVoucher(sl1.ID, id1, id2)

	sl2 := createdRandomSl(false)

	updatedV := &models.VoucherUpdateData{
		Voucher: &models.Voucher{
			ID:   createdV.ID,
			Code: utils.RandomString(consts.STRINGS_LEN),
		},
		Items: &models.VoucherItemUpdateData{
			ToInserts: []*models.VoucherItem{
				{
					DL_ID:  0,
					SL_ID:  sl2.ID,
					DEBIT:  0,
					CREDIT: creditVal,
				},
			},
			ToUpdates: []*models.VoucherItem{
				{
					ID:     id1,
					DL_ID:  0,
					SL_ID:  sl2.ID,
					DEBIT:  creditVal,
					CREDIT: 0,
				},
			},
			ToDeletes: []*int64{
				&id2,
			},
		},
	}

	if err := updateVoucher(updatedV, testDB); err != nil {
		t.Fatalf(consts.EXPECTED_NO_ERR, err)
	}
}

func TestUpdateVoucher_VidNotExist(t *testing.T) {
	sl1 := createdRandomSl(false)

	id1 := utils.RandomInt64()
	id2 := utils.RandomInt64()
	creditVal := utils.AbsInt32(utils.RandomInt32())
	createRandomVoucher(sl1.ID, id1, id2)

	sl2 := &models.SLModel{
		Code:  utils.RandomString(consts.STRINGS_LEN),
		Title: utils.RandomString(consts.STRINGS_LEN),
		HasDl: false,
	}
	sl_id2, err := CreateSLModel(sl2, testDB)
	if err != nil {
		t.Fatalf(consts.EXPECTED_NO_ERR, err)
	}

	updatedV := &models.VoucherUpdateData{
		Voucher: &models.Voucher{
			ID:   utils.RandomInt64(),
			Code: utils.RandomString(consts.STRINGS_LEN),
		},
		Items: &models.VoucherItemUpdateData{
			ToInserts: []*models.VoucherItem{
				{
					DL_ID:  0,
					SL_ID:  *sl_id2,
					DEBIT:  0,
					CREDIT: creditVal,
				},
			},
			ToUpdates: []*models.VoucherItem{
				{
					ID:     id1,
					DL_ID:  0,
					SL_ID:  *sl_id2,
					DEBIT:  creditVal,
					CREDIT: 0,
				},
			},
			ToDeletes: []*int64{
				&id2,
			},
		},
	}
	err = updateVoucher(updatedV, testDB)
	if err == nil || err.Error() != consts.MODEL_NOT_FOUND {
		t.Fatalf(consts.EXPECTED_NOT_FOUND_ERR, err)
	}
}

func TestUpdateVoucher_UpdateItemIdNotExist(t *testing.T) {
	sl1 := createdRandomSl(false)

	id2 := utils.RandomInt64()
	creditVal := utils.AbsInt32(utils.RandomInt32())
	createdV := createRandomVoucher(sl1.ID, utils.RandomInt64(), id2)

	sl2 := &models.SLModel{
		Code:  utils.RandomString(consts.STRINGS_LEN),
		Title: utils.RandomString(consts.STRINGS_LEN),
		HasDl: false,
	}
	sl_id2, err := CreateSLModel(sl2, testDB)
	if err != nil {
		t.Fatalf(consts.EXPECTED_NO_ERR, err)
	}

	updatedV := &models.VoucherUpdateData{
		Voucher: &models.Voucher{
			ID:   createdV.ID,
			Code: utils.RandomString(consts.STRINGS_LEN),
		},
		Items: &models.VoucherItemUpdateData{
			ToInserts: []*models.VoucherItem{
				{
					DL_ID:  0,
					SL_ID:  *sl_id2,
					DEBIT:  0,
					CREDIT: creditVal,
				},
			},
			ToUpdates: []*models.VoucherItem{
				{
					ID:     utils.RandomInt64(),
					DL_ID:  0,
					SL_ID:  *sl_id2,
					DEBIT:  creditVal,
					CREDIT: 0,
				},
			},
			ToDeletes: []*int64{
				&id2,
			},
		},
	}
	err = updateVoucher(updatedV, testDB)
	if err == nil || err.Error() != consts.NOT_FOUND_ITEM_TO_UPDATE {
		t.Fatalf(consts.EXPECTED_NOT_FOUND_ERR, err)
	}
}

func TestUpdateVoucher_DeleteItemIdNotExist(t *testing.T) {
	sl1 := createdRandomSl(false)

	creditVal := utils.AbsInt32(utils.RandomInt32())
	createdV := createRandomVoucher(sl1.ID, utils.RandomInt64(), utils.RandomInt64())

	sl2 := &models.SLModel{
		Code:  utils.RandomString(consts.STRINGS_LEN),
		Title: utils.RandomString(consts.STRINGS_LEN),
		HasDl: false,
	}
	sl_id2, err := CreateSLModel(sl2, testDB)
	if err != nil {
		t.Fatalf(consts.EXPECTED_NO_ERR, err)
	}

	deleteId := utils.RandomInt64()
	updatedV := &models.VoucherUpdateData{
		Voucher: &models.Voucher{
			ID:   createdV.ID,
			Code: utils.RandomString(consts.STRINGS_LEN),
		},
		Items: &models.VoucherItemUpdateData{
			ToInserts: []*models.VoucherItem{
				{
					DL_ID:  0,
					SL_ID:  *sl_id2,
					DEBIT:  0,
					CREDIT: creditVal,
				},
			},
			ToUpdates: []*models.VoucherItem{
				{
					ID:     utils.RandomInt64(),
					DL_ID:  0,
					SL_ID:  *sl_id2,
					DEBIT:  creditVal,
					CREDIT: 0,
				},
			},
			ToDeletes: []*int64{
				&deleteId,
			},
		},
	}
	err = updateVoucher(updatedV, testDB)
	if err == nil || err.Error() != consts.NOT_FOUND_ITEM_TO_DELETE {
		t.Fatalf(consts.EXPECTED_NOT_FOUND_ERR, err)
	}
}

func TestUpdateVoucher_NotBalanced(t *testing.T) {
	sl1 := createdRandomSl(false)

	id1 := utils.RandomInt64()
	id2 := utils.RandomInt64()
	creditVal := utils.AbsInt32(utils.RandomInt32())
	createdV := createRandomVoucher(sl1.ID, id1, id2)

	sl2 := &models.SLModel{
		Code:  utils.RandomString(consts.STRINGS_LEN),
		Title: utils.RandomString(consts.STRINGS_LEN),
		HasDl: false,
	}
	sl_id2, err := CreateSLModel(sl2, testDB)
	if err != nil {
		t.Fatalf(consts.EXPECTED_NO_ERR, err)
	}

	updatedV := &models.VoucherUpdateData{
		Voucher: &models.Voucher{
			ID:   createdV.ID,
			Code: utils.RandomString(consts.STRINGS_LEN),
		},
		Items: &models.VoucherItemUpdateData{
			ToInserts: []*models.VoucherItem{
				{
					DL_ID:  0,
					SL_ID:  *sl_id2,
					DEBIT:  0,
					CREDIT: creditVal,
				},
				{
					DL_ID:  0,
					SL_ID:  *sl_id2,
					DEBIT:  creditVal,
					CREDIT: 0,
				},
			},
			ToUpdates: []*models.VoucherItem{
				{
					ID:     id1,
					DL_ID:  0,
					SL_ID:  *sl_id2,
					DEBIT:  creditVal,
					CREDIT: 0,
				},
			},
			ToDeletes: []*int64{
				&id2,
			},
		},
	}

	err = updateVoucher(updatedV, testDB)
	if err == nil || err.Error() != consts.NOT_BALANCED {
		t.Fatalf(consts.EXPECTED_NOT_BALANCED_ITEMS_ERR, err)
	}
}

func TestDeleteVoucher_Valid(t *testing.T) {
	sl1 := createdRandomSl(false)

	id1 := utils.RandomInt64()
	id2 := utils.RandomInt64()
	createdV := createRandomVoucher(sl1.ID, id1, id2)

	if err := DeleteVoucher(createdV.ID, testDB); err != nil {
		t.Fatalf(consts.EXPECTED_NO_ERR, err)
	}
}

func TestDeleteVoucher_NotExistVid(t *testing.T) {
	randomVid := utils.RandomInt64()
	err := DeleteVoucher(randomVid, testDB)
	if err == nil || err.Error() != consts.MODEL_NOT_FOUND {
		t.Fatalf(consts.EXPECTED_NOT_FOUND_ERR, err)
	}
}

func TestGetVoucher_ValidArgs(t *testing.T) {
	sl1 := createdRandomSl(false)

	id1 := utils.RandomInt64()
	id2 := utils.RandomInt64()
	createdV := createRandomVoucher(sl1.ID, id1, id2)

	dataOut, err := GetVoucher(createdV.ID, testDB)
	if err != nil {
		t.Fatalf(consts.EXPECTED_NO_ERR, err)
	}
	if !dataOut.Equals(dataOut) {
		t.Fatalf(consts.EXPECTED_DIFF_RETURNED_MODEL_ERR, createdV, *dataOut)
	}
}

func TestGetVoucher_NotExistVid(t *testing.T) {
	randomVid := utils.RandomInt64()
	_, err := GetVoucher(randomVid, testDB)
	if err == nil || err.Error() != consts.MODEL_NOT_FOUND {
		t.Fatalf(consts.EXPECTED_NOT_FOUND_ERR, err)
	}
}
