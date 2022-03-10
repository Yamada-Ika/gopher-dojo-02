package divget

import (
	"errors"
	"strconv"
)

func GetURL(args []string) (url string, err error) {
	if len(args) != 3 {
		return "", errors.New("error: Invalid argument")
	}
	return args[1], nil
}

func GetParallelN(args []string) (n uint64, err error) {
	if len(args) != 3 {
		return 0, errors.New("error: Invalid argument")
	}
	return strconv.ParseUint(args[2], 10, 64)
}
