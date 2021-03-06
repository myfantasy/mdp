package mdp

import (
	"encoding/json"
	"errors"
)

// ItemString - one row with key string
type ItemString struct {
	Key       string          `json:"key,omitempty"`
	Data      json.RawMessage `json:"d,omitempty"`
	Rv        int64           `json:"rv,omitempty"`
	IsRemoved bool            `json:"rm,omitempty"`
	ShardKey  int64           `json:"sk,omitempty"`
}

// Less itm1 less then itm2
func (itm ItemString) Less(itm2 ItemString) bool {
	return itm.Key < itm2.Key
}

// ItemStringMakeFromJSONObject create ItemInt from json
func ItemStringMakeFromJSONObject(msg json.RawMessage, is *ItemStruct) (itm ItemString, err error) {

	m := make(map[string]json.RawMessage)

	err = json.Unmarshal(msg, &m)

	if err != nil {
		return itm, err
	}

	var v json.RawMessage
	var ok bool

	if is == nil {
		v, ok = m["key"]
	} else {
		v, ok = m[is.KeyName]
	}

	if !ok && is == nil {
		return itm, errors.New("key not found")
	}
	if !ok {
		return itm, errors.New("key (" + is.KeyName + ") not found")
	}

	err = json.Unmarshal(v, &itm.Key)

	if err != nil {
		return itm, ErrorNew("key unmurshal fail", err)
	}

	if is == nil {
		v, ok = m["sk"]
	} else {
		v, ok = m[is.ShardKeyName]
	}

	if ok {
		err = json.Unmarshal(v, &itm.ShardKey)
		if err != nil {
			return itm, ErrorNew("key unmurshal fail", err)
		}
	}

	itm.Data = msg
	return itm, err

}

// ItemStringStruct - one row with key string
type ItemStringStruct struct {
	Key           string              `json:"key,omitempty"`
	FieldsInt     map[string]int64    `json:"fi,omitempty"`
	FieldsString  map[string]string   `json:"fs,omitempty"`
	FieldsIntA    map[string][]int64  `json:"fia,omitempty"`
	FieldsStringA map[string][]string `json:"fsa,omitempty"`
	Data          *[]byte             `json:"d,omitempty"`
	Rv            int64               `json:"rv,omitempty"`
	IsRemoved     bool                `json:"rm,omitempty"`
	ShardKey      int64               `json:"sk,omitempty"`
}

// ItemStringStat - one row with key string
type ItemStringStat struct {
	Key       string `json:"key,omitempty"`
	Rv        int64  `json:"rv,omitempty"`
	IsRemoved bool   `json:"rm,omitempty"`
	ShardKey  int64  `json:"sk,omitempty"`
}

// Stat - get stat object
func (itm ItemString) Stat() ItemStringStat {
	return ItemStringStat{
		Key:       itm.Key,
		Rv:        itm.Rv,
		IsRemoved: itm.IsRemoved,
		ShardKey:  itm.ShardKey,
	}
}
