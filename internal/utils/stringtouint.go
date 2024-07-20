package utils

import (
	"fmt"
	"strconv"
)

func StringToUint(id string) (uint, error) {
	num, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("error converting string to uint: %w", err)
	}
	uintNum := uint(num)
	return uintNum, nil
}
