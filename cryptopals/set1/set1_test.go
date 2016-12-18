// Created on 10 June 2016
// author: youenn.laborde

package set1

import (
	"encoding/hex"
	"testing"
)

import "fmt"

var _ = fmt.Printf

func TestChallenge1(t *testing.T) {
	hex := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	b64 := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

	res, err := HexToBase64(hex)
	if err != nil {
		t.Error(err)
	}

	if res != b64 {
		t.Error("Expected", b64, "got", res)
	}
}

func TestChallenge2(t *testing.T) {
	res := "746865206b696420646f6e277420706c6179"
	a, _ := hex.DecodeString("1c0111001f010100061a024b53535009181c")
	b, _ := hex.DecodeString("686974207468652062756c6c277320657965")

	c, err := FixedXOR(a, b)
	if err != nil {
		t.Error(err)
	}

	r := hex.EncodeToString(c)
	if res != r {
		t.Error("Expected", r, "got", c)
	}
}

func TestChallenge3(t *testing.T) {
	a, _ := hex.DecodeString("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	k, _ := SingleByteXorCipher(a)

	if k != 'X' {
		t.Error("Expected", "X", "got", string(k))
	}
}

func TestChallenge5(t *testing.T) {
	res := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"
	text := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	key := "ICE"

	r := EncXOR([]byte(text), []byte(key))
	rh := hex.EncodeToString(r)

	if res != rh {
		t.Error("Expected", res, "got", rh)
	}
}
