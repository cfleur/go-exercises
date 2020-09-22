package main
import(
	"fmt"
	"math/rand"
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

func main() {
	const randSeed = 934567
	const listLength = 9
	var numbers []int = initializeSlice(randSeed,listLength)
	var printLog = true

	fmt.Printf("\nOriginal data: %v\n\n", numbers)
	
	for iter := 0; true; iter ++ {
		var numOfSwaps int

		for i := 1; i < len(numbers); i ++ {
			if numbers[i-1] > numbers[i] {
				numbers[i-1], numbers[i] = numbers[i], numbers[i-1]
				numOfSwaps ++
			}
		}

		if printLog == true {
			fmt.Printf("Iteration: %v, Swaps: %v, Data: %v\n",iter, numOfSwaps, numbers)	
		}
		
		if numOfSwaps < 1 {
			break
		}
	}

	fmt.Printf("\nSorted data: %v\n\n", numbers)
}