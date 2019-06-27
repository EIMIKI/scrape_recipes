package scrape

import (
	"strings"
)

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
	Amount      string       `json:"amount"`
	Ingredients []Ingredient `json:"ingredients"`
	Directions  []Direction  `json:"directions"`
	Err         string       `json:"err"`
}

func cleanupStr(str string) string {
	str = strings.ReplaceAll(str, " ", "")
	str = strings.ReplaceAll(str, "\n", "")
	return str
}

// ScrapedRecipe以下の構造を変更するとここも追記しなきゃいけない...
func cleanupScrapedRecipe(scrapedRecipe ScrapedRecipe) ScrapedRecipe {
	creanedScrapedRecipe := ScrapedRecipe{}
	creanedScrapedRecipe.Title = cleanupStr(scrapedRecipe.Title)
	creanedScrapedRecipe.Amount = cleanupStr(scrapedRecipe.Amount)

	for _, ingredient := range scrapedRecipe.Ingredients {
		ingredient.Amount = cleanupStr(ingredient.Amount)
		ingredient.Name = cleanupStr(ingredient.Name)

		creanedScrapedRecipe.Ingredients = append(creanedScrapedRecipe.Ingredients, ingredient)
	}
	for _, direction := range scrapedRecipe.Directions {
		direction.Position = cleanupStr(direction.Position)
		direction.Text = cleanupStr(direction.Text)

		creanedScrapedRecipe.Directions = append(creanedScrapedRecipe.Directions, direction)
	}

	return creanedScrapedRecipe
}

// ScrapeRecipe urlに合わせてスクレイピングを行う。
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

	if strings.Contains(url, "kikkoman.co.jp") && strings.Contains(url, "recipe") {
		scrapedRecipe = scrapeKikkoman(url)
	}

	return cleanupScrapedRecipe(scrapedRecipe)
}
