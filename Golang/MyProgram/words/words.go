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

var Words = []DictionaryStruct{}
var GoogleDict = []DictionaryStruct{}

type IndexData struct {
	Index int `json:"index"`
}

var Word1 DictionaryStruct
var WordValue DictionaryStruct

var IndexWord int

type ElementWithIndex struct {
	Index   int
	Element DictionaryStruct
}

func main() {
	//  открываем файл
	jsonFile, err := os.Open("EnglishForEveryone.json")
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

	jsonFileGoogle, err := os.Open("eng-rus_Google_v2.json")
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
	http.HandleFunc("/api/search", searchHandler)

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

	if r.FormValue("WordOriginal") != "" {

		WordValue.WordOriginal = r.FormValue("WordOriginal")
		WordValue.WordOriginalTranscription = r.FormValue("WordOriginalTranscription")
		WordValue.WordTranslated = r.FormValue("WordTranslated")
		WordValue.WordOriginalPartOfSpeech = r.FormValue("WordOriginalPartOfSpeech")
		WordValue.WordOriginalSynonyms = r.FormValue("WordOriginalSynonyms")
		WordValue.WordOriginalPastSimpleSingular = r.FormValue("WordOriginalPastSimpleSingular")
		WordValue.WordOriginalPastSimpleSingularTranscription = r.FormValue("WordOriginalPastSimpleSingularTranscription")
		WordValue.WordOriginalPastSimplePlural = r.FormValue("WordOriginalPastSimplePlural")
		WordValue.WordOriginalPastSimplePluralTranscription = r.FormValue("WordOriginalPastSimplePluralTranscription")
		WordValue.WordOriginalPastParticipleSingular = r.FormValue("WordOriginalPastParticipleSingular")
		WordValue.WordOriginalPastParticipleSingularTranscription = r.FormValue("WordOriginalPastParticipleSingularTranscription")
		WordValue.WordOriginalPastParticiplePlural = r.FormValue("WordOriginalPastParticiplePlural")
		WordValue.WordOriginalPastParticiplePluralTranscription = r.FormValue("WordOriginalPastParticiplePluralTranscription")

		wordExists := false
		for _, words := range Words {
			if words.WordOriginal == WordValue.WordOriginal {
				wordExists = true
				break
			}
		}

		if !wordExists {
			Words = append(Words, WordValue)

			// ... (работа с файлом и запись JSON)

			// Открываем файл для записи
			jsonFile, err := os.OpenFile("EnglishForEveryone.json", os.O_WRONLY|os.O_TRUNC, 0644)
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

	err = tmpl.Execute(w, GoogleDict)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func wordOtvet(w http.ResponseWriter, r *http.Request) {
	WordValue.WordOriginal = r.FormValue("word")

	if strings.EqualFold(WordValue.WordOriginal, Words[IndexWord].WordOriginal) {
		Words[IndexWord].Rating += 1
		fmt.Println(Words[IndexWord].Rating)

		// Открываем файл для записи
		jsonFile, err := os.OpenFile("EnglishForEveryone.json", os.O_WRONLY|os.O_TRUNC, 0644)
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
		jsonFile, err := os.OpenFile("EnglishForEveryone.json", os.O_WRONLY|os.O_TRUNC, 0644)
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

func findMinRatingIndex(words []DictionaryStruct) int {
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
	jsonFile, err := os.OpenFile("EnglishForEveryone.json", os.O_WRONLY|os.O_TRUNC, 0644)
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
func removeElementByIndex(words []DictionaryStruct, index int) []DictionaryStruct {
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
		Index                                           int    `json:"index"`
		WordOriginal                                    string `json:"WordOriginal"`
		WordTranslated                                  string `json:"WordTranslated"`
		WordOriginalTranscription                       string `json:"WordOriginalTranscription"`
		WordOriginalPastSimpleSingular                  string `json:"WordOriginalPastSimpleSingular"`
		WordOriginalPastSimpleSingularTranscription     string `json:"WordOriginalPastSimpleSingularTranscription"`
		WordOriginalPastSimplePlural                    string `json:"WordOriginalPastSimplePlural"`
		WordOriginalPastSimplePluralTranscription       string `json:"WordOriginalPastSimplePluralTranscription"`
		WordOriginalPastParticipleSingular              string `json:"WordOriginalPastParticipleSingular"`
		WordOriginalPastParticipleSingularTranscription string `json:"WordOriginalPastParticipleSingularTranscription"`
		WordOriginalPastParticiplePlural                string `json:"WordOriginalPastParticiplePlural"`
		WordOriginalPastParticiplePluralTranscription   string `json:"WordOriginalPastParticiplePluralTranscription"`
		WordOriginalSynonyms                            string `json:"WordOriginalSynonyms"`
		// WordOriginalPartOfSpeech                        string `json:"WordOriginalPartOfSpeech"`
		Rating int `json:"Rating"`
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
	Words[index].WordOriginal = requestData.WordOriginal
	Words[index].WordOriginalTranscription = requestData.WordOriginalTranscription
	Words[index].WordTranslated = requestData.WordTranslated
	// Words[index].WordOriginalPartOfSpeech = requestData.WordOriginalPartOfSpeech
	Words[index].WordOriginalSynonyms = requestData.WordOriginalSynonyms
	Words[index].Rating = requestData.Rating
	Words[index].WordOriginalPastSimpleSingular = requestData.WordOriginalPastSimpleSingular
	Words[index].WordOriginalPastSimpleSingularTranscription = requestData.WordOriginalPastSimpleSingularTranscription
	Words[index].WordOriginalPastSimplePlural = requestData.WordOriginalPastSimplePlural
	Words[index].WordOriginalPastSimplePluralTranscription = requestData.WordOriginalPastSimplePluralTranscription
	Words[index].WordOriginalPastParticipleSingular = requestData.WordOriginalPastParticipleSingular
	Words[index].WordOriginalPastParticipleSingularTranscription = requestData.WordOriginalPastParticipleSingularTranscription
	Words[index].WordOriginalPastParticiplePlural = requestData.WordOriginalPastParticiplePlural
	Words[index].WordOriginalPastParticiplePluralTranscription = requestData.WordOriginalPastParticiplePluralTranscription

	// Обновление файла данных (если есть) и другие операции, если необходимо
	// Открываем файл для записи
	jsonFile, err := os.OpenFile("EnglishForEveryone.json", os.O_WRONLY|os.O_TRUNC, 0644)
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
		Index                                           int    `json:"index"`
		WordOriginal                                    string `json:"WordOriginal"`
		WordTranslated                                  string `json:"WordTranslated"`
		WordOriginalTranscription                       string `json:"WordOriginalTranscription"`
		WordOriginalPastSimpleSingular                  string `json:"WordOriginalPastSimpleSingular"`
		WordOriginalPastSimpleSingularTranscription     string `json:"WordOriginalPastSimpleSingularTranscription"`
		WordOriginalPastSimplePlural                    string `json:"WordOriginalPastSimplePlural"`
		WordOriginalPastSimplePluralTranscription       string `json:"WordOriginalPastSimplePluralTranscription"`
		WordOriginalPastParticipleSingular              string `json:"WordOriginalPastParticipleSingular"`
		WordOriginalPastParticipleSingularTranscription string `json:"WordOriginalPastParticipleSingularTranscription"`
		WordOriginalPastParticiplePlural                string `json:"WordOriginalPastParticiplePlural"`
		WordOriginalPastParticiplePluralTranscription   string `json:"WordOriginalPastParticiplePluralTranscription"`
		WordOriginalSynonyms                            string `json:"WordOriginalSynonyms"`
		WordOriginalPartOfSpeech                        string `json:"WordOriginalPartOfSpeech"`
		Rating                                          int    `json:"Rating"`
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

	Words = append(Words, DictionaryStruct(requestData))

	// Обновление файла данных (если есть) и другие операции, если необходимо
	// Открываем файл для записи
	jsonFile, err := os.OpenFile("EnglishForEveryone.json", os.O_WRONLY|os.O_TRUNC, 0644)
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
	query := r.URL.Query().Get("q") // Измените "query" на "q"
	results := searchWords(query)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

func searchWords(query string) []DictionaryStruct {
	results := []DictionaryStruct{}
	query = strings.ToLower(query)

	for _, word := range GoogleDict {
		if strings.HasPrefix(strings.ToLower(word.WordOriginal), query) { // Измените на strings.HasPrefix

			results = append(results, word)

		}

	}

	return results
}

func wordsSearch(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(GoogleDict)
}
