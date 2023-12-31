package toolbox

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hash"
	"time"
)

// Return time in ms
func GetTime() int64 {
	return time.Now().UnixMilli()
}

// Returns SHA256 HMAC hash
func GenerateHash(key, query string) hash.Hash {
	hmac := hmac.New(sha256.New, []byte(key))
	hmac.Write([]byte(query))
	return hmac
}

func GenerateQuery(secretkey string, m map[string]string) string {
	ret := ""
	m["timestamp"] = fmt.Sprintf("%d", GetTime())
	first := true

	for k, v := range m {
		if first {
			first = false
			ret = fmt.Sprintf("%s=%s", k, v)
			fmt.Println("")
			fmt.Println(ret)
			fmt.Println("")
		} else {
			ret += fmt.Sprintf("&%s=%s", k, v)
			fmt.Println("")
			fmt.Println(ret)
			fmt.Println("")
		}
	}

	hash := GenerateHash(secretkey, ret)
	ret += fmt.Sprintf("&signature=%s", hex.EncodeToString(hash.Sum(nil)))
	return ret
}
