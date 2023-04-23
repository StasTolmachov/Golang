package main

import (
	"fmt"

	"github.com/antchfx/htmlquery"
)

func main() {
	//open url
	doc, err := htmlquery.LoadURL("https://minfin.com.ua/currency/converter/?from=eur&to=uah")
	if err != nil {
		panic(err)
	}
// find
	list, err := htmlquery.QueryAll(doc, "//td[@class='sc-5olu1-6 l04f60-2 lgnGvo']")
	if err != nil {
		panic(err)
	}

	var arrayExchange []string
	for _, n := range list {
		a := htmlquery.FindOne(n, "//*")
		if a != nil {
			// fmt.Printf("%d %s\n", i, htmlquery.InnerText(a))
			arrayExchange = append(arrayExchange, htmlquery.InnerText(a))
		} else {
			fmt.Println("пусто")
		}
	}

	fmt.Println(arrayExchange[2])

}
