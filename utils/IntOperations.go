package utils

import "strconv"

func IntegerToHex(i int64) []byte {
	res := strconv.FormatInt(i, 16)
	return []byte(res)
}
