package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	f, _ := os.Open("../README.md")
	data, _ := ioutil.ReadAll(f)

	sha := sha256.New()
	sha.Write(data)
	s := sha.Sum(nil)
	fmt.Println(hex.EncodeToString(s))
}
