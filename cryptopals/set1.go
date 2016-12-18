package main

import (
	// "encoding/base64"
	"fmt"
	"github.com/kahlys/cryptopals/set1"
	// "io/ioutil"
	"os"
)

func main() {
	/* challenge 4
	 * find the line which has been encrypted by single character xor.
	 */
	// file, _ := os.Open("chall4.txt")
	// defer file.Close()
	// line, dec, k := set1.DetectSingleCharXOR(file)
	// fmt.Println(line, dec, k)

	/* challenge 6
	 *
	 */
	// file, _ := os.Open("chall6.txt")
	// defer file.Close()
	// _, key := set1.BreakEncXOR(file)
	// fmt.Println(string(key))

	/* challenge 7
	 * aes-ecb decryption
	 */
	// file, _ := os.Open("chall7.txt")
	// defer file.Close()
	// datas, err := ioutil.ReadAll(file)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// _, err = base64.StdEncoding.Decode(datas, datas)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// key := []byte("YELLOW SUBMARINE")
	// out := set1.AesEcbDecrypt(datas, key)
	// fmt.Println(string(out))

	/* challenge 8
	 *
	 */
	file, _ := os.Open("chall8.txt")
	defer file.Close()
	list := set1.DetectAesEcb(file)
	fmt.Println(list)
}
