package infoschema

import (
	"database/sql"
	"encoding/json"
)

// JSONNullInt64 is a Nullable Json SQL field type
type JSONNullInt64 sql.NullInt64

// MarshalJSON used for marshalling the JsonNullInt64 type
func (v JSONNullInt64) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Int64)
	}

	return json.Marshal(nil)
}

func (v *JSONNullInt64) UnmarshalJSON(data []byte) error {
	var x *int64
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}

	v.Valid = false

	if x != nil {
		v.Valid = true
		v.Int64 = *x
	}

	return nil
}

type JSONNullString sql.NullString

func (v JSONNullString) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.String)
	}
	return json.Marshal(nil)
}

func (v *JSONNullString) UnmarshalJSON(data []byte) error {
	var s *string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	v.Valid = false

	if s != nil {
		v.Valid = true
		v.String = *s
	}

	return nil
}
