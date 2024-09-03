package dialog

import (
	"Chiffre_de_Vigenere/internal/cipher"
	"Chiffre_de_Vigenere/internal/variables"
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func Dialog() {
	for {
		var option string
		fmt.Println("[1] Зашифровать сообщение шифром Виженера\n[2] Расшифровать сообщение, зашифрованное шифром Виженера\n[0] Очистить терминал")
		fmt.Scan(&option)

		switch option {
		case "1":
			lang := chooseLang()

			fmt.Println("Введите сообщение, которое необходимо зашифровать...")
			message := scan()
			fmt.Println("Введите ключ для шифровки сообщения...")
			message_key := scan()

			if strings.ToLower(lang) == "ru" {
				finalMessage := cipher.ExecuteCipher("encrypt", lang, []rune(message), message_key, variables.Alphabet_RU)
				fmt.Println(finalMessage)
			} else if strings.ToLower(lang) == "en" {
				finalMessage := cipher.ExecuteCipher("encrypt", lang, []rune(message), message_key, variables.Alphabet_EN)
				fmt.Println(finalMessage)
			} else {
				fmt.Println("Неверный ввод языка для шифровки!")
			}

		case "2":
			lang := chooseLang()

			fmt.Println("Введите сообщение, которое необходимо расшифровать...")
			message := scan()
			fmt.Println("Введите ключ для дешифровки сообщения...")
			message_key := scan()

			if strings.ToLower(lang) == "ru" {
				finalMessage := cipher.ExecuteCipher("decrypt", lang, []rune(message), message_key, variables.Alphabet_RU)
				fmt.Println(finalMessage)
			} else if strings.ToLower(lang) == "en" {
				finalMessage := cipher.ExecuteCipher("decrypt", lang, []rune(message), message_key, variables.Alphabet_EN)
				fmt.Println(finalMessage)
			} else {
				fmt.Println("Неверный ввод языка для шифровки!")
			}

		case "0":
			clearTerminal()
		default:
			fmt.Println("Неверный ввод!")
		}
	}
}

func chooseLang() string {
	fmt.Println("Выберите язык для шифровки (RU/EN)")
	var lang string
	for {
		lang = scan()
		if strings.ToLower(lang) == "ru" || strings.ToLower(lang) == "en" {
			break
		}
		fmt.Println("неверный ввод (введите 'ru' или 'en')")
	}
	return lang
}

func clearTerminal() {
	if value, ok := variables.ClearOS[runtime.GOOS]; ok {
		value()
	}
}

func init() {
	variables.ClearOS = make(map[string]func())
	variables.ClearOS["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	variables.ClearOS["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	variables.ClearOS["darwin"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func scan() string {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	if err := in.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка ввода:", err)
	}
	return in.Text()
}
