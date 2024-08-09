package utils

import "encoding/json"

func IsJson(s string) error {
	var data interface{}

	if err := json.Unmarshal([]byte(s), &data); err != nil {
		return err
	}

	return nil
}
