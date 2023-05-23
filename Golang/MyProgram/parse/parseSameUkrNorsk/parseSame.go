package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// type DictionaryEntry struct {
// 	Index         int
// 	English       string
// 	Transcription string
// 	Russian       string
// 	PartOfSpeech  string
// 	Synonyms      string
// 	Rating        int
// }

type DictionaryEntry struct {
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
	// Здесь предполагается, что у вас есть два словаря: firstDict и secondDict
	var firstDict, secondDict []DictionaryEntry

	// Заполните firstDict и secondDict данными
	// ...
	//  открываем файл first
	jsonFile, err := os.Open("eng-nor_v2.json")
	if err != nil {
		fmt.Println("Ошибка создания файла:", err)
		return
	}
	defer jsonFile.Close()

	// Читаем содержимое файла
	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return
	}

	// Десериализуем JSON в структуру

	err = json.Unmarshal(jsonData, &firstDict)
	if err != nil {
		fmt.Println("Ошибка десериализации:", err)
		return
	}

	//  открываем файл second
	jsonFile2, err := os.Open("eng-ukr_v2.json")
	if err != nil {
		fmt.Println("Ошибка создания файла:", err)
		return
	}
	defer jsonFile.Close()

	// Читаем содержимое файла
	jsonData2, err := ioutil.ReadAll(jsonFile2)
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return
	}

	// Десериализуем JSON в структуру

	err = json.Unmarshal(jsonData2, &secondDict)
	if err != nil {
		fmt.Println("Ошибка десериализации:", err)
		return
	}

	// Обновите значения Transcription в firstDict, используя данные из secondDict
	for i, entry1 := range firstDict {
		for _, entry2 := range secondDict {
			if entry1.WordOriginal == entry2.WordOriginal {
				firstDict[i].WordOriginal = firstDict[i].WordTranslated

				firstDict[i].WordTranslated = entry2.WordTranslated

				break
			}
		}
	}

	// Выведите обновленный словарь firstDict
	// 	for _, entry := range firstDict {
	// 		fmt.Println(entry)
	// 	}
	// Открываем файл для записи
	jsonFile, err = os.OpenFile("nok-ukr_v2.json", os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return
	}
	defer jsonFile.Close()
	// Сериализуем структуру в JSON
	jsonData, err = json.MarshalIndent(firstDict, "", "  ")
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
