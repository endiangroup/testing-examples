package retro

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNotThursday(t *testing.T) {
	cart := NewShoppingCart(CartItem{"Test line item", 100})

	now := time.Now()
	notThursdayTime := time.Date(2020, 1, 1, 0, 0, 0, 0, now.Location())
	Thursday10(notThursdayTime, cart)

	assert.Equal(t, float64(100), cart.Total())
}

func TestThursday(t *testing.T) {
	cart := NewShoppingCart(CartItem{"Test line item", 100})

	Thursday10(time.Now(), cart)

	assert.Equal(t, float64(90), cart.Total())
}
