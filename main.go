package main

import "errors"
import "fmt"
import "os"
import "time"
import "math"

func main() {

	fmt.Println("SER 450 - Project 1")
	fmt.Println("Todd Martin")

	var t string
	var err error
	t, err = findPrimes(100000)
	if err != nil {
		fmt.Println("Input given is less than 1")
		os.Exit(1)
	}
	fmt.Printf("Time for finding primes up till 100,000: %s\n", t)
	t, err = findPrimes(1000000)
	if err != nil {
		fmt.Println("Input given is less than 1")
	}
	fmt.Printf("Time for finding primes up till 1,000,000: %s\n", t)
	t, err = findPrimes(10000000)
	if err != nil {
		fmt.Println("Input given is less than 1")
	}
	fmt.Printf("Time for finding primes up till 10,000,000: %s\n", t)
}

// reference for algorithm
// https://www.algolist.net/Algorithms/Number_theoretic/Sieve_of_Eratosthenes
func findPrimes(n int) (string, error) {
	var isComposite []bool
	var primes []int

	if n < 1 {
		return "", errors.New("Integer provided is less than 1")
	}

	// initialize array
	for i := 0; i < n; i++ {
		isComposite = append(isComposite, false)
	}

	start := time.Now()

	for  p := 2; float64(p) < math.Sqrt(float64(n)); p++ {
		if isComposite[p] == false {
			// add primte to list
			primes = append(primes, p)
			for k := p; k < n; k++ {
				if k*p < n {
					isComposite[k*p] = true;
				}
			}
		}
	}

	t := time.Now()

	return time.Duration.String(t.Sub(start)), nil
}
