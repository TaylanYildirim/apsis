package fileUtil

import (
	"apsis/util/stringUtil"
	"bufio"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFile() *map[string]string {
	resultMap := make(map[string]string)
	_, b, _, _ := runtime.Caller(0)
	root := filepath.Join(filepath.Dir(b), "../../")
	absPath, _ := filepath.Abs(root + "/output")
	fullPath := absPath + "/output.txt"
	ensureBaseDir(fullPath)
	lines, err := readLines(fullPath)
	check(err)
	for _, line := range lines {
		arr := stringUtil.Split(string(line), ":")
		if len(arr) == 2 {
			resultMap[arr[0]] = arr[1]
		}
	}
	return &resultMap
}
func readLines(path string) ([]string, error) {
	file, err := os.OpenFile(path, os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func WriteFiles(str string) {
	_, b, _, _ := runtime.Caller(0)
	root := filepath.Join(filepath.Dir(b), "../../")
	absPath, _ := filepath.Abs(root + "/output")
	fullPath := absPath + "/output.txt"

	ensureBaseDir(fullPath)
	f, err := os.OpenFile(fullPath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	check(err)
	defer f.Close()
	_, err = f.WriteString(str + "\n")
	check(err)
}

func ensureBaseDir(filePath string) error {
	baseDir := path.Dir(filePath)
	info, err := os.Stat(filePath)
	if err == nil && info.IsDir() {
		return nil
	}
	return os.MkdirAll(baseDir, 0755)
}
