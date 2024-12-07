package accountSys

import (
	"Final/internal/accountSys/consts"
	"Final/internal/accountSys/models"
	"Final/internal/accountSys/utils"
	"Final/internal/accountSys/validations/modelsValidate"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func CreateDLModel(dl *models.DLModel, db *gorm.DB) (id *int64, err error) {
	if err := utils.ValidateModel(&dl.Code, &dl.Title); err != nil {
		return nil, err
	}

	if err := utils.CreateInDb(dl, db); err != nil {
		return nil, fmt.Errorf(consts.CREATING_SL_FAILED, err)
	}

	return &dl.ID, nil
}

func updateDLModel(dl *models.DLModel, db *gorm.DB) error {
	if err := utils.ValidateModel(&dl.Code, &dl.Title); err != nil {
		return err
	}

	if !utils.IsItInDb(dl.ID, dl, db) {
		return errors.New(consts.MODEL_NOT_FOUND)
	}

	return utils.UpdateInDb(dl, db)
}

func deleteDLModel(id int64, db *gorm.DB) error {
	if !utils.IsItInDb(id, &models.DLModel{}, db) {
		return errors.New(consts.MODEL_NOT_FOUND)
	}

	existsAnyRef := existAnyRefOnModel(id, consts.DL_MODEL, db)

	if existsAnyRef {
		return errors.New(consts.THIS_DL_IS_REFERENCED)
	}

	return utils.DeleteFromDb(id, &models.DLModel{}, db)
}

func getDLModel(id int64, db *gorm.DB) (*models.DLModel, error) {
	if !utils.IsItInDb(id, &models.DLModel{}, db) {
		return nil, errors.New(consts.MODEL_NOT_FOUND)
	}

	var dl models.DLModel
	utils.GetModelFromDb(id, &dl, db)

	return &dl, nil
}

func CreateSLModel(sl *models.SLModel, db *gorm.DB) (*int64, error) {
	if err := utils.ValidateModel(&sl.Code, &sl.Title); err != nil {
		return nil, err
	}

	if err := utils.CreateInDb(sl, db); err != nil {
		return nil, fmt.Errorf(consts.CREATING_DL_FAILED, err)
	}

	return &sl.ID, nil
}

func updateSLModeldl(sl *models.SLModel, db *gorm.DB) error {
	if err := utils.ValidateModel(&sl.Code, &sl.Title); err != nil {
		return err
	}

	if !utils.IsItInDb(sl.ID, sl, db) {
		return errors.New(consts.MODEL_NOT_FOUND)
	}

	existsAnyRef := existAnyRefOnModel(sl.ID, consts.SL_MODEL, db)

	if existsAnyRef {
		return errors.New(consts.THIS_SL_IS_REFERENCED)
	}

	return utils.UpdateInDb(sl, db)
}

func deleteSLModel(id int64, db *gorm.DB) error {

	if !utils.IsItInDb(id, &models.SLModel{}, db) {
		return errors.New(consts.MODEL_NOT_FOUND)
	}

	existsAnyRef := existAnyRefOnModel(id, consts.SL_MODEL, db)
	if existsAnyRef {
		return errors.New(consts.THIS_SL_IS_REFERENCED)
	}

	return utils.DeleteFromDb(id, &models.SLModel{}, db)
}

func getSLModel(id int64, db *gorm.DB) (*models.SLModel, error) {
	if !utils.IsItInDb(id, &models.SLModel{}, db) {
		return nil, errors.New(consts.MODEL_NOT_FOUND)
	}

	var sl models.SLModel
	utils.GetModelFromDb(id, &sl, db)

	return &sl, nil
}

func CreateVoucher(v *models.VoucherInsertData, db *gorm.DB) (*int64, error) {
	randomTitle := utils.RandomString(consts.STRINGS_LEN)
	if err := utils.ValidateModel(&v.Voucher.Code, &randomTitle); err != nil {
		return nil, err
	}

	if err := modelsValidate.ValidateItems(&v.Items, db); err != nil {
		return nil, err
	}

	if err := utils.CreateInDb(v.Voucher, db); err != nil {
		return nil, fmt.Errorf(consts.CREATING_VOUCHER_FAILED, err)
	}

	if err := insertItemsToDb(&v.Items, v.Voucher.ID, db); err != nil {
		return nil, err
	}

	return &v.Voucher.ID, nil
}

func deleteItems(items *map[int64]*models.VoucherItem, toDeleteIds *[]*int64) error {
	for _, deleteId := range *toDeleteIds {
		if _, okay := (*items)[*deleteId]; !okay {
			return errors.New(consts.V_ITEM_NOT_FOUND)
		}
		delete(*items, *deleteId)
	}
	return nil
}

func updateItems(items *map[int64]*models.VoucherItem, toUpdates *[]*models.VoucherItem) error {
	for _, updateItem := range *toUpdates {
		if _, okay := (*items)[(*updateItem).ID]; !okay {
			return errors.New(consts.V_ITEM_NOT_FOUND)
		}
		(*items)[(*updateItem).ID] = updateItem
	}
	return nil
}

func insertItems(items *map[int64]*models.VoucherItem, toInsert *[]*models.VoucherItem) error {
	for _, insertItem := range *toInsert {
		if (*insertItem).ID == 0 {
			for {
				insertItem.ID = utils.RandomInt64()
				if _, exists := (*items)[insertItem.ID]; !exists {
					break
				}
			}

		} else if _, okay := (*items)[(*insertItem).ID]; okay {
			return errors.New(consts.V_ITEM_ALREADY_EXIST)
		}
		(*items)[(*insertItem).ID] = insertItem
	}
	return nil
}

func insertItemsToDb(toInserts *[]*models.VoucherItem, vId int64, db *gorm.DB) error {
	for _, item := range *toInserts {
		item.Vouch_ID = vId
		if err := utils.CreateInDb(&item, db); err != nil {
			return err
		}
	}

	return nil
}

func updateItemsInDb(toUpdates *[]*models.VoucherItem, db *gorm.DB) error {
	for _, item := range *toUpdates {
		if err := db.Save(item).Error; err != nil {
			return errors.New(consts.UPDATE_V_ITEM_IN_DB_FAILED + err.Error())
		}
	}

	return nil
}

func deleteItemsFromDb(toUpdates *[]*int64, db *gorm.DB) error {
	for _, id := range *toUpdates {
		if err := db.Delete(&models.VoucherItem{}, "id = ?", id).Error; err != nil {
			return errors.New(consts.DELETE_V_ITEM_FROM_DB_FAILED + err.Error())
		}
	}

	return nil
}

func applyItemUpdatesOnDb(items *models.VoucherItemUpdateData, vId int64, db *gorm.DB) error {
	if err := insertItemsToDb(&items.ToInserts, vId, db); err != nil {
		return err
	}

	if err := updateItemsInDb(&items.ToUpdates, db); err != nil {
		return err
	}

	if err := deleteItemsFromDb(&items.ToDeletes, db); err != nil {
		return err
	}

	return nil
}

func getVItemsFromDb(vId int64, db *gorm.DB) (items []*models.VoucherItem, err error) {
	err = db.Where(consts.FIND_BI_V_ID_CMD, vId).Find(&items).Error
	if err != nil {
		return nil, errors.New(consts.GET_IN_DB_FAILED)
	}

	return items, nil
}

func handleToUpdateItems(vItemMap map[int64]*models.VoucherItem,
	v *models.VoucherUpdateData, db *gorm.DB) error {

	err := deleteItems(&vItemMap, &v.Items.ToDeletes)
	if err != nil {
		return errors.New(consts.ERR_DELETING_ITEM + err.Error())
	}

	err = updateItems(&vItemMap, &v.Items.ToUpdates)
	if err != nil {
		return errors.New(consts.ERR_UPDATING_ITEM + err.Error())
	}

	err = insertItems(&vItemMap, &v.Items.ToInserts)
	if err != nil {
		return errors.New(consts.ERR_INSERTING_ITEM + err.Error())
	}

	vItems := utils.MapToArray(vItemMap)

	if err = modelsValidate.ValidateItems(&vItems, db); err != nil {
		return err
	}

	if err = applyItemUpdatesOnDb(v.Items, v.ID, db); err != nil {
		return err
	}

	return nil
}

func updateVoucher(v *models.VoucherUpdateData, db *gorm.DB) error {
	randomTitle := utils.RandomString(consts.STRINGS_LEN)
	if err := utils.ValidateModel(&v.Voucher.Code, &randomTitle); err != nil {
		return err
	}

	if !utils.IsItInDb(v.ID, v.Voucher, db) {
		return errors.New(consts.MODEL_NOT_FOUND)
	}

	vItems, err := getVItemsFromDb(v.ID, db)
	if err != nil {
		return err
	}

	vItemMap := utils.ArrayToMap(vItems, func(v *models.VoucherItem) int64 {
		return v.ID
	})

	if err = handleToUpdateItems(vItemMap, v, db); err != nil {
		return err
	}

	return utils.UpdateInDb(v.Voucher, db)
}

func DeleteVoucher(vId int64, db *gorm.DB) error {
	if !utils.IsItInDb(vId, &models.Voucher{}, db) {
		return errors.New(consts.MODEL_NOT_FOUND)
	}

	vItems, err := getVItemsFromDb(vId, db)
	if err != nil {
		return err
	}

	vItemIds := make([]*int64, len(vItems))
	for index, item := range vItems {
		vItemIds[index] = &item.ID
	}
	deleteItemsFromDb(&vItemIds, db)

	return utils.DeleteFromDb(vId, &models.Voucher{}, db)
}

func GetVoucher(vId int64, db *gorm.DB) (*models.VoucherInsertData, error) {
	if !utils.IsItInDb(vId, &models.Voucher{}, db) {
		return nil, errors.New(consts.MODEL_NOT_FOUND)
	}
	vModel := models.Voucher{}
	utils.GetModelFromDb(vId, &vModel, db)

	vItems, err := getVItemsFromDb(vId, db)
	if err != nil {
		return nil, err
	}

	outData := &models.VoucherInsertData{
		Voucher: &vModel,
		Items:   vItems,
	}
	return outData, nil
}

func existAnyRefOnModel(id int64, modelType bool, db *gorm.DB) bool {
	var searchCmd string
	if modelType == consts.DL_MODEL {
		searchCmd = consts.FIND_BY_DL_ID_CMD
	} else {
		searchCmd = consts.FIND_BY_SL_ID_CMD
	}

	var vItem models.VoucherItem
	err := db.Where(searchCmd, id).First(&vItem).Error

	return err == nil
}
