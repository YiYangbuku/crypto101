package main

import (
	"crypto/sha256"
	"encoding/base64"
	"flag"
	"fmt"
	"time"
)

func main() {
	text := flag.String("text", "", "The plaintext to be hash")
	before := time.Now()
	bytes := sha256.Sum256([]byte(*text))
	after := time.Now()
	fmt.Println(base64.StdEncoding.EncodeToString(bytes[:]))
	fmt.Println("cost", after.Nanosecond() - before.Nanosecond(), "ns")
}
