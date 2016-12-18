// Created on 1 July 2016
// author: youenn.kahlys.laborde

/*
 * CHALLENGE 4 : Detect single-character XOR
 * Je decrypte chaque ligne avec la clef la plus propable, grace a la fonction
 * du challenge 3. Cette fonction me donne la clef et le clairs qui a le plus
 * gros score (somme des frequences de lettres). J'obtiens donc pour chaque ligne
 * le clair probable et leur clefs correspondante. Il me suffit de prendre le clair
 * qui a le plus grand score. Ce derniere seras en anglais, et est donc le clair
 * qui a été chiffré. Je peux donc retourner la ligne du chiffré, le clair, et
 * la clef.
 */

package set1

import (
	"bufio"
	"encoding/hex"
	"os"
)

// DetectSingleCharXOR detect a line encoded with a single char xor cipher. It
// returns the line, the decoded line and the key.
func DetectSingleCharXOR(file *os.File) (enc, dec string, key byte) {

	var sc float64

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		t := scanner.Text()
		data, err := hex.DecodeString(t)
		if err != nil {
			continue
		}

		k, o := SingleByteXorCipher(data)

		//fmt.Println(string(o))

		l := score(string(o))
		if l > sc {
			sc = l
			key = k
			enc = t
			dec = string(o)
		}
	}
	return
}
