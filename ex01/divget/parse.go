package divget

import (
	"strconv"
)

func parseURL(args []string) (url string, err error) {
	return args[1], nil
}

func invalidArg(args []string) error {
	if len(args) != 3 {
		return argErr
	}
	return nil
}

func parseDivNumber(args []string) (uint64, error) {
	return strconv.ParseUint(args[2], 10, 64)
}

func Parse(args []string) (string, uint64, error) {
	if err := invalidArg(args); err != nil {
		return noURL, noDivNumber, err
	}
	url, err := parseURL(args)
	if err != nil {
		return noURL, noDivNumber, err
	}
	divN, err := parseDivNumber(args)
	if err != nil {
		return noURL, noDivNumber, err
	}
	return url, divN, nil
}
