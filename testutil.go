package main

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func string_to_array(s, sep string) ([]int, error) {
	splitted := strings.Split(s, sep)
	var res []int

	for _, v := range splitted {
		i, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		res = append(res, i)
	}

	return res, nil
}

func delete_last_if_empty(lines []string) []string {
	l := len(lines)
	if l > 0 && lines[l-1] == "" {
		return lines[:l-1]
	}
	return lines
}

func file_list(dir, mask string) ([]string, error) {
	if !fs.ValidPath(dir) {
		return nil, errors.New("not valid path")
	}
	root := os.DirFS(dir)
	files, err := fs.Glob(root, mask)
	if err != nil {
		return nil, err
	}

	return files, nil
}

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