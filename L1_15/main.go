package main

import (
	"fmt"
)

/*
var justString string

func someFunc() {
  v := createHugeString(1 &lt;&lt; 10)
  justString = v[:100]
}

В Go строка представляет собой ссылку на массив байт плюс длина.
Когда делаем v[:100], создаётся новая строка, но она всё равно указывает на тот же массив байт, что и v.
В итоге justString содержит только первые 100 символов, но удерживает в памяти весь большой v.
Если createHugeString вернул очень большую строку, память будет удерживаться, хотя нужна только малая часть.

justString объявлена как var на уровне пакета.
Это делает функцию someFunc небезопасной для конкурентного использования, усложняет тестирование и вообще так делать не круто.
*/

// Возвращает большую строку из слайса байт.
func createHugeString(size int) string {
	return string(make([]byte, size))
}

func someFunc() string {
	v := createHugeString(1 << 10)

	// копируем первые 100 байт в новый слайс
	tmp := make([]byte, 100)
	copy(tmp, v[:100])

	return string(tmp)
}

func main() {
	result := someFunc()
	fmt.Println(len(result)) // 100
}
