package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type rates [3][6][2]int

func main() {

	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)

	var rates rates

	for range t {
		for i := range 3 {
			for j := range 6 {
				fmt.Fscan(in, &rates[i][j][0], &rates[i][j][1])
			}
		}
		fmt.Fprintf(out, "%f\n", answer(rates))
	}
}

func answer(rates rates) float64 {

	maxUSD := 0.0

	exchangePath1 := []int{0, 1, 2} // R-U
	exchangePath2 := [][]int{{0,1},{0,2},{1,0},{1,2},{2,0},{2,1}} // R-E E-U 
	exchangePath3 := [][]int{{0,1,2},{0,2,1},{1,0,2},{1,2,0},{2,1,0},{2,0,1}}

	for _, b := range exchangePath1 {
			rub_usd := float64(rates[b][0][1]) / float64(rates[b][0][0])
			maxUSD = math.Max(maxUSD, 1*rub_usd)
	}

	for _, b := range exchangePath2 {
			rub_eur := float64(rates[b[0]][1][1]) / float64(rates[b[0]][1][0])
			eur_usd := float64(rates[b[1]][5][1]) / float64(rates[b[1]][5][0])
			maxUSD = math.Max(maxUSD, 1*rub_eur*eur_usd)
	}

	for _, b := range exchangePath3 {
			rub_eur := float64(rates[b[0]][1][1]) / float64(rates[b[0]][1][0])
			eur_rub := float64(rates[b[1]][4][1]) / float64(rates[b[1]][4][0])
			rub_usd := float64(rates[b[2]][0][1]) / float64(rates[b[2]][0][0])
			maxUSD = math.Max(maxUSD, 1*rub_eur*eur_rub*rub_usd)

			rub_usd = float64(rates[b[0]][0][1]) / float64(rates[b[0]][0][0])
			usd_eur := float64(rates[b[1]][3][1]) / float64(rates[b[1]][3][0])
			eur_usd := float64(rates[b[2]][5][1]) / float64(rates[b[2]][5][0])
			maxUSD = math.Max(maxUSD, 1*rub_usd*usd_eur*eur_usd)

			rub_usd = float64(rates[b[0]][0][1]) / float64(rates[b[0]][0][0])
			usd_rub := float64(rates[b[1]][2][1]) / float64(rates[b[1]][2][0])
			rub_usd2 := float64(rates[b[2]][0][1]) / float64(rates[b[2]][0][0])
			maxUSD = math.Max(maxUSD, 1*rub_usd*usd_rub*rub_usd2)
	}

	return maxUSD
}
