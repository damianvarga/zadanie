
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func apiTest() {
	url := "https://zadanie.openmed.sk/data/uuid/9556a3a9-85d0-4906-af1b-9d3391e67d68"

	// Dáta, ktoré chceš poslať
	payload := map[string]string{
		"username": "john_doe",
		"email":    "john@example.com",
	}

	// Prevod na JSON
	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("Chyba pri marshall JSON: %v", err)
	}

	// Vytvorenie POST requestu
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Chyba pri vytváraní requestu: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// HTTP klient
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Chyba pri odosielaní requestu: %v", err)
	}
	defer resp.Body.Close()

	// Výpis statusu
	fmt.Println("Status:", resp.Status)

	// Čítanie odpovede ako string
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Chyba pri čítaní odpovede: %v", err)
	}
	fmt.Println("Raw response:", string(body))

	// Ak očakávaš JSON odpoveď, dekóduj ju
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		log.Fatalf("Chyba pri dekódovaní JSON: %v", err)
	}

	fmt.Println("Decoded JSON response:", result)
}
