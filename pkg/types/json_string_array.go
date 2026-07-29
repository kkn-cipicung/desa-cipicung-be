package types

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type JSONStringArray []string

func (a *JSONStringArray) Scan(value any) error {
	if value == nil {
		*a = JSONStringArray{}
		return nil
	}

	var raw []byte
	switch v := value.(type) {
	case []byte:
		raw = v
	case string:
		raw = []byte(v)
	default:
		return fmt.Errorf("scan JSONStringArray: unsupported type %T", value)
	}

	if len(raw) == 0 {
		*a = JSONStringArray{}
		return nil
	}

	var values []string
	if err := json.Unmarshal(raw, &values); err != nil {
		return fmt.Errorf("scan JSONStringArray: %w", err)
	}
	*a = values
	return nil
}

func (a JSONStringArray) Value() (driver.Value, error) {
	if a == nil {
		return "[]", nil
	}
	data, err := json.Marshal([]string(a))
	if err != nil {
		return nil, err
	}
	return string(data), nil
}
