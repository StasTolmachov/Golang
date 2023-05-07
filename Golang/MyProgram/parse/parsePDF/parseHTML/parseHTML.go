package main

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type DictionaryStruct struct {
	WordIndex                                       int    //default
	WordOriginal                                    string //<div class="custom-table-col" data-label="Infinitive:">be [bi]</div> выбрать нужно только значение: be
	WordTranslated                                  string //<div class="custom-table-col" data-label="Перевод:">быть</div> выбрать нужно только значение: Перевод
	WordOriginalTranscription                       string //<div class="custom-table-col" data-label="Infinitive:">be [bi]</div> выбрать нужно только значение: bi
	WordOriginalPresentSimpleSingularFirst          string //default
	WordOriginalPresentSimpleSingularThird          string //default
	WordOriginalPresentSimplePlural                 string //default
	WordOriginalPresentSimpleParticiple             string //default
	WordOriginalPastSimpleSingular                  string //<div class="custom-table-col" data-label="Past Simple:">was [wɒz], were [wɜː]</div> выбрать нужно только значение: was
	WordOriginalPastSimpleSingularTranscription     string //<div class="custom-table-col" data-label="Past Simple:">was [wɒz], were [wɜː]</div> выбрать нужно только значение: wɒz
	WordOriginalPastSimplePlural                    string //<div class="custom-table-col" data-label="Past Simple:">was [wɒz], were [wɜː]</div> выбрать нужно только значение: were
	WordOriginalPastSimplePluralTranscription       string //<div class="custom-table-col" data-label="Past Simple:">was [wɒz], were [wɜː]</div> выбрать нужно только значение: wɜː
	WordOriginalPastParticipleSingular              string //<div class="custom-table-col" data-label="Past Participle:">born [bɔːn],borne [bɔːn]</div> выбрать нужно только значение первое слово: born
	WordOriginalPastParticipleSingularTranscription string //<div class="custom-table-col" data-label="Past Participle:">born [bɔːn],borne [bɔːn]</div> выбрать нужно только значение транскрипция к первому слову: bɔːn
	WordOriginalPastParticiplePlural                string //<div class="custom-table-col" data-label="Past Participle:">born [bɔːn],borne [bɔːn]</div> выбрать нужно только значение второе слово: borne
	WordOriginalPastParticiplePluralTranscription   string //<div class="custom-table-col" data-label="Past Participle:">born [bɔːn],borne [bɔːn]</div> выбрать нужно только значение транскрипция ко второму слову: bɔːn
	WordOriginalSynonyms                            string //default
	WordOriginalPartOfSpeech                        string //default
	Rating                                          int    //default
}

func main() {
	file, err := os.Open("verb.html")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	doc, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		panic(err)
	}

	wordList := make([]DictionaryStruct, 0)

	doc.Find(".custom-table-row").Each(func(i int, row *goquery.Selection) {
		var word DictionaryStruct
		word.WordIndex = i
		infinitive, _ := row.Find("div[data-label='Infinitive:']").Html()
		pastSimple, _ := row.Find("div[data-label='Past Simple:']").Html()
		pastParticiple, _ := row.Find("div[data-label='Past Participle:']").Html()
		translation, _ := row.Find("div[data-label='Перевод:']").Html()

		infinitiveParts := extractParts(infinitive)
		pastSimpleParts := extractParts(pastSimple)
		pastParticipleParts := extractParts(pastParticiple)

		word.WordOriginal = infinitiveParts[0]
		word.WordOriginalTranscription = infinitiveParts[1]

		word.WordOriginalPastSimpleSingular = pastSimpleParts[0]
		word.WordOriginalPastSimpleSingularTranscription = pastSimpleParts[1]
		word.WordOriginalPastSimplePlural = pastSimpleParts[2]
		word.WordOriginalPastSimplePluralTranscription = pastSimpleParts[3]

		word.WordOriginalPastParticipleSingular = pastParticipleParts[0]
		word.WordOriginalPastParticipleSingularTranscription = pastParticipleParts[1]
		word.WordOriginalPastParticiplePlural = pastParticipleParts[2]
		word.WordOriginalPastParticiplePluralTranscription = pastParticipleParts[3]

		word.WordTranslated = translation

		wordList = append(wordList, word)
	})

	// Открываем файл для записи
	jsonFile, err := os.OpenFile("verbs.json", os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return
	}
	defer jsonFile.Close()
	// Сериализуем структуру в JSON
	jsonData, err := json.MarshalIndent(wordList, "", "  ")
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

func extractParts(input string) []string {
	re := regexp.MustCompile(`\[.*?\]`)
	matches := re.FindAllString(input, -1)

	for i, match := range matches {
		matches[i] = strings.Trim(match, "[]")
		input = strings.Replace(input, match, "", 1)
	}

	words := strings.Split(strings.TrimSpace(input), ",")

	for i, word := range words {
		words[i] = strings.TrimSpace(word)
	}

	// Расширяем списки слов и транскрипций до длины 4
	for len(words) < 4 {
		words = append(words, "")
	}

	for len(matches) < 4 {
		matches = append(matches, "")
	}

	// Объединяем списки слов и транскрипций поочередно
	result := make([]string, 0, 8)
	for i := 0; i < 4; i++ {
		result = append(result, words[i], matches[i])
	}

	return result
}
