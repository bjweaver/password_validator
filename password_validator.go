package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const MIN_LENGTH = 8
const MAX_LENGTH = 64
const WEAK_FILE = "weak_password_list.txt"

var common_pass []string

func main() {
	common_pass = load_weak_file() // Load weak passwords file
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		word := scanner.Text()
		pass, message := test_password(word)
		if !pass {
			fmt.Println(message)
		}
	}

}

func test_password(word string) (bool, string) {

	if !check_length(word) {
		return false, word + " -> Error: bad length "
	}
	if contains(word, common_pass) {
		return false, word + " -> Error: Too Common"
	}

	new_string, isasc := isASCII(word)

	if !isasc {
		return false, new_string + " -> Error: Invalid Characters"
	}
	return true, ""

}

func contains(s string, farray []string) bool {
	for _, a := range farray {
		if a == s {
			return true
		}
	}
	return false

}

func check_length(s string) bool {
	return len(s) >= MIN_LENGTH && len(s) <= MAX_LENGTH
}

func isASCII(st string) (string, bool) {
	s := []byte(st)
	isasc := true
	// Tests if s is in the ASCII charset range
	for i := 0; i < len(s); i++ {
		if s[i] < 0 || s[i] > 127 {
			isasc = false
			s[i] = '*'

		}
	}
	if isasc == true {
		return string(s), true
	} else {
		return string(s), false
	}
}

func load_weak_file() []string {
	// Loads weak password file, and returns a slice
	file, err := os.Open(WEAK_FILE)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	a := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		a = append(a, scanner.Text())
	}
	return a
}
