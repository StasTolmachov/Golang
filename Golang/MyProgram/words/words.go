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
	"strconv"
	"strings"
)

type WordsStruct struct {
	Index         int
	English       string
	Transcription string
	Russian       string
	PartOfSpeech  string
	Synonyms      string
	Rating        int
}

var Words = []WordsStruct{}
var GoogleDict = []WordsStruct{}

type IndexData struct {
	Index int `json:"index"`
}

var Word1 WordsStruct
var WordValue WordsStruct

var IndexWord int

type ElementWithIndex struct {
	Index   int
	Element WordsStruct
}

func main() {
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

	jsonFileGoogle, err := os.Open("eng-rus_Google.json")
	if err != nil {
		fmt.Println("Ошибка создания файла:", err)
		return
	}
	defer jsonFile.Close()

	// Читаем содержимое файла
	jsonDataGoogle, err := ioutil.ReadAll(jsonFileGoogle)
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return
	}

	// Десериализуем JSON в структуру

	err = json.Unmarshal(jsonDataGoogle, &GoogleDict)
	if err != nil {
		fmt.Println("Ошибка десериализации:", err)
		return
	}

	log.Println("started http.ListenAndServe localhost:8080/word")
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/index", index)
	http.HandleFunc("/word", word)
	http.HandleFunc("/wordOtvet", wordOtvet)
	http.HandleFunc("/wordAdd", wordAdd)
	http.HandleFunc("/wordAll", wordAll)
	http.HandleFunc("/handleIndex", handleIndex)
	http.HandleFunc("/handleEdit", handleEdit)
	http.HandleFunc("/handleAdd", handleAdd)
	http.HandleFunc("/element-info/", handleElementInfo)
	http.HandleFunc("/wordsSearch", wordsSearch)
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

	// Сортируем список слов по значению Rating в порядке возрастания
	sort.Slice(Words, func(i, j int) bool {
		return Words[i].Rating < Words[j].Rating
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

	IndexWord = findMinRatingIndex(Words)

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

	if r.FormValue("English") != "" {
		// lastIndex := len(Words)
		// WordValue.Index = lastIndex + 1
		WordValue.English = r.FormValue("English")
		WordValue.Transcription = r.FormValue("Transcription")
		WordValue.Russian = r.FormValue("Russian")
		WordValue.PartOfSpeech = r.FormValue("PartOfSpeech")
		WordValue.Synonyms = r.FormValue("Synonyms")

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

	err = tmpl.Execute(w, GoogleDict)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func wordOtvet(w http.ResponseWriter, r *http.Request) {
	WordValue.English = r.FormValue("word")

	if WordValue.English == Words[IndexWord].English {
		Words[IndexWord].Rating += 1
		fmt.Println(Words[IndexWord].Rating)

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
		Words[IndexWord].Rating -= 1
		fmt.Println(Word1.Rating)
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

func findMinRatingIndex(words []WordsStruct) int {
	minIndex := 0
	minValue := words[0].Rating

	for i, word := range words {
		if word.Rating < minValue {
			minValue = word.Rating
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

func handleEdit(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var requestData struct {
		Index         int    `json:"index"`
		English       string `json:"English"`
		Transcription string `json:"Transcription"`
		Russian       string `json:"Russian"`

		PartOfSpeech string `json:"PartOfSpeech"`
		Synonyms     string `json:"Synonyms"`
		Rating       int    `json:"Rating"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	index := requestData.Index
	if index < 0 || index >= len(Words) {
		http.Error(w, "Invalid index value", http.StatusBadRequest)
		return
	}

	// Обновление элемента с новыми данными
	Words[index].English = requestData.English
	Words[index].Transcription = requestData.Transcription
	Words[index].Russian = requestData.Russian
	Words[index].PartOfSpeech = requestData.PartOfSpeech
	Words[index].Synonyms = requestData.Synonyms
	Words[index].Rating = requestData.Rating

	// Обновление файла данных (если есть) и другие операции, если необходимо
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

	w.WriteHeader(http.StatusOK)
}

func handleAdd(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var requestData struct {
		Index         int    `json:"index"`
		English       string `json:"English"`
		Transcription string `json:"Transcription"`
		Russian       string `json:"Russian"`
		PartOfSpeech  string `json:"PartOfSpeech"`
		Synonyms      string `json:"Synonyms"`
		Rating        int    `json:"Rating"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	index := requestData.Index
	if index < 0 || index >= len(Words) {
		http.Error(w, "Invalid index value", http.StatusBadRequest)
		return
	}

	Words = append(Words, WordsStruct(requestData))

	// Обновление файла данных (если есть) и другие операции, если необходимо
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

	w.WriteHeader(http.StatusOK)
}

func handleElementInfo(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("template/ElementInfo.html", "template/header.html", "template/footer.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	urlPath := r.URL.Path
	indexStr := strings.TrimPrefix(urlPath, "/element-info/")
	index, err := strconv.Atoi(indexStr)
	if err != nil {
		http.Error(w, "Invalid index format", http.StatusBadRequest)
		return
	}

	if index < 0 || index >= len(Words) {
		http.Error(w, "Index out of range", http.StatusBadRequest)
		return
	}

	elementWithIndex := ElementWithIndex{
		Index:   index,
		Element: Words[index],
	}

	err = tmpl.Execute(w, elementWithIndex)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	results := searchWords(query)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

func searchWords(query string) []WordsStruct {
	results := []WordsStruct{}
	query = strings.ToLower(query)

	for _, word := range Words {
		if strings.Contains(strings.ToLower(word.English), query) {
			results = append(results, word)
		}
	}

	return results
}
func wordsSearch(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(GoogleDict)
}
