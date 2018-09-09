package main

import (
	"bytes"
	"compress/flate"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	h1 := ""
	h2 := ""
	h3 := ""
	h4 := ""

	allH := []string{h1, h2, h3, h4}

	var allB [][]byte

	for i, h := range allH {
		b, err := hex.DecodeString(h)
		if err != nil {
			fmt.Printf("For %d: Error: %v\n", i, b)
			continue
		}
		allB = append(allB, b)
		fmt.Printf("For %d: Length: %d\n", i, len(b))
	}

	fo, err := os.Create("output.rtf")
	if err != nil {
		log.Fatal(err)
	}
	defer fo.Close()

	r := flate.NewReader(bytes.NewReader(allB[2][1:]))
	defer r.Close()
	if _, err := io.Copy(fo, r); err != nil {
		fmt.Printf("Nope.\n")
	}
	fmt.Println()

}
