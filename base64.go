package main

import (
	"encoding/base64"
	"flag"
	"fmt"
)

func main() {
	text := flag.String("text", "This is a secret", "The plaintext to be encode")
	flag.Parse()
	fmt.Println(base64.StdEncoding.EncodeToString([]byte(*text)))
}
