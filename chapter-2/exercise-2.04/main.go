package main
import(
	"fmt"
	"errors"
)

func validate(input int) error {
	if input < 0 {
		return errors.New("negative number error")
	} else if input > 100 {
		return errors.New("too high input error (over 100)")
	} else if input % 7 == 0{
		return errors.New("divisible by 7 error")
	} else {
		return nil
	}
}

func main() {
	input := 21
	if err := validate(input); err != nil {
		fmt.Println(err)
	} else if input%2 == 0 {
		fmt.Println(input, "is even")
	} else {
		fmt.Println(input, "is odd")
	}
}