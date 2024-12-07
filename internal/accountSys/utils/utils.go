package utils

import (
	"Final/internal/accountSys/consts"
	"Final/internal/accountSys/validations/strValidate"
	"errors"
	"time"

	"golang.org/x/exp/rand"
	"gorm.io/gorm"
)

func CreateInDb(model interface{}, db *gorm.DB) error {
	if err := db.Create(model).Error; err != nil {
		return errors.New(consts.CREATE_IN_DB_FAILED)
	}

	return nil
}

func IsItInDb(id int64, model interface{}, db *gorm.DB) bool {
	if err := db.Where(consts.FIND_BY_ID_CMD, id).First(model).Error; err != nil {
		return false
	}

	return true
}

func UpdateInDb(model interface{}, db *gorm.DB) error {
	if err := db.Save(model).Error; err != nil {
		return errors.New(consts.DELETE_FROM_DB_FAILED)
	}

	return nil
}

func DeleteFromDb(id int64, model interface{}, db *gorm.DB) error {
	if err := db.Delete(model, consts.FIND_BY_ID_CMD, id).Error; err != nil {
		return errors.New(consts.GET_IN_DB_FAILED)
	}

	return nil
}

func GetModelFromDb(id int64, model interface{}, db *gorm.DB) error {
	if err := db.First(model, consts.FIND_BY_ID_CMD, id).Error; err != nil {
		return errors.New(consts.GET_IN_DB_FAILED)
	}

	return nil
}

func ArrayToMap[T any, K comparable](arr []T, keySelector func(T) K) map[K]T {
	result := make(map[K]T)
	for _, item := range arr {
		key := keySelector(item)
		result[key] = item
	}
	return result
}

func MapToArray[K comparable, V any](m map[K]V) []V {
	result := make([]V, 0, len(m))
	for _, value := range m {
		result = append(result, value)
	}
	return result
}

func RandomString(length int) string {
	const charset = consts.ALL_CHARS
	rand.Seed(uint64(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}

	return string(b)
}

func BoolToInt(input bool) int {
	if input {
		return 1
	}
	return 0
}

func RandomBool() bool {
	rand.Seed(uint64(time.Now().UnixNano()))

	return rand.Intn(2) == 1
}

func AbsInt32(n int32) int32 {
	if n < 0 {
		return -n
	}
	return n
}

func ValidateModel(code, title *string) error {
	if err := strValidate.ValidateStr(code); err != nil {
		return errors.New(consts.CODE + " " + err.Error())
	}

	if err := strValidate.ValidateStr(title); err != nil {
		return errors.New(consts.TITLE + " " + err.Error())
	}

	return nil
}

func RandomInt64() int64 {
	rand.Seed(uint64(time.Now().UnixNano()))
	return rand.Int63()
}

func RandomInt32() int32 {
	rand.Seed(uint64(time.Now().UnixNano()))
	return rand.Int31()
}
