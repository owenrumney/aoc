package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func getFuel(mass int) (int) {
	return int(math.Floor(float64(mass/3)) -2)
}

func getFuelRecursive(mass int) int {
	if mass <= 0 {
		return 0
	}
	fuel := int(math.Max(float64(getFuel(mass)), 0))
	return fuel + getFuelRecursive(fuel)
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalFuel := 0
	for scanner.Scan() {
		mass, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err.Error())
		}
		totalFuel += getFuelRecursive(mass)
	}
	fmt.Printf("Total Fuel Required: %d", totalFuel)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
