package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"time"
)

func main() {
	key := []byte("somekeysomekeysomekeysomekeyabcd")
	text := flag.String("text", "This is a secret", "The plaintext to be encrypt/decrypt")
	mode := flag.Int("mode", 0, "Mode: 0 -> encrypt(default), 1 -> decrypt")
	flag.Parse()

	c, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
	}

	var result string
	before := time.Now()
	switch *mode {
	case 0:
		result = encrypt(c, *text)
	case 1:
		result = decrypt(c, *text)
	}
	after := time.Now()
	fmt.Println(result)
	err = ioutil.WriteFile("./result", []byte(result), 0644)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("cost", (after.UnixNano() - before.UnixNano()) / int64(time.Millisecond), "ms")
}

func decrypt(c cipher.Block, text string) string {
	bytes, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		fmt.Println(err)
	}
	pt := make([]byte, len(bytes))
	c.Decrypt(pt, bytes)

	return string(pt[:])
}

func encrypt(c cipher.Block, text string) string {
	bytes := []byte(text)
	// allocate space for ciphered data
	out := make([]byte, len(bytes))

	// encrypt
	c.Encrypt(out, bytes)
	// return hex string
	return base64.StdEncoding.EncodeToString(out)
}
