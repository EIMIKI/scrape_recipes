package scrape

import "github.com/PuerkitoBio/goquery"

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
}

func ScrapeRecipe(url string) ScrapedRecipe {
	doc, _ := goquery.NewDocument(url)

	scrapedRecipe := ScrapedRecipe{}
	scrapedRecipe.Title = doc.Find(".recipe-title").Text()

	ingredientSelection := doc.Find(".ingredient")
	ingredient := Ingredient{}
	ingredientSelection.Each(func(index int, s *goquery.Selection) {
		ingredient.Name = s.Find(".name").Text()
		ingredient.Amount = s.Find(".amount").Text()
		scrapedRecipe.Ingredients = append(scrapedRecipe.Ingredients, ingredient)
	})

	directionSelection := doc.Find(".step")
	direction := Direction{}
	directionSelection.Each(func(index int, s *goquery.Selection) {
		direction.Position = s.Find("h3").Text()
		direction.Text = s.Find(".step_text").Text()
		scrapedRecipe.Directions = append(scrapedRecipe.Directions, direction)
	})

	return scrapedRecipe

}
