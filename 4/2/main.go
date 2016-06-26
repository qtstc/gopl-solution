package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

const s256 = "SHA256"
const s384 = "SHA384"
const s512 = "SHA512"

var hash = flag.String("h", s256, "hash")

func main() {
	flag.Parse()
	fmt.Println(*hash)
	switch *hash {
	case s256:
		fmt.Println("256")
		fmt.Println(sha256.Sum256([]byte(flag.Args()[0])))
	case s384:
		fmt.Println("384")
		fmt.Println(sha512.Sum384([]byte(flag.Args()[0])))
	case s512:
		fmt.Println("512")
		fmt.Println(sha512.Sum512([]byte(flag.Args()[0])))
	}
}
