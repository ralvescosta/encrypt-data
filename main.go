package main

import (
	"encrypt-data/pkg/crypto"
	"encrypt-data/pkg/io"
	"log"
)

func main() {
	data, err := io.ReadInput()
	if err != nil {
		log.Fatal(err)
	}

	encrypted, err := crypto.Encrypt(data.PublicKey, data.Payload)
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(io.WriteOutput(encrypted))
}
