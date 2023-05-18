package main

type DictionaryStruct struct {
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

	WordOriginalCounterAttempts  int      // колличество попыток
	WordOriginalCounterIncorrect int      // колличество неправильных ответов
	WordOriginalCounterCorrect   int      // колличество правильных ответов
	WordOriginalDifficultyRating int      // разница правильных/неправильных ответов
	WordOriginalStatus           string   // статус: новое/учится/выученно и тд
	WordOriginalDictionary       []string // в каких словарях добавленно
}

func main() {

}
