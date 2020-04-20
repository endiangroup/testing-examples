package retro

import (
	"fmt"
	"math"
)

func ShoppingCartPrinter(cart *ShoppingCart) string {
	vouchers := []CartItem{}
	lineItems := []CartItem{}

	for _, lineItem := range cart.items {
		if lineItem.Amount < 0 {
			vouchers = append(vouchers, lineItem)
		} else {
			lineItems = append(lineItems, lineItem)
		}
	}

	outputStr := ""
	for _, lineItem := range lineItems {
		outputStr += fmt.Sprintf("%s: £%.2f\n", lineItem.Name, lineItem.Amount)
	}
	outputStr += "\n"
	for _, voucher := range vouchers {
		outputStr += fmt.Sprintf("%s: -£%.2f\n", voucher.Name, math.Abs(voucher.Amount))
	}
	outputStr += "\n"
	outputStr += fmt.Sprintf("Total: £%.2f", cart.Total())

	return outputStr
}
