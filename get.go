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
