package models

type Wallet struct {
	countMoney int64
}

func NewWallet(money int64) Wallet {
	return Wallet{
		countMoney: money,
	}
}

func (w Wallet) HowManyMoney() int64 {
	return w.countMoney
}

func (w *Wallet) AddMoney(money int64) {
	w.countMoney += money
}

func (w *Wallet) MinusMoney(money int64) {
	w.countMoney -= money
}
