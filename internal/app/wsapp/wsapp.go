// Package wsapp - WordStat Application source file.

package wsapp

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	//"sort"
	"strings"
)

// Read One Line From Stdin with desired URL.
func ReadLineFromStdIn() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return " ", nil
	}

	input = strings.TrimSpace(input)

	return input, nil
}

// Print Word Statistics map to stdout.
func PrintStatInStdOut(m map[string]int) {

	for word, count := range m {
		fmt.Println(word, " - ", count)
	}

}

// Split String Line "s" to Array thru multiple separators "seps".
func splitHtmlToWords(s string) []string {
	seps := "<>"
	splitter := func(r rune) bool {
		return strings.ContainsRune(seps, r)
	}
	return strings.FieldsFunc(s, splitter)
}

// Parse HTML file to TXT file by tags /a /span
func ParseHtmlFileToTextFile(htmlfile string, textfile string) error {

	htmldata, err := ioutil.ReadFile(htmlfile)
	if err != nil {
		return err
	}

	htmlstr := string(htmldata)

	warray := splitHtmlToWords(htmlstr)

	var wslice []string

	for indx, w := range warray {

		if w == "/a" || w == "/span" || w == "/title" {
			if indx == 0 {
				continue
			}

			wslice = append(wslice, warray[indx-1]+" ")
		}
	}

	file, err := os.Create(textfile)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, line := range wslice {
		_, err = file.WriteString(line)

		if err != nil {
			return err
		}

	}

	err = file.Close()
	if err != nil {
		return err
	}
	return nil
}

// Convert HTML file to TXT file -> delete all tags
func ConvertHtmlFileToTextFile(htmlfile string, textfile string) error {

	content, err := ioutil.ReadFile(htmlfile)

	if err != nil {
		return err
	}

	htmlstr := string(content)

	var f bool
	f = false

	var wslice []string

	word := ""

	for _, r := range htmlstr {

		if r == '>' {
			f = true
			continue
		}

		if f == true {
			if r == '<' {

				if word != "" {
					word = strings.Trim(word, "\n\t\r")
					wslice = append(wslice, word)
					word = ""
				}

				f = false
				continue
			}

			word = word + string(r)
		}
	}

	file, err := os.Create(textfile)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, line := range wslice {
		_, err = file.WriteString(line)

		if err != nil {
			return err
		}

	}

	err = file.Close()
	if err != nil {
		return err
	}
	return nil
}

// Make Request to URL "ureq" and Create two files: fname.html and fname.txt.
// fname.html is copy of target URL, fname.txt - texts from target URL only.
func MakeFilesFromURL(ureq string, fname string) error {
	resp, err := http.Get(ureq)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	hfname := fname + ".html"
	tfname := fname + ".txt"

	file, err := os.Create(hfname)

	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(body)

	if err != nil {
		return err
	}

	err = file.Close()

	if err != nil {
		return err
	}

	err = ParseHtmlFileToTextFile(hfname, tfname)

	if err != nil {
		return err
	}

	return nil
}
