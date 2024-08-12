package main

import (
	"fmt"
	"strings"
)

// Хеш-функция (простая реализация)
func hash(word string) int {
	sum := 0
	for _, r := range word {
		sum += int(r)
	}
	return sum % 10 // Используйте простое деление по модулю для хеша
}

func main() {
	text := "The quick brown fox jumps over the lazy dog. The lazy dog is brown."

	// Хеш-таблица
	wordCounts := make(map[string]int)

	words := strings.Fields(text)
	for _, word := range words {
		// Вычисляем хеш
		//index := hash(word)
		//fmt.Println(index)
		// Проверяем, есть ли уже слово в хеш-таблице
		if _, ok := wordCounts[word]; ok {
			wordCounts[word]++
		} else {
			wordCounts[word] = 1
		}
	}

	// Выводим результат
	for word, count := range wordCounts {
		fmt.Printf("%s: %d\n", word, count)
	}
}
