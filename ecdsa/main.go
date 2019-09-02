package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha512"
	"fmt"
	"log"
)

func main() {
	priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatal(err)
	}
	msg := []byte("A Very very secrect message")
	hash := sha512.Sum512(msg)

	r, s, err := ecdsa.Sign(rand.Reader, priv, hash[:])

	pub := &priv.PublicKey

	fmt.Println(ecdsa.Verify(pub, hash[:], r, s))

}
