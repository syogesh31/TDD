package main

import (
	"testing"
)

func TestMultiplication(t *testing.T) {
	fiver := money{
		amount: 5, currency: "USD",
	}
	actualTenner := fiver.times(2)
	expectedTenner := money{amount: 10, currency: "USD"}

	if actualTenner != expectedTenner {
		t.Errorf("Expected %#v , got : [%#v]", expectedTenner, actualTenner)
	}

}

func TestMultiplicationWithEuros(t *testing.T) {
	fiveEuro := money{
		currency: "EUR",
		amount:   5,
	}

	tenEuroActual := fiveEuro.times(2)
	tenEuroExpected := money{currency: "EUR", amount: 10}

	if tenEuroActual != tenEuroExpected {
		t.Errorf("Expected %#v , got : [%#v]", tenEuroExpected, tenEuroActual)
	}

}

func TestDivideKoreanWon(t *testing.T) {
	givenMoney := money{currency: "KRW", amount: 7}
	actualMoney := givenMoney.dividBy(2)
	expectedMoney := money{currency: "KRW", amount: 3.5}

	assertEqual(actualMoney, expectedMoney, t)
}

func assertEqual(actualMoney money, expectedMoney money, t *testing.T) {
	if actualMoney != expectedMoney {
		t.Errorf("Expected %#v , got : [%#v]", expectedMoney, actualMoney)
	}
}

type money struct {
	amount   float64
	currency string
}

func (m money) times(times int) money {
	return money{currency: m.currency, amount: m.amount * float64(times)}
}

func (m money) dividBy(divisor int) money {
	return money{amount: m.amount / float64(divisor), currency: m.currency}
}
