package main

import (
	"fmt"
	"path/filepath"
	"strconv"
	"testing"
)

const TaskCountMulti = 7

type task3_args struct {
	strings []string
}

type task3_test struct {
	name string
	args task3_args
	want int
}

const Task3TestDataDir = "task3-input"
const Task3CopilotDataDir = "task3-input/copilot"

func get_tests_task3() ([]task3_test, error) {

	var err error

	// testDir := Task3CopilotDataDir
	testDir := Task3TestDataDir

	answer_files, err := file_list(testDir, "*.a")
	if err != nil || len(answer_files) == 0 {
		return nil, err
	}

	tests := make([]task3_test, 0, len(answer_files)*TaskCountMulti)

	for _, fn := range answer_files {
		full_fn := filepath.Join(testDir, fn)
		test_inputs, err := read_file(fileNameWithoutExt(full_fn))
		if err != nil {
			return nil, err
		}

		n_test_sets, _ := strconv.Atoi(test_inputs[0])
		test_inputs = delete_last_if_empty(test_inputs[1:])
		test_start_line := 0

		answers, err := read_file(full_fn)
		if err != nil {
			continue
		}
		answers = delete_last_if_empty(answers)

		for l := range n_test_sets {
			n_strings, _ := strconv.Atoi(test_inputs[test_start_line])
			s, e := test_start_line+1, test_start_line+n_strings+1
			strings := test_inputs[s:e]

			answer_int, _ := strconv.Atoi(answers[l])
			test_name := fmt.Sprintf("%s#%d", fn, l+1)

			test := task3_test{test_name, task3_args{strings}, answer_int}
			tests = append(tests, test)

			test_start_line += n_strings + 1
		}

	}
	return tests, nil
}

func Test_task3_answer(t *testing.T) {
	tests, err := get_tests_task3()
	if err != nil {
		t.Fatalf("cant't load tests")
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := task3_answer(tt.args.strings); got != tt.want {
				t.Errorf("task3_answer() = %v, want %v", got, tt.want)
			}
		})
	}
}
