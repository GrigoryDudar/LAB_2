package main

import (
	"fmt"
	"math/rand"
)

func generatePassword(length int, small, large, numbers, symbols bool) string {
	var letters = ""
	if small {
		letters += "abcdefghijklmnopqrstuvwxyz"
	}
	if large {
		letters += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}
	if numbers {
		letters += "0123456789"
	}
	if symbols {
		letters += "!@#$%^&*"
	}

	if letters == "" {
		return "Помилка: виберіть хоча б один тип символів!"
	}

	password := make([]byte, length)
	for i := range password {
		password[i] = letters[rand.Intn(len(letters))]
	}

	return string(password)
}

func main() {
	var length int
	var small, large, numbers, symbols string

	fmt.Print("Введіть довжину пароля: ")
	fmt.Scan(&length)

	fmt.Print("Чи використовувати малі літери? (+/-): ")
	fmt.Scan(&small)

	fmt.Print("Чи використовувати великі літери? (+/-): ")
	fmt.Scan(&large)

	fmt.Print("Чи використовувати цифри? (+/-): ")
	fmt.Scan(&numbers)

	fmt.Print("Чи використовувати спеціальні символи? (+/-): ")
	fmt.Scan(&symbols)

	password := generatePassword(
		length,
		small == "+",
		large == "+",
		numbers == "+",
		symbols == "+",
	)

	fmt.Println("Згенерований пароль:", password)
}
