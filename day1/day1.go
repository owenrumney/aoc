package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

// getFuel is used for part 1 of day1 puzzle
func getFuel(mass int) (int) {
	return int(math.Floor(float64(mass/3)) -2)
}

// getFuelRecursive is used for part 2 of day 1 to get the fuel and the required fuel for the fuel
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

	// create a scanner to open the file
	scanner := bufio.NewScanner(file)
	totalFuel := 0

	// for each line in the input calculate the fuel and increase the totalFuel
	for scanner.Scan() {
		mass, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err.Error())
		}
		totalFuel += getFuelRecursive(mass)
	}

	// output the required fuel
	fmt.Printf("Total Fuel Required: %d", totalFuel)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
