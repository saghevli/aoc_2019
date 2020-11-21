package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func ReadInts(r io.Reader) (result []int, e error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return result, err
		}
		result = append(result, x)
	}
	return result, scanner.Err()
}

func SumOfMasses(masses []int) (total int) {
	for _, m := range masses {
		total += MassSumRecursive(m/3 - 2)
	}
	return total
}

func MassSumRecursive(mass int) (total int) {
	if mass <= 0 {
		return 0
	} else {
		return mass + MassSumRecursive(mass/3-2)
	}
}

func main() {
	r, _ := os.Open(os.Args[1])
	masses, _ := ReadInts(r)
	fmt.Println(SumOfMasses(masses))
}
