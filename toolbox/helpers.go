package toolbox

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hash"
	"time"
)

type NotImplementedError struct {
	msg string
}

func (e *NotImplementedError) Error() string {
	if e.msg != "" {
		return "This feature has not been implemented."
	}
	return e.msg
}

// Return time in ms
func GetTime() int64 {
	return time.Now().UnixMilli()
}

// Returns SHA256 HMAC hash needed for sending requests
func GenerateHash(key, query string) hash.Hash {
	hmac := hmac.New(sha256.New, []byte(key))
	hmac.Write([]byte(query))
	return hmac
}

// Generates an unsorted query from a string[string] map, includes receive window,
// current timestamp and a signature from the included secret key.
func GenerateQuery(secretkey string, recvWindow int64, m map[string]string) string {
	ret := ""
	m["timestamp"] = fmt.Sprintf("%d", GetTime())
	m["recvWindow"] = fmt.Sprintf("%d", recvWindow)
	first := true

	for k, v := range m {
		if first {
			first = false
			ret = fmt.Sprintf("%s=%s", k, v)
		} else {
			ret += fmt.Sprintf("&%s=%s", k, v)
		}
	}

	hash := GenerateHash(secretkey, ret)
	ret += fmt.Sprintf("&signature=%s", hex.EncodeToString(hash.Sum(nil)))
	return ret
}
