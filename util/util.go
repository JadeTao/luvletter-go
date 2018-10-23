package util

import (
	"database/sql"
	"encoding/json"
	"path"
	"strings"
)

// NullString handle null string
type NullString struct {
	sql.NullString
}

// MarshalJSON implement for json encoding
func (v NullString) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.String)
	}
	return json.Marshal(nil)
}

// UnmarshalJSON implement for json decoding
func (v NullString) UnmarshalJSON(data []byte) error {
	var x *string
	if err := json.Unmarshal(data, &x); err != nil {
		v.Valid = false
		return err
	}
	if x != nil {
		v.Valid = true
		v.String = *x
	}
	return nil
}

// Split ...
func Split(flag rune) func(rune) bool {
	return func(s rune) bool {
		if s == flag {
			return true
		}
		return false
	}
}

// GetFileWithoutSuffix ...
func GetFileWithoutSuffix(filename string) string {
	suffix := path.Ext(filename)
	return strings.TrimSuffix(filename, suffix)
}

// ComparePath ...
func ComparePath(A string, B string) bool {
	splitedA := strings.FieldsFunc(A, Split('/'))
	splitedB := strings.FieldsFunc(B, Split('/'))

	for index, valueA := range splitedA {
		valueB := splitedB[index]
		if valueA != "*" && valueB != "*" && valueA != valueB {
			return false
		}
	}

	return true
}
