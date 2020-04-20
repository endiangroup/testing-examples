Feature: Shopping Cart Summary
	As a Customer
	I want to see my shopping cart
	So I can see what I'm buying

	Scenario: Cart summary totals
		Given I have added a "Hairdryer" for 150 to my cart
		And I have added a "Hammer" for 50 to my cart
		When I view my cart
		Then I should see:
			| Hairdryer | £150.00 |
			| Hammer    | £50.00  |
			| Total     | £200.00 |

	Scenario: Cart summary with voucher totals
		Given I have added a "Hairdryer" for 150 to my cart
		And I have added a "Hammer" for 50 to my cart
		And I have added a voucher worth 50
		When I view my cart
		Then I should see:
			| Hairdryer | £150.00 |
			| Hammer    | £50.00  |
			| Voucher   | -£50.00 |
			| Total     | £150.00 |

	Scenario: Carty summary with tax

	Scenario: Carty summary with tax in foreign country
