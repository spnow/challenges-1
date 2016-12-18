// created on 24 june 2016
// author: youenn.kahlys.laborde

package main

import (
	"fmt"
	"strings"
)

// bruteforce tries all shift possible
func bruteforce(in string) {
	for i := 0; i < 26; i++ {
		res := encode(in, i)
		// I thinks the result will be an English text with probably usual words as "the" and "is".
		if strings.Contains(res, "THE") && strings.Contains(res, "IS") {
			fmt.Println("->", res)
		}
	}
}

func encode(in string, shift int) (out string) {
	in = strings.ToUpper(in)
	for i := 0; i < len(in); i++ {
		j := (int(in[i]) + shift)
		if j > 90 {
			j = j - 26
		}
		out = out + string(j)
	}
	return
}

func main() {
	bruteforce("ZNKIGKYGXIOVNKXOYGXKGRREURJIOVNKXCNOINOYXKGRRECKGQOSTUZYAXKNUCURJHKIGAYKOSZUURGFEZURUUQGZZNKCOQOVGMKGZZNKSUSKTZHAZOLOMAXKOZYMUZZUHKGZRKGYZROQKLOLZEEKGXYURJUXCNGZKBKXBGPJADLIVBAYKZNUYKRGYZZKTINGXGIZKXYGYZNKYURAZOUT")
}
