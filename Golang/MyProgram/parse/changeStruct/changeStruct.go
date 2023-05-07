package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type DictionaryOld struct {
	Index         int
	English       string
	Transcription string
	Russian       string
	PartOfSpeech  string
	Synonyms      string
	Rating        int
}

type DictionaryNew struct {
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
	var dictOld []DictionaryOld
	var dictNew []DictionaryNew

	//  открываем файл Old
	jsonFile, err := os.Open("words.json")
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

	err = json.Unmarshal(jsonData, &dictOld)
	if err != nil {
		fmt.Println("Ошибка десериализации:", err)
		return
	}

	//  открываем файл New
	jsonFile2, err := os.Open("EnglishForEveryone.json")
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

	err = json.Unmarshal(jsonData2, &dictNew)
	if err != nil {
		fmt.Println("Ошибка десериализации:", err)
		return
	}

	// логика приложения здесь

	for _, entryOld := range dictOld {
		entryNew := DictionaryNew{
			WordIndex:                      0,
			WordOriginal:                   entryOld.English,
			WordTranslated:                 entryOld.Russian,
			WordOriginalTranscription:      entryOld.Transcription,
			WordOriginalPastSimpleSingular: "",
			WordOriginalPastSimpleSingularTranscription:     "",
			WordOriginalPastSimplePlural:                    "",
			WordOriginalPastSimplePluralTranscription:       "",
			WordOriginalPastParticipleSingular:              "",
			WordOriginalPastParticipleSingularTranscription: "",
			WordOriginalPastParticiplePlural:                "",
			WordOriginalPastParticiplePluralTranscription:   "",
			WordOriginalSynonyms:                            "",
			WordOriginalPartOfSpeech:                        entryOld.PartOfSpeech,
			Rating:                                          0,
		}

		dictNew = append(dictNew, entryNew)
	}

	// Открываем файл для записи
	jsonFile, err = os.OpenFile("EnglishForEveryone.json", os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return
	}
	defer jsonFile.Close()
	// Сериализуем структуру в JSON
	jsonData, err = json.MarshalIndent(dictNew, "", "  ")
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
