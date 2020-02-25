package mdp

import (
	"encoding/json"
	"strconv"
)

// URLGet url for get item(s) from key tables (json)
const URLGet = "get/"

// URLSet url for set item(s) into key tables (json)
const URLSet = "set/"

// StatusCodeOK query status code
const StatusCodeOK = 200

// ItemsGetQuery query for get items
type ItemsGetQuery struct {
	TableName string `json:"table,omitempty"`

	// priority 0
	LoadAll bool `json:"load_all,omitempty"`
	// priority 1
	LoadAllCount bool `json:"load_all_count,omitempty"`

	// priority 1+2
	// Load Max value
	GetMinValue bool `json:"get_min_value,omitempty"`
	// Load Min value
	GetMaxValue bool `json:"get_max_value,omitempty"`

	// flag
	LoadShort bool `json:"load_short,omitempty"`

	// priority 3
	IKeys *[]int64  `json:"iks,omitempty"`
	SKeys *[]string `json:"sks,omitempty"`

	// priority 4
	// 0 - current value; -1 all; 1<= first limit values ask or desc ordered by key
	Limit     int    `json:"limit,omitempty"`
	OrderDesc bool   `json:"desc,omitempty"`
	IKey      int64  `json:"ik,omitempty"`
	SKey      string `json:"sk,omitempty"`

	//filter
	// Empty list or nil is all
	ShardKeys *[]int64 `json:"shard_keys,omitempty"`
}

// ItemsSetQuery query for set items
type ItemsSetQuery struct {
	TableName string `json:"table,omitempty"`

	// flag only on full object
	LoadFull bool `json:"load_full,omitempty"`

	// priority 0
	IItems *[]ItemInt         `json:"iitms,omitempty"`
	SItems *[]ItemString      `json:"sitms,omitempty"`
	Items  *[]json.RawMessage `json:"itms,omitempty"`

	// priority 1
	IItem *ItemInt         `json:"iitm,omitempty"`
	SItem *ItemString      `json:"sitm,omitempty"`
	Item  *json.RawMessage `json:"itm,omitempty"`
}

// ItemsGet items get result
type ItemsGet struct {
	ParamsErr   *Error `json:"params_err,omitempty"`
	InternalErr *Error `json:"internal_err,omitempty"`

	IItems []ItemInt    `json:"iitms,omitempty"`
	SItems []ItemString `json:"sitms,omitempty"`

	Count int `json:"cnt,omitempty"`
}

// QueryItemsGet send any query waiting ItemsGet
func (c *Connection) QueryItemsGet(path string, v interface{}) (res ItemsGet) {
	body, statusCode, err := c.DoQueryObject(path, v)
	if err != nil {
		res.InternalErr = ErrorE(err)
		return res
	}

	err = json.Unmarshal(body, &res)
	if err != nil {
		res.InternalErr = ErrorNew("Unmarshal fail = "+string(body)+", status_code = "+strconv.Itoa(statusCode), err)
		return res
	}

	if statusCode != StatusCodeOK {
		if res.InternalErr == nil && res.ParamsErr == nil {
			res.InternalErr = ErrorS("status_code = " + strconv.Itoa(statusCode))
		}
	}

	return res
}

// ItemsRawGet get IItems by query
func (c *Connection) ItemsRawGet(igq ItemsGetQuery) ItemsGet {
	return c.QueryItemsGet(URLGet, igq)
}

// ItemsRawSet set IItems by query
func (c *Connection) ItemsRawSet(isq ItemsSetQuery) ItemsGet {
	return c.QueryItemsGet(URLSet, isq)
}
