package dialog

import (
	"Chiffre_de_Vigenere/internal/cipher"
	"Chiffre_de_Vigenere/internal/variables"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Dialog() {
	var option string
	fmt.Println("1) Зашифровать сообщение шифром Виженера\n2) Расшифровать сообщение, зашифрованное шифром Виженера")
	fmt.Scan(&option)

	switch option {
	case "1":
		fmt.Println("Выберите язык для шифровки (RU/EN)")
		lang := Scan()

		fmt.Println("Введите сообщение, которое необходимо зашифровать...")
		message := Scan()
		fmt.Println("Введите ключ для шифровки сообщения...")
		message_key := Scan()

		if strings.ToLower(lang) == "ru" {
			finalMessage := cipher.ExecuteCipher("encrypt", lang, []rune(message), message_key, variables.Alphabet_RU)
			fmt.Println(finalMessage)
		} else if strings.ToLower(lang) == "en" {
			finalMessage := cipher.ExecuteCipher("encrypt", lang, []rune(message), message_key, variables.Alphabet_EN)
			fmt.Println(finalMessage)
		} else {
			fmt.Println("Неверный ввод языка для шифровки!")
		}

		Dialog()
	case "2":
		fmt.Println("Выберите язык для шифровки (RU/EN)")
		lang := Scan()

		fmt.Println("Введите сообщение, которое необходимо расшифровать...")
		message := Scan()
		fmt.Println("Введите ключ для дешифровки сообщения...")
		message_key := Scan()

		if strings.ToLower(lang) == "ru" {
			finalMessage := cipher.ExecuteCipher("decrypt", lang, []rune(message), message_key, variables.Alphabet_RU)
			fmt.Println(finalMessage)
		} else if strings.ToLower(lang) == "en" {
			finalMessage := cipher.ExecuteCipher("decrypt", lang, []rune(message), message_key, variables.Alphabet_EN)
			fmt.Println(finalMessage)
		} else {
			fmt.Println("Неверный ввод языка для шифровки!")
		}

		Dialog()
	default:
		fmt.Println("Неверный ввод!")
		Dialog()
	}
}

func Scan() string {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	if err := in.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка ввода:", err)
	}
	return in.Text()
}
