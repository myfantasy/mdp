package mdp

import (
	"encoding/json"
	"errors"

	"github.com/valyala/fasthttp"
)

// URLGetI url for get int key tables
const URLGetI = "i/get/"

// URLGetS url for get string key tables
const URLGetS = "s/get/"

// IItemsGet items get result
type IItemsGet struct {
	ParamsErr   json.RawMessage `json:"params_err,omitempty"`
	InternalErr json.RawMessage `json:"internal_err,omitempty"`

	Items []ItemInt `json:"itms,omitempty"`
	Item  *ItemInt  `json:"itm,omitempty"`
}

// SItemsGet items get result
type SItemsGet struct {
	ParamsErr   json.RawMessage `json:"params_err,omitempty"`
	InternalErr json.RawMessage `json:"internal_err,omitempty"`

	Items []ItemString `json:"itms,omitempty"`
	Item  *ItemString  `json:"itm,omitempty"`
}

// AllDataIGet Get all data from int key table
func (c *Connection) AllDataIGet(table string) (ig *IItemsGet, statusCode int, err error) {

	q := make(map[string]interface{})

	q["all"] = true
	q["tbl_name"] = table

	body, statusCode, err := c.DoQueryDict(URLGetI, q)

	if err != nil {
		return ig, 0, err
	}

	ig = &IItemsGet{}

	err = json.Unmarshal(body, ig)

	if statusCode != fasthttp.StatusOK {
		return ig, statusCode, errors.New("Fail status")
	}

	return ig, statusCode, err
}

// AllDataSGet Get all data from string key table
func (c *Connection) AllDataSGet(table string) (ig *SItemsGet, statusCode int, err error) {

	q := make(map[string]interface{})

	q["all"] = true
	q["tbl_name"] = table

	body, statusCode, err := c.DoQueryDict(URLGetS, q)

	if err != nil {
		return ig, 0, err
	}

	ig = &SItemsGet{}

	err = json.Unmarshal(body, ig)

	if statusCode != fasthttp.StatusOK {
		return ig, statusCode, errors.New("Fail status")
	}

	return ig, statusCode, err
}

// AllDataIGetE Get all data from int key table
func (c *Connection) AllDataIGetE(table string) (itms []ItemInt, statusCode int, parErr error, intErr error) {

	ig, statusCode, err := c.AllDataIGet(table)

	if ig != nil {
		if ig.Item != nil {
			itms = append(itms, *ig.Item)
		} else {
			itms = ig.Items
		}
	}
	if statusCode == 0 {
		return itms, 500, parErr, err
	}

	if statusCode >= 500 {
		return itms, statusCode, parErr, ErrorNew(string(ig.InternalErr), err)
	}

	if statusCode >= 400 {
		return itms, statusCode, ErrorNew(string(ig.ParamsErr), err), intErr
	}

	return itms, statusCode, nil, nil
}

// AllDataSGetE Get all data from string key table
func (c *Connection) AllDataSGetE(table string) (itms []ItemString, statusCode int, parErr error, intErr error) {

	ig, statusCode, err := c.AllDataSGet(table)

	if ig != nil {
		if ig.Item != nil {
			itms = append(itms, *ig.Item)
		} else {
			itms = ig.Items
		}
	}
	if statusCode == 0 {
		return itms, 500, parErr, err
	}

	if statusCode >= 500 {
		return itms, statusCode, parErr, ErrorNew(string(ig.InternalErr), err)
	}

	if statusCode >= 400 {
		return itms, statusCode, ErrorNew(string(ig.ParamsErr), err), intErr
	}

	return itms, statusCode, nil, nil
}
