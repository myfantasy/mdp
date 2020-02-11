package mdp

import (
	"encoding/json"
	"strconv"
)

// URLStructGet url for get struct
const URLStructGet = "struct/get/"

// URLStructSet url for set struct
const URLStructSet = "struct/set/"

// KeyTypeString string
const KeyTypeString = "string"

// KeyTypeInt int
const KeyTypeInt = "int"

// TableDefinition table definition
type TableDefinition struct {
	UniqueID  string `json:"id"`
	TableName string `json:"name"`
	TableType string `json:"type"`
	KeyType   string `json:"key_type"`

	StoragePlace string `json:"storage,omitempty"`

	ShardKeys []int64 `json:"shard_keys,omitempty"`

	PublicParams json.RawMessage `json:"public_params,omitempty"`
	LocalParams  json.RawMessage `json:"local_params,omitempty"`

	MetaData json.RawMessage `json:"meta,omitempty"`

	ItemsStruct *ItemStruct `json:"items_struct,omitempty"`

	Version   int64 `json:"version,omitempty"`
	IsDeleted bool  `json:"is_deleted,omitempty"`
}

// ClearLocalInfo - remove local info
func (td TableDefinition) ClearLocalInfo() TableDefinition {
	td2 := TableDefinition{
		UniqueID:     td.UniqueID,
		TableName:    td.TableName,
		TableType:    td.TableType,
		KeyType:      td.KeyType,
		ShardKeys:    td.ShardKeys,
		PublicParams: td.PublicParams,
		MetaData:     td.MetaData,
		ItemsStruct:  td.ItemsStruct,
		Version:      td.Version,
		IsDeleted:    td.IsDeleted,
	}

	return td2
}

// StructGetQuery query for get struct
type StructGetQuery struct {
	// priority 0
	TableName string `json:"table,omitempty"`

	// priority 1
	LoadAll bool `json:"load_all,omitempty"`

	// flag
	LoadInternalInfo bool `json:"load_short,omitempty"`
}

// StructSetQuery query for get struct
type StructSetQuery struct {

	// priority 0
	DropTableName string `json:"drop_table,omitempty"`

	// priority 1
	CreateTable *TableDefinition `json:"create_table,omitempty"`

	// priority 2
	AlterTable           *TableDefinition `json:"alter_table,omitempty"`
	AlterShardKeysGlobal bool             `json:"alter_shard_keys_global,omitempty"`

	GetServerInfo bool `json:"get_server_info,omitempty"`
}

// StructGet struct get result
type StructGet struct {
	ParamsErr   *Error `json:"params_err,omitempty"`
	InternalErr *Error `json:"internal_err,omitempty"`

	Tables []TableDefinition `json:"tables,omitempty"`

	LastTableVersion int64 `json:"last_table_version,omitempty"`
}

// QueryStructGet send any query waiting StructGet
func (c *Connection) QueryStructGet(path string, v interface{}) (res StructGet) {
	body, statusCode, err := c.DoQueryObject(path, v)
	if err != nil {
		res.InternalErr = ErrorE(err)
		return res
	}

	err = json.Unmarshal(body, &res)
	if err != nil {
		res.InternalErr = ErrorNew("status_code = "+strconv.Itoa(statusCode), err)
		return res
	}

	if statusCode != StatusCodeOK {
		if res.InternalErr == nil && res.ParamsErr == nil {
			res.InternalErr = ErrorS("status_code = " + strconv.Itoa(statusCode))
		}
	}

	return res
}

// StructRawGet get struct by query
func (c *Connection) StructRawGet(sgq StructGetQuery) StructGet {
	return c.QueryStructGet(URLStructGet, sgq)
}

// StructRawSet set struct by query
func (c *Connection) StructRawSet(ssq StructSetQuery) StructGet {
	return c.QueryStructGet(URLStructSet, ssq)
}
