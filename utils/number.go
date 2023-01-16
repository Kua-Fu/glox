package utils

import (
	"encoding/json"
	"fmt"
)

// GetFloatNumber get float number
func GetFloatNumber(n interface{}) (float64, error) {
	var (
		res float64
	)
	// StringLiteral, Identifier
	switch n.(type) {
	case json.Number:
		v, ok := n.(json.Number)
		if ok {
			fv, err := v.Float64()
			if err != nil {

			} else {
				res = fv
			}

		}
	case int64:
		res = float64(n.(int64))
	case int:
		res = float64(n.(int))
	case float64:
		res = n.(float64)
	case float32:
		res = float64(n.(float32))
	default:
		return res, fmt.Errorf(
			"%v, can not convert to float64",
			n)

	}
	return res, nil

}
