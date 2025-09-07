package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"encoding/json"
	"bytes"
)
// AI generovane - pytal som sa na zakladny GET request (na akukolvek URL adresu)
// ku kodu som pripojil cast kodu, generovaneho AI, zo suboru apiSend.go
func Call() []byte{
	url := "https://zadanie.openmed.sk/challenge-me-easy"
	payload := map[string]string{
		"uuid": "9556a3a9-85d0-4906-af1b-9d3391e67e36",
		"user": "Damián Varga",
	}

	// Prevod Go mapy/štruktúry na JSON
	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("Chyba pri marshall JSON: %v", err)
	}
	// Vytvorenie POST requestu
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Chyba pri vytváraní requestu: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Vykonanie requestu
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Chyba pri odosielaní requestu: %v", err)
	}
	defer resp.Body.Close()

	// Telo odpovede
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	fmt.Println("Status:", resp.Status)
	fmt.Println("Response:", string(body))
	return body
}