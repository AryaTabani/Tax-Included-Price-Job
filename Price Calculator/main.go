package main

import (
	"./cmdmanager"
	"./prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	for _, taxRate := range taxRates {
		//fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		cmdm := cmdmanager.New()
		proiceJob := prices.NewTaxIncludedPriceJob(cmdm, taxRate)
		//proiceJob := prices.NewTaxIncludedPriceJob(fm,taxRate)  both fm and cmdm are working properly
		proiceJob.Process()

	}

}
