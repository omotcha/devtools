package format

import (
	"encoding/hex"
	"errors"
	"strings"
)

func hexStringToASCII(hexstr string, sep string) (string, error) {
	// remove separator
	splits := strings.Split(hexstr, sep)
	ret := ""
	for _, sp := range splits {
		// sanity check
		if len(sp) != 2 {
			return "", errors.New("not valid hex string")
		}
		// re-combine
		ret += sp
	}
	// decode
	bs, err := hex.DecodeString(ret)
	return string(bs), err
}
