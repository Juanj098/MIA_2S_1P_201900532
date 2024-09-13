package utils

import "errors"

func ConvertTobytes(size int, unit string) (int, error) {
	switch unit {
	case "K":
		return size * 1024, nil
	case "M":
		return size * 1024 * 1024, nil
	default:
		return 0, errors.New("invalid Unit")
	}
}
