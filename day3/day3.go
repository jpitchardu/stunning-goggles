package main

import (
	"fmt"
	"strings"
)

func main() {

	line := "vJrwpWtwJgWrhcsFMMfFFhFp"

	firstCompartment := line[len(line)/2:]
	secondCompartment := line[:len(line)/2]

	var duplicate byte

	fmt.Println(duplicate)

	for i := 0; i < len(firstCompartment); i++ {
		idx := strings.IndexByte(secondCompartment, firstCompartment[i])
		if idx > -1 {
			duplicate = firstCompartment[i]
		}
	}

	fmt.Println(firstCompartment, secondCompartment)

	fmt.Println(duplicate)

}
