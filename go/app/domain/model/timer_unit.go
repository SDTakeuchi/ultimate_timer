/*
unused for now
*/

package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type TimerUnit json.RawMessage

func (tu *TimerUnit) Scan(val interface{}) error {
	bytes, ok := val.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", val))
	}
	if err := json.Unmarshal(bytes, &tu); err != nil {
		return err
	}
	return nil
}

func (tu TimerUnit) Value() (driver.Value, error) {
	if len(tu) == 0 {
		return nil, nil
	}
	return json.RawMessage(tu).MarshalJSON()
}
