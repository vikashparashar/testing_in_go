package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// print a welcome message
	intro()

	// create a channel to indicate when the user want to quit
	doneChan := make(chan bool)

	// start a goroutine to read a user input and run program
	go readUSerInput(doneChan)

	// block until the doneChan get a value
	<-doneChan
	// close the channel
	close(doneChan)
	// say goodbye
	fmt.Println("Good Bye !")
}

func readUSerInput(doneChan chan bool) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		res, done := checkNumbers(scanner)

		if done {
			doneChan <- true
			return
		}

		fmt.Println(res)
		prompt()
	}
}

func checkNumbers(scanner *bufio.Scanner) (string, bool) {
	// read user input
	scanner.Scan()

	// check to see if user want to quit
	if strings.EqualFold(scanner.Text(), "q") {
		return "", true
	}

	// try to convert what the user type into int

	numTocheck, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return "Please enter a whole number!", false
	}
	_, msg := isPrime(numTocheck)
	return msg, false
}

func intro() {
	fmt.Println()
	fmt.Println("Is It Prime ?")
	fmt.Println("-------------")
	fmt.Println("Enter a whole number , and we'll tell you if it is a prime number or not . Enter q to quit.")
	prompt()
}

func prompt() {
	fmt.Print("-> ")
}
func isPrime(n int) (bool, string) {
	// 0 and 1 are not prime by defination

	if n == 0 || n == 1 {
		return false, fmt.Sprintf("%d is not a prime number , by defination !", n)
	}

	// checking for negative numbers
	// because negative numbers are not prime

	if n < 0 {
		return false, "negative numbers are not prime , by defination !"
	}

	// use the modules operator repeatedly to see if we have a prime number
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			// not a prime number
			return false, fmt.Sprintf("%d is not a prime number because it is divisible by %d .", n, i)
		}
	}
	return true, fmt.Sprintf("%d is a prime number !", n)
}
