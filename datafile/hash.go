package datafile

import (
	"crypto/md5"
	"encoding/hex"
	"os"
)

func GetHash(filename string) string {
	sourse, err := os.ReadFile(filename)
	if err != nil {
		return ""
	}
	data := string(sourse)
	hash := md5.New()
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil))
}
