package util

import (
"bytes"
"io/ioutil"

"golang.org/x/text/encoding/simplifiedchinese"
"golang.org/x/text/transform"
)

var GBKEncoder = simplifiedchinese.GBK.NewEncoder()

func UTF8ToGBK(from []byte) (to []byte, err error) {
	reader := transform.NewReader(bytes.NewReader(from), GBKEncoder)
	return ioutil.ReadAll(reader)
}