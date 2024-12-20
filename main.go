package main

import (
	"fmt"
	"strings"
	"time"
)

var morseCode = map[string]map[rune]string{
	"en": {
		'A': ".-", 'B': "-...", 'C': "-.-.", 'D': "-..", 'E': ".",
		'F': "..-.", 'G': "--.", 'H': "....", 'I': "..", 'J': ".---",
		'K': "-.-", 'L': ".-..", 'M': "--", 'N': "-.", 'O': "---",
		'P': ".--.", 'Q': "--.-", 'R': ".-.", 'S': "...", 'T': "-",
		'U': "..-", 'V': "...-", 'W': ".--", 'X': "-..-", 'Y': "-.--",
		'Z': "--..", '1': ".----", '2': "..---", '3': "...--", '4': "....-", '5': ".....",
		'6': "-....", '7': "--...", '8': "---..", '9': "----.", '0': "-----", ' ': "/",
	},
	"ru": {
		'А': ".-", 'Б': "-...", 'В': ".--", 'Г': "--.", 'Д': "-..",
		'Е': ".", 'Ж': "...-", 'З': "--..", 'И': "..", 'Й': ".---",
		'К': "-.-", 'Л': ".-..", 'М': "--", 'Н': "-.", 'О': "---",
		'П': ".--.", 'Р': ".-.", 'С': "...", 'Т': "-", 'У': "..-",
		'Ф': "..-.", 'Х': "....", 'Ц': "-.-.", 'Ч': "---.", 'Ш': "----",
		'Щ': "--.-", 'Ы': "-.--", 'Э': "..-..", 'Ю': "..--", 'Я': ".-.-",
		'1': ".----", '2': "..---", '3': "...--", '4': "....-", '5': ".....",
		'6': "-....", '7': "--...", '8': "---..", '9': "----.", '0': "-----", ' ': "/",
	},
	"de": {
		'A': ".-", 'B': "-...", 'C': "-.-.", 'D': "-..", 'E': ".",
		'F': "..-.", 'G': "--.", 'H': "....", 'I': "..", 'J': ".---",
		'K': "-.-", 'L': ".-..", 'M': "--", 'N': "-.", 'O': "---",
		'P': ".--.", 'Q': "--.-", 'R': ".-.", 'S': "...", 'T': "-",
		'U': "..-", 'V': "...-", 'W': ".--", 'X': "-..-", 'Y': "-.--",
		'Z': "--..", '1': ".----", '2': "..---", '3': "...--", '4': "....-", '5': ".....",
		'6': "-....", '7': "--...", '8': "---..", '9': "----.", '0': "-----", ' ': "/",
	},
	"fr": {
		'A': ".-", 'B': "-...", 'C': "-.-.", 'D': "-..", 'E': ".",
		'F': "..-.", 'G': "--.", 'H': "....", 'I': "..", 'J': ".---",
		'K': "-.-", 'L': ".-..", 'M': "--", 'N': "-.", 'O': "---",
		'P': ".--.", 'Q': "--.-", 'R': ".-.", 'S': "...", 'T': "-",
		'U': "..-", 'V': "...-", 'W': ".--", 'X': "-..-", 'Y': "-.--",
		'Z': "--..", '1': ".----", '2': "..---", '3': "...--", '4': "....-", '5': ".....",
		'6': "-....", '7': "--...", '8': "---..", '9': "----.", '0': "-----", ' ': "/",
	},
}

var localization = map[string]map[string]string{
	"en": {
		"choose_language": "Choose your language:",
		"action_menu":     "Choose an action:",
		"encrypt":         "1. Encrypt message",
		"decrypt":         "2. Decrypt message",
		"time_in_morse":   "3. Get current time in Morse code",
		"exit":            "4. Exit",
		"enter_text":      "Enter text to encrypt:",
		"enter_morse":     "Enter Morse code to decrypt (use spaces between letters and '/' between words):",
		"encrypted":       "Encrypted Morse:",
		"decrypted":       "Decrypted text:",
		"current_time":    "Current time in Morse code:",
		"goodbye":         "Goodbye!",
		"continue":        "Do you want to continue? (yes/no):",
	},
	"ru": {
		"choose_language": "Выберите язык:",
		"action_menu":     "Выберите действие:",
		"encrypt":         "1. Зашифровать сообщение",
		"decrypt":         "2. Расшифровать сообщение",
		"time_in_morse":   "3. Узнать текущее время в азбуке Морзе",
		"exit":            "4. Выход",
		"enter_text":      "Введите текст для шифрования:",
		"enter_morse":     "Введите код Морзе для расшифровки (используйте пробелы между буквами и '/' между словами):",
		"encrypted":       "Зашифрованный Морзе:",
		"decrypted":       "Расшифрованный текст:",
		"current_time":    "Текущее время в азбуке Морзе:",
		"goodbye":         "До свидания!",
		"continue":        "Хотите продолжить? (да/нет):",
	},
	"de": {
		"choose_language": "Wählen Sie Ihre Sprache:",
		"action_menu":     "Wählen Sie eine Aktion:",
		"encrypt":         "1. Nachricht verschlüsseln",
		"decrypt":         "2. Nachricht entschlüsseln",
		"time_in_morse":   "3. Aktuelle Zeit im Morsecode anzeigen",
		"exit":            "4. Beenden",
		"enter_text":      "Geben Sie den Text ein, der verschlüsselt werden soll:",
		"enter_morse":     "Geben Sie den Morsecode ein, der entschlüsselt werden soll (verwenden Sie Leerzeichen zwischen Buchstaben und '/' zwischen Wörtern):",
		"encrypted":       "Verschlüsselter Morsecode:",
		"decrypted":       "Entschlüsselter Text:",
		"current_time":    "Aktuelle Zeit im Morsecode:",
		"goodbye":         "Auf Wiedersehen!",
		"continue":        "Möchten Sie fortfahren? (ja/nein):",
	},
	"fr": {
		"choose_language": "Choisissez votre langue:",
		"action_menu":     "Choisissez une action:",
		"encrypt":         "1. Chiffrer un message",
		"decrypt":         "2. Déchiffrer un message",
		"time_in_morse":   "3. Obtenir l'heure actuelle en code Morse",
		"exit":            "4. Quitter",
		"enter_text":      "Entrez le texte à chiffrer:",
		"enter_morse":     "Entrez le code Morse à déchiffrer (utilisez des espaces entre les lettres et '/' entre les mots):",
		"encrypted":       "Message chiffré en Morse:",
		"decrypted":       "Texte déchiffré:",
		"current_time":    "Heure actuelle en code Morse:",
		"goodbye":         "Au revoir!",
		"continue":        "Voulez-vous continuer? (oui/non):",
	},
}

func toMorse(input string, lang string) string {
	var morseMessage []string
	alphabet, exists := morseCode[lang]
	if !exists {
		return "Language not supported"
	}
	for _, char := range strings.ToUpper(input) {
		if morse, found := alphabet[char]; found {
			morseMessage = append(morseMessage, morse)
		} else if char == ' ' {
			morseMessage = append(morseMessage, "/")
		}
	}
	return strings.Join(morseMessage, " ")
}

func fromMorse(input string, lang string) string {
	alphabet, exists := morseCode[lang]
	if !exists {
		return "Language not supported"
	}
	reverseAlphabet := make(map[string]rune)
	for char, morse := range alphabet {
		reverseAlphabet[morse] = char
	}
	words := strings.Split(input, "/")
	var textMessage []string
	for _, word := range words {
		letters := strings.Split(word, " ")
		for _, letter := range letters {
			if char, found := reverseAlphabet[letter]; found {
				textMessage = append(textMessage, string(char))
			}
		}
		textMessage = append(textMessage, " ")
	}
	return strings.TrimSpace(strings.Join(textMessage, ""))
}

func currentTimeInMorse(lang string) string {
	now := time.Now()
	timeStr := now.Format("15:04:05")
	timeStrWithColon := strings.Replace(timeStr, ":", " : ", -1)
	return toMorse(timeStrWithColon, lang)
}

func main() {
	lang := "en"
	fmt.Println(localization[lang]["choose_language"])
	fmt.Println("1 - English, 2 - Русский, 3 - Deutsch, 4 - Français")
	var choice int
	fmt.Scanln(&choice)
	switch choice {
	case 2:
		lang = "ru"
	case 3:
		lang = "de"
	case 4:
		lang = "fr"
	}

	for {
		fmt.Println(localization[lang]["action_menu"])
		fmt.Println(localization[lang]["encrypt"])
		fmt.Println(localization[lang]["decrypt"])
		fmt.Println(localization[lang]["time_in_morse"])
		fmt.Println(localization[lang]["exit"])

		var action int
		fmt.Scanln(&action)
		switch action {
		case 1:
			fmt.Println(localization[lang]["enter_text"])
			var text string
			fmt.Scanln(&text)
			fmt.Println(localization[lang]["encrypted"], toMorse(text, lang))
		case 2:
			fmt.Println(localization[lang]["enter_morse"])
			var morse string
			fmt.Scanln(&morse)
			fmt.Println(localization[lang]["decrypted"], fromMorse(morse, lang))
		case 3:
			fmt.Println(localization[lang]["current_time"], currentTimeInMorse(lang))
		case 4:
			fmt.Println(localization[lang]["goodbye"])
			return
		}

		fmt.Println(localization[lang]["continue"])
		var cont string
		fmt.Scanln(&cont)
		if strings.ToLower(cont) != "yes" && strings.ToLower(cont) != "да" && strings.ToLower(cont) != "ja" && strings.ToLower(cont) != "oui" {
			fmt.Println(localization[lang]["goodbye"])
			break
		}
	}
}

// Language: go