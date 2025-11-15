package md5

import (
	"crypto/md5"
	"encoding/hex"
)

func GenMD5(stuID string) string {
	hash := md5.New()
	hash.Write([]byte(stuID))
	hashBytes := hash.Sum(nil)
	return hex.EncodeToString(hashBytes)
}
