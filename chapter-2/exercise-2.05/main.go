package main
import(
	"fmt"
	"time"
)

func main() {
	dayBorn := time.Friday

	switch dayBorn {
	case time.Monday:
		fmt.Println("Monday")
	case time.Tuesday:
		fmt.Println("Tuesday")
	default:
		fmt.Println(dayBorn, "is not a day")
	}
}