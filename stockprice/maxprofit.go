package main

import "fmt"

func main() {
	prices := []int{4, 11, 2, 20, 59, 80}
	maxProfit := MaxProfit(prices, 2)
	fmt.Println()
	fmt.Println("Profit :", maxProfit)
}

func MaxProfit(prices []int, k int) (totalProfit int) {
	n := len(prices)
	nCalc := 0
	buyPrice := prices[0]
	maxProfit := 0
	sellPrice := 0
	actualProfit := 0

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			profit := prices[j] - buyPrice
			if profit > maxProfit {
				actualProfit = profit
				sellPrice = prices[j]
				nCalc++
			} else {
				if i == 0 {
					fmt.Printf("Beli: %d ", buyPrice)
				} else {
					if prices[j] > 2 {
						buyPrice = 2
					} else {
						buyPrice = prices[j]
					}

					fmt.Printf("Beli: %d ", buyPrice)
				}
				nCalc++
				break
			}
		}

		if actualProfit > 0 {
			fmt.Printf("Jual: %d ", sellPrice)
			totalProfit = totalProfit + actualProfit
		}

		actualProfit = 0

		if n == nCalc {
			break
		}
	}

	return

}
