package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/tealeg/xlsx"
)

type DictionaryEntry struct {
	English       string
	Transcription string
	Russian       string
	PartOfSpeech  string
}

func main() {
	// Замените "path/to/your/file.xls" на путь к вашему файлу
	dictionary, err := readDictionary("Dict/eng-ru_basic_3000.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	// Вывести прочитанные значения словаря
	// for _, entry := range dictionary {
	// 	fmt.Println(entry)
	// }

	// Открываем файл для записи
	jsonFile, err := os.OpenFile("words.json", os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return
	}
	defer jsonFile.Close()

	// Сериализуем структуру в JSON
	jsonData, err := json.MarshalIndent(dictionary, "", "  ")
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

func readDictionary(filename string) ([]DictionaryEntry, error) {
	xlFile, err := xlsx.OpenFile(filename)
	if err != nil {
		return nil, err
	}

	var dictionary []DictionaryEntry

	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			if len(row.Cells) >= 5 {
				entry := DictionaryEntry{
					English:       row.Cells[1].String(),
					Russian:       row.Cells[3].String(),
					Transcription: row.Cells[2].String(),
					PartOfSpeech:  row.Cells[0].String(),
				}
				dictionary = append(dictionary, entry)
			}
		}
	}

	return dictionary, nil
}
