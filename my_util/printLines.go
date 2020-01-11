// Test
package my_util

import (
	"bufio"
	"fmt"
	"os"
)

func PrintLines(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(MyUpper(scanner.Text()))
	}
	if scanner.Err() != nil {
		return scanner.Err()
	}
	return nil
}
