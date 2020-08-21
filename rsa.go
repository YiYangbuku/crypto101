package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"time"
)
const key = "MIIEowIBAAKCAQEA3kiyNUbF8Q2FWYu5eeSESx5RYlQ13KnWuVyCnW/YakOAjqMaev8aD04jjYQHiAV+TVzlVa0Q9xFeNmik5YU2u1dB4jYer6V7Tnp7jxiMs3OfaekPZAlm3nKsvShkW75pyJ11a2Fkk0h7/9C2OTRAV67rpAmtGqqCqvyMaSn5SHtSDRDHPQJEJQnRWZ3gBbP93n2uPVPOXqVbjGIUF04BbUIY\nY/7QxjPJBa6cHnIWF0O67RWI8lCrkgvH9TRmxQ/Uul5kWajGvgNd7PTjMrkTMPk1XPMRrzdpQ/qzljARm4TfEejhPCiF1fJN7sB/O6B+J/VNLhigonzFm48MIqCKywIDAQABAoIBAQDD9fKlZS87w1+8AaISA5NdZez5cqPJVTEnWJcNlHyFsdyz4raFmD+xHoHZUHwTPkSKj0rerSQ3q+gSr04vybDml5ZNhsim\nsIx0MyaakFn6GgR/qicXWfKGSTJf0CwpdURsx+OsNsAj19L5Q1ZiN95BbsPuaVliGM/5TYu7anWJnwjjOdMaJBzBb5YkxDPP6emuWrM+/DtuHEOmGjByjGp6O97u0s1EUJIhD7dUIsG7X01LN1sHz4es2YOE2/TcMtzbVTJ138yQ/fi03aIikRRCdIGhOazoXuf/uxB7vXYl0pJRjocMemTblik/4Eur3TZv7UZn\nbQLWmJ1uDaIGCFeBAoGBAOBk62aZLFR8Kwv5kZ3/tAMQk7DH/UYFmCp46hFpBzKSRoj/t6F4wGNbtOGsDj6Z680vUYKIrfWvXYl7RYf49fZ6ZR3TxKPydWN68O4B6gQHzY0tD5TsA59Yst+lss3f63diQxpwR2VDBWsSkcx/TTh5RGgiS5JchIRZ52WcaOgnAoGBAP2Xr7xJ+vvCDc+03NzEkwsSzdZ+jn48KdNF\nq/J128VBU3/jqvNde4V9BS0Q77DKHhx8Ku+owXshRnPwHdSNZWbM+UkGsklJju2HrltrbKCe4y9JXEzCBAEUvYMExjTmvD/BFXfXkhLviR2/JO85xOfm0q7uyj7R74RdLw8Bq2q9AoGAQqkYWVFVnv+IFjWcsbA9vM9W4KR4tC2DR9LFzkhCMB7OP3KgDaL+nEWpFYV0cdpt93WAJGQQMbVX9zicrkkiYId5tIOt\nnfqwLs5d9oaxC2N6B9+ECnyWkubZtKDX3lsP61ZQkvEZ9JYVbPqGP5btzMIRtVHC24cKgvrmSARQEe8CgYBLWpF7r7gGq0kTWTjv2PGgrru/aufIlvQOtXs8lszxNlIrhGk9259isR2ioI+4xrZf3H2drVWg0uhudwFGMaXaADpq+rRrlpID/vdObwNeTxhO6suke1pZP+J4VijXe2CgyS0p6UKcodTqo+vDsjTS\nfAJKvOYX+KXlfRMIsIRROQKBgCaUNBbKcaWuyBV7WJWxYSnElcgZbCghnVkS3QX4fOZ0VVEJzAdTKjgzuSe8rmqXO54z57jFAxpQZmA4W8KVKtUP/pRbXO1OvifaSJYEndF6AWG4WAjxVOASySgI4NRTLz14EPGwhh4KmdybPH0ARI67fc3phAkUQjVFeEzKPh5l"

func main() {
	text := flag.String("text", "", "The plaintext to be encrypt/decrypt")
	mode := flag.Int("mode", 0, "Mode: 0 -> encrypt(default), 1 -> decrypt")
	flag.Parse()


	decodeString, _ := base64.StdEncoding.DecodeString(key)
	privateKey, err := x509.ParsePKCS1PrivateKey(decodeString)
	if err != nil {
		fmt.Println(err)
	}
	var result string
	before := time.Now()
	switch *mode {
	case 0:
		result = encryptRsa(privateKey.PublicKey, *text)
	case 1:
		result = decryptRsa(*privateKey, *text)
	}
	after := time.Now()
	fmt.Println(result)
	err = ioutil.WriteFile("./result", []byte(result), 0644)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("cost", after.Nanosecond() - before.Nanosecond(), "ns")
}

func decryptRsa(privateKey rsa.PrivateKey, text string) string {
	bytes, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		fmt.Println(err)
	}
	ciphertext, err := rsa.DecryptPKCS1v15(rand.Reader, &privateKey, bytes)
	return string(ciphertext)
}

func encryptRsa(publicKey rsa.PublicKey, text string) string {
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, &publicKey, []byte(text))
	if err != nil {
		fmt.Println(err)
	}
	return base64.StdEncoding.EncodeToString(ciphertext)
}
