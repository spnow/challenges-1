// Created on 1 July 2016
// author: youenn.kahlys.laborde

/*
 * CHALLENGE 5 : Implement repeating-key XOR
 */

package set1

func EncXOR(in, key []byte) (out []byte) {

	k := make([]byte, len(in))

	for i, _ := range k {
		k[i] = key[i%len(key)]
	}

	a, _ := FixedXOR(in, k)
	return a
}
