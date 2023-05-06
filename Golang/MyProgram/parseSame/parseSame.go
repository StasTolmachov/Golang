package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type DictionaryEntry struct {
	Index         int
	English       string
	Transcription string
	Russian       string
	PartOfSpeech  string
	Synonyms      string
	Rating        int
}

func main() {
	// Здесь предполагается, что у вас есть два словаря: firstDict и secondDict
	var firstDict, secondDict []DictionaryEntry

	// Заполните firstDict и secondDict данными
	// ...
	//  открываем файл first
	jsonFile, err := os.Open("eng-rus_Google.json")
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
	jsonFile2, err := os.Open("3000.json")
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
			if entry1.English == entry2.English {
				firstDict[i].Transcription = entry2.Transcription
				break
			}
		}
	}

	// Выведите обновленный словарь firstDict
	// 	for _, entry := range firstDict {
	// 		fmt.Println(entry)
	// 	}
	// Открываем файл для записи
	jsonFile, err = os.OpenFile("eng-rus_Google.json", os.O_WRONLY|os.O_TRUNC, 0644)
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
