// Created on 1 July 2016
// author: youenn.kahlys.laborde

/*
 * CHALLENGE 4 : Detect single-character XOR
 * The problem with ECB is that it is stateless and deterministic; the same 16
 * byte plaintext block will always produce the same 16 byte ciphertext.
 */

package set1

import (
	"bufio"
	"os"
	"strings"
)

const (
	blocksize = 16
)

// DetectSingleCharXOR detect a line encoded with a single char xor cipher. It
// returns the line, the decoded line and the key.
func DetectAesEcb(file *os.File) (out []string) {

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var ligne = 1
	for scanner.Scan() {
		t := scanner.Text()

		var blockToCount = make(map[string]int)
		for i := 0; i < len(t)/blocksize; i++ {
			j := i * blocksize
			c := strings.Count(t, t[j:j+blocksize])
			if c > 1 {
				blockToCount[t[j:j+blocksize]] = c
			}
		}

		if len(blockToCount) > 0 {
			out = append(out, t)
		}

		ligne++
	}
	return
}
