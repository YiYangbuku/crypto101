package main

import (
	"crypto/sha256"
	"encoding/base64"
	"flag"
	"fmt"
)

func main() {
	text := flag.String("text", "", "The plaintext to be hash")
	bytes := sha256.Sum256([]byte(*text))
	fmt.Println(base64.StdEncoding.EncodeToString(bytes[:]))
}
