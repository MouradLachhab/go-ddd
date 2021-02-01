package domain

const (
	MinAmount    = 0
	CadToUsdRate = 0.79
)

type Money struct {
	usd float32
}

type MoneyActions interface {
	AddUsd(m *Money, amount float32)
	AddCad(m *Money, amount float32)
}

func MoneyFromCad(amount float32) (*Money, error) {
	return &Money{
		usd: amount * CadToUsdRate,
	}, nil
}

func MoneyFromUsd(amount float32) (*Money, error) {

	return &Money{
		usd: amount,
	}, nil
}

func (m *Money) AddUsd(amount float32) {
	m.usd += amount
}

func (m *Money) AddCad(amount float32) {
	m.usd += (amount * CadToUsdRate)
}

func (m *Money) Usd() float32 {
	return m.usd
}

func (m *Money) Cad() float32 {
	return m.usd / CadToUsdRate
}
