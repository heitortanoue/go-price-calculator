package conversion

import (
	"strconv"
)

func StringsToFloats(lines []string) ([]float64, error) {
	values := make([]float64, 0, len(lines))

	for _, line := range lines {
		convertedValue, err := strconv.ParseFloat(line, 64)
		if err != nil {
			return nil, err
		}
		values = append(values, convertedValue)
	}

	return values, nil
}
