package scrape

import "strings"

// Ingredient 材料と量
type Ingredient struct {
	Name   string `json:"name"`
	Amount string `json:"amount"`
}

// Direction 手順の番号と内容
type Direction struct {
	Position string `json:"position"`
	Text     string `json:"text"`
}

// ScrapedRecipe スクレイピングされたレシピデータ
type ScrapedRecipe struct {
	Title       string       `json:"title"`
	Ingredients []Ingredient `json:"ingredients"`
	Directions  []Direction  `json:"directions"`
	Err         string       `json:"err"`
}

func ScrapeRecipe(url string) ScrapedRecipe {
	scrapedRecipe := ScrapedRecipe{Err: "Failed"}
	if strings.Contains(url, "cookpad.com") && strings.Contains(url, "recipe") {
		if !strings.Contains(url, "pro") {
			scrapedRecipe = scrapeCookpad(url)
		}
	}

	if strings.Contains(url, "park.ajinomoto.co.jp") && strings.Contains(url, "card") {
		scrapedRecipe = scrapeAjinomoto(url)
	}

	return scrapedRecipe
}
