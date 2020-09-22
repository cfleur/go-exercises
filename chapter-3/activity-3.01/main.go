package main
import(
	"fmt"
)

func calculateTax(price, tax float64) float64 {
	return price * tax
}

func main() {

	bill := map[string]map[string]float64{
		"cake" : map[string]float64{
			"cost" : 0.99,
			"tax": 0.075,
		},
		"milk" : map[string]float64{
			"cost" : 2.75,
			"tax" : 0.015,
		},
		"butter" : map[string]float64{
			"cost" : 0.87,
			"tax" : 0.02,
		},
	}

	var total float64
	for _, item := range bill {
		total += calculateTax(item["cost"], item["tax"])
	} 
	fmt.Println("Total sales tax: ", total)


}