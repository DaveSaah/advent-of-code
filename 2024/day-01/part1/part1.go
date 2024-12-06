package part1

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

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

	total := 0.0

	for i := 0; i < 1000; i++ {
		total += math.Abs(float64(lElems[i]) - float64(rElems[i]))
	}

	fmt.Printf("Part 1: %.0f\n", total)
}
