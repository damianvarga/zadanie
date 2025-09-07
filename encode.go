package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strconv"
)

//AI generovana funkcia
func EncodeBoard(height int, width int) string {
	// data := "35,30,31,32,33,34,35,30\n5,0,1,2,3,4,5,0\n11,6,7,8,9,10,11,6\n17,12,13,14,15,16,17,12\n23,18,19,20,21,22,23,18\n29,24,25,26,27,28,29,24\n35,30,31,32,33,34,35,30\n5,0,1,2,3,4,5,0"
	// Input data
	data := createBoard(height, width)
	data =  "53,45,46,47,48,49,50,51,52,53,45\n8,0,1,2,3,4,5,6,7,8,0\n17,9,10,11,12,13,14,15,16,17,9\n26,18,19,20,21,22,23,24,25,26,18\n35,27,28,29,30,31,32,33,34,35,27\n44,36,37,38,39,40,41,42,43,44,36\n53,45,46,47,48,49,50,51,52,53,45\n8,0,1,2,3,4,5,6,7,8,0"
	// Compute SHA-256 hash
	hash := sha256.Sum256([]byte(data))
	// Encode the hash in Base64
	base64Hash := base64.StdEncoding.EncodeToString(hash[:])

	fmt.Println("Input:", data)
	fmt.Printf("%x\n", hash)
	fmt.Println("SHA-256 (Base64):", base64Hash)
	// baseenc, _ := base64.StdEncoding.DecodeString(base64Hash)
	return base64Hash
}

	
func createBoard(height int, width int) string {

	data := strconv.Itoa(height*width-1)
	data += ","
	for w := 0; w < width; w++ {
		data += strconv.Itoa(width*(height-1)+w)
	data += ","
	}
	data += strconv.Itoa(width*(height-1))
	data += "\n"
	for h := 0; h < height; h++ {
		data += strconv.Itoa(h*width+width-1)
		data += ","
		for w := 0; w < width; w++ {
			data += strconv.Itoa(width*(h)+w)
			data += ","
		}
		data += strconv.Itoa(h*width)
		data += "\n"
	}
	data += strconv.Itoa(width-1)
	data += ","
	for w := 0; w < width; w++ {
		data += strconv.Itoa(w)
		data += ","
	}
	data += strconv.Itoa(0)
	fmt.Println(data)

	return data
}