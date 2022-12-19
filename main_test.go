package main

import (
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
		{"not prime", 0, false, "0 is not a prime number , by defination !"},
		{"not prime", 1, false, "1 is not a prime number , by defination !"},
		{"not prime", -7, false, "negative numbers are not prime , by defination !"},
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

/*

	by this table test we again got 72.7 % coverage for our code


	PS E:\go_testing_course\01_Prime_Application> go test -cover
	PASS
	coverage: 72.7% of statements
	ok      01_Prime_Application    0.313s


*/

/*

	At A Galance Check To See Which Of My Code Lines Are Covered By The Test And Whcih Not By Using Below
	Command .

	go test -coverprofile=coverage.out


*/
