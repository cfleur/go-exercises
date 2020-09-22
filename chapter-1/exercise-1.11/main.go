package main
import(
	"fmt"
)

func main() {
	visits := 15
	fmt.Println("Number of visits: ", visits)
	fmt.Println("First visit :", visits == 1)
	fmt.Println("Returning visit :", visits != 1)
	fmt.Println("Silver member :", visits >= 10 && visits < 21)
	fmt.Println("Gold member :", visits > 20 && visits <= 30)
	fmt.Println("Platinum member :", visits > 30)
}