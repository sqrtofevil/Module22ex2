package main

import (
	"fmt"
	"math/rand"
	"time"
)

const elem = 12

func main() {
	var array [elem]int
	var tofind int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < elem; i++ {
		array[i] = i*2 + rand.Intn(5)
	}

	fmt.Println("Массив на входе:", array)
	fmt.Println("Введите переменую для поиска:")
	fmt.Scan(&tofind)
	res, res_ := searchfirst(array, tofind)
	if !res_ {
		fmt.Println("Переменная не найдена в массиве")
	} else {
		fmt.Println("Переменная найдена в индексе", res)
	}
}

func searchfirst(array [elem]int, tofind int) (int, bool) {
	var index int
	result := false
	for i, value := range array {
		if value > tofind {
			break
		} //если число больше искомого, то нет смысла продолжать поиск по отсортированному массиву
		//сложность алгоритма в лучшем случае O(2) - тогда поиск остановится на втором элементе
		//сложность алгоритма в худшем случае О(n) - тогда нужно пройти весь массив до конца
		if value == tofind {
			index = i
			result = true
			break
		}
	}
	return index, result
}
