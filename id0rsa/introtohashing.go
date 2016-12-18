// created on 24 june 2016
// author: youenn.kahlys.laborde

package main

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	text := "id0-rsa.pub"

	// sha256
	h := sha256.New()
	h.Write([]byte(text))
	r1 := h.Sum(nil)

	// hex encoding
	r2 := hex.EncodeToString(r1)

	// md5
	r3 := md5.Sum([]byte(r2))

	// hex encoding
	r4 := hex.EncodeToString(r3[:])
	fmt.Println(r4)
}
