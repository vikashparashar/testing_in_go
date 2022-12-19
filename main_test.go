package main

import (
	"testing"
)

func Test_isPrime(t *testing.T) {

	// check for zero

	result, msg := isPrime(0)

	if result {
		t.Errorf("with %d as test parameter , got true , but excepted false", 0)
	}

	if msg != "0 is not a prime number , by defination !" {
		t.Error("wrong message returned:", msg)
	}

	// check for one
	result, msg = isPrime(1)

	if result {
		t.Errorf("with %d as test parameter , got true , but excepted false", 1)
	}

	if msg != "1 is not a prime number , by defination !" {
		t.Error("wrong message returned:", msg)
	}

	// check for positive number
	result, msg = isPrime(7)

	if result != true {
		t.Errorf("with %d as test parameter , got false , but excepted true", 7)
	}

	if msg != "7 is a prime number !" {
		t.Error("wrong message returned:", msg)
	}

	// check for negative number
	result, msg = isPrime(-32)

	if result {
		t.Errorf("with %d as test parameter , got true , but excepted false", -32)
	}

	if msg != "negative numbers are not prime , by defination !" {
		t.Error("wrong message returned:", msg)
	}

}

/*
	after taking all 4 possible outcomes
	Archieved PASS
	coverage: 63.6% of statements
	ok      01_Prime_Application    0.312s

*/

/*

	1.1 Test File Should Be With Main Package File

	1.2 Test File Name Shoud Be xxx_test.go

	1.3 Test Function Name Sould Be As  func TestXxx(t *testing.T){}

	1.4 To run test file use command
	go test
	go test -v



	1.5 To check the code coverge by test can use this command
	go test -cover

*/
