package main

import (
	"crypto/sha256"
	"encoding/base64"
	"strconv"
)

//AI generovana funkcia
func EncodeBoard(height int, width int) string {
	// Input data
	data := createBoard(height, width)
	// Compute SHA-256 hash
	hash := sha256.Sum256([]byte(data))
	// Encode the hash in Base64
	base64Hash := base64.StdEncoding.EncodeToString(hash[:])

	return base64Hash
}

// Funkcia na vytvorenie stringu, ktory reprezentuje tabulku s presahmi
func createBoard(height int, width int) string {

	// prvy riadok 
	data := strconv.Itoa(height*width-1)
	data += ","
	for w := 0; w < width; w++ {
		data += strconv.Itoa(width*(height-1)+w)
	data += ","
	}
	data += strconv.Itoa(width*(height-1))
	data += "\n"

	// druhy az predposledny riadok
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
	
	// posledny riadok
	data += strconv.Itoa(width-1)
	data += ","
	for w := 0; w < width; w++ {
		data += strconv.Itoa(w)
		data += ","
	}
	data += strconv.Itoa(0)

	return data
}