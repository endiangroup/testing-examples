package retro

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTotal(t *testing.T) {
	cart := NewShoppingCart(
		CartItem{"Item 1", 10},
		CartItem{"Item 2", 50},
		CartItem{"Item 3", -5},
	)

	assert.Equal(t, float64(55), cart.Total())
}

func TestAddItem(t *testing.T) {
	cart := NewShoppingCart()
	expectedItem := CartItem{"Item 1", 50}

	cart.AddItem(expectedItem.Name, expectedItem.Amount)

	assert.Len(t, cart.items, 1)
	assert.Equal(t, expectedItem, cart.items[0])
}
