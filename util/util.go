package util

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

// Check error
func Check(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

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
