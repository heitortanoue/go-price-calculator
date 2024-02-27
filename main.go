package main

import (
	"fmt"
	"heitortanoue/price-calculator/cmdmanager"
	"heitortanoue/price-calculator/prices"
)

func main() {
	taxes := []float64{0, 5, 15} // %

	for _, tax := range taxes {
		// fm := filemanager.New("prices.txt", fmt.Sprintf("results_%.2f.json", tax))
		cmd := cmdmanager.New()

		newJob := prices.NewTaxIncludedPriceJob(tax, cmd)
		err := newJob.Process()
		
		if err != nil {
			fmt.Println(err)
		}
	}
}
