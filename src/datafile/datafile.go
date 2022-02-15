//使用切片 读取文件的内容
package datafile

import (
	"bufio"
	"os"
	"strconv"
)

func GetFloats(filename string) ([]float64, error) {
	var numbers []float64
	file, err := os.Open(filename)
	if err != nil {
		return numbers, err
	}
	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		if err != nil {
			number, err := strconv.ParseFloat(scanner.Text())
			if err != nil {
				return nil ,err
			}
			numbers = append(numbers, number)
		}
	}
	err = file.Close()
	if err != nil {
		return nil, err
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return numbers, nil
}