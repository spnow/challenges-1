// Created on 10 June 2016
// author: youenn.laborde

/* CHALLENGE 3 : Single-byte XOR cipher
 * Je prend le chiffré et je le déchifre avec un XOR-CIPHER pour toute les clefs
 * possible. C'est à dire tout les octets de 0 à 255. Je prends le résultat
 * qui a le plus grand score (somme des fréquences des lettres). Ce dernier
 * est un texte en anglais, donc le clair qui a été chiffré.
 */

package set1

import "strings"

//import "fmt"

// English character frequency from http://www.data-compression.com/english.html
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
	freq[' '] = 0.1918182
}

// Decode an English plaintext xor'd with a single character
func SingleByteXorCipher(in []byte) (key byte, out []byte) {
	var sc float64

	var k byte
	for k = 1; k <= 255; k++ {

		x := make([]byte, len(in))
		for i, _ := range x {
			x[i] = byte(k)
		}

		tmp, err := FixedXOR(in, x)
		if err != nil {
			return
		}

		l := score((string(tmp)))
		if l > sc {
			sc = l
			key = k
			out = tmp

			//fmt.Println("\n== \nscore:", l ,k , "\n",strings.ToUpper(string(out)),"\n==\n")
		}

		if k == 255 {
			break
		}
	}
	return
}

func score(in string) (s float64) {

	in = strings.ToUpper(in)

	for _, v := range in {
		s += freq[v]
	}
	return
}
