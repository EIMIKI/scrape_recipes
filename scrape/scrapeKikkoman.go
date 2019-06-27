package scrape

import (
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func scrapeKikkoman(url string) ScrapedRecipe {
	doc, _ := goquery.NewDocument(url)

	scrapedRecipe := ScrapedRecipe{}
	scrapedRecipe.Title = doc.Find("h1").Text()
	scrapedRecipe.Amount = strings.Replace(doc.Find(".yield").Text(), "材料", "", -1)

	ingredientSelection := doc.Find(".ingredient")
	ingredient := Ingredient{}
	ingredientSelection.Each(func(index int, s *goquery.Selection) {
		ingredient.Name = s.Find(".name").Text()
		ingredient.Amount = s.Find(".amount").Text()
		scrapedRecipe.Ingredients = append(scrapedRecipe.Ingredients, ingredient)
	})

	directionSelection := doc.Find(".instruction")
	direction := Direction{}
	position := 0
	directionSelection.Each(func(index int, s *goquery.Selection) {
		position++
		direction.Position = strconv.Itoa(position)
		direction.Text = s.Text()
		scrapedRecipe.Directions = append(scrapedRecipe.Directions, direction)

	})

	scrapedRecipe.Err = ""

	return scrapedRecipe
}
