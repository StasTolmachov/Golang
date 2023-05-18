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
type Dictionary_V5 struct {
	Index                     int
	WordOriginal              string
	WordOriginalTranscription string
	WordTranslated            string
	WordOriginalSynonyms      string
	WordOriginalPartOfSpeech  string

	WordOriginalPastSimpleSingular                  string
	WordOriginalPastSimpleSingularTranscription     string
	WordOriginalPastSimplePlural                    string
	WordOriginalPastSimplePluralTranscription       string
	WordOriginalPastParticipleSingular              string
	WordOriginalPastParticipleSingularTranscription string
	WordOriginalPastParticiplePlural                string
	WordOriginalPastParticiplePluralTranscription   string

	WordOriginalCounterAttempts  int
	WordOriginalCounterIncorrect int
	WordOriginalCounterCorrect   int
	WordOriginalDifficultyRating int
	WordOriginalStatus           string
	WordOriginalDictionary       []string // в каких словарях добавленно
}

func main() {
	var dictOld []Dictionary_V2
	var dictNew []Dictionary_V5

	//  открываем файл Old
	jsonFile, err := os.Open("eng-rus_Google_v2.json")
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
	jsonFile2, err := os.Open("eng-rus_Google_v5.json")
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
		entryNew := Dictionary_V5{
			Index:                     0,
			WordOriginal:              entryOld.WordOriginal,
			WordOriginalTranscription: entryOld.WordOriginalTranscription,
			WordTranslated:            entryOld.WordTranslated,
			WordOriginalSynonyms:      entryOld.WordOriginalSynonyms,
			WordOriginalPartOfSpeech:  entryOld.WordOriginalPartOfSpeech,

			WordOriginalPastSimpleSingular:                  entryOld.WordOriginalPastSimpleSingular,
			WordOriginalPastSimpleSingularTranscription:     entryOld.WordOriginalPastSimpleSingularTranscription,
			WordOriginalPastSimplePlural:                    entryOld.WordOriginalPastSimplePlural,
			WordOriginalPastSimplePluralTranscription:       entryOld.WordOriginalPastSimplePluralTranscription,
			WordOriginalPastParticipleSingular:              entryOld.WordOriginalPastParticipleSingular,
			WordOriginalPastParticipleSingularTranscription: entryOld.WordOriginalPastParticipleSingularTranscription,
			WordOriginalPastParticiplePlural:                entryOld.WordOriginalPastParticiplePlural,
			WordOriginalPastParticiplePluralTranscription:   entryOld.WordOriginalPastParticiplePluralTranscription,

			WordOriginalCounterAttempts:  0,
			WordOriginalCounterIncorrect: 0,
			WordOriginalCounterCorrect:   0,
			WordOriginalDifficultyRating: 0,
			WordOriginalStatus:           "",
			WordOriginalDictionary:       []string{},
		}

		dictNew = append(dictNew, entryNew)
	}

	// Открываем файл для записи
	jsonFile, err = os.OpenFile("eng-rus_Google_v5.json", os.O_WRONLY|os.O_TRUNC, 0644)
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
