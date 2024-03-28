package main

// Call this via HTTP GET with a URL like:
//     http://localhost:9998/translate?lang=fr&term=hello
//
// This will return a JSON-encoded map, with a single key:
// translation (containing the translated term). It currently
// supports the following languages
//
//    de: German
//    es: Spanish
//    fr: French
//    lv: Latvian
//    mi: Maori
//    sk: Slovak
//    tr: Turkish
//    zu: Zulu

import (
	"fmt"
	"net/http"
	"strings"
	"unicode"
)

// Key = language code (ISO 639-1), value = map [key=term, value=translation]
var translations map[string]map[string]string

func init() {

	translations = map[string]map[string]string{
		"de": {
			"hello":   "hallo",
			"goodbye": "auf wiedersehen",
			"thanks":  "danke schön",
		},
		"es": {
			"hello":   "hola",
			"goodbye": "adiós",
			"thanks":  "gracias",
		},
		"fr": {
			"hello":   "bonjour",
			"goodbye": "au revoir",
			"thanks":  "merci",
		},
		"lv": {
			"hello":   "sveiks",
			"goodbye": "ardievu",
			"thanks":  "paldies",
		},
		"mi": {
			"hello":   "kia ora",
			"goodbye": "poroporoaki",
			"thanks":  "whakawhetai koe",
		},
		"sk": {
			"hello":   "ahoj",
			"goodbye": "zbohom",
			"thanks":  "ďakujem koe",
		},
		"tr": {
			"hello":   "merhaba",
			"goodbye": "güle güle",
			"thanks":  "teşekkür ederim",
		},
		"zu": {
			"hello":   "hamba kahle",
			"goodbye": "sawubona",
			"thanks":  "ngiyabonga",
		},
	}
}

// Languages currently supported are French, German, Spanish, Latvian,
// Maori, Slovak, Turkish, and Zulu
func translationHandler(w http.ResponseWriter, r *http.Request) {

	// look up translation for the specified term in the specified language
	langKeys, hasLangParam := r.URL.Query()["lang"]
	if !hasLangParam {
		http.Error(w, "Missing required 'lang' parameter.", http.StatusBadRequest)
		return
	}
	lang := strings.ToLower(langKeys[0])

	var _, isSupportedLangauge = translations[lang]
	if !isSupportedLangauge {
		msg := fmt.Sprintf("Unknown language code '%s'", lang)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	keys, hasTermParam := r.URL.Query()["term"]
	if hasTermParam {
		term := keys[0]
		key := strings.ToLower(term)
		var translation, canTranslate = translations[lang][key]
		if canTranslate {
			// if the phrase had an initial uppercase letter, reflect that in the translation
			firstLetter := term[0:1]
			asRune := []rune(firstLetter)
			if unicode.IsUpper(asRune[0]) {
				translation = strings.ToUpper(translation[0:1]) + translation[1:]
			}
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, string(translation))
		} else {
			msg := fmt.Sprintf("Unable to translate term '%s' to language '%s'", term, lang)
			http.Error(w, msg, http.StatusBadRequest)
		}
	} else {
		http.Error(w, "Missing required 'term' parameter.", http.StatusBadRequest)
	}
}

func main() {
	http.HandleFunc("/translate", translationHandler)
	fmt.Println("Starting server on port 9998")
	http.ListenAndServe(":9998", nil)
}
