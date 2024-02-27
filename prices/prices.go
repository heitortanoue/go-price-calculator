package prices

import (
	"fmt"
	"heitortanoue/price-calculator/conversion"
	"heitortanoue/price-calculator/iomanager"
	"math"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64             `json:"tax_rate"`
	Prices            []float64           `json:"prices"`
	TaxIncludedPrices map[string]float64  `json:"tax_included"`
	IOManager         iomanager.IOManager `json:"-"` // ignoramos esse campo no json
}

func (job *TaxIncludedPriceJob) Process() error {
	err := job.LoadData()
	if err != nil {
		return err
	}

	result := make(map[string]float64)

	for _, price := range job.Prices {
		calculatedPrice := price * (1 - job.TaxRate/100)
		key := fmt.Sprintf("%.2f", price)
		result[key] = math.Round(calculatedPrice*100) / 100
	}

	job.TaxIncludedPrices = result
	job.IOManager.WriteLines(job)
	return nil
}

func (job *TaxIncludedPriceJob) LoadData() error {
	lines, err := job.IOManager.ReadLines()
	if err != nil {
		return err
	}

	prices, err := conversion.StringsToFloats(lines)
	if err != nil {
		return err
	}

	job.Prices = prices
	return nil
}

func NewTaxIncludedPriceJob(taxRate float64, IOManager iomanager.IOManager) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate:   taxRate,
		IOManager: IOManager,
	}
}
