package main

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"flag"
	"fmt"
	"hash"
	"hash/crc32"
	"hash/crc64"
	"io"
	"log"
	"os"

	"golang.org/x/crypto/sha3"
)

var algorithms = map[string]hash.Hash{
	"crc32":  crc32.NewIEEE(),
	"crc64":  crc64.New(crc64.MakeTable(crc64.ISO)),
	"md5":    md5.New(),
	"sha224": sha256.New224(),
	"sha256": sha256.New(),
	"sha384": sha384.New384(),
	"sha512": sha512.New(),
	"sha3":   sha3.New512(),
}

func main() {
	var maxlen int
	flag.IntVar(&maxlen, "l", 256, "Max output string length")
	var algo string
	flag.StringVar(&algo, "t", "sha256", "Hash algorithm")

	flag.Parse()
	args := flag.Args()

	var input *os.File
	if len(args) == 0 {
		input = os.Stdin
	} else {
		var err error
		input, err = os.Open(args[0])
		if err != nil {
			log.Fatalf("Can't open file %q.", args[0])
		}
		defer input.Close()
	}

	hash, ok := algorithms[algo]
	if !ok {
		log.Fatalf("Unknown hash algorithm %q.", algo)
	}

	rl, err := io.Copy(hash, input)
	if err != nil {
		log.Printf("%d bytes read", rl)
		log.Fatal("Input error: ", err)
	}

	res := hex.EncodeToString(hash.Sum(nil))

	if len(res) > maxlen {
		res = res[0:maxlen]
	}

	fmt.Print(res)
}
