package main
import(
	"fmt"
)

func main() {
	words := map[string]int{
		"Gonna": 3,
		"You": 3,
		"Give": 2,
		"Never": 1,
		"Up":  4,
	}
	var popWord string
	var popCount int = 0

	for key, value := range words {
		if popCount < value {
			popCount = value
			popWord = key
		}
	}

	fmt.Println("Most popular word:", popWord)
	fmt.Println("With count of:", popCount)
}