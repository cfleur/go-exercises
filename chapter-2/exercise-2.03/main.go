package main
import(
	"fmt"
)

func main() {
	input := -10
	if input < 0 {
		fmt.Println("negative number error")
	} else if input%2 == 0 {
		fmt.Println(input, "is even")
	} else {
		fmt.Println(input, "is odd")
	}
}