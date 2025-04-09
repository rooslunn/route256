package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t, n int

	fmt.Fscan(in, &t)

	for range t {
		fmt.Fscan(in, &n)
		words := make([]string, n)
		for i := range n {
			fmt.Fscan(in, &words[i])
		}
		fmt.Fprintf(out, "%d\n", task3_answer(words))
	}
}

func task3_answer(words []string) int {

	evenHits := make(map[string]int)
	oddHits := make(map[string]int)

	for _, word := range words {
		var evenLeadHits, oddLeadHits int
		var evenLead, oddLead rune

		evenLead = rune(word[0])
		if len(word) > 1 {
			oddLead = rune(word[1])
		}

		for i, r := range word {
			if i%2 == 0 {
				if evenLead != r {
					continue
				}
				evenLeadHits++
			} else {
				if oddLead != r {
					break
				}
				oddLeadHits++
			}
		}

		if evenLeadHits > 0 {
			evenHitKey := fmt.Sprintf("%s_%d", string(evenLead), evenLeadHits)
			evenHits[evenHitKey]++
		}

		if oddLeadHits > 0 {
			oddHitKey := fmt.Sprintf("%s_%d", string(oddLead), oddLeadHits)
			oddHits[oddHitKey]++
		}
	}

	count_even := 0
	count_odd := 0

	for k, v := range evenHits {
		if string(k[0]) == "0" {
			continue
		}
		if v < 2 {
			continue
		}
		count_even = max(count_even, v*(v-1)/2)
	}
	for k, v := range oddHits {
		if string(k[0]) == "0" {
			continue
		}
		if v < 2 {
			continue
		}
		count_odd = max(count_odd, v*(v-1)/2)
	}

	count := max(count_even, count_odd)

	return count
}
