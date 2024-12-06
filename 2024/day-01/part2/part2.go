package part2

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func similarity(key int, elems []int) int {
	if !slices.Contains(elems, key) {
		return 0
	}

	idx := slices.Index(elems, key)
	return 1 + similarity(key, elems[idx+1:])
}

func Test(filepath string) {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Printf("Error reading file: %s", err)
	}
	defer file.Close()

	lElems := make([]int, 0, 1000)
	rElems := make([]int, 0, 1000)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "   ")
		l, _ := strconv.Atoi(line[0])
		r, _ := strconv.Atoi(line[1])
		lElems = append(lElems, l)
		rElems = append(rElems, r)
	}

	slices.Sort(lElems)
	slices.Sort(rElems)

	total := 0

	for i := 0; i < 1000; i++ {
		total += lElems[i] * similarity(lElems[i], rElems)
	}

	fmt.Printf("Part 2: %d\n", total)
}
