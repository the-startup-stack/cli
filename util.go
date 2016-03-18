package stackcli

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func unzip(archive, target string) error {
	reader, err := zip.OpenReader(archive)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(target, 0755); err != nil {
		return err
	}

	for _, file := range reader.File {
		path := filepath.Join(target, file.Name)
		if file.FileInfo().IsDir() {
			os.MkdirAll(path, file.Mode())
			continue
		}

		fileReader, err := file.Open()
		if err != nil {
			return err
		}
		defer fileReader.Close()

		targetFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}
		defer targetFile.Close()

		if _, err := io.Copy(targetFile, fileReader); err != nil {
			return err
		}
	}
	return nil
}

func reverseArray(list []string) []string {
	for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
		list[i], list[j] = list[j], list[i]
	}
	return list
}

type ByLength []string

func (a ByLength) Len() int { return len(a) }
func (a ByLength) Less(i, j int) bool {
	return len(strings.Split(a[i], "/")) < len(strings.Split(a[j], "/"))
}
func (a ByLength) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func traverseDirStructure(list []string) []string {
	sort.Sort(ByLength(list))
	return reverseArray(list)
}
