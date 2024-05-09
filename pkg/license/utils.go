package license

import (
	"encoding/base64"
	"errors"
)

func decodeKey(b []byte) ([]byte, error) {
	enc := base64.StdEncoding
	buf := make([]byte, enc.DecodedLen(len(b)))
	n, err := enc.Decode(buf, b)
	if err != nil {
		return nil, errors.New("illegal base64 data")
	}
	return buf[:n], nil
}
