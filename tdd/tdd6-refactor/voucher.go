package retro

import "time"

func Thursday10(timeNow time.Time, cart *ShoppingCart) {
	if timeNow.Weekday() == time.Thursday {
		discountAmount := cart.Total() * 0.1
		cart.AddItem("Thursday 10% Discount", -discountAmount)
	}
}
