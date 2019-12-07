package mdp

import (
	"encoding/json"
	"strconv"
)

// URLQueueGet url for get item(s) changes
const URLQueueGet = "get/ch/"

// ChangesGetQuery query for get changes items
type ChangesGetQuery struct {
	TableName string `json:"table,omitempty"`

	// priority 0
	Rv int64 `json:"rv,omitempty"`
	// 0 - default limit; 1<= first limit values ask or desc ordered by key
	Limit int `json:"limit,omitempty"`
}

// IItemsChangesGet items get changes result
type IItemsChangesGet struct {
	ParamsErr   *Error `json:"params_err,omitempty"`
	InternalErr *Error `json:"internal_err,omitempty"`

	Items []ItemInt `json:"itms,omitempty"`

	Count int `json:"cnt,omitempty"`

	MinRv int64 `json:"min_rv,omitempty"`
	MaxRv int64 `json:"max_rv,omitempty"`
}

// SItemsChangesGet items get changes result
type SItemsChangesGet struct {
	ParamsErr   *Error `json:"params_err,omitempty"`
	InternalErr *Error `json:"internal_err,omitempty"`

	Items []ItemString `json:"itms,omitempty"`

	MinRv int64 `json:"min_rv,omitempty"`
	MaxRv int64 `json:"max_rv,omitempty"`
}

// QueryIItemsChangesGet send any query waiting IItemsGet
func (c *Connection) QueryIItemsChangesGet(path string, v interface{}) (res IItemsChangesGet) {
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

// QuerySItemsChangesGet send any query waiting SItemsGet
func (c *Connection) QuerySItemsChangesGet(path string, v interface{}) (res SItemsChangesGet) {
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

// IItemsChangesRawGet get IItems by query
func (c *Connection) IItemsChangesRawGet(cgq ChangesGetQuery) IItemsChangesGet {
	return c.QueryIItemsChangesGet(URLQueueGet, cgq)
}

// SItemsChangesRawGet get IItems by query
func (c *Connection) SItemsChangesRawGet(cgq ChangesGetQuery) SItemsChangesGet {
	return c.QuerySItemsChangesGet(URLQueueGet, cgq)
}
