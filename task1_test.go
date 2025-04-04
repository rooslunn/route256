/**
 * Author: mysterious_kangaroo_14271
 */

package main

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func read_file(from_file string) ([]string, error) {
	data, err := os.ReadFile(from_file)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(data), "\n"), nil
}

func fileNameWithoutExt(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}

func Test_testCondition(t *testing.T) {

	type test struct {
		name string
		arg string
		want string
	}

	var err error
	const TestDataDir = "task1-test-input"

	root := os.DirFS(TestDataDir)
	answer_files, err := fs.Glob(root, "*.a")
	if err != nil {
		t.Fatal("can't list dir with test files")
	}

	tests := make([]test, 0, len(answer_files))
	
	for _, fn := range answer_files {
		full_fn := filepath.Join(TestDataDir, fn)
		test_inputs, err := read_file(fileNameWithoutExt(full_fn)) 
		if err != nil {
			t.Fatalf("can't read test inputs from %s\n", fn)
			continue
		}
		test_inputs = test_inputs[1:]

		answers, err := read_file(full_fn)
		if err != nil {
			t.Fatalf("can't read test answers from %s\n", fn)
			continue
		}

		for i, input := range test_inputs {
			if input == "" {
				continue
			}
			test := test{fn, input, answers[i]}
			tests = append(tests, test)
		}
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := testCondition(tt.arg); got != tt.want {
				t.Errorf("testCondition() = %v, want %v", got, tt.want)
			}
		})
	}
}

// todo: benchmark