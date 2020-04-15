package json

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type JSONReader struct {
	Name  string
	Index int
	Value json.RawMessage

	objectValue map[string]json.RawMessage
}

func NewJSONReader(b []byte) *JSONReader {
	raw := json.RawMessage(b)
	return &JSONReader{
		Value: raw,
	}
}

func (r *JSONReader) String() string {
	var s string
	if err := json.Unmarshal(r.Value, &s); err != nil {
		panic(fmt.Sprintf("Value of '%s' must be string", r.Name))
	}
	return s
}

func (r *JSONReader) Int64() int64 {
	var i int64
	if err := json.Unmarshal(r.Value, &i); err != nil {
		panic(fmt.Sprintf("Value of '%s' must be int64", r.Name))
	}
	return i
}

func (r *JSONReader) Bool() bool {
	var b bool
	if err := json.Unmarshal(r.Value, &b); err != nil {
		panic(fmt.Sprintf("Value of '%s' must be bool", r.Name))
	}
	return b
}

func (r *JSONReader) IterateArray(iter func(r *JSONReader)) (isEmpty bool) {
	var array []json.RawMessage
	if err := json.Unmarshal(r.Value, &array); err != nil {
		panic(fmt.Sprintf("Value of '%s' must be array", r.Name))
	}
	if len(array) == 0 {
		return true
	}
	for index, item := range array {
		reader := &JSONReader{
			Index: index,
			Value: item,
		}
		iter(reader)
	}
	return false
}

func (r *JSONReader) IterateObject(iter func(r *JSONReader)) (isEmpty bool) {
	if r.objectValue == nil {
		if err := json.Unmarshal(r.Value, &r.objectValue); err != nil {
			panic(fmt.Sprintf("Value of '%s' must be object", r.Name))
		}
	}
	if len(r.objectValue) == 0 {
		return true
	}
	index := 0
	for name, item := range r.objectValue {
		reader := &JSONReader{
			Name:  name,
			Index: index,
			Value: item,
		}
		iter(reader)
		index++
	}
	return false
}

func (r *JSONReader) Any(v interface{}) {
	if err := json.Unmarshal(r.Value, v); err != nil {
		panic(fmt.Sprintf("Value of '%s' must be %s", r.Name, reflect.TypeOf(v).Elem()))
	}
}

func (r *JSONReader) Property(name string) *JSONReader {
	if r.objectValue == nil {
		if err := json.Unmarshal(r.Value, &r.objectValue); err != nil {
			panic(fmt.Sprintf("Value of '%s' must be object", r.Name))
		}
	}
	if item, ok := r.objectValue[name]; ok {
		reader := &JSONReader{
			Name:  name,
			Index: -1,
			Value: item,
		}
		return reader
	} else {
		panic(fmt.Sprintf("Object '%s' must have property '%s'", r.Name, name))
	}
}

func (r *JSONReader) StringProperty(name string) string {
	return r.Property(name).String()
}

func (r *JSONReader) Int64Property(name string) int64 {
	return r.Property(name).Int64()
}

func (r *JSONReader) BoolProperty(name string) bool {
	return r.Property(name).Bool()
}

func (r *JSONReader) AnyProperty(name string, v interface{}) {
	r.Property(name).Any(v)
}
