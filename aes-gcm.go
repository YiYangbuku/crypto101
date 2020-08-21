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
	ad := flag.String("ad", "", "AdditionalData: the additional data used in gcm encryption/decryption")
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
		result = encryptGcm(gcm, *text, []byte(*ad))
	case 1:
		result = decryptGcm(gcm, *text, []byte(*ad))
	}
	after := time.Now()
	fmt.Println(result)
	err = ioutil.WriteFile("./result", []byte(result), 0644)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("cost", after.Nanosecond() - before.Nanosecond(), "ns")
}

func decryptGcm(gcm cipher.AEAD, text string, additionalData []byte) string {
	nonceSize := gcm.NonceSize()
	bytes, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		fmt.Println(err)
	}
	nonce, ciphertext := bytes[:nonceSize], bytes[nonceSize:]
	open, err := gcm.Open(nil, nonce, ciphertext, additionalData)
	if err != nil {
		fmt.Println(err)
	}
	return string(open)
}

func encryptGcm(gcm cipher.AEAD, text string, additionalData []byte) string {
	nonce := make([]byte, gcm.NonceSize())

	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err)
	}
	ciphertext := gcm.Seal(nonce, nonce, []byte(text), additionalData)
	return base64.StdEncoding.EncodeToString(ciphertext)
}
