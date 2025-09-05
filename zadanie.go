package main
import (
	"fmt"
	"strconv"
	"strings"
)


func main() {
	var h,w,i int
	fmt.Scan(&h)
	fmt.Scan(&w)
	fmt.Scan(&i)
	
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
	arrStr := intsToCSV(arr)
	fmt.Print(arrStr)

	return 
}

// generovane AI
func intsToCSV(nums []int) string {
	strs := make([]string, len(nums))
	for i, n := range nums {
		strs[i] = strconv.Itoa(n)
	}
	return strings.Join(strs, ",")
}

