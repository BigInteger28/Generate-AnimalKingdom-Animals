package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	TotalPieces = 280 // Aantal dieren dat je wilt genereren
	MinStep     = -23 // Laagste stap
	MaxStep     = 30  // Hoogste stap
)

var ForbiddenSteps = []int{-30, -27, -24, -21, -18, -15, -12, -9, -6, -3, 0, 3, 6, 9, 12, 15, 18, 21, 24, 27, 30} // Verboden stappen
var existingAnimals = [][2]int{
	{5, 8}, {32, -2}, {-11, 8}, {-16, -4}, {-4, 2}, {-13, -10}, {7, -8}, {-7, 14}, {22, 2}, {-14, -11}, {13, 7}, {4, 2}, {-4, 7}, {14, 11}, {26, 22}, {-23, -26}, {-4, 5}, {-2, -4}, {-4, 8}, {-7, 5},
	{1, 4}, {19, -5}, {4, -5}, {-1, 13}, {4, 11}, {5, 7}, {7, 5}, {8, -11}, {-17, 11}, {-17, -20}, {-11, -13}, {8, -4}, {-13, 5}, {-5, 11}, {11, -2}, {-17, 19}, {7, -8}, {-8, -2}, {5, -7}, {13, 10},
	{20, 4}, {-16, -1}, {-14, -5}, {2, -10}, {-4, 8}, {-23, -16}, {7, -5}, {-2, -10}, {7, -11}, {-14, 10}, {-1, -7}, {-4, 19}, {-10, -11}, {22, 16}, {4, 14}, {2, -19}, {7, 20}, {1, -2}, {14, 5}, {7, -14},
	{-5, -17}, {2, -4}, {-11, 1}, {-2, -23}, {-2, -4}, {-5, -8}, {8, 5}, {10, 8}, {-5, 2}, {16, -13}, {-16, -13}, {-8, 16}, {-11, 7}, {19, -23}, {-4, -7}, {2, -11}, {-11, 5}, {4, -7}, {16, 14}, {-7, -4},
	{11, -7}, {-4, 13}, {-5, 11}, {8, -1}, {-13, 5}, {7, 16}, {-16, -19}, {7, 2}, {-14, 8}, {-16, -23}, {-8, -11}, {7, -20}, {23, -13}, {-7, 2}, {16, -7}, {16, 5}, {17, 13}, {17, 4}, {5, 20}, {10, -7},
	{13, -4}, {-16, 17}, {11, 8}, {13, -19}, {-2, 7}, {19, 7}, {-8, -7}, {10, -4}, {7, 17}, {-16, 1}, {4, 7}, {-7, 4}, {-7, -4}, {1, 10}, {1, -7}, {-7, -10}, {-7, 4}, {-8, 4}, {8, -11}, {8, -11},
}

type Animal struct {
	XStep, YStep int
}

func generateUniqueStep() Animal {
	var animal Animal
	for {
		x := rand.Intn(MaxStep-MinStep) + MinStep
		y := rand.Intn(MaxStep-MinStep) + MinStep

		// Controleer of de stap uniek is, niet verboden en niet hetzelfde als een bestaand dier
		if !isStepForbidden(x, y) && !isExistingAnimal(x, y) {
			animal.XStep = x
			animal.YStep = y
			return animal
		}
	}
}

func isStepForbidden(x, y int) bool {
	for _, step := range ForbiddenSteps {
		if x == step || y == step {
			return true
		}
	}
	return false
}

func isExistingAnimal(x, y int) bool {
	for _, existingAnimal := range existingAnimals {
		if x == existingAnimal[0] && y == existingAnimal[1] {
			return true
		}
	}
	return false
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var animals [TotalPieces]Animal
	for i := 0; i < TotalPieces; i++ {
		animals[i] = generateUniqueStep()
	}
	fmt.Println("Generated animals:")
	for _, animal := range animals {
		fmt.Printf("{%d, %d}, ", animal.XStep, animal.YStep)
	}
	fmt.Println("\nDruk op een toets om af te sluiten ...")
	fmt.Scanln()
}
