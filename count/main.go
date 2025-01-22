// count подчитывает количество вхождений каждой строки в файле
package main

import (
	"fmt"
	"log"
	"sluzov/datafile"
)

func main() {
	lines, err := datafile.GetStrings("C:/1Sluzov/Test1/votes.txt")
	if err != nil {
		log.Fatal(err)
	}

	counts := make(map[string]int)

	for _, line := range lines {
		counts[line]++
	}

	for name, count := range counts {
		fmt.Printf("Votes for %s: %d\n", name, count)

	}

}
