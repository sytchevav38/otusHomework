package main

import (
	"fmt"
	"lesson7/v2/methods/models"
)

func main() {
	u := models.User{
		Name:     "Andrey",
		Age:      35,
		LastName: "Sych",
	}
	fmt.Println(u.Name)
	fmt.Println(u.GetName())
	// f := models.Friend{
	// 	Name:     "Dmitry",
	// 	Age:      34,
	// 	LastName: "Strela",
	// }

	// wallet := models.Wallet{}
	// fmt.Println(
	// 	"У меня сейчас столько денег: ",
	// 	wallet.HowManyMoney(),
	// )
	// wallet.AddMoney(100) // А как в него деньги то попали, если у нас нет countMoney в этой переменной структуре?!
	// fmt.Println(
	// 	"У меня денег после ЗП: ",
	// 	wallet.HowManyMoney(),
	// )
	// wallet.MinusMoney(12)
	// fmt.Println(
	// 	"У меня денег после магазина: ",
	// 	wallet.HowManyMoney(),
	//	)

	nowIHave := int64(100)

	wallet := models.Wallet{}

	wallet.AddMoney(nowIHave)

	nowIHave2 := models.NewWallet(7456789412315498632)

	fmt.Println(nowIHave2.HowManyMoney())
}
