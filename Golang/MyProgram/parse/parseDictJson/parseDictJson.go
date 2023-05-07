package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

type Word struct {
	English       string
	Transcription string
	Russian       string
	PartOfSpeech  string
	Synonyms      string // Изменено на string
	Rating        int
}

func main() {
	data, err := readJSON("eng-rus_GoogleTranslate_1_0-GD.json")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	words := parseData(data)
	// Открываем файл для записи
	jsonFile, err := os.OpenFile("words.json", os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return
	}
	defer jsonFile.Close()
	// Сериализуем структуру в JSON
	jsonData, err := json.MarshalIndent(words, "", "  ")
	if err != nil {
		fmt.Println("Ошибка сериализации:", err)
		return
	}
	// Записываем JSON в файл
	_, err = jsonFile.Write(jsonData)
	if err != nil {
		fmt.Println("Ошибка записи в файл:", err)
		return
	}
}

func readJSON(filename string) (map[string]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var rawData map[string]string
	err = json.Unmarshal(data, &rawData)
	if err != nil {
		return nil, err
	}

	return rawData, nil
}

func parseData(data map[string]string) []Word {
	words := make([]Word, 0, len(data))

	for english, html := range data {
		russian, partOfSpeech, translations := extractInfo(html)
		words = append(words, Word{
			English:      english,
			Russian:      russian,
			PartOfSpeech: partOfSpeech,
			Synonyms:     translations,
		})
	}

	return words
}

func extractInfo(html string) (string, string, string) { // Изменено на string
	russian := extractRussian(html)
	partOfSpeech := extractPartOfSpeech(html)
	translations := extractTranslations(html)

	return russian, partOfSpeech, translations
}

func extractRussian(html string) string {
	re := regexp.MustCompile(`<b>(.*?)</b>`)
	matches := re.FindStringSubmatch(html)
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}

func extractPartOfSpeech(html string) string {
	re := regexp.MustCompile(`<font color="green">(.*?):</font>`)
	matches := re.FindStringSubmatch(html)
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}

func extractTranslations(html string) string { // Изменено на string
	re := regexp.MustCompile(`<i>\((.*?)\)</i>`)
	matches := re.FindStringSubmatch(html)
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}
