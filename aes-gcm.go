package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"time"
)

func main() {
	key := []byte("somekeysomekeysomekeysomekeyabcd")
	text := flag.String("text", "", "The plaintext to be encrypt/decrypt")
	mode := flag.Int("mode", 0, "Mode: 0 -> encrypt(default), 1 -> decrypt")
	flag.Parse()

	c, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		fmt.Println(err)
	}
	var result string
	before := time.Now()
	switch *mode {
	case 0:
		result = encryptGcm(gcm, *text)
	case 1:
		result = decryptGcm(gcm, *text)
	}
	after := time.Now()
	fmt.Println(result)
	err = ioutil.WriteFile("./result", []byte(result), 0644)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("cost", (after.UnixNano() - before.UnixNano()) / int64(time.Millisecond), "ms")
}

func decryptGcm(gcm cipher.AEAD, text string) string {
	nonceSize := gcm.NonceSize()
	bytes, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		fmt.Println(err)
	}
	nonce, ciphertext := bytes[:nonceSize], bytes[nonceSize:]
	open, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		fmt.Println(err)
	}
	return string(open)
}

func encryptGcm(gcm cipher.AEAD, text string) string {
	nonce := make([]byte, gcm.NonceSize())

	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err)
	}
	ciphertext := gcm.Seal(nonce, nonce, []byte(text), nil)
	return base64.StdEncoding.EncodeToString(ciphertext)
}
