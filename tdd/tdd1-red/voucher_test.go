package retro

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNotThursday(t *testing.T) {
	cart := NewShoppingCart(CartItem{"Test line item", 100})

	Thursday10(cart)

	assert.Equal(t, float64(100), cart.Total())
}
