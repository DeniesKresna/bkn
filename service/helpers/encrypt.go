package helpers

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"strings"
)

func EncryptMD5(stringlist ...string) (code string) {
	stringjoin := strings.Join(stringlist, "_")
	hash := md5.Sum([]byte(stringjoin))
	code = hex.EncodeToString(hash[:])
	return code
}

func Base64Encode(input string) string {
	encoded := base64.StdEncoding.EncodeToString([]byte(input))
	return encoded
}
