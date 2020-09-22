package main
import(
	"fmt"
)

func main() {
	logLevel := "デバッグ"
	runeVal := []rune(logLevel)

	fmt.Println("\nString value and length:", logLevel, len(logLevel), 
	"\nRune value and length:", runeVal, len(runeVal))
	fmt.Println("\nIndex, rune and character:")

	for index, runeVal := range logLevel {
		fmt.Println(index, runeVal, string(runeVal))
	}
}