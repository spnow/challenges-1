// Created on 3 July 2016
// author: youenn.kahlys.laborde

package set1

import (
	"crypto/aes"
)

// Aes Ecb Decryption
func AesEcbDecrypt(ciphertext, key []byte) []byte {
	cipher, _ := aes.NewCipher([]byte(key))
	bs := 16

	i := 0
	plaintext := make([]byte, len(ciphertext))
	finalplaintext := make([]byte, len(ciphertext))
	for len(ciphertext) > bs {
		cipher.Decrypt(plaintext, ciphertext)
		ciphertext = ciphertext[bs:]
		decryptedBlock := plaintext[:bs]
		for index, element := range decryptedBlock {
			finalplaintext[(i*bs)+index] = element
		}
		i++
		plaintext = plaintext[bs:]
	}
	return finalplaintext[:]
}
