package euler

import (
	"fmt"
	"strconv"
)

const (
	MILLION = 1_000_000
)

// https://projecteuler.net/problem=1
func MultiplesOfThreeAndFive() {
	const (
		THREE = 3
		FIVE  = 5
	)

	sum := 0

	for i := 1; i < 1000; i++ {
		isMultipleOfThree := (i % THREE) == 0
		isMultipleOfFive := (i % FIVE) == 0

		if isMultipleOfFive || isMultipleOfThree {
			sum += i
		}
	}

	fmt.Printf("%d\n", sum)
}

// https://projecteuler.net/problem=2
/*
	We can use temporary variables to generate fibonacci series
	instead of a list. But for practicing golang loops we use list
*/
func EvenFibonacciNumbers() {
	const fourMillion = MILLION * 4

	var accum []int

	accum = append(accum, 0)
	accum = append(accum, 1)

	sum := 0

	for {
		last := accum[len(accum)-1]
		lastButOne := accum[len(accum)-2]

		newSum := last + lastButOne

		if newSum > fourMillion {
			break
		}

		if newSum%2 == 0 {
			sum += newSum
		}

		accum = append(accum, newSum)
	}

	fmt.Println(accum)
	fmt.Println(sum)
}

// https://projecteuler.net/problem=3
func LargestPrimeFactor() {
	const (
		N = 600851475143
	)

	primes := sieve(10000)

	for i := len(primes) - 1; i >= 0; i-- {
		if N%primes[i] == 0 {
			fmt.Println(primes[i])
			return
		}
	}
}

// https://projecteuler.net/problem=4
func LargestPalindromeProduct() {
	const (
		MIN = 100
		MAX = 999
	)

	var currentMax = -1
	for i := MIN; i <= MAX; i++ {
		for j := MIN; j <= MAX; j++ {
			if isPalindrome(i*j) && currentMax < i*j {
				currentMax = i * j
			}
		}
	}

	fmt.Println(currentMax)
}

// https://projecteuler.net/problem=5
func SmallestMultiple() {
	var max = 1

	for i := 2; i <= 20; i++ {
		max *= i
	}

	for i := 21; i <= max; i++ {
		isDivisible := true
		for j := 2; j <= 20; j++ {
			if i%j != 0 {
				isDivisible = false
				break
			}
		}

		if isDivisible {
			fmt.Println(i)
			return
		}
	}
}

// https://projecteuler.net/problem=6
func SumSquareDifference() {
	const (
		N = 100
	)

	var squaredSum int = 0
	var sum int = 0

	for i := 0; i <= N; i++ {
		squaredSum += (i * i)
		sum += i
	}

	fmt.Println(sum*sum - squaredSum)
}

// https://projecteuler.net/problem=7
func ProblemSeven() {
	var primes = sieve(200_000)

	fmt.Println(primes[10000])
}

// https://projecteuler.net/problem=9
func SpecialPythagoreanTriplet() {
	for a := 1; a < 1000; a++ {
		for b := 1; b < 1000; b++ {
			for c := 1; c < 1000; c++ {
				if a+b+c == 1000 {
					if (a*a)+(b*b) == (c*c) || (b*b)+(c*c) == (a*a) || (a*a)+(c*c) == (b*b) {
						fmt.Println(a * b * c)
						return
					}
				}
			}
		}
	}
}

// https://projecteuler.net/problem=10
func SummationOfPrimes() {
	const (
		N = 2_000_000
	)

	var primes = sieve(N)

	sum := 0
	for _, prime := range primes {
		sum += prime
	}

	fmt.Println(sum)
}

func isPalindrome(n int) bool {
	s := strconv.Itoa(n)

	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}

	return true
}

func sieve(limit int) []int {
	isPrime := make([]bool, limit+1)

	for i := 2; i < limit; i++ {
		isPrime[i] = true
	}

	for i := 2; i < limit; i++ {
		if isPrime[i] {
			for j := 2; i*j < limit; j++ {
				isPrime[i*j] = false
			}
		}
	}

	primes := []int{}

	for i := 2; i < limit; i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
	}

	return primes
}
