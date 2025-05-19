package main

import (
	"fmt"
)

// ANSI escape codes for colors and formatting
const (
	BOLD      = "\033[1m"
	NORMAL    = "\033[0m"
	LGREEN    = "\033[1;32m"
	LYELLOW   = "\033[1;33m"
	LMAGENTA  = "\033[1;35m"
	GRAY      = "\033[0;37m"
)

func clearScreen() {
	fmt.Print("\033[H\033[2J") // ANSI escape code to clear the screen
}

func main() {
	// Clear the screen
	clearScreen()

	// Display the menu
	fmt.Println()
	fmt.Printf("     %s%sМеню DNS323%s\n", BOLD, LGREEN, NORMAL)
	fmt.Println()
	fmt.Printf("%s 1 %s Комманды для удобной работы в telnet %s(Выполнить?)%s\n", LYELLOW, LGREEN, GRAY, NORMAL)
	fmt.Println()
	fmt.Printf("%s 2 %s Пути к папкам & Изменение прав доступа %s(Комманды)%s\n", LYELLOW, LGREEN, GRAY, NORMAL)
	fmt.Println()
	fmt.Printf("%s 3 %s Transmission (%sStart%s, %sStop%s, %sUpgrade%s) %s(Меню)%s\n", LYELLOW, LGREEN, LGREEN, NORMAL, LYELLOW, NORMAL, LMAGENTA, NORMAL, GRAY, NORMAL)
	fmt.Println()
	fmt.Printf("%s 4 %s Копирование (cp & rsync) %s(Комманды)%s\n", LYELLOW, LGREEN, GRAY, NORMAL)
	fmt.Println()
	fmt.Printf("%s 5 %s Создание ссылки на файл или папку %s(Комманды)%s\n", LYELLOW, LGREEN, GRAY, NORMAL)
	fmt.Println()
	fmt.Printf("%s 6 %s Установка из fun-plug & IPKG %s(Комманды)%s\n", LYELLOW, LGREEN, GRAY, NORMAL)
	fmt.Println()
	fmt.Printf("%s 7 %s Показать Трафик (%sn%sload) %s(Выполнить?)%s\n", LYELLOW, LGREEN, LYELLOW, LGREEN, GRAY, NORMAL)
	fmt.Println()
	fmt.Printf("%s 8 %s Диспетчер задач (%sh%st%sop) %s(Выполнить?)%s\n", LYELLOW, LGREEN, LYELLOW, LGREEN, GRAY, GRAY, NORMAL)
	fmt.Println()
	fmt.Printf("%s 9 %s Midnight Commander (%sm%s) %s(Выполнить?)%s\n", LYELLOW, LGREEN, LYELLOW, LGREEN, GRAY, NORMAL)
	fmt.Println()
	fmt.Printf("%s q %s Выход %s\n", LMAGENTA, LGREEN, NORMAL)
	fmt.Println()
	fmt.Println("(Введите пожалуйта номер пункта, чтобы выполнить комманды этого пункта, любой другой ввод, Выход)")
	fmt.Println()
}

