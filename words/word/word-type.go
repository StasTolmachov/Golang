package word

type WordStruct struct {
	ID             int
	Original       string
	Translated     string
	Rating         int      // разница правильных/неправильных ответов
	Status         string   // статус: новое/учится/выученно и тд
	WordDictionary []string // в каких словарях добавленно
}

var Dictionary []WordStruct
