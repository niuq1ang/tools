package json

import (
	"bytes"
	"encoding/json"
)

var (
	Null        = json.RawMessage("null")
	EmptyStr    = json.RawMessage(`""`)
	Zero        = json.RawMessage("0")
	EmptyArray  = json.RawMessage("[]")
	EmptyObject = json.RawMessage("{}")
)

func IsNull(r json.RawMessage) bool {
	return bytes.Equal(Null, r)
}

func IsMissing(r json.RawMessage) bool {
	return len(r) == 0
}

func IsMissingOrNull(r json.RawMessage) bool {
	return len(r) == 0 || bytes.Equal(Null, r)
}
