package braintree

import (
	"crypto/hmac"
	"crypto/sha1"
	"fmt"
	"io"
)

func Hexdigest(key, payload string) string {
	h := sha1.New()
	io.WriteString(h, key)
	mac := hmac.New(sha1.New, h.Sum(nil))
	mac.Write([]byte(payload))
	return fmt.Sprintf("%x", mac.Sum(nil))
}

func SecureCompare(left, right string) bool {
	if len(left) != len(right) {
		return false
	}

	return left == right
}
