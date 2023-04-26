package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
)

type WordsStruct struct {
	Rus           string `json: "Rus`
	Norg          string `json: "Norg`
	Transcription string `json: "Transcription`
	True          int    `json: "True`
}

var Words = []WordsStruct{}

type IndexData struct {
	Index int `json:"index"`
}

var Word1 WordsStruct
var WordValue WordsStruct

// var WordTemp WordsStruct
var IndexWord int

func main() {
	// Word1.Rus = "машина2"
	// Word1.Norg = "bil"
	// Word1.Transcription = "transcription"

	// Сериализуем структуру в JSON
	// jsonData, err := json.MarshalIndent(Words, "", "  ")
	// if err != nil {
	// 	fmt.Println("Ошибка сериализации:", err)
	// 	return
	// }

	//  создаем и открываем файл
	// jsonFile, err := os.Create("words.json")
	// if err != nil {
	// 	fmt.Println("Ошибка создания файла:", err)
	// 	return
	// }
	// defer jsonFile.Close()

	//  открываем файл
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

	err = json.Unmarshal(jsonData, &Words)
	if err != nil {
		fmt.Println("Ошибка десериализации:", err)
		return
	}

	// Записываем JSON в файл
	// _, err = jsonFile.Write(jsonData)
	// if err != nil {
	// 	fmt.Println("Ошибка записи в файл:", err)
	// 	return
	// }

	log.Println("started http.ListenAndServe localhost:8080/word")
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/index", index)
	http.HandleFunc("/word", word)
	http.HandleFunc("/wordOtvet", wordOtvet)
	http.HandleFunc("/wordAdd", wordAdd)
	http.HandleFunc("/wordAll", wordAll)
	http.HandleFunc("/handleIndex", handleIndex)
	// http.HandleFunc("/nextWord", nextWord)
	http.ListenAndServe(":8080", nil)

}

func index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("template/index.html", "template/header.html", "template/footer.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func wordAll(w http.ResponseWriter, r *http.Request) {

	// Сортируем список слов по значению True в порядке возрастания
	sort.Slice(Words, func(i, j int) bool {
		return Words[i].True < Words[j].True
	})

	tmpl, err := template.ParseFiles("template/wordAll.html", "template/header.html", "template/footer.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, Words)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func word(w http.ResponseWriter, r *http.Request) {

	// max := findMinTrueIndex(Words)
	// IndexWord = randomInt(max)

	IndexWord = findMinTrueIndex(Words)

	tmpl, err := template.ParseFiles("template/word.html", "template/header.html", "template/footer.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, Words[IndexWord])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
func wordAdd(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("template/wordAdd.html", "template/header.html", "template/footer.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if r.FormValue("Norg") != "" {
		WordValue.Norg = r.FormValue("Norg")
		WordValue.Transcription = r.FormValue("Transcription")
		WordValue.Rus = r.FormValue("Rus")
		Words = append(Words, WordValue)

		// Открываем файл для записи
		jsonFile, err := os.OpenFile("words.json", os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			fmt.Println("Ошибка открытия файла:", err)
			return
		}
		defer jsonFile.Close()
		// Сериализуем структуру в JSON
		jsonData, err := json.MarshalIndent(Words, "", "  ")
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

}
func wordOtvet(w http.ResponseWriter, r *http.Request) {
	WordValue.Norg = r.FormValue("word")

	if WordValue.Norg == Words[IndexWord].Norg {
		Words[IndexWord].True += 1
		fmt.Println(Words[IndexWord].True)

		// Открываем файл для записи
		jsonFile, err := os.OpenFile("words.json", os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			fmt.Println("Ошибка создания файла:", err)
			return
		}
		defer jsonFile.Close()
		// Сериализуем структуру в JSON
		jsonData, err := json.MarshalIndent(Words, "", "  ")
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

		tmpl, err := template.ParseFiles("template/wordOk.html", "template/header.html", "template/footer.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = tmpl.Execute(w, Words[IndexWord])
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	} else {
		Words[IndexWord].True -= 1
		fmt.Println(Word1.True)
		// Открываем файл для записи
		jsonFile, err := os.OpenFile("words.json", os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			fmt.Println("Ошибка создания файла:", err)
			return
		}
		defer jsonFile.Close()
		// Сериализуем структуру в JSON
		jsonData, err := json.MarshalIndent(Words, "", "  ")
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
		tmpl, err := template.ParseFiles("template/wordNot.html", "template/header.html", "template/footer.html")
		err = tmpl.Execute(w, Words[IndexWord])
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	}
}

func findMinTrueIndex(words []WordsStruct) int {
	minIndex := 0
	minValue := words[0].True

	for i, word := range words {
		if word.True < minValue {
			minValue = word.True
			minIndex = i
		}
	}

	return minIndex
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	var indexData IndexData
	err = json.Unmarshal(body, &indexData)
	if err != nil {
		http.Error(w, "Error unmarshalling JSON", http.StatusBadRequest)
		return
	}

	wordsDelete(indexData.Index)

	// responseData := map[string]string{
	// 	"message": "Индекс успешно обработан",
	// }
	// jsonResponse, err := json.Marshal(responseData)
	// if err != nil {
	// 	http.Error(w, "Error marshalling JSON", http.StatusInternalServerError)
	// 	return
	// }

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))

}

func wordsDelete(index int) {
	fmt.Println("Вызвана функция с индексом:", index)
	// Реализуйте вашу логику здесь
	Words = removeElementByIndex(Words, index)
	// Открываем файл для записи
	jsonFile, err := os.OpenFile("words.json", os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return
	}
	defer jsonFile.Close()
	// Сериализуем структуру в JSON
	jsonData, err := json.MarshalIndent(Words, "", "  ")
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
func removeElementByIndex(words []WordsStruct, index int) []WordsStruct {
	if index < 0 || index >= len(words) {
		return words
	}
	return append(words[:index], words[index+1:]...)
}
