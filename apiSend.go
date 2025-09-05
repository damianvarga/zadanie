package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)
// AI generovane
func main() {
	url := "https://zadanie.openmed.sk/challenge-me-easy"

	// Dáta, ktoré chceš poslať
	payload := map[string]string{
		"uuid": "9556a3a9-85d0-4906-af1b-9d3391e67d68",
		"result":    "22,23,18,28,24,34,35,30",
		"table_hash": "/5n2zQS74JggOKVd+UojTLaepaCHp6YcQP5Ccy4KYHM=",
	}

	// Prevod Go mapy/štruktúry na JSON
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

	// Vykonanie requestu
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Chyba pri odosielaní requestu: %v", err)
	}
	defer resp.Body.Close()

	// Výpis statusu odpovede
	fmt.Println("Status:", resp.Status)
}
