package main

/*
Применимость
- Когда есть клиент, ожидающий один интерфейс, а у нас есть несовместимый по сигнатурам/типам поставщик (SDK, легаси, и т.д.).
- Когда нужно интегрировать сторонний код без переписывания клиентской логики.
Плюсы
- Зависимости: клиент работает с нашим интерфейсом, адаптер меняется при замене провайдера.
- Локализация преобразований: конвертация типов, единиц, протоколов сосредоточена в одном месте.
Минусы
- Дополнительный слой абстракции.
- Накладные расходы, например на маршалинг/анмаршалинг данных.
Реальные примеры:
1) http.HandlerFunc:
- Сервер ждёт интерфейс http.Handler, а у нас есть обычная функция.
- http.HandlerFunc оборачивает функцию и делает её совместимой с http.Handler.
2) strings.NewReader
- У нас есть строка, но функция требует io.Reader.
- strings.NewReader превращает строку в объект, реализующий io.Reader.
3) bytes.NewBuffer
- У нас есть срез байтов, а нужен io.Reader или io.Writer.
- bytes.NewBuffer делает адаптер, который реализует оба интерфейса (io.Reader и io.Writer).
*/

import "fmt"

// Клиент ожидает такой интерфейс
type Printer interface {
	Print(text string)
}

// Легаси-класс со своим методом
type legacyPrinter struct{}

func (p legacyPrinter) legacyPrint(isPrint bool, a ...interface{}) {
	if isPrint {
		fmt.Println(a...)
	} else {
		fmt.Println("Сlient does not want to print")
	}
}

// Адаптер: реализует Printer и внутри вызывает LegacyPrinter
type PrinterAdapter struct {
	legacy *legacyPrinter
}

func (a *PrinterAdapter) Print(text string) {
	// Хардкодим, что мы, например, всегда согласны печатать, т.е. isPrint = true.
	// Имитирует доп. логику, которая нужна старому методу.
	a.legacy.legacyPrint(true, text)
}

func main() {
	// Клиентский код ждёт Printer
	var p Printer

	// Но у нас есть только LegacyPrinter
	legacy := &legacyPrinter{}

	// Оборачиваем его адаптером
	p = &PrinterAdapter{legacy: legacy}

	// И используем как обычный Printer
	p.Print("Hello, World!")
}
