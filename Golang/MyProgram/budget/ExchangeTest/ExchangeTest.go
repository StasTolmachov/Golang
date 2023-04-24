package main

import "fmt"

func main() {
	var NokNok, NokUah, EurNok, EurEur, EurUsd, EurUah, UsdNok, UsdUsd, UsdUah, UahUah float64
	var Nok, Eur, Usd, Uah string
	Nok = "Nok"
	Eur = "Eur"
	Usd = "Usd"
	Uah = "Uah"
	var TotalNok, TotalEur, TotalUsd, TotalUah float64
	var Payment string
	var Currency string
	var CurrencyValue float64 //значение курса в форме
	var Amount float64
	var Sum float64

	NokNok = 1
	// NokEur = 0.086
	// NokUsd = 0.094
	NokUah = 3.445
	EurNok = 11.655
	EurEur = 1
	EurUsd = 1.098
	EurUah = 40.151
	UsdNok = 10.615
	// UsdEur = 0.911
	UsdUsd = 1
	UsdUah = 36.569
	// UahNok = 0.29
	// UahEur = 0.025
	// UahUsd = 0.027
	UahUah = 1


	// форма на сайте
	if Payment == Nok && Currency == Nok {
		CurrencyValue = NokNok
		Sum = Amount * NokNok
	}
	if Payment == Nok && Currency == Eur {
		CurrencyValue = EurNok
		Sum = Amount * EurNok
	}
	if Payment == Nok && Currency == Usd {
		CurrencyValue = UsdNok
		Sum = Amount * UsdNok
	}
	if Payment == Nok && Currency == Uah {
		CurrencyValue = NokUah
		Sum = Amount / NokUah
		
	}

	// 16000грн с норвежской карты в нок

	Payment = Nok
	Currency = Uah
	Amount = 16000
	// 16000грн / 3,445 = 4644 кроны

	// Payment == Nok
	if Payment == Nok && Currency == Nok {
		TotalNok = Amount * NokNok
		TotalEur = Amount / EurNok
		TotalUsd = Amount / UsdNok
		TotalUah = Amount * NokUah
	}
	if Payment == Nok && Currency == Eur {
		TotalNok = Amount * EurNok
		TotalEur = Amount * EurEur
		TotalUsd = Amount * EurUsd
		TotalUah = Amount * EurUah
	}
	if Payment == Nok && Currency == Usd {
		TotalNok = Amount * UsdNok
		TotalEur = Amount / EurUsd
		TotalUsd = Amount * UsdUsd
		TotalUah = Amount * UsdUah
	}
	if Payment == Nok && Currency == Uah {
		TotalNok = Amount / NokUah
		TotalEur = Amount / EurUah
		TotalUsd = Amount / UsdUah
		TotalUah = Amount * UahUah
	}
	// Payment == Eur
	if Payment == Eur && Currency == Nok {
		TotalNok = Amount * NokNok
		TotalEur = Amount / EurNok
		TotalUsd = Amount / UsdNok
		TotalUah = Amount * NokUah
	}
	if Payment == Eur && Currency == Eur {
		TotalNok = Amount * EurNok
		TotalEur = Amount * EurEur
		TotalUsd = Amount * EurUsd
		TotalUah = Amount * EurUah
	}
	if Payment == Eur && Currency == Usd {
		TotalNok = Amount * UsdNok
		TotalEur = Amount / EurUsd
		TotalUsd = Amount * UsdUsd
		TotalUah = Amount * UsdUah
	}
	if Payment == Eur && Currency == Uah {
		TotalNok = Amount / NokUah
		TotalEur = Amount / EurUah
		TotalUsd = Amount / UsdUah
		TotalUah = Amount * UahUah
	}


	fmt.Println(TotalNok, TotalEur, TotalUsd, TotalUah)


}

