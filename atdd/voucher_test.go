package retro

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func generateWeekWithDays(includeDays ...time.Weekday) []time.Time {
	dates := []time.Time{}

	now := time.Now()
	for nextWeek := now.AddDate(0, 0, 7); now.Before(nextWeek); now = now.Add(24 * time.Hour) {
		for _, includeDay := range includeDays {
			if now.Weekday() == includeDay {
				dates = append(dates, now)
			}
		}
	}

	return dates
}

func Test_ItDoesntApplyVoucherWhenItsNotThursday(t *testing.T) {
	testCases := generateWeekWithDays(
		time.Monday,
		time.Tuesday,
		time.Wednesday,
		// Not Thursday
		time.Friday,
		time.Saturday,
		time.Sunday,
	)

	for _, notThursdayTime := range testCases {
		cart := NewShoppingCart(CartItem{"Test line item", 100})

		Thursday10(notThursdayTime, cart)

		assert.Equal(t, float64(100), cart.Total())
	}
}

func Test_ItAppliesVoucherWhenItIsThursday(t *testing.T) {
	cart := NewShoppingCart(CartItem{"Test line item", 100})

	nextThursday := generateWeekWithDays(time.Thursday)[0]
	Thursday10(nextThursday, cart)

	assert.Equal(t, float64(90), cart.Total())
}
