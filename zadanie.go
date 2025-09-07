package main
import (
	"fmt"
	"strconv"
	"strings"
	"encoding/json"
	"log"
)


type PartialResponse struct {
    SetX string  `json:"set_x"`
    SetY string	 `json:"set_y"`
    SetZ string	 `json:"set_z"`
}

func main() {
	var choice int
	fmt.Println("If you want to input values manually, enter 1.\nIf you want to access the API, enter 2.")
	fmt.Scan(&choice)
	var h,w,i int
	if choice == 1 {

		fmt.Scan(&h)
		fmt.Scan(&w)
		fmt.Scan(&i)
		EncodeBoard(h,w)
		fmt.Println()

	} else if choice == 2 {
		body := Call()
		
		// generovane AI - chcel som, aby mi dal kod, kde z JSONu vo forme []byte ziskam hodnoty h,w,i
		// Tu som pouzival AI celkom dost; kopiroval som aj par riadkov kodu  
		// spracovanie udajov z API
		var resp PartialResponse
		if err := json.Unmarshal(body, &resp); err != nil {
			fmt.Println("Error:", err)
		}

		var err error
		h, err = strconv.Atoi(resp.SetY)
		w, err = strconv.Atoi(resp.SetX)
		i, err = strconv.Atoi(resp.SetZ)
		if err != nil {
			log.Fatal(err)
		}
		
	}
	
	resultString := intsToCSV(algorithm(h,w,i))
	fmt.Print("result ")
	fmt.Print(resultString)
	fmt.Println()

	// ak ideme cez API, odosli vysledok
	if choice == 2 {
		Send(h, w, resultString)
	}
	return 
}

// generovane AI
// konverzia pola cisel na string cisel oddelenych ciarkami
func intsToCSV(nums []int) string {
	strs := make([]string, len(nums))
	for i, n := range nums {
		strs[i] = strconv.Itoa(n)
	}
	return strings.Join(strs, ",")
}

// algoritmus na najdenie susedov
func algorithm(h,w,i int) []int {
	arr := make([]int, 0)
	// AI vygenerovala zakladny algoritmus - cykly + prvu podmienku, ni,nj,neighbor,arr
	for row := -1; row <= 1; row++ {
		for col := -1; col <= 1; col++ {
			if row == 0 && col == 0 {
				continue
			}
			// ni,nj,neighbor,arr boli vygenerovane AI, premenne upravene na zaklade toho, ako ich mam pomenovane
			ni := (row + i/w) % (h)
			nj := (col + i%w) % (w)
			neighbor := ni*w+nj

			// osetrenie okrajov
			// vrchny a spodny okraj
			if h != 1 {
				if i < w && row == -1 {
					neighbor += h*w
				}
			}
			// lavy a pravy okraj
			if w != 1 {
				if i % w == 0 && col == -1 {
					neighbor += w
				}
			}
			arr = append(arr, neighbor)

		}
	}
	return arr
}