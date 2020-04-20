package retro

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/cucumber/godog"
	"github.com/cucumber/messages-go/v10"
)

var (
	cart   = NewShoppingCart()
	output = ""
)

func TestMain(m *testing.M) {
	format := "progress"
	for _, arg := range os.Args[1:] {
		if arg == "-test.v=true" { // go test transforms -v option
			format = "pretty"
			break
		}
	}

	status := godog.RunWithOptions("godog", func(s *godog.Suite) {
		s.Step(`^I have added a "([^"]*)" for (\d+) to my cart$`, iHaveAddedAForToMyCart)
		s.Step(`^I have added a voucher worth (\d+)$`, iHaveAddedAVoucherWorth)
		s.Step(`^I view my cart$`, iViewMyCart)
		s.Step(`^I should see:$`, iShouldSee)
		s.AfterScenario(func(_ *messages.Pickle, _ error) {
			cart = NewShoppingCart()
			output = ""
		})
		godog.SuiteContext(s)
	}, godog.Options{
		Format: format,
		Paths:  []string{"features"},
	})

	if st := m.Run(); st > status {
		status = st
	}
	os.Exit(status)
}

func iHaveAddedAForToMyCart(name string, amount int) error {
	cart.AddItem(name, float64(amount))

	return nil
}

func iViewMyCart() error {
	output = ShoppingCartPrinter(cart)

	return nil
}

func iShouldSee(table *messages.PickleStepArgument_PickleTable) error {
	for _, row := range table.GetRows() {
		name := row.Cells[0].Value
		amount := row.Cells[1].Value

		expectedLine := fmt.Sprintf("%s: %s", name, amount)

		if !strings.Contains(output, expectedLine) {
			return errors.New("Cant find line: " + expectedLine + " in: " + output)
		}
	}

	return nil
}

func iHaveAddedAVoucherWorth(value int) error {
	cart.AddItem("Voucher", -float64(value))

	return nil
}
