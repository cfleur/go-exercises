package main
import(
	"fmt"
	"time"
)

func main() {
	dayBorn := time.Friday

	switch dayBorn {
	case time.Monday, time.Tuesday, time.Wednesday,
		time.Thursday, time.Friday:
		fmt.Println("Weekday")
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend")
	default:
		fmt.Println(dayBorn, "is not a day")
	}
}