package gutils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"testing"
)

func Test_GenerateRSAKey(t *testing.T) {
	d := md5.Sum([]byte("ass"))
	fmt.Println(hex.EncodeToString(d[:]))
	p, pr, _ := GenerateRSAKey(2048)
	fmt.Println(string(p))
	fmt.Println(string(pr))
}
