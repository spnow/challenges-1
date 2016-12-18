// Created on 10 June 2016
// author: youenn.kahlys.laborde

// Convert hex to base64

package set1

import (
	"encoding/base64"
	"encoding/hex"
)

func HexToBase64(in string) (out string, err error) {
	data, err := hex.DecodeString(in)
	if err != nil {
		return
	}
	out = base64.StdEncoding.EncodeToString(data)
	return
}
