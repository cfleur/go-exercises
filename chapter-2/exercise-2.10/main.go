package main
import(
	"fmt"
)

func main() {
	config := map[string]string {
		"debug": "1",
		"LogLevel": "warn",
		"version": "1.2.1",
	}
	for key, value := range config {
		fmt.Println(key, "=", value)
	}
}