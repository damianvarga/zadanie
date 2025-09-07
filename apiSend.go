package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)
// AI generovane
func Send(height int, width int, result string) {
	url := "https://zadanie.openmed.sk/challenge-me-easy"
	table_hash := EncodeBoard(height, width)
	
	// Dáta, ktoré chceš poslať
	payload := map[string]string{
		"uuid": "9556a3a9-85d0-4906-af1b-9d3391e67e42",
		"result":    result,
		"table_hash": table_hash,
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

	// Telo odpovede
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	// Výpis statusu odpovede
	fmt.Println("Status:", resp.Status)
	fmt.Println("Response:", string(body))

}
