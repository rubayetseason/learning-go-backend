package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
)

func main() {
	secret := []byte("this is secret")
	message := []byte("hello world")

	h := hmac.New(sha256.New, secret)
	h.Write(message)

	text := h.Sum(nil)

	fmt.Println(text)
}
