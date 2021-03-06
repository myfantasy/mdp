package mdp

import (
	"encoding/json"
	"errors"
)

// ItemInt - one row with key string
type ItemInt struct {
	Key       int64           `json:"key,omitempty"`
	Data      json.RawMessage `json:"d,omitempty"`
	Rv        int64           `json:"rv,omitempty"`
	IsRemoved bool            `json:"rm,omitempty"`
	ShardKey  int64           `json:"sk,omitempty"`
}

// Less itm1 less then itm2
func (itm ItemInt) Less(itm2 ItemInt) bool {
	return itm.Key < itm2.Key
}

// ItemIntMakeFromJSONObject create ItemInt from json
func ItemIntMakeFromJSONObject(msg json.RawMessage, is *ItemStruct) (itm ItemInt, err error) {

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

// ItemIntStruct - one row with key string
type ItemIntStruct struct {
	Key           int64               `json:"key,omitempty"`
	FieldsInt     map[string]int64    `json:"fi,omitempty"`
	FieldsString  map[string]string   `json:"fs,omitempty"`
	FieldsIntA    map[string][]int64  `json:"fia,omitempty"`
	FieldsStringA map[string][]string `json:"fsa,omitempty"`
	Data          *[]byte             `json:"d,omitempty"`
	Rv            int64               `json:"rv,omitempty"`
	IsRemoved     bool                `json:"rm,omitempty"`
	ShardKey      int64               `json:"sk,omitempty"`
}

// ItemIntStat - one row with key string
type ItemIntStat struct {
	Key       int64 `json:"key,omitempty"`
	Rv        int64 `json:"rv,omitempty"`
	IsRemoved bool  `json:"rm,omitempty"`
	ShardKey  int64 `json:"sk,omitempty"`
}

// Stat - get stat object
func (itm ItemInt) Stat() ItemIntStat {
	return ItemIntStat{
		Key:       itm.Key,
		Rv:        itm.Rv,
		IsRemoved: itm.IsRemoved,
		ShardKey:  itm.ShardKey,
	}
}
