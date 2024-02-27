package cmdmanager

import "fmt"

type CMDManager struct {
}

func (cmd CMDManager) ReadLines() ([]string, error) {
	fmt.Println("Insira os valores. Pressione ENTER após cada valor.")

	var prices []string

	for {
		var price string
		fmt.Scan(&price)
		if price == "0" {
			break
		}
		prices = append(prices, price)
	}

	return prices, nil
}

func (cmd CMDManager) WriteLines(data any) error {
	_, err := fmt.Println(data)
	return err
}

func New() CMDManager {
	return CMDManager{}
}
