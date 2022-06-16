package main

import (
	"fmt"
	"math"
)

func main() {
	const questionOne string = "Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш?"
	const questionTwo string = "Скільки груш ми можемо купити?"
	const questionThree string = "Скільки яблук ми можемо купити?"
	const questionFour string = "Чи ми можемо купити 2 груші та 2 яблука?"
	const applePrice float64 = 5.99
	const pearPrice float64 = 7.0
	const money float64 = 23.0

	// 1
	fmt.Println(questionOne)
	var result float64 = applePrice*9 + pearPrice*8
	var msg string = fmt.Sprint("9 яблук та 8 груш коштують - ", math.Ceil(result*100)/100)
	fmt.Println(msg)
	// 2
	fmt.Println(questionTwo)
	result = math.Floor(money / pearPrice)
	msg = fmt.Sprint("Ми можемо купити ", result, " груші")
	fmt.Println(msg)
	// 3
	fmt.Println(questionThree)
	result = math.Floor(money / applePrice)
	msg = fmt.Sprint("Ми можемо купити ", result, " яблука")
	fmt.Println(msg)
	// 4
	fmt.Println(questionFour)
	result = applePrice*2 + pearPrice*2
	if money >= result {
		fmt.Println("Так")
	} else {
		fmt.Println("Ні")
	}

}
