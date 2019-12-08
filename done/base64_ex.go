package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	const foobar = `foo bar`
	encoding := base64.StdEncoding
	encodedFooBar := make([]byte, encoding.EncodedLen(len(foobar)))
	encoding.Encode(encodedFooBar, []byte(foobar))
	fmt.Printf("%s\n", encodedFooBar)

	fmt.Println(base64.StdEncoding.EncodeToString([]byte(`foo bar`)))
	decoded, err := base64.StdEncoding.DecodeString(`biws`)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", decoded)
}
