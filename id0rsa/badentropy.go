// created on 26 june 2016
// author: youenn.kahlys.laborde

package main

import (
	"crypto/aes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	startDate := time.Date(2016, time.January, 25, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2016, time.February, 1, 0, 0, 0, 0, time.UTC)
	start := startDate.Unix()
	end := endDate.Unix()
	text, _ := hex.DecodeString("a99210d796a1e37503febf65c329c1b2")
	Decrypt(text, start, end)
}

// Brute force aes decryption
func Decrypt(text []byte, start, end int64) {

	var res string
	var scr float64

	for i := start; i <= end; i++ {
		key := md5.Sum([]byte(strconv.FormatInt(i, 10)))

		text := string(aesEcbDecrypt(text, key[:]))
		plaintext := strings.ToUpper(text)

		if score(plaintext) > scr {
			res = text
			scr = score(plaintext)
		}
	}
	fmt.Println(res)
}

// Aes Ecb Decryption
func aesEcbDecrypt(ciphertext, key []byte) []byte {
	cipher, _ := aes.NewCipher([]byte(key))
	bs := 16

	i := 0
	plaintext := make([]byte, len(ciphertext))
	finalplaintext := make([]byte, len(ciphertext))
	for len(ciphertext) > 0 {
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

// In order to recognize english
var freq map[rune]float64

func init() {
	freq = make(map[rune]float64)
	freq['A'] = 0.0651738
	freq['B'] = 0.0124248
	freq['C'] = 0.0217339
	freq['D'] = 0.0349835
	freq['E'] = 0.1041442
	freq['F'] = 0.0197881
	freq['G'] = 0.0158610
	freq['H'] = 0.0492888
	freq['I'] = 0.0558094
	freq['J'] = 0.0009033
	freq['K'] = 0.0050529
	freq['L'] = 0.0331490
	freq['M'] = 0.0202124
	freq['N'] = 0.0564513
	freq['O'] = 0.0596302
	freq['P'] = 0.0137645
	freq['Q'] = 0.0008606
	freq['R'] = 0.0497563
	freq['S'] = 0.0515760
	freq['T'] = 0.0729357
	freq['U'] = 0.0225134
	freq['V'] = 0.0082903
	freq['W'] = 0.0171272
	freq['X'] = 0.0013692
	freq['Y'] = 0.0145984
	freq['Z'] = 0.0007836
}

func score(in string) (s float64) {
	for _, v := range in {
		s += freq[v]
	}
	return
}
