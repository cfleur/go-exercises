package main
import(
	"fmt"
	"time"
)

func main() {
	switch dayBorn := time.Sunday; {
	case dayBorn == time.Sunday || dayBorn == time.Saturday: 
		fmt.Println("Weekend")
	default:
		fmt.Println(dayBorn, "Weekday")
	}
}