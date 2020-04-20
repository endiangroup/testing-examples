package retro

import "time"

func Thursday10(cart *ShoppingCart) {
	if time.Now().Weekday() == time.Thursday {
		discountAmount := cart.Total() * 0.1
		cart.AddItem("Thursday 10% Discount", -discountAmount)
	}
}
