package retro

func NewShoppingCart(items ...CartItem) *ShoppingCart {
	return &ShoppingCart{
		items: items,
	}
}

type ShoppingCart struct {
	items []CartItem
}

func (s *ShoppingCart) Total() float64 {
	total := float64(0)
	for _, item := range s.items {
		total += item.Amount
	}

	return total
}

func (s *ShoppingCart) AddItem(name string, amount float64) {
	s.items = append(s.items, CartItem{Name: name, Amount: amount})
}

type CartItem struct {
	Name   string
	Amount float64
}
