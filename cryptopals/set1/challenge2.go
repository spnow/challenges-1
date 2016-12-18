// Created on 10 June 2016
// author: youenn.kahlys.laborde

// Fixed XOR

package set1

import "errors"

func FixedXOR(a, b []byte) (r []byte, err error) {

	if len(a) != len(b) {
		err = errors.New("set1: not same lenght")
		return
	}

	r = make([]byte, len(a))
	for i, _ := range a {
		r[i] = a[i] ^ b[i]
	}

	return
}
