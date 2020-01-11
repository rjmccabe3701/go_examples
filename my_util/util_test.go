/*
Run
	go test
in this directory and it will run all methods
in this file starting with "Test"
*/
package my_util

import "testing"

import "fmt"

func TestMyUpper(t *testing.T) {
	want := "HELLO"
	if got := MyUpper("Hello"); got != want {
		t.Errorf("Bad!")
	}
	fmt.Println("Good")
}

//This failes (as it should)
func TestNothing(t *testing.T) {
	t.Errorf("bad")
}

// Doesn't start with "Test" so is not run
func WillNotRun(t *testing.T) {
	t.Errorf("bad")
}
