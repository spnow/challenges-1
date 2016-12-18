// Created on 1 July 2016
// author: youenn.kahlys.laborde

// CHALLENGE 6 : Break repeating-key XOR
//
//      Let KEYSIZE be the guessed length of the key; try values from 2 to (say) 40.
//      Write a function to compute the edit distance/Hamming distance between two
//      strings. The Hamming distance is just the number of differing bits.
//
//      For each KEYSIZE, take the first KEYSIZE worth of bytes, and the second
//      KEYSIZE worth of bytes, and find the edit distance between them. Normalize
//      this result by dividing by KEYSIZE.
//
//      The KEYSIZE with the smallest normalized edit distance is probably the key.
//      You could proceed perhaps with the smallest 2-3 KEYSIZE values. Or take 4
//      KEYSIZE blocks instead of 2 and average the distances.
//
//      Now that you probably know the KEYSIZE: break the ciphertext into blocks of
//      KEYSIZE length.
//
//      Now transpose the blocks: make a block that is the first byte of every block,
//      and a block that is the second byte of every block, and so on.
//
//      Solve each block as if it was single-character XOR. You already have code to do
//      this.
//
//      For each block, the single-byte XOR key that produces the best looking histogram
//      is the repeating-key XOR key byte for that block. Put them together and you have
//      the key.

package set1

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	maxKeyLen       = 40
	nProbableKeyLen = 3
)

// DetectSingleCharXOR detect a line encoded with a single char xor cipher. It
// returns the line, the decoded line and the key.
func BreakEncXOR(file *os.File) (out, key []byte) {

	datas, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	_, err = base64.StdEncoding.Decode(datas, datas)
	if err != nil {
		fmt.Println(err)
	}

	// Find KeyLen from 2 to 40
	var keyLen int
	var tmpDist float64
	for tmpLen := 2; tmpLen < maxKeyLen; tmpLen++ {
		c := averageHamming(datas, tmpLen)

		if c < tmpDist || tmpDist == 0 {
			keyLen = tmpLen
			tmpDist = c
		}
	}

	// subtext
	var subText = make(map[int][]byte)
	for i, v := range datas {
		k := i % keyLen
		subText[k] = append(subText[k], v)
	}

	key = make([]byte, keyLen)
	for i, _ := range key {
		//fmt.Println(subText[i])
		key[i], _ = SingleByteXorCipher(subText[i])
	}
	out = EncXOR(datas, key)
	return
}

func Hamming(a, b []byte) int {
	switch bytes.Compare(a, b) {
	case 0: // a == b
	case 1: // a > b
		temp := make([]byte, len(a))
		copy(temp, b)
		b = temp
	case -1: // a < b
		temp := make([]byte, len(b))
		copy(temp, a)
		a = temp
	}
	d := 0
	for i, x := range a {
		y := b[i]
		xor := x ^ y
		for y := xor; y > 0; y >>= 1 {
			// check if lowest bit is 1
			if int(y&1) == 1 {
				d++
			}
		}
	}
	return d
}

func averageHamming(in []byte, s int) float64 {

	var d []int
	for len(in) >= 2*s {
		d = append(d, Hamming(in[:s], in[s:2*s])/s)
		in = in[2*s:]
	}

	var t int
	for _, v := range d {
		t = t + v
	}

	return float64(t) / float64(len(d))
}
