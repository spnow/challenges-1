package main

import (
	"encoding/hex"
	"flag"
	"fmt"
)

func main() {
	var ip = flag.String("m", "4C6520666C6167206465206365206368616C6C656E6765206573743A203261633337363438316165353436636436383964356239313237356433323465", "hexadecimal string to decode")

	flag.Parse()

	x, err := hex.DecodeString(*ip)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(x))
}
