package main

import (
	"examples/avito"
	"fmt"
)

func main() {
	// Пример ввода
	goods := []int{8, 3, 5}
	buyerNeeds := []int{5, 6}

	// Вычисляем неудовлетворённость
	result := avito.CalculateDissatisfaction(goods, buyerNeeds)
	fmt.Println("Total Dissatisfaction:", result) // Ожидается: 1
}
