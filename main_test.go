package main

import (
	"bufio"
	"io"
	"os"
	"strings"
	"testing"
)

func Test_isPrime(t *testing.T) {

	// this method is called table test

	primeTests := []struct {
		name       string
		testNumber int
		expected   bool
		msg        string
	}{
		{"prime", 7, true, "7 is a prime number !"},
		{"not prime", 8, false, "8 is not a prime number because it is divisible by 2 ."},
		{"zero", 0, false, "0 is not a prime number , by defination !"},
		{"one", 1, false, "1 is not a prime number , by defination !"},
		{"negative number", -7, false, "negative numbers are not prime , by defination !"},
	}

	for _, e := range primeTests {
		result, msg := isPrime(e.testNumber)
		if e.expected && !result {
			t.Errorf("%s expected true but got false", e.name)
		}
		if !e.expected && result {
			t.Errorf("%s expected false but got true", e.name)
		}
		if e.msg != msg {
			t.Errorf("%s : expected %s but got %s", e.name, e.msg, msg)
		}
	}
}

func Test_prompt(t *testing.T) {

	// save a copy of os.stdout
	oldOut := os.Stdout

	// create a read and write pipe
	r, w, _ := os.Pipe()
	os.Stdout = w
	prompt()
	// close the writer

	_ = w.Close()

	// reset os.Stdout what is was before

	os.Stdout = oldOut

	// read the ooutput of our promp function from our read pipe
	out, _ := io.ReadAll(r)
	if string(out) != "-> " {
		t.Errorf("incorrect promt : expted -> but got %s", string(out))
	}

}

/*
git commit -m "writing test for intro function"
*/

func Test_intro(t *testing.T) {

	// save a copy of os.stdout
	oldOut := os.Stdout

	// create a read and write pipe
	r, w, _ := os.Pipe()
	os.Stdout = w
	intro()
	// close the writer

	_ = w.Close()

	// reset os.Stdout what is was before

	os.Stdout = oldOut

	// read the ooutput of our promp function from our read pipe
	out, _ := io.ReadAll(r)

	if !strings.Contains(string(out), "Enter a whole number") {
		t.Errorf("intro text not correct ; got %s", string(out))
	}

}

func Test_checkNumbers(t *testing.T) {
	// res, _ := checkNumbers(bufio.NewScanner(os.Stdin))
	// if !strings.EqualFold(res, "7 is a prime number!") {
	// 	t.Errorf(("wrong return value"))
	// }

	input := strings.NewReader("7")
	reader := bufio.NewScanner(input)
	res, _ := checkNumbers(reader)
	if !strings.EqualFold(res, "7 is a prime number !") {
		t.Errorf("incorrect value returned; got %s", res)
	}
}
