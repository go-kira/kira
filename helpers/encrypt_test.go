package helpers

import (
	"encoding/hex"
	"fmt"
)

func ExampleEncrypt() {
	key := []byte("9d8b23c2529ced916abaf60599fb3110")
	text := []byte("Testing encrypt func")

	doTest, _ := Encrypt(text, key)

	fmt.Printf("%x\n", doTest)
}

func ExampleDecrypt() {
	key := []byte("9d8b23c2529ced916abaf60599fb3110")
	text, _ := hex.DecodeString("537235e0ba1c4551c1787ab68ceb4bc3c6e738f0e3b3e8656322932e2a56969b4bcf5afb53e07d082b5cd61e8a451433")

	doTest, _ := Decrypt(text, key)

	fmt.Printf("%s\n", doTest)
	// fmt.Println(err)

	// Output: Testing encrypt func
}
