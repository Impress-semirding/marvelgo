package main

import (
	"crypto/md5"
	"encoding/hex"
)

func main() {

	h := md5.New()
	h.Write([]byte("abcd"))
	pass := hex.EncodeToString(h.Sum(nil))
	println(pass)


}
