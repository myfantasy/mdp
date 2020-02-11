package mdp

import (
	"encoding/json"
	"strconv"
)

// URLStructStorageGet url for get struct
const URLStructStorageGet = "struct/storage_get/"

// URLStructStorageSet url for set struct
const URLStructStorageSet = "struct/storage_set/"

// StorageDefinition storage definition
type StorageDefinition struct {
	StorageName      string `json:"name"`
	LastTableVersion int64  `json:"last_table_version,omitempty"`
	LocalStoragePath string `json:"local_path"`
}

// ClearLocalInfo - remove local info
func (sd StorageDefinition) ClearLocalInfo() StorageDefinition {
	sd2 := StorageDefinition{
		StorageName:      sd.StorageName,
		LastTableVersion: sd.LastTableVersion,
	}

	return sd2
}

// StructStorageGetQuery query for get struct
type StructStorageGetQuery struct {
	Name bool `json:"get_name,omitempty"`

	LocalPlace bool `json:"get_local_place,omitempty"`
}

// StructStorageSetQuery query for get struct
type StructStorageSetQuery struct {
	Name string `json:"name,omitempty"`

	// Stop
	Stop bool `json:"stop,omitempty"`
}

// StructStorageGet struct get result
type StructStorageGet struct {
	ParamsErr   *Error `json:"params_err,omitempty"`
	InternalErr *Error `json:"internal_err,omitempty"`

	Storage StorageDefinition `json:"storage,omitempty"`
}

// QueryStructStorageGet send any query waiting StructStorageGet
func (c *Connection) QueryStructStorageGet(path string, v interface{}) (res StructStorageGet) {
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

// StructStorageRawGet get struct by query
func (c *Connection) StructStorageRawGet(sgq StructStorageGetQuery) StructStorageGet {
	return c.QueryStructStorageGet(URLStructStorageGet, sgq)
}

// StructStorageRawSet set struct by query
func (c *Connection) StructStorageRawSet(ssq StructStorageSetQuery) StructStorageGet {
	return c.QueryStructStorageGet(URLStructStorageSet, ssq)
}
