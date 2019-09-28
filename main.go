package main

import (
	"encoding/hex"
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
)

// Repeatedly xoring the bytestream of key with plaintext.
//
// Returns the ciphertext as a string
func DecryptRepeatingKeyXOR(plaintext, key string) string {

	// Convert Your input text into a sequence of bytes
	byteStream, _ := hex.DecodeString(plaintext)

	// Creates space for cipher.
	cipherStream := make([]byte, len(byteStream))

	// Initialize the key as bytes and coressponding params.
	byteKey := []byte(key)
	keySize := len(byteKey)

	// This is like auto in C++.
	for i := range byteStream {
		cipherStream[i] = byteStream[i] ^ byteKey[i%keySize]
	}

	return string(cipherStream)
}

// Repeatedly xoring the bytestream of key with plaintext.
//
// Returns the ciphertext as a string
func RepeatingKeyXOR(plaintext, key string) string {

	// Convert Your input text into a sequence of bytes
	byteStream := []byte(plaintext)

	// Creates space for cipher.
	cipherStream := make([]byte, len(byteStream))

	// Initialize the key as bytes and coressponding params.
	byteKey := []byte(key)
	keySize := len(byteKey)

	// This is like auto in C++.
	for i := range byteStream {
		cipherStream[i] = byteStream[i] ^ byteKey[i%keySize]
	}

	return hex.EncodeToString(cipherStream)
}

/************* Main entrypoint - CLI Configurations **************/
/*****************************************************************/
/*****************************************************************/
/*****************************************************************/
/*****************************************************************/
/*****************************************************************/
/*****************************************************************/
/*****************************************************************/
/*****************************************************************/
func main() {
	app := cli.NewApp()
	app.Name = "repeatxor"
	app.Usage = " Gives you repeating key xor."
	app.Action = func(c *cli.Context) error {
		if c.Args().Get(0) == "decrypt" {
			enc := DecryptRepeatingKeyXOR(c.Args().Get(1), c.Args().Get(2))
			fmt.Printf("Your decypted string is on the line below: \n%v\n", enc)
		} else {
			enc := RepeatingKeyXOR(c.Args().Get(0), c.Args().Get(1))
			fmt.Printf("Your encrypted string is on the line below: \n%v\n", enc)
		}
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

/*****************************************************************/
/*****************************************************************/
/*****************************************************************/
/*****************************************************************/
/*****************************************************************/
