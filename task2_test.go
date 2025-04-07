package main

import (
	"fmt"
	"path/filepath"
	"strconv"
	"testing"
)

type args struct {
	rates rates
}

type test struct {
	name string
	args args
	want float64
}

const TestDataDir = "task2-input"

func get_tests() ([]test, error) {

	var err error

	answer_files, err := file_list(TestDataDir, "*.a")
	if err != nil || len(answer_files) == 0 {
		return nil, err
	}

	tests := make([]test, 0, len(answer_files)*3)

	for _, fn := range answer_files {
		full_fn := filepath.Join(TestDataDir, fn)
		test_inputs, err := read_file(fileNameWithoutExt(full_fn))
		if err != nil {
			return nil, err
		}

		n_test_sets, _ := strconv.Atoi(test_inputs[0])
		test_inputs = delete_last_if_empty(test_inputs[1:])

		answers, err := read_file(full_fn)
		if err != nil {
			continue
		}
		answers = delete_last_if_empty(answers)

		var test_name string
		var answer_float float64
		var rates rates
		var n_line int

		for n := range n_test_sets {
			for i := range 3 {
				for j := range 6 {
					c, err := string_to_array(test_inputs[n_line], " ")
					if err != nil {
						return nil, err
					}
					rates[i][j][0] = c[0]
					rates[i][j][1] = c[1]
					n_line++
				}
			}
			answer_float, _ = strconv.ParseFloat(answers[n], 64)
			test_name = fmt.Sprintf("%s#%d", fn, n+1)
			test := test{test_name, args{rates}, answer_float}
			tests = append(tests, test)
		}

	}
	return tests, nil
}

func Test_answer(t *testing.T) {
	tests, err := get_tests();
	if err != nil {
		t.Fatalf("can't get test data")
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := int(answer(tt.args.rates) * 1e6)
			want := int(tt.want * 1e6)
			if got != want {
				t.Errorf("answer() = %v, want %v", got, want)
			}
		})
	}
}
