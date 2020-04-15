package json

import (
	"bytes"
	"encoding/json"
)

// String
var (
	MissingRawString = RawString(nil)
	NullRawString    = RawString{nil}
)

// 这里用 slice，而不是用 struct 或者 pointer 实现，是因为只有 slice 能够同时在
// marshal 和 unmarshal 时区分 null 和 missing，其它实现方式只能在其中一种情况下
// 区分
type RawString []*string

func NewRawString(s *string, missing bool) RawString {
	var r RawString = make([]*string, 0, 1)
	if !missing {
		r = append(r, s)
	}
	return r
}

func (r RawString) String() string {
	if len(r) == 0 || r[0] == nil {
		return ""
	} else {
		return *r[0]
	}
}

func (r RawString) PString() *string {
	if len(r) == 0 || r[0] == nil {
		return nil
	} else {
		p := new(string)
		*p = *r[0]
		return p
	}
}

func (r RawString) Missing() bool {
	return len(r) == 0
}

func (r RawString) Null() bool {
	return len(r) > 0 && r[0] == nil
}

func (r RawString) MissingOrNull() bool {
	return len(r) == 0 || r[0] == nil
}

func (r RawString) Zero() bool {
	return len(r) == 0 || r[0] == nil || *r[0] == ""
}

func (r RawString) MarshalJSON() ([]byte, error) {
	w := NewJSONWriter()
	if r.Null() {
		w.Null()
	} else {
		w.String(*r[0])
	}
	return w.ToBytes(), nil
}

func (r *RawString) UnmarshalJSON(b []byte) error {
	if len(b) == 0 {
		return nil
	}
	slice := make([]*string, 1)
	*r = slice
	if bytes.Equal(b, Null) {
		return nil
	}
	slice[0] = new(string)
	return json.Unmarshal(b, slice[0])
}

var (
	MissingRawInt64 = RawInt64(nil)
	NullRawInt64    = RawInt64{nil}
)

// Int64
type RawInt64 []*int64

func NewRawInt64(i *int64, missing bool) RawInt64 {
	var r RawInt64 = make([]*int64, 0, 1)
	if !missing {
		r = append(r, i)
	}
	return r
}

func (r RawInt64) Int64() int64 {
	if len(r) == 0 || r[0] == nil {
		return 0
	} else {
		return *r[0]
	}
}

func (r RawInt64) PInt64() *int64 {
	if len(r) == 0 || r[0] == nil {
		return nil
	} else {
		p := new(int64)
		*p = *r[0]
		return p
	}
}

func (r RawInt64) Missing() bool {
	return len(r) == 0
}

func (r RawInt64) Null() bool {
	return len(r) > 0 && r[0] == nil
}

func (r RawInt64) MissingOrNull() bool {
	return len(r) == 0 || r[0] == nil
}

func (r RawInt64) Zero() bool {
	return len(r) == 0 || r[0] == nil || *r[0] == 0
}

func (r RawInt64) MarshalJSON() ([]byte, error) {
	w := NewJSONWriter()
	if r.Null() {
		w.Null()
	} else {
		w.Int64(*r[0])
	}
	return w.ToBytes(), nil
}

func (r *RawInt64) UnmarshalJSON(b []byte) error {
	if len(b) == 0 {
		return nil
	}
	slice := make([]*int64, 1)
	*r = slice
	if bytes.Equal(b, Null) {
		return nil
	}
	slice[0] = new(int64)
	return json.Unmarshal(b, slice[0])
}

// Interface
var (
	MissingRawInterface = RawInterface(nil)
	NullRawInterface    = RawInterface{nil}
)

type RawInterface []*interface{}

func NewRawInterface(v interface{}) RawInterface {
	return RawInterface{&v}
}

func (r RawInterface) Interface() interface{} {
	if len(r) == 0 || r[0] == nil {
		return nil
	} else {
		return *r[0]
	}
}

func (r RawInterface) Missing() bool {
	return len(r) == 0
}

func (r RawInterface) Null() bool {
	return len(r) > 0 && r[0] == nil
}

func (r RawInterface) MissingOrNull() bool {
	return len(r) == 0 || r[0] == nil
}

func (r RawInterface) Zero() bool {
	return len(r) == 0 || r[0] == nil || *r[0] == nil
}

func (r RawInterface) MarshalJSON() ([]byte, error) {
	if r.Null() {
		return Null[:], nil
	} else {
		return json.Marshal(*r[0])
	}
}

func (r *RawInterface) UnmarshalJSON(b []byte) error {
	if len(b) == 0 {
		return nil
	}
	slice := make([]*interface{}, 1)
	*r = slice
	if bytes.Equal(b, Null) {
		return nil
	}
	slice[0] = new(interface{})
	return json.Unmarshal(b, slice[0])
}
