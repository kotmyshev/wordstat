package filestat

import (
	"bufio"
	"os"
	"strings"
)

// Split String Line "s" to Array thru multiple separators "seps".
func splitLineToWords(s string, seps string) []string {
	splitter := func(r rune) bool {
		return strings.ContainsRune(seps, r)
	}
	return strings.FieldsFunc(s, splitter)
}

// Calulate Statistics Words in File "filename" with multiple separators.
// File read with bufio.NewScanner -> line by line
func WordStatFromFile(filename string, separators string) (map[string]int, error) {

	stat := make(map[string]int)

	file, err := os.Open(filename)
	if err != nil {
		return stat, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		words := splitLineToWords(scanner.Text(), separators)

		for _, w := range words {
			w = strings.ToUpper(w)
			stat[w]++
		}
	}
	err = file.Close()
	if err != nil {
		return stat, err
	}
	if scanner.Err() != nil {
		return stat, scanner.Err()
	}

	return stat, nil
}
