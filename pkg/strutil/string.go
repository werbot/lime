package strutil

import (
	"strconv"
	"strings"
)

// ToInt32 is convert string to int32
func ToInt32(s string) int32 {
	parsed, err := strconv.ParseInt(strings.TrimSpace(s), 10, 32)
	if err != nil {
		return 256
	}
	return int32(parsed)
}
