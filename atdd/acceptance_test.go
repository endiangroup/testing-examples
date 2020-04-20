package retro

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Customer_WantsToSeeTheBreakdownOfTheirShoppingCartWithVouchers(t *testing.T) {
	cart := NewShoppingCart(
		CartItem{"Hairdryer", 150},
		CartItem{"Voucher code", -10},
		CartItem{"Hammer", 30},
		CartItem{"Toothpaste", 15},
	)
	expectedCart := `Hairdryer: £150.00
Hammer: £30.00
Toothpaste: £15.00

Voucher code: -£10.00

Total: £185.00`

	printedCart := ShoppingCartPrinter(cart)

	assert.Equal(t, expectedCart, printedCart)
}
