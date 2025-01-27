package avito

import (
	"math"
	"sort"
)

type UserStat struct {
	UserId int
	Steps  int
}

func FindChampions(statistics [][]UserStat) map[string]interface{} {
	// Карта для суммирования шагов по userId
	stepMap := make(map[int]int)

	// Суммируем шаги
	for _, dayStats := range statistics {
		for _, stat := range dayStats {
			stepMap[stat.UserId] += stat.Steps
		}
	}

	// Находим максимальное количество шагов
	maxSteps := 0
	for _, steps := range stepMap {
		if steps > maxSteps {
			maxSteps = steps
		}
	}

	// Находим всех пользователей с максимальными шагами
	var userIds []int
	for userId, steps := range stepMap {
		if steps == maxSteps {
			userIds = append(userIds, userId)
		}
	}

	// Возвращаем результат
	return map[string]interface{}{
		"userIds": userIds,
		"steps":   maxSteps,
	}
}

func CalculateDissatisfaction(goods []int, buyerNeeds []int) int {
	// Сортируем массив товаров
	sort.Ints(goods)

	totalDissatisfaction := 0

	// Для каждого покупателя ищем ближайший товар
	for _, need := range buyerNeeds {
		// Используем двоичный поиск для нахождения ближайшего товара
		idx := sort.Search(len(goods), func(i int) bool {
			return goods[i] >= need
		})

		// Определяем неудовлетворённость
		dissatisfaction := math.MaxInt32
		if idx < len(goods) {
			// Если товар больше или равен нужде, проверяем разницу
			dissatisfaction = abs(goods[idx] - need)
		}
		if idx > 0 {
			// Если есть товар меньше нужды, проверяем его разницу
			dissatisfaction = min(dissatisfaction, abs(goods[idx-1]-need))
		}

		// Суммируем неудовлетворённость
		totalDissatisfaction += dissatisfaction
	}

	return totalDissatisfaction
}

// Вспомогательная функция для нахождения модуля числа
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Вспомогательная функция для нахождения минимального числа
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
