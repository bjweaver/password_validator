package main

import (
	"testing"
)

func TestLen(t *testing.T) {
	const SHORTP = "abcd"
	const P = "aaaaaaaa"
	const LONGP = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	// error =
	pass := check_length(P) == true && check_length(SHORTP) == false && check_length(LONGP) == false
	if !pass {
		t.Errorf("Lenth check failed")
	}
}

func TestASCII(t *testing.T) {
	pass := false
	const ASTRING = "abcdff"
	const BSTRING = "abcd\x8ac"

	if _, result := isASCII(ASTRING); result == true {
		if _, result := isASCII(BSTRING); result == false {
			pass = true
		}
	}

	// pass := isASCII(ASTRING) == true && isASCII(BSTRING) == false

	if !pass {
		t.Errorf("ASCII check failed")
	}
}

func TestCommon(t *testing.T) {
	common_pass := load_weak_file()
	pass := contains("password", common_pass) == true
	if !pass {
		t.Errorf("common word check failed")
	}
}

func TestRegression(t *testing.T) {
	common_pass = load_weak_file()

	pass := true
	var m string

	result, _ := test_password("dgdsgbdgegsg")
	if result == false {
		pass = false
	}

	result, _ = test_password("a")
	if result == true {
		pass = false
	}

	result, m = test_password("password")
	if result == true {
		pass = false
	}

	if !pass {
		t.Errorf("error " + m)
	}

}
