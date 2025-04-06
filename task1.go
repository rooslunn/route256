/**
 * Author: mysterious_kangaroo_14271
 */

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main_task1() {

	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n_test int
	var s_test string
	fmt.Fscan(in, &n_test)

	for range n_test {
		fmt.Fscan(in, &s_test)
		fmt.Fprint(out, testCondition(s_test) + "\n")
	}

}

func testCondition(s string) string {

	ace := s[0]
	l := len(s)
	i := 1
	res := "YES"

	for i < l {

		if s[i] == ace {
			i++
			continue
		}

		if (i+1 < l) && (s[i+1] == ace) {
			i = i + 2
			continue
		}

		res = "NO"
		break
	}

	return res
}
