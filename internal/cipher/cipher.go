package cipher

import (
	"Chiffre_de_Vigenere/internal/variables"
	"regexp"
	"strings"
	"unicode"
)

func formatMessageKey(message_key string) []rune {
	message_key = strings.ToLower(message_key)
	message_key = strings.ReplaceAll(message_key, " ", "")

	return []rune(message_key)
}

// Parameters: ACTION stands for "decrypt" or "encrypt"; LANG is a language of message and message key;
// MESSAGE and _MESSAGE_KEY stands for the message you want to encrypt/decrypt and for the key for encryption/decryption respectively;
// ALPHABET is a map[string]int where key is a letter of an alphabet and value is a letter number;
//
// This function returns the result of the encryption/decryption of given message using given key.
func ExecuteCipher(action, lang string, message []rune, _message_key string, alphabet map[string]int) string {
	var (
		encrypted_message  string
		symbol             rune
		num1, num2, newNum int
		specials           *regexp.Regexp
		isUpper            bool = false
	)

	message_key := formatMessageKey(_message_key)
	if len(message) > len(message_key) {
		for i := 0; i < len(message_key); i++ {
			if len(message) == len(message_key) {
				break
			}
			message_key = append(message_key, message_key[i])
		}
	}

	switch lang {
	case "en":
		specials, _ = regexp.Compile(variables.EngRegularExp)
	case "ru":
		specials, _ = regexp.Compile(variables.RuRegularExp)
	}

	j := 0
	for i := 0; i < len(message); i++ {

		if unicode.IsUpper(message[i]) {
			isUpper = true
			symbol = unicode.ToLower(message[i])
		} else {
			symbol = message[i]
		}

		if match := specials.MatchString(string(symbol)); !match {
			if isUpper {
				symbol = unicode.ToUpper(symbol)
				isUpper = false
			}
			encrypted_message += string(symbol)
			// log.Printf("SKIP \tLetter №%d; \tSymbol: %c;", i, symbol)

			continue
		} else {
			for key, value := range alphabet {
				if string(symbol) == key {
					num1 = value
				}
				if string(message_key[j]) == key {
					num2 = value
				}
			}
			j++

			switch action {
			case "encrypt":
				newNum = num1 + num2
				if newNum >= len(alphabet) {
					newNum -= len(alphabet)
				}
			case "decrypt":
				newNum = num1 - num2
				if newNum < 0 {
					newNum += len(alphabet)
				}
			}

			for key, value := range alphabet {
				if newNum == value {
					if isUpper {
						key = strings.ToUpper(key)
						isUpper = false
					}
					encrypted_message += key
					// log.Printf("NOSKIP \tLetter №%d; \tSymbol: %c; \tMessageNum: %d; \tKeyNum: %d; \tFinalNumber: %d; \tFinalLetter: %s", i, symbol, num1, num2, newNum, key)
				}
			}
		}
	}

	return encrypted_message
}
