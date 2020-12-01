package util

import (
	"bufio"
	"io/ioutil"
	"os"
)

const (
	PROTOCOL = "tcp"
	PORT     = ":9999"
)

// Loads and html file and returns the file string
// with the content
func LoadHtml(a string) string {
	html, _ := ioutil.ReadFile(a)

	return string(html)
}

// Scans from the user input and then returns the
// complete string
func ScanString() string {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	return scanner.Text()
}
