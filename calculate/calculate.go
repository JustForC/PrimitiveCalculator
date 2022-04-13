package calculate

import (
	"strings"

	"github.com/JustForC/PrimitiveCalculator/model"
)

func CheckSymbol(symbol string, number *model.Number) {
	if strings.Compare("+", symbol) == 0 {
		Plus(number)
	} else if strings.Compare("-", symbol) == 0 {
		Minus(number)
	} else if strings.Compare("*", symbol) == 0 {
		Times(number)
	} else if strings.Compare("/", symbol) == 0 {
		Divide(number)
	} else {
		Modular(number)
	}
}

func Plus(number *model.Number) {
	number.Calculation = number.NumberOne + number.NumberTwo
}

func Minus(number *model.Number) {
	number.Calculation = number.NumberOne - number.NumberTwo
}

func Times(number *model.Number) {
	number.Calculation = number.NumberOne * number.NumberTwo
}

func Divide(number *model.Number) {
	number.Calculation = number.NumberOne / number.NumberTwo
}

func Modular(number *model.Number) {
	number.Calculation = number.NumberOne % number.NumberTwo
}
