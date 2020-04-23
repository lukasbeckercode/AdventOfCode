package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("in.txt")
	var mod, fuel []int64
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, err := strconv.ParseInt(scanner.Text(), 10, 32)
		if err != nil {
			log.Fatal(err)
		}
		mod = append(mod, i)
	}
	var x int64 = 0
	for _, v := range mod {
		n := calcFuel(v)
		x += n
		for {
			n = calcFuel(n)
			if n > 0 {
				x += n
			} else {
				break
			}

		}

		fuel = append(fuel, x)
		x = 0
	}

	fmt.Println(fuel)
	var ans int64 = 0
	for _, v := range fuel {
		ans += v
	}
	fmt.Println(ans)
}

func calcFuel(i int64) int64 {
	var n = float64(i / 3)
	n = math.Ceil(n)
	n -= 2

	return int64(n)
}
