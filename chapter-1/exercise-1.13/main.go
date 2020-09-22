package main
import(
	"fmt"
	"time"
)

func main() {
	var count1 *int
	count2 := new(int)
	countTemp := 5
	count3 := &countTemp
	t1 := &time.Time{}
	tTemp := time.Now()
	t2 := &tTemp
	fmt.Printf("count1: %v %T\n", count1, count1)
	// fmt.Printf("countTemp: %#v\n", countTemp)
	fmt.Printf("count2: %#v\n", count2)
	fmt.Printf("count3: %#v\n", count3)
	fmt.Printf("time: %#v\n", t1)
	fmt.Printf("time: %#v\n", t2)
}