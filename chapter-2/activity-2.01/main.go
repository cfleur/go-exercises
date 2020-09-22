package main
import(
	"fmt"
	"strconv"
)

func getWord(num int) string {
/* 
	This creates an array that holds all the the possible values 
	of a number: The number itself, "Fizz", "Buzz" or "Fizz" and "Buzz".
	This is done to make the functionality only two conditional statements
	are needed rather than 3 or a nested conditional statement. It tries to 
	minimize repitition of checks. 
	
	If a number's value should be something other than itself, then it's place 
	in the array is set to be an empty string. The array is then concatinated
	into a single string that is returned.
	
	Caveats:
	Unfortunately, this method requires a nested for loop. Therefore, if there
	is a need for having a long array of words to concatinate, it will lose
	efficiency */

	words := []string{strconv.Itoa(num)}
	var word string

	if num % 3 == 0 {
		words[0] = ""
		words = append(words, "Fizz")
	} 
	
	if num % 5 == 0 {
		words[0] = ""
		words = append(words, "Buzz")
	}

	for w := 0; w < len(words); w ++ {
		word += words[w]
	} 
	return word
}

func main() {
	for i := 1; i <= 100; i ++ {
		word := getWord(i)
		fmt.Println(word)
	}
}