package json

import (
	"bytes"
	"encoding/json"
	"strconv"
)

type JSONWriter bytes.Buffer

func NewJSONWriter() *JSONWriter {
	return new(JSONWriter)
}

func (w *JSONWriter) ToBytes() []byte {
	return (*bytes.Buffer)(w).Bytes()
}

func (w *JSONWriter) ToString() string {
	return (*bytes.Buffer)(w).String()
}

func (w *JSONWriter) String(s string) {
	writeString((*bytes.Buffer)(w), s, true)
}

func (w *JSONWriter) Int64(i int64) {
	(*bytes.Buffer)(w).WriteString(strconv.FormatInt(i, 10))
}

func (w *JSONWriter) Bool(b bool) {
	(*bytes.Buffer)(w).WriteString(strconv.FormatBool(b))
}

func (w *JSONWriter) Object(write func(w *JSONWriter)) {
	buf := (*bytes.Buffer)(w)
	buf.WriteByte('{')
	write(w)
	buf.WriteByte('}')
}

func (w *JSONWriter) Array(write func(w *JSONWriter)) {
	buf := (*bytes.Buffer)(w)
	buf.WriteByte('[')
	write(w)
	buf.WriteByte(']')
}

func (w *JSONWriter) Any(v interface{}) {
	enc := json.NewEncoder((*bytes.Buffer)(w))
	enc.Encode(v)
}

func (w *JSONWriter) Comma() {
	(*bytes.Buffer)(w).WriteByte(',')
}

func (w *JSONWriter) Null() {
	(*bytes.Buffer)(w).WriteString("null")
}

func (w *JSONWriter) Property(name string, isLast bool, write func(w *JSONWriter)) {
	w.String(name)
	buf := (*bytes.Buffer)(w)
	buf.WriteByte(':')
	write(w)
	if !isLast {
		buf.WriteByte(',')
	}
}

func (w *JSONWriter) StringProperty(name string, value string, isLast bool) {
	w.Property(name, isLast, func(w *JSONWriter) {
		w.String(value)
	})
}

func (w *JSONWriter) Int64Property(name string, value int64, isLast bool) {
	w.Property(name, isLast, func(w *JSONWriter) {
		w.Int64(value)
	})
}

func (w *JSONWriter) BoolProperty(name string, value bool, isLast bool) {
	w.Property(name, isLast, func(w *JSONWriter) {
		w.Bool(value)
	})
}

func (w *JSONWriter) AnyProperty(name string, value interface{}, isLast bool) {
	w.Property(name, isLast, func(w *JSONWriter) {
		w.Any(value)
	})
}

func (w *JSONWriter) ObjectProperty(name string, isLast bool, write func(w *JSONWriter)) {
	w.Property(name, isLast, func(w *JSONWriter) {
		w.Object(write)
	})
}

func (w *JSONWriter) ArrayProperty(name string, isLast bool, write func(w *JSONWriter)) {
	w.Property(name, isLast, func(w *JSONWriter) {
		w.Array(write)
	})
}
