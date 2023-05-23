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
	Ukr         string
	PartOfSpeech  string
	Synonyms      string // Изменено на string
	Rating        int
}

type DictionaryStruct struct {
	Index                                           int
	WordOriginal                                    string
	WordTranslated                                  string
	WordOriginalTranscription                       string
	WordOriginalPastSimpleSingular                  string
	WordOriginalPastSimpleSingularTranscription     string
	WordOriginalPastSimplePlural                    string
	WordOriginalPastSimplePluralTranscription       string
	WordOriginalPastParticipleSingular              string
	WordOriginalPastParticipleSingularTranscription string
	WordOriginalPastParticiplePlural                string
	WordOriginalPastParticiplePluralTranscription   string
	WordOriginalSynonyms                            string
	WordOriginalPartOfSpeech                        string
	Rating                                          int
}
type Dictionary_V2 struct {
	WordIndex                                       int
	WordOriginal                                    string
	WordTranslated                                  string
	WordOriginalTranscription                       string
	WordOriginalPastSimpleSingular                  string
	WordOriginalPastSimpleSingularTranscription     string
	WordOriginalPastSimplePlural                    string
	WordOriginalPastSimplePluralTranscription       string
	WordOriginalPastParticipleSingular              string
	WordOriginalPastParticipleSingularTranscription string
	WordOriginalPastParticiplePlural                string
	WordOriginalPastParticiplePluralTranscription   string
	WordOriginalSynonyms                            string
	WordOriginalPartOfSpeech                        string
	Rating                                          int
}

func main() {
	data, err := readJSON("eng-ukr.json")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	words := parseData(data)
	// Открываем файл для записи
	jsonFile, err := os.OpenFile("eng-ukr_v2.json", os.O_WRONLY|os.O_TRUNC, 0644)
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
		ukr := extractInfo(html)
		words = append(words, Word{
			English: english,
			Ukr:   ukr,
		})
	}

	return words
}

func extractInfo(html string) string {
	ukr := extractUkr(html)
	return ukr
}

func extractUkr(html string) string {
	re := regexp.MustCompile(`<div style="margin-left:2em">(.*?)</div>`)
	matches := re.FindStringSubmatch(html)
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}

func extractEng(html string) string { // Изменено на string
	re := regexp.MustCompile(`<i>\((.*?)\)</i>`)
	matches := re.FindStringSubmatch(html)
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}
