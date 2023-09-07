package format

import (
	"fmt"
	"testing"
)

func TestHexStringToASCII(t *testing.T) {
	s, err := hexStringToASCII("53:4d:32:53:69:67:6e:4b:65:79:31:39", ":")
	if err == nil {
		fmt.Println(s)
	} else {
		fmt.Println(err)
	}
}
