package bppformula

import (
	"testing"
)

func TestS(t *testing.T)() {
	// Expected values taken from http://www.davidhbailey.com/dhbpapers/bbp-alg.pdf
	// page 4

	// Single tests for summations
	d := 1000000
	tests := map[int]map[string]string{
		0: {
			"expected":	"0.1810395338014360678534893462",
			"calc":		S(1, 1, d).Text('f', 30)[0:30],
		},
		1: {
			"expected":	"0.7760655498078074613722975943",
			"calc":		S(1, 4, d).Text('f', 30)[0:30],
		},
		2: {
			"expected":	"0.3624585640705741420683343355",
			"calc":		S(1, 5, d).Text('f', 30)[0:30],
		},
		3: {
			"expected":	"0.3861386739520148480012151865",
			"calc":		S(1, 6, d).Text('f', 30)[0:30],
		},
	}

	for i := 0; i < 4; i++ {
		if (tests[i]["calc"] != tests[i]["expected"]) {
			t.Error("Test", i, ") Expected", tests[i]["expected"], "got", tests[i]["calc"])
		}
	}

}

func TestPi(t *testing.T)() {
	res, _ := ToStringBase(Pi(1), 16, 8)
	if res != "243F6A88" {
		t.Error("Expected 243F6A88, got", res)
	}

	res, _ = ToStringBase(Pi(1000001), 16, 8)
	if res != "6C65E52C" {
		t.Error("Expected 6C65E52C, got", res)
	}
}
