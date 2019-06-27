package scrape

import "github.com/PuerkitoBio/goquery"

func scrapeCookpad(url string) ScrapedRecipe {
	doc, _ := goquery.NewDocument(url)

	scrapedRecipe := ScrapedRecipe{}
	scrapedRecipe.Title = doc.Find(".recipe-title").Text()
	scrapedRecipe.Amount = doc.Find(".yield").Text()

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

	scrapedRecipe.Err = ""

	return scrapedRecipe

}
