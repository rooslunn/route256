package main

import (
	"fmt"
	"math"
)


type rates [3][6][2]int

func answer(rates rates) float64 {
	var bestRUR_USD, bestRUR_EUR, bestEUR_USD float64

	for i := range 3 {
		bestRUR_USD = math.Max(bestRUR_USD, float64(rates[i][0][1])/float64(rates[i][0][0]))
		bestRUR_EUR = math.Max(bestRUR_EUR, float64(rates[i][1][1])/float64(rates[i][1][0]))
		bestEUR_USD = math.Max(bestEUR_USD, float64(rates[i][5][1])/float64(rates[i][5][0]))
	}

	maxUSD_1 := 1 * bestRUR_USD
	maxUSD_2 := 1 * bestRUR_EUR * bestEUR_USD

	return math.Max(maxUSD_1, maxUSD_2)
}

func main() {
	var t int
	fmt.Scan(&t)

	var rates rates 

	for n := range t {
		for i := range 3 {
			for j := range 6 {
				fmt.Scan(&rates[i][j][0], &rates[i][j][1])
			}
		}
		fmt.Printf("solve: for set %d you got max dollars %f\n", n+1, answer(rates))
	}
}

// func main() {
	// var t int
	// fmt.Scan(&t)

	// var rates rates 

	// for n := range t {
	// 	for i := range 3 {
	// 		for j := range 6 {
	// 			fmt.Scan(&rates[i][j][0], &rates[i][j][1])
	// 		}
	// 	}
	// 	fmt.Printf("solve: for set %d you got max dollars %f\n", n+1, solve(rates))
	// }
// }

// func solve(rates rates) float64 {

// 	maxDollars := 0.0

// 	// Вариант 1: Посетить банк A первым
// 	// Рубли -> Доллары в A
// 	dollarsA := 1.0 * float64(rates[0][0][1]) / float64(rates[0][0][0])
// 	// Доллары A -> Рубли в B -> Доллары в C
// 	rublesB := dollarsA * float64(rates[1][2][1]) / float64(rates[1][2][0])
// 	dollarsC1 := rublesB * float64(rates[2][0][1]) / float64(rates[2][0][0])
// 	maxDollars = math.Max(maxDollars, dollarsC1)
// 	// Доллары A -> Евро в B -> Доллары в C
// 	eurosB := dollarsA * float64(rates[1][3][1]) / float64(rates[1][3][0])
// 	dollarsC2 := eurosB * float64(rates[2][5][1]) / float64(rates[2][5][0])
// 	maxDollars = math.Max(maxDollars, dollarsC2)

// 	// Рубли -> Евро в A
// 	eurosA := 1.0 * float64(rates[0][1][1]) / float64(rates[0][1][0])
// 	// Евро A -> Рубли в B -> Доллары в C
// 	rublesB1 := eurosA * float64(rates[1][4][1]) / float64(rates[1][4][0])
// 	dollarsC3 := rublesB1 * float64(rates[2][0][1]) / float64(rates[2][0][0])
// 	maxDollars = math.Max(maxDollars, dollarsC3)
// 	// Евро A -> Доллары в B -> Доллары в C
// 	dollarsB1 := eurosA * float64(rates[1][5][1]) / float64(rates[1][5][0])
// 	maxDollars = math.Max(maxDollars, dollarsB1)

// 	// Вариант 2: Посетить банк B первым
// 	// Рубли -> Доллары в B
// 	dollarsB := 1.0 * float64(rates[1][0][1]) / float64(rates[1][0][0])
// 	// Доллары B -> Рубли в A -> Доллары в C
// 	rublesA := dollarsB * float64(rates[0][2][1]) / float64(rates[0][2][0])
// 	dollarsC4 := rublesA * float64(rates[2][0][1]) / float64(rates[2][0][0])
// 	maxDollars = math.Max(maxDollars, dollarsC4)
// 	// Доллары B -> Евро в A -> Доллары в C
// 	eurosA1 := dollarsB * float64(rates[0][3][1]) / float64(rates[0][3][0])
// 	dollarsC5 := eurosA1 * float64(rates[2][5][1]) / float64(rates[2][5][0])
// 	maxDollars = math.Max(maxDollars, dollarsC5)

// 	// Рубли -> Евро в B
// 	eurosB1 := 1.0 * float64(rates[1][1][1]) / float64(rates[1][1][0])
// 	// Евро B -> Рубли в A -> Доллары в C
// 	rublesA1 := eurosB1 * float64(rates[0][4][1]) / float64(rates[0][4][0])
// 	dollarsC6 := rublesA1 * float64(rates[2][0][1]) / float64(rates[2][0][0])
// 	maxDollars = math.Max(maxDollars, dollarsC6)
// 	// Евро B -> Доллары в A -> Доллары в C
// 	dollarsA1 := eurosB1 * float64(rates[0][5][1]) / float64(rates[0][5][0])
// 	maxDollars = math.Max(maxDollars, dollarsA1)

// 	// Вариант 3: Посетить банк C первым
// 	// Рубли -> Доллары в C
// 	dollarsC := 1.0 * float64(rates[2][0][1]) / float64(rates[2][0][0])
// 	// Доллары C -> Рубли в A -> Доллары в B
// 	rublesA2 := dollarsC * float64(rates[0][2][1]) / float64(rates[0][2][0])
// 	dollarsB2 := rublesA2 * float64(rates[1][0][1]) / float64(rates[1][0][0])
// 	maxDollars = math.Max(maxDollars, dollarsB2)
// 	// Доллары C -> Евро в A -> Доллары в B
// 	eurosA2 := dollarsC * float64(rates[0][3][1]) / float64(rates[0][3][0])
// 	dollarsB3 := eurosA2 * float64(rates[1][5][1]) / float64(rates[1][5][0])
// 	maxDollars = math.Max(maxDollars, dollarsB3)

// 	// Рубли -> Евро в C
// 	eurosC := 1.0 * float64(rates[2][1][1]) / float64(rates[2][1][0])
// 	// Евро C -> Рубли в A -> Доллары в B
// 	rublesA3 := eurosC * float64(rates[0][4][1]) / float64(rates[0][4][0])
// 	dollarsB4 := rublesA3 * float64(rates[1][0][1]) / float64(rates[1][0][0])
// 	maxDollars = math.Max(maxDollars, dollarsB4)
// 	// Евро C -> Доллары в A -> Доллары в B
// 	dollarsA2 := eurosC * float64(rates[0][5][1]) / float64(rates[0][5][0])
// 	maxDollars = math.Max(maxDollars, dollarsA2)

// 	return maxDollars
// }
