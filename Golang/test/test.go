package main

import (
	"fmt"
	"regexp"
)

type WordTranslation struct {
	English    string
	Transcript string
	Russian    string
}

func main() {
	input := "|zoo /zˈuː/|зоопарк|ABC /eibiːsiː/|азбука, алфавит|Abyssinia /əbisiniə/|Эфиопия|Adam's apple /ədɑːmzæpl/|адамово яблоко|Addis Ababa /ədaizæbəbə/|Аддис‐Абеба|Aden /ədn/|Аден|Afghanistan /æfgænistɑːn/|Афганистан|"
	re := regexp.MustCompile(`\|([^/]*)(/[^|]*\|)?([^|]*)`)
	matches := re.FindAllStringSubmatch(input, -1)

	for _, match := range matches {
		wordTranslation := WordTranslation{
			English:    match[1],
			Transcript: match[2],
			Russian:    match[3],
		}
		if len(wordTranslation.Transcript) > 0 {
			wordTranslation.Transcript = wordTranslation.Transcript[1 : len(wordTranslation.Transcript)-1]
		}
		fmt.Printf("English: %s, Transcript: %s, Russian: %s\n", wordTranslation.English, wordTranslation.Transcript, wordTranslation.Russian)
	}
}
