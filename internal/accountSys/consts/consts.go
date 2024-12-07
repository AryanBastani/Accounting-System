package consts

const SEPERATOR = " "

const UNUNIQUE_ID = "ID already exists. Please use a unique ID"
const UNUNIQUE_TITLE = "Title already exists. Please choose a different title"

const EMPTY_STR = "must not be empty"
const LONG_STR = "must be at least 64 characters long"

const ID = "ID"
const CODE = "Code"
const TITLE = "Title"

const EMPTY_CODE = CODE + SEPERATOR + EMPTY_STR
const LONG_CODE = CODE + SEPERATOR + LONG_STR

const EMPTY_TITLE = TITLE + SEPERATOR + EMPTY_STR
const LONG_TITLE = TITLE + SEPERATOR + LONG_STR

const DL_MODEL_NOT_FOUND = "DLModel not found"
const SL_MODEL_NOT_FOUND = "SLModel not found"
const VOUCHER_NOT_FOUND = "voucher not found"

const EMPTY_SL_ID = "sl_id shouldn't be empty"
const SL_NOT_FOUND = "the sl model with given id not found"

const SL_HAS_DL_BUT_DL_ID_IS_NIL = "sl has dl. so you should choose a dl_id"
const SL_HASNT_DL_BUT_DL_ID_ISNT_NIL = "sl hasn't dl. so you shouldn't choose any dl_id"
const INVALID_CREDIT_DEBIT = "one of the credit and debit values should be positive and the other one be zero"
const INVALID_NUM_OF_ITEMS = "number of items should be at least 2 and at most 500"
const NOT_BALANCED = "the values of credit and debit shouold be balanced in items"

const V_ITEM_NOT_FOUND = "the voucherItem id not found"
const ERR_DELETING_ITEM = "got error while deleting items: "
const ERR_UPDATING_ITEM = "got error while updating items: "
const ERR_INSERTING_ITEM = "got error while inserting items: "
const V_ITEM_ALREADY_EXIST = "the voucherItem id is already exist"

const NOT_FOUND_ITEM_TO_UPDATE = ERR_UPDATING_ITEM + V_ITEM_NOT_FOUND
const NOT_FOUND_ITEM_TO_DELETE = ERR_DELETING_ITEM + V_ITEM_NOT_FOUND

const THIS_DL_IS_REFERENCED = "this dl is referenced by some VoucherItems"
const THIS_SL_IS_REFERENCED = "this sl is referenced by some VoucherItems"

const DL_MODEL = true
const SL_MODEL = false

const DL_UPDATE_NOT_FOUND = "failed to update DLModel: model not found"

const CREATING_SL_FAILED = "failed to create SLModel: %w"
const CREATING_DL_FAILED = "failed to create DLModel: %w"
const CREATING_VOUCHER_FAILED = "failed to create Voucher: %w"
const UPDATE_DL_FAILED = "failed to update DLModel: %s"
const MODEL_NOT_FOUND = "model not found"

const CREATE_IN_DB_FAILED = "failed to create model from database"
const DELETE_FROM_DB_FAILED = "failed to delete model from database"
const UPDATE_IN_DB_FAILED = "failed to update model in database"
const GET_IN_DB_FAILED = "failed to get model from database"

const UPDATE_V_ITEM_IN_DB_FAILED = "failed to update VoucherItem: "
const DELETE_V_ITEM_FROM_DB_FAILED = "failed to delete VoucherItem: "

const DL_ID_COLUMN = "dl_id"
const SL_ID_COLUMN = "sl_id"

const FIND_BY_ID_CMD = "id = ?"
const FIND_BI_V_ID_CMD = "vouch_id = ?"
const FIND_BY_DL_ID_CMD = "dl_id = ?"
const FIND_BY_SL_ID_CMD = "sl_id = ?"

const CONNECT_DB_FAILED = "failed to connect to test database: %v"
const MIGRATE_DB_FAILED = "failed to migrate test database: %v"

const EXPECTED_NO_ERR = "expected no error, got: %v"
const EXPECTED_EMPTY_TITLE_ERR = "expected error for the empty TITLE. but got: %v"
const EXPECTED_EMPTY_CODE_ERR = "expected error for the empty CODE. but got: %v"
const EXPECTED_LONG_TITLE_ERR = "expected error for the long TITLE. but got: %v"
const EXPECTED_LONG_CODE_ERR = "expected error for the long CODE. but got: %v"
const NOT_EXPECTED_INSERT_ERR = "not expected error for inserting. but got: %v"
const EXPECTED_NOT_FOUND_ERR = "expected error for not found model, but got: %v"
const EXPECTED_DIFF_RETURNED_MODEL_ERR = "the returned model is not the expected one. exepted: %v, got: %v"
const EXPECTED_REFERENCED_ERR = "expected error for referenced model. but got: %v"
const EXPECTED_EMPTY_SL_ID_ERR = "expected error for empty sl_id, but got: %v"
const EXPECTED_NOT_EXIST_SL_ID_ERR = "expected error for not existing sl_id, but got: %v"
const EXPECTED_EMPTY_DL_ID_ERR = "expected error for empty dl_id, but got: %v"
const EXPECTED_NOT_EXIST_DL_ID_ERR = "expected error for not existing dl_id, but got: %v"
const EXPECTED_NOT_EMPTY_DL_ID_ERR = "expected error for not nil dl_id, but got: %v"
const EXPECTED_NEG_CREDIT_ERR = "expected error for negative credit, but got: %v"
const EXPECTED_NEG_DEBIT_ERR = "expected error for negative debit, but got: %v"
const EXPECTED_NEG_DEB_AND_CRED_ERR = "expected error for negative debit and credit, but got: %v"
const EXPECTED_ZERO_DEB_AND_CRED_ERR = "expected error for zero debit and credit, but got: %v"
const EXPECTED_POS_DEB_AND_CRED_ERR = "expected error for positive debit and credit, but got: %v"
const EXPECTED_OUT_OF_LIMIT_ITEMS_ERR = "expected error for under limit num of items, but got: %v"
const EXPECTED_NOT_BALANCED_ITEMS_ERR = "expected error for not balanced items, but got: %v"

const ALL_CHARS = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

const STRINGS_LEN = 64
