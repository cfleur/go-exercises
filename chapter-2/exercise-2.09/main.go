package main
import(
	"fmt"
)

func main() {
	names := []string{"Dude1", "Chick1", "Dude2", "SheMale1"}
	fmt.Println("\n\n Here is a list of names:\n", names, "\n")
	for index:= 0; index < len(names); index++ {
		fmt.Println("\t", names[index], "is number", index+1)
	}
}