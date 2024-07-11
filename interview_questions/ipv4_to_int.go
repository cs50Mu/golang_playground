package main

import (
	"fmt"
	"strconv"
	"strings"
)

func ipv4ToInt(s string) (uint32, error) {
	splitted := strings.Split(s, ".")
	if len(splitted) != 4 {
		return 0, fmt.Errorf("not a ipv4 addr")
	}

	var ints []uint32
	for _, s := range splitted {
		// converted, err := stringToInt32(s)
		converted, err := stringToInt8V2(s)
		if err != nil {
			return 0, err
		}
		ints = append(ints, uint32(converted))
	}

	ints[0] = ints[0] << 24
	ints[1] = ints[1] << 16
	ints[2] = ints[2] << 8

	return ints[0] | ints[1] | ints[2] | ints[3], nil
}

// ref: https://stackoverflow.com/questions/30299649/parse-string-to-specific-type-of-int-int8-int16-int32-int64
func stringToInt8(s string) (uint8, error) {
	s = strings.Trim(s, " ")
	converted, err := strconv.ParseUint(s, 10, 8)
	if err != nil {
		return 0, err
	}
	return uint8(converted), nil
}

// Sscan is a little easy to understand
func stringToInt8V2(s string) (uint8, error) {
	s = strings.Trim(s, " ")
	var res uint8
	_, err := fmt.Sscan(s, &res)
	if err != nil {
		return 0, err
	}
	return res, nil
}
