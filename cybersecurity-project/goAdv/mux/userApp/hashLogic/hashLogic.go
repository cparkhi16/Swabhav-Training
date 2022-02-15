package hashLogic

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

var salt string = "yogesh"

func HashString(s string) string {
	fmt.Println("hash string-", s)
	h := sha1.New()
	h.Write([]byte(s + salt))
	bs := hex.EncodeToString(h.Sum(nil))
	return bs
}
