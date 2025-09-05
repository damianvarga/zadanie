package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

//AI generovane
func main() {
	// Input data
	data := "35,30,31,32,33,34,35,30\n5,0,1,2,3,4,5,0\n11,6,7,8,9,10,11,6\n17,12,13,14,15,16,17,12\n23,18,19,20,21,22,23,18\n29,24,25,26,27,28,29,24\n35,30,31,32,33,34,35,30\n5,0,1,2,3,4,5,0"
	// data = "15,12,13,14,15,12\n3,0,1,2,3,0\n7,4,5,6,7,4\n11,8,9,10,11,8\n15,12,13,14,15,12\n3,0,1,2,3,0"
	// Compute SHA-256 hash
	hash := sha256.Sum256([]byte(data))

	// Encode the hash in Base64
	base64Hash := base64.StdEncoding.EncodeToString(hash[:])

	fmt.Println("Input:", data)
	fmt.Println("SHA-256 (Base64):", base64Hash)
	baseenc, _ := base64.StdEncoding.DecodeString(base64Hash)
	fmt.Printf("%x", baseenc)
}
