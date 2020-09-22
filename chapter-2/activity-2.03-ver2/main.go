package main
import(
	"fmt"
	"math/rand"
	"time"
)

func initializeSlice(seed int64, length int) []int {
	// set seed for random number generator:
	rand.Seed(seed)

	return rand.Perm(length)

/*	Note: Perm returns, as a slice of n ints, 
	a pseudo-random permutation of the integers [0,n)
	more info at: https://golang.org/pkg/math/rand/#Perm
	*/
}

func bubbleSort(numbers []int, itr *int) ([]int, *int) {
/*	Recursive version. Exit case is that the number of swaps done in an iteration 
	(one pass through the numbers list) is less than 1. 

	To keep track of the number of times this function is
	called, a pointer is passed to an external veriable, and its value is incremented
	once in each cycle to count this. Make sure that it is initialized as 0 externally
	to get an accurate count. 
*/
	var numOfSwaps int

		for i := 1; i < len(numbers); i ++ {
			if numbers[i-1] > numbers[i] {
				numbers[i-1], numbers[i] = numbers[i], numbers[i-1]
				numOfSwaps ++
			}
		}

		*itr ++
		fmt.Printf("Iteration: %v, Swaps: %v, Data: %v\n", *itr, numOfSwaps, numbers)

		if numOfSwaps < 1 {
			return numbers, itr
		}
	return bubbleSort(numbers, itr)
}

func main() {
	// Choose a random seed. Using a constant will result in the same behaviour
	// every time the program is run:
	randSeed := time.Now().UnixNano()
	listLength := 25
	// initialize random list of intergers:
	numbers := initializeSlice(randSeed, listLength)
	// create a variable for checking how many times the bubble sort ran:
	iterations := 0 

	fmt.Printf("\n\n\nBUBBLE SORT\n\n\n")
	fmt.Printf("Original data: %v\n\n", numbers)

	bubbleSort(numbers, &iterations)

	fmt.Printf("\nSorted data: %v\nSorted in %v iterations.\n\n", numbers, iterations)
}