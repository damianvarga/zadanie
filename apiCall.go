package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)
// AI generovane - pytal som sa na zakladny GET request (na akukolvek URL adresu)
func main() {
	url := "https://zadanie.openmed.sk/"

	// GET request
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error making GET request: %v", err)
	}
	defer resp.Body.Close()

	// Telo odpovede
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	fmt.Println("Status:", resp.Status)
	fmt.Println("Response:", string(body))
}