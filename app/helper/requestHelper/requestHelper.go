package requestHelper

import (
	"errors"
	"strconv"
)

func CheckIDInt(id string) (int, error) {
	// check id
	// conv to int
	if newId, err := strconv.Atoi(id); id != "" || err == nil {
		return newId, nil
	}
	return 0, errors.New("failed")
}
